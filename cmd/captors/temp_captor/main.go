package main

import (
	"encoding/json"
	"log"
	"math/rand"
	"time"

	"github.com/Evrard-Nil/middleware/internal/captor"
	"github.com/Evrard-Nil/middleware/internal/donneestruct"
	"github.com/Evrard-Nil/middleware/internal/mqtt_client"
)

var c captor.Captor

func init() {
	c.GetConf(donneestruct.TEMP)
}

func main() {
	connection := mqtt_client.Connect(c.ClientID)
	defer connection.Disconnect(250)
	mqtt_client.Publish(connection, c.Qos, c.Topic, getDonnees, c.TimeBtwData)
}

func getDonnees() []byte {

	generatedData := donneestruct.DonneesCapteur{
		CapteurID:  c.GenerateCapteurID(0, 5),
		AeroportID: c.GenerateAeroportID(0, 14),
		Nature:     c.Nature,
		Valeur:     c.GenerateValeur(-10, 40),
		Date:       time.Now().AddDate(rand.Intn(20), rand.Intn(11), rand.Intn(31)),
	}
	json, err := json.Marshal(generatedData)
	if err != nil {
		log.Fatalf("Can't marshall data: %s", err)
	}
	return json
}
