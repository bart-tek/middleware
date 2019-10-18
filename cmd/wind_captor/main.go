package main

import (
	"crypto/tls"
	"flag"
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"

	"github.com/Evrard-Nil/middleware/internal/confcaptorstruct"
	"github.com/Evrard-Nil/middleware/internal/donneestruct"
	"github.com/Evrard-Nil/middleware/internal/enumnature"

	MQTT "github.com/eclipse/paho.mqtt.golang"
	mqtt "github.com/eclipse/paho.mqtt.golang"
)

func main() {
	publish()
}

func getDonnees() donneestruct.DonneesCapteur {

	return donneestruct.DonneesCapteur{
		CapteurID:  generateCapteurID(),
		AeroportID: generateAeroportID(),
		Nature:     enumnature.WIND,
		Valeur:     generateValeur(),
		Date:       time.Now(),
	}

}

func generateValeur() float32 {

	var min float32 = 0
	var max float32 = 140
	return min + rand.Float32()*(max-min)

}

func generateCapteurID() int {
	min := 1
	max := 5
	return rand.Intn(max-min) + min
}

func generateAeroportID() string {
	min := 0
	max := 14

	var c confcaptorstruct.ConfCaptorStruct
	c.GetConf()

	return c.ListeAeroport[rand.Intn(max-min)+min]
}

func connect() mqtt.Client {

	hostname, _ := os.Hostname()

	server := flag.String("server", "farmer.cloudmqtt.com:15652", "The full url of the MQTT server to connect to ex: tcp://127.0.0.1:1883")

	clientid := flag.String("clientid", hostname+strconv.Itoa(time.Now().Second()), "A clientid for the connection")
	username := flag.String("username", "pvpuovcq", "A username to authenticate to the MQTT server")
	password := flag.String("password", "h56KR9mXu9Xu", "Password to match username")

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

	client := MQTT.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	} else {
		fmt.Printf("Connected to %s\n", *server)
	}

	return client
}

func publish() {

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	client := connect()
	defer client.Disconnect(250)

	qos := flag.Int("qos", 0, "The QoS to subscribe to messages at")
	topicWind := flag.String("topicWind", "captor/wind", "wind topic")

	flag.Parse()
loop:
	for {
		select {
		default:
			client.Publish(*topicWind, byte(*qos), false, getDonnees().String())
			time.Sleep(time.Second)
		case <-c:
			break loop
		}
	}
}
