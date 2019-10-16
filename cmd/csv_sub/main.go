package main

import (
	"crypto/tls"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"

	"github.com/Evrard-Nil/middleware/internal/donneestruct"
	"github.com/Evrard-Nil/middleware/internal/enumnature"

	MQTT "github.com/eclipse/paho.mqtt.golang"
)

// onWinReceived is the function called when we get a MQTT message on the "captor/wind" channel
//
func onWindReceived(client MQTT.Client, message MQTT.Message) {
	var sensorData donneestruct.DonneesCapteur

	if err := json.Unmarshal(message.Payload(), &sensorData); err != nil {
		log.Printf("%s", err)
	}

	WriteCsv(sensorData.Date.String(),
		sensorData.AeroportID,
		strconv.Itoa(sensorData.CapteurID),
		sensorData.Nature,
		fmt.Sprintf("%f", sensorData.Valeur))
}

// onPressReceived is the function called when we get a MQTT message on the "captor/pressure" channel
//
func onPressReceived(client MQTT.Client, message MQTT.Message) {
	var sensorData donneestruct.DonneesCapteur

	if err := json.Unmarshal(message.Payload(), &sensorData); err != nil {
		log.Printf("%s", err)
	}

	WriteCsv(sensorData.Date.String(),
		sensorData.AeroportID,
		strconv.Itoa(sensorData.CapteurID),
		sensorData.Nature,
		fmt.Sprintf("%f", sensorData.Valeur))
}

// onTempReceived is the function called when we get a MQTT message on the "captor/temperature" channel
//
func onTempReceived(client MQTT.Client, message MQTT.Message) {
	var sensorData donneestruct.DonneesCapteur

	if err := json.Unmarshal(message.Payload(), &sensorData); err != nil {
		log.Printf("%s", err)
	}

	WriteCsv(sensorData.Date.String(),
		sensorData.AeroportID,
		strconv.Itoa(sensorData.CapteurID),
		sensorData.Nature,
		fmt.Sprintf("%f", sensorData.Valeur))
}

// main function for csv_subscriber
//
func main() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	hostname, _ := os.Hostname()

	server := flag.String("server", "farmer.cloudmqtt.com:15652", "The full url of the MQTT server to connect to ex: tcp://127.0.0.1:1883")
	topicWind := flag.String("topicWind", "captor/wind", "wind topic")
	topicPress := flag.String("topicPress", "captor/pressure", "pressure topic")
	topicTemp := flag.String("topicTemp", "captor/temperature", "temperature topic")
	qos := flag.Int("qos", 0, "The QoS to subscribe to messages at")
	clientid := flag.String("clientid", hostname+strconv.Itoa(time.Now().Second()), "A clientid for the connection")
	username := flag.String("username", "pvpuovcq", "A username to authenticate to the MQTT server")
	password := flag.String("password", "h56KR9mXu9Xu", "Password to match username")
	flag.Parse()

	connOpts := MQTT.NewClientOptions().AddBroker(*server).SetClientID(*clientid).SetCleanSession(true)
	if *username != "" {
		connOpts.SetUsername(*username)
		if *password != "" {
			connOpts.SetPassword(*password)
		}
	}
	tlsConfig := &tls.Config{InsecureSkipVerify: true, ClientAuth: tls.NoClientCert}
	connOpts.SetTLSConfig(tlsConfig)

	connOpts.OnConnect = func(c MQTT.Client) {
		if windToken := c.Subscribe(*topicWind, byte(*qos), onWindReceived); windToken.Wait() && windToken.Error() != nil {
			panic(windToken.Error())
		}
		if pressToken := c.Subscribe(*topicPress, byte(*qos), onPressReceived); pressToken.Wait() && pressToken.Error() != nil {
			panic(pressToken.Error())
		}
		if tempToken := c.Subscribe(*topicTemp, byte(*qos), onTempReceived); tempToken.Wait() && tempToken.Error() != nil {
			panic(tempToken.Error())
		}
	}

	client := MQTT.NewClient(connOpts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	} else {
		fmt.Printf("Connected to %s\n", *server)
	}

	// Test du publish
	captorStruct := donneestruct.DonneesCapteur{
		CapteurID:  0,
		AeroportID: "NTE",
		Nature:     enumnature.PRES,
		Valeur:     4,
		Date:       time.Now(),
	}

	for i := 0; i < 20; i++ {
		captorStruct.Valeur = captorStruct.Valeur + (float32(i) * 0.3)
		captorStruct.Date = captorStruct.Date.Add(20 * time.Second)
		captorString, err := json.Marshal(captorStruct)
		if err != nil {
			log.Printf("%s", err)
		}
		client.Publish("captor/pressure", 0, false, captorString)
	}

	<-c
}
