package libc

func Llround(x float64) int64 {
	return int64(Round(x))
}
