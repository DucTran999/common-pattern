package bridge_test

import (
	"bytes"
	"os"
	"patterns/structural/bridge"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_PrintDocument(t *testing.T) {
	t.Parallel()
	// Capture stdout to verify printed output
	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

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

	// Restore stdout and read captured output
	err := w.Close()
	require.NoError(t, err)
	os.Stdout = old
	var buf bytes.Buffer
	_, err = buf.ReadFrom(r)
	require.NoError(t, err)
	output := buf.String()

	require.Contains(t, output, "Printing using Cannon")
	assert.Contains(t, output, "Printing using Epson")
}
