package bridge

import "fmt"

type Printer interface {
	PrintFile()
}

// Epson Printer
type Epson struct{}

func (e *Epson) PrintFile() {
	fmt.Println("Printing using Epson")
}

// Cannon Printer
type Cannon struct{}

func (c *Cannon) PrintFile() {
	fmt.Println("Printing using Cannon")
}
