package libc

func getuid() uint32 {
	return uint32(__syscall0(int64(24)))
}
