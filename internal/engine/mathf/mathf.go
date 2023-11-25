package mathf

import (
	"math"
)

const (
	Pi      float32 = 3.14159265358979323846264338327950288419716939937510582097494459
	Epsilon float32 = 1e-6
)

func Floor(x float32) float32 {
	return float32(int(x))
}

func Sin(x float32) float32 {
	return float32(math.Sin(float64(x)))
}

func Cos(x float32) float32 {
	return float32(math.Cos(float64(x)))
}

func Tan(x float32) float32 {
	return float32(math.Tan(float64(x)))
}
