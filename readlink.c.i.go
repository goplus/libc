package libc

import unsafe "unsafe"

func Readlink(path *int8, buf *int8, bufsize uint64) int64 {
	var dummy [1]int8
	if !(bufsize != 0) {
		buf = (*int8)(unsafe.Pointer(&dummy))
		bufsize = uint64(1)
	}
	var r int32 = int32(__syscall3(int64(58), int64(uintptr(unsafe.Pointer(path))), int64(uintptr(unsafe.Pointer(buf))), int64(bufsize)))
	if uintptr(unsafe.Pointer(buf)) == uintptr(unsafe.Pointer((*int8)(unsafe.Pointer(&dummy)))) && r > int32(0) {
		r = int32(0)
	}
	return __syscall_ret(uint64(r))
}
