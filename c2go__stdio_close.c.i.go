package libc

func dummy(fd int32) int32 {
	return fd
}
func __stdio_close(f *struct__IO_FILE) int32 {
	return int32(__syscall_ret(uint64(__syscall1(int64(6), int64(__aio_close(f.fd))))))
}
