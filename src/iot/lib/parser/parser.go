package parser

import (
	//"encoding/json"
	"strconv"
	//"log"
	"encoding/binary"
	"strings"
	"github.com/StabbyCutyou/buffstreams"
	"github.com/orcaman/concurrent-map"
	"github.com/benmanns/goworker"
	"iot/lib/formatter"
	"github.com/revel/revel"
)


var SGU_TCP_CONNECTION cmap.ConcurrentMap

func Wrap(conn *buffstreams.Client)map[string]interface{} {


	incoming := Incoming{}
	var result map[string]int64
	result = make(map[string]int64)

	packet_data := conn.Data

	packet_config := GetSguPacket()

	delim := int(packet_config.Delim)

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

	//byte_arr = preparePacket8(packet_data[9:23])
	//timestamp:= int64(binary.BigEndian.Uint64(byte_arr))


	byte_arr = preparePacket8(packet_data[23:27])
	seq_no := int64(binary.BigEndian.Uint64(preparePacket(byte_arr)))
	incoming.SeqNo = seq_no

	byte_arr = preparePacket(packet_data[27:29])
	packet_type := int(binary.BigEndian.Uint32([]byte(byte_arr)))

	packet_description := packet_config.Packets

	revel.INFO.Println("Packet Received:","packet_type=>",packet_type,"(",formatter.ToHex(packet_type),")","packet_length=>",packet_length,"sgu_id=>",sgu_id,"(",formatter.ToHex(sgu_id),")")

	client,_ := buffstreams.TcpClients.Get(conn.Address)
	SGU_TCP_CONNECTION.Set(strconv.FormatUint(sgu_id,10),client)


	var repeat_parameter []Parameters
	last_offset := 0
	iterate := 0
		for offset,val :=range packet_description[packet_type].Parameters{
			off := 0
			len :=0
			if strings.Contains(offset,"repeat_"){
				off,_ = strconv.Atoi(strings.Split(offset,"repeat_")[1])
				len,_ = strconv.Atoi(val.Length)

				//save for repeat
				ma := val
				repeat_parameter = append(repeat_parameter, ma)
			}else {
				off,_ = strconv.Atoi(offset)
				len,_ = strconv.Atoi(val.Length)
			}

			if val.Out_type == "int64"{
				byte_arr = preparePacket8(packet_data[off:off+len])
				result[val.Name] = int64(binary.BigEndian.Uint64([]byte(byte_arr)))
			}else{
				byte_arr = preparePacket(packet_data[off:off+len])
				result[val.Name] = int64(binary.BigEndian.Uint32([]byte(byte_arr)))
			}

			last_offset = off+len

			if strings.Contains(val.Name, "num_"){
				iterate = int(result[val.Name])
			}
		}




	for i:=0;i<iterate-1;i++{
		for j:=0;j<len(repeat_parameter);j++{
			pa := repeat_parameter[j]
			len,_ := strconv.Atoi(pa.Length)
			if pa.Out_type == "int64"{
				byte_arr = preparePacket8(packet_data[last_offset:last_offset+len])
				result[pa.Name+"_"+strconv.Itoa(i+1)] = int64(binary.BigEndian.Uint64([]byte(byte_arr)))
			}else{
				byte_arr = preparePacket(packet_data[last_offset:last_offset+len])
				result[pa.Name+"_"+strconv.Itoa(i+1)] = int64(binary.BigEndian.Uint32([]byte(byte_arr)))
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

		revel.INFO.Println("Response Packet:", params[1],"(",formatter.ToHex(params[1]),")", "Enqueued")
	}


	revel.INFO.Println(result)
	return nil

}

func readPacket(arr []string, i int, j int) string{
	result := ""
	for ;i<=j;i++ {
		result+=arr[i]
	}
	return result
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


const ollo = `
 {
   "name": "ollo",
   "method": "post",
   "type":"hex",
   "url":[
   		"ollo",
   		"havells"
	],
	"db":{
		"db_url":[
			"mysql"],
		"db_type":"POSTGRES"
	},
   "parameters": {
     "0":{
       "in_type":"string",
       "out_type":"int64",
       "op":"swap8",
       "name":"Current",
       "dbcol":"Current",
       "indexed":"true",
       "description":"Power Consumption",
       "length":"4"
     },
     "4":{
       "in_type":"string",
       "out_type":"float64",
       "op":"swap16",
       "name":"Power",
       "dbcol":"power",
       "indexed":"true",
       "description":"Power Consumption",
       "length":"4"
     }
   }
 }
    `
const LNT  =`{
   "name": "'Vb, L&T Nova WM30KFC3CRS",
   "method": "post",
   "type":"hex",
   "url":[
   		"ollo",
   		"havells"
	],
	"db":{
		"db_url":[
			"mysql"],
		"db_type":"POSTGRES"
	},
   "parameters": {
     "6":{
       "in_type":"string",
       "out_type":"int64",
       "op":"int64",
       "name":"Vr",
       "dbcol":"Vr",
       "indexed":"true",
       "description":"Vr",
       "length":"8"
     },
     "14":{
       "in_type":"string",
       "out_type":"int64",
       "op":"int64",
       "name":"Vy",
       "dbcol":"Vy",
       "indexed":"true",
       "description":"Vy",
       "length":"8"
     },
     "22":{
       "in_type":"string",
       "out_type":"int64",
       "op":"int64",
       "name":"Vb",
       "dbcol":"Vb",
       "indexed":"true",
       "description":"Vb",
       "length":"8"
     },
     "30":{
       "in_type":"string",
       "out_type":"int64",
       "op":"int64",
       "name":"Ir",
       "dbcol":"Ir",
       "indexed":"true",
       "description":"Ir",
       "length":"8"
     },
     "38":{
       "in_type":"string",
       "out_type":"int64",
       "op":"int64",
       "name":"Iy",
       "dbcol":"Iy",
       "indexed":"true",
       "description":"Iy",
       "length":"8"
     },
     "46":{
       "in_type":"string",
       "out_type":"int64",
       "op":"int64",
       "name":"Ib",
       "dbcol":"Ib",
       "indexed":"true",
       "description":"Ib",
       "length":"8"
     },
     "54":{
       "in_type":"string",
       "out_type":"int64",
       "op":"int64",
       "name":"KWr",
       "dbcol":"KWr",
       "indexed":"true",
       "description":"KWr",
       "length":"8"
     },
     "62":{
       "in_type":"string",
       "out_type":"int64",
       "op":"int64",
       "name":"KWy",
       "dbcol":"KWy",
       "indexed":"true",
       "description":"KWy",
       "length":"8"
     },
     "70":{
       "in_type":"string",
       "out_type":"int64",
       "op":"int64",
       "name":"KWb",
       "dbcol":"KWb",
       "indexed":"true",
       "description":"KWb",
       "length":"8"
     },
     "78":{
       "in_type":"string",
       "out_type":"int64",
       "op":"int64",
       "name":"AvgV",
       "dbcol":"AvgV",
       "indexed":"true",
       "description":"AvgV",
       "length":"8"
     },
     "86":{
       "in_type":"string",
       "out_type":"int64",
       "op":"int64",
       "name":"AvgI",
       "dbcol":"AvgI",
       "indexed":"true",
       "description":"AvgI",
       "length":"8"
     },
     "94":{
       "in_type":"string",
       "out_type":"int64",
       "op":"int64",
       "name":"KW",
       "dbcol":"KW",
       "indexed":"true",
       "description":"KW",
       "length":"8"
     },
     "102":{
       "in_type":"string",
       "out_type":"int64",
       "op":"int64",
       "name":"FREQ",
       "dbcol":"FREQ",
       "indexed":"true",
       "description":"FREQ",
       "length":"8"
     },
     "110":{
       "in_type":"string",
       "out_type":"int64",
       "op":"int64",
       "name":"KWH",
       "dbcol":"KWH",
       "indexed":"true",
       "description":"KWH",
       "length":"8"
     },
     "118":{
       "in_type":"string",
       "out_type":"int64",
       "op":"int64",
       "name":"ActPower",
       "dbcol":"ActPower",
       "indexed":"true",
       "description":"ActPower",
       "length":"8"
     },
     "126":{
       "in_type":"string",
       "out_type":"int64",
       "op":"int64",
       "name":"PF",
       "dbcol":"PF",
       "indexed":"true",
       "description":"PF",
       "length":"8"
     }
   }
 }`
