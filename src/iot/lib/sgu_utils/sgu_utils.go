package sgu_utils

import (
	"github.com/StabbyCutyou/buffstreams"
	//"github.com/golang/protobuf/proto"
	"fmt"
	//"time"
	//"github.com/StabbyCutyou/buffstreams/test/message"
	"strconv"
	"strings"
	"iot/lib/parser"
)


func ParseInputPackets(conn *buffstreams.Client)  {
	fmt.Print("called")
	string_data := convert(conn.Data)
	fmt.Print(string_data)
	parser.Wrap(conn)
	/*name := "Server"
	date := time.Now().UnixNano()
	data := "Reply from Server"
	wmsg := &message.Note{Name: &name, Date: &date, Comment: &data}
	msgBytes, err := proto.Marshal(wmsg)
	if err != nil {
		fmt.Print(err)
	}

	client,_ := buffstreams.TcpClients.Get(conn.Address)
	conv,ok := client.(*buffstreams.TCPConn)
	if !ok{
		fmt.Print("Invalid Connection")
	}
	wr,e:=conv.Write(msgBytes)
	fmt.Print("line",wr)
	if e!=nil{
		fmt.Print(e.Error())
	}*/
}

func convert( b []byte ) string {
	s := make([]string,len(b))
	for i := range b {
		s[i] = strconv.Itoa(int(b[i]))
	}
	return strings.Join(s,",")
}