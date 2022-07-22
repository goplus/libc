package libc

func Ldexpf(x float32, n int32) float32 {
	return Scalbnf(x, n)
}
