package libc

func setreuid(ruid uint32, euid uint32) int32 {
	return __setxid(int32(126), int32(ruid), int32(euid), int32(0))
}
