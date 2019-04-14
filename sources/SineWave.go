package sources

import (
	"fmt"
	"math"
	"time"

	MQTT "github.com/eclipse/paho.mqtt.golang"
)

// SineWave generates a sine wave and sends it over MQTT
func SineWave(cxn MQTT.Client) {
	for i := 0.0; ; i = i + 0.01 {
		cxn.Publish("rocket_view/data/sine", 0, false, fmt.Sprintf("%.3f", math.Sin(i)))
		time.Sleep(25 * time.Millisecond)
	}
}
