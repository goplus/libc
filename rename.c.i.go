package libc

import unsafe "unsafe"

func Rename(old *int8, new *int8) int32 {
	return int32(__syscall2_r1(int64(128), int64(uintptr(unsafe.Pointer(old))), int64(uintptr(unsafe.Pointer(new)))))
}
