package libc

import unsafe "unsafe"

var _cgos_pi_lgamma_r float64 = 3.1415926535897931
var _cgos_a0_lgamma_r float64 = 0.077215664901532866
var _cgos_a1_lgamma_r float64 = 0.32246703342411359
var _cgos_a2_lgamma_r float64 = 0.067352301053129268
var _cgos_a3_lgamma_r float64 = 0.020580808432516733
var _cgos_a4_lgamma_r float64 = 0.0073855508608140288
var _cgos_a5_lgamma_r float64 = 0.0028905138367341563
var _cgos_a6_lgamma_r float64 = 0.0011927076318336207
var _cgos_a7_lgamma_r float64 = 5.1006979215351134e-4
var _cgos_a8_lgamma_r float64 = 2.2086279071390839e-4
var _cgos_a9_lgamma_r float64 = 1.0801156724758394e-4
var _cgos_a10_lgamma_r float64 = 2.5214456545125733e-5
var _cgos_a11_lgamma_r float64 = 4.4864094961891516e-5
var _cgos_tc_lgamma_r float64 = 1.4616321449683622
var _cgos_tf_lgamma_r float64 = -0.12148629053584961
var _cgos_tt_lgamma_r float64 = -3.6386769970395054e-18
var _cgos_t0_lgamma_r float64 = 0.48383612272381005
var _cgos_t1_lgamma_r float64 = -0.14758772299459391
var _cgos_t2_lgamma_r float64 = 0.064624940239133385
var _cgos_t3_lgamma_r float64 = -0.032788541075985965
var _cgos_t4_lgamma_r float64 = 0.017970675081182039
var _cgos_t5_lgamma_r float64 = -0.010314224129834144
var _cgos_t6_lgamma_r float64 = 0.0061005387024629133
var _cgos_t7_lgamma_r float64 = -0.0036845201678113826
var _cgos_t8_lgamma_r float64 = 0.0022596478090061247
var _cgos_t9_lgamma_r float64 = -0.0014034646998923284
var _cgos_t10_lgamma_r float64 = 8.8108188243765401e-4
var _cgos_t11_lgamma_r float64 = -5.3859530535674055e-4
var _cgos_t12_lgamma_r float64 = 3.1563207090362595e-4
var _cgos_t13_lgamma_r float64 = -3.1275416837512086e-4
var _cgos_t14_lgamma_r float64 = 3.3552919263551907e-4
var _cgos_u0_lgamma_r float64 = -0.077215664901532866
var _cgos_u1_lgamma_r float64 = 0.63282706402509337
var _cgos_u2_lgamma_r float64 = 1.4549225013723477
var _cgos_u3_lgamma_r float64 = 0.97771752796337275
var _cgos_u4_lgamma_r float64 = 0.22896372806469245
var _cgos_u5_lgamma_r float64 = 0.013381091853678766
var _cgos_v1_lgamma_r float64 = 2.4559779371304113
var _cgos_v2_lgamma_r float64 = 2.128489763798934
var _cgos_v3_lgamma_r float64 = 0.76928515045667278
var _cgos_v4_lgamma_r float64 = 0.10422264559336913
var _cgos_v5_lgamma_r float64 = 0.0032170924228242391
var _cgos_s0_lgamma_r float64 = -0.077215664901532866
var _cgos_s1_lgamma_r float64 = 0.21498241596060885
var _cgos_s2_lgamma_r float64 = 0.32577879640893098
var _cgos_s3_lgamma_r float64 = 0.14635047265246445
var _cgos_s4_lgamma_r float64 = 0.026642270303363861
var _cgos_s5_lgamma_r float64 = 0.0018402845140733772
var _cgos_s6_lgamma_r float64 = 3.1947532658410087e-5
var _cgos_r1_lgamma_r float64 = 1.3920053346762105
var _cgos_r2_lgamma_r float64 = 0.72193554756713807
var _cgos_r3_lgamma_r float64 = 0.17193386563280308
var _cgos_r4_lgamma_r float64 = 0.01864591917156529
var _cgos_r5_lgamma_r float64 = 7.779424963818936e-4
var _cgos_r6_lgamma_r float64 = 7.3266843074462564e-6
var _cgos_w0_lgamma_r float64 = 0.41893853320467273
var _cgos_w1_lgamma_r float64 = 0.083333333333332968
var _cgos_w2_lgamma_r float64 = -0.0027777777772877554
var _cgos_w3_lgamma_r float64 = 7.9365055864301956e-4
var _cgos_w4_lgamma_r float64 = -5.9518755745033996e-4
var _cgos_w5_lgamma_r float64 = 8.3633991899628213e-4
var _cgos_w6_lgamma_r float64 = -0.0016309293409657527

func _cgos_sin_pi_lgamma_r(x float64) float64 {
	var n int32
	x = 2 * (x*0.5 - Floor(x*0.5))
	n = int32(x * 4)
	n = (n + int32(1)) / int32(2)
	x -= float64(float32(n) * 0.5)
	x *= _cgos_pi_lgamma_r
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
	type _cgoa_18_lgamma_r struct {
		f float64
	}
	var u _cgoa_18_lgamma_r
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
		t = _cgos_sin_pi_lgamma_r(x)
		if t == 0 {
			return 1 / (x - x)
		}
		if t > 0 {
			*signgamp = -1
		} else {
			t = -t
		}
		nadj = Log(_cgos_pi_lgamma_r / (t * x))
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
				y = x - (_cgos_tc_lgamma_r - 1)
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
				y = x - _cgos_tc_lgamma_r
				i = int32(1)
			} else {
				y = x - 1
				i = int32(2)
			}
		}
		switch i {
		case int32(0):
			z = y * y
			p1 = _cgos_a0_lgamma_r + z*(_cgos_a2_lgamma_r+z*(_cgos_a4_lgamma_r+z*(_cgos_a6_lgamma_r+z*(_cgos_a8_lgamma_r+z*_cgos_a10_lgamma_r))))
			p2 = z * (_cgos_a1_lgamma_r + z*(_cgos_a3_lgamma_r+z*(_cgos_a5_lgamma_r+z*(_cgos_a7_lgamma_r+z*(_cgos_a9_lgamma_r+z*_cgos_a11_lgamma_r)))))
			p = y*p1 + p2
			r += p - 0.5*y
			break
		case int32(1):
			z = y * y
			w = z * y
			p1 = _cgos_t0_lgamma_r + w*(_cgos_t3_lgamma_r+w*(_cgos_t6_lgamma_r+w*(_cgos_t9_lgamma_r+w*_cgos_t12_lgamma_r)))
			p2 = _cgos_t1_lgamma_r + w*(_cgos_t4_lgamma_r+w*(_cgos_t7_lgamma_r+w*(_cgos_t10_lgamma_r+w*_cgos_t13_lgamma_r)))
			p3 = _cgos_t2_lgamma_r + w*(_cgos_t5_lgamma_r+w*(_cgos_t8_lgamma_r+w*(_cgos_t11_lgamma_r+w*_cgos_t14_lgamma_r)))
			p = z*p1 - (_cgos_tt_lgamma_r - w*(p2+y*p3))
			r += _cgos_tf_lgamma_r + p
			break
		case int32(2):
			p1 = y * (_cgos_u0_lgamma_r + y*(_cgos_u1_lgamma_r+y*(_cgos_u2_lgamma_r+y*(_cgos_u3_lgamma_r+y*(_cgos_u4_lgamma_r+y*_cgos_u5_lgamma_r)))))
			p2 = 1 + y*(_cgos_v1_lgamma_r+y*(_cgos_v2_lgamma_r+y*(_cgos_v3_lgamma_r+y*(_cgos_v4_lgamma_r+y*_cgos_v5_lgamma_r))))
			r += -0.5*y + p1/p2
		}
	} else if ix < uint32(1075838976) {
		i = int32(x)
		y = x - float64(i)
		p = y * (_cgos_s0_lgamma_r + y*(_cgos_s1_lgamma_r+y*(_cgos_s2_lgamma_r+y*(_cgos_s3_lgamma_r+y*(_cgos_s4_lgamma_r+y*(_cgos_s5_lgamma_r+y*_cgos_s6_lgamma_r))))))
		q = 1 + y*(_cgos_r1_lgamma_r+y*(_cgos_r2_lgamma_r+y*(_cgos_r3_lgamma_r+y*(_cgos_r4_lgamma_r+y*(_cgos_r5_lgamma_r+y*_cgos_r6_lgamma_r)))))
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
		w = _cgos_w0_lgamma_r + z*(_cgos_w1_lgamma_r+y*(_cgos_w2_lgamma_r+y*(_cgos_w3_lgamma_r+y*(_cgos_w4_lgamma_r+y*(_cgos_w5_lgamma_r+y*_cgos_w6_lgamma_r)))))
		r = (x-0.5)*(t-1) + w
	} else {
		r = x * (Log(x) - 1)
	}
	if sign != 0 {
		r = nadj - r
	}
	return r
}
