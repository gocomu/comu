// demo package simulating a realtime generation and processing.
// Start the example from your terminal and type a letter + enter.
package main

import (
	"flag"

	"github.com/gocomu/comu"
)

var (
	sampleRate      = flag.Int("samplerate", 44100, "sample rate")
	channelNum      = flag.Int("channelnum", 2, "number of channel")
	bitDepthInBytes = flag.Int("bitdepthinbytes", 2, "bit depth in bytes")
)

func main() {
	comu.fftt()
	comu.encode()
	comu.R()
	// oscio := comuio.NewOSCio("8765", "localhost", "8765")
	// oscio.Server.Handle("/message/address", func(msg *osc.Message) {
	// 	osc.PrintMessage(msg)
	// })
	//TODO: clock auto receivers

	// oscio.Message("/message/address", int32(666), true, "Hello")
	//TODO: bundle messages
	//TODO: clock auto messages

}
