package nes

import (
	"math"
	"math/rand"
)

type RandomEs[T Float] struct {
	sigma float64 // 学习率
	droup float64 // 抑制率

	minR float64
}

func NewRandomEs[T Float](sigma float64, droup float64) *RandomEs[T] {
	return &RandomEs[T]{
		sigma: 0.5,
		droup: 0.5,
		minR:  math.MaxFloat32,
	}
}

func (n *RandomEs[T]) Run(f LossFunc[T], w []T) {
	// 生成的参数值
	wTry := make([]T, len(w))

	// 生成随机噪声样本
	lenTry := int(float64(len(wTry)) * n.droup)
	for i := 0; i < lenTry; i++ {
		wTry[i] = T(n.sigma * (rand.Float64()*2 - 1))
	}
	rand.Shuffle(len(wTry), func(i, j int) {
		wTry[i], wTry[j] = wTry[j], wTry[i]
	})
	for i := 0; i < len(w); i++ {
		wTry[i] += w[i]
	}

	R := f(wTry)
	if R < T(n.minR) {
		n.minR = float64(R)
		copy(w, wTry)
	}
}
