package libc

import unsafe "unsafe"

func Posix_openpt(flags int32) int32 {
	var r int32 = open((*int8)(unsafe.Pointer(&[10]int8{'/', 'd', 'e', 'v', '/', 'p', 't', 'm', 'x', '\x00'})), flags)
	if r < int32(0) && *__errno_location() == int32(28) {
		*__errno_location() = int32(11)
	}
	return r
}
func Grantpt(fd int32) int32 {
	return int32(0)
}
func Unlockpt(fd int32) int32 {
	var unlock int32 = int32(0)
	return Ioctl(fd, int32(1074025521), &unlock)
}
func __ptsname_r(fd int32, buf *int8, len uint64) int32 {
	var pty int32
	var err int32
	if !(buf != nil) {
		len = uint64(0)
	}
	if func() (_cgo_ret int32) {
		_cgo_addr := &err
		*_cgo_addr = int32(__syscall3(int64(54), int64(fd), int64(2147767344), int64(uintptr(unsafe.Pointer(&pty)))))
		return *_cgo_addr
	}() != 0 {
		return -err
	}
	if uint64(Snprintf(buf, len, (*int8)(unsafe.Pointer(&[12]int8{'/', 'd', 'e', 'v', '/', 'p', 't', 's', '/', '%', 'd', '\x00'})), pty)) >= len {
		return int32(34)
	}
	return int32(0)
}
