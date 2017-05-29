package parser

import (
	"io/ioutil"
	"log"
	"fmt"
	"gopkg.in/yaml.v2"
)


type Sgu_packet struct {
	Delim    int `json:"delim"`
	Packets  map[int]Packets `json:"packets"`
}

type Packets struct {
	Parameters map[string]Parameters	`json:"parameters"`
}

type Parameters struct {
	Name     string `json:"name"`
	Length   string `json:"length"`
	In_type  string `json:"in_type"`
	Out_type string `json:"out_type"`
}

func GetConf() Sgu_packet {

	yamlFile, err := ioutil.ReadFile("packets.yml")
	if err != nil {
		log.Printf("yamlFile.Get err   #%v ", err)
	}
	c := Sgu_packet{}
	err = yaml.Unmarshal(yamlFile, &c)
	fmt.Println(c)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}
	return c
}
