package libc

import unsafe "unsafe"

func strtox_cgo18_strtod(s *int8, p **int8, prec int32) float64 {
	var f struct__IO_FILE
	func() *uint8 {
		(&f).buf = func() (_cgo_ret *uint8) {
			_cgo_addr := &(&f).rpos
			*_cgo_addr = (*uint8)(unsafe.Pointer(s))
			return *_cgo_addr
		}()
		return func() (_cgo_ret *uint8) {
			_cgo_addr := &(&f).rend
			*_cgo_addr = (*uint8)(unsafe.Pointer(uintptr(18446744073709551615)))
			return *_cgo_addr
		}()
	}()
	__shlim(&f, int64(0))
	var y float64 = __floatscan(&f, prec, int32(1))
	var cnt int64 = (&f).shcnt + int64(uintptr(unsafe.Pointer((&f).rpos))-uintptr(unsafe.Pointer((&f).buf)))
	if p != nil {
		*p = func() *int8 {
			if cnt != 0 {
				return (*int8)(unsafe.Pointer(uintptr(unsafe.Pointer((*int8)(unsafe.Pointer(s)))) + uintptr(cnt)))
			} else {
				return (*int8)(unsafe.Pointer(s))
			}
		}()
	}
	return y
}
func Strtof(s *int8, p **int8) float32 {
	return float32(strtox_cgo18_strtod(s, p, int32(0)))
}
func Strtod(s *int8, p **int8) float64 {
	return float64(strtox_cgo18_strtod(s, p, int32(1)))
}
func Strtold(s *int8, p **int8) float64 {
	return strtox_cgo18_strtod(s, p, int32(2))
}
