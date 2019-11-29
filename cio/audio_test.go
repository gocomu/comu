package cio

import (
	"github.com/go-audio/audio"
	"testing"
)

var audioio *AudioIO

func TestAudioIO(t *testing.T) {
	t.Run("NewAudioIO", func(t *testing.T) {
		audioio = NewAudioIO(PortAudio, 2, BS2048)
	})

	t.Run("NewAudioIO", func(t *testing.T) {
		testbuf := &audio.FloatBuffer{
			Data:   make([]float64, BS2048),
			Format: audio.FormatStereo44100,
		}
		audioio.PortAudioOut(testbuf)
	})
}
