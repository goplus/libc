package libc

func X__lseek(fd int32, offset int64, whence int32) int64 {
	return int64(__syscall3_r1(int64(199), int64(fd), int64(offset), int64(whence)))
}
