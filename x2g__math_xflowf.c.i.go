package libc

func __math_xflowf(sign uint32, y float32) float32 {
	return eval_as_float(fp_barrierf(func() float32 {
		if sign != 0 {
			return -y
		} else {
			return y
		}
	}()) * y)
}
