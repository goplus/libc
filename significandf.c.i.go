package libc

func Significandf(x float32) float32 {
	return Scalbnf(x, -Ilogbf(x))
}
