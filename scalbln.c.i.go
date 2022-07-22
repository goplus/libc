package libc

func Scalbln(x float64, n int64) float64 {
	if n > int64(2147483647) {
		n = int64(2147483647)
	} else if n < int64(-2147483648) {
		n = int64(-2147483648)
	}
	return Scalbn(x, int32(n))
}
