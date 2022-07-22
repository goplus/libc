package libc

func Llroundl(x float64) int64 {
	return int64(Roundl(x))
}
