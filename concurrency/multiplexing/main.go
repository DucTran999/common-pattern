package main

import (
	"context"
	"fmt"
	"log"
	"os/signal"
	"patterns/concurrency/multiplexing/components"
	"syscall"
)

// initRouter initializes the router with IPs, broadcast, and listen channels.
func initRouter(broadcast chan string, routerListen components.UnicastChan) (components.Router, error) {
	ips := []string{
		"192.186.1.12",
		"192.186.1.13",
		"192.186.1.14",
		"192.186.1.17",
	}

	return components.NewRouter(ips, broadcast, routerListen)
}

func main() {
	// Communication channels
	broadcastChan := make(chan string)
	routerListenChan := make(chan string)
	ackChan := make(chan string)

	// Initialize core components
	switchDevice := components.NewSwitch(broadcastChan, ackChan, routerListenChan)
	router, err := initRouter(broadcastChan, routerListenChan)
	if err != nil {
		log.Fatalf("failed to initialize router: %v", err)
	}

	// Create devices
	devices := []components.Device{
		components.NewDevice("Device 1", "192.186.1.12", "00:1A:2B:3C:4D:5E", ackChan),
		components.NewDevice("Device 2", "192.186.1.13", "00:1A:2B:3C:4D:5F", ackChan),
		components.NewDevice("Device 3", "192.186.1.14", "00:1A:2B:3C:4D:60", ackChan),
		components.NewDevice("Device 4", "192.186.1.17", "00:1A:2B:3C:4D:50", ackChan),
	}

	// Register devices with the switch
	var unicasts []components.UnicastChan
	for _, dev := range devices {
		unicasts = append(unicasts, dev.Unicast())
	}
	switchDevice.RegisterDeviceUnicast(unicasts...)

	// Start all goroutines
	go router.SendArp()
	go switchDevice.Listen()
	for _, dev := range devices {
		go dev.Listen()
	}

	// Graceful shutdown on SIGINT or SIGTERM
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	<-ctx.Done()
	fmt.Println("Shutdown app")
}
