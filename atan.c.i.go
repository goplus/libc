package libc

import unsafe "unsafe"

var atanhi_cgos__atan [4]float64 = [4]float64{0.46364760900080609, 0.78539816339744828, 0.98279372324732905, 1.5707963267948966}
var atanlo_cgos__atan [4]float64 = [4]float64{2.2698777452961687e-17, 3.061616997868383e-17, 1.3903311031230998e-17, 6.123233995736766e-17}
var aT_cgos__atan [11]float64 = [11]float64{0.33333333333332932, -0.19999999999876483, 0.14285714272503466, -0.11111110405462356, 0.090908871334365065, -0.0769187620504483, 0.066610731373875312, -0.058335701337905735, 0.049768779946159324, -0.036531572744216916, 0.016285820115365782}

func Atan(x float64) float64 {
	var w float64
	var s1 float64
	var s2 float64
	var z float64
	var ix uint32
	var sign uint32
	var id int32
	for {
		ix = uint32(*(*uint64)(unsafe.Pointer(&_cgoz_18_atan{x})) >> int32(32))
		if true {
			break
		}
	}
	sign = ix >> int32(31)
	ix &= uint32(2147483647)
	if ix >= uint32(1141899264) {
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
		}() != 0 {
			return x
		}
		z = *(*float64)(unsafe.Pointer(uintptr(unsafe.Pointer((*float64)(unsafe.Pointer(&atanhi_cgos__atan)))) + uintptr(int32(3))*8)) + float64(7.52316385e-37)
		return func() float64 {
			if sign != 0 {
				return -z
			} else {
				return z
			}
		}()
	}
	if ix < uint32(1071382528) {
		if ix < uint32(1044381696) {
			if ix < uint32(1048576) {
				for {
					if 4 == 4 {
						fp_force_evalf(float32(x))
					} else if 4 == 8 {
						fp_force_eval(float64(float32(x)))
					} else {
						fp_force_evall(float64(float32(x)))
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
		x = Fabs(x)
		if ix < uint32(1072889856) {
			if ix < uint32(1072037888) {
				id = int32(0)
				x = (2*x - 1) / (2 + x)
			} else {
				id = int32(1)
				x = (x - 1) / (x + 1)
			}
		} else if ix < uint32(1073971200) {
			id = int32(2)
			x = (x - 1.5) / (1 + 1.5*x)
		} else {
			id = int32(3)
			x = -1 / x
		}
	}
	z = x * x
	w = z * z
	s1 = z * (*(*float64)(unsafe.Pointer(uintptr(unsafe.Pointer((*float64)(unsafe.Pointer(&aT_cgos__atan)))) + uintptr(int32(0))*8)) + w*(*(*float64)(unsafe.Pointer(uintptr(unsafe.Pointer((*float64)(unsafe.Pointer(&aT_cgos__atan)))) + uintptr(int32(2))*8))+w*(*(*float64)(unsafe.Pointer(uintptr(unsafe.Pointer((*float64)(unsafe.Pointer(&aT_cgos__atan)))) + uintptr(int32(4))*8))+w*(*(*float64)(unsafe.Pointer(uintptr(unsafe.Pointer((*float64)(unsafe.Pointer(&aT_cgos__atan)))) + uintptr(int32(6))*8))+w*(*(*float64)(unsafe.Pointer(uintptr(unsafe.Pointer((*float64)(unsafe.Pointer(&aT_cgos__atan)))) + uintptr(int32(8))*8))+w**(*float64)(unsafe.Pointer(uintptr(unsafe.Pointer((*float64)(unsafe.Pointer(&aT_cgos__atan)))) + uintptr(int32(10))*8)))))))
	s2 = w * (*(*float64)(unsafe.Pointer(uintptr(unsafe.Pointer((*float64)(unsafe.Pointer(&aT_cgos__atan)))) + uintptr(int32(1))*8)) + w*(*(*float64)(unsafe.Pointer(uintptr(unsafe.Pointer((*float64)(unsafe.Pointer(&aT_cgos__atan)))) + uintptr(int32(3))*8))+w*(*(*float64)(unsafe.Pointer(uintptr(unsafe.Pointer((*float64)(unsafe.Pointer(&aT_cgos__atan)))) + uintptr(int32(5))*8))+w*(*(*float64)(unsafe.Pointer(uintptr(unsafe.Pointer((*float64)(unsafe.Pointer(&aT_cgos__atan)))) + uintptr(int32(7))*8))+w**(*float64)(unsafe.Pointer(uintptr(unsafe.Pointer((*float64)(unsafe.Pointer(&aT_cgos__atan)))) + uintptr(int32(9))*8))))))
	if id < int32(0) {
		return x - x*(s1+s2)
	}
	z = *(*float64)(unsafe.Pointer(uintptr(unsafe.Pointer((*float64)(unsafe.Pointer(&atanhi_cgos__atan)))) + uintptr(id)*8)) - (x*(s1+s2) - *(*float64)(unsafe.Pointer(uintptr(unsafe.Pointer((*float64)(unsafe.Pointer(&atanlo_cgos__atan)))) + uintptr(id)*8)) - x)
	return func() float64 {
		if sign != 0 {
			return -z
		} else {
			return z
		}
	}()
}

type _cgoz_18_atan struct {
	_f float64
}
