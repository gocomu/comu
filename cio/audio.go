package cio

import (
	"fmt"
	"log"
	"time"

	"github.com/go-audio/audio"
	"github.com/gordonklaus/portaudio"
)

type bufferSize int

const (
	// BS64 bufferSize
	BS64 bufferSize = 64 << iota
	// BS128 bufferSize
	BS128
	// BS256 bufferSize
	BS256
	// BS512 bufferSize
	BS512
	// BS1024 bufferSize
	BS1024
	// BS2048 bufferSize
	BS2048
	// BS4096 bufferSize
	BS4096
	// BS8192 bufferSize
	BS8192
)

type out int

const (
	// PortAudio out
	PortAudio = out(iota)
	// Oto out
	Oto
)

// AudioIO holds audio input/output information
type AudioIO struct {
	//BufChan          chan *audio.FloatBuffer
	numChan    int
	bufferSize bufferSize
	stream     *portaudio.Stream
	Out        []float32
}

// NewAudioIO starts sound output stream
func NewAudioIO(audioOutput out, numberOfChannels int, bs bufferSize) *AudioIO {
	fmt.Println(`
	COMU initializing
	`)

	aio := &AudioIO{
		//BufChan:          make(chan *audio.FloatBuffer),
		numChan:    numberOfChannels,
		bufferSize: bs,
		Out:        make([]float32, bs),
	}

	switch audioOutput {
	case PortAudio:
		aio.portAudio()

	case Oto:
		// aio.oto()
	}

	return aio
}

func (aio *AudioIO) portAudio() {
	portaudio.Initialize()
	//defer portaudio.Terminate()
	stream, err := portaudio.OpenDefaultStream(0, aio.numChan, 44100, len(aio.Out), &aio.Out)
	if err != nil {
		log.Fatal(err)
	}
	aio.stream = stream
	//defer stream.Close()

	if err := stream.Start(); err != nil {
		log.Fatal(err)
	}

	time.Sleep(1 * time.Second)
	fmt.Println(`
	COMU started
	`)
}

// PortAudioOut takes a buffer and writes it to PortAudio's stream
func (aio *AudioIO) PortAudioOut(buf *audio.FloatBuffer) {
	// portaudio doesn't support float64 so we need to copy our data over to the
	// destination buffer.
	for i := range buf.Data {
		aio.Out[i] = float32(buf.Data[i])
	}

	// write to the stream
	if err := aio.stream.Write(); err != nil {
		log.Printf("error writing to stream : %v\n", err)
	}
}
