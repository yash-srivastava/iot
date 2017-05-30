package parser

import (
	"iot/lib/utils"
	"github.com/revel/revel"
	"iot/lib/formatter"
)

func HandlePackets(packet_type int, input map[string]uint64){
	switch packet_type {

	case 0x0003:{
		var scuids []uint64
		iterate := input["iterate"]
		sgu_id := input["incoming_sgu_id"]
		for i := 0; i < int(iterate); i++ {
			if i == 0 {
				scuids = append(scuids, input["scuid"])
			} else {
				scuids = append(scuids, input["scuid_" + utils.ToStr(i)])
			}
		}
		tmp := Scu{}
		tmp.ScuIds = scuids
		SGU_SCU_LIST.Set(utils.ToStr(sgu_id), tmp)
		revel.INFO.Println("Following SCUs Found:", scuids, " For SGU:", formatter.Prettify(sgu_id))
	}

	case 0x0004:{

		var newscuids []uint64
		sgu_id := input["incoming_sgu_id"]

		scu,_ := SGU_SCU_LIST.Get(utils.ToStr(sgu_id))

		incoming := Scu{}
		err := formatter.GetStructFromInterface(scu, &incoming)
		if err!=nil{
			revel.ERROR.Println(err)
		}

		scu_ids := incoming.ScuIds

		rem_scu_id := input["scuid"]
		for i := 0; i < len(scu_ids); i++ {
			if rem_scu_id == scu_ids[i]{
				continue
			}
			newscuids = append(newscuids,scu_ids[i])
		}

		incoming.ScuIds = newscuids

		SGU_SCU_LIST.Set(utils.ToStr(sgu_id), incoming)
		revel.INFO.Println("Following SCU Removed:", formatter.Prettify(rem_scu_id), " For SGU:", formatter.Prettify(sgu_id))
	}
	}
}
