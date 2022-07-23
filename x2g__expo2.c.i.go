package libc

import unsafe "unsafe"

const _cgos_k___expo2 int32 = int32(2043)

var _cgos_kln2___expo2 float64 = 1416.0996898839683

func __expo2(x float64, sign float64) float64 {
	var scale float64
	for {
		scale = *(*float64)(unsafe.Pointer(&_cgoz_18___expo2{9205357638345293824}))
		if true {
			break
		}
	}
	return Exp(x-_cgos_kln2___expo2) * (sign * scale) * scale
}

type _cgoz_18___expo2 struct {
	_i uint64
}
