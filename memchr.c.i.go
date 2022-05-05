package libc

import unsafe "unsafe"

func memchr(src unsafe.Pointer, c int32, n uint32) unsafe.Pointer {
	var s *uint8 = (*uint8)(src)
	c = int32(uint8(c))
	for ; uint64(uint32(uintptr(unsafe.Pointer(s))))&(4-uint64(1)) != 0 && n != 0 && int32(*s) != c; func() uint32 {
		*(*uintptr)(unsafe.Pointer(&s))++
		return func() (_cgo_ret uint32) {
			_cgo_addr := &n
			_cgo_ret = *_cgo_addr
			*_cgo_addr--
			return
		}()
	}() {
	}
	if n != 0 && int32(*s) != c {
		type word = uint32
		var w *uint32
		var k uint32 = uint32(4294967295) / uint32(255) * uint32(c)
		for w = (*uint32)(unsafe.Pointer(s)); uint64(n) >= 4 && !((*w^k-uint32(4294967295)/uint32(255)) & ^(*w^k) & (uint32(4294967295)/uint32(255)*uint32(255/2+1)) != 0); func() uint32 {
			*(*uintptr)(unsafe.Pointer(&w)) += 4
			return func() (_cgo_ret uint32) {
				_cgo_addr := &n
				*_cgo_addr -= uint32(4)
				return *_cgo_addr
			}()
		}() {
		}
		s = (*uint8)(unsafe.Pointer(w))
	}
	for ; n != 0 && int32(*s) != c; func() uint32 {
		*(*uintptr)(unsafe.Pointer(&s))++
		return func() (_cgo_ret uint32) {
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
