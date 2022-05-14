package libc

import unsafe "unsafe"

func Memset(dest unsafe.Pointer, c int32, n uint64) unsafe.Pointer {
	var s *uint8 = (*uint8)(dest)
	var k uint64
	if !(n != 0) {
		return dest
	}
	*(*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(s)) + uintptr(int32(0)))) = uint8(c)
	*(*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(s)) + uintptr(n-uint64(1)))) = uint8(c)
	if n <= uint64(2) {
		return dest
	}
	*(*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(s)) + uintptr(int32(1)))) = uint8(c)
	*(*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(s)) + uintptr(int32(2)))) = uint8(c)
	*(*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(s)) + uintptr(n-uint64(2)))) = uint8(c)
	*(*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(s)) + uintptr(n-uint64(3)))) = uint8(c)
	if n <= uint64(6) {
		return dest
	}
	*(*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(s)) + uintptr(int32(3)))) = uint8(c)
	*(*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(s)) + uintptr(n-uint64(4)))) = uint8(c)
	if n <= uint64(8) {
		return dest
	}
	k = -uint64(uintptr(unsafe.Pointer(s))) & uint64(3)
	*(*uintptr)(unsafe.Pointer(&s)) += uintptr(k)
	n -= k
	n &= uint64(18446744073709551612)
	type u32 = uint32
	type u64 = uint64
	var c32 uint32 = 16843009 * uint32(uint8(c))
	*(*uint32)(unsafe.Pointer((*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(s)) + uintptr(int32(0)))))) = c32
	*(*uint32)(unsafe.Pointer((*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer((*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(s))+uintptr(n))))) - uintptr(int32(4)))))) = c32
	if n <= uint64(8) {
		return dest
	}
	*(*uint32)(unsafe.Pointer((*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(s)) + uintptr(int32(4)))))) = c32
	*(*uint32)(unsafe.Pointer((*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(s)) + uintptr(int32(8)))))) = c32
	*(*uint32)(unsafe.Pointer((*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer((*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(s))+uintptr(n))))) - uintptr(int32(12)))))) = c32
	*(*uint32)(unsafe.Pointer((*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer((*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(s))+uintptr(n))))) - uintptr(int32(8)))))) = c32
	if n <= uint64(24) {
		return dest
	}
	*(*uint32)(unsafe.Pointer((*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(s)) + uintptr(int32(12)))))) = c32
	*(*uint32)(unsafe.Pointer((*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(s)) + uintptr(int32(16)))))) = c32
	*(*uint32)(unsafe.Pointer((*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(s)) + uintptr(int32(20)))))) = c32
	*(*uint32)(unsafe.Pointer((*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(s)) + uintptr(int32(24)))))) = c32
	*(*uint32)(unsafe.Pointer((*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer((*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(s))+uintptr(n))))) - uintptr(int32(28)))))) = c32
	*(*uint32)(unsafe.Pointer((*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer((*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(s))+uintptr(n))))) - uintptr(int32(24)))))) = c32
	*(*uint32)(unsafe.Pointer((*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer((*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(s))+uintptr(n))))) - uintptr(int32(20)))))) = c32
	*(*uint32)(unsafe.Pointer((*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer((*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(s))+uintptr(n))))) - uintptr(int32(16)))))) = c32
	k = uint64(24) + uint64(uintptr(unsafe.Pointer(s)))&uint64(4)
	*(*uintptr)(unsafe.Pointer(&s)) += uintptr(k)
	n -= k
	var c64 uint64 = uint64(c32) | uint64(c32)<<int32(32)
	for ; n >= uint64(32); func() *uint8 {
		n -= uint64(32)
		return func() (_cgo_ret *uint8) {
			_cgo_addr := &s
			*(*uintptr)(unsafe.Pointer(&*_cgo_addr)) += uintptr(int32(32))
			return *_cgo_addr
		}()
	}() {
		*(*uint64)(unsafe.Pointer((*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(s)) + uintptr(int32(0)))))) = c64
		*(*uint64)(unsafe.Pointer((*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(s)) + uintptr(int32(8)))))) = c64
		*(*uint64)(unsafe.Pointer((*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(s)) + uintptr(int32(16)))))) = c64
		*(*uint64)(unsafe.Pointer((*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(s)) + uintptr(int32(24)))))) = c64
	}
	return dest
}
