package job_worker

import (
	"github.com/StabbyCutyou/buffstreams"
	"iot/lib/sgu_utils"
	"errors"
	"fmt"
	"iot/lib/formatter"
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
			fmt.Print(err)
		}
		sgu_utils.ParseInputPackets(&client)
	}
	return nil
}
