package libc

import unsafe "unsafe"

func jrand48(s *uint16) int64 {
	return int64(int32(__rand48_step(s, (*uint16)(unsafe.Pointer(uintptr(unsafe.Pointer((*uint16)(unsafe.Pointer(&__seed48))))+uintptr(int32(3))*2))) >> int32(16)))
}
func mrand48() int64 {
	return jrand48((*uint16)(unsafe.Pointer(&__seed48)))
}
