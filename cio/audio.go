package cio

import (
	"log"

	"github.com/go-audio/audio"
	"github.com/gordonklaus/portaudio"
)

type out int

const (
	PortAudio = out(iota)
	Oto
)

type AIO struct {
	Audio   chan *audio.FloatBuffer
	Trigger chan bool
}

//func NewOutput(audioOutput out, numberOfChannels, bufferSize int) chan *audio.FloatBuffer {
func NewOutput(audioOutput out, numberOfChannels, bufferSize int) *AIO {
	aio := &AIO{
		Audio: make(chan *audio.FloatBuffer),
	}
	switch audioOutput {
	case PortAudio:
		go aio.portAudio(numberOfChannels, bufferSize)

	case Oto:
		// go aio.oto(numberOfChannels, bufferSize)
	}

	return aio
}

func (aio *AIO) portAudio(numberOfChannels, bufferSize int) {
	portaudio.Initialize()
	defer portaudio.Terminate()
	out := make([]float32, bufferSize)
	stream, err := portaudio.OpenDefaultStream(0, numberOfChannels, 44100, len(out), &out)
	if err != nil {
		log.Fatal(err)
	}
	defer stream.Close()

	if err := stream.Start(); err != nil {
		log.Fatal(err)
	}

	for {
		buf := <-aio.Audio
		f64ToF32Copy(out, buf.Data)
		// write to the stream
		if err := stream.Write(); err != nil {
			log.Printf("error writing to stream : %v\n", err)
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
