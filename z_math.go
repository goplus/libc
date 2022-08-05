package libc

import "math"

func X__builtin_inf() float64 {
	return math.Inf(0)
}

func X__builtin_inff() float32 {
	return float32(X__builtin_inf())
}

func X__builtin_nan(*int8) float64 {
	return math.NaN()
}

func X__builtin_nanf(*int8) float32 {
	return float32(math.NaN())
}

/*
func Sin(x float64) float64 {
	return math.Sin(x)
}
*/
