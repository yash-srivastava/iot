package parser

import (
	"iot/lib/utils"
	"github.com/revel/revel"
	"iot/lib/formatter"
	"strconv"
	"math"
	"encoding/binary"
	"iot/conf"
	"iot/lib/sender"
)

func HandlePackets(packet_type int, input map[string]interface{}){
	switch packet_type {

	case 0x0003:{
		var scuids []uint64
		iterate := utils.ToUint64(input["iterate"])
		sgu_id := utils.ToUint64(input["incoming_sgu_id"])
		for i := 0; i < int(iterate); i++ {
			if i == 0 {
				scuids = append(scuids, utils.ToUint64(input["scuid"]))
			} else {
				scuids = append(scuids, utils.ToUint64(input["scuid_" + utils.ToStr(i)]))
			}
		}
		tmp := conf.Scu{}
		tmp.ScuIds = scuids
		conf.SGU_SCU_LIST.Set(utils.ToStr(sgu_id), tmp)
		revel.INFO.Println("Following SCUs Found:", scuids, " For SGU:", formatter.Prettify(sgu_id))
	}

	case 0x0004:{

		var newscuids []uint64
		sgu_id := utils.ToUint64(input["incoming_sgu_id"])

		scu,_ := conf.SGU_SCU_LIST.Get(utils.ToStr(sgu_id))

		incoming := conf.Scu{}
		err := formatter.GetStructFromInterface(scu, &incoming)
		if err!=nil{
			revel.ERROR.Println(err)
		}

		scu_ids := incoming.ScuIds

		rem_scu_id := utils.ToUint64(input["scuid"])
		for i := 0; i < len(scu_ids); i++ {
			if rem_scu_id == scu_ids[i]{
				continue
			}
			newscuids = append(newscuids,scu_ids[i])
		}

		incoming.ScuIds = newscuids

		conf.SGU_SCU_LIST.Set(utils.ToStr(sgu_id), incoming)
		revel.INFO.Println("Following SCU Removed:", formatter.Prettify(rem_scu_id), " For SGU:", formatter.Prettify(sgu_id))
	}

	case 0x0005:{
		sgu_id := input["incoming_sgu_id"]

		scu,_ := conf.SGU_SCU_LIST.Get(utils.ToStr(sgu_id))
		incoming := conf.Scu{}
		err := formatter.GetStructFromInterface(scu, &incoming)
		if err!=nil{
			revel.ERROR.Println(err)
		}
		incoming.ScuIds = append(incoming.ScuIds, utils.ToUint64(input["scuid"]))

		conf.SGU_SCU_LIST.Set(utils.ToStr(sgu_id), incoming)
		revel.INFO.Println("Following SCU Added:", formatter.Prettify(input["scuid"]), " For SGU:", formatter.Prettify(sgu_id))
	}

	case 0x3001:{
		status := utils.ToInt(input["status"])

		if status == 1{
			sgu_id := utils.ToUint64(input["incoming_sgu_id"])
			scu_id := utils.ToUint64(input["scuid"])
			get_set := utils.ToInt(input["get_set"])
			packet := sender.Packet_3000{}
			packet.SguId = sgu_id
			packet.ScuId = scu_id
			packet.GetSet = get_set
			conf.Retry_3000.Set(sender.Get300Hash(packet),false)
		}
	}
	}
}

func HandleCustomPackets(packet_type int, packet_data []byte, start_from int) map[string]interface{}{
	packet_des := conf.CUSTOM_PACKET_CONFIG.Packets
	response := make(map[string]interface{})

	for offset,val :=range packet_des[packet_type].Parameters{
		off,_ := strconv.Atoi(offset)
		len,_ := strconv.Atoi(val.Length)

		swappedstr := make([]byte, len)
		for i:=len/2;i<len;i++{
			swappedstr = append(swappedstr, packet_data[start_from+off+i])
		}
		for i:=0;i<len/2;i++{
			swappedstr = append(swappedstr, packet_data[start_from+off+i])
		}

		value:= binary.BigEndian.Uint64(swappedstr)
		data := math.Float32frombits(uint32(value))
		response[val.Name] = data
	}
	return response
}
