package parser

import (
	"strconv"
	"encoding/binary"
	"strings"
	"github.com/StabbyCutyou/buffstreams"
	"github.com/benmanns/goworker"
	"iot/lib/formatter"
	"github.com/revel/revel"
	"iot/lib/utils"
	"iot/conf"
)


func Wrap(conn *buffstreams.Client)map[string]interface{} {


	incoming := conf.Incoming{}
	var result map[string]interface{}
	result = make(map[string]interface{})

	packet_data := conn.Data

	delim := int(conf.PACKET_CONFIG.Delim)

	byte_arr :=preparePacket(packet_data[0:1])

	val:= int(binary.BigEndian.Uint32(byte_arr))
	if val != delim{
		revel.WARN.Println("Failed To Match Start Delim=>",val," Delim=>",delim)
		return nil
	}

	byte_arr = preparePacket(packet_data[1:3])
	packet_length := int(binary.BigEndian.Uint32(byte_arr))


	byte_arr = preparePacket8(packet_data[3:9])
	sgu_id := (binary.BigEndian.Uint64(byte_arr))

	incoming.SguId = sgu_id

	byte_arr = preparePacket8(packet_data[9:23])
	timestamp:= int64(binary.BigEndian.Uint64(byte_arr))


	byte_arr = preparePacket8(packet_data[23:27])
	seq_no := int64(binary.BigEndian.Uint64(preparePacket(byte_arr)))
	incoming.SeqNo = seq_no

	byte_arr = preparePacket(packet_data[27:29])
	packet_type := int(binary.BigEndian.Uint32([]byte(byte_arr)))

	packet_description := conf.PACKET_CONFIG.Packets

	revel.WARN.Println("Packet Received:","packet_type=>",formatter.Prettify(packet_type),"| description=>",packet_description[packet_type].Description,"| packet_length=>",packet_length,"| sgu_id=>",formatter.Prettify(sgu_id),"| seq_no=>",formatter.Prettify(seq_no))

	result["incoming_sgu_id"] = utils.ToUint64(sgu_id)
	result["incoming_timestamp"] = utils.ToUint64(timestamp)

	client,_ := buffstreams.TcpClients.Get(conn.Address)
	conf.SGU_TCP_CONNECTION.Set(strconv.FormatUint(sgu_id,10),client)


	var repeat_parameter conf.Packets
	repeat_parameter.Parameters = make(map[string]conf.Parameters)
	last_offset := 0
	iterate := 0

	for offset,val :=range packet_description[packet_type].Parameters{

		splitted_arr := strings.Split(offset,"_")
		splitted_len := GetLength(splitted_arr)
		off,_ := strconv.Atoi(splitted_arr[splitted_len-1])
		len,_ := strconv.Atoi(val.Length)
		if strings.Contains(offset,"repeat_"){
			//save for repeat
			ma := val
			repeat_parameter.Parameters[offset] = ma
		}

		if val.Out_type == "int64"{
			byte_arr = preparePacket8(packet_data[off:off+len])
			result[val.Name] = (binary.BigEndian.Uint64([]byte(byte_arr)))
		}else{
			byte_arr = preparePacket(packet_data[off:off+len])
			result[val.Name] = uint64(binary.BigEndian.Uint32([]byte(byte_arr)))
		}

		last_offset = off+len

		if strings.Contains(val.Name, "num_"){
			iterate = utils.ToInt(result[val.Name])
		}

		if strings.Contains(offset, "length_"){
			custom_response := HandleCustomPackets(packet_type, packet_data,off+len)
			for ck,cv:=range custom_response {
				result[ck] = cv
			}
			last_offset += utils.ToInt(result[val.Name])
		}
	}


	result["iterate"] = utils.ToUint64(iterate)

	for i:=0;i<iterate-1;i++{
		suffix := "_"+strconv.Itoa(i+1)
		for off,v:=range repeat_parameter.Parameters{
			len,_ := strconv.Atoi(v.Length)
			if v.Out_type == "int64"{
				byte_arr = preparePacket8(packet_data[last_offset:last_offset+len])
				result[v.Name+suffix] = (binary.BigEndian.Uint64([]byte(byte_arr)))
			}else{
				byte_arr = preparePacket(packet_data[last_offset:last_offset+len])
				result[v.Name+suffix] = uint64(binary.BigEndian.Uint32([]byte(byte_arr)))
			}
			if strings.Contains(off, "length_") {
				custom_response := HandleCustomPackets(packet_type, packet_data,last_offset+len)
				for ck,cv:=range custom_response {
					result[ck+suffix] = cv
				}
				last_offset += utils.ToInt(result[v.Name+suffix])
			}
			last_offset += len
		}
	}


	if packet_description[packet_type].Response_packet != -1{

		params := make([]interface{}, 3)
		params[0] = "send_response_packets"
		params[1] = packet_description[packet_type].Response_packet
		params[2] = incoming

		payload := goworker.Payload{"packets", params}
		job := goworker.Job{"packet_queue", payload}
		goworker.Enqueue(&job)

		revel.INFO.Println("Response Packet:", formatter.Prettify(params[1]), "Enqueued")
	}

	revel.WARN.Println(result)
	conf.Producer.SendMessage("important",result)
	HandlePackets(packet_type, result)
	return nil

}

func readPacket(arr []string, i int, j int) string{
	result := ""
	for ;i<=j;i++ {
		result+=arr[i]
	}
	return result
}

func GetLength(arr []string)int{
	return len(arr)
}

func GetStringLength(arr string)int{
	return len(arr)
}
func preparePacket(arr []byte) []byte{
	var result []byte

	tmp := byte(0)
	len := len(arr)
	for i:=len;i<4;i++{
		result=append(result,tmp)
	}
	for k,_:=range arr{
		result = append(result, arr[k])
	}
	return (result)
}

func preparePacket8(arr []byte) []byte{
	var result []byte

	tmp := byte(0)
	len := len(arr)
	for i:=len;i<8;i++{
		result=append(result,tmp)
	}
	for k,_:=range arr{
		result = append(result, arr[k])
	}
	return result
}