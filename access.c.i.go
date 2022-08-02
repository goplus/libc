package libc

import unsafe "unsafe"

func access(filename *int8, amode int32) int32 {
	return int32(__syscall2_r1(int64(33), int64(uintptr(unsafe.Pointer(filename))), int64(amode)))
}
