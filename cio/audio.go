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
	BS64 bufferSize = 64 << iota
	BS128
	BS256
	BS512
	BS1024
	BS2048
	BS4096
	BS8192
)

type out int

const (
	PortAudio = out(iota)
	Oto
)

type AudioIO struct {
	//BufChan          chan *audio.FloatBuffer
	numberOfChannels int
	bufferSize       bufferSize
	stream           *portaudio.Stream
	Out              []float32
}

func NewAudioIO(audioOutput out, numberOfChannels int, bufferSize bufferSize) *AudioIO {
	aio := &AudioIO{
		//BufChan:          make(chan *audio.FloatBuffer),
		numberOfChannels: numberOfChannels,
		bufferSize:       bufferSize,
		Out:              make([]float32, bufferSize),
	}
	switch audioOutput {
	case PortAudio:
		aio.portAudio()

	case Oto:
		// go aio.oto()
	}

	return aio
}

func (aio *AudioIO) portAudio() {
	portaudio.Initialize()
	//defer portaudio.Terminate()
	stream, err := portaudio.OpenDefaultStream(0, aio.numberOfChannels, 44100, len(aio.Out), &aio.Out)
	if err != nil {
		log.Fatal(err)
	}
	aio.stream = stream
	//defer stream.Close()

	if err := stream.Start(); err != nil {
		log.Fatal(err)
	}
	fmt.Println("COMU initializing")
	time.Sleep(1 * time.Second)
}

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
