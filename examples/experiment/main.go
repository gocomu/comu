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
}
