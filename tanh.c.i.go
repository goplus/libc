package libc

import unsafe "unsafe"

func Tanh(x float64) float64 {
	type _cgoa_18_tanh struct {
		f float64
	}
	var u _cgoa_18_tanh
	u.f = x
	var w uint32
	var sign int32
	var t float64
	sign = int32(*(*uint64)(unsafe.Pointer(&u)) >> int32(63))
	*(*uint64)(unsafe.Pointer(&u)) &= 9223372036854775807
	x = u.f
	w = uint32(*(*uint64)(unsafe.Pointer(&u)) >> int32(32))
	if w > uint32(1071748074) {
		if w > uint32(1077149696) {
			t = float64(int32(1)) - float64(int32(0))/x
		} else {
			t = Expm1(float64(int32(2)) * x)
			t = float64(int32(1)) - float64(int32(2))/(t+float64(int32(2)))
		}
	} else if w > uint32(1070618798) {
		t = Expm1(float64(int32(2)) * x)
		t = t / (t + float64(int32(2)))
	} else if w >= uint32(1048576) {
		t = Expm1(float64(-2) * x)
		t = -t / (t + float64(int32(2)))
	} else {
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
		t = x
	}
	return func() float64 {
		if sign != 0 {
			return -t
		} else {
			return t
		}
	}()
}
