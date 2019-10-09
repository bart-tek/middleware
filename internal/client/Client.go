package client

import (
	"fmt"
	"log"
	"net/url"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

/// Allows client to connect to broker
///
/// @params clientID (string), brokerURI (url.URL)
///
/// @return client
///
func Connect(clientID string, brokerURI *url.URL) mqtt.Client {
	opts := CreateClientOptions(clientID, brokerURI)
	client := mqtt.NewClient(opts)
	token := client.Connect()
	for !token.WaitTimeout(3 * time.Second) {
	}
	if err := token.Error(); err != nil {
		log.Fatal(err)
	}
	return client
}

/// Creates the options for the client connection to the broker
///
/// @params clientID (string), brokerURI (url.URL)
///
/// @return return
///
func CreateClientOptions(clientID string, brokerURI *url.URL) *mqtt.ClientOptions {
	opts := mqtt.NewClientOptions()
	opts.AddBroker(fmt.Sprintf("tcp://%s", brokerURI.Host))
	opts.SetUsername(brokerURI.User.Username())
	password, _ := brokerURI.User.Password()
	opts.SetPassword(password)
	opts.SetClientID(clientID)
	return opts
}

/// Listen to a broker and prints the msg topic and data
///
/// @params brokerURI (url.URL), topic (string)
///
/// @return void
///
func Listen(brokerURI *url.URL, topic string) {
	client := Connect("sub", brokerURI)
	client.Subscribe(topic, 0, func(client mqtt.Client, msg mqtt.Message) {
		fmt.Printf("* [%s] %s\n", msg.Topic(), string(msg.Payload()))
	})
}
