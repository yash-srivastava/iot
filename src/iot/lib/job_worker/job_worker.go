package job_worker

import (
	"github.com/StabbyCutyou/buffstreams"
	"iot/lib/sgu_utils"
	"iot/lib/formatter"
	"iot/lib/sender"
	"github.com/revel/revel"
	"iot/conf"
	"github.com/jrallison/go-workers"
	"iot/lib/utils"
)

func ProcessPacket(message *workers.Msg ) {

	params,_ := message.Map()
	msg := params["args"].(map[string]interface{})

	action := msg["action"]

	if action == "parse_sgu_packets"{
		client := buffstreams.Client{}
		err := formatter.GetStructFromInterface(msg["client"], &client)

		if err!=nil{
			revel.ERROR.Println(err)
		}
		sgu_utils.ParseInputPackets(&client)
	}else if action == "send_response_packets"{
		incoming := conf.Incoming{}

		packet_type := utils.ToInt(msg["packet_type"])

		err := formatter.GetStructFromInterface(msg["incoming"], &incoming)
		if err!=nil{
			revel.ERROR.Println(err)
		}
		sender.SendResponsePacket(packet_type, incoming)
	}else if action == "send_3000"{
		sender.SendServerPacket(0x3000, msg["params"])
	}else if action == "send_8000"{
		sender.SendServerPacket(0x8000, msg["params"])
	}
}
