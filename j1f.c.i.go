package libc

import unsafe "unsafe"

var invsqrtpi_cgos__j1f float32 = float32(0.56418961287000002)
var tpi_cgos__j1f float32 = float32(0.63661974668999999)

func common_cgos__j1f(ix uint32, x float32, y1 int32, sign int32) float32 {
	var z float64
	var s float64
	var c float64
	var ss float64
	var cc float64
	s = float64(Sinf(x))
	if y1 != 0 {
		s = -s
	}
	c = float64(Cosf(x))
	cc = s - c
	if ix < uint32(2130706432) {
		ss = -s - c
		z = float64(Cosf(float32(int32(2)) * x))
		if s*c > float64(int32(0)) {
			cc = z / ss
		} else {
			ss = z / cc
		}
		if ix < uint32(1484783616) {
			if y1 != 0 {
				ss = -ss
			}
			cc = float64(ponef_cgos__j1f(x))*cc - float64(qonef_cgos__j1f(x))*ss
		}
	}
	if sign != 0 {
		cc = -cc
	}
	return float32(float64(invsqrtpi_cgos__j1f) * cc / float64(Sqrtf(x)))
}

var r00_cgos__j1f float32 = float32(-0.0625)
var r01_cgos__j1f float32 = float32(0.0014070566976)
var r02_cgos__j1f float32 = float32(-1.5995563443999999e-5)
var r03_cgos__j1f float32 = float32(4.9672799206999999e-8)
var s01_cgos__j1f float32 = float32(0.019153760746)
var s02_cgos__j1f float32 = float32(1.8594678841e-4)
var s03_cgos__j1f float32 = float32(1.1771846857000001e-6)
var s04_cgos__j1f float32 = float32(5.0463624389999999e-9)
var s05_cgos__j1f float32 = float32(1.2354227016e-11)

func J1f(x float32) float32 {
	var z float32
	var r float32
	var s float32
	var ix uint32
	var sign int32
	for {
		ix = *(*uint32)(unsafe.Pointer(&_cgoz_18_j1f{x}))
		if true {
			break
		}
	}
	sign = int32(ix >> int32(31))
	ix &= uint32(2147483647)
	if ix >= uint32(2139095040) {
		return float32(int32(1)) / (x * x)
	}
	if ix >= uint32(1073741824) {
		return common_cgos__j1f(ix, Fabsf(x), int32(0), sign)
	}
	if ix >= uint32(956301312) {
		z = x * x
		r = z * (r00_cgos__j1f + z*(r01_cgos__j1f+z*(r02_cgos__j1f+z*r03_cgos__j1f)))
		s = float32(int32(1)) + z*(s01_cgos__j1f+z*(s02_cgos__j1f+z*(s03_cgos__j1f+z*(s04_cgos__j1f+z*s05_cgos__j1f))))
		z = 0.5 + r/s
	} else {
		z = float32(0.5)
	}
	return z * x
}

type _cgoz_18_j1f struct {
	_f float32
}

var U0_cgos__j1f [5]float32 = [5]float32{float32(-0.19605709612), float32(0.050443872809), float32(-0.0019125689287000001), float32(2.3525259166e-5), float32(-9.1909917899000003e-8)}
var V0_cgos__j1f [5]float32 = [5]float32{float32(0.019916731864), float32(2.025525755e-4), float32(1.3560879779e-6), float32(6.227414584e-9), float32(1.6655924903e-11)}

func Y1f(x float32) float32 {
	var z float32
	var u float32
	var v float32
	var ix uint32
	for {
		ix = *(*uint32)(unsafe.Pointer(&_cgoz_19_j1f{x}))
		if true {
			break
		}
	}
	if ix&uint32(2147483647) == uint32(0) {
		return func() float32 {
			return float32(-1)
		}() / 0
	}
	if ix>>int32(31) != 0 {
		return func() float32 {
			return float32(int32(0))
		}() / 0
	}
	if ix >= uint32(2139095040) {
		return float32(int32(1)) / x
	}
	if ix >= uint32(1073741824) {
		return common_cgos__j1f(ix, x, int32(1), int32(0))
	}
	if ix < uint32(855638016) {
		return -tpi_cgos__j1f / x
	}
	z = x * x
	u = *(*float32)(unsafe.Pointer(uintptr(unsafe.Pointer((*float32)(unsafe.Pointer(&U0_cgos__j1f)))) + uintptr(int32(0))*4)) + z*(*(*float32)(unsafe.Pointer(uintptr(unsafe.Pointer((*float32)(unsafe.Pointer(&U0_cgos__j1f)))) + uintptr(int32(1))*4))+z*(*(*float32)(unsafe.Pointer(uintptr(unsafe.Pointer((*float32)(unsafe.Pointer(&U0_cgos__j1f)))) + uintptr(int32(2))*4))+z*(*(*float32)(unsafe.Pointer(uintptr(unsafe.Pointer((*float32)(unsafe.Pointer(&U0_cgos__j1f)))) + uintptr(int32(3))*4))+z**(*float32)(unsafe.Pointer(uintptr(unsafe.Pointer((*float32)(unsafe.Pointer(&U0_cgos__j1f)))) + uintptr(int32(4))*4)))))
	v = 1 + z*(*(*float32)(unsafe.Pointer(uintptr(unsafe.Pointer((*float32)(unsafe.Pointer(&V0_cgos__j1f)))) + uintptr(int32(0))*4))+z*(*(*float32)(unsafe.Pointer(uintptr(unsafe.Pointer((*float32)(unsafe.Pointer(&V0_cgos__j1f)))) + uintptr(int32(1))*4))+z*(*(*float32)(unsafe.Pointer(uintptr(unsafe.Pointer((*float32)(unsafe.Pointer(&V0_cgos__j1f)))) + uintptr(int32(2))*4))+z*(*(*float32)(unsafe.Pointer(uintptr(unsafe.Pointer((*float32)(unsafe.Pointer(&V0_cgos__j1f)))) + uintptr(int32(3))*4))+z**(*float32)(unsafe.Pointer(uintptr(unsafe.Pointer((*float32)(unsafe.Pointer(&V0_cgos__j1f)))) + uintptr(int32(4))*4))))))
	return x*(u/v) + tpi_cgos__j1f*(J1f(x)*Logf(x)-1/x)
}

type _cgoz_19_j1f struct {
	_f float32
}

var pr8_cgos__j1f [6]float32 = [6]float32{float32(0), float32(0.1171875), float32(13.239480972000001), float32(412.05184937000001), float32(3874.7453612999998), float32(7914.4794922000001)}
var ps8_cgos__j1f [5]float32 = [5]float32{float32(114.20736694), float32(3650.9309082), float32(36956.207030999998), float32(97602.796875), float32(30804.271484000001)}
var pr5_cgos__j1f [6]float32 = [6]float32{float32(1.3199052093999999e-11), float32(0.11718749254999999), float32(6.8027510642999998), float32(108.30818176), float32(517.63616943), float32(528.71520996000004)}
var ps5_cgos__j1f [5]float32 = [5]float32{float32(59.280597686999997), float32(991.40142821999996), float32(5353.2670897999997), float32(7844.6904297000001), float32(1504.046875)}
var pr3_cgos__j1f [6]float32 = [6]float32{float32(3.0250391081000001e-9), float32(0.1171868667), float32(3.932977438), float32(35.119403839), float32(91.055007935000006), float32(48.559066772000001)}
var ps3_cgos__j1f [5]float32 = [5]float32{float32(34.791309357000003), float32(336.76245117000002), float32(1046.8714600000001), float32(890.81134033000001), float32(103.78793335)}
var pr2_cgos__j1f [6]float32 = [6]float32{float32(1.0771083225e-7), float32(0.11717621982), float32(2.3685150145999998), float32(12.242610931), float32(17.693971634), float32(5.0735230445999999)}
var ps2_cgos__j1f [5]float32 = [5]float32{float32(21.436485291), float32(125.29022980000001), float32(232.27647400000001), float32(117.67937469), float32(8.3646392822000006)}

func ponef_cgos__j1f(x float32) float32 {
	var p *float32
	var q *float32
	var z float32
	var r float32
	var s float32
	var ix uint32
	for {
		ix = *(*uint32)(unsafe.Pointer(&_cgoz_20_j1f{x}))
		if true {
			break
		}
	}
	ix &= uint32(2147483647)
	if ix >= uint32(1090519040) {
		p = (*float32)(unsafe.Pointer(&pr8_cgos__j1f))
		q = (*float32)(unsafe.Pointer(&ps8_cgos__j1f))
	} else if ix >= uint32(1083274219) {
		p = (*float32)(unsafe.Pointer(&pr5_cgos__j1f))
		q = (*float32)(unsafe.Pointer(&ps5_cgos__j1f))
	} else if ix >= uint32(1077336343) {
		p = (*float32)(unsafe.Pointer(&pr3_cgos__j1f))
		q = (*float32)(unsafe.Pointer(&ps3_cgos__j1f))
	} else {
		p = (*float32)(unsafe.Pointer(&pr2_cgos__j1f))
		q = (*float32)(unsafe.Pointer(&ps2_cgos__j1f))
	}
	z = 1 / (x * x)
	r = *(*float32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + uintptr(int32(0))*4)) + z*(*(*float32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + uintptr(int32(1))*4))+z*(*(*float32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + uintptr(int32(2))*4))+z*(*(*float32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + uintptr(int32(3))*4))+z*(*(*float32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + uintptr(int32(4))*4))+z**(*float32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + uintptr(int32(5))*4))))))
	s = 1 + z*(*(*float32)(unsafe.Pointer(uintptr(unsafe.Pointer(q)) + uintptr(int32(0))*4))+z*(*(*float32)(unsafe.Pointer(uintptr(unsafe.Pointer(q)) + uintptr(int32(1))*4))+z*(*(*float32)(unsafe.Pointer(uintptr(unsafe.Pointer(q)) + uintptr(int32(2))*4))+z*(*(*float32)(unsafe.Pointer(uintptr(unsafe.Pointer(q)) + uintptr(int32(3))*4))+z**(*float32)(unsafe.Pointer(uintptr(unsafe.Pointer(q)) + uintptr(int32(4))*4))))))
	return 1 + r/s
}

type _cgoz_20_j1f struct {
	_f float32
}

var qr8_cgos__j1f [6]float32 = [6]float32{float32(0), float32(-0.1025390625), float32(-16.271753311000001), float32(-759.60174560999997), float32(-11849.806640999999), float32(-48438.511719000002)}
var qs8_cgos__j1f [6]float32 = [6]float32{float32(161.39537048), float32(7825.3862305000002), float32(133875.34375), float32(719657.75), float32(666601.25), float32(-294490.25)}
var qr5_cgos__j1f [6]float32 = [6]float32{float32(-2.0897993404999998e-11), float32(-0.1025390476), float32(-8.0564479828), float32(-183.66960144000001), float32(-1373.1937256000001), float32(-2612.4443359000002)}
var qs5_cgos__j1f [6]float32 = [6]float32{float32(81.276550292999999), float32(1991.7987060999999), float32(17468.484375), float32(49851.425780999998), float32(27948.074218999998), float32(-4719.1835938000004)}
var qr3_cgos__j1f [6]float32 = [6]float32{float32(-5.0783124372e-9), float32(-0.10253783314999999), float32(-4.6101160049000001), float32(-57.847221374999997), float32(-228.24453735), float32(-219.21012877999999)}
var qs3_cgos__j1f [6]float32 = [6]float32{float32(47.665153502999999), float32(673.86511229999996), float32(3380.1528320000002), float32(5547.7290039), float32(1903.1191406), float32(-135.20118712999999)}
var qr2_cgos__j1f [6]float32 = [6]float32{float32(-1.7838172539000001e-7), float32(-0.10251704603), float32(-2.7522056102999999), float32(-19.663616180000002), float32(-42.325313567999999), float32(-21.371921538999999)}
var qs2_cgos__j1f [6]float32 = [6]float32{float32(29.533363342000001), float32(252.98155212), float32(757.50280762), float32(739.39318848000005), float32(155.94900512999999), float32(-4.9594988823000001)}

func qonef_cgos__j1f(x float32) float32 {
	var p *float32
	var q *float32
	var s float32
	var r float32
	var z float32
	var ix uint32
	for {
		ix = *(*uint32)(unsafe.Pointer(&_cgoz_21_j1f{x}))
		if true {
			break
		}
	}
	ix &= uint32(2147483647)
	if ix >= uint32(1090519040) {
		p = (*float32)(unsafe.Pointer(&qr8_cgos__j1f))
		q = (*float32)(unsafe.Pointer(&qs8_cgos__j1f))
	} else if ix >= uint32(1083274219) {
		p = (*float32)(unsafe.Pointer(&qr5_cgos__j1f))
		q = (*float32)(unsafe.Pointer(&qs5_cgos__j1f))
	} else if ix >= uint32(1077336343) {
		p = (*float32)(unsafe.Pointer(&qr3_cgos__j1f))
		q = (*float32)(unsafe.Pointer(&qs3_cgos__j1f))
	} else {
		p = (*float32)(unsafe.Pointer(&qr2_cgos__j1f))
		q = (*float32)(unsafe.Pointer(&qs2_cgos__j1f))
	}
	z = 1 / (x * x)
	r = *(*float32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + uintptr(int32(0))*4)) + z*(*(*float32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + uintptr(int32(1))*4))+z*(*(*float32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + uintptr(int32(2))*4))+z*(*(*float32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + uintptr(int32(3))*4))+z*(*(*float32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + uintptr(int32(4))*4))+z**(*float32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + uintptr(int32(5))*4))))))
	s = 1 + z*(*(*float32)(unsafe.Pointer(uintptr(unsafe.Pointer(q)) + uintptr(int32(0))*4))+z*(*(*float32)(unsafe.Pointer(uintptr(unsafe.Pointer(q)) + uintptr(int32(1))*4))+z*(*(*float32)(unsafe.Pointer(uintptr(unsafe.Pointer(q)) + uintptr(int32(2))*4))+z*(*(*float32)(unsafe.Pointer(uintptr(unsafe.Pointer(q)) + uintptr(int32(3))*4))+z*(*(*float32)(unsafe.Pointer(uintptr(unsafe.Pointer(q)) + uintptr(int32(4))*4))+z**(*float32)(unsafe.Pointer(uintptr(unsafe.Pointer(q)) + uintptr(int32(5))*4)))))))
	return (0.375 + r/s) / x
}

type _cgoz_21_j1f struct {
	_f float32
}
