package libc

import unsafe "unsafe"

func __getauxval(item uint64) uint64 {
	var auxv *uint64 = __libc.auxv
	if item == uint64(23) {
		return uint64(__libc.secure)
	}
	for ; *auxv != 0; *(*uintptr)(unsafe.Pointer(&auxv)) += uintptr(int32(2)) * 8 {
		if *auxv == item {
			return *(*uint64)(unsafe.Pointer(uintptr(unsafe.Pointer(auxv)) + uintptr(int32(1))*8))
		}
	}
	*__errno_location() = int32(2)
	return uint64(0)
}
