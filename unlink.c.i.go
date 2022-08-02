package libc

import unsafe "unsafe"

func unlink(path *int8) int32 {
	return int32(__syscall1_r1(int64(10), int64(uintptr(unsafe.Pointer(path)))))
}
