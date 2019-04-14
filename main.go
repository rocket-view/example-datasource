package main

import (
	"math/rand"

	"./sources"
	MQTT "github.com/eclipse/paho.mqtt.golang"
)

const alphabet = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func main() {
	opts := MQTT.NewClientOptions().AddBroker("tcp://127.0.0.1:1883")
	cs := make([]byte, 23-len("rocketui"))
	for i := range cs {
		cs[i] = alphabet[rand.Int63()%int64(len(alphabet))]
	}
	opts.SetClientID(string(cs))
	c := MQTT.NewClient(opts)
	if token := c.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}
	go sources.SineWave(c)
	select {}
}
