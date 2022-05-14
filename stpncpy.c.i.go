package libc

import unsafe "unsafe"

func __stpncpy(d *int8, s *int8, n uint64) *int8 {
	type word = uint64
	var wd *uint64
	var ws *uint64
	if uint64(uintptr(unsafe.Pointer(s)))&7 == uint64(uintptr(unsafe.Pointer(d)))&7 {
		for ; uint64(uintptr(unsafe.Pointer(s)))&7 != 0 && n != 0 && int32(func() (_cgo_ret int8) {
			_cgo_addr := &*d
			*_cgo_addr = *s
			return *_cgo_addr
		}()) != 0; func() *int8 {
			func() *int8 {
				n--
				return func() (_cgo_ret *int8) {
					_cgo_addr := &s
					_cgo_ret = *_cgo_addr
					*(*uintptr)(unsafe.Pointer(_cgo_addr))++
					return
				}()
			}()
			return func() (_cgo_ret *int8) {
				_cgo_addr := &d
				_cgo_ret = *_cgo_addr
				*(*uintptr)(unsafe.Pointer(_cgo_addr))++
				return
			}()
		}() {
		}
		if !(n != 0) || !(*s != 0) {
			goto tail
		}
		wd = (*uint64)(unsafe.Pointer(d))
		ws = (*uint64)(unsafe.Pointer(s))
		for ; n >= 8 && !((*ws-72340172838076673) & ^*ws & 9259542123273814144 != 0); func() *uint64 {
			func() *uint64 {
				n -= uint64(8)
				return func() (_cgo_ret *uint64) {
					_cgo_addr := &ws
					_cgo_ret = *_cgo_addr
					*(*uintptr)(unsafe.Pointer(_cgo_addr)) += 8
					return
				}()
			}()
			return func() (_cgo_ret *uint64) {
				_cgo_addr := &wd
				_cgo_ret = *_cgo_addr
				*(*uintptr)(unsafe.Pointer(_cgo_addr)) += 8
				return
			}()
		}() {
			*wd = *ws
		}
		d = (*int8)(unsafe.Pointer(wd))
		s = (*int8)(unsafe.Pointer(ws))
	}
	for ; n != 0 && int32(func() (_cgo_ret int8) {
		_cgo_addr := &*d
		*_cgo_addr = *s
		return *_cgo_addr
	}()) != 0; func() *int8 {
		func() *int8 {
			n--
			return func() (_cgo_ret *int8) {
				_cgo_addr := &s
				_cgo_ret = *_cgo_addr
				*(*uintptr)(unsafe.Pointer(_cgo_addr))++
				return
			}()
		}()
		return func() (_cgo_ret *int8) {
			_cgo_addr := &d
			_cgo_ret = *_cgo_addr
			*(*uintptr)(unsafe.Pointer(_cgo_addr))++
			return
		}()
	}() {
	}
tail:
	Memset(unsafe.Pointer(d), int32(0), n)
	return d
}
