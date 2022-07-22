package libc

func __math_xflow(sign uint32, y float64) float64 {
	return eval_as_double(fp_barrier(func() float64 {
		if sign != 0 {
			return -y
		} else {
			return y
		}
	}()) * y)
}
