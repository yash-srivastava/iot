package dial_tcp

import (
"log"
"strconv"
"time"

"github.com/StabbyCutyou/buffstreams"
"github.com/StabbyCutyou/buffstreams/test/message"
"github.com/golang/protobuf/proto"
	"fmt"
)

// Test client to send a sample payload of data endlessly
// By default it points locally, but it can point to any network address
// TODO Make that externally configurable to make automating the test easier
func Connect() {
	cfg := &buffstreams.TCPConnConfig{
		MaxMessageSize: 2048,
		Address:        buffstreams.FormatAddress("127.0.0.1", strconv.Itoa(6000)),
	}
	/*name := "Client"
	date := time.Now().UnixNano()
	data := "This is an intenntionally long and rambling sentence to pad out the size of the message."
	msg := &message.Note{Name: &name, Date: &date, Comment: &data}
	msgBytes, err := proto.Marshal(msg)
	if err != nil {
		log.Print(err)
	}*/
	btw, err := buffstreams.DialTCP(cfg)
	if err != nil {
		log.Fatal(err)
	}
	var msgBytes []byte
	msgBytes = []byte{0x7e,0x00,0x2a,0x34,0x30,0x36,0x37,0x38,0x33,0x32,0x30,0x31,0x36,0x30,0x31,0x32,0x32,0x31,0x38,0x35,0x30,0x33,0x34,0xf3,0x00,0x00,0x00,0x00,0x06,0x05,0x01,0x03,0x00,0xff,0x03,0x00,0xff,0x03,0x00,0xff,0x03,0x00,0xff,0x03,0x00}
	fmt.Print(msgBytes)
	_, err = btw.Write(msgBytes)
	if err != nil {
		log.Print("There was an error")
		log.Print(err)
	}
	time.Sleep(100 * time.Millisecond)
	readBytes := make([]byte, 4096)
	msgLen, err := btw.Read(readBytes)
	if err != nil {
		log.Printf("Address %s: Failure to read from connection", err)
		btw.Close()
		return
	}
	read_msg := &message.Note{}
	err = proto.Unmarshal(readBytes[:msgLen], read_msg)
	if err!=nil{
		fmt.Print(err)
	}
	fmt.Println("dsds")
	fmt.Print(read_msg)



}