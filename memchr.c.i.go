package libc

import unsafe "unsafe"

func Memchr(src unsafe.Pointer, c int32, n uint64) unsafe.Pointer {
	var s *uint8 = (*uint8)(src)
	c = int32(uint8(c))
	for ; uint64(uintptr(unsafe.Pointer(s)))&7 != 0 && n != 0 && int32(*s) != c; func() uint64 {
		*(*uintptr)(unsafe.Pointer(&s))++
		return func() (_cgo_ret uint64) {
			_cgo_addr := &n
			_cgo_ret = *_cgo_addr
			*_cgo_addr--
			return
		}()
	}() {
	}
	if n != 0 && int32(*s) != c {
		type word = uint64
		var w *uint64
		var k uint64 = 72340172838076673 * uint64(c)
		for w = (*uint64)(unsafe.Pointer(s)); n >= 8 && !((*w^k-72340172838076673) & ^(*w^k) & 9259542123273814144 != 0); func() uint64 {
			*(*uintptr)(unsafe.Pointer(&w)) += 8
			return func() (_cgo_ret uint64) {
				_cgo_addr := &n
				*_cgo_addr -= uint64(8)
				return *_cgo_addr
			}()
		}() {
		}
		s = (*uint8)(unsafe.Pointer(w))
	}
	for ; n != 0 && int32(*s) != c; func() uint64 {
		*(*uintptr)(unsafe.Pointer(&s))++
		return func() (_cgo_ret uint64) {
			_cgo_addr := &n
			_cgo_ret = *_cgo_addr
			*_cgo_addr--
			return
		}()
	}() {
	}
	return func() unsafe.Pointer {
		if n != 0 {
			return unsafe.Pointer(s)
		} else {
			return nil
		}
	}()
}
