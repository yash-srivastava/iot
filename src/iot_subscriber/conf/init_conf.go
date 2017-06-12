package conf

import (
	"io/ioutil"
	"log"
	"gopkg.in/yaml.v2"
)

var(
	CONSTANTS map[string]string
)

func Init()  {
	load_constants()
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