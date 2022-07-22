package libc

import unsafe "unsafe"

var B1_cgo18_cbrtf uint32 = uint32(709958130)
var B2_cgo19_cbrtf uint32 = uint32(642849266)

func Cbrtf(x float32) float32 {
	var r float64
	var T float64
	type _cgoa_20_cbrtf struct {
		f float32
	}
	var u _cgoa_20_cbrtf
	u.f = x
	var hx uint32 = *(*uint32)(unsafe.Pointer(&u)) & uint32(2147483647)
	if hx >= uint32(2139095040) {
		return x + x
	}
	if hx < uint32(8388608) {
		if hx == uint32(0) {
			return x
		}
		u.f = x * 16777216
		hx = *(*uint32)(unsafe.Pointer(&u)) & uint32(2147483647)
		hx = hx/uint32(3) + B2_cgo19_cbrtf
	} else {
		hx = hx/uint32(3) + B1_cgo18_cbrtf
	}
	*(*uint32)(unsafe.Pointer(&u)) &= uint32(2147483648)
	*(*uint32)(unsafe.Pointer(&u)) |= hx
	T = float64(u.f)
	r = T * T * T
	T = T * (float64(x) + float64(x) + r) / (float64(x) + r + r)
	r = T * T * T
	T = T * (float64(x) + float64(x) + r) / (float64(x) + r + r)
	return float32(T)
}
