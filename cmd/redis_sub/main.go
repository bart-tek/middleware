package main

import (
	"encoding/json"
	"log"
	"os"
	"os/signal"
	"strconv"
	"syscall"

	"github.com/Evrard-Nil/middleware/internal/mqtt_client"
	"github.com/Evrard-Nil/middleware/internal/redis_client"

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
	// log.Printf("Received %s value: %s\n", sensorData.Nature, message.Payload())
	key := sensorData.AeroportID + ":" + sensorData.Nature + ":" + strconv.Itoa(sensorData.Date.Year())
	log.Printf("Inserting in key %s -- value: %v", key, sensorData.Valeur)
	_, err := redisCli.Do("ZADD", key, strconv.Itoa(int(sensorData.Date.Unix())), message.Payload())
	if err != nil {
		log.Printf("%s", err)
	}
}

func main() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	confRedis := redis_client.GetConf()
	redisCli = redis_client.ConnectToRedis(confRedis)
	defer redisCli.Close()

	mQTTCli = mqtt_client.Connect("redis_sub")
	mQTTCli.Subscribe("captors/#", 0, onValueReceived)
	defer mQTTCli.Disconnect(250)
	<-c
}
