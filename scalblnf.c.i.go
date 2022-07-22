package libc

func Scalblnf(x float32, n int64) float32 {
	if n > int64(2147483647) {
		n = int64(2147483647)
	} else if n < int64(-2147483648) {
		n = int64(-2147483648)
	}
	return Scalbnf(x, int32(n))
}
