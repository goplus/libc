package libc

import unsafe "unsafe"

func Lcong48(p *uint16) {
	Memcpy(unsafe.Pointer((*uint16)(unsafe.Pointer(&__seed48))), unsafe.Pointer(p), 14)
}
