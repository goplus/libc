package libc

var S1_cgos____sin float64 = -0.16666666666666632
var S2_cgos____sin float64 = 0.0083333333333224895
var S3_cgos____sin float64 = -1.9841269829857949e-4
var S4_cgos____sin float64 = 2.7557313707070068e-6
var S5_cgos____sin float64 = -2.5050760253406863e-8
var S6_cgos____sin float64 = 1.5896909952115501e-10

func __sin(x float64, y float64, iy int32) float64 {
	var z float64
	var r float64
	var v float64
	var w float64
	z = x * x
	w = z * z
	r = S2_cgos____sin + z*(S3_cgos____sin+z*S4_cgos____sin) + z*w*(S5_cgos____sin+z*S6_cgos____sin)
	v = z * x
	if iy == int32(0) {
		return x + v*(S1_cgos____sin+z*r)
	} else {
		return x - (z*(0.5*y-v*r) - y - v*S1_cgos____sin)
	}
	return 0
}
