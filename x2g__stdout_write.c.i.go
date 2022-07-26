package libc

import unsafe "unsafe"

func __stdout_write(f *Struct__IO_FILE, buf *uint8, len uint64) uint64 {
	var wsz Struct_winsize
	f.Write = __stdio_write
	if !(f.Flags&uint32(64) != 0) && __syscall3(int64(54), int64(f.Fd), int64(21523), int64(uintptr(unsafe.Pointer(&wsz)))) != 0 {
		f.Lbf = -1
	}
	return __stdio_write(f, buf, len)
}
