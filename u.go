package main

import (
        "fmt"
        
        "github.com/mjibson/go-dsp/fft"
)

func fftt() {
        fmt.Println(fft.FFTReal([]float64 {1, 2, 3}))
}