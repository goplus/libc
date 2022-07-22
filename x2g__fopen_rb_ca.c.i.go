package libc

import unsafe "unsafe"

func __fopen_rb_ca(filename *int8, f *struct__IO_FILE, buf *uint8, len uint64) *struct__IO_FILE {
	Memset(unsafe.Pointer(f), int32(0), 232)
	f.fd = int32(__syscall2_r1(int64(5), int64(uintptr(unsafe.Pointer(filename))), int64(655360)))
	if f.fd < int32(0) {
		return (*struct__IO_FILE)(nil)
	}
	__syscall3(int64(92), int64(f.fd), int64(2), int64(1))
	f.flags = uint32(9)
	f.buf = (*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(buf)) + uintptr(int32(8))))
	f.buf_size = len - uint64(8)
	f.read = __stdio_read
	f.seek = __stdio_seek
	f.close = __stdio_close
	f.lock = -1
	return f
}
