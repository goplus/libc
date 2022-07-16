package libc

import "math"

func __builtin_inf() float64 {
	return math.Inf(0)
}
func __builtin_inff() float32 {
	return float32(__builtin_inf())
}
func __builtin_nan(*int8) float64 {
	return math.NaN()
}
func __builtin_nanf(*int8) float32 {
	return float32(math.NaN())
}
