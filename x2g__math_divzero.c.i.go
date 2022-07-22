package libc

func __math_divzero(sign uint32) float64 {
	return fp_barrier(func() float64 {
		if sign != 0 {
			return -1
		} else {
			return 1
		}
	}()) / 0
}
