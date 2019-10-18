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
	MQTT "github.com/eclipse/paho.mqtt.golang"
	"github.com/gomodule/redigo/redis"
)

var redisCli redis.Conn
var mQTTCli MQTT.Client

func onValueReceived(client MQTT.Client, message MQTT.Message) {
	var sensorData donneestruct.DonneesCapteur
	if err := json.Unmarshal(message.Payload(), &sensorData); err != nil {
		log.Printf("%s", err)
	}
	fmt.Printf("Received "+sensorData.Nature+" value: %s\n", message.Payload())
	key := sensorData.AeroportID + "." + strconv.Itoa(sensorData.Date.Year())
	_, err := redisCli.Do("ZADD", key, message.Payload())
	if err != nil {
		log.Printf("%s", err)
	}
}

func main() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	redisCli = newRedisClient("redis-10932.c1.us-west-2-2.ec2.cloud.redislabs.com:10932", "uutPD4Eh1qkYtGWxiuYvfXE7Ri5N7oPQ")
	mQTTCli = newMQQTClient()
	defer redisCli.Close()
	defer mQTTCli.Disconnect(250)
	<-c
}

func newRedisClient(addr string, pass string) redis.Conn {
	client, err := redis.Dial("tcp", addr, redis.DialPassword(pass))
	if err != nil {
		log.Fatal(err)
	} else {
		log.Printf("Succesfully connected to Redis at %s\n", addr)
	}
	return client
}

func newMQQTClient() MQTT.Client {
	hostname, _ := os.Hostname()
	server := flag.String("server", "farmer.cloudmqtt.com:15652", "The full url of the MQTT server to connect")
	captorTopic := flag.String("topicWind", "captor/*", "Topic")
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
		if windToken := c.Subscribe(*captorTopic, byte(*qos), onValueReceived); windToken.Wait() && windToken.Error() != nil {
			panic(windToken.Error())
		}
	}

	client := MQTT.NewClient(connOpts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	} else {
		log.Printf("Succesfully connected to MQQT at %s\n", *server)
	}
	return client
}
