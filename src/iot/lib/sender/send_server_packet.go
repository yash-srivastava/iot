package sender

import (
	"iot/lib/formatter"
	"github.com/revel/revel"
	"iot/conf"
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
			if k == "scuid"{
				val := convertToByteArray(packet.ScuId, 8)
				response = add_byte_array_to_response(v.Offset, v.Length,val, response)
			}else if k == "get_set"{
				val := convertToByteArray(uint64(packet.GetSet), 1)
				response = add_byte_array_to_response(v.Offset, v.Length,val, response)
			}
			if packet.GetSet !=0 {
				if k == "pwm"{
					val := convertToByteArray(uint64(packet.Pwm), 1)
					response = add_byte_array_to_response(v.Offset, v.Length,val, response)
				}else if k == "op1"{
					val := convertToByteArray(uint64(packet.Op1), 1)
					response = add_byte_array_to_response(v.Offset, v.Length,val, response)
				}else if k == "op2"{
					val := convertToByteArray(uint64(packet.Op2), 1)
					response = add_byte_array_to_response(v.Offset, v.Length,val, response)
				}else if k == "op3"{
					val := convertToByteArray(uint64(packet.Op3), 1)
					response = add_byte_array_to_response(v.Offset, v.Length,val, response)
				}else if k == "op4"{
					val := convertToByteArray(uint64(packet.Op4), 1)
					response = add_byte_array_to_response(v.Offset, v.Length,val, response)
				}
			}


		}

		revel.INFO.Println("Resp=>",response)
		revel.INFO.Println("Sending Packet:","packet_type=>",formatter.Prettify(packet_type),"| description=>",packet_description[packet_type].Description,"| sgu_id=>",formatter.Prettify(sgu_id))
		_,e:=conn.Write(response)
		if e!=nil{
			revel.ERROR.Print(e.Error())
		}

	}

	}
}
