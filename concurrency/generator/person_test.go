package generator_test

import (
	"patterns/concurrency/generator"
	"testing"
)

func TestPersonChat(t *testing.T) {
	senderToReceiver := make(chan string, 1)
	receiverToSender := make(chan string, 1)
	nonceChan := make(chan int, 10)

	sender := generator.NewPerson("Alice", []string{"Hello"}, senderToReceiver, receiverToSender, 0)
	receiver := generator.NewPerson("Bob", []string{"Hello"}, receiverToSender, senderToReceiver, 1)

	// Use a simple Caesar shift for testing
	go sender.Chat()
	go receiver.Chat()

	// Provide a nonce and close the channel to end chat
	nonces := []int{3, 4, 5}
	for _, n := range nonces {
		receiver.SecretChan <- n
		sender.SecretChan <- n
	}
	close(nonceChan)
}
