package libc

import unsafe "unsafe"

func _cgos_mul32_sqrtf(a uint32, b uint32) uint32 {
	return uint32(uint64(a) * uint64(b) >> int32(32))
}
func Sqrtf(x float32) float32 {
	var ix uint32
	var m uint32
	var m1 uint32
	var m0 uint32
	var even uint32
	var ey uint32
	ix = *(*uint32)(unsafe.Pointer(&_cgoz_18_sqrtf{x}))
	if func() int64 {
		if ix-uint32(8388608) >= uint32(2130706432) {
			return 1
		} else {
			return 0
		}
	}() == int64(0) {
		if ix*uint32(2) == uint32(0) {
			return x
		}
		if ix == uint32(2139095040) {
			return x
		}
		if ix > uint32(2139095040) {
			return __math_invalidf(x)
		}
		ix = *(*uint32)(unsafe.Pointer(&_cgoz_19_sqrtf{x * 8388608.0}))
		ix -= uint32(192937984)
	}
	even = ix & uint32(8388608)
	m1 = ix<<int32(8) | uint32(2147483648)
	m0 = ix << int32(7) & uint32(2147483647)
	m = func() uint32 {
		if even != 0 {
			return m0
		} else {
			return m1
		}
	}()
	ey = ix >> int32(1)
	ey += uint32(532676608)
	ey &= uint32(2139095040)
	const _cgos_three_sqrtf uint32 = uint32(3221225472)
	var r uint32
	var s uint32
	var d uint32
	var u uint32
	var i uint32
	i = ix >> int32(17) % uint32(128)
	r = uint32(*(*uint16)(unsafe.Pointer(uintptr(unsafe.Pointer((*uint16)(unsafe.Pointer(&__rsqrt_tab)))) + uintptr(i)*2))) << int32(16)
	s = _cgos_mul32_sqrtf(m, r)
	d = _cgos_mul32_sqrtf(s, r)
	u = _cgos_three_sqrtf - d
	r = _cgos_mul32_sqrtf(r, u) << int32(1)
	s = _cgos_mul32_sqrtf(s, u) << int32(1)
	d = _cgos_mul32_sqrtf(s, r)
	u = _cgos_three_sqrtf - d
	s = _cgos_mul32_sqrtf(s, u)
	s = (s - uint32(1)) >> int32(6)
	var d0 uint32
	var d1 uint32
	var d2 uint32
	var y float32
	var t float32
	d0 = m<<int32(16) - s*s
	d1 = s - d0
	d2 = d1 + s + uint32(1)
	s += d1 >> int32(31)
	s &= uint32(8388607)
	s |= ey
	y = *(*float32)(unsafe.Pointer(&_cgoz_20_sqrtf{s}))
	if int32(1) != 0 {
		var tiny uint32 = uint32(func() int32 {
			if func() int64 {
				if d2 == uint32(0) {
					return 1
				} else {
					return 0
				}
			}() == int64(0) {
				return int32(0)
			} else {
				return int32(16777216)
			}
		}())
		tiny |= (d1 ^ d2) & uint32(2147483648)
		t = *(*float32)(unsafe.Pointer(&_cgoz_21_sqrtf{tiny}))
		y = eval_as_float(y + t)
	}
	return y
}

type _cgoz_18_sqrtf struct {
	_f float32
}
type _cgoz_19_sqrtf struct {
	_f float32
}
type _cgoz_20_sqrtf struct {
	_i uint32
}
type _cgoz_21_sqrtf struct {
	_i uint32
}
