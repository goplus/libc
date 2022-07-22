package libc

import unsafe "unsafe"

var pi_cgo18_lgamma_r float64 = 3.1415926535897931
var a0_cgo19_lgamma_r float64 = 0.077215664901532866
var a1_cgo20_lgamma_r float64 = 0.32246703342411359
var a2_cgo21_lgamma_r float64 = 0.067352301053129268
var a3_cgo22_lgamma_r float64 = 0.020580808432516733
var a4_cgo23_lgamma_r float64 = 0.0073855508608140288
var a5_cgo24_lgamma_r float64 = 0.0028905138367341563
var a6_cgo25_lgamma_r float64 = 0.0011927076318336207
var a7_cgo26_lgamma_r float64 = 5.1006979215351134e-4
var a8_cgo27_lgamma_r float64 = 2.2086279071390839e-4
var a9_cgo28_lgamma_r float64 = 1.0801156724758394e-4
var a10_cgo29_lgamma_r float64 = 2.5214456545125733e-5
var a11_cgo30_lgamma_r float64 = 4.4864094961891516e-5
var tc_cgo31_lgamma_r float64 = 1.4616321449683622
var tf_cgo32_lgamma_r float64 = -0.12148629053584961
var tt_cgo33_lgamma_r float64 = -3.6386769970395054e-18
var t0_cgo34_lgamma_r float64 = 0.48383612272381005
var t1_cgo35_lgamma_r float64 = -0.14758772299459391
var t2_cgo36_lgamma_r float64 = 0.064624940239133385
var t3_cgo37_lgamma_r float64 = -0.032788541075985965
var t4_cgo38_lgamma_r float64 = 0.017970675081182039
var t5_cgo39_lgamma_r float64 = -0.010314224129834144
var t6_cgo40_lgamma_r float64 = 0.0061005387024629133
var t7_cgo41_lgamma_r float64 = -0.0036845201678113826
var t8_cgo42_lgamma_r float64 = 0.0022596478090061247
var t9_cgo43_lgamma_r float64 = -0.0014034646998923284
var t10_cgo44_lgamma_r float64 = 8.8108188243765401e-4
var t11_cgo45_lgamma_r float64 = -5.3859530535674055e-4
var t12_cgo46_lgamma_r float64 = 3.1563207090362595e-4
var t13_cgo47_lgamma_r float64 = -3.1275416837512086e-4
var t14_cgo48_lgamma_r float64 = 3.3552919263551907e-4
var u0_cgo49_lgamma_r float64 = -0.077215664901532866
var u1_cgo50_lgamma_r float64 = 0.63282706402509337
var u2_cgo51_lgamma_r float64 = 1.4549225013723477
var u3_cgo52_lgamma_r float64 = 0.97771752796337275
var u4_cgo53_lgamma_r float64 = 0.22896372806469245
var u5_cgo54_lgamma_r float64 = 0.013381091853678766
var v1_cgo55_lgamma_r float64 = 2.4559779371304113
var v2_cgo56_lgamma_r float64 = 2.128489763798934
var v3_cgo57_lgamma_r float64 = 0.76928515045667278
var v4_cgo58_lgamma_r float64 = 0.10422264559336913
var v5_cgo59_lgamma_r float64 = 0.0032170924228242391
var s0_cgo60_lgamma_r float64 = -0.077215664901532866
var s1_cgo61_lgamma_r float64 = 0.21498241596060885
var s2_cgo62_lgamma_r float64 = 0.32577879640893098
var s3_cgo63_lgamma_r float64 = 0.14635047265246445
var s4_cgo64_lgamma_r float64 = 0.026642270303363861
var s5_cgo65_lgamma_r float64 = 0.0018402845140733772
var s6_cgo66_lgamma_r float64 = 3.1947532658410087e-5
var r1_cgo67_lgamma_r float64 = 1.3920053346762105
var r2_cgo68_lgamma_r float64 = 0.72193554756713807
var r3_cgo69_lgamma_r float64 = 0.17193386563280308
var r4_cgo70_lgamma_r float64 = 0.01864591917156529
var r5_cgo71_lgamma_r float64 = 7.779424963818936e-4
var r6_cgo72_lgamma_r float64 = 7.3266843074462564e-6
var w0_cgo73_lgamma_r float64 = 0.41893853320467273
var w1_cgo74_lgamma_r float64 = 0.083333333333332968
var w2_cgo75_lgamma_r float64 = -0.0027777777772877554
var w3_cgo76_lgamma_r float64 = 7.9365055864301956e-4
var w4_cgo77_lgamma_r float64 = -5.9518755745033996e-4
var w5_cgo78_lgamma_r float64 = 8.3633991899628213e-4
var w6_cgo79_lgamma_r float64 = -0.0016309293409657527

func sin_pi_cgo80_lgamma_r(x float64) float64 {
	var n int32
	x = 2 * (x*0.5 - Floor(x*0.5))
	n = int32(x * 4)
	n = (n + int32(1)) / int32(2)
	x -= float64(float32(n) * 0.5)
	x *= pi_cgo18_lgamma_r
	switch n {
	default:
		fallthrough
	case int32(0):
		return __sin(x, 0, int32(0))
	case int32(1):
		return __cos(x, 0)
	case int32(2):
		return __sin(-x, 0, int32(0))
	case int32(3):
		return -__cos(x, 0)
	}
	return 0
}
func __lgamma_r(x float64, signgamp *int32) float64 {
	type _cgoa_81_lgamma_r struct {
		f float64
	}
	var u _cgoa_81_lgamma_r
	u.f = x
	var t float64
	var y float64
	var z float64
	var nadj float64
	var p float64
	var p1 float64
	var p2 float64
	var p3 float64
	var q float64
	var r float64
	var w float64
	var ix uint32
	var sign int32
	var i int32
	*signgamp = int32(1)
	sign = int32(*(*uint64)(unsafe.Pointer(&u)) >> int32(63))
	ix = uint32(*(*uint64)(unsafe.Pointer(&u)) >> int32(32) & uint64(2147483647))
	if ix >= uint32(2146435072) {
		return x * x
	}
	if ix < uint32(999292928) {
		if sign != 0 {
			x = -x
			*signgamp = -1
		}
		return -Log(x)
	}
	if sign != 0 {
		x = -x
		t = sin_pi_cgo80_lgamma_r(x)
		if t == 0 {
			return 1 / (x - x)
		}
		if t > 0 {
			*signgamp = -1
		} else {
			t = -t
		}
		nadj = Log(pi_cgo18_lgamma_r / (t * x))
	}
	if (ix == uint32(1072693248) || ix == uint32(1073741824)) && uint32(*(*uint64)(unsafe.Pointer(&u))) == uint32(0) {
		r = float64(int32(0))
	} else if ix < uint32(1073741824) {
		if ix <= uint32(1072483532) {
			r = -Log(x)
			if ix >= uint32(1072130372) {
				y = 1 - x
				i = int32(0)
			} else if ix >= uint32(1070442081) {
				y = x - (tc_cgo31_lgamma_r - 1)
				i = int32(1)
			} else {
				y = x
				i = int32(2)
			}
		} else {
			r = float64(0)
			if ix >= uint32(1073460419) {
				y = 2 - x
				i = int32(0)
			} else if ix >= uint32(1072936132) {
				y = x - tc_cgo31_lgamma_r
				i = int32(1)
			} else {
				y = x - 1
				i = int32(2)
			}
		}
		switch i {
		case int32(0):
			z = y * y
			p1 = a0_cgo19_lgamma_r + z*(a2_cgo21_lgamma_r+z*(a4_cgo23_lgamma_r+z*(a6_cgo25_lgamma_r+z*(a8_cgo27_lgamma_r+z*a10_cgo29_lgamma_r))))
			p2 = z * (a1_cgo20_lgamma_r + z*(a3_cgo22_lgamma_r+z*(a5_cgo24_lgamma_r+z*(a7_cgo26_lgamma_r+z*(a9_cgo28_lgamma_r+z*a11_cgo30_lgamma_r)))))
			p = y*p1 + p2
			r += p - 0.5*y
			break
		case int32(1):
			z = y * y
			w = z * y
			p1 = t0_cgo34_lgamma_r + w*(t3_cgo37_lgamma_r+w*(t6_cgo40_lgamma_r+w*(t9_cgo43_lgamma_r+w*t12_cgo46_lgamma_r)))
			p2 = t1_cgo35_lgamma_r + w*(t4_cgo38_lgamma_r+w*(t7_cgo41_lgamma_r+w*(t10_cgo44_lgamma_r+w*t13_cgo47_lgamma_r)))
			p3 = t2_cgo36_lgamma_r + w*(t5_cgo39_lgamma_r+w*(t8_cgo42_lgamma_r+w*(t11_cgo45_lgamma_r+w*t14_cgo48_lgamma_r)))
			p = z*p1 - (tt_cgo33_lgamma_r - w*(p2+y*p3))
			r += tf_cgo32_lgamma_r + p
			break
		case int32(2):
			p1 = y * (u0_cgo49_lgamma_r + y*(u1_cgo50_lgamma_r+y*(u2_cgo51_lgamma_r+y*(u3_cgo52_lgamma_r+y*(u4_cgo53_lgamma_r+y*u5_cgo54_lgamma_r)))))
			p2 = 1 + y*(v1_cgo55_lgamma_r+y*(v2_cgo56_lgamma_r+y*(v3_cgo57_lgamma_r+y*(v4_cgo58_lgamma_r+y*v5_cgo59_lgamma_r))))
			r += -0.5*y + p1/p2
		}
	} else if ix < uint32(1075838976) {
		i = int32(x)
		y = x - float64(i)
		p = y * (s0_cgo60_lgamma_r + y*(s1_cgo61_lgamma_r+y*(s2_cgo62_lgamma_r+y*(s3_cgo63_lgamma_r+y*(s4_cgo64_lgamma_r+y*(s5_cgo65_lgamma_r+y*s6_cgo66_lgamma_r))))))
		q = 1 + y*(r1_cgo67_lgamma_r+y*(r2_cgo68_lgamma_r+y*(r3_cgo69_lgamma_r+y*(r4_cgo70_lgamma_r+y*(r5_cgo71_lgamma_r+y*r6_cgo72_lgamma_r)))))
		r = 0.5*y + p/q
		z = float64(1)
		switch i {
		case int32(7):
			z *= y + 6
		case int32(6):
			z *= y + 5
		case int32(5):
			z *= y + 4
		case int32(4):
			z *= y + 3
		case int32(3):
			z *= y + 2
			r += Log(z)
			break
		}
	} else if ix < uint32(1133510656) {
		t = Log(x)
		z = 1 / x
		y = z * z
		w = w0_cgo73_lgamma_r + z*(w1_cgo74_lgamma_r+y*(w2_cgo75_lgamma_r+y*(w3_cgo76_lgamma_r+y*(w4_cgo77_lgamma_r+y*(w5_cgo78_lgamma_r+y*w6_cgo79_lgamma_r)))))
		r = (x-0.5)*(t-1) + w
	} else {
		r = x * (Log(x) - 1)
	}
	if sign != 0 {
		r = nadj - r
	}
	return r
}
