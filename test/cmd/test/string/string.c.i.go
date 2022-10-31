package string

import (
	unsafe "unsafe"
	libc "github.com/goplus/libc"
	common "github.com/goplus/libc/test/common"
	testing "testing"
)

func _cgo_main() int32 {
	var b [32]int8
	var s *int8
	var i int32
	*(*int8)(unsafe.Pointer(uintptr(unsafe.Pointer((*int8)(unsafe.Pointer(&b)))) + uintptr(int32(16)))) = int8('a')
	*(*int8)(unsafe.Pointer(uintptr(unsafe.Pointer((*int8)(unsafe.Pointer(&b)))) + uintptr(int32(17)))) = int8('b')
	*(*int8)(unsafe.Pointer(uintptr(unsafe.Pointer((*int8)(unsafe.Pointer(&b)))) + uintptr(int32(18)))) = int8('c')
	*(*int8)(unsafe.Pointer(uintptr(unsafe.Pointer((*int8)(unsafe.Pointer(&b)))) + uintptr(int32(19)))) = int8(0)
	if !(uintptr(unsafe.Pointer(func() (_cgo_ret *int8) {
		_cgo_addr := &s
		*_cgo_addr = libc.Strcpy((*int8)(unsafe.Pointer(&b)), (*int8)(unsafe.Pointer(uintptr(unsafe.Pointer((*int8)(unsafe.Pointer(&b))))+uintptr(int32(16)))))
		return *_cgo_addr
	}())) == uintptr(unsafe.Pointer((*int8)(unsafe.Pointer(&b))))) {
		func() int32 {
			common.T_printf((*int8)(unsafe.Pointer(&[91]int8{'/', 'U', 's', 'e', 'r', 's', '/', 'v', 'f', 'c', '/', 'g', 'o', 'p', 'l', 'u', 's', '/', 'l', 'i', 'b', 'c', '/', 't', 'e', 's', 't', '/', 's', 'r', 'c', '/', 'f', 'u', 'n', 'c', 't', 'i', 'o', 'n', 'a', 'l', '/', 's', 't', 'r', 'i', 'n', 'g', '.', 'c', ':', '2', '8', ':', ' ', '%', 's', ' ', 'f', 'a', 'i', 'l', 'e', 'd', ' ', '(', 'w', 'r', 'o', 'n', 'g', ' ', 'r', 'e', 't', 'u', 'r', 'n', ' ', '%', 'p', ' ', '!', '=', ' ', '%', 'p', ')', '\n', '\x00'})), (*int8)(unsafe.Pointer(&[16]int8{'s', 't', 'r', 'c', 'p', 'y', '(', 'b', ',', ' ', 'b', '+', '1', '6', ')', '\x00'})), s, (*int8)(unsafe.Pointer(&b)))
			return int32(0)
		}()
	}
	if !!(libc.Strcmp(s, (*int8)(unsafe.Pointer(&[4]int8{'a', 'b', 'c', '\x00'}))) != 0) {
		func() int32 {
			common.T_printf((*int8)(unsafe.Pointer(&[75]int8{'/', 'U', 's', 'e', 'r', 's', '/', 'v', 'f', 'c', '/', 'g', 'o', 'p', 'l', 'u', 's', '/', 'l', 'i', 'b', 'c', '/', 't', 'e', 's', 't', '/', 's', 'r', 'c', '/', 'f', 'u', 'n', 'c', 't', 'i', 'o', 'n', 'a', 'l', '/', 's', 't', 'r', 'i', 'n', 'g', '.', 'c', ':', '2', '9', ':', ' ', '[', '%', 's', ']', ' ', '!', '=', ' ', '[', '%', 's', ']', ' ', '(', '%', 's', ')', '\n', '\x00'})), s, (*int8)(unsafe.Pointer(&[4]int8{'a', 'b', 'c', '\x00'})), (*int8)(unsafe.Pointer(&[29]int8{'s', 't', 'r', 'c', 'p', 'y', ' ', 'g', 'a', 'v', 'e', ' ', 'i', 'n', 'c', 'o', 'r', 'r', 'e', 'c', 't', ' ', 's', 't', 'r', 'i', 'n', 'g', '\x00'})))
			return int32(0)
		}()
	}
	if !(uintptr(unsafe.Pointer(func() (_cgo_ret *int8) {
		_cgo_addr := &s
		*_cgo_addr = libc.Strcpy((*int8)(unsafe.Pointer(uintptr(unsafe.Pointer((*int8)(unsafe.Pointer(&b))))+uintptr(int32(1)))), (*int8)(unsafe.Pointer(uintptr(unsafe.Pointer((*int8)(unsafe.Pointer(&b))))+uintptr(int32(16)))))
		return *_cgo_addr
	}())) == uintptr(unsafe.Pointer((*int8)(unsafe.Pointer(uintptr(unsafe.Pointer((*int8)(unsafe.Pointer(&b))))+uintptr(int32(1))))))) {
		func() int32 {
			common.T_printf((*int8)(unsafe.Pointer(&[91]int8{'/', 'U', 's', 'e', 'r', 's', '/', 'v', 'f', 'c', '/', 'g', 'o', 'p', 'l', 'u', 's', '/', 'l', 'i', 'b', 'c', '/', 't', 'e', 's', 't', '/', 's', 'r', 'c', '/', 'f', 'u', 'n', 'c', 't', 'i', 'o', 'n', 'a', 'l', '/', 's', 't', 'r', 'i', 'n', 'g', '.', 'c', ':', '3', '0', ':', ' ', '%', 's', ' ', 'f', 'a', 'i', 'l', 'e', 'd', ' ', '(', 'w', 'r', 'o', 'n', 'g', ' ', 'r', 'e', 't', 'u', 'r', 'n', ' ', '%', 'p', ' ', '!', '=', ' ', '%', 'p', ')', '\n', '\x00'})), (*int8)(unsafe.Pointer(&[18]int8{'s', 't', 'r', 'c', 'p', 'y', '(', 'b', '+', '1', ',', ' ', 'b', '+', '1', '6', ')', '\x00'})), s, (*int8)(unsafe.Pointer(uintptr(unsafe.Pointer((*int8)(unsafe.Pointer(&b))))+uintptr(int32(1)))))
			return int32(0)
		}()
	}
	if !!(libc.Strcmp(s, (*int8)(unsafe.Pointer(&[4]int8{'a', 'b', 'c', '\x00'}))) != 0) {
		func() int32 {
			common.T_printf((*int8)(unsafe.Pointer(&[75]int8{'/', 'U', 's', 'e', 'r', 's', '/', 'v', 'f', 'c', '/', 'g', 'o', 'p', 'l', 'u', 's', '/', 'l', 'i', 'b', 'c', '/', 't', 'e', 's', 't', '/', 's', 'r', 'c', '/', 'f', 'u', 'n', 'c', 't', 'i', 'o', 'n', 'a', 'l', '/', 's', 't', 'r', 'i', 'n', 'g', '.', 'c', ':', '3', '1', ':', ' ', '[', '%', 's', ']', ' ', '!', '=', ' ', '[', '%', 's', ']', ' ', '(', '%', 's', ')', '\n', '\x00'})), s, (*int8)(unsafe.Pointer(&[4]int8{'a', 'b', 'c', '\x00'})), (*int8)(unsafe.Pointer(&[29]int8{'s', 't', 'r', 'c', 'p', 'y', ' ', 'g', 'a', 'v', 'e', ' ', 'i', 'n', 'c', 'o', 'r', 'r', 'e', 'c', 't', ' ', 's', 't', 'r', 'i', 'n', 'g', '\x00'})))
			return int32(0)
		}()
	}
	if !(uintptr(unsafe.Pointer(func() (_cgo_ret *int8) {
		_cgo_addr := &s
		*_cgo_addr = libc.Strcpy((*int8)(unsafe.Pointer(uintptr(unsafe.Pointer((*int8)(unsafe.Pointer(&b))))+uintptr(int32(2)))), (*int8)(unsafe.Pointer(uintptr(unsafe.Pointer((*int8)(unsafe.Pointer(&b))))+uintptr(int32(16)))))
		return *_cgo_addr
	}())) == uintptr(unsafe.Pointer((*int8)(unsafe.Pointer(uintptr(unsafe.Pointer((*int8)(unsafe.Pointer(&b))))+uintptr(int32(2))))))) {
		func() int32 {
			common.T_printf((*int8)(unsafe.Pointer(&[91]int8{'/', 'U', 's', 'e', 'r', 's', '/', 'v', 'f', 'c', '/', 'g', 'o', 'p', 'l', 'u', 's', '/', 'l', 'i', 'b', 'c', '/', 't', 'e', 's', 't', '/', 's', 'r', 'c', '/', 'f', 'u', 'n', 'c', 't', 'i', 'o', 'n', 'a', 'l', '/', 's', 't', 'r', 'i', 'n', 'g', '.', 'c', ':', '3', '2', ':', ' ', '%', 's', ' ', 'f', 'a', 'i', 'l', 'e', 'd', ' ', '(', 'w', 'r', 'o', 'n', 'g', ' ', 'r', 'e', 't', 'u', 'r', 'n', ' ', '%', 'p', ' ', '!', '=', ' ', '%', 'p', ')', '\n', '\x00'})), (*int8)(unsafe.Pointer(&[18]int8{'s', 't', 'r', 'c', 'p', 'y', '(', 'b', '+', '2', ',', ' ', 'b', '+', '1', '6', ')', '\x00'})), s, (*int8)(unsafe.Pointer(uintptr(unsafe.Pointer((*int8)(unsafe.Pointer(&b))))+uintptr(int32(2)))))
			return int32(0)
		}()
	}
	if !!(libc.Strcmp(s, (*int8)(unsafe.Pointer(&[4]int8{'a', 'b', 'c', '\x00'}))) != 0) {
		func() int32 {
			common.T_printf((*int8)(unsafe.Pointer(&[75]int8{'/', 'U', 's', 'e', 'r', 's', '/', 'v', 'f', 'c', '/', 'g', 'o', 'p', 'l', 'u', 's', '/', 'l', 'i', 'b', 'c', '/', 't', 'e', 's', 't', '/', 's', 'r', 'c', '/', 'f', 'u', 'n', 'c', 't', 'i', 'o', 'n', 'a', 'l', '/', 's', 't', 'r', 'i', 'n', 'g', '.', 'c', ':', '3', '3', ':', ' ', '[', '%', 's', ']', ' ', '!', '=', ' ', '[', '%', 's', ']', ' ', '(', '%', 's', ')', '\n', '\x00'})), s, (*int8)(unsafe.Pointer(&[4]int8{'a', 'b', 'c', '\x00'})), (*int8)(unsafe.Pointer(&[29]int8{'s', 't', 'r', 'c', 'p', 'y', ' ', 'g', 'a', 'v', 'e', ' ', 'i', 'n', 'c', 'o', 'r', 'r', 'e', 'c', 't', ' ', 's', 't', 'r', 'i', 'n', 'g', '\x00'})))
			return int32(0)
		}()
	}
	if !(uintptr(unsafe.Pointer(func() (_cgo_ret *int8) {
		_cgo_addr := &s
		*_cgo_addr = libc.Strcpy((*int8)(unsafe.Pointer(uintptr(unsafe.Pointer((*int8)(unsafe.Pointer(&b))))+uintptr(int32(3)))), (*int8)(unsafe.Pointer(uintptr(unsafe.Pointer((*int8)(unsafe.Pointer(&b))))+uintptr(int32(16)))))
		return *_cgo_addr
	}())) == uintptr(unsafe.Pointer((*int8)(unsafe.Pointer(uintptr(unsafe.Pointer((*int8)(unsafe.Pointer(&b))))+uintptr(int32(3))))))) {
		func() int32 {
			common.T_printf((*int8)(unsafe.Pointer(&[91]int8{'/', 'U', 's', 'e', 'r', 's', '/', 'v', 'f', 'c', '/', 'g', 'o', 'p', 'l', 'u', 's', '/', 'l', 'i', 'b', 'c', '/', 't', 'e', 's', 't', '/', 's', 'r', 'c', '/', 'f', 'u', 'n', 'c', 't', 'i', 'o', 'n', 'a', 'l', '/', 's', 't', 'r', 'i', 'n', 'g', '.', 'c', ':', '3', '4', ':', ' ', '%', 's', ' ', 'f', 'a', 'i', 'l', 'e', 'd', ' ', '(', 'w', 'r', 'o', 'n', 'g', ' ', 'r', 'e', 't', 'u', 'r', 'n', ' ', '%', 'p', ' ', '!', '=', ' ', '%', 'p', ')', '\n', '\x00'})), (*int8)(unsafe.Pointer(&[18]int8{'s', 't', 'r', 'c', 'p', 'y', '(', 'b', '+', '3', ',', ' ', 'b', '+', '1', '6', ')', '\x00'})), s, (*int8)(unsafe.Pointer(uintptr(unsafe.Pointer((*int8)(unsafe.Pointer(&b))))+uintptr(int32(3)))))
			return int32(0)
		}()
	}
	if !!(libc.Strcmp(s, (*int8)(unsafe.Pointer(&[4]int8{'a', 'b', 'c', '\x00'}))) != 0) {
		func() int32 {
			common.T_printf((*int8)(unsafe.Pointer(&[75]int8{'/', 'U', 's', 'e', 'r', 's', '/', 'v', 'f', 'c', '/', 'g', 'o', 'p', 'l', 'u', 's', '/', 'l', 'i', 'b', 'c', '/', 't', 'e', 's', 't', '/', 's', 'r', 'c', '/', 'f', 'u', 'n', 'c', 't', 'i', 'o', 'n', 'a', 'l', '/', 's', 't', 'r', 'i', 'n', 'g', '.', 'c', ':', '3', '5', ':', ' ', '[', '%', 's', ']', ' ', '!', '=', ' ', '[', '%', 's', ']', ' ', '(', '%', 's', ')', '\n', '\x00'})), s, (*int8)(unsafe.Pointer(&[4]int8{'a', 'b', 'c', '\x00'})), (*int8)(unsafe.Pointer(&[29]int8{'s', 't', 'r', 'c', 'p', 'y', ' ', 'g', 'a', 'v', 'e', ' ', 'i', 'n', 'c', 'o', 'r', 'r', 'e', 'c', 't', ' ', 's', 't', 'r', 'i', 'n', 'g', '\x00'})))
			return int32(0)
		}()
	}
	if !(uintptr(unsafe.Pointer(func() (_cgo_ret *int8) {
		_cgo_addr := &s
		*_cgo_addr = libc.Strcpy((*int8)(unsafe.Pointer(uintptr(unsafe.Pointer((*int8)(unsafe.Pointer(&b))))+uintptr(int32(1)))), (*int8)(unsafe.Pointer(uintptr(unsafe.Pointer((*int8)(unsafe.Pointer(&b))))+uintptr(int32(17)))))
		return *_cgo_addr
	}())) == uintptr(unsafe.Pointer((*int8)(unsafe.Pointer(uintptr(unsafe.Pointer((*int8)(unsafe.Pointer(&b))))+uintptr(int32(1))))))) {
		func() int32 {
			common.T_printf((*int8)(unsafe.Pointer(&[91]int8{'/', 'U', 's', 'e', 'r', 's', '/', 'v', 'f', 'c', '/', 'g', 'o', 'p', 'l', 'u', 's', '/', 'l', 'i', 'b', 'c', '/', 't', 'e', 's', 't', '/', 's', 'r', 'c', '/', 'f', 'u', 'n', 'c', 't', 'i', 'o', 'n', 'a', 'l', '/', 's', 't', 'r', 'i', 'n', 'g', '.', 'c', ':', '3', '7', ':', ' ', '%', 's', ' ', 'f', 'a', 'i', 'l', 'e', 'd', ' ', '(', 'w', 'r', 'o', 'n', 'g', ' ', 'r', 'e', 't', 'u', 'r', 'n', ' ', '%', 'p', ' ', '!', '=', ' ', '%', 'p', ')', '\n', '\x00'})), (*int8)(unsafe.Pointer(&[18]int8{'s', 't', 'r', 'c', 'p', 'y', '(', 'b', '+', '1', ',', ' ', 'b', '+', '1', '7', ')', '\x00'})), s, (*int8)(unsafe.Pointer(uintptr(unsafe.Pointer((*int8)(unsafe.Pointer(&b))))+uintptr(int32(1)))))
			return int32(0)
		}()
	}
	if !!(libc.Strcmp(s, (*int8)(unsafe.Pointer(&[3]int8{'b', 'c', '\x00'}))) != 0) {
		func() int32 {
			common.T_printf((*int8)(unsafe.Pointer(&[75]int8{'/', 'U', 's', 'e', 'r', 's', '/', 'v', 'f', 'c', '/', 'g', 'o', 'p', 'l', 'u', 's', '/', 'l', 'i', 'b', 'c', '/', 't', 'e', 's', 't', '/', 's', 'r', 'c', '/', 'f', 'u', 'n', 'c', 't', 'i', 'o', 'n', 'a', 'l', '/', 's', 't', 'r', 'i', 'n', 'g', '.', 'c', ':', '3', '8', ':', ' ', '[', '%', 's', ']', ' ', '!', '=', ' ', '[', '%', 's', ']', ' ', '(', '%', 's', ')', '\n', '\x00'})), s, (*int8)(unsafe.Pointer(&[3]int8{'b', 'c', '\x00'})), (*int8)(unsafe.Pointer(&[29]int8{'s', 't', 'r', 'c', 'p', 'y', ' ', 'g', 'a', 'v', 'e', ' ', 'i', 'n', 'c', 'o', 'r', 'r', 'e', 'c', 't', ' ', 's', 't', 'r', 'i', 'n', 'g', '\x00'})))
			return int32(0)
		}()
	}
	if !(uintptr(unsafe.Pointer(func() (_cgo_ret *int8) {
		_cgo_addr := &s
		*_cgo_addr = libc.Strcpy((*int8)(unsafe.Pointer(uintptr(unsafe.Pointer((*int8)(unsafe.Pointer(&b))))+uintptr(int32(2)))), (*int8)(unsafe.Pointer(uintptr(unsafe.Pointer((*int8)(unsafe.Pointer(&b))))+uintptr(int32(18)))))
		return *_cgo_addr
	}())) == uintptr(unsafe.Pointer((*int8)(unsafe.Pointer(uintptr(unsafe.Pointer((*int8)(unsafe.Pointer(&b))))+uintptr(int32(2))))))) {
		func() int32 {
			common.T_printf((*int8)(unsafe.Pointer(&[91]int8{'/', 'U', 's', 'e', 'r', 's', '/', 'v', 'f', 'c', '/', 'g', 'o', 'p', 'l', 'u', 's', '/', 'l', 'i', 'b', 'c', '/', 't', 'e', 's', 't', '/', 's', 'r', 'c', '/', 'f', 'u', 'n', 'c', 't', 'i', 'o', 'n', 'a', 'l', '/', 's', 't', 'r', 'i', 'n', 'g', '.', 'c', ':', '3', '9', ':', ' ', '%', 's', ' ', 'f', 'a', 'i', 'l', 'e', 'd', ' ', '(', 'w', 'r', 'o', 'n', 'g', ' ', 'r', 'e', 't', 'u', 'r', 'n', ' ', '%', 'p', ' ', '!', '=', ' ', '%', 'p', ')', '\n', '\x00'})), (*int8)(unsafe.Pointer(&[18]int8{'s', 't', 'r', 'c', 'p', 'y', '(', 'b', '+', '2', ',', ' ', 'b', '+', '1', '8', ')', '\x00'})), s, (*int8)(unsafe.Pointer(uintptr(unsafe.Pointer((*int8)(unsafe.Pointer(&b))))+uintptr(int32(2)))))
			return int32(0)
		}()
	}
	if !!(libc.Strcmp(s, (*int8)(unsafe.Pointer(&[2]int8{'c', '\x00'}))) != 0) {
		func() int32 {
			common.T_printf((*int8)(unsafe.Pointer(&[75]int8{'/', 'U', 's', 'e', 'r', 's', '/', 'v', 'f', 'c', '/', 'g', 'o', 'p', 'l', 'u', 's', '/', 'l', 'i', 'b', 'c', '/', 't', 'e', 's', 't', '/', 's', 'r', 'c', '/', 'f', 'u', 'n', 'c', 't', 'i', 'o', 'n', 'a', 'l', '/', 's', 't', 'r', 'i', 'n', 'g', '.', 'c', ':', '4', '0', ':', ' ', '[', '%', 's', ']', ' ', '!', '=', ' ', '[', '%', 's', ']', ' ', '(', '%', 's', ')', '\n', '\x00'})), s, (*int8)(unsafe.Pointer(&[2]int8{'c', '\x00'})), (*int8)(unsafe.Pointer(&[29]int8{'s', 't', 'r', 'c', 'p', 'y', ' ', 'g', 'a', 'v', 'e', ' ', 'i', 'n', 'c', 'o', 'r', 'r', 'e', 'c', 't', ' ', 's', 't', 'r', 'i', 'n', 'g', '\x00'})))
			return int32(0)
		}()
	}
	if !(uintptr(unsafe.Pointer(func() (_cgo_ret *int8) {
		_cgo_addr := &s
		*_cgo_addr = libc.Strcpy((*int8)(unsafe.Pointer(uintptr(unsafe.Pointer((*int8)(unsafe.Pointer(&b))))+uintptr(int32(3)))), (*int8)(unsafe.Pointer(uintptr(unsafe.Pointer((*int8)(unsafe.Pointer(&b))))+uintptr(int32(19)))))
		return *_cgo_addr
	}())) == uintptr(unsafe.Pointer((*int8)(unsafe.Pointer(uintptr(unsafe.Pointer((*int8)(unsafe.Pointer(&b))))+uintptr(int32(3))))))) {
		func() int32 {
			common.T_printf((*int8)(unsafe.Pointer(&[91]int8{'/', 'U', 's', 'e', 'r', 's', '/', 'v', 'f', 'c', '/', 'g', 'o', 'p', 'l', 'u', 's', '/', 'l', 'i', 'b', 'c', '/', 't', 'e', 's', 't', '/', 's', 'r', 'c', '/', 'f', 'u', 'n', 'c', 't', 'i', 'o', 'n', 'a', 'l', '/', 's', 't', 'r', 'i', 'n', 'g', '.', 'c', ':', '4', '1', ':', ' ', '%', 's', ' ', 'f', 'a', 'i', 'l', 'e', 'd', ' ', '(', 'w', 'r', 'o', 'n', 'g', ' ', 'r', 'e', 't', 'u', 'r', 'n', ' ', '%', 'p', ' ', '!', '=', ' ', '%', 'p', ')', '\n', '\x00'})), (*int8)(unsafe.Pointer(&[18]int8{'s', 't', 'r', 'c', 'p', 'y', '(', 'b', '+', '3', ',', ' ', 'b', '+', '1', '9', ')', '\x00'})), s, (*int8)(unsafe.Pointer(uintptr(unsafe.Pointer((*int8)(unsafe.Pointer(&b))))+uintptr(int32(3)))))
			return int32(0)
		}()
	}
	if !!(libc.Strcmp(s, (*int8)(unsafe.Pointer(&[1]int8{'\x00'}))) != 0) {
		func() int32 {
			common.T_printf((*int8)(unsafe.Pointer(&[75]int8{'/', 'U', 's', 'e', 'r', 's', '/', 'v', 'f', 'c', '/', 'g', 'o', 'p', 'l', 'u', 's', '/', 'l', 'i', 'b', 'c', '/', 't', 'e', 's', 't', '/', 's', 'r', 'c', '/', 'f', 'u', 'n', 'c', 't', 'i', 'o', 'n', 'a', 'l', '/', 's', 't', 'r', 'i', 'n', 'g', '.', 'c', ':', '4', '2', ':', ' ', '[', '%', 's', ']', ' ', '!', '=', ' ', '[', '%', 's', ']', ' ', '(', '%', 's', ')', '\n', '\x00'})), s, (*int8)(unsafe.Pointer(&[1]int8{'\x00'})), (*int8)(unsafe.Pointer(&[29]int8{'s', 't', 'r', 'c', 'p', 'y', ' ', 'g', 'a', 'v', 'e', ' ', 'i', 'n', 'c', 'o', 'r', 'r', 'e', 'c', 't', ' ', 's', 't', 'r', 'i', 'n', 'g', '\x00'})))
			return int32(0)
		}()
	}
	if !(uintptr(unsafe.Pointer(func() (_cgo_ret *int8) {
		_cgo_addr := &s
		*_cgo_addr = (*int8)(libc.Memset(unsafe.Pointer((*int8)(unsafe.Pointer(&b))), 'x', 32))
		return *_cgo_addr
	}())) == uintptr(unsafe.Pointer((*int8)(unsafe.Pointer(&b))))) {
		func() int32 {
			common.T_printf((*int8)(unsafe.Pointer(&[91]int8{'/', 'U', 's', 'e', 'r', 's', '/', 'v', 'f', 'c', '/', 'g', 'o', 'p', 'l', 'u', 's', '/', 'l', 'i', 'b', 'c', '/', 't', 'e', 's', 't', '/', 's', 'r', 'c', '/', 'f', 'u', 'n', 'c', 't', 'i', 'o', 'n', 'a', 'l', '/', 's', 't', 'r', 'i', 'n', 'g', '.', 'c', ':', '4', '4', ':', ' ', '%', 's', ' ', 'f', 'a', 'i', 'l', 'e', 'd', ' ', '(', 'w', 'r', 'o', 'n', 'g', ' ', 'r', 'e', 't', 'u', 'r', 'n', ' ', '%', 'p', ' ', '!', '=', ' ', '%', 'p', ')', '\n', '\x00'})), (*int8)(unsafe.Pointer(&[25]int8{'m', 'e', 'm', 's', 'e', 't', '(', 'b', ',', ' ', '\'', 'x', '\'', ',', ' ', 's', 'i', 'z', 'e', 'o', 'f', ' ', 'b', ')', '\x00'})), s, (*int8)(unsafe.Pointer(&b)))
			return int32(0)
		}()
	}
	if !(uintptr(unsafe.Pointer(func() (_cgo_ret *int8) {
		_cgo_addr := &s
		*_cgo_addr = libc.Strncpy((*int8)(unsafe.Pointer(&b)), (*int8)(unsafe.Pointer(&[4]int8{'a', 'b', 'c', '\x00'})), 31)
		return *_cgo_addr
	}())) == uintptr(unsafe.Pointer((*int8)(unsafe.Pointer(&b))))) {
		func() int32 {
			common.T_printf((*int8)(unsafe.Pointer(&[91]int8{'/', 'U', 's', 'e', 'r', 's', '/', 'v', 'f', 'c', '/', 'g', 'o', 'p', 'l', 'u', 's', '/', 'l', 'i', 'b', 'c', '/', 't', 'e', 's', 't', '/', 's', 'r', 'c', '/', 'f', 'u', 'n', 'c', 't', 'i', 'o', 'n', 'a', 'l', '/', 's', 't', 'r', 'i', 'n', 'g', '.', 'c', ':', '4', '5', ':', ' ', '%', 's', ' ', 'f', 'a', 'i', 'l', 'e', 'd', ' ', '(', 'w', 'r', 'o', 'n', 'g', ' ', 'r', 'e', 't', 'u', 'r', 'n', ' ', '%', 'p', ' ', '!', '=', ' ', '%', 'p', ')', '\n', '\x00'})), (*int8)(unsafe.Pointer(&[32]int8{'s', 't', 'r', 'n', 'c', 'p', 'y', '(', 'b', ',', ' ', '"', 'a', 'b', 'c', '"', ',', ' ', 's', 'i', 'z', 'e', 'o', 'f', ' ', 'b', ' ', '-', ' ', '1', ')', '\x00'})), s, (*int8)(unsafe.Pointer(&b)))
			return int32(0)
		}()
	}
	if !(func() (_cgo_ret int32) {
		_cgo_addr := &i
		*_cgo_addr = libc.Memcmp(unsafe.Pointer((*int8)(unsafe.Pointer(&b))), unsafe.Pointer((*int8)(unsafe.Pointer(&[8]int8{'a', 'b', 'c', '\x00', '\x00', '\x00', '\x00', '\x00'}))), uint64(8))
		return *_cgo_addr
	}() == int32(0)) {
		func() int32 {
			common.T_printf((*int8)(unsafe.Pointer(&[100]int8{'/', 'U', 's', 'e', 'r', 's', '/', 'v', 'f', 'c', '/', 'g', 'o', 'p', 'l', 'u', 's', '/', 'l', 'i', 'b', 'c', '/', 't', 'e', 's', 't', '/', 's', 'r', 'c', '/', 'f', 'u', 'n', 'c', 't', 'i', 'o', 'n', 'a', 'l', '/', 's', 't', 'r', 'i', 'n', 'g', '.', 'c', ':', '4', '6', ':', ' ', '%', 's', ' ', 'f', 'a', 'i', 'l', 'e', 'd', ' ', '(', 's', 't', 'r', 'n', 'c', 'p', 'y', ' ', 'f', 'a', 'i', 'l', 's', ' ', 't', 'o', ' ', 'z', 'e', 'r', 'o', '-', 'p', 'a', 'd', ' ', 'd', 'e', 's', 't', ')', '\n', '\x00'})), (*int8)(unsafe.Pointer(&[28]int8{'m', 'e', 'm', 'c', 'm', 'p', '(', 'b', ',', ' ', '"', 'a', 'b', 'c', '\\', '0', '\\', '0', '\\', '0', '\\', '0', '"', ',', ' ', '8', ')', '\x00'})), i, int32(0))
			return int32(0)
		}()
	}
	if !(func() (_cgo_ret int32) {
		_cgo_addr := &i
		*_cgo_addr = int32(*(*int8)(unsafe.Pointer(uintptr(unsafe.Pointer((*int8)(unsafe.Pointer(&b)))) + uintptr(31))))
		return *_cgo_addr
	}() == 'x') {
		func() int32 {
			common.T_printf((*int8)(unsafe.Pointer(&[114]int8{'/', 'U', 's', 'e', 'r', 's', '/', 'v', 'f', 'c', '/', 'g', 'o', 'p', 'l', 'u', 's', '/', 'l', 'i', 'b', 'c', '/', 't', 'e', 's', 't', '/', 's', 'r', 'c', '/', 'f', 'u', 'n', 'c', 't', 'i', 'o', 'n', 'a', 'l', '/', 's', 't', 'r', 'i', 'n', 'g', '.', 'c', ':', '4', '7', ':', ' ', '%', 's', ' ', 'f', 'a', 'i', 'l', 'e', 'd', ' ', '(', 's', 't', 'r', 'n', 'c', 'p', 'y', ' ', 'o', 'v', 'e', 'r', 'r', 'u', 'n', 's', ' ', 'b', 'u', 'f', 'f', 'e', 'r', ' ', 'w', 'h', 'e', 'n', ' ', 'n', ' ', '>', ' ', 's', 't', 'r', 'l', 'e', 'n', '(', 's', 'r', 'c', ')', ')', '\n', '\x00'})), (*int8)(unsafe.Pointer(&[16]int8{'b', '[', 's', 'i', 'z', 'e', 'o', 'f', ' ', 'b', ' ', '-', ' ', '1', ']', '\x00'})), i, 'x')
			return int32(0)
		}()
	}
	*(*int8)(unsafe.Pointer(uintptr(unsafe.Pointer((*int8)(unsafe.Pointer(&b)))) + uintptr(int32(3)))) = int8('x')
	*(*int8)(unsafe.Pointer(uintptr(unsafe.Pointer((*int8)(unsafe.Pointer(&b)))) + uintptr(int32(4)))) = int8(0)
	libc.Strncpy((*int8)(unsafe.Pointer(&b)), (*int8)(unsafe.Pointer(&[4]int8{'a', 'b', 'c', '\x00'})), uint64(3))
	if !(func() (_cgo_ret int32) {
		_cgo_addr := &i
		*_cgo_addr = int32(*(*int8)(unsafe.Pointer(uintptr(unsafe.Pointer((*int8)(unsafe.Pointer(&b)))) + uintptr(int32(2)))))
		return *_cgo_addr
	}() == 'c') {
		func() int32 {
			common.T_printf((*int8)(unsafe.Pointer(&[115]int8{'/', 'U', 's', 'e', 'r', 's', '/', 'v', 'f', 'c', '/', 'g', 'o', 'p', 'l', 'u', 's', '/', 'l', 'i', 'b', 'c', '/', 't', 'e', 's', 't', '/', 's', 'r', 'c', '/', 'f', 'u', 'n', 'c', 't', 'i', 'o', 'n', 'a', 'l', '/', 's', 't', 'r', 'i', 'n', 'g', '.', 'c', ':', '5', '1', ':', ' ', '%', 's', ' ', 'f', 'a', 'i', 'l', 'e', 'd', ' ', '(', 's', 't', 'r', 'n', 'c', 'p', 'y', ' ', 'f', 'a', 'i', 'l', 's', ' ', 't', 'o', ' ', 'c', 'o', 'p', 'y', ' ', 'l', 'a', 's', 't', ' ', 'b', 'y', 't', 'e', ':', ' ', '%', 'h', 'h', 'u', ' ', '!', '=', ' ', '%', 'h', 'h', 'u', ')', '\n', '\x00'})), (*int8)(unsafe.Pointer(&[5]int8{'b', '[', '2', ']', '\x00'})), i, 'c')
			return int32(0)
		}()
	}
	if !(func() (_cgo_ret int32) {
		_cgo_addr := &i
		*_cgo_addr = int32(*(*int8)(unsafe.Pointer(uintptr(unsafe.Pointer((*int8)(unsafe.Pointer(&b)))) + uintptr(int32(3)))))
		return *_cgo_addr
	}() == 'x') {
		func() int32 {
			common.T_printf((*int8)(unsafe.Pointer(&[125]int8{'/', 'U', 's', 'e', 'r', 's', '/', 'v', 'f', 'c', '/', 'g', 'o', 'p', 'l', 'u', 's', '/', 'l', 'i', 'b', 'c', '/', 't', 'e', 's', 't', '/', 's', 'r', 'c', '/', 'f', 'u', 'n', 'c', 't', 'i', 'o', 'n', 'a', 'l', '/', 's', 't', 'r', 'i', 'n', 'g', '.', 'c', ':', '5', '2', ':', ' ', '%', 's', ' ', 'f', 'a', 'i', 'l', 'e', 'd', ' ', '(', 's', 't', 'r', 'n', 'c', 'p', 'y', ' ', 'o', 'v', 'e', 'r', 'r', 'u', 'n', 's', ' ', 'b', 'u', 'f', 'f', 'e', 'r', ' ', 't', 'o', ' ', 'n', 'u', 'l', 'l', '-', 't', 'e', 'r', 'm', 'i', 'n', 'a', 't', 'e', ':', ' ', '%', 'h', 'h', 'u', ' ', '!', '=', ' ', '%', 'h', 'h', 'u', ')', '\n', '\x00'})), (*int8)(unsafe.Pointer(&[5]int8{'b', '[', '3', ']', '\x00'})), i, 'x')
			return int32(0)
		}()
	}
	if !(func() (_cgo_ret int32) {
		_cgo_addr := &i
		*_cgo_addr = func() int32 {
			if !(libc.Strncmp((*int8)(unsafe.Pointer(&[5]int8{'a', 'b', 'c', 'd', '\x00'})), (*int8)(unsafe.Pointer(&[5]int8{'a', 'b', 'c', 'e', '\x00'})), uint64(3)) != 0) {
				return 1
			} else {
				return 0
			}
		}()
		return *_cgo_addr
	}() == int32(1)) {
		func() int32 {
			common.T_printf((*int8)(unsafe.Pointer(&[93]int8{'/', 'U', 's', 'e', 'r', 's', '/', 'v', 'f', 'c', '/', 'g', 'o', 'p', 'l', 'u', 's', '/', 'l', 'i', 'b', 'c', '/', 't', 'e', 's', 't', '/', 's', 'r', 'c', '/', 'f', 'u', 'n', 'c', 't', 'i', 'o', 'n', 'a', 'l', '/', 's', 't', 'r', 'i', 'n', 'g', '.', 'c', ':', '5', '4', ':', ' ', '%', 's', ' ', 'f', 'a', 'i', 'l', 'e', 'd', ' ', '(', 's', 't', 'r', 'n', 'c', 'm', 'p', ' ', 'c', 'o', 'm', 'p', 'a', 'r', 'e', 's', ' ', 'p', 'a', 's', 't', ' ', 'n', ')', '\n', '\x00'})), (*int8)(unsafe.Pointer(&[28]int8{'!', 's', 't', 'r', 'n', 'c', 'm', 'p', '(', '"', 'a', 'b', 'c', 'd', '"', ',', ' ', '"', 'a', 'b', 'c', 'e', '"', ',', ' ', '3', ')', '\x00'})), i, int32(1))
			return int32(0)
		}()
	}
	if !(func() (_cgo_ret int32) {
		_cgo_addr := &i
		*_cgo_addr = func() int32 {
			if !!(libc.Strncmp((*int8)(unsafe.Pointer(&[4]int8{'a', 'b', 'c', '\x00'})), (*int8)(unsafe.Pointer(&[4]int8{'a', 'b', 'd', '\x00'})), uint64(3)) != 0) {
				return 1
			} else {
				return 0
			}
		}()
		return *_cgo_addr
	}() == int32(1)) {
		func() int32 {
			common.T_printf((*int8)(unsafe.Pointer(&[105]int8{'/', 'U', 's', 'e', 'r', 's', '/', 'v', 'f', 'c', '/', 'g', 'o', 'p', 'l', 'u', 's', '/', 'l', 'i', 'b', 'c', '/', 't', 'e', 's', 't', '/', 's', 'r', 'c', '/', 'f', 'u', 'n', 'c', 't', 'i', 'o', 'n', 'a', 'l', '/', 's', 't', 'r', 'i', 'n', 'g', '.', 'c', ':', '5', '5', ':', ' ', '%', 's', ' ', 'f', 'a', 'i', 'l', 'e', 'd', ' ', '(', 's', 't', 'r', 'n', 'c', 'm', 'p', ' ', 'f', 'a', 'i', 'l', 's', ' ', 't', 'o', ' ', 'c', 'o', 'm', 'p', 'a', 'r', 'e', ' ', 'n', '-', '1', 's', 't', ' ', 'b', 'y', 't', 'e', ')', '\n', '\x00'})), (*int8)(unsafe.Pointer(&[27]int8{'!', '!', 's', 't', 'r', 'n', 'c', 'm', 'p', '(', '"', 'a', 'b', 'c', '"', ',', ' ', '"', 'a', 'b', 'd', '"', ',', ' ', '3', ')', '\x00'})), i, int32(1))
			return int32(0)
		}()
	}
	libc.Strcpy((*int8)(unsafe.Pointer(&b)), (*int8)(unsafe.Pointer(&[4]int8{'a', 'b', 'c', '\x00'})))
	if !(uintptr(unsafe.Pointer(func() (_cgo_ret *int8) {
		_cgo_addr := &s
		*_cgo_addr = libc.Strncat((*int8)(unsafe.Pointer(&b)), (*int8)(unsafe.Pointer(&[7]int8{'1', '2', '3', '4', '5', '6', '\x00'})), uint64(3))
		return *_cgo_addr
	}())) == uintptr(unsafe.Pointer((*int8)(unsafe.Pointer(&b))))) {
		func() int32 {
			common.T_printf((*int8)(unsafe.Pointer(&[78]int8{'/', 'U', 's', 'e', 'r', 's', '/', 'v', 'f', 'c', '/', 'g', 'o', 'p', 'l', 'u', 's', '/', 'l', 'i', 'b', 'c', '/', 't', 'e', 's', 't', '/', 's', 'r', 'c', '/', 'f', 'u', 'n', 'c', 't', 'i', 'o', 'n', 'a', 'l', '/', 's', 't', 'r', 'i', 'n', 'g', '.', 'c', ':', '5', '8', ':', ' ', '%', 's', ' ', 'f', 'a', 'i', 'l', 'e', 'd', ' ', '(', '%', 'p', ' ', '!', '=', ' ', '%', 'p', ')', '\n', '\x00'})), (*int8)(unsafe.Pointer(&[24]int8{'s', 't', 'r', 'n', 'c', 'a', 't', '(', 'b', ',', ' ', '"', '1', '2', '3', '4', '5', '6', '"', ',', ' ', '3', ')', '\x00'})), s, (*int8)(unsafe.Pointer(&b)))
			return int32(0)
		}()
	}
	if !(func() (_cgo_ret int32) {
		_cgo_addr := &i
		*_cgo_addr = int32(*(*int8)(unsafe.Pointer(uintptr(unsafe.Pointer((*int8)(unsafe.Pointer(&b)))) + uintptr(int32(6)))))
		return *_cgo_addr
	}() == int32(0)) {
		func() int32 {
			common.T_printf((*int8)(unsafe.Pointer(&[107]int8{'/', 'U', 's', 'e', 'r', 's', '/', 'v', 'f', 'c', '/', 'g', 'o', 'p', 'l', 'u', 's', '/', 'l', 'i', 'b', 'c', '/', 't', 'e', 's', 't', '/', 's', 'r', 'c', '/', 'f', 'u', 'n', 'c', 't', 'i', 'o', 'n', 'a', 'l', '/', 's', 't', 'r', 'i', 'n', 'g', '.', 'c', ':', '5', '9', ':', ' ', '%', 's', ' ', 'f', 'a', 'i', 'l', 'e', 'd', ' ', '(', 's', 't', 'r', 'n', 'c', 'a', 't', ' ', 'f', 'a', 'i', 'l', 'e', 'd', ' ', 't', 'o', ' ', 'n', 'u', 'l', 'l', '-', 't', 'e', 'r', 'm', 'i', 'n', 'a', 't', 'e', ' ', '(', '%', 'd', ')', ')', '\n', '\x00'})), (*int8)(unsafe.Pointer(&[5]int8{'b', '[', '6', ']', '\x00'})), i, int32(0))
			return int32(0)
		}()
	}
	if !!(libc.Strcmp(s, (*int8)(unsafe.Pointer(&[7]int8{'a', 'b', 'c', '1', '2', '3', '\x00'}))) != 0) {
		func() int32 {
			common.T_printf((*int8)(unsafe.Pointer(&[75]int8{'/', 'U', 's', 'e', 'r', 's', '/', 'v', 'f', 'c', '/', 'g', 'o', 'p', 'l', 'u', 's', '/', 'l', 'i', 'b', 'c', '/', 't', 'e', 's', 't', '/', 's', 'r', 'c', '/', 'f', 'u', 'n', 'c', 't', 'i', 'o', 'n', 'a', 'l', '/', 's', 't', 'r', 'i', 'n', 'g', '.', 'c', ':', '6', '0', ':', ' ', '[', '%', 's', ']', ' ', '!', '=', ' ', '[', '%', 's', ']', ' ', '(', '%', 's', ')', '\n', '\x00'})), s, (*int8)(unsafe.Pointer(&[7]int8{'a', 'b', 'c', '1', '2', '3', '\x00'})), (*int8)(unsafe.Pointer(&[30]int8{'s', 't', 'r', 'n', 'c', 'a', 't', ' ', 'g', 'a', 'v', 'e', ' ', 'i', 'n', 'c', 'o', 'r', 'r', 'e', 'c', 't', ' ', 's', 't', 'r', 'i', 'n', 'g', '\x00'})))
			return int32(0)
		}()
	}
	libc.Strcpy((*int8)(unsafe.Pointer(&b)), (*int8)(unsafe.Pointer(&[21]int8{'a', 'a', 'a', 'b', 'a', 'b', 'c', 'c', 'd', 'd', '0', '0', '0', '1', '1', '2', '2', '2', '2', '3', '\x00'})))
	if !(uintptr(unsafe.Pointer(func() (_cgo_ret *int8) {
		_cgo_addr := &s
		*_cgo_addr = libc.Strchr((*int8)(unsafe.Pointer(&b)), 'b')
		return *_cgo_addr
	}())) == uintptr(unsafe.Pointer((*int8)(unsafe.Pointer(uintptr(unsafe.Pointer((*int8)(unsafe.Pointer(&b))))+uintptr(int32(3))))))) {
		func() int32 {
			common.T_printf((*int8)(unsafe.Pointer(&[78]int8{'/', 'U', 's', 'e', 'r', 's', '/', 'v', 'f', 'c', '/', 'g', 'o', 'p', 'l', 'u', 's', '/', 'l', 'i', 'b', 'c', '/', 't', 'e', 's', 't', '/', 's', 'r', 'c', '/', 'f', 'u', 'n', 'c', 't', 'i', 'o', 'n', 'a', 'l', '/', 's', 't', 'r', 'i', 'n', 'g', '.', 'c', ':', '6', '3', ':', ' ', '%', 's', ' ', 'f', 'a', 'i', 'l', 'e', 'd', ' ', '(', '%', 'p', ' ', '!', '=', ' ', '%', 'p', ')', '\n', '\x00'})), (*int8)(unsafe.Pointer(&[15]int8{'s', 't', 'r', 'c', 'h', 'r', '(', 'b', ',', ' ', '\'', 'b', '\'', ')', '\x00'})), s, (*int8)(unsafe.Pointer(uintptr(unsafe.Pointer((*int8)(unsafe.Pointer(&b))))+uintptr(int32(3)))))
			return int32(0)
		}()
	}
	if !(uintptr(unsafe.Pointer(func() (_cgo_ret *int8) {
		_cgo_addr := &s
		*_cgo_addr = libc.Strrchr((*int8)(unsafe.Pointer(&b)), 'b')
		return *_cgo_addr
	}())) == uintptr(unsafe.Pointer((*int8)(unsafe.Pointer(uintptr(unsafe.Pointer((*int8)(unsafe.Pointer(&b))))+uintptr(int32(5))))))) {
		func() int32 {
			common.T_printf((*int8)(unsafe.Pointer(&[78]int8{'/', 'U', 's', 'e', 'r', 's', '/', 'v', 'f', 'c', '/', 'g', 'o', 'p', 'l', 'u', 's', '/', 'l', 'i', 'b', 'c', '/', 't', 'e', 's', 't', '/', 's', 'r', 'c', '/', 'f', 'u', 'n', 'c', 't', 'i', 'o', 'n', 'a', 'l', '/', 's', 't', 'r', 'i', 'n', 'g', '.', 'c', ':', '6', '4', ':', ' ', '%', 's', ' ', 'f', 'a', 'i', 'l', 'e', 'd', ' ', '(', '%', 'p', ' ', '!', '=', ' ', '%', 'p', ')', '\n', '\x00'})), (*int8)(unsafe.Pointer(&[16]int8{'s', 't', 'r', 'r', 'c', 'h', 'r', '(', 'b', ',', ' ', '\'', 'b', '\'', ')', '\x00'})), s, (*int8)(unsafe.Pointer(uintptr(unsafe.Pointer((*int8)(unsafe.Pointer(&b))))+uintptr(int32(5)))))
			return int32(0)
		}()
	}
	if !(func() (_cgo_ret int32) {
		_cgo_addr := &i
		*_cgo_addr = int32(libc.Strspn((*int8)(unsafe.Pointer(&b)), (*int8)(unsafe.Pointer(&[5]int8{'a', 'b', 'c', 'd', '\x00'}))))
		return *_cgo_addr
	}() == int32(10)) {
		func() int32 {
			common.T_printf((*int8)(unsafe.Pointer(&[78]int8{'/', 'U', 's', 'e', 'r', 's', '/', 'v', 'f', 'c', '/', 'g', 'o', 'p', 'l', 'u', 's', '/', 'l', 'i', 'b', 'c', '/', 't', 'e', 's', 't', '/', 's', 'r', 'c', '/', 'f', 'u', 'n', 'c', 't', 'i', 'o', 'n', 'a', 'l', '/', 's', 't', 'r', 'i', 'n', 'g', '.', 'c', ':', '6', '5', ':', ' ', '%', 's', ' ', 'f', 'a', 'i', 'l', 'e', 'd', ' ', '(', '%', 'd', ' ', '!', '=', ' ', '%', 'd', ')', '\n', '\x00'})), (*int8)(unsafe.Pointer(&[18]int8{'s', 't', 'r', 's', 'p', 'n', '(', 'b', ',', ' ', '"', 'a', 'b', 'c', 'd', '"', ')', '\x00'})), i, int32(10))
			return int32(0)
		}()
	}
	if !(func() (_cgo_ret int32) {
		_cgo_addr := &i
		*_cgo_addr = int32(libc.Strcspn((*int8)(unsafe.Pointer(&b)), (*int8)(unsafe.Pointer(&[5]int8{'0', '1', '2', '3', '\x00'}))))
		return *_cgo_addr
	}() == int32(10)) {
		func() int32 {
			common.T_printf((*int8)(unsafe.Pointer(&[78]int8{'/', 'U', 's', 'e', 'r', 's', '/', 'v', 'f', 'c', '/', 'g', 'o', 'p', 'l', 'u', 's', '/', 'l', 'i', 'b', 'c', '/', 't', 'e', 's', 't', '/', 's', 'r', 'c', '/', 'f', 'u', 'n', 'c', 't', 'i', 'o', 'n', 'a', 'l', '/', 's', 't', 'r', 'i', 'n', 'g', '.', 'c', ':', '6', '6', ':', ' ', '%', 's', ' ', 'f', 'a', 'i', 'l', 'e', 'd', ' ', '(', '%', 'd', ' ', '!', '=', ' ', '%', 'd', ')', '\n', '\x00'})), (*int8)(unsafe.Pointer(&[19]int8{'s', 't', 'r', 'c', 's', 'p', 'n', '(', 'b', ',', ' ', '"', '0', '1', '2', '3', '"', ')', '\x00'})), i, int32(10))
			return int32(0)
		}()
	}
	if !(uintptr(unsafe.Pointer(func() (_cgo_ret *int8) {
		_cgo_addr := &s
		*_cgo_addr = libc.Strpbrk((*int8)(unsafe.Pointer(&b)), (*int8)(unsafe.Pointer(&[5]int8{'0', '1', '2', '3', '\x00'})))
		return *_cgo_addr
	}())) == uintptr(unsafe.Pointer((*int8)(unsafe.Pointer(uintptr(unsafe.Pointer((*int8)(unsafe.Pointer(&b))))+uintptr(int32(10))))))) {
		func() int32 {
			common.T_printf((*int8)(unsafe.Pointer(&[78]int8{'/', 'U', 's', 'e', 'r', 's', '/', 'v', 'f', 'c', '/', 'g', 'o', 'p', 'l', 'u', 's', '/', 'l', 'i', 'b', 'c', '/', 't', 'e', 's', 't', '/', 's', 'r', 'c', '/', 'f', 'u', 'n', 'c', 't', 'i', 'o', 'n', 'a', 'l', '/', 's', 't', 'r', 'i', 'n', 'g', '.', 'c', ':', '6', '7', ':', ' ', '%', 's', ' ', 'f', 'a', 'i', 'l', 'e', 'd', ' ', '(', '%', 'd', ' ', '!', '=', ' ', '%', 'd', ')', '\n', '\x00'})), (*int8)(unsafe.Pointer(&[19]int8{'s', 't', 'r', 'p', 'b', 'r', 'k', '(', 'b', ',', ' ', '"', '0', '1', '2', '3', '"', ')', '\x00'})), s, (*int8)(unsafe.Pointer(uintptr(unsafe.Pointer((*int8)(unsafe.Pointer(&b))))+uintptr(int32(10)))))
			return int32(0)
		}()
	}
	libc.Strcpy((*int8)(unsafe.Pointer(&b)), (*int8)(unsafe.Pointer(&[20]int8{'a', 'b', 'c', ' ', ' ', ' ', '1', '2', '3', ';', ' ', 'x', 'y', 'z', ';', ' ', 'f', 'o', 'o', '\x00'})))
	if !(uintptr(unsafe.Pointer(func() (_cgo_ret *int8) {
		_cgo_addr := &s
		*_cgo_addr = libc.Strtok((*int8)(unsafe.Pointer(&b)), (*int8)(unsafe.Pointer(&[2]int8{' ', '\x00'})))
		return *_cgo_addr
	}())) == uintptr(unsafe.Pointer((*int8)(unsafe.Pointer(&b))))) {
		func() int32 {
			common.T_printf((*int8)(unsafe.Pointer(&[78]int8{'/', 'U', 's', 'e', 'r', 's', '/', 'v', 'f', 'c', '/', 'g', 'o', 'p', 'l', 'u', 's', '/', 'l', 'i', 'b', 'c', '/', 't', 'e', 's', 't', '/', 's', 'r', 'c', '/', 'f', 'u', 'n', 'c', 't', 'i', 'o', 'n', 'a', 'l', '/', 's', 't', 'r', 'i', 'n', 'g', '.', 'c', ':', '7', '0', ':', ' ', '%', 's', ' ', 'f', 'a', 'i', 'l', 'e', 'd', ' ', '(', '%', 'p', ' ', '!', '=', ' ', '%', 'p', ')', '\n', '\x00'})), (*int8)(unsafe.Pointer(&[15]int8{'s', 't', 'r', 't', 'o', 'k', '(', 'b', ',', ' ', '"', ' ', '"', ')', '\x00'})), s, (*int8)(unsafe.Pointer(&b)))
			return int32(0)
		}()
	}
	if !!(libc.Strcmp(s, (*int8)(unsafe.Pointer(&[4]int8{'a', 'b', 'c', '\x00'}))) != 0) {
		func() int32 {
			common.T_printf((*int8)(unsafe.Pointer(&[75]int8{'/', 'U', 's', 'e', 'r', 's', '/', 'v', 'f', 'c', '/', 'g', 'o', 'p', 'l', 'u', 's', '/', 'l', 'i', 'b', 'c', '/', 't', 'e', 's', 't', '/', 's', 'r', 'c', '/', 'f', 'u', 'n', 'c', 't', 'i', 'o', 'n', 'a', 'l', '/', 's', 't', 'r', 'i', 'n', 'g', '.', 'c', ':', '7', '1', ':', ' ', '[', '%', 's', ']', ' ', '!', '=', ' ', '[', '%', 's', ']', ' ', '(', '%', 's', ')', '\n', '\x00'})), s, (*int8)(unsafe.Pointer(&[4]int8{'a', 'b', 'c', '\x00'})), (*int8)(unsafe.Pointer(&[14]int8{'s', 't', 'r', 't', 'o', 'k', ' ', 'r', 'e', 's', 'u', 'l', 't', '\x00'})))
			return int32(0)
		}()
	}
	if !(uintptr(unsafe.Pointer(func() (_cgo_ret *int8) {
		_cgo_addr := &s
		*_cgo_addr = libc.Strtok(nil, (*int8)(unsafe.Pointer(&[2]int8{';', '\x00'})))
		return *_cgo_addr
	}())) == uintptr(unsafe.Pointer((*int8)(unsafe.Pointer(uintptr(unsafe.Pointer((*int8)(unsafe.Pointer(&b))))+uintptr(int32(4))))))) {
		func() int32 {
			common.T_printf((*int8)(unsafe.Pointer(&[78]int8{'/', 'U', 's', 'e', 'r', 's', '/', 'v', 'f', 'c', '/', 'g', 'o', 'p', 'l', 'u', 's', '/', 'l', 'i', 'b', 'c', '/', 't', 'e', 's', 't', '/', 's', 'r', 'c', '/', 'f', 'u', 'n', 'c', 't', 'i', 'o', 'n', 'a', 'l', '/', 's', 't', 'r', 'i', 'n', 'g', '.', 'c', ':', '7', '3', ':', ' ', '%', 's', ' ', 'f', 'a', 'i', 'l', 'e', 'd', ' ', '(', '%', 'p', ' ', '!', '=', ' ', '%', 'p', ')', '\n', '\x00'})), (*int8)(unsafe.Pointer(&[18]int8{'s', 't', 'r', 't', 'o', 'k', '(', 'N', 'U', 'L', 'L', ',', ' ', '"', ';', '"', ')', '\x00'})), s, (*int8)(unsafe.Pointer(uintptr(unsafe.Pointer((*int8)(unsafe.Pointer(&b))))+uintptr(int32(4)))))
			return int32(0)
		}()
	}
	if !!(libc.Strcmp(s, (*int8)(unsafe.Pointer(&[6]int8{' ', ' ', '1', '2', '3', '\x00'}))) != 0) {
		func() int32 {
			common.T_printf((*int8)(unsafe.Pointer(&[75]int8{'/', 'U', 's', 'e', 'r', 's', '/', 'v', 'f', 'c', '/', 'g', 'o', 'p', 'l', 'u', 's', '/', 'l', 'i', 'b', 'c', '/', 't', 'e', 's', 't', '/', 's', 'r', 'c', '/', 'f', 'u', 'n', 'c', 't', 'i', 'o', 'n', 'a', 'l', '/', 's', 't', 'r', 'i', 'n', 'g', '.', 'c', ':', '7', '4', ':', ' ', '[', '%', 's', ']', ' ', '!', '=', ' ', '[', '%', 's', ']', ' ', '(', '%', 's', ')', '\n', '\x00'})), s, (*int8)(unsafe.Pointer(&[6]int8{' ', ' ', '1', '2', '3', '\x00'})), (*int8)(unsafe.Pointer(&[14]int8{'s', 't', 'r', 't', 'o', 'k', ' ', 'r', 'e', 's', 'u', 'l', 't', '\x00'})))
			return int32(0)
		}()
	}
	if !(uintptr(unsafe.Pointer(func() (_cgo_ret *int8) {
		_cgo_addr := &s
		*_cgo_addr = libc.Strtok(nil, (*int8)(unsafe.Pointer(&[3]int8{' ', ';', '\x00'})))
		return *_cgo_addr
	}())) == uintptr(unsafe.Pointer((*int8)(unsafe.Pointer(uintptr(unsafe.Pointer((*int8)(unsafe.Pointer(&b))))+uintptr(int32(11))))))) {
		func() int32 {
			common.T_printf((*int8)(unsafe.Pointer(&[78]int8{'/', 'U', 's', 'e', 'r', 's', '/', 'v', 'f', 'c', '/', 'g', 'o', 'p', 'l', 'u', 's', '/', 'l', 'i', 'b', 'c', '/', 't', 'e', 's', 't', '/', 's', 'r', 'c', '/', 'f', 'u', 'n', 'c', 't', 'i', 'o', 'n', 'a', 'l', '/', 's', 't', 'r', 'i', 'n', 'g', '.', 'c', ':', '7', '6', ':', ' ', '%', 's', ' ', 'f', 'a', 'i', 'l', 'e', 'd', ' ', '(', '%', 'p', ' ', '!', '=', ' ', '%', 'p', ')', '\n', '\x00'})), (*int8)(unsafe.Pointer(&[19]int8{'s', 't', 'r', 't', 'o', 'k', '(', 'N', 'U', 'L', 'L', ',', ' ', '"', ' ', ';', '"', ')', '\x00'})), s, (*int8)(unsafe.Pointer(uintptr(unsafe.Pointer((*int8)(unsafe.Pointer(&b))))+uintptr(int32(11)))))
			return int32(0)
		}()
	}
	if !!(libc.Strcmp(s, (*int8)(unsafe.Pointer(&[4]int8{'x', 'y', 'z', '\x00'}))) != 0) {
		func() int32 {
			common.T_printf((*int8)(unsafe.Pointer(&[75]int8{'/', 'U', 's', 'e', 'r', 's', '/', 'v', 'f', 'c', '/', 'g', 'o', 'p', 'l', 'u', 's', '/', 'l', 'i', 'b', 'c', '/', 't', 'e', 's', 't', '/', 's', 'r', 'c', '/', 'f', 'u', 'n', 'c', 't', 'i', 'o', 'n', 'a', 'l', '/', 's', 't', 'r', 'i', 'n', 'g', '.', 'c', ':', '7', '7', ':', ' ', '[', '%', 's', ']', ' ', '!', '=', ' ', '[', '%', 's', ']', ' ', '(', '%', 's', ')', '\n', '\x00'})), s, (*int8)(unsafe.Pointer(&[4]int8{'x', 'y', 'z', '\x00'})), (*int8)(unsafe.Pointer(&[14]int8{'s', 't', 'r', 't', 'o', 'k', ' ', 'r', 'e', 's', 'u', 'l', 't', '\x00'})))
			return int32(0)
		}()
	}
	if !(uintptr(unsafe.Pointer(func() (_cgo_ret *int8) {
		_cgo_addr := &s
		*_cgo_addr = libc.Strtok(nil, (*int8)(unsafe.Pointer(&[3]int8{' ', ';', '\x00'})))
		return *_cgo_addr
	}())) == uintptr(unsafe.Pointer((*int8)(unsafe.Pointer(uintptr(unsafe.Pointer((*int8)(unsafe.Pointer(&b))))+uintptr(int32(16))))))) {
		func() int32 {
			common.T_printf((*int8)(unsafe.Pointer(&[78]int8{'/', 'U', 's', 'e', 'r', 's', '/', 'v', 'f', 'c', '/', 'g', 'o', 'p', 'l', 'u', 's', '/', 'l', 'i', 'b', 'c', '/', 't', 'e', 's', 't', '/', 's', 'r', 'c', '/', 'f', 'u', 'n', 'c', 't', 'i', 'o', 'n', 'a', 'l', '/', 's', 't', 'r', 'i', 'n', 'g', '.', 'c', ':', '7', '9', ':', ' ', '%', 's', ' ', 'f', 'a', 'i', 'l', 'e', 'd', ' ', '(', '%', 'p', ' ', '!', '=', ' ', '%', 'p', ')', '\n', '\x00'})), (*int8)(unsafe.Pointer(&[19]int8{'s', 't', 'r', 't', 'o', 'k', '(', 'N', 'U', 'L', 'L', ',', ' ', '"', ' ', ';', '"', ')', '\x00'})), s, (*int8)(unsafe.Pointer(uintptr(unsafe.Pointer((*int8)(unsafe.Pointer(&b))))+uintptr(int32(16)))))
			return int32(0)
		}()
	}
	if !!(libc.Strcmp(s, (*int8)(unsafe.Pointer(&[4]int8{'f', 'o', 'o', '\x00'}))) != 0) {
		func() int32 {
			common.T_printf((*int8)(unsafe.Pointer(&[75]int8{'/', 'U', 's', 'e', 'r', 's', '/', 'v', 'f', 'c', '/', 'g', 'o', 'p', 'l', 'u', 's', '/', 'l', 'i', 'b', 'c', '/', 't', 'e', 's', 't', '/', 's', 'r', 'c', '/', 'f', 'u', 'n', 'c', 't', 'i', 'o', 'n', 'a', 'l', '/', 's', 't', 'r', 'i', 'n', 'g', '.', 'c', ':', '8', '0', ':', ' ', '[', '%', 's', ']', ' ', '!', '=', ' ', '[', '%', 's', ']', ' ', '(', '%', 's', ')', '\n', '\x00'})), s, (*int8)(unsafe.Pointer(&[4]int8{'f', 'o', 'o', '\x00'})), (*int8)(unsafe.Pointer(&[14]int8{'s', 't', 'r', 't', 'o', 'k', ' ', 'r', 'e', 's', 'u', 'l', 't', '\x00'})))
			return int32(0)
		}()
	}
	libc.Memset(unsafe.Pointer((*int8)(unsafe.Pointer(&b))), 'x', 32)
	if !(func() (_cgo_ret int32) {
		_cgo_addr := &i
		*_cgo_addr = int32(libc.Strlcpy((*int8)(unsafe.Pointer(&b)), (*int8)(unsafe.Pointer(&[4]int8{'a', 'b', 'c', '\x00'})), 31))
		return *_cgo_addr
	}() == int32(3)) {
		func() int32 {
			common.T_printf((*int8)(unsafe.Pointer(&[85]int8{'/', 'U', 's', 'e', 'r', 's', '/', 'v', 'f', 'c', '/', 'g', 'o', 'p', 'l', 'u', 's', '/', 'l', 'i', 'b', 'c', '/', 't', 'e', 's', 't', '/', 's', 'r', 'c', '/', 'f', 'u', 'n', 'c', 't', 'i', 'o', 'n', 'a', 'l', '/', 's', 't', 'r', 'i', 'n', 'g', '.', 'c', ':', '8', '3', ':', ' ', '%', 's', ' ', 'f', 'a', 'i', 'l', 'e', 'd', ' ', '(', 'l', 'e', 'n', 'g', 't', 'h', ' ', '%', 'd', ' ', '!', '=', ' ', '%', 'd', ')', '\n', '\x00'})), (*int8)(unsafe.Pointer(&[32]int8{'s', 't', 'r', 'l', 'c', 'p', 'y', '(', 'b', ',', ' ', '"', 'a', 'b', 'c', '"', ',', ' ', 's', 'i', 'z', 'e', 'o', 'f', ' ', 'b', ' ', '-', ' ', '1', ')', '\x00'})), i, int32(3))
			return int32(0)
		}()
	}
	if !(func() (_cgo_ret int32) {
		_cgo_addr := &i
		*_cgo_addr = int32(*(*int8)(unsafe.Pointer(uintptr(unsafe.Pointer((*int8)(unsafe.Pointer(&b)))) + uintptr(int32(3)))))
		return *_cgo_addr
	}() == int32(0)) {
		func() int32 {
			common.T_printf((*int8)(unsafe.Pointer(&[118]int8{'/', 'U', 's', 'e', 'r', 's', '/', 'v', 'f', 'c', '/', 'g', 'o', 'p', 'l', 'u', 's', '/', 'l', 'i', 'b', 'c', '/', 't', 'e', 's', 't', '/', 's', 'r', 'c', '/', 'f', 'u', 'n', 'c', 't', 'i', 'o', 'n', 'a', 'l', '/', 's', 't', 'r', 'i', 'n', 'g', '.', 'c', ':', '8', '4', ':', ' ', '%', 's', ' ', 'f', 'a', 'i', 'l', 'e', 'd', ' ', '(', 's', 't', 'r', 'l', 'c', 'p', 'y', ' ', 'd', 'i', 'd', ' ', 'n', 'o', 't', ' ', 'n', 'u', 'l', 'l', '-', 't', 'e', 'r', 'm', 'i', 'n', 'a', 't', 'e', ' ', 's', 'h', 'o', 'r', 't', ' ', 's', 't', 'r', 'i', 'n', 'g', ' ', '(', '%', 'd', ')', ')', '\n', '\x00'})), (*int8)(unsafe.Pointer(&[5]int8{'b', '[', '3', ']', '\x00'})), i, int32(0))
			return int32(0)
		}()
	}
	if !(func() (_cgo_ret int32) {
		_cgo_addr := &i
		*_cgo_addr = int32(*(*int8)(unsafe.Pointer(uintptr(unsafe.Pointer((*int8)(unsafe.Pointer(&b)))) + uintptr(int32(4)))))
		return *_cgo_addr
	}() == 'x') {
		func() int32 {
			common.T_printf((*int8)(unsafe.Pointer(&[100]int8{'/', 'U', 's', 'e', 'r', 's', '/', 'v', 'f', 'c', '/', 'g', 'o', 'p', 'l', 'u', 's', '/', 'l', 'i', 'b', 'c', '/', 't', 'e', 's', 't', '/', 's', 'r', 'c', '/', 'f', 'u', 'n', 'c', 't', 'i', 'o', 'n', 'a', 'l', '/', 's', 't', 'r', 'i', 'n', 'g', '.', 'c', ':', '8', '5', ':', ' ', '%', 's', ' ', 'f', 'a', 'i', 'l', 'e', 'd', ' ', '(', 's', 't', 'r', 'l', 'c', 'p', 'y', ' ', 'w', 'r', 'o', 't', 'e', ' ', 'e', 'x', 't', 'r', 'a', ' ', 'b', 'y', 't', 'e', 's', ' ', '(', '%', 'd', ')', ')', '\n', '\x00'})), (*int8)(unsafe.Pointer(&[5]int8{'b', '[', '4', ']', '\x00'})), i, 'x')
			return int32(0)
		}()
	}
	libc.Memset(unsafe.Pointer((*int8)(unsafe.Pointer(&b))), 'x', 32)
	if !(func() (_cgo_ret int32) {
		_cgo_addr := &i
		*_cgo_addr = int32(libc.Strlcpy((*int8)(unsafe.Pointer(&b)), (*int8)(unsafe.Pointer(&[4]int8{'a', 'b', 'c', '\x00'})), uint64(2)))
		return *_cgo_addr
	}() == int32(3)) {
		func() int32 {
			common.T_printf((*int8)(unsafe.Pointer(&[85]int8{'/', 'U', 's', 'e', 'r', 's', '/', 'v', 'f', 'c', '/', 'g', 'o', 'p', 'l', 'u', 's', '/', 'l', 'i', 'b', 'c', '/', 't', 'e', 's', 't', '/', 's', 'r', 'c', '/', 'f', 'u', 'n', 'c', 't', 'i', 'o', 'n', 'a', 'l', '/', 's', 't', 'r', 'i', 'n', 'g', '.', 'c', ':', '8', '8', ':', ' ', '%', 's', ' ', 'f', 'a', 'i', 'l', 'e', 'd', ' ', '(', 'l', 'e', 'n', 'g', 't', 'h', ' ', '%', 'd', ' ', '!', '=', ' ', '%', 'd', ')', '\n', '\x00'})), (*int8)(unsafe.Pointer(&[21]int8{'s', 't', 'r', 'l', 'c', 'p', 'y', '(', 'b', ',', ' ', '"', 'a', 'b', 'c', '"', ',', ' ', '2', ')', '\x00'})), i, int32(3))
			return int32(0)
		}()
	}
	if !(func() (_cgo_ret int32) {
		_cgo_addr := &i
		*_cgo_addr = int32(*(*int8)(unsafe.Pointer(uintptr(unsafe.Pointer((*int8)(unsafe.Pointer(&b)))) + uintptr(int32(0)))))
		return *_cgo_addr
	}() == 'a') {
		func() int32 {
			common.T_printf((*int8)(unsafe.Pointer(&[103]int8{'/', 'U', 's', 'e', 'r', 's', '/', 'v', 'f', 'c', '/', 'g', 'o', 'p', 'l', 'u', 's', '/', 'l', 'i', 'b', 'c', '/', 't', 'e', 's', 't', '/', 's', 'r', 'c', '/', 'f', 'u', 'n', 'c', 't', 'i', 'o', 'n', 'a', 'l', '/', 's', 't', 'r', 'i', 'n', 'g', '.', 'c', ':', '8', '9', ':', ' ', '%', 's', ' ', 'f', 'a', 'i', 'l', 'e', 'd', ' ', '(', 's', 't', 'r', 'l', 'c', 'p', 'y', ' ', 'd', 'i', 'd', ' ', 'n', 'o', 't', ' ', 'c', 'o', 'p', 'y', ' ', 'c', 'h', 'a', 'r', 'a', 'c', 't', 'e', 'r', ' ', '%', 'd', ')', '\n', '\x00'})), (*int8)(unsafe.Pointer(&[5]int8{'b', '[', '0', ']', '\x00'})), i, 'a')
			return int32(0)
		}()
	}
	if !(func() (_cgo_ret int32) {
		_cgo_addr := &i
		*_cgo_addr = int32(*(*int8)(unsafe.Pointer(uintptr(unsafe.Pointer((*int8)(unsafe.Pointer(&b)))) + uintptr(int32(1)))))
		return *_cgo_addr
	}() == int32(0)) {
		func() int32 {
			common.T_printf((*int8)(unsafe.Pointer(&[117]int8{'/', 'U', 's', 'e', 'r', 's', '/', 'v', 'f', 'c', '/', 'g', 'o', 'p', 'l', 'u', 's', '/', 'l', 'i', 'b', 'c', '/', 't', 'e', 's', 't', '/', 's', 'r', 'c', '/', 'f', 'u', 'n', 'c', 't', 'i', 'o', 'n', 'a', 'l', '/', 's', 't', 'r', 'i', 'n', 'g', '.', 'c', ':', '9', '0', ':', ' ', '%', 's', ' ', 'f', 'a', 'i', 'l', 'e', 'd', ' ', '(', 's', 't', 'r', 'l', 'c', 'p', 'y', ' ', 'd', 'i', 'd', ' ', 'n', 'o', 't', ' ', 'n', 'u', 'l', 'l', '-', 't', 'e', 'r', 'm', 'i', 'n', 'a', 't', 'e', ' ', 'l', 'o', 'n', 'g', ' ', 's', 't', 'r', 'i', 'n', 'g', ' ', '(', '%', 'd', ')', ')', '\n', '\x00'})), (*int8)(unsafe.Pointer(&[5]int8{'b', '[', '1', ']', '\x00'})), i, int32(0))
			return int32(0)
		}()
	}
	libc.Memset(unsafe.Pointer((*int8)(unsafe.Pointer(&b))), 'x', 32)
	if !(func() (_cgo_ret int32) {
		_cgo_addr := &i
		*_cgo_addr = int32(libc.Strlcpy((*int8)(unsafe.Pointer(&b)), (*int8)(unsafe.Pointer(&[4]int8{'a', 'b', 'c', '\x00'})), uint64(3)))
		return *_cgo_addr
	}() == int32(3)) {
		func() int32 {
			common.T_printf((*int8)(unsafe.Pointer(&[85]int8{'/', 'U', 's', 'e', 'r', 's', '/', 'v', 'f', 'c', '/', 'g', 'o', 'p', 'l', 'u', 's', '/', 'l', 'i', 'b', 'c', '/', 't', 'e', 's', 't', '/', 's', 'r', 'c', '/', 'f', 'u', 'n', 'c', 't', 'i', 'o', 'n', 'a', 'l', '/', 's', 't', 'r', 'i', 'n', 'g', '.', 'c', ':', '9', '3', ':', ' ', '%', 's', ' ', 'f', 'a', 'i', 'l', 'e', 'd', ' ', '(', 'l', 'e', 'n', 'g', 't', 'h', ' ', '%', 'd', ' ', '!', '=', ' ', '%', 'd', ')', '\n', '\x00'})), (*int8)(unsafe.Pointer(&[21]int8{'s', 't', 'r', 'l', 'c', 'p', 'y', '(', 'b', ',', ' ', '"', 'a', 'b', 'c', '"', ',', ' ', '3', ')', '\x00'})), i, int32(3))
			return int32(0)
		}()
	}
	if !(func() (_cgo_ret int32) {
		_cgo_addr := &i
		*_cgo_addr = int32(*(*int8)(unsafe.Pointer(uintptr(unsafe.Pointer((*int8)(unsafe.Pointer(&b)))) + uintptr(int32(2)))))
		return *_cgo_addr
	}() == int32(0)) {
		func() int32 {
			common.T_printf((*int8)(unsafe.Pointer(&[121]int8{'/', 'U', 's', 'e', 'r', 's', '/', 'v', 'f', 'c', '/', 'g', 'o', 'p', 'l', 'u', 's', '/', 'l', 'i', 'b', 'c', '/', 't', 'e', 's', 't', '/', 's', 'r', 'c', '/', 'f', 'u', 'n', 'c', 't', 'i', 'o', 'n', 'a', 'l', '/', 's', 't', 'r', 'i', 'n', 'g', '.', 'c', ':', '9', '4', ':', ' ', '%', 's', ' ', 'f', 'a', 'i', 'l', 'e', 'd', ' ', '(', 's', 't', 'r', 'l', 'c', 'p', 'y', ' ', 'd', 'i', 'd', ' ', 'n', 'o', 't', ' ', 'n', 'u', 'l', 'l', '-', 't', 'e', 'r', 'm', 'i', 'n', 'a', 't', 'e', ' ', 'l', '-', 'l', 'e', 'n', 'g', 't', 'h', ' ', 's', 't', 'r', 'i', 'n', 'g', ' ', '(', '%', 'd', ')', ')', '\n', '\x00'})), (*int8)(unsafe.Pointer(&[5]int8{'b', '[', '2', ']', '\x00'})), i, int32(0))
			return int32(0)
		}()
	}
	if !(func() (_cgo_ret int32) {
		_cgo_addr := &i
		*_cgo_addr = int32(libc.Strlcpy(nil, (*int8)(unsafe.Pointer(&[4]int8{'a', 'b', 'c', '\x00'})), uint64(0)))
		return *_cgo_addr
	}() == int32(3)) {
		func() int32 {
			common.T_printf((*int8)(unsafe.Pointer(&[85]int8{'/', 'U', 's', 'e', 'r', 's', '/', 'v', 'f', 'c', '/', 'g', 'o', 'p', 'l', 'u', 's', '/', 'l', 'i', 'b', 'c', '/', 't', 'e', 's', 't', '/', 's', 'r', 'c', '/', 'f', 'u', 'n', 'c', 't', 'i', 'o', 'n', 'a', 'l', '/', 's', 't', 'r', 'i', 'n', 'g', '.', 'c', ':', '9', '6', ':', ' ', '%', 's', ' ', 'f', 'a', 'i', 'l', 'e', 'd', ' ', '(', 'l', 'e', 'n', 'g', 't', 'h', ' ', '%', 'd', ' ', '!', '=', ' ', '%', 'd', ')', '\n', '\x00'})), (*int8)(unsafe.Pointer(&[24]int8{'s', 't', 'r', 'l', 'c', 'p', 'y', '(', 'N', 'U', 'L', 'L', ',', ' ', '"', 'a', 'b', 'c', '"', ',', ' ', '0', ')', '\x00'})), i, int32(3))
			return int32(0)
		}()
	}
	libc.Memcpy(unsafe.Pointer((*int8)(unsafe.Pointer(&b))), unsafe.Pointer((*int8)(unsafe.Pointer(&[9]int8{'a', 'b', 'c', '\x00', '\x00', '\x00', 'x', '\x00', '\x00'}))), uint64(8))
	if !(func() (_cgo_ret int32) {
		_cgo_addr := &i
		*_cgo_addr = int32(libc.Strlcat((*int8)(unsafe.Pointer(&b)), (*int8)(unsafe.Pointer(&[4]int8{'1', '2', '3', '\x00'})), 32))
		return *_cgo_addr
	}() == int32(6)) {
		func() int32 {
			common.T_printf((*int8)(unsafe.Pointer(&[85]int8{'/', 'U', 's', 'e', 'r', 's', '/', 'v', 'f', 'c', '/', 'g', 'o', 'p', 'l', 'u', 's', '/', 'l', 'i', 'b', 'c', '/', 't', 'e', 's', 't', '/', 's', 'r', 'c', '/', 'f', 'u', 'n', 'c', 't', 'i', 'o', 'n', 'a', 'l', '/', 's', 't', 'r', 'i', 'n', 'g', '.', 'c', ':', '9', '9', ':', ' ', '%', 's', ' ', 'f', 'a', 'i', 'l', 'e', 'd', ' ', '(', 'l', 'e', 'n', 'g', 't', 'h', ' ', '%', 'd', ' ', '!', '=', ' ', '%', 'd', ')', '\n', '\x00'})), (*int8)(unsafe.Pointer(&[28]int8{'s', 't', 'r', 'l', 'c', 'a', 't', '(', 'b', ',', ' ', '"', '1', '2', '3', '"', ',', ' ', 's', 'i', 'z', 'e', 'o', 'f', ' ', 'b', ')', '\x00'})), i, int32(6))
			return int32(0)
		}()
	}
	if !!(libc.Strcmp((*int8)(unsafe.Pointer(&b)), (*int8)(unsafe.Pointer(&[7]int8{'a', 'b', 'c', '1', '2', '3', '\x00'}))) != 0) {
		func() int32 {
			common.T_printf((*int8)(unsafe.Pointer(&[76]int8{'/', 'U', 's', 'e', 'r', 's', '/', 'v', 'f', 'c', '/', 'g', 'o', 'p', 'l', 'u', 's', '/', 'l', 'i', 'b', 'c', '/', 't', 'e', 's', 't', '/', 's', 'r', 'c', '/', 'f', 'u', 'n', 'c', 't', 'i', 'o', 'n', 'a', 'l', '/', 's', 't', 'r', 'i', 'n', 'g', '.', 'c', ':', '1', '0', '0', ':', ' ', '[', '%', 's', ']', ' ', '!', '=', ' ', '[', '%', 's', ']', ' ', '(', '%', 's', ')', '\n', '\x00'})), (*int8)(unsafe.Pointer(&b)), (*int8)(unsafe.Pointer(&[7]int8{'a', 'b', 'c', '1', '2', '3', '\x00'})), (*int8)(unsafe.Pointer(&[15]int8{'s', 't', 'r', 'l', 'c', 'a', 't', ' ', 'r', 'e', 's', 'u', 'l', 't', '\x00'})))
			return int32(0)
		}()
	}
	libc.Memcpy(unsafe.Pointer((*int8)(unsafe.Pointer(&b))), unsafe.Pointer((*int8)(unsafe.Pointer(&[9]int8{'a', 'b', 'c', '\x00', '\x00', '\x00', 'x', '\x00', '\x00'}))), uint64(8))
	if !(func() (_cgo_ret int32) {
		_cgo_addr := &i
		*_cgo_addr = int32(libc.Strlcat((*int8)(unsafe.Pointer(&b)), (*int8)(unsafe.Pointer(&[4]int8{'1', '2', '3', '\x00'})), uint64(6)))
		return *_cgo_addr
	}() == int32(6)) {
		func() int32 {
			common.T_printf((*int8)(unsafe.Pointer(&[86]int8{'/', 'U', 's', 'e', 'r', 's', '/', 'v', 'f', 'c', '/', 'g', 'o', 'p', 'l', 'u', 's', '/', 'l', 'i', 'b', 'c', '/', 't', 'e', 's', 't', '/', 's', 'r', 'c', '/', 'f', 'u', 'n', 'c', 't', 'i', 'o', 'n', 'a', 'l', '/', 's', 't', 'r', 'i', 'n', 'g', '.', 'c', ':', '1', '0', '3', ':', ' ', '%', 's', ' ', 'f', 'a', 'i', 'l', 'e', 'd', ' ', '(', 'l', 'e', 'n', 'g', 't', 'h', ' ', '%', 'd', ' ', '!', '=', ' ', '%', 'd', ')', '\n', '\x00'})), (*int8)(unsafe.Pointer(&[21]int8{'s', 't', 'r', 'l', 'c', 'a', 't', '(', 'b', ',', ' ', '"', '1', '2', '3', '"', ',', ' ', '6', ')', '\x00'})), i, int32(6))
			return int32(0)
		}()
	}
	if !!(libc.Strcmp((*int8)(unsafe.Pointer(&b)), (*int8)(unsafe.Pointer(&[6]int8{'a', 'b', 'c', '1', '2', '\x00'}))) != 0) {
		func() int32 {
			common.T_printf((*int8)(unsafe.Pointer(&[76]int8{'/', 'U', 's', 'e', 'r', 's', '/', 'v', 'f', 'c', '/', 'g', 'o', 'p', 'l', 'u', 's', '/', 'l', 'i', 'b', 'c', '/', 't', 'e', 's', 't', '/', 's', 'r', 'c', '/', 'f', 'u', 'n', 'c', 't', 'i', 'o', 'n', 'a', 'l', '/', 's', 't', 'r', 'i', 'n', 'g', '.', 'c', ':', '1', '0', '4', ':', ' ', '[', '%', 's', ']', ' ', '!', '=', ' ', '[', '%', 's', ']', ' ', '(', '%', 's', ')', '\n', '\x00'})), (*int8)(unsafe.Pointer(&b)), (*int8)(unsafe.Pointer(&[6]int8{'a', 'b', 'c', '1', '2', '\x00'})), (*int8)(unsafe.Pointer(&[15]int8{'s', 't', 'r', 'l', 'c', 'a', 't', ' ', 'r', 'e', 's', 'u', 'l', 't', '\x00'})))
			return int32(0)
		}()
	}
	if !(func() (_cgo_ret int32) {
		_cgo_addr := &i
		*_cgo_addr = int32(*(*int8)(unsafe.Pointer(uintptr(unsafe.Pointer((*int8)(unsafe.Pointer(&b)))) + uintptr(int32(6)))))
		return *_cgo_addr
	}() == 'x') {
		func() int32 {
			common.T_printf((*int8)(unsafe.Pointer(&[105]int8{'/', 'U', 's', 'e', 'r', 's', '/', 'v', 'f', 'c', '/', 'g', 'o', 'p', 'l', 'u', 's', '/', 'l', 'i', 'b', 'c', '/', 't', 'e', 's', 't', '/', 's', 'r', 'c', '/', 'f', 'u', 'n', 'c', 't', 'i', 'o', 'n', 'a', 'l', '/', 's', 't', 'r', 'i', 'n', 'g', '.', 'c', ':', '1', '0', '5', ':', ' ', '%', 's', ' ', 'f', 'a', 'i', 'l', 'e', 'd', ' ', '(', 's', 't', 'r', 'l', 'c', 'a', 't', ' ', 'w', 'r', 'o', 't', 'e', ' ', 'p', 'a', 's', 't', ' ', 's', 't', 'r', 'i', 'n', 'g', ' ', '%', 'd', ' ', '!', '=', ' ', '%', 'd', ')', '\n', '\x00'})), (*int8)(unsafe.Pointer(&[5]int8{'b', '[', '6', ']', '\x00'})), i, 'x')
			return int32(0)
		}()
	}
	libc.Memcpy(unsafe.Pointer((*int8)(unsafe.Pointer(&b))), unsafe.Pointer((*int8)(unsafe.Pointer(&[9]int8{'a', 'b', 'c', '\x00', '\x00', '\x00', 'x', '\x00', '\x00'}))), uint64(8))
	if !(func() (_cgo_ret int32) {
		_cgo_addr := &i
		*_cgo_addr = int32(libc.Strlcat((*int8)(unsafe.Pointer(&b)), (*int8)(unsafe.Pointer(&[4]int8{'1', '2', '3', '\x00'})), uint64(4)))
		return *_cgo_addr
	}() == int32(6)) {
		func() int32 {
			common.T_printf((*int8)(unsafe.Pointer(&[86]int8{'/', 'U', 's', 'e', 'r', 's', '/', 'v', 'f', 'c', '/', 'g', 'o', 'p', 'l', 'u', 's', '/', 'l', 'i', 'b', 'c', '/', 't', 'e', 's', 't', '/', 's', 'r', 'c', '/', 'f', 'u', 'n', 'c', 't', 'i', 'o', 'n', 'a', 'l', '/', 's', 't', 'r', 'i', 'n', 'g', '.', 'c', ':', '1', '0', '8', ':', ' ', '%', 's', ' ', 'f', 'a', 'i', 'l', 'e', 'd', ' ', '(', 'l', 'e', 'n', 'g', 't', 'h', ' ', '%', 'd', ' ', '!', '=', ' ', '%', 'd', ')', '\n', '\x00'})), (*int8)(unsafe.Pointer(&[21]int8{'s', 't', 'r', 'l', 'c', 'a', 't', '(', 'b', ',', ' ', '"', '1', '2', '3', '"', ',', ' ', '4', ')', '\x00'})), i, int32(6))
			return int32(0)
		}()
	}
	if !!(libc.Strcmp((*int8)(unsafe.Pointer(&b)), (*int8)(unsafe.Pointer(&[4]int8{'a', 'b', 'c', '\x00'}))) != 0) {
		func() int32 {
			common.T_printf((*int8)(unsafe.Pointer(&[76]int8{'/', 'U', 's', 'e', 'r', 's', '/', 'v', 'f', 'c', '/', 'g', 'o', 'p', 'l', 'u', 's', '/', 'l', 'i', 'b', 'c', '/', 't', 'e', 's', 't', '/', 's', 'r', 'c', '/', 'f', 'u', 'n', 'c', 't', 'i', 'o', 'n', 'a', 'l', '/', 's', 't', 'r', 'i', 'n', 'g', '.', 'c', ':', '1', '0', '9', ':', ' ', '[', '%', 's', ']', ' ', '!', '=', ' ', '[', '%', 's', ']', ' ', '(', '%', 's', ')', '\n', '\x00'})), (*int8)(unsafe.Pointer(&b)), (*int8)(unsafe.Pointer(&[4]int8{'a', 'b', 'c', '\x00'})), (*int8)(unsafe.Pointer(&[15]int8{'s', 't', 'r', 'l', 'c', 'a', 't', ' ', 'r', 'e', 's', 'u', 'l', 't', '\x00'})))
			return int32(0)
		}()
	}
	libc.Memcpy(unsafe.Pointer((*int8)(unsafe.Pointer(&b))), unsafe.Pointer((*int8)(unsafe.Pointer(&[9]int8{'a', 'b', 'c', '\x00', '\x00', '\x00', 'x', '\x00', '\x00'}))), uint64(8))
	if !(func() (_cgo_ret int32) {
		_cgo_addr := &i
		*_cgo_addr = int32(libc.Strlcat((*int8)(unsafe.Pointer(&b)), (*int8)(unsafe.Pointer(&[4]int8{'1', '2', '3', '\x00'})), uint64(3)))
		return *_cgo_addr
	}() == int32(6)) {
		func() int32 {
			common.T_printf((*int8)(unsafe.Pointer(&[86]int8{'/', 'U', 's', 'e', 'r', 's', '/', 'v', 'f', 'c', '/', 'g', 'o', 'p', 'l', 'u', 's', '/', 'l', 'i', 'b', 'c', '/', 't', 'e', 's', 't', '/', 's', 'r', 'c', '/', 'f', 'u', 'n', 'c', 't', 'i', 'o', 'n', 'a', 'l', '/', 's', 't', 'r', 'i', 'n', 'g', '.', 'c', ':', '1', '1', '2', ':', ' ', '%', 's', ' ', 'f', 'a', 'i', 'l', 'e', 'd', ' ', '(', 'l', 'e', 'n', 'g', 't', 'h', ' ', '%', 'd', ' ', '!', '=', ' ', '%', 'd', ')', '\n', '\x00'})), (*int8)(unsafe.Pointer(&[21]int8{'s', 't', 'r', 'l', 'c', 'a', 't', '(', 'b', ',', ' ', '"', '1', '2', '3', '"', ',', ' ', '3', ')', '\x00'})), i, int32(6))
			return int32(0)
		}()
	}
	if !!(libc.Strcmp((*int8)(unsafe.Pointer(&b)), (*int8)(unsafe.Pointer(&[4]int8{'a', 'b', 'c', '\x00'}))) != 0) {
		func() int32 {
			common.T_printf((*int8)(unsafe.Pointer(&[76]int8{'/', 'U', 's', 'e', 'r', 's', '/', 'v', 'f', 'c', '/', 'g', 'o', 'p', 'l', 'u', 's', '/', 'l', 'i', 'b', 'c', '/', 't', 'e', 's', 't', '/', 's', 'r', 'c', '/', 'f', 'u', 'n', 'c', 't', 'i', 'o', 'n', 'a', 'l', '/', 's', 't', 'r', 'i', 'n', 'g', '.', 'c', ':', '1', '1', '3', ':', ' ', '[', '%', 's', ']', ' ', '!', '=', ' ', '[', '%', 's', ']', ' ', '(', '%', 's', ')', '\n', '\x00'})), (*int8)(unsafe.Pointer(&b)), (*int8)(unsafe.Pointer(&[4]int8{'a', 'b', 'c', '\x00'})), (*int8)(unsafe.Pointer(&[15]int8{'s', 't', 'r', 'l', 'c', 'a', 't', ' ', 'r', 'e', 's', 'u', 'l', 't', '\x00'})))
			return int32(0)
		}()
	}
	return common.T_status
}
func TestMain(t *testing.T) {
	if _cgo_ret := _cgo_main(); _cgo_ret != 0 {
		t.Fatal("exit status", _cgo_ret)
	}
}
