package libc

func fsync(fd int32) int32 {
	return int32(__syscall_ret(uint64(__syscall_cp(int64(95), int64(fd), int64(0), int64(0), int64(0), int64(0), int64(0)))))
}
