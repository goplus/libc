package libc

func Llroundf(x float32) int64 {
	return int64(Roundf(x))
}
