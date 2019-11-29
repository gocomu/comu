// demo package simulating a realtime generation and processing.
// Start the example from your terminal and type a letter + enter.
package main

import (
	"fmt"
	"log"

	"github.com/go-audio/audio"
	"github.com/go-audio/generator"
	"github.com/go-audio/transforms"
	"github.com/gocomu/comu"
	"github.com/gocomu/comu/cio"
)

func main() {
	// Audio output
	// arg1 cio.out: PortAudio, Oto
	// arg2 int: number of channels
	// arg3 int: buffer size
	comuIO := cio.NewAudioIO(cio.PortAudio, 2, cio.BS2048)

	buf := &audio.FloatBuffer{
		Data:   make([]float64, cio.BS2048),
		Format: audio.FormatStereo44100,
	}

	osc := generator.NewOsc(generator.WaveSine, 440.0, buf.Format.SampleRate)
	osc.Amplitude = 0.5

	tempo := comu.NewClock(120.0)
	sine := comu.NewPattern(tempo, osc)
	go sine.Four2TheFloor([]int{1, 0, 1, 0, 1, 0})

	go func() {
		//time.Sleep(5 * time.Second)
		//tempo.NewBPM(120)
		//tempo.BPMchange <- 120.0
		for {
			<-tempo.Beat.C
			if tempo.BeatCounter == 4 {
				fmt.Println("beat nu. 8")
				// tempo.BPMchange <- 240.0
				tempo.BPMchange <- 240.0
			}
		}
	}()

	for {
		// populate the out buffer
		if err := osc.Fill(buf); err != nil {
			log.Printf("error filling up the buffer")
		}

		transforms.StereoPan(buf, 0.0)
		//transforms.Gain(buf, 1)

		// pass populated buffer to port-audio stream
		comuIO.PortAudioOut(buf)
	}

}
