package libc

var C1_cgos____cos float64 = 0.041666666666666602
var C2_cgos____cos float64 = -0.001388888888887411
var C3_cgos____cos float64 = 2.4801587289476729e-5
var C4_cgos____cos float64 = -2.7557314351390663e-7
var C5_cgos____cos float64 = 2.0875723212981748e-9
var C6_cgos____cos float64 = -1.1359647557788195e-11

func __cos(x float64, y float64) float64 {
	var hz float64
	var z float64
	var r float64
	var w float64
	z = x * x
	w = z * z
	r = z*(C1_cgos____cos+z*(C2_cgos____cos+z*C3_cgos____cos)) + w*w*(C4_cgos____cos+z*(C5_cgos____cos+z*C6_cgos____cos))
	hz = 0.5 * z
	w = 1 - hz
	return w + (1 - w - hz + (z*r - x*y))
}
