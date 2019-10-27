package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"os/signal"
	"strconv"
	"syscall"

	"github.com/Evrard-Nil/middleware/internal/donneestruct"
	"github.com/Evrard-Nil/middleware/internal/mqtt_client"
	"github.com/Evrard-Nil/middleware/internal/redis_client"
	MQTT "github.com/eclipse/paho.mqtt.golang"
	"github.com/gomodule/redigo/redis"
)

var redisCli redis.Conn
var mQTTCli MQTT.Client

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
		"WIND",
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
		"PRESS",
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
		"TEMP",
		fmt.Sprintf("%f", sensorData.Valeur))
}

// main function for csv_subscriber
//
func main() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	confRedis := redis_client.GetConf()
	redisCli = redis_client.ConnectToRedis(confRedis)
	defer redisCli.Close()

	mQTTCli = mqtt_client.Connect("redis_sub")
	mQTTCli.Subscribe("captors/temperature", 0, onTempReceived)
	mQTTCli.Subscribe("captors/pressure", 0, onPressReceived)
	mQTTCli.Subscribe("captors/wind", 0, onWindReceived)
	defer mQTTCli.Disconnect(250)

	<-c
}
