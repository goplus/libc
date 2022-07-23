package libc

import unsafe "unsafe"

func initgroups(user *int8, gid uint32) int32 {
	var groups [32]uint32
	var count int32 = int32(32)
	if getgrouplist(user, gid, (*uint32)(unsafe.Pointer(&groups)), &count) < int32(0) {
		return -1
	}
	return setgroups(uint64(count), (*uint32)(unsafe.Pointer(&groups)))
}
