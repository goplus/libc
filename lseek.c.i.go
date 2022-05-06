package libc

func __lseek(fd int32, offset int64, whence int32) int64 {
	return int64(__syscall_ret(uint64(__syscall3(int64(199), int64(fd), int64(offset), int64(whence)))))
}
