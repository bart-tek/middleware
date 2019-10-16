package main

import (
	"crypto/tls"
	"flag"
	"fmt"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"

	MQTT "github.com/eclipse/paho.mqtt.golang"
)

func onWindReceived(client MQTT.Client, message MQTT.Message) {
	fmt.Printf("Received wind value: %s\n", message.Payload())
}

func onPressReceived(client MQTT.Client, message MQTT.Message) {
	fmt.Printf("Received pressure value: %s\n", message.Payload())
}

func onTempReceived(client MQTT.Client, message MQTT.Message) {
	fmt.Printf("Received temperature value: %s\n", message.Payload())
}
func main() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	hostname, _ := os.Hostname()

	server := flag.String("server", "farmer.cloudmqtt.com:15652", "The full url of the MQTT server to connect")
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
	defer client.Disconnect(250)

	<-c
}
