package publisher

import (
	"github.com/aws/aws-sdk-go/service/sns"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/aws"
	"iot/lib/utils"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/revel/revel"
)

func Pub(msg map[string]interface{}){

	result := make(map[string]*sns.MessageAttributeValue)
	for k,v:=range(msg){
		tmp := sns.MessageAttributeValue{}
		tmp.BinaryValue = []byte(utils.ToStr(v))
		//tmp.StringValue = aws.String(utils.ToStr(v))
		tmp.DataType = aws.String("Binary")
		result[k] = &tmp
	}

	sess := session.New(&aws.Config{
		Region:      aws.String(revel.Config.StringDefault("aws_region", "aws_region")),
		Credentials: credentials.NewStaticCredentials(revel.Config.StringDefault("aws_access_key", "aws_access_key"),revel.Config.StringDefault("aws_secret_key", "aws_secret_key"),""),
		MaxRetries:  aws.Int(5),
	})

	svc := sns.New(sess)

	// params will be sent to the publish call included here is the bare minimum params to send a message.
	params := &sns.PublishInput{
		MessageAttributes: result,
		Message: aws.String("Called"), // This is the message itself (can be XML / JSON / Text - anything you want)
		TopicArn: aws.String(revel.Config.StringDefault("aws_topic_arn", "aws_topic_arn")),  //Get this from the Topic in the AWS console.
	}

	resp, err := svc.Publish(params)   //Call to puclish the message

	if err != nil {                    //Check for errors
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		revel.ERROR.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	revel.INFO.Println(resp)
}