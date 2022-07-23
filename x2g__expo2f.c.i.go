package libc

import unsafe "unsafe"

const _cgos_k___expo2f int32 = int32(235)

var _cgos_kln2___expo2f float32 = 162.889587

func __expo2f(x float32, sign float32) float32 {
	var scale float32
	for {
		scale = *(*float32)(unsafe.Pointer(&_cgoz_18___expo2f{2046820352}))
		if true {
			break
		}
	}
	return Expf(x-_cgos_kln2___expo2f) * (sign * scale) * scale
}

type _cgoz_18___expo2f struct {
	_i uint32
}
