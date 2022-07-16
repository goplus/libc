package string_memcpy

import (
	unsafe "unsafe"
	libc "github.com/goplus/libc"
	common "github.com/goplus/libc/test/common"
	testing "testing"
)

var buf_cgo1_string_memcpy [512]int8
var pmemcpy_cgo2_string_memcpy func(unsafe.Pointer, unsafe.Pointer, uint64) unsafe.Pointer

func aligned_cgo3_string_memcpy(p unsafe.Pointer) unsafe.Pointer {
	return unsafe.Pointer(uintptr((uint64(uintptr(p)) + uint64(63)) & uint64(18446744073709551552)))
}
func test_align_cgo4_string_memcpy(dalign int32, salign int32, len int32) {
	var src *int8 = (*int8)(aligned_cgo3_string_memcpy(unsafe.Pointer((*int8)(unsafe.Pointer(&buf_cgo1_string_memcpy)))))
	var dst *int8 = (*int8)(aligned_cgo3_string_memcpy(unsafe.Pointer((*int8)(unsafe.Pointer(uintptr(unsafe.Pointer((*int8)(unsafe.Pointer(&buf_cgo1_string_memcpy)))) + uintptr(int32(128)))))))
	var want *int8 = (*int8)(aligned_cgo3_string_memcpy(unsafe.Pointer((*int8)(unsafe.Pointer(uintptr(unsafe.Pointer((*int8)(unsafe.Pointer(&buf_cgo1_string_memcpy)))) + uintptr(int32(256)))))))
	var p *int8
	var i int32
	if salign+len > int32(80) || dalign+len > int32(80) {
		libc.Abort()
	}
	for i = int32(0); i < int32(80); i++ {
		*(*int8)(unsafe.Pointer(uintptr(unsafe.Pointer(src)) + uintptr(i))) = int8('#')
		*(*int8)(unsafe.Pointer(uintptr(unsafe.Pointer(dst)) + uintptr(i))) = func() (_cgo_ret int8) {
			_cgo_addr := &*(*int8)(unsafe.Pointer(uintptr(unsafe.Pointer(want)) + uintptr(i)))
			*_cgo_addr = int8(' ')
			return *_cgo_addr
		}()
	}
	for i = int32(0); i < len; i++ {
		*(*int8)(unsafe.Pointer(uintptr(unsafe.Pointer(src)) + uintptr(salign+i))) = func() (_cgo_ret int8) {
			_cgo_addr := &*(*int8)(unsafe.Pointer(uintptr(unsafe.Pointer(want)) + uintptr(dalign+i)))
			*_cgo_addr = int8('0' + i)
			return *_cgo_addr
		}()
	}
	p = (*int8)(pmemcpy_cgo2_string_memcpy(unsafe.Pointer((*int8)(unsafe.Pointer(uintptr(unsafe.Pointer(dst))+uintptr(dalign)))), unsafe.Pointer((*int8)(unsafe.Pointer(uintptr(unsafe.Pointer(src))+uintptr(salign)))), uint64(len)))
	if uintptr(unsafe.Pointer(p)) != uintptr(unsafe.Pointer((*int8)(unsafe.Pointer(uintptr(unsafe.Pointer(dst))+uintptr(dalign))))) {
		common.T_printf((*int8)(unsafe.Pointer(&[63]int8{'s', 'r', 'c', '/', 'f', 'u', 'n', 'c', 't', 'i', 'o', 'n', 'a', 'l', '/', 's', 't', 'r', 'i', 'n', 'g', '_', 'm', 'e', 'm', 'c', 'p', 'y', '.', 'c', ':', '3', '3', ':', ' ', 'm', 'e', 'm', 'c', 'p', 'y', '(', '%', 'p', ',', '.', '.', '.', ')', ' ', 'r', 'e', 't', 'u', 'r', 'n', 'e', 'd', ' ', '%', 'p', '\n', '\x00'})), (*int8)(unsafe.Pointer(uintptr(unsafe.Pointer(dst))+uintptr(dalign))), p)
	}
	for i = int32(0); i < int32(80); i++ {
		if int32(*(*int8)(unsafe.Pointer(uintptr(unsafe.Pointer(dst)) + uintptr(i)))) != int32(*(*int8)(unsafe.Pointer(uintptr(unsafe.Pointer(want)) + uintptr(i)))) {
			common.T_printf((*int8)(unsafe.Pointer(&[74]int8{'s', 'r', 'c', '/', 'f', 'u', 'n', 'c', 't', 'i', 'o', 'n', 'a', 'l', '/', 's', 't', 'r', 'i', 'n', 'g', '_', 'm', 'e', 'm', 'c', 'p', 'y', '.', 'c', ':', '3', '6', ':', ' ', 'm', 'e', 'm', 'c', 'p', 'y', '(', 'a', 'l', 'i', 'g', 'n', ' ', '%', 'd', ',', ' ', 'a', 'l', 'i', 'g', 'n', ' ', '%', 'd', ',', ' ', '%', 'd', ')', ' ', 'f', 'a', 'i', 'l', 'e', 'd', '\n', '\x00'})), dalign, salign, len)
			common.T_printf((*int8)(unsafe.Pointer(&[12]int8{'g', 'o', 't', ' ', ':', ' ', '%', '.', '*', 's', '\n', '\x00'})), dalign+len+int32(1), dst)
			common.T_printf((*int8)(unsafe.Pointer(&[12]int8{'w', 'a', 'n', 't', ':', ' ', '%', '.', '*', 's', '\n', '\x00'})), dalign+len+int32(1), want)
			break
		}
	}
}
func _cgo_main() int32 {
	var i int32
	var j int32
	var k int32
	pmemcpy_cgo2_string_memcpy = libc.Memcpy
	for i = int32(0); i < int32(16); i++ {
		for j = int32(0); j < int32(16); j++ {
			for k = int32(0); k < int32(64); k++ {
				test_align_cgo4_string_memcpy(i, j, k)
			}
		}
	}
	return common.T_status
}
func TestMain(t *testing.T) {
	if _cgo_ret := _cgo_main(); _cgo_ret != 0 {
		t.Fatal("exit status", _cgo_ret)
	}
}
