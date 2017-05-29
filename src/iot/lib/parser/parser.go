package parser

import (
	//"encoding/json"
	"strconv"
	//"log"
	"fmt"
	"encoding/binary"
	"strings"
)

/*type asset struct {
	Name       string                `json:"name"`
	Method     string                `json:"method"`
	Typ        string                `json:"type"`
	Url        []string              `json:"url"`
	Db         db                    `json:"db"`
	Parameters map[string]parameters `json:"parameters"`
}

type db struct {  `json:"db_type"`
}

type parameters struct {
	Name     st
	DbUrl  []string `json:"db_url"`
	DbType string ring `json:"name"`
	Dbcol    string `json:"dbcol"`
	Indexed  string `json:"indexed"`
	Desc     string `json:"description"`
	Len      string `json:"length"`
	In_type  string `json:"in_type"`
	Out_type string `json:"out_type"`
	Op       string `json:"op"`
}

type output struct {
	Data []data `json:"data"`
}

type data struct {
	Key   string      `json:"key"`
	Value interface{} `json:"value"`
}

var string_packet string*/

func Wrap(input []byte)map[string]interface{} {
	smap := make(map[string]int)
	smap["swap8"] = 8
	smap["swap16"] = 16
	smap["swap32"] = 32
	packet_data := input

	packet_config := GetConf()

	delim := int(packet_config.Delim)

	byte_arr :=preparePacket(packet_data[0:1])

	fmt.Print("data=>",byte_arr)
	val:= int(binary.BigEndian.Uint32(byte_arr))
	if val != delim{
		fmt.Print("Failed To Match Start Delim=>",val," ppp=>",delim)
		return nil
	}

	byte_arr = preparePacket(packet_data[1:3])
	fmt.Print("data=>",byte_arr)
	packet_length := int(binary.BigEndian.Uint32(byte_arr))

	fmt.Print("length=>",packet_length)

	byte_arr = preparePacket8(packet_data[3:9])
	sgu_id := (binary.BigEndian.Uint64(byte_arr))

	fmt.Print("sgu_id=>",sgu_id)

	byte_arr = preparePacket8(packet_data[9:23])
	timestamp:= int64(binary.BigEndian.Uint64(byte_arr))

	fmt.Print("timestamp=>",timestamp)

	byte_arr = preparePacket8(packet_data[23:27])
	seq_no := int64(binary.BigEndian.Uint64(preparePacket(byte_arr)))

	fmt.Print("seq_no=>",seq_no)

	byte_arr = preparePacket(packet_data[27:29])
	packet_type := int(binary.BigEndian.Uint32([]byte(byte_arr)))

	fmt.Print("packet_type=>",packet_type)

	packet_description := packet_config.Packets
	fmt.Print("packet_des=>",packet_description)

	var result map[string]int64
	result = make(map[string]int64)
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


			fmt.Println("val",val)

			fmt.Println("len,off=>",len,off)
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
			fmt.Println("Data for=", packet_data[off:off+len], " off=", off, " len=", off+len)
		}




	for i:=0;i<iterate-1;i++{
		fmt.Print("@@@@")
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
	fmt.Print(result)
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
	fmt.Println("inaa=>",arr)
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
