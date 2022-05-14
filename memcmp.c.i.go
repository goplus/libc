package libc

import unsafe "unsafe"

func Memcmp(vl unsafe.Pointer, vr unsafe.Pointer, n uint64) int32 {
	var l *uint8 = (*uint8)(vl)
	var r *uint8 = (*uint8)(vr)
	for ; n != 0 && int32(*l) == int32(*r); func() *uint8 {
		func() *uint8 {
			n--
			return func() (_cgo_ret *uint8) {
				_cgo_addr := &l
				_cgo_ret = *_cgo_addr
				*(*uintptr)(unsafe.Pointer(_cgo_addr))++
				return
			}()
		}()
		return func() (_cgo_ret *uint8) {
			_cgo_addr := &r
			_cgo_ret = *_cgo_addr
			*(*uintptr)(unsafe.Pointer(_cgo_addr))++
			return
		}()
	}() {
	}
	return func() int32 {
		if n != 0 {
			return int32(*l) - int32(*r)
		} else {
			return int32(0)
		}
	}()
}
