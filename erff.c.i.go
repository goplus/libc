package libc

import unsafe "unsafe"

var _cgos_erx__erff float32 = float32(0.84506291151000001)
var _cgos_efx8__erff float32 = float32(1.027033329)
var _cgos_pp0__erff float32 = float32(0.12837916613)
var _cgos_pp1__erff float32 = float32(-0.32504209876000001)
var _cgos_pp2__erff float32 = float32(-0.028481749818000002)
var _cgos_pp3__erff float32 = float32(-0.0057702702470000004)
var _cgos_pp4__erff float32 = float32(-2.3763017451999999e-5)
var _cgos_qq1__erff float32 = float32(0.39791721106)
var _cgos_qq2__erff float32 = float32(0.065022252500000002)
var _cgos_qq3__erff float32 = float32(0.0050813062117000003)
var _cgos_qq4__erff float32 = float32(1.3249473704e-4)
var _cgos_qq5__erff float32 = float32(-3.9602282412999997e-6)
var _cgos_pa0__erff float32 = float32(-0.0023621185682999998)
var _cgos_pa1__erff float32 = float32(0.41485610604000001)
var _cgos_pa2__erff float32 = float32(-0.37220788001999999)
var _cgos_pa3__erff float32 = float32(0.31834661960999999)
var _cgos_pa4__erff float32 = float32(-0.11089469491999999)
var _cgos_pa5__erff float32 = float32(0.035478305071999998)
var _cgos_pa6__erff float32 = float32(-0.0021663755177999998)
var _cgos_qa1__erff float32 = float32(0.10642088205)
var _cgos_qa2__erff float32 = float32(0.54039794207000003)
var _cgos_qa3__erff float32 = float32(0.071828655899000005)
var _cgos_qa4__erff float32 = float32(0.12617121637000001)
var _cgos_qa5__erff float32 = float32(0.013637083583)
var _cgos_qa6__erff float32 = float32(0.011984500103)
var _cgos_ra0__erff float32 = float32(-0.0098649440333000004)
var _cgos_ra1__erff float32 = float32(-0.6938585639)
var _cgos_ra2__erff float32 = float32(-10.558626175000001)
var _cgos_ra3__erff float32 = float32(-62.375331879000001)
var _cgos_ra4__erff float32 = float32(-162.39666747999999)
var _cgos_ra5__erff float32 = float32(-184.60508727999999)
var _cgos_ra6__erff float32 = float32(-81.287437439000001)
var _cgos_ra7__erff float32 = float32(-9.8143291473000004)
var _cgos_sa1__erff float32 = float32(19.651271820000002)
var _cgos_sa2__erff float32 = float32(137.65776062)
var _cgos_sa3__erff float32 = float32(434.56588744999999)
var _cgos_sa4__erff float32 = float32(645.38726807)
var _cgos_sa5__erff float32 = float32(429.00814818999999)
var _cgos_sa6__erff float32 = float32(108.63500214)
var _cgos_sa7__erff float32 = float32(6.5702495575000004)
var _cgos_sa8__erff float32 = float32(-0.060424413532000003)
var _cgos_rb0__erff float32 = float32(-0.0098649431019999997)
var _cgos_rb1__erff float32 = float32(-0.79928326607)
var _cgos_rb2__erff float32 = float32(-17.757955550999998)
var _cgos_rb3__erff float32 = float32(-160.63638305999999)
var _cgos_rb4__erff float32 = float32(-637.56646728999999)
var _cgos_rb5__erff float32 = float32(-1025.0950928)
var _cgos_rb6__erff float32 = float32(-483.51919556000001)
var _cgos_sb1__erff float32 = float32(30.338060379000002)
var _cgos_sb2__erff float32 = float32(325.79251098999998)
var _cgos_sb3__erff float32 = float32(1536.7296143000001)
var _cgos_sb4__erff float32 = float32(3199.8581543)
var _cgos_sb5__erff float32 = float32(2553.0502929999998)
var _cgos_sb6__erff float32 = float32(474.52853393999999)
var _cgos_sb7__erff float32 = float32(-22.440952300999999)

func _cgos_erfc1__erff(x float32) float32 {
	var s float32
	var P float32
	var Q float32
	s = Fabsf(x) - float32(int32(1))
	P = _cgos_pa0__erff + s*(_cgos_pa1__erff+s*(_cgos_pa2__erff+s*(_cgos_pa3__erff+s*(_cgos_pa4__erff+s*(_cgos_pa5__erff+s*_cgos_pa6__erff)))))
	Q = float32(int32(1)) + s*(_cgos_qa1__erff+s*(_cgos_qa2__erff+s*(_cgos_qa3__erff+s*(_cgos_qa4__erff+s*(_cgos_qa5__erff+s*_cgos_qa6__erff)))))
	return float32(int32(1)) - _cgos_erx__erff - P/Q
}
func _cgos_erfc2__erff(ix uint32, x float32) float32 {
	var s float32
	var R float32
	var S float32
	var z float32
	if ix < uint32(1067450368) {
		return _cgos_erfc1__erff(x)
	}
	x = Fabsf(x)
	s = float32(int32(1)) / (x * x)
	if ix < uint32(1077336941) {
		R = _cgos_ra0__erff + s*(_cgos_ra1__erff+s*(_cgos_ra2__erff+s*(_cgos_ra3__erff+s*(_cgos_ra4__erff+s*(_cgos_ra5__erff+s*(_cgos_ra6__erff+s*_cgos_ra7__erff))))))
		S = 1 + s*(_cgos_sa1__erff+s*(_cgos_sa2__erff+s*(_cgos_sa3__erff+s*(_cgos_sa4__erff+s*(_cgos_sa5__erff+s*(_cgos_sa6__erff+s*(_cgos_sa7__erff+s*_cgos_sa8__erff)))))))
	} else {
		R = _cgos_rb0__erff + s*(_cgos_rb1__erff+s*(_cgos_rb2__erff+s*(_cgos_rb3__erff+s*(_cgos_rb4__erff+s*(_cgos_rb5__erff+s*_cgos_rb6__erff)))))
		S = 1 + s*(_cgos_sb1__erff+s*(_cgos_sb2__erff+s*(_cgos_sb3__erff+s*(_cgos_sb4__erff+s*(_cgos_sb5__erff+s*(_cgos_sb6__erff+s*_cgos_sb7__erff))))))
	}
	for {
		ix = *(*uint32)(unsafe.Pointer(&_cgoz_18_erff{x}))
		if true {
			break
		}
	}
	for {
		z = *(*float32)(unsafe.Pointer(&_cgoz_19_erff{ix & uint32(4294959104)}))
		if true {
			break
		}
	}
	return Expf(-z*z-0.5625) * Expf((z-x)*(z+x)+R/S) / x
}

type _cgoz_18_erff struct {
	_f float32
}
type _cgoz_19_erff struct {
	_i uint32
}

func Erff(x float32) float32 {
	var r float32
	var s float32
	var z float32
	var y float32
	var ix uint32
	var sign int32
	for {
		ix = *(*uint32)(unsafe.Pointer(&_cgoz_20_erff{x}))
		if true {
			break
		}
	}
	sign = int32(ix >> int32(31))
	ix &= uint32(2147483647)
	if ix >= uint32(2139095040) {
		return float32(int32(1)-int32(2)*sign) + float32(int32(1))/x
	}
	if ix < uint32(1062731776) {
		if ix < uint32(830472192) {
			return 0.125 * (float32(int32(8))*x + _cgos_efx8__erff*x)
		}
		z = x * x
		r = _cgos_pp0__erff + z*(_cgos_pp1__erff+z*(_cgos_pp2__erff+z*(_cgos_pp3__erff+z*_cgos_pp4__erff)))
		s = float32(int32(1)) + z*(_cgos_qq1__erff+z*(_cgos_qq2__erff+z*(_cgos_qq3__erff+z*(_cgos_qq4__erff+z*_cgos_qq5__erff))))
		y = r / s
		return x + x*y
	}
	if ix < uint32(1086324736) {
		y = float32(int32(1)) - _cgos_erfc2__erff(ix, x)
	} else {
		y = float32(int32(1)) - 7.52316385e-37
	}
	return func() float32 {
		if sign != 0 {
			return -y
		} else {
			return y
		}
	}()
}

type _cgoz_20_erff struct {
	_f float32
}

func Erfcf(x float32) float32 {
	var r float32
	var s float32
	var z float32
	var y float32
	var ix uint32
	var sign int32
	for {
		ix = *(*uint32)(unsafe.Pointer(&_cgoz_21_erff{x}))
		if true {
			break
		}
	}
	sign = int32(ix >> int32(31))
	ix &= uint32(2147483647)
	if ix >= uint32(2139095040) {
		return float32(int32(2)*sign) + float32(int32(1))/x
	}
	if ix < uint32(1062731776) {
		if ix < uint32(595591168) {
			return 1 - x
		}
		z = x * x
		r = _cgos_pp0__erff + z*(_cgos_pp1__erff+z*(_cgos_pp2__erff+z*(_cgos_pp3__erff+z*_cgos_pp4__erff)))
		s = 1 + z*(_cgos_qq1__erff+z*(_cgos_qq2__erff+z*(_cgos_qq3__erff+z*(_cgos_qq4__erff+z*_cgos_qq5__erff))))
		y = r / s
		if sign != 0 || ix < uint32(1048576000) {
			return 1 - (x + x*y)
		}
		return 0.5 - (x - 0.5 + x*y)
	}
	if ix < uint32(1105199104) {
		return func() float32 {
			if sign != 0 {
				return float32(int32(2)) - _cgos_erfc2__erff(ix, x)
			} else {
				return _cgos_erfc2__erff(ix, x)
			}
		}()
	}
	return func() float32 {
		if sign != 0 {
			return float32(int32(2)) - 7.52316385e-37
		} else {
			return 7.52316385e-37 * 7.52316385e-37
		}
	}()
}

type _cgoz_21_erff struct {
	_f float32
}
