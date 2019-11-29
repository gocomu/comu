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

type AudioIO struct {
	//BufChan          chan *audio.FloatBuffer
	numberOfChannels int
	bufferSize       int
	stream           *portaudio.Stream
	Out              []float32
}

func NewAudioIO(audioOutput out, numberOfChannels, bufferSize int) *AudioIO {
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
}

func (aio *AudioIO) PortAudioOut(out []float32, buf *audio.FloatBuffer) {
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
