package libc

func Llrintl(x float64) int64 {
	return Llrint(float64(x))
}
