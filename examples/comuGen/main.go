// demo package simulating a realtime generation and processing.
// Start the example from your terminal and type a letter + enter.
package main

import (
	"fmt"
	"log"
	"math"
	"time"

	"github.com/go-audio/audio"
	"github.com/go-audio/generator"
	"github.com/go-audio/transforms"
	"github.com/gocomu/comu"
	"github.com/gocomu/comu/cio"
)

func main() {
	var ar = []float64{5.0, 6.0, 7.0}
	for i, value := range ar {
		o := math.Nextafter(ar[i], value)
		fmt.Println(o)
	}
	fmt.Println(`
	
	
	`)
	// Audio output
	// arg1 cio.out: PortAudio, Oto
	// arg2 int: number of channels
	// arg3 cio.bufferSize
	comuIO := cio.NewAudioIO(cio.PortAudio, 2, cio.BS2048)

	buf := &audio.FloatBuffer{
		Data:   make([]float64, cio.BS2048),
		Format: audio.FormatStereo44100,
	}

	osc := generator.NewOsc(generator.WaveSine, 5000.0, buf.Format.SampleRate)
	//osc.Amplitude = 0.5

	tempo := comu.NewClock(120.0)
	sine := comu.NewPattern(tempo, osc)
	go sine.Four2TheFloor([]int{1, 0, 1, 0, 1, 0})

	// go func() {
	// 	//time.Sleep(5 * time.Second)
	// 	//tempo.NewBPM(120)
	// 	//tempo.BPMchange <- 120.0
	// 	for {
	// 		<-tempo.Beat.C
	// 		if tempo.BeatCounter == 4 {
	// 			fmt.Println("beat nu. 8")
	// 			// tempo.BPMchange <- 240.0
	// 			tempo.BPMchange <- 240.0
	// 		}
	// 	}
	// }()

	// make two channels two for communication with oscillator's goroutine
	panwait := make(chan float64)
	ampwait := make(chan float64)

	// init two goroutines that will increment "smoothly" (time operation)
	// go func() {
	// 	for {
	// 		time.Sleep(1 * time.Second)
	// 		panwait <- 0.0

	// 		time.Sleep(1 * time.Second)
	// 		panwait <- 1.0
	// 	}
	// }()
	go func() {
		for i := 0.0; i < 1.0; i = i + 0.01 {
			time.Sleep(100 * time.Millisecond)
			ampwait <- i
		}
	}()

	// variables to hold channels new sent values
	var pan float64
	var amp float64

	for {
		// populate the out buffer
		if err := osc.Fill(buf); err != nil {
			log.Printf("error filling up the buffer")
		}

		// non-blocking
		select {
		case pan = <-panwait:
		case amp = <-ampwait:
		default:
		}

		transforms.Gain(buf, amp)
		transforms.StereoPan(buf, pan)

		// pass populated buffer to port-audio stream
		comuIO.PortAudioOut(buf)
	}

}
