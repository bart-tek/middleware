package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"

	"github.com/Evrard-Nil/middleware/internal/client"

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

	client := client.Connect()

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
		client.Publish("captor/wind", 0, false, captorString)
	}

	<-c
}
