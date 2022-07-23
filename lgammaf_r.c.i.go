package libc

import unsafe "unsafe"

var pi_cgos__lgammaf_r float32 = float32(3.1415927410000002)
var a0_cgos__lgammaf_r float32 = float32(0.077215664089000005)
var a1_cgos__lgammaf_r float32 = float32(0.32246702909000002)
var a2_cgos__lgammaf_r float32 = float32(0.067352302372000003)
var a3_cgos__lgammaf_r float32 = float32(0.020580807701000001)
var a4_cgos__lgammaf_r float32 = float32(0.0073855509982000004)
var a5_cgos__lgammaf_r float32 = float32(0.0028905137442000002)
var a6_cgos__lgammaf_r float32 = float32(0.0011927076848)
var a7_cgos__lgammaf_r float32 = float32(5.1006977445999997e-4)
var a8_cgos__lgammaf_r float32 = float32(2.2086278476999999e-4)
var a9_cgos__lgammaf_r float32 = float32(1.0801156894999999e-4)
var a10_cgos__lgammaf_r float32 = float32(2.52144564e-5)
var a11_cgos__lgammaf_r float32 = float32(4.4864096708000002e-5)
var tc_cgos__lgammaf_r float32 = float32(1.4616321324999999)
var tf_cgos__lgammaf_r float32 = float32(-0.12148628384)
var tt_cgos__lgammaf_r float32 = float32(6.6971006517999998e-9)
var t0_cgos__lgammaf_r float32 = float32(0.48383611441000002)
var t1_cgos__lgammaf_r float32 = float32(-0.14758771658)
var t2_cgos__lgammaf_r float32 = float32(0.064624942838999999)
var t3_cgos__lgammaf_r float32 = float32(-0.032788541168)
var t4_cgos__lgammaf_r float32 = float32(0.017970675602999999)
var t5_cgos__lgammaf_r float32 = float32(-0.010314224287999999)
var t6_cgos__lgammaf_r float32 = float32(0.0061005386524000003)
var t7_cgos__lgammaf_r float32 = float32(-0.0036845202557999998)
var t8_cgos__lgammaf_r float32 = float32(0.0022596477065)
var t9_cgos__lgammaf_r float32 = float32(-0.0014034647029)
var t10_cgos__lgammaf_r float32 = float32(8.8108185445999996e-4)
var t11_cgos__lgammaf_r float32 = float32(-5.3859531181000002e-4)
var t12_cgos__lgammaf_r float32 = float32(3.1563205994e-4)
var t13_cgos__lgammaf_r float32 = float32(-3.1275415677000002e-4)
var t14_cgos__lgammaf_r float32 = float32(3.3552918466999999e-4)
var u0_cgos__lgammaf_r float32 = float32(-0.077215664089000005)
var u1_cgos__lgammaf_r float32 = float32(0.63282704352999997)
var u2_cgos__lgammaf_r float32 = float32(1.4549225568999999)
var u3_cgos__lgammaf_r float32 = float32(0.97771751881000001)
var u4_cgos__lgammaf_r float32 = float32(0.22896373272000001)
var u5_cgos__lgammaf_r float32 = float32(0.013381091878)
var v1_cgos__lgammaf_r float32 = float32(2.4559779167000002)
var v2_cgos__lgammaf_r float32 = float32(2.1284897326999999)
var v3_cgos__lgammaf_r float32 = float32(0.76928514242000001)
var v4_cgos__lgammaf_r float32 = float32(0.10422264785)
var v5_cgos__lgammaf_r float32 = float32(0.0032170924824000001)
var s0_cgos__lgammaf_r float32 = float32(-0.077215664089000005)
var s1_cgos__lgammaf_r float32 = float32(0.21498242021)
var s2_cgos__lgammaf_r float32 = float32(0.32577878237000002)
var s3_cgos__lgammaf_r float32 = float32(0.14635047317)
var s4_cgos__lgammaf_r float32 = float32(0.026642270385999998)
var s5_cgos__lgammaf_r float32 = float32(0.0018402845599)
var s6_cgos__lgammaf_r float32 = float32(3.1947532988999999e-5)
var r1_cgos__lgammaf_r float32 = float32(1.3920053243999999)
var r2_cgos__lgammaf_r float32 = float32(0.72193557023999999)
var r3_cgos__lgammaf_r float32 = float32(0.17193385958999999)
var r4_cgos__lgammaf_r float32 = float32(0.018645919859)
var r5_cgos__lgammaf_r float32 = float32(7.7794247772999998e-4)
var r6_cgos__lgammaf_r float32 = float32(7.3266842264000001e-6)
var w0_cgos__lgammaf_r float32 = float32(0.41893854737000003)
var w1_cgos__lgammaf_r float32 = float32(0.083333335817000003)
var w2_cgos__lgammaf_r float32 = float32(-0.002777777845)
var w3_cgos__lgammaf_r float32 = float32(7.9365057172000003e-4)
var w4_cgos__lgammaf_r float32 = float32(-5.9518753551000001e-4)
var w5_cgos__lgammaf_r float32 = float32(8.3633989560999997e-4)
var w6_cgos__lgammaf_r float32 = float32(-0.0016309292987)

func sin_pi_cgos__lgammaf_r(x float32) float32 {
	var y float64
	var n int32
	x = float32(int32(2)) * (x*0.5 - Floorf(x*0.5))
	n = int32(x * float32(int32(4)))
	n = (n + int32(1)) / int32(2)
	y = float64(x - float32(n)*0.5)
	y *= float64(3.1415926535897931)
	switch n {
	default:
		fallthrough
	case int32(0):
		return __sindf(y)
	case int32(1):
		return __cosdf(y)
	case int32(2):
		return __sindf(-y)
	case int32(3):
		return -__cosdf(y)
	}
	return 0
}
func __lgammaf_r(x float32, signgamp *int32) float32 {
	type _cgoa_18_lgammaf_r struct {
		f float32
	}
	var u _cgoa_18_lgammaf_r
	u.f = x
	var t float32
	var y float32
	var z float32
	var nadj float32
	var p float32
	var p1 float32
	var p2 float32
	var p3 float32
	var q float32
	var r float32
	var w float32
	var ix uint32
	var i int32
	var sign int32
	*signgamp = int32(1)
	sign = int32(*(*uint32)(unsafe.Pointer(&u)) >> int32(31))
	ix = *(*uint32)(unsafe.Pointer(&u)) & uint32(2147483647)
	if ix >= uint32(2139095040) {
		return x * x
	}
	if ix < uint32(889192448) {
		if sign != 0 {
			*signgamp = -1
			x = -x
		}
		return -Logf(x)
	}
	if sign != 0 {
		x = -x
		t = sin_pi_cgos__lgammaf_r(x)
		if t == 0 {
			return 1 / (x - x)
		}
		if t > 0 {
			*signgamp = -1
		} else {
			t = -t
		}
		nadj = Logf(pi_cgos__lgammaf_r / (t * x))
	}
	if ix == uint32(1065353216) || ix == uint32(1073741824) {
		r = float32(int32(0))
	} else if ix < uint32(1073741824) {
		if ix <= uint32(1063675494) {
			r = -Logf(x)
			if ix >= uint32(1060850208) {
				y = 1 - x
				i = int32(0)
			} else if ix >= uint32(1047343880) {
				y = x - (tc_cgos__lgammaf_r - 1)
				i = int32(1)
			} else {
				y = x
				i = int32(2)
			}
		} else {
			r = float32(0)
			if ix >= uint32(1071490584) {
				y = 2 - x
				i = int32(0)
			} else if ix >= uint32(1067296288) {
				y = x - tc_cgos__lgammaf_r
				i = int32(1)
			} else {
				y = x - 1
				i = int32(2)
			}
		}
		switch i {
		case int32(0):
			z = y * y
			p1 = a0_cgos__lgammaf_r + z*(a2_cgos__lgammaf_r+z*(a4_cgos__lgammaf_r+z*(a6_cgos__lgammaf_r+z*(a8_cgos__lgammaf_r+z*a10_cgos__lgammaf_r))))
			p2 = z * (a1_cgos__lgammaf_r + z*(a3_cgos__lgammaf_r+z*(a5_cgos__lgammaf_r+z*(a7_cgos__lgammaf_r+z*(a9_cgos__lgammaf_r+z*a11_cgos__lgammaf_r)))))
			p = y*p1 + p2
			r += p - 0.5*y
			break
		case int32(1):
			z = y * y
			w = z * y
			p1 = t0_cgos__lgammaf_r + w*(t3_cgos__lgammaf_r+w*(t6_cgos__lgammaf_r+w*(t9_cgos__lgammaf_r+w*t12_cgos__lgammaf_r)))
			p2 = t1_cgos__lgammaf_r + w*(t4_cgos__lgammaf_r+w*(t7_cgos__lgammaf_r+w*(t10_cgos__lgammaf_r+w*t13_cgos__lgammaf_r)))
			p3 = t2_cgos__lgammaf_r + w*(t5_cgos__lgammaf_r+w*(t8_cgos__lgammaf_r+w*(t11_cgos__lgammaf_r+w*t14_cgos__lgammaf_r)))
			p = z*p1 - (tt_cgos__lgammaf_r - w*(p2+y*p3))
			r += tf_cgos__lgammaf_r + p
			break
		case int32(2):
			p1 = y * (u0_cgos__lgammaf_r + y*(u1_cgos__lgammaf_r+y*(u2_cgos__lgammaf_r+y*(u3_cgos__lgammaf_r+y*(u4_cgos__lgammaf_r+y*u5_cgos__lgammaf_r)))))
			p2 = 1 + y*(v1_cgos__lgammaf_r+y*(v2_cgos__lgammaf_r+y*(v3_cgos__lgammaf_r+y*(v4_cgos__lgammaf_r+y*v5_cgos__lgammaf_r))))
			r += -0.5*y + p1/p2
		}
	} else if ix < uint32(1090519040) {
		i = int32(x)
		y = x - float32(i)
		p = y * (s0_cgos__lgammaf_r + y*(s1_cgos__lgammaf_r+y*(s2_cgos__lgammaf_r+y*(s3_cgos__lgammaf_r+y*(s4_cgos__lgammaf_r+y*(s5_cgos__lgammaf_r+y*s6_cgos__lgammaf_r))))))
		q = 1 + y*(r1_cgos__lgammaf_r+y*(r2_cgos__lgammaf_r+y*(r3_cgos__lgammaf_r+y*(r4_cgos__lgammaf_r+y*(r5_cgos__lgammaf_r+y*r6_cgos__lgammaf_r)))))
		r = 0.5*y + p/q
		z = float32(1)
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
			r += Logf(z)
			break
		}
	} else if ix < uint32(1551892480) {
		t = Logf(x)
		z = 1 / x
		y = z * z
		w = w0_cgos__lgammaf_r + z*(w1_cgos__lgammaf_r+y*(w2_cgos__lgammaf_r+y*(w3_cgos__lgammaf_r+y*(w4_cgos__lgammaf_r+y*(w5_cgos__lgammaf_r+y*w6_cgos__lgammaf_r)))))
		r = (x-0.5)*(t-1) + w
	} else {
		r = x * (Logf(x) - 1)
	}
	if sign != 0 {
		r = nadj - r
	}
	return r
}
