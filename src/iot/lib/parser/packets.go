package parser

import (
	"io/ioutil"
	"log"
	"gopkg.in/yaml.v2"
)

func GetConf() map[string]interface{} {

	yamlFile, err := ioutil.ReadFile("packets.yml")
	if err != nil {
		log.Printf("yamlFile.Get err   #%v ", err)
	}
	var c map[string]interface{}
	err = yaml.Unmarshal(yamlFile, &c)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}
	return c
}
