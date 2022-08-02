package libc

func getegid() uint32 {
	return uint32(__syscall0(int64(43)))
}
