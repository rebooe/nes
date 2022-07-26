package nes

import (
	"log"
	"testing"
)

type float float32

func loss[T float](n int) LossFunc[T] {
	r := make([]T, n)
	for i := 0; i < n; i++ {
		r[i] = 1
	}

	return func(w []T) T {
		var sum T
		for i := 0; i < n; i++ {
			sum += (r[i] - w[i]) * (r[i] - w[i])
		}
		return -sum / T(n)
	}
}

func Test_Run(t *testing.T) {
	nes := NewNes[float](WithPopulation[float](50))
	w := make([]float, 100)
	f := loss(len(w))

	nes.Run(f, w, 10)
	log.Printf("R: %v\n", f(w))
}
