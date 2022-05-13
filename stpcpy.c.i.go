package libc

import unsafe "unsafe"

func __stpcpy(d *int8, s *int8) *int8 {
	type word = uint64
	var wd *uint64
	var ws *uint64
	if uint64(uintptr(unsafe.Pointer(s)))%8 == uint64(uintptr(unsafe.Pointer(d)))%8 {
		for ; uint64(uintptr(unsafe.Pointer(s)))%8 != 0; func() *int8 {
			*(*uintptr)(unsafe.Pointer(&s))++
			return func() (_cgo_ret *int8) {
				_cgo_addr := &d
				_cgo_ret = *_cgo_addr
				*(*uintptr)(unsafe.Pointer(_cgo_addr))++
				return
			}()
		}() {
			if !(func() (_cgo_ret int8) {
				_cgo_addr := &*d
				*_cgo_addr = *s
				return *_cgo_addr
			}() != 0) {
				return d
			}
		}
		wd = (*uint64)(unsafe.Pointer(d))
		ws = (*uint64)(unsafe.Pointer(s))
		for ; !((*ws-uint64(18446744073709551615)/uint64(255)) & ^*ws & (uint64(18446744073709551615)/uint64(255)*uint64(255/2+1)) != 0); *func() (_cgo_ret *uint64) {
			_cgo_addr := &wd
			_cgo_ret = *_cgo_addr
			*(*uintptr)(unsafe.Pointer(_cgo_addr)) += 8
			return
		}() = *func() (_cgo_ret *uint64) {
			_cgo_addr := &ws
			_cgo_ret = *_cgo_addr
			*(*uintptr)(unsafe.Pointer(_cgo_addr)) += 8
			return
		}() {
		}
		d = (*int8)(unsafe.Pointer(wd))
		s = (*int8)(unsafe.Pointer(ws))
	}
	for ; func() (_cgo_ret int8) {
		_cgo_addr := &*d
		*_cgo_addr = *s
		return *_cgo_addr
	}() != 0; func() *int8 {
		*(*uintptr)(unsafe.Pointer(&s))++
		return func() (_cgo_ret *int8) {
			_cgo_addr := &d
			_cgo_ret = *_cgo_addr
			*(*uintptr)(unsafe.Pointer(_cgo_addr))++
			return
		}()
	}() {
	}
	return d
}
