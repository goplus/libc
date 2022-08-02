package libc

func geteuid() uint32 {
	return uint32(__syscall0(int64(25)))
}
