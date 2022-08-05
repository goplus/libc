package string_strcspn

import (
	unsafe "unsafe"
	libc "github.com/goplus/libc"
	common "github.com/goplus/libc/test/common"
	testing "testing"
)

func _cgo_main() int32 {
	var i int32
	var a [128]int8
	var s [256]int8
	for i = int32(0); i < int32(128); i++ {
		*(*int8)(unsafe.Pointer(uintptr(unsafe.Pointer((*int8)(unsafe.Pointer(&a)))) + uintptr(i))) = int8((i + int32(1)) & int32(127))
	}
	for i = int32(0); i < int32(256); i++ {
		*(*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer((*uint8)(unsafe.Pointer((*int8)(unsafe.Pointer(&s)))))) + uintptr(i))) = uint8(i + int32(1))
	}
	{
		var p *int8 = (*int8)(unsafe.Pointer(&[1]int8{'\x00'}))
		var q *int8 = (*int8)(unsafe.Pointer(&[1]int8{'\x00'}))
		var r uint64 = libc.Strcspn(p, q)
		if r != uint64(0) {
			common.T_printf((*int8)(unsafe.Pointer(&[108]int8{'/', 'U', 's', 'e', 'r', 's', '/', 'x', 'u', 's', 'h', 'i', 'w', 'e', 'i', '/', 'w', 'o', 'r', 'k', '/', 'l', 'i', 'b', 'c', '/', 't', 'e', 's', 't', '/', 's', 'r', 'c', '/', 'f', 'u', 'n', 'c', 't', 'i', 'o', 'n', 'a', 'l', '/', 's', 't', 'r', 'i', 'n', 'g', '_', 's', 't', 'r', 'c', 's', 'p', 'n', '.', 'c', ':', '2', '4', ':', ' ', 's', 't', 'r', 'c', 's', 'p', 'n', '(', '%', 's', ',', '%', 's', ')', ' ', 'r', 'e', 't', 'u', 'r', 'n', 'e', 'd', ' ', '%', 'l', 'u', ',', ' ', 'w', 'a', 'n', 't', 'e', 'd', ' ', '%', 'l', 'u', '\n', '\x00'})), (*int8)(unsafe.Pointer(&[3]int8{'"', '"', '\x00'})), (*int8)(unsafe.Pointer(&[3]int8{'"', '"', '\x00'})), uint64(r), uint64(0))
		}
	}
	{
		var p *int8 = (*int8)(unsafe.Pointer(&[2]int8{'a', '\x00'}))
		var q *int8 = (*int8)(unsafe.Pointer(&[1]int8{'\x00'}))
		var r uint64 = libc.Strcspn(p, q)
		if r != uint64(1) {
			common.T_printf((*int8)(unsafe.Pointer(&[108]int8{'/', 'U', 's', 'e', 'r', 's', '/', 'x', 'u', 's', 'h', 'i', 'w', 'e', 'i', '/', 'w', 'o', 'r', 'k', '/', 'l', 'i', 'b', 'c', '/', 't', 'e', 's', 't', '/', 's', 'r', 'c', '/', 'f', 'u', 'n', 'c', 't', 'i', 'o', 'n', 'a', 'l', '/', 's', 't', 'r', 'i', 'n', 'g', '_', 's', 't', 'r', 'c', 's', 'p', 'n', '.', 'c', ':', '2', '5', ':', ' ', 's', 't', 'r', 'c', 's', 'p', 'n', '(', '%', 's', ',', '%', 's', ')', ' ', 'r', 'e', 't', 'u', 'r', 'n', 'e', 'd', ' ', '%', 'l', 'u', ',', ' ', 'w', 'a', 'n', 't', 'e', 'd', ' ', '%', 'l', 'u', '\n', '\x00'})), (*int8)(unsafe.Pointer(&[4]int8{'"', 'a', '"', '\x00'})), (*int8)(unsafe.Pointer(&[3]int8{'"', '"', '\x00'})), uint64(r), uint64(1))
		}
	}
	{
		var p *int8 = (*int8)(unsafe.Pointer(&[1]int8{'\x00'}))
		var q *int8 = (*int8)(unsafe.Pointer(&[2]int8{'a', '\x00'}))
		var r uint64 = libc.Strcspn(p, q)
		if r != uint64(0) {
			common.T_printf((*int8)(unsafe.Pointer(&[108]int8{'/', 'U', 's', 'e', 'r', 's', '/', 'x', 'u', 's', 'h', 'i', 'w', 'e', 'i', '/', 'w', 'o', 'r', 'k', '/', 'l', 'i', 'b', 'c', '/', 't', 'e', 's', 't', '/', 's', 'r', 'c', '/', 'f', 'u', 'n', 'c', 't', 'i', 'o', 'n', 'a', 'l', '/', 's', 't', 'r', 'i', 'n', 'g', '_', 's', 't', 'r', 'c', 's', 'p', 'n', '.', 'c', ':', '2', '6', ':', ' ', 's', 't', 'r', 'c', 's', 'p', 'n', '(', '%', 's', ',', '%', 's', ')', ' ', 'r', 'e', 't', 'u', 'r', 'n', 'e', 'd', ' ', '%', 'l', 'u', ',', ' ', 'w', 'a', 'n', 't', 'e', 'd', ' ', '%', 'l', 'u', '\n', '\x00'})), (*int8)(unsafe.Pointer(&[3]int8{'"', '"', '\x00'})), (*int8)(unsafe.Pointer(&[4]int8{'"', 'a', '"', '\x00'})), uint64(r), uint64(0))
		}
	}
	{
		var p *int8 = (*int8)(unsafe.Pointer(&[4]int8{'a', 'b', 'c', '\x00'}))
		var q *int8 = (*int8)(unsafe.Pointer(&[4]int8{'c', 'd', 'e', '\x00'}))
		var r uint64 = libc.Strcspn(p, q)
		if r != uint64(2) {
			common.T_printf((*int8)(unsafe.Pointer(&[108]int8{'/', 'U', 's', 'e', 'r', 's', '/', 'x', 'u', 's', 'h', 'i', 'w', 'e', 'i', '/', 'w', 'o', 'r', 'k', '/', 'l', 'i', 'b', 'c', '/', 't', 'e', 's', 't', '/', 's', 'r', 'c', '/', 'f', 'u', 'n', 'c', 't', 'i', 'o', 'n', 'a', 'l', '/', 's', 't', 'r', 'i', 'n', 'g', '_', 's', 't', 'r', 'c', 's', 'p', 'n', '.', 'c', ':', '2', '7', ':', ' ', 's', 't', 'r', 'c', 's', 'p', 'n', '(', '%', 's', ',', '%', 's', ')', ' ', 'r', 'e', 't', 'u', 'r', 'n', 'e', 'd', ' ', '%', 'l', 'u', ',', ' ', 'w', 'a', 'n', 't', 'e', 'd', ' ', '%', 'l', 'u', '\n', '\x00'})), (*int8)(unsafe.Pointer(&[6]int8{'"', 'a', 'b', 'c', '"', '\x00'})), (*int8)(unsafe.Pointer(&[6]int8{'"', 'c', 'd', 'e', '"', '\x00'})), uint64(r), uint64(2))
		}
	}
	{
		var p *int8 = (*int8)(unsafe.Pointer(&[4]int8{'a', 'b', 'c', '\x00'}))
		var q *int8 = (*int8)(unsafe.Pointer(&[4]int8{'c', 'c', 'c', '\x00'}))
		var r uint64 = libc.Strcspn(p, q)
		if r != uint64(2) {
			common.T_printf((*int8)(unsafe.Pointer(&[108]int8{'/', 'U', 's', 'e', 'r', 's', '/', 'x', 'u', 's', 'h', 'i', 'w', 'e', 'i', '/', 'w', 'o', 'r', 'k', '/', 'l', 'i', 'b', 'c', '/', 't', 'e', 's', 't', '/', 's', 'r', 'c', '/', 'f', 'u', 'n', 'c', 't', 'i', 'o', 'n', 'a', 'l', '/', 's', 't', 'r', 'i', 'n', 'g', '_', 's', 't', 'r', 'c', 's', 'p', 'n', '.', 'c', ':', '2', '8', ':', ' ', 's', 't', 'r', 'c', 's', 'p', 'n', '(', '%', 's', ',', '%', 's', ')', ' ', 'r', 'e', 't', 'u', 'r', 'n', 'e', 'd', ' ', '%', 'l', 'u', ',', ' ', 'w', 'a', 'n', 't', 'e', 'd', ' ', '%', 'l', 'u', '\n', '\x00'})), (*int8)(unsafe.Pointer(&[6]int8{'"', 'a', 'b', 'c', '"', '\x00'})), (*int8)(unsafe.Pointer(&[6]int8{'"', 'c', 'c', 'c', '"', '\x00'})), uint64(r), uint64(2))
		}
	}
	{
		var p *int8 = (*int8)(unsafe.Pointer(&[4]int8{'a', 'b', 'c', '\x00'}))
		var q *int8 = (*int8)(unsafe.Pointer(&a))
		var r uint64 = libc.Strcspn(p, q)
		if r != uint64(0) {
			common.T_printf((*int8)(unsafe.Pointer(&[108]int8{'/', 'U', 's', 'e', 'r', 's', '/', 'x', 'u', 's', 'h', 'i', 'w', 'e', 'i', '/', 'w', 'o', 'r', 'k', '/', 'l', 'i', 'b', 'c', '/', 't', 'e', 's', 't', '/', 's', 'r', 'c', '/', 'f', 'u', 'n', 'c', 't', 'i', 'o', 'n', 'a', 'l', '/', 's', 't', 'r', 'i', 'n', 'g', '_', 's', 't', 'r', 'c', 's', 'p', 'n', '.', 'c', ':', '2', '9', ':', ' ', 's', 't', 'r', 'c', 's', 'p', 'n', '(', '%', 's', ',', '%', 's', ')', ' ', 'r', 'e', 't', 'u', 'r', 'n', 'e', 'd', ' ', '%', 'l', 'u', ',', ' ', 'w', 'a', 'n', 't', 'e', 'd', ' ', '%', 'l', 'u', '\n', '\x00'})), (*int8)(unsafe.Pointer(&[6]int8{'"', 'a', 'b', 'c', '"', '\x00'})), (*int8)(unsafe.Pointer(&[2]int8{'a', '\x00'})), uint64(r), uint64(0))
		}
	}
	{
		var p *int8 = (*int8)(unsafe.Pointer(&[7]int8{-1, -128, ' ', 'a', 'b', 'c', '\x00'}))
		var q *int8 = (*int8)(unsafe.Pointer(&a))
		var r uint64 = libc.Strcspn(p, q)
		if r != uint64(2) {
			common.T_printf((*int8)(unsafe.Pointer(&[108]int8{'/', 'U', 's', 'e', 'r', 's', '/', 'x', 'u', 's', 'h', 'i', 'w', 'e', 'i', '/', 'w', 'o', 'r', 'k', '/', 'l', 'i', 'b', 'c', '/', 't', 'e', 's', 't', '/', 's', 'r', 'c', '/', 'f', 'u', 'n', 'c', 't', 'i', 'o', 'n', 'a', 'l', '/', 's', 't', 'r', 'i', 'n', 'g', '_', 's', 't', 'r', 'c', 's', 'p', 'n', '.', 'c', ':', '3', '0', ':', ' ', 's', 't', 'r', 'c', 's', 'p', 'n', '(', '%', 's', ',', '%', 's', ')', ' ', 'r', 'e', 't', 'u', 'r', 'n', 'e', 'd', ' ', '%', 'l', 'u', ',', ' ', 'w', 'a', 'n', 't', 'e', 'd', ' ', '%', 'l', 'u', '\n', '\x00'})), (*int8)(unsafe.Pointer(&[15]int8{'"', '\\', 'x', 'f', 'f', '\\', 'x', '8', '0', ' ', 'a', 'b', 'c', '"', '\x00'})), (*int8)(unsafe.Pointer(&[2]int8{'a', '\x00'})), uint64(r), uint64(2))
		}
	}
	{
		var p *int8 = (*int8)(unsafe.Pointer(&s))
		var q *int8 = (*int8)(unsafe.Pointer(&[2]int8{-1, '\x00'}))
		var r uint64 = libc.Strcspn(p, q)
		if r != uint64(254) {
			common.T_printf((*int8)(unsafe.Pointer(&[108]int8{'/', 'U', 's', 'e', 'r', 's', '/', 'x', 'u', 's', 'h', 'i', 'w', 'e', 'i', '/', 'w', 'o', 'r', 'k', '/', 'l', 'i', 'b', 'c', '/', 't', 'e', 's', 't', '/', 's', 'r', 'c', '/', 'f', 'u', 'n', 'c', 't', 'i', 'o', 'n', 'a', 'l', '/', 's', 't', 'r', 'i', 'n', 'g', '_', 's', 't', 'r', 'c', 's', 'p', 'n', '.', 'c', ':', '3', '1', ':', ' ', 's', 't', 'r', 'c', 's', 'p', 'n', '(', '%', 's', ',', '%', 's', ')', ' ', 'r', 'e', 't', 'u', 'r', 'n', 'e', 'd', ' ', '%', 'l', 'u', ',', ' ', 'w', 'a', 'n', 't', 'e', 'd', ' ', '%', 'l', 'u', '\n', '\x00'})), (*int8)(unsafe.Pointer(&[2]int8{'s', '\x00'})), (*int8)(unsafe.Pointer(&[7]int8{'"', '\\', 'x', 'f', 'f', '"', '\x00'})), uint64(r), uint64(254))
		}
	}
	return common.T_status
}
func TestMain(t *testing.T) {
	if _cgo_ret := _cgo_main(); _cgo_ret != 0 {
		t.Fatal("exit status", _cgo_ret)
	}
}
