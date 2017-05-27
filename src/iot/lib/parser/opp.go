package parser

import (
	"encoding/hex"
	"iot/lib/utils"
	"log"
	"math"
	"strconv"
)

//reverse byte array
//input=>byte array output=>byte array

func reversebyte(arr []byte) []byte {
	for i := 0; i < len(arr)/2; i++ {
		j := len(arr) - i - 1
		arr[i], arr[j] = arr[j], arr[i]
	}
	return arr
}

//compare no. of bytes in hex string with max value
//input=>("1234",4) output=>true

func chkBytes(sval string, max int) bool {
	if len(sval) > max {
		log.Println("Error input must be of atmost ", max/2, " bytes")
		return false
	}
	return true
}

//function for constructing byte array on the basis of swap and number of bytes

func get(sval string, byt int, swap int) []byte {
	hval := make([]byte, byt)
	hval, _ = hex.DecodeString(sval)
	reversebyte(hval)
	tmp := make([]byte, byt)
	ind := 0
	cn := 0
	var ttl []byte
	for i := len(hval) - 1; i >= 0; i-- {
		ttl = append(ttl, hval[i])
		cn++
		if cn == swap/8 {
			for z := len(ttl) - 1; z >= 0; z-- {
				tmp[ind] = ttl[z]
				ind++
			}
			cn = 0
			ttl = ttl[:0]
		}
	}
	for i := len(hval); i < byt; i++ {
		tmp[ind] = byte(0)
		ind++
	}
	reversebyte(tmp)
	return (tmp)
}

//function to convert string to int
//input=>string, no. of bytes(4,8), unsigned conversion(true,false), 2scomplement(true,false)  output=>interface

func ToInt(sval string, by int, unsigned bool, twoscomplement bool) interface{} {
	if unsigned {
		data, _ := strconv.ParseUint((sval), 16, by*8)
		if twoscomplement {
			data = ^data + 1
			log.Println(data)
			return data
		} else {
			log.Println(data)
			return data
		}
	} else {
		data, _ := strconv.ParseInt((sval), 16, by*8)
		if twoscomplement {
			data = ^data + 1
			log.Println(data)
			return data
		} else {
			log.Println(data)
			return data
		}
	}
}

//function to convert string to float
//input=>string, no. of bytes(4,8)  output=>interface

func ToFloat(sval string, by int) interface{} {
	if by == 4 {
		i, err := strconv.ParseUint((sval), 16, 32)
		if err != nil {
			log.Println(err.Error())
			return -1
		}
		data := math.Float32frombits(uint32(i))
		return data
	} else {
		i, err := strconv.ParseUint((sval), 16, 64)
		if err != nil {
			log.Println(err.Error())
			return -1
		}
		data := math.Float64frombits(uint64(i))
		return data
	}
}

//function to implement swap on any data-type and return the value in specified type
//input=>data, required return type, swap(ex.8,16,32)	output=>interface

func Swap(val interface{}, ty string, swap int) interface{} {
	switch t := val.(type) {

	default:
		log.Fatalf("unexpected type %T\n", t)
		return nil

	case string:
		sval := val.(string)
		if ty == "int32" {
			if !chkBytes(sval, 8) {
				return -1
			}
			return SwapInt(sval, 4, swap, false)

		} else if ty == "int64" {
			if !chkBytes(sval, 16) {
				return -1
			}
			return SwapInt(sval, 8, swap, false)
		} else if ty == "uint32" {
			if !chkBytes(sval, 8) {
				return -1
			}
			return SwapInt(sval, 4, swap, true)
		} else if ty == "uint64" {
			if !chkBytes(sval, 16) {
				return -1
			}
			return SwapInt(sval, 8, swap, true)
		} else if ty == "float32" {
			if !chkBytes(sval, 8) {
				return -1
			}
			return SwapFloat(sval, 4, swap)
		} else if ty == "float64" {
			if !chkBytes(sval, 16) {
				return -1
			}
			return SwapFloat(sval, 8, swap)
		} else if ty == "ascii" {
			if !chkBytes(sval, 16) {
				return -1
			}
			return SwapAscii(sval, swap)
		}
	case float64, int64, float32, int32:
		if ty == "int32" {
			ival := utils.ToInt32(val)
			sval := strconv.FormatInt(int64(ival), 16)
			if !chkBytes(sval, 8) {
				return -1
			}
			return SwapInt(sval, 4, swap, false)
		} else if ty == "int64" {
			ival := utils.ToInt64(val)
			sval := strconv.FormatInt(ival, 16)
			if !chkBytes(sval, 16) {
				return -1
			}
			return SwapInt(sval, 8, swap, false)
		} else if ty == "uint32" {
			ival := utils.ToUint32(val)
			sval := strconv.FormatUint(uint64(ival), 16)
			if !chkBytes(sval, 8) {
				return -1
			}
			return SwapInt(sval, 4, swap, true)
		} else if ty == "uint64" {
			ival := utils.ToUint64(val)
			sval := strconv.FormatUint(ival, 16)
			if !chkBytes(sval, 16) {
				return -1
			}
			return SwapInt(sval, 8, swap, true)
		} else if ty == "float32" {
			ival := utils.ToFloat32(val)
			hx := math.Float32bits(ival)
			sval := strconv.FormatUint(uint64(hx), 16)
			if !chkBytes(sval, 8) {
				return -1
			}
			return SwapFloat(sval, 4, swap)
		} else if ty == "float64" {
			ival := utils.ToFloat64(val)
			hx := math.Float64bits(ival)
			sval := strconv.FormatUint((hx), 16)
			if !chkBytes(sval, 16) {
				return -1
			}
			return SwapFloat(sval, 8, swap)
		} else if ty == "ascii" {
			ival := utils.ToInt64(val)
			sval := strconv.FormatInt(ival, 16)
			if !chkBytes(sval, 16) {
				return -1
			}
			return SwapAscii(sval, swap)
		}
	}
	return 0
}

//function to convert data of any type to specific type
//input=>data, required type 	output=>interface

func Tonative(val interface{}, ty string) interface{} {
	switch t := val.(type) {

	default:
		log.Fatalf("unexpected type %T\n", t)
		return nil

	case string:
		sval := val.(string)
		if ty == "int32" {
			if !chkBytes(sval, 8) {
				return -1
			}
			return ToInt(sval, 4, false, false)

		} else if ty == "int64" {
			if !chkBytes(sval, 16) {
				return -1
			}
			return ToInt(sval, 8, false, false)
		} else if ty == "uint32" {
			if !chkBytes(sval, 8) {
				return -1
			}
			return ToInt(sval, 4, true, false)
		} else if ty == "uint64" {
			if !chkBytes(sval, 16) {
				return -1
			}
			return ToInt(sval, 8, true, false)
		} else if ty == "float32" {
			if !chkBytes(sval, 8) {
				return -1
			}
			return ToFloat(sval, 4)
		} else if ty == "float64" {
			if !chkBytes(sval, 16) {
				return -1
			}
			return ToFloat(sval, 8)
		} else if ty == "ascii" {
			if !chkBytes(sval, 16) {
				return -1
			}
			hval, _ := hex.DecodeString(sval)
			data := string((hval)[:])
			log.Println(data)
			return data
		}
	case float64, int64, float32, int32:
		if ty == "int32" {
			data := utils.ToInt32(val)
			log.Println(data)
			return data
		} else if ty == "int64" {
			data := utils.ToInt64(val)
			log.Println(data)
			return data
		} else if ty == "uint32" {
			data := utils.ToUint32(val)
			log.Println(data)
			return data
		} else if ty == "uint64" {
			data := utils.ToUint64(val)
			log.Println(data)
			return data
		} else if ty == "float32" {
			data := utils.ToFloat32(val)
			log.Println(data)
			return data
		} else if ty == "float64" {
			data := utils.ToFloat64(val)
			log.Println(data)
			return data
		} else if ty == "ascii" {
			ival := utils.ToInt64(val)
			sval := strconv.FormatInt(ival, 16)
			if !chkBytes(sval, 16) {
				return -1
			}
			hval, _ := hex.DecodeString(sval)
			data := string((hval)[:])
			log.Println(data)
			return data
		}
	}
	return 0
}

//function to find 2scomplement
//input=>data, required type 	output=>interface

func Twoscomplement(val interface{}, ty string) interface{} {
	switch t := val.(type) {

	default:
		log.Fatalf("unexpected type %T\n", t)
		return nil

	case string:
		sval := val.(string)
		if ty == "int32" {
			if !chkBytes(sval, 8) {
				return -1
			}
			return ToInt(sval, 4, false, true)
		} else if ty == "int64" {
			if !chkBytes(sval, 16) {
				return -1
			}
			return ToInt(sval, 8, false, true)
		} else if ty == "uint32" {
			if !chkBytes(sval, 8) {
				return -1
			}
			return ToInt(sval, 4, true, true)
		} else if ty == "uint64" {
			if !chkBytes(sval, 16) {
				return -1
			}
			return ToInt(sval, 8, true, true)
		} else if ty == "ascii" {
			if !chkBytes(sval, 16) {
				return -1
			}
			tmp, _ := strconv.ParseInt((sval), 16, 64)
			tmp = ^tmp + 1
			he := strconv.FormatInt(tmp, 16)
			hval, _ := hex.DecodeString(he)
			data := string((hval)[:])
			log.Println(data)
			return data
		}
	case float64, int64, float32, int32:
		if ty == "int32" {
			ival := utils.ToInt32(val)
			data := ^ival + 1
			log.Println(data)
			return data
		} else if ty == "int64" {
			ival := utils.ToInt64(val)
			data := ^ival + 1
			log.Println(data)
			return data
		} else if ty == "uint32" {
			ival := utils.ToUint32(val)
			data := ^ival + 1
			log.Println(data)
			return data
		} else if ty == "uint64" {
			ival := utils.ToUint64(val)
			data := ^ival + 1
			log.Println(data)
			return data
		} else if ty == "ascii" {
			ival := utils.ToInt64(val)
			tmp := ^ival + 1
			he := strconv.FormatInt(tmp, 16)
			hval, _ := hex.DecodeString(he)
			data := string((hval)[:])
			log.Println(data)
			return data
		}
	}
	return 0
}

func main() {
	var x float64
	x = 28022902

	Twoscomplement(x, "int32")
}