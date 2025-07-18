package bridge

import (
	"log"
)

type Printer interface {
	PrintFile()
}

// Epson Printer
type Epson struct{}

func (e *Epson) PrintFile() {
	log.Println("Printing using Epson")
}

// Cannon Printer
type Cannon struct{}

func (c *Cannon) PrintFile() {
	log.Println("Printing using Cannon")
}
