package libc

import unsafe "unsafe"

func _cgos_sq__hypot(hi *float64, lo *float64, x float64) {
	var xh float64
	var xl float64
	var xc float64
	xc = float64(x) * (134217728 + float64(int32(1)))
	xh = x - xc + xc
	xl = x - xh
	*hi = float64(x) * x
	*lo = xh*xh - *hi + float64(int32(2))*xh*xl + xl*xl
}
func Hypot(x float64, y float64) float64 {
	type _cgoa_18_hypot struct {
		f float64
	}
	var ux _cgoa_18_hypot
	ux.f = x
	var uy _cgoa_18_hypot
	uy.f = y
	var ut _cgoa_18_hypot
	var ex int32
	var ey int32
	var hx float64
	var lx float64
	var hy float64
	var ly float64
	var z float64
	*(*uint64)(unsafe.Pointer(&ux)) &= 9223372036854775807
	*(*uint64)(unsafe.Pointer(&uy)) &= 9223372036854775807
	if *(*uint64)(unsafe.Pointer(&ux)) < *(*uint64)(unsafe.Pointer(&uy)) {
		ut = ux
		ux = uy
		uy = ut
	}
	ex = int32(*(*uint64)(unsafe.Pointer(&ux)) >> int32(52))
	ey = int32(*(*uint64)(unsafe.Pointer(&uy)) >> int32(52))
	x = ux.f
	y = uy.f
	if ey == int32(2047) {
		return y
	}
	if ex == int32(2047) || *(*uint64)(unsafe.Pointer(&uy)) == uint64(0) {
		return x
	}
	if ex-ey > int32(64) {
		return x + y
	}
	z = float64(int32(1))
	if ex > 1533 {
		z = float64(5.2601359015483735e+210)
		x *= float64(1.9010915662951598e-211)
		y *= float64(1.9010915662951598e-211)
	} else if ey < 573 {
		z = float64(1.9010915662951598e-211)
		x *= float64(5.2601359015483735e+210)
		y *= float64(5.2601359015483735e+210)
	}
	_cgos_sq__hypot(&hx, &lx, x)
	_cgos_sq__hypot(&hy, &ly, y)
	return z * Sqrt(ly+lx+hy+hx)
}
