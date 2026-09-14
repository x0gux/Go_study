package mqtt

import (
	"fmt"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

var ConnectOptions *mqtt.ClientOptions
var Client mqtt.Client

func Connect() {
	ConnectOptions = mqtt.NewClientOptions()
	ConnectOptions.AddBroker("mqtt://192.168.0.170:1883")
	ConnectOptions.SetClientID("sexyguy")
	ConnectOptions.SetUsername("tusk")
	ConnectOptions.SetPassword("Tusk12345!")

	Client = mqtt.NewClient(ConnectOptions)
	token := Client.Connect()
	if token.Wait() && token.Error() != nil {
		fmt.Println(token.Error())
	}
}
