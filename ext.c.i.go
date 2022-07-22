package libc

import unsafe "unsafe"

func _flushlbf() {
	Fflush(nil)
}
func __fsetlocking(f *struct__IO_FILE, type_ int32) int32 {
	return int32(0)
}
func __fwriting(f *struct__IO_FILE) int32 {
	return func() int32 {
		if f.flags&uint32(4) != 0 || f.wend != nil {
			return 1
		} else {
			return 0
		}
	}()
}
func __freading(f *struct__IO_FILE) int32 {
	return func() int32 {
		if f.flags&uint32(8) != 0 || f.rend != nil {
			return 1
		} else {
			return 0
		}
	}()
}
func __freadable(f *struct__IO_FILE) int32 {
	return func() int32 {
		if !(f.flags&uint32(4) != 0) {
			return 1
		} else {
			return 0
		}
	}()
}
func __fwritable(f *struct__IO_FILE) int32 {
	return func() int32 {
		if !(f.flags&uint32(8) != 0) {
			return 1
		} else {
			return 0
		}
	}()
}
func __flbf(f *struct__IO_FILE) int32 {
	return func() int32 {
		if f.lbf >= int32(0) {
			return 1
		} else {
			return 0
		}
	}()
}
func __fbufsize(f *struct__IO_FILE) uint64 {
	return f.buf_size
}
func __fpending(f *struct__IO_FILE) uint64 {
	return uint64(func() int64 {
		if f.wend != nil {
			return int64(uintptr(unsafe.Pointer(f.wpos)) - uintptr(unsafe.Pointer(f.wbase)))
		} else {
			return int64(0)
		}
	}())
}
func __fpurge(f *struct__IO_FILE) int32 {
	f.wpos = func() (_cgo_ret *uint8) {
		_cgo_addr := &f.wbase
		*_cgo_addr = func() (_cgo_ret *uint8) {
			_cgo_addr := &f.wend
			*_cgo_addr = (*uint8)(nil)
			return *_cgo_addr
		}()
		return *_cgo_addr
	}()
	f.rpos = func() (_cgo_ret *uint8) {
		_cgo_addr := &f.rend
		*_cgo_addr = (*uint8)(nil)
		return *_cgo_addr
	}()
	return int32(0)
}
