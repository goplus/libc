package libc

import unsafe "unsafe"

func memset(dest unsafe.Pointer, c int32, n uint32) unsafe.Pointer {
	var s *uint8 = (*uint8)(dest)
	var k uint32
	if !(n != 0) {
		return dest
	}
	*(*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(s)) + uintptr(0))) = uint8(c)
	*(*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(s)) + uintptr(n-uint32(1)))) = uint8(c)
	if n <= uint32(2) {
		return dest
	}
	*(*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(s)) + uintptr(1))) = uint8(c)
	*(*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(s)) + uintptr(2))) = uint8(c)
	*(*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(s)) + uintptr(n-uint32(2)))) = uint8(c)
	*(*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(s)) + uintptr(n-uint32(3)))) = uint8(c)
	if n <= uint32(6) {
		return dest
	}
	*(*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(s)) + uintptr(3))) = uint8(c)
	*(*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(s)) + uintptr(n-uint32(4)))) = uint8(c)
	if n <= uint32(8) {
		return dest
	}
	k = -uint32(uintptr(unsafe.Pointer(s))) & uint32(3)
	*(*uintptr)(unsafe.Pointer(&s)) += uintptr(k)
	n -= k
	n &= uint32(4294967292)
	type u32 = uint32
	type u64 = uint64
	var c32 uint32 = uint32(4294967295) / uint32(255) * uint32(uint8(c))
	*(*uint32)(unsafe.Pointer((*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(s)) + uintptr(0))))) = c32
	*(*uint32)(unsafe.Pointer((*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer((*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(s))+uintptr(n))))) - uintptr(4))))) = c32
	if n <= uint32(8) {
		return dest
	}
	*(*uint32)(unsafe.Pointer((*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(s)) + uintptr(4))))) = c32
	*(*uint32)(unsafe.Pointer((*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(s)) + uintptr(8))))) = c32
	*(*uint32)(unsafe.Pointer((*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer((*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(s))+uintptr(n))))) - uintptr(12))))) = c32
	*(*uint32)(unsafe.Pointer((*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer((*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(s))+uintptr(n))))) - uintptr(8))))) = c32
	if n <= uint32(24) {
		return dest
	}
	*(*uint32)(unsafe.Pointer((*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(s)) + uintptr(12))))) = c32
	*(*uint32)(unsafe.Pointer((*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(s)) + uintptr(16))))) = c32
	*(*uint32)(unsafe.Pointer((*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(s)) + uintptr(20))))) = c32
	*(*uint32)(unsafe.Pointer((*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(s)) + uintptr(24))))) = c32
	*(*uint32)(unsafe.Pointer((*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer((*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(s))+uintptr(n))))) - uintptr(28))))) = c32
	*(*uint32)(unsafe.Pointer((*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer((*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(s))+uintptr(n))))) - uintptr(24))))) = c32
	*(*uint32)(unsafe.Pointer((*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer((*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(s))+uintptr(n))))) - uintptr(20))))) = c32
	*(*uint32)(unsafe.Pointer((*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer((*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(s))+uintptr(n))))) - uintptr(16))))) = c32
	k = uint32(24) + uint32(uintptr(unsafe.Pointer(s)))&uint32(4)
	*(*uintptr)(unsafe.Pointer(&s)) += uintptr(k)
	n -= k
	var c64 uint64 = uint64(c32) | uint64(c32)<<32
	for ; n >= uint32(32); func() *uint8 {
		n -= uint32(32)
		return func() (_cgo_ret *uint8) {
			_cgo_addr := &s
			*(*uintptr)(unsafe.Pointer(&*_cgo_addr)) += uintptr(32)
			return *_cgo_addr
		}()
	}() {
		*(*uint64)(unsafe.Pointer((*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(s)) + uintptr(0))))) = c64
		*(*uint64)(unsafe.Pointer((*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(s)) + uintptr(8))))) = c64
		*(*uint64)(unsafe.Pointer((*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(s)) + uintptr(16))))) = c64
		*(*uint64)(unsafe.Pointer((*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(s)) + uintptr(24))))) = c64
	}
	return dest
}
