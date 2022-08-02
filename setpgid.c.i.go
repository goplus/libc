package libc

func setpgid(pid int32, pgid int32) int32 {
	return int32(__syscall2_r1(int64(82), int64(pid), int64(pgid)))
}
