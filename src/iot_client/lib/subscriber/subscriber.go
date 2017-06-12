package subscriber

import (
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/aws"
	"fmt"
	"github.com/aws/aws-sdk-go/service/sqs"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"encoding/json"
	"iot/lib/utils"
)


type Message struct {
	Msg	map[string]interface{}
}

type Data struct {
	Type  	string
	Value	string
}
func Sub(){
	sess := session.New(&aws.Config{
		Region:      aws.String("us-west-2"),
		Credentials: credentials.NewStaticCredentials("AKIAJ2PQZJE7JD4GUQYA","AQjrMM/qVvzr/0VvPI88uztKZTXazr0mP/tKBJnu",""),
		MaxRetries:  aws.Int(5),
	})

	q := sqs.New(sess)
	messages, err := q.ReceiveMessage(&sqs.ReceiveMessageInput{
		QueueUrl:aws.String("https://sqs.us-west-2.amazonaws.com/454446851063/havells-v1"),
		MaxNumberOfMessages: aws.Int64(10),
	})
	if err!=nil{
		fmt.Println(err.Error())
	}
	for _,v:=range messages.Messages{
		ma := make(map[string]interface{})
		json.Unmarshal([]byte(*(v.Body)), &ma)


		if ma["MessageAttributes"] == nil{
			continue
		}
		v:=ma["MessageAttributes"].(map[string]interface{})
		resp := map[string]Data{}
		for k,vv:=range v{
			tmp := Data{}
			convv := vv.(map[string]interface{})
			tmp.Value = utils.ToStr(convv["Value"])
			resp[k] = tmp
		}
		fmt.Println("ans=>",resp)
	}
	fmt.Println(messages)

}