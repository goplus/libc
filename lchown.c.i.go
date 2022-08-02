package libc

import unsafe "unsafe"

func lchown(path *int8, uid uint32, gid uint32) int32 {
	return int32(__syscall3_r1(int64(364), int64(uintptr(unsafe.Pointer(path))), int64(uid), int64(gid)))
}
