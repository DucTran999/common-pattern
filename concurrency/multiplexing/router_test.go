package multiplexing_test

import (
	"patterns/concurrency/multiplexing"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_NewRouter_MissingBroadcast(t *testing.T) {
	listen := make(chan string)
	ipList := []string{"192.168.1.10"}

	router, err := multiplexing.NewRouter(ipList, nil, listen)

	require.ErrorIs(t, multiplexing.ErrMissingBroadcastChannel, err)
	assert.Nil(t, router)
}

func Test_NewRouter_MissingIpList(t *testing.T) {
	listen := make(chan string)
	broadcast := make(chan string)
	router, err := multiplexing.NewRouter([]string{}, broadcast, listen)

	require.ErrorIs(t, multiplexing.ErrEmptyIPList, err)
	assert.Nil(t, router)
}

func Test_SendArp_FoundMAC(t *testing.T) {
	listen := make(chan string)
	broadcast := make(chan string)
	ipList := []string{"192.168.1.10", "192.168.1.1"}

	router, err := multiplexing.NewRouter(ipList, broadcast, listen)
	require.NoError(t, err)

	go router.SendArp()
	broadcastIp := <-broadcast
	assert.Equal(t, "192.168.1.10", broadcastIp)

	// Fake MAC
	listen <- "00:1A:2B:3C:4D:50"

	broadcastIp = <-broadcast
	assert.Equal(t, "192.168.1.1", broadcastIp)

	// Fake MAC not found
	listen <- ""
}
