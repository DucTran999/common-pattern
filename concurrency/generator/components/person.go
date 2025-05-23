package components

import (
	"fmt"
	"log"

	"github.com/DucTran999/shared-pkg/scrypto/caesar"
)

type person struct {
	name         string
	SecretChan   chan int
	sendChan     chan string
	receivedChan chan string
	messages     []string
	turn         int
}

func NewPerson(
	name string, messages []string, sendChan, receivedChan chan string,
	turn int,
) *person {
	return &person{
		name:         name,
		messages:     messages,
		sendChan:     sendChan,
		receivedChan: receivedChan,
		SecretChan:   make(chan int),
		turn:         turn,
	}
}

func (p *person) Chat() {
	sent := p.turn
	msgIndex := 0

	for nonce := range p.SecretChan {
		if sent%2 == 0 { // my turn to SEND
			if msgIndex >= len(p.messages) {
				log.Printf("%s: all messages sent - closing chat", p.name)
				close(p.sendChan)
				return
			}
			p.sendChan <- caesar.CaesarEncrypt(p.messages[msgIndex], nonce)
			msgIndex++
		} else { // my turn to RECEIVE
			cipher := <-p.receivedChan
			log.Printf("%s got cipher: %q", p.name, cipher)
			log.Printf("%s received message: %q", p.name, caesar.CaesarDecrypt(cipher, nonce))
			fmt.Println("---------------------------------------------------------")
		}
		sent++
	}
}
