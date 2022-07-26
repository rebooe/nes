# nes
Openai proposes [a simple evolutionary strategy](https://openai.com/blog/evolution-strategies/) based on natural gradients 

# Installation
>go get -u github.com/rebooe/nes  
>go version >= 1.18 

# Usage
    package main

    import (
        "log"

        "github.com/rebooe/nes"
    )

    func loss[T nes.Float](w []T) T {
        var (
            sum T
            r   = []T{1, 0.3, 0.9, -1, -0.5}
        )
        for i := 0; i < len(w); i++ {
            sum += (r[i] - w[i]) * (r[i] - w[i])
        }
        return -sum / T(len(w))
    }

    func main() {
        w := make([]float32, 5)
        es := nes.NewNes[float32]()

        es.Run(loss[float32], w, 100)
        log.Printf("w: %v\n, R: %v\n", w, loss(w))
    }