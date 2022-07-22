package libc

import unsafe "unsafe"

const k_cgo18___expo2f int32 = int32(235)

var kln2_cgo19___expo2f float32 = 162.889587

func __expo2f(x float32, sign float32) float32 {
	var scale float32
	for {
		scale = *(*float32)(unsafe.Pointer(&_cgoz_20___expo2f{2046820352}))
		if true {
			break
		}
	}
	return Expf(x-kln2_cgo19___expo2f) * (sign * scale) * scale
}

type _cgoz_20___expo2f struct {
	_i uint32
}
