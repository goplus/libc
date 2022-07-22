package libc

func Ldexp(x float64, n int32) float64 {
	return Scalbn(x, n)
}
