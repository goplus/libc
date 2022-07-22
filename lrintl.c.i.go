package libc

func Lrintl(x float64) int64 {
	return Lrint(float64(x))
}
