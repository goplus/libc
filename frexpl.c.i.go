package libc

func Frexpl(x float64, e *int32) float64 {
	return float64(Frexp(float64(x), e))
}
