package libc

import unsafe "unsafe"

func nrand48(s *uint16) int64 {
	return int64(__rand48_step(s, (*uint16)(unsafe.Pointer(uintptr(unsafe.Pointer((*uint16)(unsafe.Pointer(&__seed48))))+uintptr(int32(3))*2))) >> int32(17))
}
func lrand48() int64 {
	return nrand48((*uint16)(unsafe.Pointer(&__seed48)))
}
