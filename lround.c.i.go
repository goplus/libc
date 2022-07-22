package libc

func Lround(x float64) int64 {
	return int64(Round(x))
}
