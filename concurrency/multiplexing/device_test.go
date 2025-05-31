package multiplexing_test

import (
	"patterns/concurrency/multiplexing"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func Test_DeviceListen_MatchingIP(t *testing.T) {
	expectedMAC := "AA:BB:CC:DD:EE:FF"

	switchChan := make(chan string, 1) // buffered to prevent blocking
	d := multiplexing.NewDevice("TestDevice", "192.168.1.10", "AA:BB:CC:DD:EE:FF", switchChan)

	// Start the Listen goroutine
	go d.Listen()

	// Send matching IP to unicast channel
	d.Unicast() <- "192.168.1.10"

	// Wait and check response
	select {
	case resp := <-switchChan:
		assert.Equal(t, expectedMAC, resp)
	case <-time.After(1 * time.Second):
		t.Fatal("timeout waiting for MAC address response")
	}
}

func Test_DeviceListen_NonMatchingIP(t *testing.T) {
	expectedEmptyMAC := ""
	switchChan := make(chan string, 1)
	d := multiplexing.NewDevice("TestDevice", "192.168.1.10", "AA:BB:CC:DD:EE:FF", switchChan)

	// Start the Listen goroutine
	go d.Listen()

	// Send non-matching IP to unicast channel
	d.Unicast() <- "192.168.1.99"

	// Wait and check response
	select {
	case resp := <-switchChan:
		assert.Equal(t, expectedEmptyMAC, resp)
	case <-time.After(1 * time.Second):
		t.Fatal("timeout waiting for empty response")
	}
}
