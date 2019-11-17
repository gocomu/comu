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
	"github.com/gocomu/comu/tempo"
	"github.com/gordonklaus/portaudio"
)

var (
	sampleRate      = flag.Int("samplerate", 44100, "sample rate")
	channelNum      = flag.Int("channelnum", 2, "number of channel")
	bitDepthInBytes = flag.Int("bitdepthinbytes", 2, "bit depth in bytes")
)

func main() {
	//fftt()
	//encode()
	//r()

	bufferSize := 512
	buf := &audio.FloatBuffer{
		Data:   make([]float64, bufferSize),
		Format: audio.FormatMono44100,
	}
	currentNote := 440.0
	osc := generator.NewOsc(generator.WaveSine, currentNote, buf.Format.SampleRate)
	osc.Amplitude = 1

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt, os.Kill)

	gainControl := 0.0
	currentVol := osc.Amplitude

	fmt.Println(`This is a demo, press a key followed by enter, the played note should change.
Use the - and + keys follow by enter to decrease or increase the volume\nPress q or ctrl-c to exit.
Note that the sound will come out of your default sound card.`)

	tempo := tempo.NewTempo(120.0)
	tempoChange := 30.0

	go func() {
		for {
			select {
			case <-tempo.BPMchannel:
				//fmt.Println("received message", nextBeat)
				osc.SetFreq(currentNote)
				currentNote = currentNote + 20.0
				tempo.BPMchange <- tempoChange
				tempoChange = tempoChange + 10.0
			default:
				//fmt.Println("no message received")
			}
		}
	}()

	// Audio output
	portaudio.Initialize()
	defer portaudio.Terminate()
	out := make([]float32, bufferSize)
	stream, err := portaudio.OpenDefaultStream(0, 1, 44100, len(out), &out)
	if err != nil {
		log.Fatal(err)
	}
	defer stream.Close()

	if err := stream.Start(); err != nil {
		log.Fatal(err)
	}
	defer stream.Stop()

	for {

		// populate the out buffer
		if err := osc.Fill(buf); err != nil {
			log.Printf("error filling up the buffer")
		}
		// apply vol control if needed (applied as a transform instead of a control
		// on the osc)
		if gainControl != 0 {
			currentVol += gainControl
			if currentVol < 0.1 {
				currentVol = 0
			}
			if currentVol > 6 {
				currentVol = 6
			}
			fmt.Printf("new vol %f.2", currentVol)
			gainControl = 0
		}
		transforms.Gain(buf, currentVol)

		f64ToF32Copy(out, buf.Data)

		// write to the stream
		if err := stream.Write(); err != nil {
			log.Printf("error writing to stream : %v\n", err)
		}
		select {
		case <-sig:
			fmt.Println("\tCiao!")
			return
		default:
		}
	}
}

// portaudio doesn't support float64 so we need to copy our data over to the
// destination buffer.
func f64ToF32Copy(dst []float32, src []float64) {
	for i := range src {
		dst[i] = float32(src[i])
	}
}

//  start: First value of the sequence.
//  end:   The sequence is ended upon reaching the end value.
//  step:  step will be used as the increment between elements in the sequence.
//         step should be given as a positive number.
func stepSeq(start, end, step int) []int {
	if step <= 0 || end < start {
		return []int{}
	}
	s := make([]int, 0, 1+(end-start)/step)
	for start <= end {
		s = append(s, start)
		start += step
	}
	return s
}

func seq(min, max int) []int {
	a := make([]int, max-min+1)
	for i := range a {
		a[i] = min + i
	}
	return a
}

// Returns a slice of elements with exact count.
// step will be used as the increment between elements in the sequence.
// step should be given as a positive, negative or zero number.
func ser(start, count, step int) []int {
	s := make([]int, count)
	for i := range s {
		s[i] = start
		start += step
	}
	return s
}
