package conf

import (
	"io/ioutil"
	"log"
	"gopkg.in/yaml.v2"
)


type Sgu_packet struct {
	Delim    int `json:"delim"`
	Packets  map[int]Packets `json:"packets"`
}

type Packets struct {
	Description string `json:"description"`
	Response_packet int `json:"response_packet"`
	Publish 	int `json:"publish"`
	Parameters map[string]Parameters	`json:"parameters"`
}

type Parameters struct {
	Name     string `json:"name"`
	Length   string `json:"length"`
	In_type  string `json:"in_type"`
	Out_type string `json:"out_type"`
}


type Sgu_response_packet struct {
	Delim    int `json:"delim"`
	Response_packets  map[int]Response_packets `json:"response_packets" yaml:"packets"`
}

type Response_packets struct {
	Description string `json:"description"`
	Length int `json:"length"`
	Response_parameters map[string]Response_parameters `json:"response_parameters" yaml:"parameters"`
}

type Response_parameters struct {
	Offset     int `json:"name"`
	Length   int `json:"length"`
}

type Incoming struct {
	SguId     uint64 `json:"sgu_id"`
	SeqNo   int64 `json:"seq_no"`
}

type Scu struct {
	ScuIds     []uint64 `json:"scu_ids"`
}

type Sgu struct {
	SguIds     []uint64 `json:"sgu_ids"`
}

type Server_packet struct {
	Delim    int `json:"delim" yaml:"delim"`
	Response_packets  map[int]Response_packets `json:"packets" yaml:"packets"`
}


func GetSguPacket() Sgu_packet {

	yamlFile, err := ioutil.ReadFile("packets.yml")
	if err != nil {
		log.Printf("yamlFile.Get err   #%v ", err)
	}
	c := Sgu_packet{}
	err = yaml.Unmarshal(yamlFile, &c)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}
	return c
}

func GetSguResponsePacket() Sgu_response_packet {

	yamlFile, err := ioutil.ReadFile("response_packets.yml")
	if err != nil {
		log.Printf("yamlFile.Get err   #%v ", err)
	}
	c := Sgu_response_packet{}
	err = yaml.Unmarshal(yamlFile, &c)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}
	return c
}


func GetCustomPackets() Sgu_packet {

	yamlFile, err := ioutil.ReadFile("custom_packets.yml")
	if err != nil {
		log.Printf("yamlFile.Get err   #%v ", err)
	}
	c := Sgu_packet{}
	err = yaml.Unmarshal(yamlFile, &c)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}
	return c
}

func GetServerPacket() Server_packet {

	yamlFile, err := ioutil.ReadFile("server_packets.yml")
	if err != nil {
		log.Printf("yamlFile.Get err   #%v ", err)
	}
	c := Server_packet{}
	err = yaml.Unmarshal(yamlFile, &c)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}
	return c
}