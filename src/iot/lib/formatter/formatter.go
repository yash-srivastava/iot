package formatter

import (
	"encoding/json"
	"iot/lib/utils"
	"encoding/binary"
	"encoding/hex"
	"strings"
)

func GetStructFromInterface(data interface{} , structure interface{} ) (error) {
	bodyBytes, err := json.Marshal(data)
	if err !=nil {
		return err
	}
	err =json.Unmarshal(bodyBytes, &structure)
	return  err
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
