package libc

import unsafe "unsafe"

func __fopen_rb_ca(filename *int8, f *Struct__IO_FILE, buf *uint8, len uint64) *Struct__IO_FILE {
	Memset(unsafe.Pointer(f), int32(0), 232)
	f.Fd = int32(__syscall2_r1(int64(5), int64(uintptr(unsafe.Pointer(filename))), int64(655360)))
	if f.Fd < int32(0) {
		return (*Struct__IO_FILE)(nil)
	}
	__syscall3(int64(92), int64(f.Fd), int64(2), int64(1))
	f.Flags = uint32(9)
	f.Buf = (*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(buf)) + uintptr(int32(8))))
	f.Buf_size = len - uint64(8)
	f.Read = __stdio_read
	f.Seek = __stdio_seek
	f.Close = __stdio_close
	f.Lock = -1
	return f
}
