package utils

import (
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
	"strconv"
	"strings"
)

type Reader struct {
	*bytes.Buffer
}

// So that it implements the io.ReadCloser interface
func (Reader) Close() error { return nil }

func MapToString(m interface{}) string {
	b := &bytes.Buffer{}
	json.NewEncoder(b).Encode(&m)
	return b.String()
}

func MapToBody(m interface{}) io.ReadCloser {
	result := &bytes.Buffer{}
	json.NewEncoder(result).Encode(&m)
	return Reader{result}
}

func BodyToMap(b io.ReadCloser) map[string]interface{} {
	var m map[string]interface{}

	d := json.NewDecoder(b)
	d.UseNumber()
	d.Decode(&m)

	return m
}

//counter -> no of clone copies
func CloneBody(body io.ReadCloser, counter int) []io.ReadCloser {
	result := make([]io.ReadCloser, counter)

	buf, _ := ioutil.ReadAll(body)
	for i := 0; i < counter; i++ {
		result[i] = Reader{bytes.NewBuffer(buf)}
	}
	return result
}

func ConvertBytesToString( b []byte ) string {
	s := make([]string,len(b))
	for i := range b {
		s[i] = strconv.Itoa(int(b[i]))
	}
	return strings.Join(s,"")
}