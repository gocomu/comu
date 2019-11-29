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
	BufChan          chan *audio.FloatBuffer
	numberOfChannels int
	bufferSize       int
	stream           *portaudio.Stream
	//Out              []float32
}

func NewAudioIO(audioOutput out, numberOfChannels, bufferSize int, out []float32) *AudioIO {
	aio := &AudioIO{
		BufChan:          make(chan *audio.FloatBuffer),
		numberOfChannels: numberOfChannels,
		bufferSize:       bufferSize,
		//Out:              make([]float32, bufferSize),
	}
	switch audioOutput {
	case PortAudio:
		aio.portAudio(out)

	case Oto:
		// go aio.oto()
	}

	return aio
}

func (aio *AudioIO) portAudio(out []float32) {
	portaudio.Initialize()
	//defer portaudio.Terminate()
	//out := make([]float32, aio.bufferSize)
	// stream, err := portaudio.OpenDefaultStream(0, aio.numberOfChannels, 44100, len(aio.Out), &aio.Out)
	stream, err := portaudio.OpenDefaultStream(0, aio.numberOfChannels, 44100, len(out), &out)
	if err != nil {
		log.Fatal(err)
	}
	aio.stream = stream
	//defer stream.Close()

	if err := stream.Start(); err != nil {
		log.Fatal(err)
	}
}

func (aio *AudioIO) PortAudioFunc(out []float32, buf *audio.FloatBuffer) {
	// portaudio doesn't support float64 so we need to copy our data over to the
	// destination buffer.
	for i := range buf.Data {
		out[i] = float32(buf.Data[i])
	}

	// write to the stream
	if err := aio.stream.Write(); err != nil {
		log.Printf("error writing to stream : %v\n", err)
	}
}
