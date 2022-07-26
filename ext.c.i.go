package libc

import unsafe "unsafe"

func _flushlbf() {
	Fflush(nil)
}
func __fsetlocking(f *Struct__IO_FILE, type_ int32) int32 {
	return int32(0)
}
func __fwriting(f *Struct__IO_FILE) int32 {
	return func() int32 {
		if f.Flags&uint32(4) != 0 || f.Wend != nil {
			return 1
		} else {
			return 0
		}
	}()
}
func __freading(f *Struct__IO_FILE) int32 {
	return func() int32 {
		if f.Flags&uint32(8) != 0 || f.Rend != nil {
			return 1
		} else {
			return 0
		}
	}()
}
func __freadable(f *Struct__IO_FILE) int32 {
	return func() int32 {
		if !(f.Flags&uint32(4) != 0) {
			return 1
		} else {
			return 0
		}
	}()
}
func __fwritable(f *Struct__IO_FILE) int32 {
	return func() int32 {
		if !(f.Flags&uint32(8) != 0) {
			return 1
		} else {
			return 0
		}
	}()
}
func __flbf(f *Struct__IO_FILE) int32 {
	return func() int32 {
		if f.Lbf >= int32(0) {
			return 1
		} else {
			return 0
		}
	}()
}
func __fbufsize(f *Struct__IO_FILE) uint64 {
	return f.Buf_size
}
func __fpending(f *Struct__IO_FILE) uint64 {
	return uint64(func() int64 {
		if f.Wend != nil {
			return int64(uintptr(unsafe.Pointer(f.Wpos)) - uintptr(unsafe.Pointer(f.Wbase)))
		} else {
			return int64(0)
		}
	}())
}
func __fpurge(f *Struct__IO_FILE) int32 {
	f.Wpos = func() (_cgo_ret *uint8) {
		_cgo_addr := &f.Wbase
		*_cgo_addr = func() (_cgo_ret *uint8) {
			_cgo_addr := &f.Wend
			*_cgo_addr = (*uint8)(nil)
			return *_cgo_addr
		}()
		return *_cgo_addr
	}()
	f.Rpos = func() (_cgo_ret *uint8) {
		_cgo_addr := &f.Rend
		*_cgo_addr = (*uint8)(nil)
		return *_cgo_addr
	}()
	return int32(0)
}
