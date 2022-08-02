package libc

import unsafe "unsafe"

func fchownat(fd int32, path *int8, uid uint32, gid uint32, flag int32) int32 {
	return int32(__syscall5_r1(int64(468), int64(fd), int64(uintptr(unsafe.Pointer(path))), int64(uid), int64(gid), int64(flag)))
}
