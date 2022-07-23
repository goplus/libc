package libc

import unsafe "unsafe"

var _cgos_B1__cbrt uint32 = uint32(715094163)
var _cgos_B2__cbrt uint32 = uint32(696219795)
var _cgos_P0__cbrt float64 = 1.8759518242717701
var _cgos_P1__cbrt float64 = -1.8849797954337717
var _cgos_P2__cbrt float64 = 1.6214297201053545
var _cgos_P3__cbrt float64 = -0.75839793477876605
var _cgos_P4__cbrt float64 = 0.14599619288661245

func Cbrt(x float64) float64 {
	type _cgoa_18_cbrt struct {
		f float64
	}
	var u _cgoa_18_cbrt
	u.f = x
	var r float64
	var s float64
	var t float64
	var w float64
	var hx uint32 = uint32(*(*uint64)(unsafe.Pointer(&u)) >> int32(32) & uint64(2147483647))
	if hx >= uint32(2146435072) {
		return x + x
	}
	if hx < uint32(1048576) {
		u.f = x * 18014398509481984
		hx = uint32(*(*uint64)(unsafe.Pointer(&u)) >> int32(32) & uint64(2147483647))
		if hx == uint32(0) {
			return x
		}
		hx = hx/uint32(3) + _cgos_B2__cbrt
	} else {
		hx = hx/uint32(3) + _cgos_B1__cbrt
	}
	*(*uint64)(unsafe.Pointer(&u)) &= 9223372036854775808
	*(*uint64)(unsafe.Pointer(&u)) |= uint64(hx) << int32(32)
	t = u.f
	r = t * t * (t / x)
	t = t * (_cgos_P0__cbrt + r*(_cgos_P1__cbrt+r*_cgos_P2__cbrt) + r*r*r*(_cgos_P3__cbrt+r*_cgos_P4__cbrt))
	u.f = t
	*(*uint64)(unsafe.Pointer(&u)) = (*(*uint64)(unsafe.Pointer(&u)) + uint64(2147483648)) & uint64(18446744072635809792)
	t = u.f
	s = t * t
	r = x / s
	w = t + t
	r = (r - t) / (w + r)
	t = t + t*r
	return t
}
