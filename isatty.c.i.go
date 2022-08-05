package libc

import unsafe "unsafe"

func isatty(fd int32) int32 {
	var wsz Struct_winsize
	var r uint64 = uint64(__syscall3_r1(int64(54), int64(fd), int64(21523), int64(uintptr(unsafe.Pointer(&wsz)))))
	if r == uint64(0) {
		return int32(1)
	}
	if *X__errno_location() != int32(9) {
		*X__errno_location() = int32(25)
	}
	return int32(0)
}
