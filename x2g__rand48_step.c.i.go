package libc

import unsafe "unsafe"

func __rand48_step(xi *uint16, lc *uint16) uint64 {
	var a uint64
	var x uint64
	x = uint64(uint32(*(*uint16)(unsafe.Pointer(uintptr(unsafe.Pointer(xi)) + uintptr(int32(0))*2)))|(uint32(*(*uint16)(unsafe.Pointer(uintptr(unsafe.Pointer(xi)) + uintptr(int32(1))*2)))+uint32(0))<<int32(16)) | (uint64(*(*uint16)(unsafe.Pointer(uintptr(unsafe.Pointer(xi)) + uintptr(int32(2))*2)))+uint64(0))<<int32(32)
	a = uint64(uint32(*(*uint16)(unsafe.Pointer(uintptr(unsafe.Pointer(lc)) + uintptr(int32(0))*2)))|(uint32(*(*uint16)(unsafe.Pointer(uintptr(unsafe.Pointer(lc)) + uintptr(int32(1))*2)))+uint32(0))<<int32(16)) | (uint64(*(*uint16)(unsafe.Pointer(uintptr(unsafe.Pointer(lc)) + uintptr(int32(2))*2)))+uint64(0))<<int32(32)
	x = a*x + uint64(*(*uint16)(unsafe.Pointer(uintptr(unsafe.Pointer(lc)) + uintptr(int32(3))*2)))
	*(*uint16)(unsafe.Pointer(uintptr(unsafe.Pointer(xi)) + uintptr(int32(0))*2)) = uint16(x)
	*(*uint16)(unsafe.Pointer(uintptr(unsafe.Pointer(xi)) + uintptr(int32(1))*2)) = uint16(x >> int32(16))
	*(*uint16)(unsafe.Pointer(uintptr(unsafe.Pointer(xi)) + uintptr(int32(2))*2)) = uint16(x >> int32(32))
	return x & uint64(281474976710655)
}
