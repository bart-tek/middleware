package main

import (
	"encoding/json"
	"flag"
	"log"
	"math/rand"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/Evrard-Nil/middleware/internal/client"
	"github.com/Evrard-Nil/middleware/internal/confcaptorstruct"
	"github.com/Evrard-Nil/middleware/internal/donneestruct"
	"github.com/Evrard-Nil/middleware/internal/enumnature"
)

var c confcaptorstruct.ConfCaptorStruct

func init() {
	c.GetConf()
}

func main() {
	publish()
}

func getDonnees() []byte {

	generatedData := donneestruct.DonneesCapteur{
		CapteurID:  generateCapteurID(),
		AeroportID: generateAeroportID(),
		Nature:     enumnature.WIND,
		Valeur:     generateValeur(),
		Date:       time.Now(),
	}
	json, err := json.Marshal(generatedData)
	if err != nil {
		log.Fatalf("Can't marshall data: %s", err)
	}
	return json
}

func generateValeur() float32 {

	var min float32 = -10
	var max float32 = 40
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

	return c.ListeAeroport[rand.Intn(max-min)+min]
}

func publish() {

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	connection := client.Connect()
	defer connection.Disconnect(250)

	qos := flag.Int("qos", client.GetQOS(), "The QoS to subscribe to messages at")
	topicTemp := flag.String("topicTemp", "captor/temperature", "temp topic")

	flag.Parse()
loop:
	for {
		select {
		default:
			connection.Publish(*topicTemp, byte(*qos), false, getDonnees())
			time.Sleep(time.Second * 5)
		case <-c:
			break loop
		}
	}
}
