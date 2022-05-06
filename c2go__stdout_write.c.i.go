package libc

import unsafe "unsafe"

func __stdout_write(f *struct__IO_FILE, buf *uint8, len uint64) uint64 {
	var wsz struct_winsize
	f.write = __stdio_write
	if !(f.flags&uint32(64) != 0) && __syscall3(int64(54), int64(f.fd), int64(21523), int64(uintptr(unsafe.Pointer(&wsz)))) != 0 {
		f.lbf = int32(-1)
	}
	return __stdio_write(f, buf, len)
}
