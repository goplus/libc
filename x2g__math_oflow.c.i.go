package libc

func __math_oflow(sign uint32) float64 {
	return __math_xflow(sign, 3.1050361846014179e+231)
}
