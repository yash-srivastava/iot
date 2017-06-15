package job_worker

import (
	"iot_subscriber/lib/message_handler"
	"github.com/jrallison/go-workers"
)

func ProcessData(message *workers.Msg) {

	params,_ := message.Map()
	msg := params["args"].(map[string]interface{})

	action := msg["action"]

	if action == "save_in_db"{
		message_handler.Handle(msg["data"])
	}
}
