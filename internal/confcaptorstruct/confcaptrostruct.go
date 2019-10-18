package confcaptorstruct

import (
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v2"
)

// ConfCaptorStruct Structure récupérant la configuration des capteurs depuis un ficheir de de configuration .yaml
type ConfCaptorStruct struct {
	ListeAeroport []string `yaml:"airports"`
}

// GetConf Méthode qui lit le fichier de configuration, le parse et peuple la structure avec la configuration récupérée
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
