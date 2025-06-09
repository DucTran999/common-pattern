package bridge

import "testing"

func Test_PrintDocument(t *testing.T) {
	// Print document using cannon printer
	doc := Document{
		Printer: &Cannon{},
	}
	doc.Print()

	// Print document using Epson printer
	docColorful := Document{
		Printer: &Epson{},
	}
	docColorful.Print()
}
