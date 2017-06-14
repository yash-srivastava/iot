package subscriber

import (
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/sqs"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"encoding/json"
	"github.com/revel/revel"
	"github.com/benmanns/goworker"
)

func SqsSubscribe(){
	sess := session.New(&aws.Config{
		Region:      aws.String("us-west-2"),
		Credentials: credentials.NewStaticCredentials("AKIAIWOKS4IKMGHPCMQQ","3CFxcS8760bXMUPC5/Tk7dyY69Kv9+aTT4ZVAHCY",""),
		MaxRetries:  aws.Int(5),
	})

	q := sqs.New(sess)

	receive_params := &sqs.ReceiveMessageInput{
		QueueUrl:            aws.String("https://sqs.us-west-2.amazonaws.com/454446851063/havells-v1"),
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

				params := make([]interface{}, 2)
				params[0] = "save_in_db"
				params[1] = ma["MessageAttributes"]

				payload := goworker.Payload{"subscribers", params}
				job := goworker.Job{"packet_subscriber_queue", payload}
				goworker.Enqueue(&job)
				revel.INFO.Println("Job Enqueued")

				// Delete Message
				delete_params := &sqs.DeleteMessageInput{
					QueueUrl:      aws.String("https://sqs.us-west-2.amazonaws.com/454446851063/havells-v1"),
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