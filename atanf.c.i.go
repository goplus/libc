package libc

import unsafe "unsafe"

var atanhi_cgos__atanf [4]float32 = [4]float32{float32(0.46364760398999999), float32(0.78539812565), float32(0.98279368876999995), float32(1.5707962513)}
var atanlo_cgos__atanf [4]float32 = [4]float32{float32(5.0121582440000004e-9), float32(3.7748947078999999e-8), float32(3.4473217170000002e-8), float32(7.5497894159e-8)}
var aT_cgos__atanf [5]float32 = [5]float32{float32(0.33333328365999998), float32(-0.19999158382000001), float32(0.14253635705000001), float32(-0.10648017377000001), float32(0.061687607318)}

func Atanf(x float32) float32 {
	var w float32
	var s1 float32
	var s2 float32
	var z float32
	var ix uint32
	var sign uint32
	var id int32
	for {
		ix = *(*uint32)(unsafe.Pointer(&_cgoz_18_atanf{x}))
		if true {
			break
		}
	}
	sign = ix >> int32(31)
	ix &= uint32(2147483647)
	if ix >= uint32(1283457024) {
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
		}() != 0 {
			return x
		}
		z = *(*float32)(unsafe.Pointer(uintptr(unsafe.Pointer((*float32)(unsafe.Pointer(&atanhi_cgos__atanf)))) + uintptr(int32(3))*4)) + 7.52316385e-37
		return func() float32 {
			if sign != 0 {
				return -z
			} else {
				return z
			}
		}()
	}
	if ix < uint32(1054867456) {
		if ix < uint32(964689920) {
			if ix < uint32(8388608) {
				for {
					if 4 == 4 {
						fp_force_evalf(x * x)
					} else if 4 == 8 {
						fp_force_eval(float64(x * x))
					} else {
						fp_force_evall(float64(x * x))
					}
					if true {
						break
					}
				}
			}
			return x
		}
		id = -1
	} else {
		x = Fabsf(x)
		if ix < uint32(1066926080) {
			if ix < uint32(1060110336) {
				id = int32(0)
				x = (2*x - 1) / (2 + x)
			} else {
				id = int32(1)
				x = (x - 1) / (x + 1)
			}
		} else if ix < uint32(1075576832) {
			id = int32(2)
			x = (x - 1.5) / (1 + 1.5*x)
		} else {
			id = int32(3)
			x = -1 / x
		}
	}
	z = x * x
	w = z * z
	s1 = z * (*(*float32)(unsafe.Pointer(uintptr(unsafe.Pointer((*float32)(unsafe.Pointer(&aT_cgos__atanf)))) + uintptr(int32(0))*4)) + w*(*(*float32)(unsafe.Pointer(uintptr(unsafe.Pointer((*float32)(unsafe.Pointer(&aT_cgos__atanf)))) + uintptr(int32(2))*4))+w**(*float32)(unsafe.Pointer(uintptr(unsafe.Pointer((*float32)(unsafe.Pointer(&aT_cgos__atanf)))) + uintptr(int32(4))*4))))
	s2 = w * (*(*float32)(unsafe.Pointer(uintptr(unsafe.Pointer((*float32)(unsafe.Pointer(&aT_cgos__atanf)))) + uintptr(int32(1))*4)) + w**(*float32)(unsafe.Pointer(uintptr(unsafe.Pointer((*float32)(unsafe.Pointer(&aT_cgos__atanf)))) + uintptr(int32(3))*4)))
	if id < int32(0) {
		return x - x*(s1+s2)
	}
	z = *(*float32)(unsafe.Pointer(uintptr(unsafe.Pointer((*float32)(unsafe.Pointer(&atanhi_cgos__atanf)))) + uintptr(id)*4)) - (x*(s1+s2) - *(*float32)(unsafe.Pointer(uintptr(unsafe.Pointer((*float32)(unsafe.Pointer(&atanlo_cgos__atanf)))) + uintptr(id)*4)) - x)
	return func() float32 {
		if sign != 0 {
			return -z
		} else {
			return z
		}
	}()
}

type _cgoz_18_atanf struct {
	_f float32
}
