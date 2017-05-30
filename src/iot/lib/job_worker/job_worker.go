package job_worker

import (
	"github.com/StabbyCutyou/buffstreams"
	"iot/lib/sgu_utils"
	"errors"
	"iot/lib/formatter"
	"iot/lib/parser"
	"iot/lib/sender"
	"encoding/json"
	"github.com/revel/revel"
)

func ProcessPacket(task string, args ...interface{}) error {
	name, ok := args[0].(string)
	if !ok {
		return errors.New("Invalid Worker")
	}
	if name == "parse_sgu_packets"{
		client := buffstreams.Client{}
		err := formatter.GetStructFromInterface(args[1], &client)

		if err!=nil{
			revel.ERROR.Println(err)
		}
		sgu_utils.ParseInputPackets(&client)
	}else if name == "send_response_packets"{
		incoming := parser.Incoming{}

		packet_type := args[1].(json.Number)

		err := formatter.GetStructFromInterface(args[2], &incoming)
		if err!=nil{
			revel.ERROR.Println(err)
		}
		pack_type,_ := (packet_type.Int64())
		sender.SendResponsePacket(int(pack_type), incoming)
	}
	return nil
}
