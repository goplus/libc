package libc

import unsafe "unsafe"

func Strncmp(_l *int8, _r *int8, n uint64) int32 {
	var l *uint8 = (*uint8)(unsafe.Pointer(_l))
	var r *uint8 = (*uint8)(unsafe.Pointer(_r))
	if !(func() (_cgo_ret uint64) {
		_cgo_addr := &n
		_cgo_ret = *_cgo_addr
		*_cgo_addr--
		return
	}() != 0) {
		return int32(0)
	}
	for ; int32(*l) != 0 && int32(*r) != 0 && n != 0 && int32(*l) == int32(*r); func() uint64 {
		func() *uint8 {
			*(*uintptr)(unsafe.Pointer(&l))++
			return func() (_cgo_ret *uint8) {
				_cgo_addr := &r
				_cgo_ret = *_cgo_addr
				*(*uintptr)(unsafe.Pointer(_cgo_addr))++
				return
			}()
		}()
		return func() (_cgo_ret uint64) {
			_cgo_addr := &n
			_cgo_ret = *_cgo_addr
			*_cgo_addr--
			return
		}()
	}() {
	}
	return int32(*l) - int32(*r)
}
