package libc

func frexpl(x float64, e *int32) float64 {
	return float64(frexp(float64(x), e))
}
