package libc

import unsafe "unsafe"

func Seed48(s *uint16) *uint16 {
	Memcpy(unsafe.Pointer((*uint16)(unsafe.Pointer(&_cgos_p_seed48))), unsafe.Pointer((*uint16)(unsafe.Pointer(&__seed48))), 6)
	Memcpy(unsafe.Pointer((*uint16)(unsafe.Pointer(&__seed48))), unsafe.Pointer(s), 6)
	return (*uint16)(unsafe.Pointer(&_cgos_p_seed48))
}

var _cgos_p_seed48 [3]uint16
