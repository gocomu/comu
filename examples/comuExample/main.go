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
	"github.com/gocomu/comu"
	"github.com/gocomu/comu/comuio"
	"github.com/gocomu/comu/pattern"
	"github.com/gordonklaus/portaudio"
	"github.com/hypebeast/go-osc/osc"
)

var (
	sampleRate      = flag.Int("samplerate", 44100, "sample rate")
	channelNum      = flag.Int("channelnum", 2, "number of channel")
	bitDepthInBytes = flag.Int("bitdepthinbytes", 2, "bit depth in bytes")
)

func main() {
	//comu.fftt()
	//comu.encode()
	//comu.R()
	oscio := comuio.NewOSCio("8765", "localhost", "8765")
	oscio.Server.Handle("/message/address", func(msg *osc.Message) {
		osc.PrintMessage(msg)
	})
	//TODO: clock auto receivers

	oscio.Message("/message/address", int32(666), true, "Hello")
	//TODO: bundle messages
	//TODO: clock auto messages

	bufferSize := 512
	buf := &audio.FloatBuffer{
		Data:   make([]float64, bufferSize),
		Format: audio.FormatMono44100,
	}
	osc := generator.NewOsc(generator.WaveSine, 440.0, buf.Format.SampleRate)
	osc.Amplitude = 0.5

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt, os.Kill)

	//gainControl := 0.0
	//currentVol := osc.Amplitude

	tempo := comu.NewClock(120.0)
	sine := pattern.NewPattern(tempo, osc)
	go sine.Four2TheFloor([]int{1, 0, 1, 0, 1, 0})

	// go func() {
	// 	fmt.Println("tempo change")
	// 	time.Sleep(10 * time.Second)
	// 	tempo.NewBPM(60.0)
	// }()

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
		//transforms.Gain(buf, 1)

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
