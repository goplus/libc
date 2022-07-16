package libc

import unsafe "unsafe"

func wcsncasecmp(l *uint32, r *uint32, n uint64) int32 {
	if !(func() (_cgo_ret uint64) {
		_cgo_addr := &n
		_cgo_ret = *_cgo_addr
		*_cgo_addr--
		return
	}() != 0) {
		return int32(0)
	}
	for ; *l != 0 && *r != 0 && n != 0 && (*l == *r || towlower(*l) == towlower(*r)); func() uint64 {
		func() *uint32 {
			*(*uintptr)(unsafe.Pointer(&l)) += 4
			return func() (_cgo_ret *uint32) {
				_cgo_addr := &r
				_cgo_ret = *_cgo_addr
				*(*uintptr)(unsafe.Pointer(_cgo_addr)) += 4
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
	return int32(towlower(*l) - towlower(*r))
}
