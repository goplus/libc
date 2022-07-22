package libc

func __math_invalidf(x float32) float32 {
	return (x - x) / (x - x)
}
