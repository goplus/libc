package libc

import unsafe "unsafe"

var _cgos_pi_lgammaf_r float32 = float32(3.1415927410000002)
var _cgos_a0_lgammaf_r float32 = float32(0.077215664089000005)
var _cgos_a1_lgammaf_r float32 = float32(0.32246702909000002)
var _cgos_a2_lgammaf_r float32 = float32(0.067352302372000003)
var _cgos_a3_lgammaf_r float32 = float32(0.020580807701000001)
var _cgos_a4_lgammaf_r float32 = float32(0.0073855509982000004)
var _cgos_a5_lgammaf_r float32 = float32(0.0028905137442000002)
var _cgos_a6_lgammaf_r float32 = float32(0.0011927076848)
var _cgos_a7_lgammaf_r float32 = float32(5.1006977445999997e-4)
var _cgos_a8_lgammaf_r float32 = float32(2.2086278476999999e-4)
var _cgos_a9_lgammaf_r float32 = float32(1.0801156894999999e-4)
var _cgos_a10_lgammaf_r float32 = float32(2.52144564e-5)
var _cgos_a11_lgammaf_r float32 = float32(4.4864096708000002e-5)
var _cgos_tc_lgammaf_r float32 = float32(1.4616321324999999)
var _cgos_tf_lgammaf_r float32 = float32(-0.12148628384)
var _cgos_tt_lgammaf_r float32 = float32(6.6971006517999998e-9)
var _cgos_t0_lgammaf_r float32 = float32(0.48383611441000002)
var _cgos_t1_lgammaf_r float32 = float32(-0.14758771658)
var _cgos_t2_lgammaf_r float32 = float32(0.064624942838999999)
var _cgos_t3_lgammaf_r float32 = float32(-0.032788541168)
var _cgos_t4_lgammaf_r float32 = float32(0.017970675602999999)
var _cgos_t5_lgammaf_r float32 = float32(-0.010314224287999999)
var _cgos_t6_lgammaf_r float32 = float32(0.0061005386524000003)
var _cgos_t7_lgammaf_r float32 = float32(-0.0036845202557999998)
var _cgos_t8_lgammaf_r float32 = float32(0.0022596477065)
var _cgos_t9_lgammaf_r float32 = float32(-0.0014034647029)
var _cgos_t10_lgammaf_r float32 = float32(8.8108185445999996e-4)
var _cgos_t11_lgammaf_r float32 = float32(-5.3859531181000002e-4)
var _cgos_t12_lgammaf_r float32 = float32(3.1563205994e-4)
var _cgos_t13_lgammaf_r float32 = float32(-3.1275415677000002e-4)
var _cgos_t14_lgammaf_r float32 = float32(3.3552918466999999e-4)
var _cgos_u0_lgammaf_r float32 = float32(-0.077215664089000005)
var _cgos_u1_lgammaf_r float32 = float32(0.63282704352999997)
var _cgos_u2_lgammaf_r float32 = float32(1.4549225568999999)
var _cgos_u3_lgammaf_r float32 = float32(0.97771751881000001)
var _cgos_u4_lgammaf_r float32 = float32(0.22896373272000001)
var _cgos_u5_lgammaf_r float32 = float32(0.013381091878)
var _cgos_v1_lgammaf_r float32 = float32(2.4559779167000002)
var _cgos_v2_lgammaf_r float32 = float32(2.1284897326999999)
var _cgos_v3_lgammaf_r float32 = float32(0.76928514242000001)
var _cgos_v4_lgammaf_r float32 = float32(0.10422264785)
var _cgos_v5_lgammaf_r float32 = float32(0.0032170924824000001)
var _cgos_s0_lgammaf_r float32 = float32(-0.077215664089000005)
var _cgos_s1_lgammaf_r float32 = float32(0.21498242021)
var _cgos_s2_lgammaf_r float32 = float32(0.32577878237000002)
var _cgos_s3_lgammaf_r float32 = float32(0.14635047317)
var _cgos_s4_lgammaf_r float32 = float32(0.026642270385999998)
var _cgos_s5_lgammaf_r float32 = float32(0.0018402845599)
var _cgos_s6_lgammaf_r float32 = float32(3.1947532988999999e-5)
var _cgos_r1_lgammaf_r float32 = float32(1.3920053243999999)
var _cgos_r2_lgammaf_r float32 = float32(0.72193557023999999)
var _cgos_r3_lgammaf_r float32 = float32(0.17193385958999999)
var _cgos_r4_lgammaf_r float32 = float32(0.018645919859)
var _cgos_r5_lgammaf_r float32 = float32(7.7794247772999998e-4)
var _cgos_r6_lgammaf_r float32 = float32(7.3266842264000001e-6)
var _cgos_w0_lgammaf_r float32 = float32(0.41893854737000003)
var _cgos_w1_lgammaf_r float32 = float32(0.083333335817000003)
var _cgos_w2_lgammaf_r float32 = float32(-0.002777777845)
var _cgos_w3_lgammaf_r float32 = float32(7.9365057172000003e-4)
var _cgos_w4_lgammaf_r float32 = float32(-5.9518753551000001e-4)
var _cgos_w5_lgammaf_r float32 = float32(8.3633989560999997e-4)
var _cgos_w6_lgammaf_r float32 = float32(-0.0016309292987)

func _cgos_sin_pi_lgammaf_r(x float32) float32 {
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
		t = _cgos_sin_pi_lgammaf_r(x)
		if t == 0.0 {
			return 1.0 / (x - x)
		}
		if t > 0.0 {
			*signgamp = -1
		} else {
			t = -t
		}
		nadj = Logf(_cgos_pi_lgammaf_r / (t * x))
	}
	if ix == uint32(1065353216) || ix == uint32(1073741824) {
		r = float32(int32(0))
	} else if ix < uint32(1073741824) {
		if ix <= uint32(1063675494) {
			r = -Logf(x)
			if ix >= uint32(1060850208) {
				y = 1.0 - x
				i = int32(0)
			} else if ix >= uint32(1047343880) {
				y = x - (_cgos_tc_lgammaf_r - 1.0)
				i = int32(1)
			} else {
				y = x
				i = int32(2)
			}
		} else {
			r = float32(0.0)
			if ix >= uint32(1071490584) {
				y = 2.0 - x
				i = int32(0)
			} else if ix >= uint32(1067296288) {
				y = x - _cgos_tc_lgammaf_r
				i = int32(1)
			} else {
				y = x - 1.0
				i = int32(2)
			}
		}
		switch i {
		case int32(0):
			z = y * y
			p1 = _cgos_a0_lgammaf_r + z*(_cgos_a2_lgammaf_r+z*(_cgos_a4_lgammaf_r+z*(_cgos_a6_lgammaf_r+z*(_cgos_a8_lgammaf_r+z*_cgos_a10_lgammaf_r))))
			p2 = z * (_cgos_a1_lgammaf_r + z*(_cgos_a3_lgammaf_r+z*(_cgos_a5_lgammaf_r+z*(_cgos_a7_lgammaf_r+z*(_cgos_a9_lgammaf_r+z*_cgos_a11_lgammaf_r)))))
			p = y*p1 + p2
			r += p - 0.5*y
			break
		case int32(1):
			z = y * y
			w = z * y
			p1 = _cgos_t0_lgammaf_r + w*(_cgos_t3_lgammaf_r+w*(_cgos_t6_lgammaf_r+w*(_cgos_t9_lgammaf_r+w*_cgos_t12_lgammaf_r)))
			p2 = _cgos_t1_lgammaf_r + w*(_cgos_t4_lgammaf_r+w*(_cgos_t7_lgammaf_r+w*(_cgos_t10_lgammaf_r+w*_cgos_t13_lgammaf_r)))
			p3 = _cgos_t2_lgammaf_r + w*(_cgos_t5_lgammaf_r+w*(_cgos_t8_lgammaf_r+w*(_cgos_t11_lgammaf_r+w*_cgos_t14_lgammaf_r)))
			p = z*p1 - (_cgos_tt_lgammaf_r - w*(p2+y*p3))
			r += _cgos_tf_lgammaf_r + p
			break
		case int32(2):
			p1 = y * (_cgos_u0_lgammaf_r + y*(_cgos_u1_lgammaf_r+y*(_cgos_u2_lgammaf_r+y*(_cgos_u3_lgammaf_r+y*(_cgos_u4_lgammaf_r+y*_cgos_u5_lgammaf_r)))))
			p2 = 1.0 + y*(_cgos_v1_lgammaf_r+y*(_cgos_v2_lgammaf_r+y*(_cgos_v3_lgammaf_r+y*(_cgos_v4_lgammaf_r+y*_cgos_v5_lgammaf_r))))
			r += -0.5*y + p1/p2
		}
	} else if ix < uint32(1090519040) {
		i = int32(x)
		y = x - float32(i)
		p = y * (_cgos_s0_lgammaf_r + y*(_cgos_s1_lgammaf_r+y*(_cgos_s2_lgammaf_r+y*(_cgos_s3_lgammaf_r+y*(_cgos_s4_lgammaf_r+y*(_cgos_s5_lgammaf_r+y*_cgos_s6_lgammaf_r))))))
		q = 1.0 + y*(_cgos_r1_lgammaf_r+y*(_cgos_r2_lgammaf_r+y*(_cgos_r3_lgammaf_r+y*(_cgos_r4_lgammaf_r+y*(_cgos_r5_lgammaf_r+y*_cgos_r6_lgammaf_r)))))
		r = 0.5*y + p/q
		z = float32(1.0)
		switch i {
		case int32(7):
			z *= y + 6.0
		case int32(6):
			z *= y + 5.0
		case int32(5):
			z *= y + 4.0
		case int32(4):
			z *= y + 3.0
		case int32(3):
			z *= y + 2.0
			r += Logf(z)
			break
		}
	} else if ix < uint32(1551892480) {
		t = Logf(x)
		z = 1.0 / x
		y = z * z
		w = _cgos_w0_lgammaf_r + z*(_cgos_w1_lgammaf_r+y*(_cgos_w2_lgammaf_r+y*(_cgos_w3_lgammaf_r+y*(_cgos_w4_lgammaf_r+y*(_cgos_w5_lgammaf_r+y*_cgos_w6_lgammaf_r)))))
		r = (x-0.5)*(t-1.0) + w
	} else {
		r = x * (Logf(x) - 1.0)
	}
	if sign != 0 {
		r = nadj - r
	}
	return r
}
