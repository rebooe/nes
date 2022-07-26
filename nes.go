package nes

import (
	"math"
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

type Float interface {
	~float32 | ~float64
}

type LossFunc[T Float] func([]T) T

type Nes[T Float] struct {
	npop  int // 人口数量
	alpha T   // 学习率
	sigma T   // 噪声衰减系数
}

func NewNes[T Float](op ...NesConfigFunc[T]) *Nes[T] {
	nes := &Nes[T]{
		npop:  10,
		alpha: 0.1,
		sigma: 1,
	}

	for _, f := range op {
		f(nes)
	}
	return nes
}

func (n *Nes[T]) Run(f LossFunc[T], w []T, num int) {
	// 随机噪声样本
	N := make([][]T, n.npop)
	for i := 0; i < len(N); i++ {
		N[i] = make([]T, len(w))
	}
	// 期望回报
	R := make([]T, n.npop)
	// 更新参数
	NT := make([]T, n.npop)
	// w 加上噪声后的值
	wTry := make([]T, len(w))

	for ; num > 0; num-- {
		// 生成随机噪声样本
		for i := 0; i < len(N); i++ {
			randn(N[i], n.sigma)
		}

		for i := 0; i < n.npop; i++ {
			for j := 0; j < len(w); j++ {
				wTry[j] = w[j] + N[i][j]
			}

			R[i] = f(wTry)
		}
		// 标准化
		standardize(R)

		for i := 0; i < len(w); i++ {
			// 获取转置矩阵值
			for j := 0; j < n.npop; j++ {
				NT[j] = N[j][i]
			}

			w[i] = w[i] + n.alpha*dot(NT, R)
		}
	}
}

func randn[T Float](x []T, sigma T) {
	for i := 0; i < len(x); i++ {
		r := T(rand.Float64()) * sigma
		x[i] = r
	}
}

// 奖励标准化
func standardize[T Float](R []T) {
	var sum T
	for i := 0; i < len(R); i++ {
		sum += R[i]
	}
	mean := sum / T(len(R))

	sum = 0
	for i := 0; i < len(R); i++ {
		sum += (mean - R[i]) * (mean - R[i])
	}
	std := T(math.Sqrt(float64(sum / T(len(R)))))

	for i := 0; i < len(R); i++ {
		R[i] = (R[i] - mean) / std
	}
}

func dot[T Float](a, b []T) T {
	var num T
	for i := 0; i < len(a); i++ {
		num += a[i] * b[i]
	}
	return num
}

type NesConfigFunc[T Float] func(n *Nes[T])

func WithPopulation[T Float](npop int) NesConfigFunc[T] {
	if npop <= 1 {
		panic("人口数量必须大于1")
	}
	return func(n *Nes[T]) {
		n.npop = npop
	}
}

func WithLearnRate[T Float](alpha T) NesConfigFunc[T] {
	return func(n *Nes[T]) {
		n.alpha = alpha
	}
}

func WithSigma[T Float](sigma T) NesConfigFunc[T] {
	return func(n *Nes[T]) {
		n.sigma = sigma
	}
}
