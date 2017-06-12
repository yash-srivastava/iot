package conf

import (
	"github.com/orcaman/concurrent-map"
	"io/ioutil"
	"log"
	"gopkg.in/yaml.v2"
)

var(
	SGU_TCP_CONNECTION cmap.ConcurrentMap
	SGU_SCU_LIST cmap.ConcurrentMap
	Retry_3000 cmap.ConcurrentMap
	PACKET_CONFIG Sgu_packet
	RESPONSE_PACKET_CONFIG Sgu_response_packet
	CUSTOM_PACKET_CONFIG Sgu_packet
	SERVER_PACKET_CONFIG Server_packet
	CONSTANTS map[string]string
)

func Init()  {
	load_constants()
	SGU_TCP_CONNECTION = cmap.New()
	SGU_SCU_LIST = cmap.New()
	Retry_3000 = cmap.New()
	PACKET_CONFIG = GetSguPacket()
	RESPONSE_PACKET_CONFIG = GetSguResponsePacket()
	CUSTOM_PACKET_CONFIG = GetCustomPackets()
	SERVER_PACKET_CONFIG = GetServerPacket()
}

func load_constants(){
	yamlFile, err := ioutil.ReadFile("constants.yml")
	if err != nil {
		log.Printf("yamlFile.Get err   #%v ", err)
	}
	CONSTANTS = make(map[string]string)
	err = yaml.Unmarshal(yamlFile, &CONSTANTS)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}
}