package libc

import unsafe "unsafe"

func fchdir(fd int32) int32 {
	var ret int32 = int32(__syscall1(int64(13), int64(fd)))
	if ret != -9 || __syscall2(int64(92), int64(fd), int64(1)) < int64(0) {
		return int32(__syscall_ret(uint64(ret)))
	}
	var buf [27]int8
	__procfdname((*int8)(unsafe.Pointer(&buf)), uint32(fd))
	return int32(__syscall1_r1(int64(12), int64(uintptr(unsafe.Pointer((*int8)(unsafe.Pointer(&buf)))))))
}
