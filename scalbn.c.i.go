package libc

import unsafe "unsafe"

func scalbn(x float64, n int32) float64 {
	type _cgoa_176 struct {
		f float64
	}
	var u _cgoa_176
	var y float64 = x
	if n > int32(1023) {
		y *= float64(8.9884656743115795e+307)
		n -= int32(1023)
		if n > int32(1023) {
			y *= float64(8.9884656743115795e+307)
			n -= int32(1023)
			if n > int32(1023) {
				n = int32(1023)
			}
		}
	} else if n < -1022 {
		y *= float64(2.2250738585072014e-308 * 9007199254740992)
		n += 969
		if n < -1022 {
			y *= float64(2.2250738585072014e-308 * 9007199254740992)
			n += 969
			if n < -1022 {
				n = -1022
			}
		}
	}
	*(*uint64)(unsafe.Pointer(&u)) = uint64(int32(1023)+n) << int32(52)
	x = y * u.f
	return x
}
