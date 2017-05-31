package sgu_utils

import (
	"github.com/StabbyCutyou/buffstreams"

	"iot/lib/parser"
	"github.com/revel/revel"
	"iot/lib/utils"
)


func ParseInputPackets(conn *buffstreams.Client)  {
	string_data := utils.ConvertBytesToString(conn.Data)
	revel.INFO.Println(string_data)
	parser.Wrap(conn)
}

