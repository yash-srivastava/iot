package sender

import (
	"iot/lib/formatter"
	"github.com/revel/revel"
	"iot/conf"
	"gopkg.in/oleiade/reflections.v1"
	"iot/lib/utils"
)

func SendServerPacket(packet_type int, params interface{}){
	switch packet_type {

	case 0x3000:{
		packet := Packet_3000{}
		err := formatter.GetStructFromInterface(params, &packet)

		if err!=nil{
			revel.ERROR.Println(err)
			return
		}

		packet_description := conf.SERVER_PACKET_CONFIG.Response_packets

		delim := conf.SERVER_PACKET_CONFIG.Delim

		packet_type := int(packet_type)

		length := packet_description[packet_type].Length

		if packet.GetSet == 0{
			length -=5
		}

		seq_no := 0
		sgu_id := packet.SguId

		conn := sguConnection(sgu_id)

		if conn==nil{
			return
		}

		if !scuPresent(sgu_id, packet.ScuId) {
			revel.WARN.Println("SCU=>", formatter.Prettify(packet.ScuId),"not connected to specified SGU=>", formatter.Prettify(sgu_id))
			return
		}

		response := AddCommonParameters(byte(delim),sgu_id,uint64(seq_no),length,packet_type)

		for k,v := range packet_description[packet_type].Response_parameters {
			val,err := reflections.GetField(packet, k)
			if err != nil {
				revel.INFO.Println(k,"not present for packet_type", utils.ToStr(packet_type))
				panic(k+" not present for packet_type " + utils.ToStr(packet_type))
			}

			uint_val := utils.ToUint64(val)
			byte_val := convertToByteArray(uint_val, v.Length)

			if k == "Pwm" || k == "Op1" || k == "Op2" || k == "Op3" || k == "Op4"{
				if packet.GetSet !=0{
					response = add_byte_array_to_response(v.Offset, v.Length,byte_val, response)
				}
			}else {
				response = add_byte_array_to_response(v.Offset, v.Length,byte_val, response)
			}

		}
		revel.INFO.Println("Sending Packet:","packet_type=>",formatter.Prettify(packet_type),"| description=>",packet_description[packet_type].Description,"| sgu_id=>",formatter.Prettify(sgu_id))
		revel.WARN.Println("Packet:",packet)
		_,e:=conn.Write(response)
		if e!=nil{
			revel.ERROR.Print(e.Error())
		}

	}
	case 0x8000:{
		packet := Packet_8000{}
		err := formatter.GetStructFromInterface(params, &packet)

		if err!=nil{
			revel.ERROR.Println(err)
			return
		}

		packet_description := conf.SERVER_PACKET_CONFIG.Response_packets

		delim := conf.SERVER_PACKET_CONFIG.Delim

		packet_type := int(packet_type)

		length := packet_description[packet_type].Length

		exprArr := []byte(packet.Expression)

		if packet.GetSet == 1{
			length += len(exprArr)
		}

		seq_no := 0
		sgu_id := packet.SguId

		conn := sguConnection(sgu_id)

		if conn==nil{
			return
		}

		if !scuPresent(sgu_id, packet.ScuId) {
			revel.WARN.Println("SCU=>", formatter.Prettify(packet.ScuId),"not connected to specified SGU=>", formatter.Prettify(sgu_id))
			return
		}

		response := AddCommonParameters(byte(delim),sgu_id,uint64(seq_no),length,packet_type)

		for k,v := range packet_description[packet_type].Response_parameters {

			if k == "Expr"{
				if packet.GetSet !=0{
					response = add_byte_array_to_response(v.Offset, len(exprArr), exprArr, response)
				}
			}else{
				val,err := reflections.GetField(packet, k)
				if err != nil {
					revel.INFO.Println(k,"not present for packet_type", utils.ToStr(packet_type))
					panic(k+" not present for packet_type " + utils.ToStr(packet_type))
				}

				uint_val := utils.ToUint64(val)
				byte_val := convertToByteArray(uint_val, v.Length)
				response = add_byte_array_to_response(v.Offset, v.Length,byte_val, response)
			}
		}
		revel.INFO.Println("Sending Packet:","packet_type=>",formatter.Prettify(packet_type),"| description=>",packet_description[packet_type].Description,"| sgu_id=>",formatter.Prettify(sgu_id))
		revel.WARN.Println("Packet:",packet)
		_,e:=conn.Write(response)
		if e!=nil{
			revel.ERROR.Print(e.Error())
		}

	}

	}
}
