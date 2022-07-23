package libc

import unsafe "unsafe"

var invsqrtpi_cgos__j0 float64 = 0.56418958354775628
var tpi_cgos__j0 float64 = 0.63661977236758138

func common_cgos__j0(ix uint32, x float64, y0 int32) float64 {
	var s float64
	var c float64
	var ss float64
	var cc float64
	var z float64
	s = Sin(x)
	c = Cos(x)
	if y0 != 0 {
		c = -c
	}
	cc = s + c
	if ix < uint32(2145386496) {
		ss = s - c
		z = -Cos(float64(int32(2)) * x)
		if s*c < float64(int32(0)) {
			cc = z / ss
		} else {
			ss = z / cc
		}
		if ix < uint32(1207959552) {
			if y0 != 0 {
				ss = -ss
			}
			cc = pzero_cgos__j0(x)*cc - qzero_cgos__j0(x)*ss
		}
	}
	return invsqrtpi_cgos__j0 * cc / Sqrt(x)
}

var R02_cgos__j0 float64 = 0.015624999999999995
var R03_cgos__j0 float64 = -1.8997929423885472e-4
var R04_cgos__j0 float64 = 1.8295404953270067e-6
var R05_cgos__j0 float64 = -4.6183268853210319e-9
var S01_cgos__j0 float64 = 0.015619102946489001
var S02_cgos__j0 float64 = 1.1692678466333745e-4
var S03_cgos__j0 float64 = 5.1354655020731811e-7
var S04_cgos__j0 float64 = 1.1661400333379e-9

func J0(x float64) float64 {
	var z float64
	var r float64
	var s float64
	var ix uint32
	for {
		ix = uint32(*(*uint64)(unsafe.Pointer(&_cgoz_18_j0{x})) >> int32(32))
		if true {
			break
		}
	}
	ix &= uint32(2147483647)
	if ix >= uint32(2146435072) {
		return float64(int32(1)) / (x * x)
	}
	x = Fabs(x)
	if ix >= uint32(1073741824) {
		return common_cgos__j0(ix, x, int32(0))
	}
	if ix >= uint32(1059061760) {
		z = x * x
		r = z * (R02_cgos__j0 + z*(R03_cgos__j0+z*(R04_cgos__j0+z*R05_cgos__j0)))
		s = float64(int32(1)) + z*(S01_cgos__j0+z*(S02_cgos__j0+z*(S03_cgos__j0+z*S04_cgos__j0)))
		return (float64(int32(1))+x/float64(int32(2)))*(float64(int32(1))-x/float64(int32(2))) + z*(r/s)
	}
	if ix >= uint32(939524096) {
		x = 0.25 * x * x
	}
	return float64(int32(1)) - x
}

type _cgoz_18_j0 struct {
	_f float64
}

var u00_cgos__j0 float64 = -0.073804295108687232
var u01_cgos__j0 float64 = 0.17666645250918112
var u02_cgos__j0 float64 = -0.01381856719455969
var u03_cgos__j0 float64 = 3.4745343209368365e-4
var u04_cgos__j0 float64 = -3.8140705372436416e-6
var u05_cgos__j0 float64 = 1.9559013703502292e-8
var u06_cgos__j0 float64 = -3.982051941321034e-11
var v01_cgos__j0 float64 = 0.01273048348341237
var v02_cgos__j0 float64 = 7.6006862735035325e-5
var v03_cgos__j0 float64 = 2.5915085184045781e-7
var v04_cgos__j0 float64 = 4.4111031133267547e-10

func Y0(x float64) float64 {
	var z float64
	var u float64
	var v float64
	var ix uint32
	var lx uint32
	for {
		var __u uint64 = *(*uint64)(unsafe.Pointer(&_cgoz_19_j0{x}))
		ix = uint32(__u >> int32(32))
		lx = uint32(__u)
		if true {
			break
		}
	}
	if ix<<int32(1)|lx == uint32(0) {
		return func() float64 {
			return float64(-1)
		}() / 0
	}
	if ix>>int32(31) != 0 {
		return func() float64 {
			return float64(int32(0))
		}() / 0
	}
	if ix >= uint32(2146435072) {
		return float64(int32(1)) / x
	}
	if ix >= uint32(1073741824) {
		return common_cgos__j0(ix, x, int32(1))
	}
	if ix >= uint32(1044381696) {
		z = x * x
		u = u00_cgos__j0 + z*(u01_cgos__j0+z*(u02_cgos__j0+z*(u03_cgos__j0+z*(u04_cgos__j0+z*(u05_cgos__j0+z*u06_cgos__j0)))))
		v = 1 + z*(v01_cgos__j0+z*(v02_cgos__j0+z*(v03_cgos__j0+z*v04_cgos__j0)))
		return u/v + tpi_cgos__j0*(J0(x)*Log(x))
	}
	return u00_cgos__j0 + tpi_cgos__j0*Log(x)
}

type _cgoz_19_j0 struct {
	_f float64
}

var pR8_cgos__j0 [6]float64 = [6]float64{0, -0.070312499999990036, -8.081670412753498, -257.06310567970485, -2485.2164100942882, -5253.0438049072955}
var pS8_cgos__j0 [5]float64 = [5]float64{116.53436461966818, 3833.7447536412183, 40597.857264847255, 116752.97256437592, 47627.728414673096}
var pR5_cgos__j0 [6]float64 = [6]float64{-1.141254646918945e-11, -0.070312494087359928, -4.1596106447058778, -67.674765226516726, -331.23129964917297, -346.43338836560491}
var pS5_cgos__j0 [5]float64 = [5]float64{60.753938269230034, 1051.2523059570458, 5978.9709433385578, 9625.4451435777446, 2406.0581592293911}
var pR3_cgos__j0 [6]float64 = [6]float64{-2.5470460177195192e-9, -0.070311961638148165, -2.4090322154952961, -21.965977473488309, -58.079170470173757, -31.44794705948885}
var pS3_cgos__j0 [5]float64 = [5]float64{35.856033805520973, 361.51398305030386, 1193.6078379211153, 1127.9967985690741, 173.58093081333575}
var pR2_cgos__j0 [6]float64 = [6]float64{-8.8753433303252641e-8, -0.070303099548362474, -1.4507384678095299, -7.6356961382352777, -11.193166886035675, -3.2336457935133534}
var pS2_cgos__j0 [5]float64 = [5]float64{22.220299753208881, 136.20679421821521, 270.47027865808349, 153.87539420832033, 14.657617694825619}

func pzero_cgos__j0(x float64) float64 {
	var p *float64
	var q *float64
	var z float64
	var r float64
	var s float64
	var ix uint32
	for {
		ix = uint32(*(*uint64)(unsafe.Pointer(&_cgoz_20_j0{x})) >> int32(32))
		if true {
			break
		}
	}
	ix &= uint32(2147483647)
	if ix >= uint32(1075838976) {
		p = (*float64)(unsafe.Pointer(&pR8_cgos__j0))
		q = (*float64)(unsafe.Pointer(&pS8_cgos__j0))
	} else if ix >= uint32(1074933387) {
		p = (*float64)(unsafe.Pointer(&pR5_cgos__j0))
		q = (*float64)(unsafe.Pointer(&pS5_cgos__j0))
	} else if ix >= uint32(1074191213) {
		p = (*float64)(unsafe.Pointer(&pR3_cgos__j0))
		q = (*float64)(unsafe.Pointer(&pS3_cgos__j0))
	} else {
		p = (*float64)(unsafe.Pointer(&pR2_cgos__j0))
		q = (*float64)(unsafe.Pointer(&pS2_cgos__j0))
	}
	z = 1 / (x * x)
	r = *(*float64)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + uintptr(int32(0))*8)) + z*(*(*float64)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + uintptr(int32(1))*8))+z*(*(*float64)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + uintptr(int32(2))*8))+z*(*(*float64)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + uintptr(int32(3))*8))+z*(*(*float64)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + uintptr(int32(4))*8))+z**(*float64)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + uintptr(int32(5))*8))))))
	s = 1 + z*(*(*float64)(unsafe.Pointer(uintptr(unsafe.Pointer(q)) + uintptr(int32(0))*8))+z*(*(*float64)(unsafe.Pointer(uintptr(unsafe.Pointer(q)) + uintptr(int32(1))*8))+z*(*(*float64)(unsafe.Pointer(uintptr(unsafe.Pointer(q)) + uintptr(int32(2))*8))+z*(*(*float64)(unsafe.Pointer(uintptr(unsafe.Pointer(q)) + uintptr(int32(3))*8))+z**(*float64)(unsafe.Pointer(uintptr(unsafe.Pointer(q)) + uintptr(int32(4))*8))))))
	return 1 + r/s
}

type _cgoz_20_j0 struct {
	_f float64
}

var qR8_cgos__j0 [6]float64 = [6]float64{0, 0.073242187499993505, 11.768206468225269, 557.67338025640186, 8859.1972075646863, 37014.626777688783}
var qS8_cgos__j0 [6]float64 = [6]float64{163.77602689568982, 8098.3449465644981, 142538.29141912048, 803309.2571195144, 840501.57981906051, -343899.29353786662}
var qR5_cgos__j0 [6]float64 = [6]float64{1.8408596359451553e-11, 0.073242176661268477, 5.8356350896205695, 135.11157728644983, 1027.243765961641, 1989.9778586460538}
var qS5_cgos__j0 [6]float64 = [6]float64{82.776610223653776, 2077.8141642139299, 18847.288778571809, 56751.112289494733, 35976.753842511447, -5354.3427560194477}
var qR3_cgos__j0 [6]float64 = [6]float64{4.3774101408973862e-9, 0.073241118004291145, 3.3442313751617072, 42.621844074541265, 170.8080913405656, 166.73394869665117}
var qS3_cgos__j0 [6]float64 = [6]float64{48.758872972458718, 709.68922105660602, 3704.1482262011136, 6460.4251675256892, 2516.3336892036896, -149.24745183615639}
var qR2_cgos__j0 [6]float64 = [6]float64{1.5044444488698327e-7, 0.073223426596307928, 1.99819174093816, 14.495602934788574, 31.666231750478154, 16.252707571092927}
var qS2_cgos__j0 [6]float64 = [6]float64{30.365584835521918, 269.34811860804984, 844.78375759532014, 882.93584511248855, 212.66638851179883, -5.3109549388266695}

func qzero_cgos__j0(x float64) float64 {
	var p *float64
	var q *float64
	var s float64
	var r float64
	var z float64
	var ix uint32
	for {
		ix = uint32(*(*uint64)(unsafe.Pointer(&_cgoz_21_j0{x})) >> int32(32))
		if true {
			break
		}
	}
	ix &= uint32(2147483647)
	if ix >= uint32(1075838976) {
		p = (*float64)(unsafe.Pointer(&qR8_cgos__j0))
		q = (*float64)(unsafe.Pointer(&qS8_cgos__j0))
	} else if ix >= uint32(1074933387) {
		p = (*float64)(unsafe.Pointer(&qR5_cgos__j0))
		q = (*float64)(unsafe.Pointer(&qS5_cgos__j0))
	} else if ix >= uint32(1074191213) {
		p = (*float64)(unsafe.Pointer(&qR3_cgos__j0))
		q = (*float64)(unsafe.Pointer(&qS3_cgos__j0))
	} else {
		p = (*float64)(unsafe.Pointer(&qR2_cgos__j0))
		q = (*float64)(unsafe.Pointer(&qS2_cgos__j0))
	}
	z = 1 / (x * x)
	r = *(*float64)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + uintptr(int32(0))*8)) + z*(*(*float64)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + uintptr(int32(1))*8))+z*(*(*float64)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + uintptr(int32(2))*8))+z*(*(*float64)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + uintptr(int32(3))*8))+z*(*(*float64)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + uintptr(int32(4))*8))+z**(*float64)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + uintptr(int32(5))*8))))))
	s = 1 + z*(*(*float64)(unsafe.Pointer(uintptr(unsafe.Pointer(q)) + uintptr(int32(0))*8))+z*(*(*float64)(unsafe.Pointer(uintptr(unsafe.Pointer(q)) + uintptr(int32(1))*8))+z*(*(*float64)(unsafe.Pointer(uintptr(unsafe.Pointer(q)) + uintptr(int32(2))*8))+z*(*(*float64)(unsafe.Pointer(uintptr(unsafe.Pointer(q)) + uintptr(int32(3))*8))+z*(*(*float64)(unsafe.Pointer(uintptr(unsafe.Pointer(q)) + uintptr(int32(4))*8))+z**(*float64)(unsafe.Pointer(uintptr(unsafe.Pointer(q)) + uintptr(int32(5))*8)))))))
	return (-0.125 + r/s) / x
}

type _cgoz_21_j0 struct {
	_f float64
}
