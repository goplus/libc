package libc

import unsafe "unsafe"

var B1_cgo18_cbrt uint32 = uint32(715094163)
var B2_cgo19_cbrt uint32 = uint32(696219795)
var P0_cgo20_cbrt float64 = 1.8759518242717701
var P1_cgo21_cbrt float64 = -1.8849797954337717
var P2_cgo22_cbrt float64 = 1.6214297201053545
var P3_cgo23_cbrt float64 = -0.75839793477876605
var P4_cgo24_cbrt float64 = 0.14599619288661245

func Cbrt(x float64) float64 {
	type _cgoa_25_cbrt struct {
		f float64
	}
	var u _cgoa_25_cbrt
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
		hx = hx/uint32(3) + B2_cgo19_cbrt
	} else {
		hx = hx/uint32(3) + B1_cgo18_cbrt
	}
	*(*uint64)(unsafe.Pointer(&u)) &= 9223372036854775808
	*(*uint64)(unsafe.Pointer(&u)) |= uint64(hx) << int32(32)
	t = u.f
	r = t * t * (t / x)
	t = t * (P0_cgo20_cbrt + r*(P1_cgo21_cbrt+r*P2_cgo22_cbrt) + r*r*r*(P3_cgo23_cbrt+r*P4_cgo24_cbrt))
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
