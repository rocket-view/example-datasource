package sources

import (
	"fmt"
	"math"
	"time"

	MQTT "github.com/eclipse/paho.mqtt.golang"
)

// BeatWave generates a sine wave with beats and sends it over MQTT
func BeatWave(cxn MQTT.Client) {
	for i := 0.0; ; i = i + 0.01 {
		cxn.Publish("rocket_view/data/beat", 0, false, fmt.Sprintf("%.3f", math.Sin(i) + math.Sin(5 * i)))
		time.Sleep(25 * time.Millisecond)
	}
}
