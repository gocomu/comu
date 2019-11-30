// demo package simulating a realtime generation and processing.
// Start the example from your terminal and type a letter + enter.
package main

import (
	"time"

	"github.com/gocomu/comu/cio"
	"github.com/hypebeast/go-osc/osc"
)

func main() {
	// start a new osc server and client
	oscio := cio.NewOscIO("8765", "localhost", "8765")
	// init a handle
	oscio.Server.Handle("/message/address", func(msg *osc.Message) {
		osc.PrintMessage(msg)
	})

	time.Sleep(4 * time.Second)

	// send a message
	oscio.Message("/message/address", int32(666), true, "Hello")

	time.Sleep(1 * time.Second)
}
