package libc

import unsafe "unsafe"

var _cgos_pi_atan2 float64 = 3.1415926535897931
var _cgos_pi_lo_atan2 float64 = 1.2246467991473532e-16

func Atan2(y float64, x float64) float64 {
	var z float64
	var m uint32
	var lx uint32
	var ly uint32
	var ix uint32
	var iy uint32
	if func() int32 {
		if 8 == 4 {
			return func() int32 {
				if X__FLOAT_BITS(float32(x))&uint32(2147483647) > uint32(2139095040) {
					return 1
				} else {
					return 0
				}
			}()
		} else {
			return func() int32 {
				if 8 == 8 {
					return func() int32 {
						if X__DOUBLE_BITS(x)&9223372036854775807 > 9218868437227405312 {
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
		if 8 == 4 {
			return func() int32 {
				if X__FLOAT_BITS(float32(y))&uint32(2147483647) > uint32(2139095040) {
					return 1
				} else {
					return 0
				}
			}()
		} else {
			return func() int32 {
				if 8 == 8 {
					return func() int32 {
						if X__DOUBLE_BITS(y)&9223372036854775807 > 9218868437227405312 {
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
		var __u uint64 = *(*uint64)(unsafe.Pointer(&_cgoz_18_atan2{x}))
		ix = uint32(__u >> int32(32))
		lx = uint32(__u)
		if true {
			break
		}
	}
	for {
		var __u uint64 = *(*uint64)(unsafe.Pointer(&_cgoz_19_atan2{y}))
		iy = uint32(__u >> int32(32))
		ly = uint32(__u)
		if true {
			break
		}
	}
	if ix-uint32(1072693248)|lx == uint32(0) {
		return Atan(y)
	}
	m = iy>>int32(31)&uint32(1) | ix>>int32(30)&uint32(2)
	ix = ix & uint32(2147483647)
	iy = iy & uint32(2147483647)
	if iy|ly == uint32(0) {
		switch m {
		case uint32(0):
			fallthrough
		case uint32(1):
			return y
		case uint32(2):
			return _cgos_pi_atan2
		case uint32(3):
			return -_cgos_pi_atan2
		}
	}
	if ix|lx == uint32(0) {
		return func() float64 {
			if m&uint32(1) != 0 {
				return -_cgos_pi_atan2 / float64(int32(2))
			} else {
				return _cgos_pi_atan2 / float64(int32(2))
			}
		}()
	}
	if ix == uint32(2146435072) {
		if iy == uint32(2146435072) {
			switch m {
			case uint32(0):
				return _cgos_pi_atan2 / float64(int32(4))
			case uint32(1):
				return -_cgos_pi_atan2 / float64(int32(4))
			case uint32(2):
				return float64(int32(3)) * _cgos_pi_atan2 / float64(int32(4))
			case uint32(3):
				return float64(-3) * _cgos_pi_atan2 / float64(int32(4))
			}
		} else {
			switch m {
			case uint32(0):
				return float64(0.0)
			case uint32(1):
				return float64(-0.0)
			case uint32(2):
				return _cgos_pi_atan2
			case uint32(3):
				return -_cgos_pi_atan2
			}
		}
	}
	if ix+uint32(67108864) < iy || iy == uint32(2146435072) {
		return func() float64 {
			if m&uint32(1) != 0 {
				return -_cgos_pi_atan2 / float64(int32(2))
			} else {
				return _cgos_pi_atan2 / float64(int32(2))
			}
		}()
	}
	if m&uint32(2) != 0 && iy+uint32(67108864) < ix {
		z = float64(int32(0))
	} else {
		z = Atan(Fabs(y / x))
	}
	switch m {
	case uint32(0):
		return z
	case uint32(1):
		return -z
	case uint32(2):
		return _cgos_pi_atan2 - (z - _cgos_pi_lo_atan2)
	default:
		return z - _cgos_pi_lo_atan2 - _cgos_pi_atan2
	}
	return 0
}

type _cgoz_18_atan2 struct {
	_f float64
}
type _cgoz_19_atan2 struct {
	_f float64
}
