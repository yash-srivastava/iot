package tcp_server

import (
	"github.com/StabbyCutyou/buffstreams"
	"fmt"
	"strconv"
	"github.com/jrallison/go-workers"
)

var (
	bm *buffstreams.Manager
)
func Start_tcp_server()*buffstreams.Manager  {
	fmt.Println("TCP SERVER STARTED")
	cfg := buffstreams.TCPListenerConfig {
		EnableLogging: true, // true will have log messages printed to stdout/stderr, via log
		MaxMessageSize: 4096,
		Address: buffstreams.FormatAddress("", strconv.Itoa(62000)),
		Callback: HandleTcpRequest, // Any function type that adheres to this signature, you'll need to deserialize in here if need be
	}


	bm := buffstreams.NewManager()
	err := bm.StartListening(cfg)
	if err!=nil {
		fmt.Println(err)
	}
	return bm
}

func GetConnectionManager()*buffstreams.Manager{
	return bm
}
func HandleTcpRequest(conn buffstreams.Client) error {
	var duplicate buffstreams.Client
	duplicate = conn
	params := make(map[string]interface{})
	params["action"] = "parse_sgu_packets"
	params["client"] = duplicate

	workers.Enqueue("packets","parse_sgu_packets", params)

	return nil
}