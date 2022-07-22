package libc

func __math_invalid(x float64) float64 {
	return (x - x) / (x - x)
}
