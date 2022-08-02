package libc

func getppid() int32 {
	return int32(__syscall0(int64(39)))
}
