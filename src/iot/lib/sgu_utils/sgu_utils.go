package sgu_utils

import (
	"github.com/StabbyCutyou/buffstreams"
	"strconv"
	"strings"
	"iot/lib/parser"
	"github.com/revel/revel"
)


func ParseInputPackets(conn *buffstreams.Client)  {
	string_data := convert(conn.Data)
	revel.INFO.Println(string_data)
	parser.Wrap(conn)
}

func convert( b []byte ) string {
	s := make([]string,len(b))
	for i := range b {
		s[i] = strconv.Itoa(int(b[i]))
	}
	return strings.Join(s,",")
}