// demo package simulating a realtime generation and processing.
// Start the example from your terminal and type a letter + enter.
package main

import (
	"flag"
	"time"

	"github.com/gocomu/comu/comuio"
	"github.com/hypebeast/go-osc/osc"
)

var (
	sampleRate      = flag.Int("samplerate", 44100, "sample rate")
	channelNum      = flag.Int("channelnum", 2, "number of channel")
	bitDepthInBytes = flag.Int("bitdepthinbytes", 2, "bit depth in bytes")
)

func main() {
	// start a new osc server and client
	oscio := comuio.NewOSCio("8765", "localhost", "8765")
	// init a handle
	oscio.Server.Handle("/message/address", func(msg *osc.Message) {
		osc.PrintMessage(msg)
	})

	time.Sleep(4 * time.Second)

	// send a message
	oscio.Message("/message/address", int32(666), true, "Hello")

	time.Sleep(1 * time.Second)
}
