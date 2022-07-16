package libc

import unsafe "unsafe"

func wcsncpy(d *uint32, s *uint32, n uint64) *uint32 {
	var a *uint32 = d
	for n != 0 && *s != 0 {
		func() uint32 {
			n--
			return func() (_cgo_ret uint32) {
				_cgo_addr := &*func() (_cgo_ret *uint32) {
					_cgo_addr := &d
					_cgo_ret = *_cgo_addr
					*(*uintptr)(unsafe.Pointer(_cgo_addr)) += 4
					return
				}()
				*_cgo_addr = *func() (_cgo_ret *uint32) {
					_cgo_addr := &s
					_cgo_ret = *_cgo_addr
					*(*uintptr)(unsafe.Pointer(_cgo_addr)) += 4
					return
				}()
				return *_cgo_addr
			}()
		}()
	}
	wmemset(d, uint32(0), n)
	return a
}
