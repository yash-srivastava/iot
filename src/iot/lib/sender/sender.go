package sender

import (
	"iot/lib/parser"
	"strconv"
	"github.com/StabbyCutyou/buffstreams"
	"encoding/binary"
	"iot/lib/formatter"
	"github.com/revel/revel"
)

func SendResponsePacket(pack_type int, incoming parser.Incoming){
	packet_config := parser.GetSguResponsePacket()

	packet_description := packet_config.Response_packets

	delim := packet_config.Delim

	packet_type := int(pack_type)

	length := packet_description[packet_type].Length

	var response []byte
	response = make([]byte, length)

	sgu_id := incoming.SguId
	seq_no := incoming.SeqNo


	response = convertToByteArray(uint64(delim),1)
	response = append(response, convertToByteArray(uint64(length),2)...)
	response = append(response, convertToByteArray(sgu_id,6)...)
	response = append(response, convertToByteArray(uint64(seq_no),4)...)
	response = append(response, convertToByteArray(uint64(packet_type),2)...)




	for key,_ :=range packet_description[packet_type].Response_parameters{
		if key=="status"{
			response = append(response, byte(1))
		}
	}

	client,_ := parser.SGU_TCP_CONNECTION.Get(strconv.FormatUint(sgu_id,10))
	conv,ok := client.(*buffstreams.TCPConn)
	if !ok{
		revel.WARN.Println("Invalid Connection")
	}

	revel.INFO.Println("Sending Packet:","packet_type=>",packet_type,"(",formatter.ToHex(packet_type),")","sgu_id=>",sgu_id,"(",formatter.ToHex(sgu_id),")")
	_,e:=conv.Write(response)
	if e!=nil{
		revel.ERROR.Print(e.Error())
	}

}

func convertToByteArray (val uint64, len int)[]byte{
	value := uint64(val)
	bs := make([]byte, 8)
	binary.BigEndian.PutUint64(bs,value)
	return bs[8-len:]
}
