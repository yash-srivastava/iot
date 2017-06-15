package formatter

import (
	"iot_subscriber/lib/utils"
	"encoding/binary"
	"encoding/hex"
	"strings"
	"iot_subscriber/lib/subscriber"
	"encoding/base64"
	"github.com/revel/revel"
)

func GetMessageFromInterface(data interface{}) map[string]subscriber.Data  {
	msg_attr := data.(map[string]interface{})
	resp := map[string]subscriber.Data{}
	for k,v:=range msg_attr{
		tmp := subscriber.Data{}
		converted := v.(map[string]interface{})
		str_val := converted["Value"].(string)
		revel.INFO.Println(str_val)

		data, err := base64.StdEncoding.DecodeString(str_val)

		if err != nil {
			revel.ERROR.Println(err.Error())
		}
		revel.INFO.Println(string(data[:]))
		tmp.Value = string(data[:])
		resp[k] = tmp
	}
	return resp
}

func ToHex(val interface{}) string{
	input := utils.ToUint64(val)
	src := make([]byte, 8)
	binary.BigEndian.PutUint64(src, input)
	dst := make([]byte, hex.EncodedLen(len(src)))
	hex.Encode(dst, src)
	return strings.TrimLeft(string(dst),"0")
}

func Prettify(val interface{}) string  {
	return  utils.ToStr(val)+" ("+ToHex(val)+")"
}
