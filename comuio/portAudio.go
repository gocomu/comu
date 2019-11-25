package comuio

import (
	"log"

	"github.com/go-audio/audio"
	"github.com/gordonklaus/portaudio"
)

func PortAudio(bufchan chan *audio.FloatBuffer, bufferSize int) {
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
		buf := <-bufchan
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
