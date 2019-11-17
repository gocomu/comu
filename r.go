package main

import (
	"log"

	"github.com/go-audio/generator/euclidean"
)

func r() {
	output := euclidean.Rhythm(2, 20)
	log.Println(output)
}
