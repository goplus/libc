package libc

import unsafe "unsafe"

func Pipe(fd *int32) int32 {
	return int32(__syscall1_r1(int64(42), int64(uintptr(unsafe.Pointer(fd)))))
}
