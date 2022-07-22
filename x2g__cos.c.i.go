package libc

var C1_cgo20___cos float64 = 0.041666666666666602
var C2_cgo21___cos float64 = -0.001388888888887411
var C3_cgo22___cos float64 = 2.4801587289476729e-5
var C4_cgo23___cos float64 = -2.7557314351390663e-7
var C5_cgo24___cos float64 = 2.0875723212981748e-9
var C6_cgo25___cos float64 = -1.1359647557788195e-11

func __cos(x float64, y float64) float64 {
	var hz float64
	var z float64
	var r float64
	var w float64
	z = x * x
	w = z * z
	r = z*(C1_cgo20___cos+z*(C2_cgo21___cos+z*C3_cgo22___cos)) + w*w*(C4_cgo23___cos+z*(C5_cgo24___cos+z*C6_cgo25___cos))
	hz = 0.5 * z
	w = 1 - hz
	return w + (1 - w - hz + (z*r - x*y))
}
