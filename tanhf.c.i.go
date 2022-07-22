package libc

import unsafe "unsafe"

func Tanhf(x float32) float32 {
	type _cgoa_18_tanhf struct {
		f float32
	}
	var u _cgoa_18_tanhf
	u.f = x
	var w uint32
	var sign int32
	var t float32
	sign = int32(*(*uint32)(unsafe.Pointer(&u)) >> int32(31))
	*(*uint32)(unsafe.Pointer(&u)) &= uint32(2147483647)
	x = u.f
	w = *(*uint32)(unsafe.Pointer(&u))
	if w > uint32(1057791828) {
		if w > uint32(1092616192) {
			t = float32(int32(1)) + float32(int32(0))/x
		} else {
			t = Expm1f(float32(int32(2)) * x)
			t = float32(int32(1)) - float32(int32(2))/(t+float32(int32(2)))
		}
	} else if w > uint32(1048757624) {
		t = Expm1f(float32(int32(2)) * x)
		t = t / (t + float32(int32(2)))
	} else if w >= uint32(8388608) {
		t = Expm1f(float32(-2) * x)
		t = -t / (t + float32(int32(2)))
	} else {
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
		t = x
	}
	return func() float32 {
		if sign != 0 {
			return -t
		} else {
			return t
		}
	}()
}
