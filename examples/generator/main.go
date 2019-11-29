// demo package simulating a realtime generation and processing.
// Start the example from your terminal and type a letter + enter.
package main

import (
	"log"

	"github.com/go-audio/audio"
	"github.com/go-audio/generator"
	"github.com/gocomu/comu"
	"github.com/gocomu/comu/cio"
)

func main() {
	bufferSize := 512
	buf := &audio.FloatBuffer{
		Data:   make([]float64, bufferSize),
		Format: audio.FormatStereo44100,
	}
	out := make([]float32, bufferSize)

	osc := generator.NewOsc(generator.WaveSine, 440.0, buf.Format.SampleRate)
	osc.Amplitude = 0.5

	tempo := comu.NewClock(240.0)
	sine := comu.NewPattern(tempo, osc)
	go sine.Four2TheFloor([]int{1, 0, 1, 0, 1, 0})

	// go func() {
	// 	fmt.Println("tempo change")
	// 	time.Sleep(10 * time.Second)
	// 	tempo.NewBPM(float64(rand.Intn(250)))

	// if p.clock.BeatCounter == 8 {
	// 	//fmt.Println("tempo change")
	// 	p.clock.BPMchange <- 120.0
	// }
	// }()

	// Audio output
	// arg1 cio.out: PortAudio, Oto
	// arg2 int: number of channels
	// arg3 int: buffer size
	comuIO := cio.NewAudioIO(cio.PortAudio, 2, bufferSize, out)
	for {
		// populate the out buffer
		if err := osc.Fill(buf); err != nil {
			log.Printf("error filling up the buffer")
		}

		//transforms.Gain(buf, 1)
		//transforms.StereoPan(buf, rand.Float64())

		comuIO.PortAudioFunc(out, buf)
	}

}
