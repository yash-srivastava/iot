package subscriber

import (
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/sqs"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"encoding/json"
	"github.com/revel/revel"
	"github.com/jrallison/go-workers"
)

func SqsSubscribe(){
	sess := session.New(&aws.Config{
		Region:      aws.String(revel.Config.StringDefault("aws_region", "aws_region")),
		Credentials: credentials.NewStaticCredentials(revel.Config.StringDefault("aws_access_key", "aws_access_key"),revel.Config.StringDefault("aws_secret_key", "aws_secret_key"),""),
		MaxRetries:  aws.Int(5),
	})

	q := sqs.New(sess)

	receive_params := &sqs.ReceiveMessageInput{
		QueueUrl:            aws.String(revel.Config.StringDefault("aws_queue_url", "aws_queue_url")),
		MaxNumberOfMessages: aws.Int64(10),
		VisibilityTimeout:   aws.Int64(30),
		WaitTimeSeconds:     aws.Int64(20),
	}
	revel.INFO.Println("Listening For Messages")
	for ; ; {
		messages, err := q.ReceiveMessage(receive_params)
		if err!=nil{
			revel.ERROR.Println(err.Error())
			continue
		}
		if len(messages.Messages) > 0{

			//Read Messages
			for _,msg:=range messages.Messages{

				ma := make(map[string]interface{})
				json.Unmarshal([]byte(*(msg.Body)), &ma)
				if ma["MessageAttributes"] == nil{
					continue
				}
				/*msg_attr := ma["MessageAttributes"].(map[string]interface{})
				resp := map[string]Data{}
				for k,v:=range msg_attr{
					tmp := Data{}
					converted := v.(map[string]interface{})
					tmp.Value = utils.ToStr(converted["Value"])
					resp[k] = tmp
				}*/

				params := make(map[string]interface{})
				params["action"] = "save_in_db"
				params["data"] = ma["MessageAttributes"]

				workers.Enqueue("packet_subscriber_queue", "save_in_db", params)
				revel.INFO.Println("Job Enqueued")

				// Delete Message
				delete_params := &sqs.DeleteMessageInput{
					QueueUrl:      aws.String(revel.Config.StringDefault("aws_queue_url", "aws_queue_url")),
					ReceiptHandle: msg.ReceiptHandle,

				}
				_, err := q.DeleteMessage(delete_params)
				if err != nil {
					revel.ERROR.Println(err.Error())
				}
				revel.INFO.Println("Message ID:",*msg.MessageId,"has been deleted")
			}
		}

	}


}