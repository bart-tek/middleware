package captor

import (
	"go/build"
	"io/ioutil"
	"log"
	"math/rand"
	"os"

	"gopkg.in/yaml.v2"
)

//Captor Structure
type Captor struct {
	ListeAeroport []string `yaml:"airports"`
	Topic         string   `yaml:"topic"`
	Qos           int      `yaml:"qos"`
	TimeBtwData   int      `yaml:"timeBtwData"`
	ClientID      string   `yaml:"cliendID"`
	Nature        string   `yaml:"nature"`
}

//GenerateValeur Generates value of captor
func (c *Captor) GenerateValeur(min float32, max float32) float32 {
	return min + rand.Float32()*(max-min)
}

//GenerateCapteurID Generates random ID to simulate multiple captors
func (c *Captor) GenerateCapteurID(min int, max int) int {
	return rand.Intn(max-min) + min
}

//GenerateAeroportID Generates airportID
func (c *Captor) GenerateAeroportID(min int, max int) string {
	return c.ListeAeroport[rand.Intn(max-min)+min]
}

//GetConf Retrieve conf from configs folder
func (c *Captor) GetConf(nat string) {
	gopath := os.Getenv("GOPATH")
	if gopath == "" {
		gopath = build.Default.GOPATH
	}
	projectPath := gopath + "/src/github.com/Evrard-Nil/middleware"
	yamlFile, err := ioutil.ReadFile(projectPath + "/configs/conf_captor_" + nat + ".yaml")
	if err != nil {
		log.Printf("yamlFile.Get err   #%v ", err)
	}
	err = yaml.Unmarshal(yamlFile, c)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}
}
