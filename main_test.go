package main

import (
	"bytes"
	"log"
	"os"
	"testing"
)

func TestPrintSomething(t *testing.T) {

	output := captureOutput(func() {
		PrintSomething("input value")
	})
	testOutput(output, "Printing: input value\n", t)

}

func captureOutput(f func()) string {
	var buf bytes.Buffer
	log.SetOutput(&buf)
	log.SetFlags(0)
	f()
	log.SetOutput(os.Stderr)
	return buf.String()
}

func testOutput(output string, expected string, t *testing.T) {
	if output != expected {
		t.Error("UNEXPECTED OUTPUT", "Expected:", expected, "Output:", output)
	}
}
