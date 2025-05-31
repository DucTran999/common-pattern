package multiplexing_test

import (
	"patterns/concurrency/multiplexing"
	"testing"
)

func Test_SwitchListen(t *testing.T) {
	broadcast := make(chan string, 1)
	ack := make(chan string, 1)
	routerUnicast := make(chan string, 1)
	deviceReceived := make(chan string, 1)

	s := multiplexing.NewSwitch(broadcast, ack, routerUnicast)
	s.RegisterDeviceUnicast(deviceReceived)

	go s.Listen()

	// send IP want to ask
	s.BroadcastChan() <- "192.168.1.23"
	close(broadcast)

	// Get IP from switch
	<-deviceReceived

	// Device send fake MAC via ack channel
	ack <- "AA:BB:CC:DD:EE:FF"
}
