package libc

import unsafe "unsafe"

func Erand48(s *uint16) float64 {
	type _cgoa_19_drand48 struct {
		u uint64
	}
	var x _cgoa_19_drand48
	x.u = uint64(4607182418800017408) | __rand48_step(s, (*uint16)(unsafe.Pointer(uintptr(unsafe.Pointer((*uint16)(unsafe.Pointer(&__seed48))))+uintptr(int32(3))*2)))<<int32(4)
	return *(*float64)(unsafe.Pointer(&x)) - 1
}
func Drand48() float64 {
	return Erand48((*uint16)(unsafe.Pointer(&__seed48)))
}
