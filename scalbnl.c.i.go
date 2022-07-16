package libc

func Scalbnl(x float64, n int32) float64 {
	return float64(Scalbn(float64(x), n))
}
