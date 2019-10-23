package mqtt_client

import (
	"crypto/tls"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/signal"
	"strconv"
	"syscall"
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
	projectPath := os.Getenv("GOPATH") + "/src/github.com/Evrard-Nil/middleware"
	yamlFile, err := ioutil.ReadFile(projectPath + "/configs/confBroker.yaml")
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
func Connect(clientID string) mqtt.Client {
	server := flag.String("server", conf.AdresseBroker+":"+strconv.Itoa(conf.PortBroker), "The full url of the MQTT server to connect to ex: tcp://127.0.0.1:1883")
	clientid := flag.String("clientid", clientID, "A clientid for the connection")
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

type payloadGen func() []byte

// Publish send values to a topic with mqtt client
func Publish(connection mqtt.Client, qosA int, topic string, payload payloadGen, sleepTime int) {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	qos := flag.Int("qos", qosA, "The QoS to subscribe to messages at")
	topicPress := flag.String("topicPress", topic, "Topic to publish on")

	flag.Parse()
loop:
	for {
		select {
		default:
			log.Printf("Publishing new val")
			connection.Publish(*topicPress, byte(*qos), false, payload())
			time.Sleep(time.Duration(sleepTime) * time.Second)
		case <-c:
			break loop
		}
	}
}
