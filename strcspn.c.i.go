package libc

import unsafe "unsafe"

func Strcspn(s *int8, c *int8) uint64 {
	var a *int8 = s
	var byteset [4]uint64
	if !(*(*int8)(unsafe.Pointer(uintptr(unsafe.Pointer(c)) + uintptr(int32(0)))) != 0) || !(*(*int8)(unsafe.Pointer(uintptr(unsafe.Pointer(c)) + uintptr(int32(1)))) != 0) {
		return uint64(uintptr(unsafe.Pointer(__strchrnul(s, int32(*c)))) - uintptr(unsafe.Pointer(a)))
	}
	Memset(unsafe.Pointer((*uint64)(unsafe.Pointer(&byteset))), int32(0), 32)
	for ; int32(*c) != 0 && func() (_cgo_ret uint64) {
		_cgo_addr := &*(*uint64)(unsafe.Pointer(uintptr(unsafe.Pointer((*uint64)(unsafe.Pointer(&byteset)))) + uintptr(uint64(*(*uint8)(unsafe.Pointer(c)))/64)*8))
		*_cgo_addr |= uint64(1) << (uint64(*(*uint8)(unsafe.Pointer(c))) % 64)
		return *_cgo_addr
	}() != 0; *(*uintptr)(unsafe.Pointer(&c))++ {
	}
	for ; int32(*s) != 0 && !(*(*uint64)(unsafe.Pointer(uintptr(unsafe.Pointer((*uint64)(unsafe.Pointer(&byteset)))) + uintptr(uint64(*(*uint8)(unsafe.Pointer(s)))/64)*8))&(uint64(1)<<(uint64(*(*uint8)(unsafe.Pointer(s)))%64)) != 0); *(*uintptr)(unsafe.Pointer(&s))++ {
	}
	return uint64(uintptr(unsafe.Pointer(s)) - uintptr(unsafe.Pointer(a)))
}
