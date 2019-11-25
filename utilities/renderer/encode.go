package renderer

import (
	"fmt"
	"os"

	"github.com/go-audio/wav"
)

func encode() {
	f, err := os.Open("tiegine.wav")
	if err != nil {
		panic(fmt.Sprintf("couldn't open audio file - %v", err))
	}

	// Decode the original audio file
	// and collect audio content and information.
	d := wav.NewDecoder(f)
	buf, err := d.FullPCMBuffer()
	if err != nil {
		panic(err)
	}
	f.Close()
	fmt.Println("Old file ->", d)
	fmt.Println("Old file ->", buf)

	// Destination file
	out, err := os.Create("test.wav")
	if err != nil {
		panic(fmt.Sprintf("couldn't create output file - %v", err))
	}

	// setup the encoder and write all the frames
	e := wav.NewEncoder(out, 44100, 24, 2, 1)
	if err = e.Write(buf); err != nil {
		panic(err)
	}
	// close the encoder to make sure the headers are properly
	// set and the data is flushed.
	if err = e.Close(); err != nil {
		panic(err)
	}
	out.Close()

	// reopen to confirm things worked well
	out, err = os.Open("test.wav")
	if err != nil {
		panic(err)
	}
	d2 := wav.NewDecoder(out)
	d2.ReadInfo()
	fmt.Println("New file ->", d2)
	out.Close()

	// Output:
	// Old file -> Format: WAVE - 1 channels @ 22050 / 16 bits - Duration: 0.204172 seconds
	// New file -> Format: WAVE - 1 channels @ 22050 / 16 bits - Duration: 0.204172 seconds
}
