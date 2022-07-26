package libc

import unsafe "unsafe"

func __fdopen(fd int32, mode *int8) *Struct__IO_FILE {
	var f *Struct__IO_FILE
	var wsz Struct_winsize
	if !(Strchr((*int8)(unsafe.Pointer(&[4]int8{'r', 'w', 'a', '\x00'})), int32(*mode)) != nil) {
		*__errno_location() = int32(22)
		return (*Struct__IO_FILE)(nil)
	}
	if !(func() (_cgo_ret *Struct__IO_FILE) {
		_cgo_addr := &f
		*_cgo_addr = (*Struct__IO_FILE)(Malloc(1264))
		return *_cgo_addr
	}() != nil) {
		return (*Struct__IO_FILE)(nil)
	}
	Memset(unsafe.Pointer(f), int32(0), 232)
	if !(Strchr(mode, '+') != nil) {
		f.Flags = uint32(func() int32 {
			if int32(*mode) == 'r' {
				return int32(8)
			} else {
				return int32(4)
			}
		}())
	}
	if Strchr(mode, 'e') != nil {
		__syscall3(int64(92), int64(fd), int64(2), int64(1))
	}
	if int32(*mode) == 'a' {
		var flags int32 = int32(__syscall2(int64(92), int64(fd), int64(3)))
		if !(flags&int32(1024) != 0) {
			__syscall3(int64(92), int64(fd), int64(4), int64(flags|int32(1024)))
		}
		f.Flags |= uint32(128)
	}
	f.Fd = fd
	f.Buf = (*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer((*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer((*uint8)(unsafe.Pointer(f))))+uintptr(232))))) + uintptr(int32(8))))
	f.Buf_size = uint64(1024)
	f.Lbf = -1
	if !(f.Flags&uint32(8) != 0) && !(__syscall3(int64(54), int64(fd), int64(21523), int64(uintptr(unsafe.Pointer(&wsz)))) != 0) {
		f.Lbf = int32('\n')
	}
	f.Read = __stdio_read
	f.Write = __stdio_write
	f.Seek = __stdio_seek
	f.Close = __stdio_close
	if !(__libc.threaded != 0) {
		f.Lock = -1
	}
	return __ofl_add(f)
}
