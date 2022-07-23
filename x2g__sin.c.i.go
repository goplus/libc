package libc

var _cgos_S1____sin float64 = -0.16666666666666632
var _cgos_S2____sin float64 = 0.0083333333333224895
var _cgos_S3____sin float64 = -1.9841269829857949e-4
var _cgos_S4____sin float64 = 2.7557313707070068e-6
var _cgos_S5____sin float64 = -2.5050760253406863e-8
var _cgos_S6____sin float64 = 1.5896909952115501e-10

func __sin(x float64, y float64, iy int32) float64 {
	var z float64
	var r float64
	var v float64
	var w float64
	z = x * x
	w = z * z
	r = _cgos_S2____sin + z*(_cgos_S3____sin+z*_cgos_S4____sin) + z*w*(_cgos_S5____sin+z*_cgos_S6____sin)
	v = z * x
	if iy == int32(0) {
		return x + v*(_cgos_S1____sin+z*r)
	} else {
		return x - (z*(0.5*y-v*r) - y - v*_cgos_S1____sin)
	}
	return 0
}
