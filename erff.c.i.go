package libc

import unsafe "unsafe"

var erx_cgo18_erff float32 = float32(0.84506291151000001)
var efx8_cgo19_erff float32 = float32(1.027033329)
var pp0_cgo20_erff float32 = float32(0.12837916613)
var pp1_cgo21_erff float32 = float32(-0.32504209876000001)
var pp2_cgo22_erff float32 = float32(-0.028481749818000002)
var pp3_cgo23_erff float32 = float32(-0.0057702702470000004)
var pp4_cgo24_erff float32 = float32(-2.3763017451999999e-5)
var qq1_cgo25_erff float32 = float32(0.39791721106)
var qq2_cgo26_erff float32 = float32(0.065022252500000002)
var qq3_cgo27_erff float32 = float32(0.0050813062117000003)
var qq4_cgo28_erff float32 = float32(1.3249473704e-4)
var qq5_cgo29_erff float32 = float32(-3.9602282412999997e-6)
var pa0_cgo30_erff float32 = float32(-0.0023621185682999998)
var pa1_cgo31_erff float32 = float32(0.41485610604000001)
var pa2_cgo32_erff float32 = float32(-0.37220788001999999)
var pa3_cgo33_erff float32 = float32(0.31834661960999999)
var pa4_cgo34_erff float32 = float32(-0.11089469491999999)
var pa5_cgo35_erff float32 = float32(0.035478305071999998)
var pa6_cgo36_erff float32 = float32(-0.0021663755177999998)
var qa1_cgo37_erff float32 = float32(0.10642088205)
var qa2_cgo38_erff float32 = float32(0.54039794207000003)
var qa3_cgo39_erff float32 = float32(0.071828655899000005)
var qa4_cgo40_erff float32 = float32(0.12617121637000001)
var qa5_cgo41_erff float32 = float32(0.013637083583)
var qa6_cgo42_erff float32 = float32(0.011984500103)
var ra0_cgo43_erff float32 = float32(-0.0098649440333000004)
var ra1_cgo44_erff float32 = float32(-0.6938585639)
var ra2_cgo45_erff float32 = float32(-10.558626175000001)
var ra3_cgo46_erff float32 = float32(-62.375331879000001)
var ra4_cgo47_erff float32 = float32(-162.39666747999999)
var ra5_cgo48_erff float32 = float32(-184.60508727999999)
var ra6_cgo49_erff float32 = float32(-81.287437439000001)
var ra7_cgo50_erff float32 = float32(-9.8143291473000004)
var sa1_cgo51_erff float32 = float32(19.651271820000002)
var sa2_cgo52_erff float32 = float32(137.65776062)
var sa3_cgo53_erff float32 = float32(434.56588744999999)
var sa4_cgo54_erff float32 = float32(645.38726807)
var sa5_cgo55_erff float32 = float32(429.00814818999999)
var sa6_cgo56_erff float32 = float32(108.63500214)
var sa7_cgo57_erff float32 = float32(6.5702495575000004)
var sa8_cgo58_erff float32 = float32(-0.060424413532000003)
var rb0_cgo59_erff float32 = float32(-0.0098649431019999997)
var rb1_cgo60_erff float32 = float32(-0.79928326607)
var rb2_cgo61_erff float32 = float32(-17.757955550999998)
var rb3_cgo62_erff float32 = float32(-160.63638305999999)
var rb4_cgo63_erff float32 = float32(-637.56646728999999)
var rb5_cgo64_erff float32 = float32(-1025.0950928)
var rb6_cgo65_erff float32 = float32(-483.51919556000001)
var sb1_cgo66_erff float32 = float32(30.338060379000002)
var sb2_cgo67_erff float32 = float32(325.79251098999998)
var sb3_cgo68_erff float32 = float32(1536.7296143000001)
var sb4_cgo69_erff float32 = float32(3199.8581543)
var sb5_cgo70_erff float32 = float32(2553.0502929999998)
var sb6_cgo71_erff float32 = float32(474.52853393999999)
var sb7_cgo72_erff float32 = float32(-22.440952300999999)

func erfc1_cgo73_erff(x float32) float32 {
	var s float32
	var P float32
	var Q float32
	s = Fabsf(x) - float32(int32(1))
	P = pa0_cgo30_erff + s*(pa1_cgo31_erff+s*(pa2_cgo32_erff+s*(pa3_cgo33_erff+s*(pa4_cgo34_erff+s*(pa5_cgo35_erff+s*pa6_cgo36_erff)))))
	Q = float32(int32(1)) + s*(qa1_cgo37_erff+s*(qa2_cgo38_erff+s*(qa3_cgo39_erff+s*(qa4_cgo40_erff+s*(qa5_cgo41_erff+s*qa6_cgo42_erff)))))
	return float32(int32(1)) - erx_cgo18_erff - P/Q
}
func erfc2_cgo74_erff(ix uint32, x float32) float32 {
	var s float32
	var R float32
	var S float32
	var z float32
	if ix < uint32(1067450368) {
		return erfc1_cgo73_erff(x)
	}
	x = Fabsf(x)
	s = float32(int32(1)) / (x * x)
	if ix < uint32(1077336941) {
		R = ra0_cgo43_erff + s*(ra1_cgo44_erff+s*(ra2_cgo45_erff+s*(ra3_cgo46_erff+s*(ra4_cgo47_erff+s*(ra5_cgo48_erff+s*(ra6_cgo49_erff+s*ra7_cgo50_erff))))))
		S = 1 + s*(sa1_cgo51_erff+s*(sa2_cgo52_erff+s*(sa3_cgo53_erff+s*(sa4_cgo54_erff+s*(sa5_cgo55_erff+s*(sa6_cgo56_erff+s*(sa7_cgo57_erff+s*sa8_cgo58_erff)))))))
	} else {
		R = rb0_cgo59_erff + s*(rb1_cgo60_erff+s*(rb2_cgo61_erff+s*(rb3_cgo62_erff+s*(rb4_cgo63_erff+s*(rb5_cgo64_erff+s*rb6_cgo65_erff)))))
		S = 1 + s*(sb1_cgo66_erff+s*(sb2_cgo67_erff+s*(sb3_cgo68_erff+s*(sb4_cgo69_erff+s*(sb5_cgo70_erff+s*(sb6_cgo71_erff+s*sb7_cgo72_erff))))))
	}
	for {
		ix = *(*uint32)(unsafe.Pointer(&_cgoz_75_erff{x}))
		if true {
			break
		}
	}
	for {
		z = *(*float32)(unsafe.Pointer(&_cgoz_76_erff{ix & uint32(4294959104)}))
		if true {
			break
		}
	}
	return Expf(-z*z-0.5625) * Expf((z-x)*(z+x)+R/S) / x
}

type _cgoz_75_erff struct {
	_f float32
}
type _cgoz_76_erff struct {
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
		ix = *(*uint32)(unsafe.Pointer(&_cgoz_77_erff{x}))
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
			return 0.125 * (float32(int32(8))*x + efx8_cgo19_erff*x)
		}
		z = x * x
		r = pp0_cgo20_erff + z*(pp1_cgo21_erff+z*(pp2_cgo22_erff+z*(pp3_cgo23_erff+z*pp4_cgo24_erff)))
		s = float32(int32(1)) + z*(qq1_cgo25_erff+z*(qq2_cgo26_erff+z*(qq3_cgo27_erff+z*(qq4_cgo28_erff+z*qq5_cgo29_erff))))
		y = r / s
		return x + x*y
	}
	if ix < uint32(1086324736) {
		y = float32(int32(1)) - erfc2_cgo74_erff(ix, x)
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

type _cgoz_77_erff struct {
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
		ix = *(*uint32)(unsafe.Pointer(&_cgoz_78_erff{x}))
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
		r = pp0_cgo20_erff + z*(pp1_cgo21_erff+z*(pp2_cgo22_erff+z*(pp3_cgo23_erff+z*pp4_cgo24_erff)))
		s = 1 + z*(qq1_cgo25_erff+z*(qq2_cgo26_erff+z*(qq3_cgo27_erff+z*(qq4_cgo28_erff+z*qq5_cgo29_erff))))
		y = r / s
		if sign != 0 || ix < uint32(1048576000) {
			return 1 - (x + x*y)
		}
		return 0.5 - (x - 0.5 + x*y)
	}
	if ix < uint32(1105199104) {
		return func() float32 {
			if sign != 0 {
				return float32(int32(2)) - erfc2_cgo74_erff(ix, x)
			} else {
				return erfc2_cgo74_erff(ix, x)
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

type _cgoz_78_erff struct {
	_f float32
}
