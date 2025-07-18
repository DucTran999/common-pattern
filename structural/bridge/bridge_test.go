package bridge_test

import (
	"patterns/structural/bridge"
	"testing"
)

func Test_PrintDocument(t *testing.T) {
	t.Parallel()
	// Print document using cannon printer
	cannonDoc := bridge.Document{
		Printer: &bridge.Cannon{},
	}
	cannonDoc.Print()

	// Print document using Epson printer
	epsonDoc := bridge.Document{
		Printer: &bridge.Epson{},
	}
	epsonDoc.Print()
}
