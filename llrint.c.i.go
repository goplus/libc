package libc

func Llrint(x float64) int64 {
	return int64(Rint(x))
}
