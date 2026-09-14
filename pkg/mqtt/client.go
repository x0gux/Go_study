package mqtt

import (
	"fmt"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

var ConnectOptions *mqtt.ClientOptions
var Client mqtt.Client

func Connect() {
	ConnectOptions = mqtt.NewClientOptions()
	ConnectOptions.AddBroker("mqtt://localhost:1883")
	ConnectOptions.SetClientID("goshop_mqtt")
	ConnectOptions.SetUsername("admin")
	ConnectOptions.SetPassword("1234")

	Client = mqtt.NewClient(ConnectOptions)
	token := Client.Connect()
	if token.Wait() && token.Error() != nil {
		fmt.Println(token.Error())
	}
}
