package sender

import (
	"iot/lib/parser"
	"fmt"
	"strconv"
	"github.com/StabbyCutyou/buffstreams"
	"encoding/binary"
)

func SendResponsePacket(pack_type int, incoming parser.Incoming){
	packet_config := parser.GetSguResponsePacket()

	packet_description := packet_config.Response_packets
	fmt.Print("packet_des$$$$$=>",packet_description)

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
			fmt.Println("Appending=>>>",key)
			response = append(response, byte(1))
		}
	}

	client,_ := parser.SGU_TCP_CONNECTION.Get(strconv.FormatUint(sgu_id,10))
	conv,ok := client.(*buffstreams.TCPConn)
	if !ok{
		fmt.Print("Invalid Connection")
	}
	wr,e:=conv.Write(response)
	fmt.Print("line",wr)
	if e!=nil{
		fmt.Print(e.Error())
	}

}

func convertToByteArray (val uint64, len int)[]byte{
	value := uint64(val)
	bs := make([]byte, 8)
	binary.BigEndian.PutUint64(bs,value)
	return bs[8-len:]
}
