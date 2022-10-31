package libc

func __math_divzerof(sign uint32) float32 {
	return fp_barrierf(func() float32 {
		if sign != 0 {
			return -1.0
		} else {
			return 1.0
		}
	}()) / 0.0
}
