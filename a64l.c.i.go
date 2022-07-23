package libc

import unsafe "unsafe"

var _cgos_digits_a64l [65]int8 = [65]int8{'.', '/', '0', '1', '2', '3', '4', '5', '6', '7', '8', '9', 'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J', 'K', 'L', 'M', 'N', 'O', 'P', 'Q', 'R', 'S', 'T', 'U', 'V', 'W', 'X', 'Y', 'Z', 'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z', '\x00'}

func A64l(s *int8) int64 {
	var e int32
	var x uint32 = uint32(0)
	for e = int32(0); e < int32(36) && int32(*s) != 0; func() *int8 {
		e += int32(6)
		return func() (_cgo_ret *int8) {
			_cgo_addr := &s
			_cgo_ret = *_cgo_addr
			*(*uintptr)(unsafe.Pointer(_cgo_addr))++
			return
		}()
	}() {
		var d *int8 = Strchr((*int8)(unsafe.Pointer(&_cgos_digits_a64l)), int32(*s))
		if !(d != nil) {
			break
		}
		x |= uint32(uintptr(unsafe.Pointer(d))-uintptr(unsafe.Pointer((*int8)(unsafe.Pointer(&_cgos_digits_a64l))))) << e
	}
	return int64(int32(x))
}
func L64a(x0 int64) *int8 {
	var p *int8
	var x uint32 = uint32(x0)
	for p = (*int8)(unsafe.Pointer(&_cgos_s_a64l)); x != 0; func() uint32 {
		*(*uintptr)(unsafe.Pointer(&p))++
		return func() (_cgo_ret uint32) {
			_cgo_addr := &x
			*_cgo_addr >>= int32(6)
			return *_cgo_addr
		}()
	}() {
		*p = *(*int8)(unsafe.Pointer(uintptr(unsafe.Pointer((*int8)(unsafe.Pointer(&_cgos_digits_a64l)))) + uintptr(x&uint32(63))))
	}
	*p = int8(0)
	return (*int8)(unsafe.Pointer(&_cgos_s_a64l))
}

var _cgos_s_a64l [7]int8
