package libc

import unsafe "unsafe"

func Memccpy(dest unsafe.Pointer, src unsafe.Pointer, c int32, n uint64) unsafe.Pointer {
	var d *uint8 = (*uint8)(dest)
	var s *uint8 = (*uint8)(src)
	c = int32(uint8(c))
	type word = uint64
	var wd *uint64
	var ws *uint64
	if uint64(uintptr(unsafe.Pointer(s)))&7 == uint64(uintptr(unsafe.Pointer(d)))&7 {
		for ; uint64(uintptr(unsafe.Pointer(s)))&7 != 0 && n != 0 && int32(func() (_cgo_ret uint8) {
			_cgo_addr := &*d
			*_cgo_addr = *s
			return *_cgo_addr
		}()) != c; func() *uint8 {
			func() *uint8 {
				n--
				return func() (_cgo_ret *uint8) {
					_cgo_addr := &s
					_cgo_ret = *_cgo_addr
					*(*uintptr)(unsafe.Pointer(_cgo_addr))++
					return
				}()
			}()
			return func() (_cgo_ret *uint8) {
				_cgo_addr := &d
				_cgo_ret = *_cgo_addr
				*(*uintptr)(unsafe.Pointer(_cgo_addr))++
				return
			}()
		}() {
		}
		if uint64(uintptr(unsafe.Pointer(s)))&7 != 0 {
			goto tail
		}
		var k uint64 = 72340172838076673 * uint64(c)
		wd = (*uint64)(unsafe.Pointer(d))
		ws = (*uint64)(unsafe.Pointer(s))
		for ; n >= 8 && !((*ws^k-72340172838076673) & ^(*ws^k) & 9259542123273814144 != 0); func() *uint64 {
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
		d = (*uint8)(unsafe.Pointer(wd))
		s = (*uint8)(unsafe.Pointer(ws))
	}
	for ; n != 0 && int32(func() (_cgo_ret uint8) {
		_cgo_addr := &*d
		*_cgo_addr = *s
		return *_cgo_addr
	}()) != c; func() *uint8 {
		func() *uint8 {
			n--
			return func() (_cgo_ret *uint8) {
				_cgo_addr := &s
				_cgo_ret = *_cgo_addr
				*(*uintptr)(unsafe.Pointer(_cgo_addr))++
				return
			}()
		}()
		return func() (_cgo_ret *uint8) {
			_cgo_addr := &d
			_cgo_ret = *_cgo_addr
			*(*uintptr)(unsafe.Pointer(_cgo_addr))++
			return
		}()
	}() {
	}
tail:
	if n != 0 {
		return unsafe.Pointer((*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(d)) + uintptr(int32(1)))))
	}
	return unsafe.Pointer(nil)
}
