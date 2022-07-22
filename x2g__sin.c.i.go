package libc

var S1_cgo18___sin float64 = -0.16666666666666632
var S2_cgo19___sin float64 = 0.0083333333333224895
var S3_cgo20___sin float64 = -1.9841269829857949e-4
var S4_cgo21___sin float64 = 2.7557313707070068e-6
var S5_cgo22___sin float64 = -2.5050760253406863e-8
var S6_cgo23___sin float64 = 1.5896909952115501e-10

func __sin(x float64, y float64, iy int32) float64 {
	var z float64
	var r float64
	var v float64
	var w float64
	z = x * x
	w = z * z
	r = S2_cgo19___sin + z*(S3_cgo20___sin+z*S4_cgo21___sin) + z*w*(S5_cgo22___sin+z*S6_cgo23___sin)
	v = z * x
	if iy == int32(0) {
		return x + v*(S1_cgo18___sin+z*r)
	} else {
		return x - (z*(0.5*y-v*r) - y - v*S1_cgo18___sin)
	}
	return 0
}
