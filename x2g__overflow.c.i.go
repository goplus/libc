package libc

import unsafe "unsafe"

func __overflow(f *struct__IO_FILE, _c int32) int32 {
	var c uint8 = uint8(_c)
	if !(f.wend != nil) && __towrite(f) != 0 {
		return -1
	}
	if uintptr(unsafe.Pointer(f.wpos)) != uintptr(unsafe.Pointer(f.wend)) && int32(c) != f.lbf {
		return int32(func() (_cgo_ret uint8) {
			_cgo_addr := &*func() (_cgo_ret *uint8) {
				_cgo_addr := &f.wpos
				_cgo_ret = *_cgo_addr
				*(*uintptr)(unsafe.Pointer(_cgo_addr))++
				return
			}()
			*_cgo_addr = c
			return *_cgo_addr
		}())
	}
	if f.write(f, &c, uint64(1)) != uint64(1) {
		return -1
	}
	return int32(c)
}
