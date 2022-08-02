package nes

import (
	"log"
	"math"
	"testing"
)

func Test_RandomRun(t *testing.T) {
	nes := RandomEs[float]{
		sigma: 0.5,
		droup: 0.5,
		minR:  math.MaxFloat32,
	}
	w := make([]float, 25)
	f := loss(len(w))

	for i := 0; i < 100; i++ {
		nes.Run(f, w)
		log.Printf("R: %v\n", f(w))
	}
}
