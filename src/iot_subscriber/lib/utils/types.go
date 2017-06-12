package utils

import (
	"encoding/json"
	"strconv"
	"time"
)

func Blank(id interface{}) bool {

	// lets only do type assertion and not type conversio for basic types.
	// 3x slow.. type-converion 6x slow.. but ok
	switch id.(type) {
	case int, int32, int64:
		if id == 0 {
			return true
		}
	case uint, uint32, uint64:
		if id == 0 {
			return true
		}

	case float32, float64:
		if id == 0.0 {
			return true
		}

	case json.Number:
		n, _ := strconv.ParseFloat(string(id.(json.Number)), 64)
		if n == 0.0 {
			return true
		}

	case nil:
		return true

	case string:
		if id == "" {
			return true
		}

	case bool:
		if id == false {
			return true
		}
	}
	return false
}

func Present(id interface{}) bool {
	return !Blank(id)
}

func ToStr(id interface{}) string {
	switch id.(type) {
	case uint:
		return strconv.FormatUint(uint64(id.(uint)), 10)
	case uint8:
		return strconv.FormatUint(uint64(id.(uint8)), 10)
	case uint16:
		return strconv.FormatUint(uint64(id.(uint16)), 10)
	case uint32:
		return strconv.FormatUint(uint64(id.(uint32)), 10)
	case uint64:
		return strconv.FormatUint(uint64(id.(uint64)), 10)

	case int:
		return strconv.FormatInt(int64(id.(int)), 10)
	case int8:
		return strconv.FormatInt(int64(id.(int8)), 10)
	case int16:
		return strconv.FormatInt(int64(id.(int16)), 10)
	case int32:
		return strconv.FormatInt(int64(id.(int32)), 10)
	case int64:
		return strconv.FormatInt(int64(id.(int64)), 10)

	case float32:
		return strconv.FormatFloat(float64(id.(float32)), 'E', -1, 32)
	case float64:
		return strconv.FormatFloat(float64(id.(float64)), 'E', -1, 64)

	case json.Number:
		return string(id.(json.Number))

	case string:
		return id.(string)
	}
	return ""
}

func ToUint64(id interface{}) uint64 {
	switch id.(type) {
	case uint:
		return uint64(id.(uint))
	case uint8:
		return uint64(id.(uint8))
	case uint16:
		return uint64(id.(uint16))
	case uint32:
		return uint64(id.(uint32))
	case uint64:
		return id.(uint64)

	case int:
		return uint64(id.(int))
	case int8:
		return uint64(id.(int8))
	case int16:
		return uint64(id.(int16))
	case int32:
		return uint64(id.(int32))
	case int64:
		return uint64(id.(int64))

	case float32:
		return uint64(id.(float32))
	case float64:
		return uint64(id.(float64))

	case json.Number:
		n, err := strconv.ParseUint(string(id.(json.Number)), 10, 64)
		if err == nil {
			return n
		}

	case string:
		n, err := strconv.Atoi(id.(string))
		if err == nil {
			return uint64(n)
		}
	}
	return 0
}

func ToInt64(id interface{}) int64 {
	switch id.(type) {
	case uint:
		return int64(id.(uint))
	case uint8:
		return int64(id.(uint8))
	case uint16:
		return int64(id.(uint16))
	case uint32:
		return int64(id.(uint32))
	case uint64:
		return int64(id.(uint64))

	case int:
		return int64(id.(int))
	case int8:
		return int64(id.(int8))
	case int16:
		return int64(id.(int16))
	case int32:
		return int64(id.(int32))
	case int64:
		return id.(int64)

	case float32:
		return int64(id.(float32))
	case float64:
		return int64(id.(float64))

	case json.Number:
		n, err := strconv.ParseInt(string(id.(json.Number)), 10, 64)
		if err == nil {
			return n
		}

	case string:
		n, err := strconv.Atoi(id.(string))
		if err == nil {
			return int64(n)
		}
	}
	return 0
}
func ToUint32(id interface{}) uint32 {
	switch id.(type) {
	case uint:
		return uint32(id.(uint))
	case uint8:
		return uint32(id.(uint8))
	case uint16:
		return uint32(id.(uint16))
	case uint32:
		return (id.(uint32))
	case uint64:
		return uint32(id.(uint64))

	case int:
		return uint32(id.(int))
	case int8:
		return uint32(id.(int8))
	case int16:
		return uint32(id.(int16))
	case int32:
		return uint32(id.(int32))
	case int64:
		return uint32(id.(int64))

	case float32:
		return uint32(id.(float32))
	case float64:
		return uint32(id.(float64))

	case json.Number:
		n, err := strconv.ParseUint(string(id.(json.Number)), 10, 32)
		if err == nil {
			return uint32(n)
		}

	case string:
		n, err := strconv.Atoi(id.(string))
		if err == nil {
			return uint32(n)
		}
	}
	return 0
}

func ToInt32(id interface{}) int32 {
	switch id.(type) {
	case uint:
		return int32(id.(uint))
	case uint8:
		return int32(id.(uint8))
	case uint16:
		return int32(id.(uint16))
	case uint32:
		return int32(id.(uint32))
	case uint64:
		return int32(id.(uint64))

	case int:
		return int32(id.(int))
	case int8:
		return int32(id.(int8))
	case int16:
		return int32(id.(int16))
	case int32:
		return (id.(int32))
	case int64:
		return int32(id.(int64))

	case float32:
		return int32(id.(float32))
	case float64:
		return int32(id.(float64))

	case json.Number:
		n, err := strconv.ParseInt(string(id.(json.Number)), 10, 32)
		if err == nil {
			return int32(n)
		}

	case string:
		n, err := strconv.Atoi(id.(string))
		if err == nil {
			return int32(n)
		}
	}
	return 0
}

func ToFloat64(id interface{}) float64 {
	switch id.(type) {
	case uint:
		return float64(id.(uint))
	case uint8:
		return float64(id.(uint8))
	case uint16:
		return float64(id.(uint16))
	case uint32:
		return float64(id.(uint32))
	case uint64:
		return float64(id.(uint64))

	case int:
		return float64(id.(int))
	case int8:
		return float64(id.(int8))
	case int16:
		return float64(id.(int16))
	case int32:
		return float64(id.(int32))
	case int64:
		return float64(id.(int64))

	case float32:
		return float64(id.(float32))
	case float64:
		return (id.(float64))

	case json.Number:
		n, err := strconv.ParseFloat(string(id.(json.Number)), 64)
		if err == nil {
			return n
		}

	case string:
		n, err := strconv.ParseFloat(id.(string), 64)
		if err == nil {
			return (n)
		}
	}
	return 0
}

func ToFloat32(id interface{}) float32 {
	switch id.(type) {
	case uint:
		return float32(id.(uint))
	case uint8:
		return float32(id.(uint8))
	case uint16:
		return float32(id.(uint16))
	case uint32:
		return float32(id.(uint32))
	case uint64:
		return float32(id.(uint64))

	case int:
		return float32(id.(int))
	case int8:
		return float32(id.(int8))
	case int16:
		return float32(id.(int16))
	case int32:
		return float32(id.(int32))
	case int64:
		return float32(id.(int64))

	case float32:
		return (id.(float32))
	case float64:
		return float32(id.(float64))

	case json.Number:
		n, err := strconv.ParseFloat(string(id.(json.Number)), 32)
		if err == nil {
			return float32(n)
		}

	case string:
		n, err := strconv.ParseFloat(id.(string), 32)
		if err == nil {
			return float32(n)
		}
	}
	return 0
}

func ToInt(id interface{}) int {
	return int(ToInt64(id))
}

func ToId(id interface{}) uint64 {
	return ToUint64(id)
}

func ToTime(id interface{}) int64 {
	return int64(ToId(id))
}

// unixtimestamp in Millisecond
func UnixTime() int64 {
	return int64(time.Nanosecond) * (time.Now().UnixNano() / int64(time.Millisecond))
}