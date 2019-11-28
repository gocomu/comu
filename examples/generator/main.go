// demo package simulating a realtime generation and processing.
// Start the example from your terminal and type a letter + enter.
package main

import (
	"flag"
	"log"

	"github.com/go-audio/audio"
	"github.com/go-audio/generator"
	"github.com/go-audio/transforms"
	"github.com/gocomu/comu"
	"github.com/gocomu/comu/cio"
	"github.com/gocomu/comu/pattern"
)

var (
	sampleRate      = flag.Int("samplerate", 44100, "sample rate")
	channelNum      = flag.Int("channelnum", 1, "number of channel")
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

	tempo := comu.NewClock(240.0)
	sine := pattern.NewPattern(tempo, osc)
	go sine.Four2TheFloor([]int{1, 0, 1, 0, 1, 0})

	// go func() {
	// 	fmt.Println("tempo change")
	// 	time.Sleep(10 * time.Second)
	// 	tempo.NewBPM(60.0)
	// }()

	// Audio output
	// arg1 int: portAudio, oto
	// arg2 int: number of channels
	// arg3 int: buffersize
	outChannels := cio.NewOutput(cio.PortAudio, *channelNum, bufferSize)

	for {

		// populate the out buffer
		if err := osc.Fill(buf); err != nil {
			log.Printf("error filling up the buffer")
		}
		// apply vol control if needed (applied as a transform instead of a control
		// on the osc)
		//transforms.Gain(buf, 1)
		transforms.StereoPan(buf, 0.0)

		outChannels.Audio <- buf
	}
}
