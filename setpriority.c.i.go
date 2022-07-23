package libc

func setpriority(which int32, who uint32, prio int32) int32 {
	return int32(__syscall3_r1(int64(96), int64(which), int64(who), int64(prio)))
}
