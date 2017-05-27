package parser

import (
	"encoding/hex"
	"log"
	"math"
	"strconv"
)

//main function for swap implementation and returning integer
//input=>data, no. of bytes(ex.4,8) expected, swap(ex.8,16,32), unsigned(true,false)		output=>interface

func SwapInt(sval string, by int, swap int, unsigned bool) interface{} {
	if len(sval)%2 != 0 {
		sval = "0" + sval
	}
	tmp := get(sval, by, swap)
	log.Println(hex.EncodeToString(tmp))
	if unsigned {
		data, _ := strconv.ParseUint(hex.EncodeToString(tmp), 16, by*8)
		log.Println(data)
		return data
	} else {
		data, _ := strconv.ParseInt(hex.EncodeToString(tmp), 16, by*8)
		log.Println(data)
		return data
	}
}

//main function for swap implementation and returning floating number
//input=>data, no. of bytes(ex.4,8) expected, swap(ex.8,16,32)	output=>interface

func SwapFloat(sval string, by int, swap int) interface{} {
	if len(sval)%2 != 0 {
		sval = "0" + sval
	}
	tmp := get(sval, by, swap)
	log.Println(hex.EncodeToString(tmp))
	i, err := strconv.ParseUint(hex.EncodeToString(tmp), 16, by*8)
	if err != nil {
		log.Println(err.Error())
		return -1
	}
	if by == 4 {
		data := math.Float32frombits(uint32(i))
		log.Println(data)
		return data
	} else {
		data := math.Float64frombits(uint64(i))
		log.Println(data)
		return data
	}

}

//main function for swap implementation and returning ascii
//input=>data, swap(ex.8,16,32)	output=>interface

func SwapAscii(sval string, swap int) interface{} {
	tmp := get(sval, 8, swap)
	log.Println(hex.EncodeToString(tmp))
	x := len(sval) / 2
	data := string(tmp[8-x:])
	log.Println(data)
	return data
}