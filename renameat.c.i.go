package libc

import unsafe "unsafe"

func Renameat(oldfd int32, old *int8, newfd int32, new *int8) int32 {
	return int32(__syscall4_r1(int64(465), int64(oldfd), int64(uintptr(unsafe.Pointer(old))), int64(newfd), int64(uintptr(unsafe.Pointer(new)))))
}
