package libc

import unsafe "unsafe"

func fchown(fd int32, uid uint32, gid uint32) int32 {
	var ret int32 = int32(__syscall3(int64(123), int64(fd), int64(uid), int64(gid)))
	if ret != -9 || __syscall2(int64(92), int64(fd), int64(1)) < int64(0) {
		return int32(__syscall_ret(uint64(ret)))
	}
	var buf [27]int8
	__procfdname((*int8)(unsafe.Pointer(&buf)), uint32(fd))
	return int32(__syscall3_r1(int64(16), int64(uintptr(unsafe.Pointer((*int8)(unsafe.Pointer(&buf))))), int64(uid), int64(gid)))
}
