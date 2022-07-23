package libc

import unsafe "unsafe"

var erx_cgos__erff float32 = float32(0.84506291151000001)
var efx8_cgos__erff float32 = float32(1.027033329)
var pp0_cgos__erff float32 = float32(0.12837916613)
var pp1_cgos__erff float32 = float32(-0.32504209876000001)
var pp2_cgos__erff float32 = float32(-0.028481749818000002)
var pp3_cgos__erff float32 = float32(-0.0057702702470000004)
var pp4_cgos__erff float32 = float32(-2.3763017451999999e-5)
var qq1_cgos__erff float32 = float32(0.39791721106)
var qq2_cgos__erff float32 = float32(0.065022252500000002)
var qq3_cgos__erff float32 = float32(0.0050813062117000003)
var qq4_cgos__erff float32 = float32(1.3249473704e-4)
var qq5_cgos__erff float32 = float32(-3.9602282412999997e-6)
var pa0_cgos__erff float32 = float32(-0.0023621185682999998)
var pa1_cgos__erff float32 = float32(0.41485610604000001)
var pa2_cgos__erff float32 = float32(-0.37220788001999999)
var pa3_cgos__erff float32 = float32(0.31834661960999999)
var pa4_cgos__erff float32 = float32(-0.11089469491999999)
var pa5_cgos__erff float32 = float32(0.035478305071999998)
var pa6_cgos__erff float32 = float32(-0.0021663755177999998)
var qa1_cgos__erff float32 = float32(0.10642088205)
var qa2_cgos__erff float32 = float32(0.54039794207000003)
var qa3_cgos__erff float32 = float32(0.071828655899000005)
var qa4_cgos__erff float32 = float32(0.12617121637000001)
var qa5_cgos__erff float32 = float32(0.013637083583)
var qa6_cgos__erff float32 = float32(0.011984500103)
var ra0_cgos__erff float32 = float32(-0.0098649440333000004)
var ra1_cgos__erff float32 = float32(-0.6938585639)
var ra2_cgos__erff float32 = float32(-10.558626175000001)
var ra3_cgos__erff float32 = float32(-62.375331879000001)
var ra4_cgos__erff float32 = float32(-162.39666747999999)
var ra5_cgos__erff float32 = float32(-184.60508727999999)
var ra6_cgos__erff float32 = float32(-81.287437439000001)
var ra7_cgos__erff float32 = float32(-9.8143291473000004)
var sa1_cgos__erff float32 = float32(19.651271820000002)
var sa2_cgos__erff float32 = float32(137.65776062)
var sa3_cgos__erff float32 = float32(434.56588744999999)
var sa4_cgos__erff float32 = float32(645.38726807)
var sa5_cgos__erff float32 = float32(429.00814818999999)
var sa6_cgos__erff float32 = float32(108.63500214)
var sa7_cgos__erff float32 = float32(6.5702495575000004)
var sa8_cgos__erff float32 = float32(-0.060424413532000003)
var rb0_cgos__erff float32 = float32(-0.0098649431019999997)
var rb1_cgos__erff float32 = float32(-0.79928326607)
var rb2_cgos__erff float32 = float32(-17.757955550999998)
var rb3_cgos__erff float32 = float32(-160.63638305999999)
var rb4_cgos__erff float32 = float32(-637.56646728999999)
var rb5_cgos__erff float32 = float32(-1025.0950928)
var rb6_cgos__erff float32 = float32(-483.51919556000001)
var sb1_cgos__erff float32 = float32(30.338060379000002)
var sb2_cgos__erff float32 = float32(325.79251098999998)
var sb3_cgos__erff float32 = float32(1536.7296143000001)
var sb4_cgos__erff float32 = float32(3199.8581543)
var sb5_cgos__erff float32 = float32(2553.0502929999998)
var sb6_cgos__erff float32 = float32(474.52853393999999)
var sb7_cgos__erff float32 = float32(-22.440952300999999)

func erfc1_cgos__erff(x float32) float32 {
	var s float32
	var P float32
	var Q float32
	s = Fabsf(x) - float32(int32(1))
	P = pa0_cgos__erff + s*(pa1_cgos__erff+s*(pa2_cgos__erff+s*(pa3_cgos__erff+s*(pa4_cgos__erff+s*(pa5_cgos__erff+s*pa6_cgos__erff)))))
	Q = float32(int32(1)) + s*(qa1_cgos__erff+s*(qa2_cgos__erff+s*(qa3_cgos__erff+s*(qa4_cgos__erff+s*(qa5_cgos__erff+s*qa6_cgos__erff)))))
	return float32(int32(1)) - erx_cgos__erff - P/Q
}
func erfc2_cgos__erff(ix uint32, x float32) float32 {
	var s float32
	var R float32
	var S float32
	var z float32
	if ix < uint32(1067450368) {
		return erfc1_cgos__erff(x)
	}
	x = Fabsf(x)
	s = float32(int32(1)) / (x * x)
	if ix < uint32(1077336941) {
		R = ra0_cgos__erff + s*(ra1_cgos__erff+s*(ra2_cgos__erff+s*(ra3_cgos__erff+s*(ra4_cgos__erff+s*(ra5_cgos__erff+s*(ra6_cgos__erff+s*ra7_cgos__erff))))))
		S = 1 + s*(sa1_cgos__erff+s*(sa2_cgos__erff+s*(sa3_cgos__erff+s*(sa4_cgos__erff+s*(sa5_cgos__erff+s*(sa6_cgos__erff+s*(sa7_cgos__erff+s*sa8_cgos__erff)))))))
	} else {
		R = rb0_cgos__erff + s*(rb1_cgos__erff+s*(rb2_cgos__erff+s*(rb3_cgos__erff+s*(rb4_cgos__erff+s*(rb5_cgos__erff+s*rb6_cgos__erff)))))
		S = 1 + s*(sb1_cgos__erff+s*(sb2_cgos__erff+s*(sb3_cgos__erff+s*(sb4_cgos__erff+s*(sb5_cgos__erff+s*(sb6_cgos__erff+s*sb7_cgos__erff))))))
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
			return 0.125 * (float32(int32(8))*x + efx8_cgos__erff*x)
		}
		z = x * x
		r = pp0_cgos__erff + z*(pp1_cgos__erff+z*(pp2_cgos__erff+z*(pp3_cgos__erff+z*pp4_cgos__erff)))
		s = float32(int32(1)) + z*(qq1_cgos__erff+z*(qq2_cgos__erff+z*(qq3_cgos__erff+z*(qq4_cgos__erff+z*qq5_cgos__erff))))
		y = r / s
		return x + x*y
	}
	if ix < uint32(1086324736) {
		y = float32(int32(1)) - erfc2_cgos__erff(ix, x)
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
		r = pp0_cgos__erff + z*(pp1_cgos__erff+z*(pp2_cgos__erff+z*(pp3_cgos__erff+z*pp4_cgos__erff)))
		s = 1 + z*(qq1_cgos__erff+z*(qq2_cgos__erff+z*(qq3_cgos__erff+z*(qq4_cgos__erff+z*qq5_cgos__erff))))
		y = r / s
		if sign != 0 || ix < uint32(1048576000) {
			return 1 - (x + x*y)
		}
		return 0.5 - (x - 0.5 + x*y)
	}
	if ix < uint32(1105199104) {
		return func() float32 {
			if sign != 0 {
				return float32(int32(2)) - erfc2_cgos__erff(ix, x)
			} else {
				return erfc2_cgos__erff(ix, x)
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
