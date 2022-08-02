package libc

func getpgid(pid int32) int32 {
	return int32(__syscall1_r1(int64(151), int64(pid)))
}
