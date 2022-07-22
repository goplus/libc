package libc

func Scalblnl(x float64, n int64) float64 {
	return float64(Scalbln(float64(x), n))
}
