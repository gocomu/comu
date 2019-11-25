// demo package simulating a realtime generation and processing.
// Start the example from your terminal and type a letter + enter.
package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/signal"

	"github.com/go-audio/audio"
	"github.com/go-audio/generator"
	"github.com/go-audio/transforms"
	"github.com/gocomu/comu"
	"github.com/gocomu/comu/comuio"
	"github.com/gocomu/comu/pattern"
)

var (
	sampleRate      = flag.Int("samplerate", 44100, "sample rate")
	channelNum      = flag.Int("channelnum", 2, "number of channel")
	bitDepthInBytes = flag.Int("bitdepthinbytes", 2, "bit depth in bytes")
)

func main() {
	bufferSize := 512
	buf := &audio.FloatBuffer{
		Data:   make([]float64, bufferSize),
		Format: audio.FormatMono44100,
	}
	osc := generator.NewOsc(generator.WaveSine, 440.0, buf.Format.SampleRate)
	osc.Amplitude = 0.5

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt, os.Kill)

	tempo := comu.NewClock(120.0)
	sine := pattern.NewPattern(tempo, osc)
	go sine.Four2TheFloor([]int{1, 0, 1, 0, 1, 0})

	// go func() {
	// 	fmt.Println("tempo change")
	// 	time.Sleep(10 * time.Second)
	// 	tempo.NewBPM(60.0)
	// }()

	// Audio output
	to := make(chan *audio.FloatBuffer, 1)
	go comuio.PortAudio(to, bufferSize)

	for {

		// populate the out buffer
		if err := osc.Fill(buf); err != nil {
			log.Printf("error filling up the buffer")
		}
		// apply vol control if needed (applied as a transform instead of a control
		// on the osc)
		//transforms.Gain(buf, 1)
		transforms.StereoPan(buf, 0.5)

		to <- buf
		select {
		case <-sig:
			fmt.Println("\tExiting..")
			return
		default:
		}
	}
}
