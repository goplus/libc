package libc

import unsafe "unsafe"

func _cgos_mul32_sqrt(a uint32, b uint32) uint32 {
	return uint32(uint64(a) * uint64(b) >> int32(32))
}
func _cgos_mul64_sqrt(a uint64, b uint64) uint64 {
	var ahi uint64 = a >> int32(32)
	var alo uint64 = a & uint64(4294967295)
	var bhi uint64 = b >> int32(32)
	var blo uint64 = b & uint64(4294967295)
	return ahi*bhi + ahi*blo>>int32(32) + alo*bhi>>int32(32)
}
func Sqrt(x float64) float64 {
	var ix uint64
	var top uint64
	var m uint64
	ix = *(*uint64)(unsafe.Pointer(&_cgoz_18_sqrt{x}))
	top = ix >> int32(52)
	if func() int64 {
		if top-uint64(1) >= uint64(2046) {
			return 1
		} else {
			return 0
		}
	}() == int64(0) {
		if ix*uint64(2) == uint64(0) {
			return x
		}
		if ix == uint64(9218868437227405312) {
			return x
		}
		if ix > uint64(9218868437227405312) {
			return __math_invalid(x)
		}
		ix = *(*uint64)(unsafe.Pointer(&_cgoz_19_sqrt{x * 4503599627370496}))
		top = ix >> int32(52)
		top -= uint64(52)
	}
	var even int32 = int32(top & uint64(1))
	m = ix<<int32(11) | uint64(9223372036854775808)
	if even != 0 {
		m >>= int32(1)
	}
	top = (top + uint64(1023)) >> int32(1)
	var r uint64
	var s uint64
	var d uint64
	var u uint64
	var i uint64
	i = ix >> int32(46) % uint64(128)
	r = uint64(uint32(*(*uint16)(unsafe.Pointer(uintptr(unsafe.Pointer((*uint16)(unsafe.Pointer(&__rsqrt_tab)))) + uintptr(i)*2))) << int32(16))
	s = uint64(_cgos_mul32_sqrt(uint32(m>>int32(32)), uint32(r)))
	d = uint64(_cgos_mul32_sqrt(uint32(s), uint32(r)))
	u = _cgos_three_sqrt - d
	r = uint64(_cgos_mul32_sqrt(uint32(r), uint32(u)) << int32(1))
	s = uint64(_cgos_mul32_sqrt(uint32(s), uint32(u)) << int32(1))
	d = uint64(_cgos_mul32_sqrt(uint32(s), uint32(r)))
	u = _cgos_three_sqrt - d
	r = uint64(_cgos_mul32_sqrt(uint32(r), uint32(u)) << int32(1))
	r = r << int32(32)
	s = _cgos_mul64_sqrt(m, r)
	d = _cgos_mul64_sqrt(s, r)
	u = _cgos_three_sqrt<<int32(32) - d
	s = _cgos_mul64_sqrt(s, u)
	s = (s - uint64(2)) >> int32(9)
	var d0 uint64
	var d1 uint64
	var d2 uint64
	var y float64
	var t float64
	d0 = m<<int32(42) - s*s
	d1 = s - d0
	d2 = d1 + s + uint64(1)
	s += d1 >> int32(63)
	s &= uint64(4503599627370495)
	s |= top << int32(52)
	y = *(*float64)(unsafe.Pointer(&_cgoz_20_sqrt{s}))
	if int32(1) != 0 {
		var tiny uint64 = uint64(func() int64 {
			if func() int64 {
				if d2 == uint64(0) {
					return 1
				} else {
					return 0
				}
			}() == int64(0) {
				return int64(0)
			} else {
				return int64(4503599627370496)
			}
		}())
		tiny |= (d1 ^ d2) & uint64(9223372036854775808)
		t = *(*float64)(unsafe.Pointer(&_cgoz_21_sqrt{tiny}))
		y = eval_as_double(y + t)
	}
	return y
}

type _cgoz_18_sqrt struct {
	_f float64
}
type _cgoz_19_sqrt struct {
	_f float64
}

var _cgos_three_sqrt uint64 = uint64(3221225472)

type _cgoz_20_sqrt struct {
	_i uint64
}
type _cgoz_21_sqrt struct {
	_i uint64
}
