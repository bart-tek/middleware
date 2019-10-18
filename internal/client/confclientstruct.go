package client

import (
	"crypto/tls"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
	"gopkg.in/yaml.v2"
)

var conf ConfClientStruct

func init() {
	conf.GetConf()
}

// ConfClientStruct Structure récupérant la configuration des capteurs depuis un fichier de de configuration .yaml
type ConfClientStruct struct {
	AdresseBroker string `yaml:"adresse_broker"`
	PortBroker    int    `yaml:"port_broker"`
	NiveauQos     int    `yaml:"niveau_qos"`
	Username      string `yaml:"username"`
	Password      string `yaml:"password"`
}

// GetConf Méthode qui lit le fichier de configuration, le parse et peuple la structure avec la configuration récupérée
func (c *ConfClientStruct) GetConf() *ConfClientStruct {

	yamlFile, err := ioutil.ReadFile("internal/conf/confBroker.yaml")
	if err != nil {
		log.Printf("yamlFile.Get err   #%v ", err)
	}
	err = yaml.Unmarshal(yamlFile, c)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}

	return c
}

// GetQOS Retourne le niveau de QoS configuré
func GetQOS() int {
	return conf.NiveauQos
}

// Connect Initialise une connection avec le broker MQTT
func Connect() mqtt.Client {

	hostname, _ := os.Hostname()

	server := flag.String("server", conf.AdresseBroker+":"+strconv.Itoa(conf.PortBroker), "The full url of the MQTT server to connect to ex: tcp://127.0.0.1:1883")

	clientid := flag.String("clientid", hostname+strconv.Itoa(time.Now().Second()), "A clientid for the connection")
	username := flag.String("username", conf.Username, "A username to authenticate to the MQTT server")
	password := flag.String("password", conf.Password, "Password to match username")

	flag.Parse()

	opts := mqtt.NewClientOptions()
	opts.AddBroker(*server)
	opts.SetClientID(*clientid)
	opts.SetCleanSession(true)
	if *username != "" {
		opts.SetUsername(*username)
		if *password != "" {
			opts.SetPassword(*password)
		}
	}

	tlsConfig := &tls.Config{InsecureSkipVerify: true, ClientAuth: tls.NoClientCert}
	opts.SetTLSConfig(tlsConfig)

	client := mqtt.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	} else {
		fmt.Printf("Connected to %s\n", *server)
	}

	return client
}
