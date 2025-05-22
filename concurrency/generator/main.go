package main

import (
	"context"
	"fmt"
	"os/signal"
	"patterns/concurrency/generator/components"
	"syscall"
	"time"
)

func main() {
	bobSentChan := make(chan string)
	aliceSentChan := make(chan string)

	bob := components.NewPerson("Bob", []string{
		"Hi, Alice",
		"Would you like to join me for dinner at the restaurant tonight at 7:00 pm?",
	}, bobSentChan, aliceSentChan, 0)
	alice := components.NewPerson("Alice", []string{
		"Hi, Bob",
		"Sounds great! See you at 7!",
	}, aliceSentChan, bobSentChan, 1)

	// Main context
	ctx, cancel := context.WithCancel(context.Background())

	secretChan := components.Generator(ctx)
	go func() {
		for s := range secretChan {
			secret := s
			go func(val int) { bob.SecretChan <- val }(secret)
			go func(val int) { alice.SecretChan <- val }(secret)
		}
	}()

	go bob.Chat()
	go alice.Chat()

	closeApp(cancel)
}

func closeApp(cancel context.CancelFunc) {
	// Create a context that listens for interrupt or terminate signals
	shutdownCtx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop() // Clean up signal handlers when function exits

	<-shutdownCtx.Done() // Block until a signal is received
	fmt.Println("Shutdown signal received...")

	// Cancel the main context (e.g., passed to workers, generators, etc.)
	cancel()

	// Optional: Wait a bit to allow graceful cleanup
	time.Sleep(2 * time.Second)
}
