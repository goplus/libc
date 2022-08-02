package libc

import unsafe "unsafe"

func Read(fd int32, buf unsafe.Pointer, count uint64) int64 {
	return __syscall_ret(uint64(__syscall_cp(int64(3), int64(fd), int64(uintptr(buf)), int64(count), int64(0), int64(0), int64(0))))
}
