package generator_test

import (
	"patterns/concurrency/generator"
	"testing"
)

func TestPersonChat(t *testing.T) {
	senderToReceiver := make(chan string, 1)
	receiverToSender := make(chan string, 1)
	nonceChan := make(chan int)

	sender := generator.NewPerson("Alice", []string{"Hello"}, senderToReceiver, receiverToSender, 0)
	receiver := generator.NewPerson("Bob", []string{"Hello"}, receiverToSender, senderToReceiver, 1)

	// Provide a nonce and close the channel to end chat
	go func() {
		nonceChan <- 3
		nonceChan <- 7
		nonceChan <- 26
		nonceChan <- 24
		nonceChan <- 26
		nonceChan <- 22
		close(nonceChan)
	}()

	// Use a simple Caesar shift for testing
	go sender.Chat()
	go receiver.Chat()

	sender.SecretChan = nonceChan
	receiver.SecretChan = nonceChan
}
