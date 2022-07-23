package libc

import unsafe "unsafe"

var _cgos_pi__atan2f float32 = float32(3.1415927410000002)
var _cgos_pi_lo__atan2f float32 = float32(-8.7422776572999997e-8)

func Atan2f(y float32, x float32) float32 {
	var z float32
	var m uint32
	var ix uint32
	var iy uint32
	if func() int32 {
		if 4 == 4 {
			return func() int32 {
				if X__FLOAT_BITS(x)&uint32(2147483647) > uint32(2139095040) {
					return 1
				} else {
					return 0
				}
			}()
		} else {
			return func() int32 {
				if 4 == 8 {
					return func() int32 {
						if X__DOUBLE_BITS(float64(x))&9223372036854775807 > 9218868437227405312 {
							return 1
						} else {
							return 0
						}
					}()
				} else {
					return func() int32 {
						if X__fpclassifyl(float64(x)) == int32(0) {
							return 1
						} else {
							return 0
						}
					}()
				}
			}()
		}
	}() != 0 || func() int32 {
		if 4 == 4 {
			return func() int32 {
				if X__FLOAT_BITS(y)&uint32(2147483647) > uint32(2139095040) {
					return 1
				} else {
					return 0
				}
			}()
		} else {
			return func() int32 {
				if 4 == 8 {
					return func() int32 {
						if X__DOUBLE_BITS(float64(y))&9223372036854775807 > 9218868437227405312 {
							return 1
						} else {
							return 0
						}
					}()
				} else {
					return func() int32 {
						if X__fpclassifyl(float64(y)) == int32(0) {
							return 1
						} else {
							return 0
						}
					}()
				}
			}()
		}
	}() != 0 {
		return x + y
	}
	for {
		ix = *(*uint32)(unsafe.Pointer(&_cgoz_18_atan2f{x}))
		if true {
			break
		}
	}
	for {
		iy = *(*uint32)(unsafe.Pointer(&_cgoz_19_atan2f{y}))
		if true {
			break
		}
	}
	if ix == uint32(1065353216) {
		return Atanf(y)
	}
	m = iy>>int32(31)&uint32(1) | ix>>int32(30)&uint32(2)
	ix &= uint32(2147483647)
	iy &= uint32(2147483647)
	if iy == uint32(0) {
		switch m {
		case uint32(0):
			fallthrough
		case uint32(1):
			return y
		case uint32(2):
			return _cgos_pi__atan2f
		case uint32(3):
			return -_cgos_pi__atan2f
		}
	}
	if ix == uint32(0) {
		return func() float32 {
			if m&uint32(1) != 0 {
				return -_cgos_pi__atan2f / float32(int32(2))
			} else {
				return _cgos_pi__atan2f / float32(int32(2))
			}
		}()
	}
	if ix == uint32(2139095040) {
		if iy == uint32(2139095040) {
			switch m {
			case uint32(0):
				return _cgos_pi__atan2f / float32(int32(4))
			case uint32(1):
				return -_cgos_pi__atan2f / float32(int32(4))
			case uint32(2):
				return float32(int32(3)) * _cgos_pi__atan2f / float32(int32(4))
			case uint32(3):
				return float32(-3) * _cgos_pi__atan2f / float32(int32(4))
			}
		} else {
			switch m {
			case uint32(0):
				return float32(0)
			case uint32(1):
				return float32(-0)
			case uint32(2):
				return _cgos_pi__atan2f
			case uint32(3):
				return -_cgos_pi__atan2f
			}
		}
	}
	if ix+uint32(218103808) < iy || iy == uint32(2139095040) {
		return func() float32 {
			if m&uint32(1) != 0 {
				return -_cgos_pi__atan2f / float32(int32(2))
			} else {
				return _cgos_pi__atan2f / float32(int32(2))
			}
		}()
	}
	if m&uint32(2) != 0 && iy+uint32(218103808) < ix {
		z = float32(0)
	} else {
		z = Atanf(Fabsf(y / x))
	}
	switch m {
	case uint32(0):
		return z
	case uint32(1):
		return -z
	case uint32(2):
		return _cgos_pi__atan2f - (z - _cgos_pi_lo__atan2f)
	default:
		return z - _cgos_pi_lo__atan2f - _cgos_pi__atan2f
	}
	return 0
}

type _cgoz_18_atan2f struct {
	_f float32
}
type _cgoz_19_atan2f struct {
	_f float32
}
