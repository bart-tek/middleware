package confcaptorstruct

import (
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v2"
)

type ConfCaptorStruct struct {
	ListeAeroport []string `yaml:"airports"`
}

func (c *ConfCaptorStruct) GetConf() *ConfCaptorStruct {
	yamlFile, err := ioutil.ReadFile("internal/conf/confCaptor.yaml")
	if err != nil {
		log.Printf("yamlFile.Get err   #%v ", err)
	}
	err = yaml.Unmarshal(yamlFile, c)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}

	return c
}
