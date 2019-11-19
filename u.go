package comu

import (
	"fmt"

	"github.com/mjibson/go-dsp/fft"
)

func fftt() {
	fmt.Println(fft.FFTReal([]float64{1, 2, 3}))
}

// output := euclidean.Rhythm(2, 20)
// log.Println(output)
