package libc

func Nexttowardl(x float64, y float64) float64 {
	return Nextafterl(x, y)
}
