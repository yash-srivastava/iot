package sender

import (
	"strconv"
	"github.com/StabbyCutyou/buffstreams"
	"encoding/binary"
	"iot/lib/formatter"
	"github.com/revel/revel"
	"iot/conf"
	"iot/lib/utils"
)

func SendResponsePacket(pack_type int, incoming conf.Incoming){

	packet_description := conf.RESPONSE_PACKET_CONFIG.Response_packets

	delim := conf.RESPONSE_PACKET_CONFIG.Delim

	packet_type := int(pack_type)

	length := packet_description[packet_type].Length

	sgu_id := incoming.SguId
	seq_no := incoming.SeqNo

	conn := sguConnection(sgu_id)

	if conn==nil{
		return
	}

	response := AddCommonParameters(byte(delim),sgu_id,uint64(seq_no),length,packet_type)

	for key,_ :=range packet_description[packet_type].Response_parameters{
		if key=="status"{
			response = append(response, byte(1))
		}
	}



	revel.INFO.Println("Sending Packet:","packet_type=>",formatter.Prettify(packet_type),"| description=>",packet_description[packet_type].Description,"| sgu_id=>",formatter.Prettify(sgu_id),"| seq_no=>",formatter.Prettify(seq_no))
	_,e:=conn.Write(response)
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

func sguConnection(sgu_id uint64) (*buffstreams.TCPConn){
	client,_ := conf.SGU_TCP_CONNECTION.Get(strconv.FormatUint(sgu_id,10))
	conv,ok := client.(*buffstreams.TCPConn)
	if !ok{
		revel.WARN.Println("Invalid Connection, SGU=>", formatter.Prettify(sgu_id), "Not Connected")
		return nil
	}
	return conv
}

func scuPresent(sgu_id uint64, scu_id uint64) bool{
	scu,_ := conf.SGU_SCU_LIST.Get(utils.ToStr(sgu_id))
	incoming := conf.Scu{}
	err := formatter.GetStructFromInterface(scu, &incoming)
	if err !=nil {
		revel.ERROR.Println(err)
		return false
	}
	for i:=0;i<len(incoming.ScuIds);i++{
		if incoming.ScuIds[i] == scu_id{
			return true
		}
	}
	return false
}


func add_byte_array_to_response(offset int, length int,val []byte, data []byte) []byte{
	for i := offset;i <  offset + length ; i++{
		data[i] = val [i-offset]
	}
	return data
}