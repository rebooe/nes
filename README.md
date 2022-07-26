# nes
Openai proposes [a simple evolutionary strategy](https://openai.com/blog/evolution-strategies/) based on natural gradients 

# Installation
>go get github.com/rebooe/nes

# Usage
    package main

    import (
        "log"
        "github.com/rebooe/nes"
    )

    func main() {
        nes.NewNes[float32]()
        w := make([]float32, 5)
	    nes.Run(loss, w, 100)
        log.Printf("w: %v\n, R: %v\n", w, f(w))
    }

    func loss[T float32](w []T) T {
        var (
            sum T
            r = []float32{1, 0.3, 0.9, -1, -0.5}
        )
		for i := 0; i < n; i++ {
			sum += (r[i] - w[i]) * (r[i] - w[i])
		}
		return -sum / T(n)
    }