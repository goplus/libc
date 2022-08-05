package libc

import unsafe "unsafe"

func Fopen(filename *int8, mode *int8) *Struct__IO_FILE {
	var f *Struct__IO_FILE
	var fd int32
	var flags int32
	if !(Strchr((*int8)(unsafe.Pointer(&[4]int8{'r', 'w', 'a', '\x00'})), int32(*mode)) != nil) {
		*X__errno_location() = int32(22)
		return (*Struct__IO_FILE)(nil)
	}
	flags = __fmodeflags(mode)
	fd = int32(__syscall3_r1(int64(5), int64(uintptr(unsafe.Pointer(filename))), int64(flags|int32(131072)), int64(438)))
	if fd < int32(0) {
		return (*Struct__IO_FILE)(nil)
	}
	if flags&int32(524288) != 0 {
		__syscall3(int64(92), int64(fd), int64(2), int64(1))
	}
	f = __fdopen(fd, mode)
	if f != nil {
		return f
	}
	__syscall1(int64(6), int64(fd))
	return (*Struct__IO_FILE)(nil)
}
