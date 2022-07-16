package string_memset

import (
	unsafe "unsafe"
	libc "github.com/goplus/libc"
	common "github.com/goplus/libc/test/common"
	testing "testing"
)

var buf_cgo1_string_memset [400]int8
var buf2_cgo2_string_memset [400]int8
var pmemset_cgo3_string_memset func(unsafe.Pointer, int32, uint64) unsafe.Pointer

func aligned_cgo4_string_memset(p unsafe.Pointer) *int8 {
	return (*int8)(unsafe.Pointer(uintptr((uint64(uintptr(p)) + uint64(63)) & uint64(18446744073709551552))))
}
func test_align_cgo5_string_memset(align int32, len int32) {
	var s *int8 = (*int8)(unsafe.Pointer(uintptr(unsafe.Pointer(aligned_cgo4_string_memset(unsafe.Pointer((*int8)(unsafe.Pointer(uintptr(unsafe.Pointer((*int8)(unsafe.Pointer(&buf_cgo1_string_memset))))+uintptr(int32(64)))))))) + uintptr(align)))
	var want *int8 = (*int8)(unsafe.Pointer(uintptr(unsafe.Pointer(aligned_cgo4_string_memset(unsafe.Pointer((*int8)(unsafe.Pointer(uintptr(unsafe.Pointer((*int8)(unsafe.Pointer(&buf2_cgo2_string_memset))))+uintptr(int32(64)))))))) + uintptr(align)))
	var p *int8
	var i int32
	if int64(len+int32(64)) > int64(uintptr(unsafe.Pointer((*int8)(unsafe.Pointer(uintptr(unsafe.Pointer((*int8)(unsafe.Pointer(&buf_cgo1_string_memset))))+uintptr(int32(400))))))-uintptr(unsafe.Pointer(s))) || int64(len+int32(64)) > int64(uintptr(unsafe.Pointer((*int8)(unsafe.Pointer(uintptr(unsafe.Pointer((*int8)(unsafe.Pointer(&buf2_cgo2_string_memset))))+uintptr(int32(400))))))-uintptr(unsafe.Pointer(want))) {
		libc.Abort()
	}
	for i = int32(0); i < int32(400); i++ {
		*(*int8)(unsafe.Pointer(uintptr(unsafe.Pointer((*int8)(unsafe.Pointer(&buf_cgo1_string_memset)))) + uintptr(i))) = func() (_cgo_ret int8) {
			_cgo_addr := &*(*int8)(unsafe.Pointer(uintptr(unsafe.Pointer((*int8)(unsafe.Pointer(&buf2_cgo2_string_memset)))) + uintptr(i)))
			*_cgo_addr = int8(' ')
			return *_cgo_addr
		}()
	}
	for i = int32(0); i < len; i++ {
		*(*int8)(unsafe.Pointer(uintptr(unsafe.Pointer(want)) + uintptr(i))) = int8('#')
	}
	p = (*int8)(pmemset_cgo3_string_memset(unsafe.Pointer(s), '#', uint64(len)))
	if uintptr(unsafe.Pointer(p)) != uintptr(unsafe.Pointer(s)) {
		common.T_printf((*int8)(unsafe.Pointer(&[63]int8{'s', 'r', 'c', '/', 'f', 'u', 'n', 'c', 't', 'i', 'o', 'n', 'a', 'l', '/', 's', 't', 'r', 'i', 'n', 'g', '_', 'm', 'e', 'm', 's', 'e', 't', '.', 'c', ':', '3', '2', ':', ' ', 'm', 'e', 'm', 's', 'e', 't', '(', '%', 'p', ',', '.', '.', '.', ')', ' ', 'r', 'e', 't', 'u', 'r', 'n', 'e', 'd', ' ', '%', 'p', '\n', '\x00'})), s, p)
	}
	for i = -64; i < len+int32(64); i++ {
		if int32(*(*int8)(unsafe.Pointer(uintptr(unsafe.Pointer(s)) + uintptr(i)))) != int32(*(*int8)(unsafe.Pointer(uintptr(unsafe.Pointer(want)) + uintptr(i)))) {
			common.T_printf((*int8)(unsafe.Pointer(&[79]int8{'s', 'r', 'c', '/', 'f', 'u', 'n', 'c', 't', 'i', 'o', 'n', 'a', 'l', '/', 's', 't', 'r', 'i', 'n', 'g', '_', 'm', 'e', 'm', 's', 'e', 't', '.', 'c', ':', '3', '5', ':', ' ', 'm', 'e', 'm', 's', 'e', 't', '(', 'a', 'l', 'i', 'g', 'n', ' ', '%', 'd', ',', ' ', '\'', '#', '\'', ',', ' ', '%', 'd', ')', ' ', 'f', 'a', 'i', 'l', 'e', 'd', ' ', 'a', 't', ' ', 'p', 'o', 's', ' ', '%', 'd', '\n', '\x00'})), align, len, i)
			common.T_printf((*int8)(unsafe.Pointer(&[14]int8{'g', 'o', 't', ' ', ':', ' ', '\'', '%', '.', '*', 's', '\'', '\n', '\x00'})), len+int32(128), (*int8)(unsafe.Pointer(uintptr(unsafe.Pointer(s))-uintptr(int32(64)))))
			common.T_printf((*int8)(unsafe.Pointer(&[14]int8{'w', 'a', 'n', 't', ':', ' ', '\'', '%', '.', '*', 's', '\'', '\n', '\x00'})), len+int32(128), (*int8)(unsafe.Pointer(uintptr(unsafe.Pointer(want))-uintptr(int32(64)))))
			break
		}
	}
}
func test_value_cgo6_string_memset(c int32) {
	var i int32
	pmemset_cgo3_string_memset(unsafe.Pointer((*int8)(unsafe.Pointer(&buf_cgo1_string_memset))), c, uint64(10))
	for i = int32(0); i < int32(10); i++ {
		if int32(uint8(*(*int8)(unsafe.Pointer(uintptr(unsafe.Pointer((*int8)(unsafe.Pointer(&buf_cgo1_string_memset)))) + uintptr(i))))) != int32(uint8(c)) {
			common.T_printf((*int8)(unsafe.Pointer(&[62]int8{'s', 'r', 'c', '/', 'f', 'u', 'n', 'c', 't', 'i', 'o', 'n', 'a', 'l', '/', 's', 't', 'r', 'i', 'n', 'g', '_', 'm', 'e', 'm', 's', 'e', 't', '.', 'c', ':', '4', '9', ':', ' ', 'm', 'e', 'm', 's', 'e', 't', '(', '%', 'd', ')', ' ', 'f', 'a', 'i', 'l', 'e', 'd', ':', ' ', 'g', 'o', 't', ' ', '%', 'd', '\n', '\x00'})), c, int32(*(*int8)(unsafe.Pointer(uintptr(unsafe.Pointer((*int8)(unsafe.Pointer(&buf_cgo1_string_memset)))) + uintptr(i)))))
			break
		}
	}
}
func _cgo_main() int32 {
	var i int32
	var j int32
	pmemset_cgo3_string_memset = libc.Memset
	for i = int32(0); i < int32(16); i++ {
		for j = int32(0); j < int32(200); j++ {
			test_align_cgo5_string_memset(i, j)
		}
	}
	test_value_cgo6_string_memset('c')
	test_value_cgo6_string_memset(int32(0))
	test_value_cgo6_string_memset(-1)
	test_value_cgo6_string_memset(-5)
	test_value_cgo6_string_memset(int32(171))
	return common.T_status
}
func TestMain(t *testing.T) {
	if _cgo_ret := _cgo_main(); _cgo_ret != 0 {
		t.Fatal("exit status", _cgo_ret)
	}
}
