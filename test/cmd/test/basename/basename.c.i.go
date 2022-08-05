package basename

import (
	libc "github.com/goplus/libc"
	unsafe "unsafe"
	common "github.com/goplus/libc/test/common"
	testing "testing"
)

func _cgo_main() int32 {
	if libc.Strcmp(libc.Basename(nil), (*int8)(unsafe.Pointer(&[2]int8{'.', '\x00'}))) != int32(0) {
		common.T_printf((*int8)(unsafe.Pointer(&[102]int8{'/', 'U', 's', 'e', 'r', 's', '/', 'x', 'u', 's', 'h', 'i', 'w', 'e', 'i', '/', 'w', 'o', 'r', 'k', '/', 'l', 'i', 'b', 'c', '/', 't', 'e', 's', 't', '/', 's', 'r', 'c', '/', 'f', 'u', 'n', 'c', 't', 'i', 'o', 'n', 'a', 'l', '/', 'b', 'a', 's', 'e', 'n', 'a', 'm', 'e', '.', 'c', ':', '1', '7', ':', ' ', 'b', 'a', 's', 'e', 'n', 'a', 'm', 'e', '(', '0', ')', ' ', 'r', 'e', 't', 'u', 'r', 'n', 'e', 'd', ' ', '"', '%', 's', '"', ';', ' ', 'e', 'x', 'p', 'e', 'c', 't', 'e', 'd', ' ', '"', '.', '"', '\n', '\x00'})), libc.Basename(nil))
	}
	{
		var tmp [100]int8
		var got *int8 = libc.Basename(libc.Strcpy((*int8)(unsafe.Pointer(&tmp)), (*int8)(unsafe.Pointer(&[1]int8{'\x00'}))))
		if libc.Strcmp((*int8)(unsafe.Pointer(&[2]int8{'.', '\x00'})), got) != int32(0) {
			common.T_printf((*int8)(unsafe.Pointer(&[96]int8{'/', 'U', 's', 'e', 'r', 's', '/', 'x', 'u', 's', 'h', 'i', 'w', 'e', 'i', '/', 'w', 'o', 'r', 'k', '/', 'l', 'i', 'b', 'c', '/', 't', 'e', 's', 't', '/', 's', 'r', 'c', '/', 'f', 'u', 'n', 'c', 't', 'i', 'o', 'n', 'a', 'l', '/', 'b', 'a', 's', 'e', 'n', 'a', 'm', 'e', '.', 'c', ':', '1', '8', ':', ' ', 'b', 'a', 's', 'e', 'n', 'a', 'm', 'e', '(', '"', '%', 's', '"', ')', ' ', 'g', 'o', 't', ' ', '"', '%', 's', '"', ' ', 'w', 'a', 'n', 't', ' ', '"', '%', 's', '"', '\n', '\x00'})), (*int8)(unsafe.Pointer(&[1]int8{'\x00'})), got, (*int8)(unsafe.Pointer(&[2]int8{'.', '\x00'})))
		}
	}
	{
		var tmp [100]int8
		var got *int8 = libc.Basename(libc.Strcpy((*int8)(unsafe.Pointer(&tmp)), (*int8)(unsafe.Pointer(&[9]int8{'/', 'u', 's', 'r', '/', 'l', 'i', 'b', '\x00'}))))
		if libc.Strcmp((*int8)(unsafe.Pointer(&[4]int8{'l', 'i', 'b', '\x00'})), got) != int32(0) {
			common.T_printf((*int8)(unsafe.Pointer(&[96]int8{'/', 'U', 's', 'e', 'r', 's', '/', 'x', 'u', 's', 'h', 'i', 'w', 'e', 'i', '/', 'w', 'o', 'r', 'k', '/', 'l', 'i', 'b', 'c', '/', 't', 'e', 's', 't', '/', 's', 'r', 'c', '/', 'f', 'u', 'n', 'c', 't', 'i', 'o', 'n', 'a', 'l', '/', 'b', 'a', 's', 'e', 'n', 'a', 'm', 'e', '.', 'c', ':', '1', '9', ':', ' ', 'b', 'a', 's', 'e', 'n', 'a', 'm', 'e', '(', '"', '%', 's', '"', ')', ' ', 'g', 'o', 't', ' ', '"', '%', 's', '"', ' ', 'w', 'a', 'n', 't', ' ', '"', '%', 's', '"', '\n', '\x00'})), (*int8)(unsafe.Pointer(&[9]int8{'/', 'u', 's', 'r', '/', 'l', 'i', 'b', '\x00'})), got, (*int8)(unsafe.Pointer(&[4]int8{'l', 'i', 'b', '\x00'})))
		}
	}
	{
		var tmp [100]int8
		var got *int8 = libc.Basename(libc.Strcpy((*int8)(unsafe.Pointer(&tmp)), (*int8)(unsafe.Pointer(&[6]int8{'/', 'u', 's', 'r', '/', '\x00'}))))
		if libc.Strcmp((*int8)(unsafe.Pointer(&[4]int8{'u', 's', 'r', '\x00'})), got) != int32(0) {
			common.T_printf((*int8)(unsafe.Pointer(&[96]int8{'/', 'U', 's', 'e', 'r', 's', '/', 'x', 'u', 's', 'h', 'i', 'w', 'e', 'i', '/', 'w', 'o', 'r', 'k', '/', 'l', 'i', 'b', 'c', '/', 't', 'e', 's', 't', '/', 's', 'r', 'c', '/', 'f', 'u', 'n', 'c', 't', 'i', 'o', 'n', 'a', 'l', '/', 'b', 'a', 's', 'e', 'n', 'a', 'm', 'e', '.', 'c', ':', '2', '0', ':', ' ', 'b', 'a', 's', 'e', 'n', 'a', 'm', 'e', '(', '"', '%', 's', '"', ')', ' ', 'g', 'o', 't', ' ', '"', '%', 's', '"', ' ', 'w', 'a', 'n', 't', ' ', '"', '%', 's', '"', '\n', '\x00'})), (*int8)(unsafe.Pointer(&[6]int8{'/', 'u', 's', 'r', '/', '\x00'})), got, (*int8)(unsafe.Pointer(&[4]int8{'u', 's', 'r', '\x00'})))
		}
	}
	{
		var tmp [100]int8
		var got *int8 = libc.Basename(libc.Strcpy((*int8)(unsafe.Pointer(&tmp)), (*int8)(unsafe.Pointer(&[5]int8{'u', 's', 'r', '/', '\x00'}))))
		if libc.Strcmp((*int8)(unsafe.Pointer(&[4]int8{'u', 's', 'r', '\x00'})), got) != int32(0) {
			common.T_printf((*int8)(unsafe.Pointer(&[96]int8{'/', 'U', 's', 'e', 'r', 's', '/', 'x', 'u', 's', 'h', 'i', 'w', 'e', 'i', '/', 'w', 'o', 'r', 'k', '/', 'l', 'i', 'b', 'c', '/', 't', 'e', 's', 't', '/', 's', 'r', 'c', '/', 'f', 'u', 'n', 'c', 't', 'i', 'o', 'n', 'a', 'l', '/', 'b', 'a', 's', 'e', 'n', 'a', 'm', 'e', '.', 'c', ':', '2', '1', ':', ' ', 'b', 'a', 's', 'e', 'n', 'a', 'm', 'e', '(', '"', '%', 's', '"', ')', ' ', 'g', 'o', 't', ' ', '"', '%', 's', '"', ' ', 'w', 'a', 'n', 't', ' ', '"', '%', 's', '"', '\n', '\x00'})), (*int8)(unsafe.Pointer(&[5]int8{'u', 's', 'r', '/', '\x00'})), got, (*int8)(unsafe.Pointer(&[4]int8{'u', 's', 'r', '\x00'})))
		}
	}
	{
		var tmp [100]int8
		var got *int8 = libc.Basename(libc.Strcpy((*int8)(unsafe.Pointer(&tmp)), (*int8)(unsafe.Pointer(&[2]int8{'/', '\x00'}))))
		if libc.Strcmp((*int8)(unsafe.Pointer(&[2]int8{'/', '\x00'})), got) != int32(0) {
			common.T_printf((*int8)(unsafe.Pointer(&[96]int8{'/', 'U', 's', 'e', 'r', 's', '/', 'x', 'u', 's', 'h', 'i', 'w', 'e', 'i', '/', 'w', 'o', 'r', 'k', '/', 'l', 'i', 'b', 'c', '/', 't', 'e', 's', 't', '/', 's', 'r', 'c', '/', 'f', 'u', 'n', 'c', 't', 'i', 'o', 'n', 'a', 'l', '/', 'b', 'a', 's', 'e', 'n', 'a', 'm', 'e', '.', 'c', ':', '2', '2', ':', ' ', 'b', 'a', 's', 'e', 'n', 'a', 'm', 'e', '(', '"', '%', 's', '"', ')', ' ', 'g', 'o', 't', ' ', '"', '%', 's', '"', ' ', 'w', 'a', 'n', 't', ' ', '"', '%', 's', '"', '\n', '\x00'})), (*int8)(unsafe.Pointer(&[2]int8{'/', '\x00'})), got, (*int8)(unsafe.Pointer(&[2]int8{'/', '\x00'})))
		}
	}
	{
		var tmp [100]int8
		var got *int8 = libc.Basename(libc.Strcpy((*int8)(unsafe.Pointer(&tmp)), (*int8)(unsafe.Pointer(&[4]int8{'/', '/', '/', '\x00'}))))
		if libc.Strcmp((*int8)(unsafe.Pointer(&[2]int8{'/', '\x00'})), got) != int32(0) {
			common.T_printf((*int8)(unsafe.Pointer(&[96]int8{'/', 'U', 's', 'e', 'r', 's', '/', 'x', 'u', 's', 'h', 'i', 'w', 'e', 'i', '/', 'w', 'o', 'r', 'k', '/', 'l', 'i', 'b', 'c', '/', 't', 'e', 's', 't', '/', 's', 'r', 'c', '/', 'f', 'u', 'n', 'c', 't', 'i', 'o', 'n', 'a', 'l', '/', 'b', 'a', 's', 'e', 'n', 'a', 'm', 'e', '.', 'c', ':', '2', '3', ':', ' ', 'b', 'a', 's', 'e', 'n', 'a', 'm', 'e', '(', '"', '%', 's', '"', ')', ' ', 'g', 'o', 't', ' ', '"', '%', 's', '"', ' ', 'w', 'a', 'n', 't', ' ', '"', '%', 's', '"', '\n', '\x00'})), (*int8)(unsafe.Pointer(&[4]int8{'/', '/', '/', '\x00'})), got, (*int8)(unsafe.Pointer(&[2]int8{'/', '\x00'})))
		}
	}
	{
		var tmp [100]int8
		var got *int8 = libc.Basename(libc.Strcpy((*int8)(unsafe.Pointer(&tmp)), (*int8)(unsafe.Pointer(&[13]int8{'/', '/', 'u', 's', 'r', '/', '/', 'l', 'i', 'b', '/', '/', '\x00'}))))
		if libc.Strcmp((*int8)(unsafe.Pointer(&[4]int8{'l', 'i', 'b', '\x00'})), got) != int32(0) {
			common.T_printf((*int8)(unsafe.Pointer(&[96]int8{'/', 'U', 's', 'e', 'r', 's', '/', 'x', 'u', 's', 'h', 'i', 'w', 'e', 'i', '/', 'w', 'o', 'r', 'k', '/', 'l', 'i', 'b', 'c', '/', 't', 'e', 's', 't', '/', 's', 'r', 'c', '/', 'f', 'u', 'n', 'c', 't', 'i', 'o', 'n', 'a', 'l', '/', 'b', 'a', 's', 'e', 'n', 'a', 'm', 'e', '.', 'c', ':', '2', '4', ':', ' ', 'b', 'a', 's', 'e', 'n', 'a', 'm', 'e', '(', '"', '%', 's', '"', ')', ' ', 'g', 'o', 't', ' ', '"', '%', 's', '"', ' ', 'w', 'a', 'n', 't', ' ', '"', '%', 's', '"', '\n', '\x00'})), (*int8)(unsafe.Pointer(&[13]int8{'/', '/', 'u', 's', 'r', '/', '/', 'l', 'i', 'b', '/', '/', '\x00'})), got, (*int8)(unsafe.Pointer(&[4]int8{'l', 'i', 'b', '\x00'})))
		}
	}
	{
		var tmp [100]int8
		var got *int8 = libc.Basename(libc.Strcpy((*int8)(unsafe.Pointer(&tmp)), (*int8)(unsafe.Pointer(&[2]int8{'.', '\x00'}))))
		if libc.Strcmp((*int8)(unsafe.Pointer(&[2]int8{'.', '\x00'})), got) != int32(0) {
			common.T_printf((*int8)(unsafe.Pointer(&[96]int8{'/', 'U', 's', 'e', 'r', 's', '/', 'x', 'u', 's', 'h', 'i', 'w', 'e', 'i', '/', 'w', 'o', 'r', 'k', '/', 'l', 'i', 'b', 'c', '/', 't', 'e', 's', 't', '/', 's', 'r', 'c', '/', 'f', 'u', 'n', 'c', 't', 'i', 'o', 'n', 'a', 'l', '/', 'b', 'a', 's', 'e', 'n', 'a', 'm', 'e', '.', 'c', ':', '2', '5', ':', ' ', 'b', 'a', 's', 'e', 'n', 'a', 'm', 'e', '(', '"', '%', 's', '"', ')', ' ', 'g', 'o', 't', ' ', '"', '%', 's', '"', ' ', 'w', 'a', 'n', 't', ' ', '"', '%', 's', '"', '\n', '\x00'})), (*int8)(unsafe.Pointer(&[2]int8{'.', '\x00'})), got, (*int8)(unsafe.Pointer(&[2]int8{'.', '\x00'})))
		}
	}
	{
		var tmp [100]int8
		var got *int8 = libc.Basename(libc.Strcpy((*int8)(unsafe.Pointer(&tmp)), (*int8)(unsafe.Pointer(&[3]int8{'.', '.', '\x00'}))))
		if libc.Strcmp((*int8)(unsafe.Pointer(&[3]int8{'.', '.', '\x00'})), got) != int32(0) {
			common.T_printf((*int8)(unsafe.Pointer(&[96]int8{'/', 'U', 's', 'e', 'r', 's', '/', 'x', 'u', 's', 'h', 'i', 'w', 'e', 'i', '/', 'w', 'o', 'r', 'k', '/', 'l', 'i', 'b', 'c', '/', 't', 'e', 's', 't', '/', 's', 'r', 'c', '/', 'f', 'u', 'n', 'c', 't', 'i', 'o', 'n', 'a', 'l', '/', 'b', 'a', 's', 'e', 'n', 'a', 'm', 'e', '.', 'c', ':', '2', '6', ':', ' ', 'b', 'a', 's', 'e', 'n', 'a', 'm', 'e', '(', '"', '%', 's', '"', ')', ' ', 'g', 'o', 't', ' ', '"', '%', 's', '"', ' ', 'w', 'a', 'n', 't', ' ', '"', '%', 's', '"', '\n', '\x00'})), (*int8)(unsafe.Pointer(&[3]int8{'.', '.', '\x00'})), got, (*int8)(unsafe.Pointer(&[3]int8{'.', '.', '\x00'})))
		}
	}
	return common.T_status
}
func TestMain(t *testing.T) {
	if _cgo_ret := _cgo_main(); _cgo_ret != 0 {
		t.Fatal("exit status", _cgo_ret)
	}
}
