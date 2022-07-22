package libc

import unsafe "unsafe"

func Hypotf(x float32, y float32) float32 {
	type _cgoa_18_hypotf struct {
		f float32
	}
	var ux _cgoa_18_hypotf
	ux.f = x
	var uy _cgoa_18_hypotf
	uy.f = y
	var ut _cgoa_18_hypotf
	var z float32
	*(*uint32)(unsafe.Pointer(&ux)) &= 2147483647
	*(*uint32)(unsafe.Pointer(&uy)) &= 2147483647
	if *(*uint32)(unsafe.Pointer(&ux)) < *(*uint32)(unsafe.Pointer(&uy)) {
		ut = ux
		ux = uy
		uy = ut
	}
	x = ux.f
	y = uy.f
	if *(*uint32)(unsafe.Pointer(&uy)) == uint32(2139095040) {
		return y
	}
	if *(*uint32)(unsafe.Pointer(&ux)) >= uint32(2139095040) || *(*uint32)(unsafe.Pointer(&uy)) == uint32(0) || *(*uint32)(unsafe.Pointer(&ux))-*(*uint32)(unsafe.Pointer(&uy)) >= uint32(209715200) {
		return x + y
	}
	z = float32(int32(1))
	if *(*uint32)(unsafe.Pointer(&ux)) >= uint32(1568669696) {
		z = float32(1.23794004e+27)
		x *= float32(8.07793567e-28)
		y *= float32(8.07793567e-28)
	} else if *(*uint32)(unsafe.Pointer(&uy)) < uint32(562036736) {
		z = float32(8.07793567e-28)
		x *= float32(1.23794004e+27)
		y *= float32(1.23794004e+27)
	}
	return z * Sqrtf(float32(float64(x)*float64(x)+float64(y)*float64(y)))
}
