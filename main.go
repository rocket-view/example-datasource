package main

import (
	"math/rand"
	"time"

	"./sources"
	MQTT "github.com/eclipse/paho.mqtt.golang"
)

const alphabet = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	opts := MQTT.NewClientOptions().AddBroker("tcp://127.0.0.1:1883")
	cs := make([]byte, 23-len("rocketui"))
	for i := range cs {
		cs[i] = alphabet[rand.Int63()%int64(len(alphabet))]
	}
	opts.SetClientID("rocketui" + string(cs))
	c := MQTT.NewClient(opts)
	if token := c.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}
	go sources.SineWave(c)
	select {}
}
