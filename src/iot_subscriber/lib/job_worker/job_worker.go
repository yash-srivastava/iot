package job_worker

import (
	"errors"
	"iot/lib/formatter"
	"iot/lib/sender"
	"encoding/json"
	"github.com/revel/revel"
	"iot/conf"
	"iot_subscriber/lib/message_handler"
)

func ProcessData(task string, args ...interface{}) error {
	name, ok := args[0].(string)
	if !ok {
		return errors.New("Invalid Worker")
	}
	if name == "save_in_db"{
		message_handler.Handle(args[1])
	}else if name == "send_response_packets"{
		incoming := conf.Incoming{}

		packet_type := args[1].(json.Number)

		err := formatter.GetStructFromInterface(args[2], &incoming)
		if err!=nil{
			revel.ERROR.Println(err)
			return err
		}
		pack_type,_ := (packet_type.Int64())
		sender.SendResponsePacket(int(pack_type), incoming)
	}else if name == "send_3000"{
		sender.SendServerPacket(0x3000, args[1])
	}
	return nil
}
