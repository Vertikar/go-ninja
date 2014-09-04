package ninja

import (
	"fmt"

	MQTT "git.eclipse.org/gitroot/paho/org.eclipse.paho.mqtt.golang.git"
	"github.com/bitly/go-simplejson"
	"github.com/ninjasphere/go-ninja/logger"
)

type JsonMessageHandler func(string, *simplejson.Json)

// ChannelBus context for channel related bus operations.
type ChannelBus struct {
	name     string
	protocol string
	device   *DeviceBus
	channel  <-chan MQTT.Receipt
	log      *logger.Logger
}

// NewChannelBus Build a new channel bus for the supplied device
func NewChannelBus(name string, protocol string, d *DeviceBus) *ChannelBus {
	log := logger.GetLogger(fmt.Sprintf("channel.%s.%s", name, protocol))
	return &ChannelBus{
		name:     name,
		protocol: protocol,
		device:   d,
		log:      log,
	}
}

// SendEvent Publish an event on the channel bus.
func (cb *ChannelBus) SendEvent(event string, payload *simplejson.Json) error {

	json, err := payload.MarshalJSON()
	if err != nil {
		return err
	}

	cb.log.Infof("Sending event:%s payload:%s", event, json)

	receipt := cb.device.driver.mqtt.Publish(MQTT.QoS(0), "$driver/"+cb.device.driver.id+"/device/"+cb.device.id+"/channel/"+cb.name+"/"+cb.protocol+"/event/"+event, json)
	<-receipt

	return nil
}