package libc

func __math_uflow(sign uint32) float64 {
	return __math_xflow(sign, 1.2882297539194267e-231)
}
