package libc

import unsafe "unsafe"

func getgroups(count int32, list *uint32) int32 {
	return int32(__syscall2_r1(int64(79), int64(count), int64(uintptr(unsafe.Pointer(list)))))
}
