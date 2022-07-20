package libc

import unsafe "unsafe"

func srand48(seed int64) {
	seed48((*uint16)(unsafe.Pointer(&[3]uint16{uint16(13070), uint16(seed), uint16(seed >> int32(16))})))
}
