package libc

func Ldexpl(x float64, n int32) float64 {
	return Scalbnl(x, n)
}
