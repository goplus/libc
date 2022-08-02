package libc

func getgid() uint32 {
	return uint32(__syscall0(int64(47)))
}
