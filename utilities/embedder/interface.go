package embedder

import "github.com/go-audio/audio"

type ProjectLibrary interface {
	Read(...[]string) audio.IntBuffer
}
