package libc

import unsafe "unsafe"

func _cgos_strtox_strtod(s *int8, p **int8, prec int32) float64 {
	var f Struct__IO_FILE
	func() *uint8 {
		(&f).Buf = func() (_cgo_ret *uint8) {
			_cgo_addr := &(&f).Rpos
			*_cgo_addr = (*uint8)(unsafe.Pointer(s))
			return *_cgo_addr
		}()
		return func() (_cgo_ret *uint8) {
			_cgo_addr := &(&f).Rend
			*_cgo_addr = (*uint8)(unsafe.Pointer(uintptr(18446744073709551615)))
			return *_cgo_addr
		}()
	}()
	__shlim(&f, int64(0))
	var y float64 = __floatscan(&f, prec, int32(1))
	var cnt int64 = (&f).Shcnt + int64(uintptr(unsafe.Pointer((&f).Rpos))-uintptr(unsafe.Pointer((&f).Buf)))
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
	return float32(_cgos_strtox_strtod(s, p, int32(0)))
}
func Strtod(s *int8, p **int8) float64 {
	return float64(_cgos_strtox_strtod(s, p, int32(1)))
}
func Strtold(s *int8, p **int8) float64 {
	return _cgos_strtox_strtod(s, p, int32(2))
}
