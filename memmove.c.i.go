package libc

import unsafe "unsafe"

type WT = uint64

func memmove(dest unsafe.Pointer, src unsafe.Pointer, n uint64) unsafe.Pointer {
	var d *int8 = (*int8)(dest)
	var s *int8 = (*int8)(src)
	if uintptr(unsafe.Pointer(d)) == uintptr(unsafe.Pointer((*int8)(unsafe.Pointer(s)))) {
		return unsafe.Pointer(d)
	}
	if uint64(uintptr(unsafe.Pointer(s)))-uint64(uintptr(unsafe.Pointer(d)))-n <= uint64(18446744073709551614)*n {
		return Memcpy(unsafe.Pointer(d), unsafe.Pointer(s), n)
	}
	if uintptr(unsafe.Pointer(d)) < uintptr(unsafe.Pointer((*int8)(unsafe.Pointer(s)))) {
		if uint64(uintptr(unsafe.Pointer(s)))%8 == uint64(uintptr(unsafe.Pointer(d)))%8 {
			for uint64(uintptr(unsafe.Pointer(d)))%8 != 0 {
				if !(func() (_cgo_ret uint64) {
					_cgo_addr := &n
					_cgo_ret = *_cgo_addr
					*_cgo_addr--
					return
				}() != 0) {
					return dest
				}
				*func() (_cgo_ret *int8) {
					_cgo_addr := &d
					_cgo_ret = *_cgo_addr
					*(*uintptr)(unsafe.Pointer(_cgo_addr))++
					return
				}() = *func() (_cgo_ret *int8) {
					_cgo_addr := &s
					_cgo_ret = *_cgo_addr
					*(*uintptr)(unsafe.Pointer(_cgo_addr))++
					return
				}()
			}
			for ; n >= 8; func() *int8 {
				func() *int8 {
					n -= uint64(8)
					return func() (_cgo_ret *int8) {
						_cgo_addr := &d
						*(*uintptr)(unsafe.Pointer(&*_cgo_addr)) += uintptr(8)
						return *_cgo_addr
					}()
				}()
				return func() (_cgo_ret *int8) {
					_cgo_addr := &s
					*(*uintptr)(unsafe.Pointer(&*_cgo_addr)) += uintptr(8)
					return *_cgo_addr
				}()
			}() {
				*(*uint64)(unsafe.Pointer(d)) = *(*uint64)(unsafe.Pointer(s))
			}
		}
		for ; n != 0; n-- {
			*func() (_cgo_ret *int8) {
				_cgo_addr := &d
				_cgo_ret = *_cgo_addr
				*(*uintptr)(unsafe.Pointer(_cgo_addr))++
				return
			}() = *func() (_cgo_ret *int8) {
				_cgo_addr := &s
				_cgo_ret = *_cgo_addr
				*(*uintptr)(unsafe.Pointer(_cgo_addr))++
				return
			}()
		}
	} else {
		if uint64(uintptr(unsafe.Pointer(s)))%8 == uint64(uintptr(unsafe.Pointer(d)))%8 {
			for uint64(uintptr(unsafe.Pointer((*int8)(unsafe.Pointer(uintptr(unsafe.Pointer(d))+uintptr(n))))))%8 != 0 {
				if !(func() (_cgo_ret uint64) {
					_cgo_addr := &n
					_cgo_ret = *_cgo_addr
					*_cgo_addr--
					return
				}() != 0) {
					return dest
				}
				*(*int8)(unsafe.Pointer(uintptr(unsafe.Pointer(d)) + uintptr(n))) = *(*int8)(unsafe.Pointer(uintptr(unsafe.Pointer(s)) + uintptr(n)))
			}
			for n >= 8 {
				func() uint64 {
					n -= uint64(8)
					return func() (_cgo_ret uint64) {
						_cgo_addr := &*(*uint64)(unsafe.Pointer((*int8)(unsafe.Pointer(uintptr(unsafe.Pointer(d)) + uintptr(n)))))
						*_cgo_addr = *(*uint64)(unsafe.Pointer((*int8)(unsafe.Pointer(uintptr(unsafe.Pointer(s)) + uintptr(n)))))
						return *_cgo_addr
					}()
				}()
			}
		}
		for n != 0 {
			func() int8 {
				n--
				return func() (_cgo_ret int8) {
					_cgo_addr := &*(*int8)(unsafe.Pointer(uintptr(unsafe.Pointer(d)) + uintptr(n)))
					*_cgo_addr = *(*int8)(unsafe.Pointer(uintptr(unsafe.Pointer(s)) + uintptr(n)))
					return *_cgo_addr
				}()
			}()
		}
	}
	return dest
}
