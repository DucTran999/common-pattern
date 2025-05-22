package components

import (
	"fmt"
	"log"
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
		if sent%2 == 0 {
			go func() {
				if msgIndex < len(p.messages) {
					p.sendChan <- CaesarEncrypt(p.messages[msgIndex], nonce)
					msgIndex++
					sent++
				} else {
					log.Fatal("end")
				}
			}()
		} else {
			go func() {
				cipher := <-p.receivedChan
				log.Printf("%s got cipher: %v", p.name, cipher)
				log.Printf("%s received message: %v", p.name, CaesarDecrypt(cipher, nonce))
				fmt.Println("---------------------------------------------------------")
				sent++
			}()
		}
	}
}
