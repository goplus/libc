package libc

import unsafe "unsafe"

func Scalbnf(x float32, n int32) float32 {
	type _cgoa_18_scalbnf struct {
		f float32
	}
	var u _cgoa_18_scalbnf
	var y float32 = x
	if n > int32(127) {
		y *= float32(1.70141183e+38)
		n -= int32(127)
		if n > int32(127) {
			y *= float32(1.70141183e+38)
			n -= int32(127)
			if n > int32(127) {
				n = int32(127)
			}
		}
	} else if n < -126 {
		y *= float32(1.17549435e-38 * 16777216)
		n += 102
		if n < -126 {
			y *= float32(1.17549435e-38 * 16777216)
			n += 102
			if n < -126 {
				n = -126
			}
		}
	}
	*(*uint32)(unsafe.Pointer(&u)) = uint32(int32(127)+n) << int32(23)
	x = y * u.f
	return x
}
