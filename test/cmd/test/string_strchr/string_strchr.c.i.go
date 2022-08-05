package string_strchr

import (
	unsafe "unsafe"
	libc "github.com/goplus/libc"
	common "github.com/goplus/libc/test/common"
	testing "testing"
)

var _cgos_buf_string_strchr [512]int8

func _cgos_aligned_string_strchr(p unsafe.Pointer) unsafe.Pointer {
	return unsafe.Pointer(uintptr((uint64(uintptr(p)) + uint64(63)) & uint64(18446744073709551552)))
}
func _cgos_aligncpy_string_strchr(p unsafe.Pointer, len uint64, a uint64) unsafe.Pointer {
	return libc.Memcpy(unsafe.Pointer((*int8)(unsafe.Pointer(uintptr(unsafe.Pointer((*int8)(_cgos_aligned_string_strchr(unsafe.Pointer((*int8)(unsafe.Pointer(&_cgos_buf_string_strchr)))))))+uintptr(a)))), p, len)
}
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
		var align int32
		for align = int32(0); align < int32(8); align++ {
			var p *int8 = (*int8)(_cgos_aligncpy_string_strchr(unsafe.Pointer((*int8)(unsafe.Pointer(&[5]int8{'\x00', 'a', 'a', 'a', '\x00'}))), 5, uint64(align)))
			var q *int8 = libc.Strchr(p, 'a')
			if q != nil {
				common.T_printf((*int8)(unsafe.Pointer(&[121]int8{'/', 'U', 's', 'e', 'r', 's', '/', 'x', 'u', 's', 'h', 'i', 'w', 'e', 'i', '/', 'w', 'o', 'r', 'k', '/', 'l', 'i', 'b', 'c', '/', 't', 'e', 's', 't', '/', 's', 'r', 'c', '/', 'f', 'u', 'n', 'c', 't', 'i', 'o', 'n', 'a', 'l', '/', 's', 't', 'r', 'i', 'n', 'g', '_', 's', 't', 'r', 'c', 'h', 'r', '.', 'c', ':', '4', '9', ':', ' ', 's', 't', 'r', 'c', 'h', 'r', '(', '%', 's', ',', '%', 's', ')', ' ', 'w', 'i', 't', 'h', ' ', 'a', 'l', 'i', 'g', 'n', '=', '%', 'd', ' ', 'r', 'e', 't', 'u', 'r', 'n', 'e', 'd', ' ', 's', 't', 'r', '+', '%', 'd', ',', ' ', 'w', 'a', 'n', 't', 'e', 'd', ' ', '0', '\n', '\x00'})), (*int8)(unsafe.Pointer(&[8]int8{'"', '\\', '0', 'a', 'a', 'a', '"', '\x00'})), (*int8)(unsafe.Pointer(&[4]int8{'\'', 'a', '\'', '\x00'})), align, uintptr(unsafe.Pointer(q))-uintptr(unsafe.Pointer(p)))
			}
		}
	}
	{
		var align int32
		for align = int32(0); align < int32(8); align++ {
			var p *int8 = (*int8)(_cgos_aligncpy_string_strchr(unsafe.Pointer((*int8)(unsafe.Pointer(&[5]int8{'a', '\x00', 'b', 'b', '\x00'}))), 5, uint64(align)))
			var q *int8 = libc.Strchr(p, 'b')
			if q != nil {
				common.T_printf((*int8)(unsafe.Pointer(&[121]int8{'/', 'U', 's', 'e', 'r', 's', '/', 'x', 'u', 's', 'h', 'i', 'w', 'e', 'i', '/', 'w', 'o', 'r', 'k', '/', 'l', 'i', 'b', 'c', '/', 't', 'e', 's', 't', '/', 's', 'r', 'c', '/', 'f', 'u', 'n', 'c', 't', 'i', 'o', 'n', 'a', 'l', '/', 's', 't', 'r', 'i', 'n', 'g', '_', 's', 't', 'r', 'c', 'h', 'r', '.', 'c', ':', '5', '0', ':', ' ', 's', 't', 'r', 'c', 'h', 'r', '(', '%', 's', ',', '%', 's', ')', ' ', 'w', 'i', 't', 'h', ' ', 'a', 'l', 'i', 'g', 'n', '=', '%', 'd', ' ', 'r', 'e', 't', 'u', 'r', 'n', 'e', 'd', ' ', 's', 't', 'r', '+', '%', 'd', ',', ' ', 'w', 'a', 'n', 't', 'e', 'd', ' ', '0', '\n', '\x00'})), (*int8)(unsafe.Pointer(&[8]int8{'"', 'a', '\\', '0', 'b', 'b', '"', '\x00'})), (*int8)(unsafe.Pointer(&[4]int8{'\'', 'b', '\'', '\x00'})), align, uintptr(unsafe.Pointer(q))-uintptr(unsafe.Pointer(p)))
			}
		}
	}
	{
		var align int32
		for align = int32(0); align < int32(8); align++ {
			var p *int8 = (*int8)(_cgos_aligncpy_string_strchr(unsafe.Pointer((*int8)(unsafe.Pointer(&[5]int8{'a', 'b', '\x00', 'c', '\x00'}))), 5, uint64(align)))
			var q *int8 = libc.Strchr(p, 'c')
			if q != nil {
				common.T_printf((*int8)(unsafe.Pointer(&[121]int8{'/', 'U', 's', 'e', 'r', 's', '/', 'x', 'u', 's', 'h', 'i', 'w', 'e', 'i', '/', 'w', 'o', 'r', 'k', '/', 'l', 'i', 'b', 'c', '/', 't', 'e', 's', 't', '/', 's', 'r', 'c', '/', 'f', 'u', 'n', 'c', 't', 'i', 'o', 'n', 'a', 'l', '/', 's', 't', 'r', 'i', 'n', 'g', '_', 's', 't', 'r', 'c', 'h', 'r', '.', 'c', ':', '5', '1', ':', ' ', 's', 't', 'r', 'c', 'h', 'r', '(', '%', 's', ',', '%', 's', ')', ' ', 'w', 'i', 't', 'h', ' ', 'a', 'l', 'i', 'g', 'n', '=', '%', 'd', ' ', 'r', 'e', 't', 'u', 'r', 'n', 'e', 'd', ' ', 's', 't', 'r', '+', '%', 'd', ',', ' ', 'w', 'a', 'n', 't', 'e', 'd', ' ', '0', '\n', '\x00'})), (*int8)(unsafe.Pointer(&[8]int8{'"', 'a', 'b', '\\', '0', 'c', '"', '\x00'})), (*int8)(unsafe.Pointer(&[4]int8{'\'', 'c', '\'', '\x00'})), align, uintptr(unsafe.Pointer(q))-uintptr(unsafe.Pointer(p)))
			}
		}
	}
	{
		var align int32
		for align = int32(0); align < int32(8); align++ {
			var p *int8 = (*int8)(_cgos_aligncpy_string_strchr(unsafe.Pointer((*int8)(unsafe.Pointer(&[6]int8{'a', 'b', 'c', '\x00', 'd', '\x00'}))), 6, uint64(align)))
			var q *int8 = libc.Strchr(p, 'd')
			if q != nil {
				common.T_printf((*int8)(unsafe.Pointer(&[121]int8{'/', 'U', 's', 'e', 'r', 's', '/', 'x', 'u', 's', 'h', 'i', 'w', 'e', 'i', '/', 'w', 'o', 'r', 'k', '/', 'l', 'i', 'b', 'c', '/', 't', 'e', 's', 't', '/', 's', 'r', 'c', '/', 'f', 'u', 'n', 'c', 't', 'i', 'o', 'n', 'a', 'l', '/', 's', 't', 'r', 'i', 'n', 'g', '_', 's', 't', 'r', 'c', 'h', 'r', '.', 'c', ':', '5', '2', ':', ' ', 's', 't', 'r', 'c', 'h', 'r', '(', '%', 's', ',', '%', 's', ')', ' ', 'w', 'i', 't', 'h', ' ', 'a', 'l', 'i', 'g', 'n', '=', '%', 'd', ' ', 'r', 'e', 't', 'u', 'r', 'n', 'e', 'd', ' ', 's', 't', 'r', '+', '%', 'd', ',', ' ', 'w', 'a', 'n', 't', 'e', 'd', ' ', '0', '\n', '\x00'})), (*int8)(unsafe.Pointer(&[9]int8{'"', 'a', 'b', 'c', '\\', '0', 'd', '"', '\x00'})), (*int8)(unsafe.Pointer(&[4]int8{'\'', 'd', '\'', '\x00'})), align, uintptr(unsafe.Pointer(q))-uintptr(unsafe.Pointer(p)))
			}
		}
	}
	{
		var align int32
		for align = int32(0); align < int32(8); align++ {
			var p *int8 = (*int8)(_cgos_aligncpy_string_strchr(unsafe.Pointer((*int8)(unsafe.Pointer(&[10]int8{'a', 'b', 'c', ' ', 'a', 'b', 'c', '\x00', 'x', '\x00'}))), 10, uint64(align)))
			var q *int8 = libc.Strchr(p, 'x')
			if q != nil {
				common.T_printf((*int8)(unsafe.Pointer(&[121]int8{'/', 'U', 's', 'e', 'r', 's', '/', 'x', 'u', 's', 'h', 'i', 'w', 'e', 'i', '/', 'w', 'o', 'r', 'k', '/', 'l', 'i', 'b', 'c', '/', 't', 'e', 's', 't', '/', 's', 'r', 'c', '/', 'f', 'u', 'n', 'c', 't', 'i', 'o', 'n', 'a', 'l', '/', 's', 't', 'r', 'i', 'n', 'g', '_', 's', 't', 'r', 'c', 'h', 'r', '.', 'c', ':', '5', '3', ':', ' ', 's', 't', 'r', 'c', 'h', 'r', '(', '%', 's', ',', '%', 's', ')', ' ', 'w', 'i', 't', 'h', ' ', 'a', 'l', 'i', 'g', 'n', '=', '%', 'd', ' ', 'r', 'e', 't', 'u', 'r', 'n', 'e', 'd', ' ', 's', 't', 'r', '+', '%', 'd', ',', ' ', 'w', 'a', 'n', 't', 'e', 'd', ' ', '0', '\n', '\x00'})), (*int8)(unsafe.Pointer(&[13]int8{'"', 'a', 'b', 'c', ' ', 'a', 'b', 'c', '\\', '0', 'x', '"', '\x00'})), (*int8)(unsafe.Pointer(&[4]int8{'\'', 'x', '\'', '\x00'})), align, uintptr(unsafe.Pointer(q))-uintptr(unsafe.Pointer(p)))
			}
		}
	}
	{
		var align int32
		for align = int32(0); align < int32(8); align++ {
			var p *int8 = (*int8)(_cgos_aligncpy_string_strchr(unsafe.Pointer((*int8)(unsafe.Pointer(&a))), 128, uint64(align)))
			var q *int8 = libc.Strchr(p, int32(128))
			if q != nil {
				common.T_printf((*int8)(unsafe.Pointer(&[121]int8{'/', 'U', 's', 'e', 'r', 's', '/', 'x', 'u', 's', 'h', 'i', 'w', 'e', 'i', '/', 'w', 'o', 'r', 'k', '/', 'l', 'i', 'b', 'c', '/', 't', 'e', 's', 't', '/', 's', 'r', 'c', '/', 'f', 'u', 'n', 'c', 't', 'i', 'o', 'n', 'a', 'l', '/', 's', 't', 'r', 'i', 'n', 'g', '_', 's', 't', 'r', 'c', 'h', 'r', '.', 'c', ':', '5', '4', ':', ' ', 's', 't', 'r', 'c', 'h', 'r', '(', '%', 's', ',', '%', 's', ')', ' ', 'w', 'i', 't', 'h', ' ', 'a', 'l', 'i', 'g', 'n', '=', '%', 'd', ' ', 'r', 'e', 't', 'u', 'r', 'n', 'e', 'd', ' ', 's', 't', 'r', '+', '%', 'd', ',', ' ', 'w', 'a', 'n', 't', 'e', 'd', ' ', '0', '\n', '\x00'})), (*int8)(unsafe.Pointer(&[2]int8{'a', '\x00'})), (*int8)(unsafe.Pointer(&[4]int8{'1', '2', '8', '\x00'})), align, uintptr(unsafe.Pointer(q))-uintptr(unsafe.Pointer(p)))
			}
		}
	}
	{
		var align int32
		for align = int32(0); align < int32(8); align++ {
			var p *int8 = (*int8)(_cgos_aligncpy_string_strchr(unsafe.Pointer((*int8)(unsafe.Pointer(&a))), 128, uint64(align)))
			var q *int8 = libc.Strchr(p, int32(255))
			if q != nil {
				common.T_printf((*int8)(unsafe.Pointer(&[121]int8{'/', 'U', 's', 'e', 'r', 's', '/', 'x', 'u', 's', 'h', 'i', 'w', 'e', 'i', '/', 'w', 'o', 'r', 'k', '/', 'l', 'i', 'b', 'c', '/', 't', 'e', 's', 't', '/', 's', 'r', 'c', '/', 'f', 'u', 'n', 'c', 't', 'i', 'o', 'n', 'a', 'l', '/', 's', 't', 'r', 'i', 'n', 'g', '_', 's', 't', 'r', 'c', 'h', 'r', '.', 'c', ':', '5', '5', ':', ' ', 's', 't', 'r', 'c', 'h', 'r', '(', '%', 's', ',', '%', 's', ')', ' ', 'w', 'i', 't', 'h', ' ', 'a', 'l', 'i', 'g', 'n', '=', '%', 'd', ' ', 'r', 'e', 't', 'u', 'r', 'n', 'e', 'd', ' ', 's', 't', 'r', '+', '%', 'd', ',', ' ', 'w', 'a', 'n', 't', 'e', 'd', ' ', '0', '\n', '\x00'})), (*int8)(unsafe.Pointer(&[2]int8{'a', '\x00'})), (*int8)(unsafe.Pointer(&[4]int8{'2', '5', '5', '\x00'})), align, uintptr(unsafe.Pointer(q))-uintptr(unsafe.Pointer(p)))
			}
		}
	}
	{
		var align int32
		for align = int32(0); align < int32(8); align++ {
			var p *int8 = (*int8)(_cgos_aligncpy_string_strchr(unsafe.Pointer((*int8)(unsafe.Pointer(&[1]int8{'\x00'}))), 1, uint64(align)))
			var q *int8 = libc.Strchr(p, int32(0))
			if uintptr(unsafe.Pointer(q)) == uintptr(unsafe.Pointer(nil)) {
				common.T_printf((*int8)(unsafe.Pointer(&[121]int8{'/', 'U', 's', 'e', 'r', 's', '/', 'x', 'u', 's', 'h', 'i', 'w', 'e', 'i', '/', 'w', 'o', 'r', 'k', '/', 'l', 'i', 'b', 'c', '/', 't', 'e', 's', 't', '/', 's', 'r', 'c', '/', 'f', 'u', 'n', 'c', 't', 'i', 'o', 'n', 'a', 'l', '/', 's', 't', 'r', 'i', 'n', 'g', '_', 's', 't', 'r', 'c', 'h', 'r', '.', 'c', ':', '5', '7', ':', ' ', 's', 't', 'r', 'c', 'h', 'r', '(', '%', 's', ',', '%', 's', ')', ' ', 'w', 'i', 't', 'h', ' ', 'a', 'l', 'i', 'g', 'n', '=', '%', 'd', ' ', 'r', 'e', 't', 'u', 'r', 'n', 'e', 'd', ' ', '0', ',', ' ', 'w', 'a', 'n', 't', 'e', 'd', ' ', 's', 't', 'r', '+', '%', 'd', '\n', '\x00'})), (*int8)(unsafe.Pointer(&[3]int8{'"', '"', '\x00'})), (*int8)(unsafe.Pointer(&[2]int8{'0', '\x00'})), align, int32(0))
			} else if int64(uintptr(unsafe.Pointer(q))-uintptr(unsafe.Pointer(p))) != int64(0) {
				common.T_printf((*int8)(unsafe.Pointer(&[126]int8{'/', 'U', 's', 'e', 'r', 's', '/', 'x', 'u', 's', 'h', 'i', 'w', 'e', 'i', '/', 'w', 'o', 'r', 'k', '/', 'l', 'i', 'b', 'c', '/', 't', 'e', 's', 't', '/', 's', 'r', 'c', '/', 'f', 'u', 'n', 'c', 't', 'i', 'o', 'n', 'a', 'l', '/', 's', 't', 'r', 'i', 'n', 'g', '_', 's', 't', 'r', 'c', 'h', 'r', '.', 'c', ':', '5', '7', ':', ' ', 's', 't', 'r', 'c', 'h', 'r', '(', '%', 's', ',', '%', 's', ')', ' ', 'w', 'i', 't', 'h', ' ', 'a', 'l', 'i', 'g', 'n', '=', '%', 'd', ' ', 'r', 'e', 't', 'u', 'r', 'n', 'e', 'd', ' ', 's', 't', 'r', '+', '%', 'd', ',', ' ', 'w', 'a', 'n', 't', 'e', 'd', ' ', 's', 't', 'r', '+', '%', 'd', '\n', '\x00'})), (*int8)(unsafe.Pointer(&[3]int8{'"', '"', '\x00'})), (*int8)(unsafe.Pointer(&[2]int8{'0', '\x00'})), align, uintptr(unsafe.Pointer(q))-uintptr(unsafe.Pointer(p)), int32(0))
			}
		}
	}
	{
		var align int32
		for align = int32(0); align < int32(8); align++ {
			var p *int8 = (*int8)(_cgos_aligncpy_string_strchr(unsafe.Pointer((*int8)(unsafe.Pointer(&[2]int8{'a', '\x00'}))), 2, uint64(align)))
			var q *int8 = libc.Strchr(p, 'a')
			if uintptr(unsafe.Pointer(q)) == uintptr(unsafe.Pointer(nil)) {
				common.T_printf((*int8)(unsafe.Pointer(&[121]int8{'/', 'U', 's', 'e', 'r', 's', '/', 'x', 'u', 's', 'h', 'i', 'w', 'e', 'i', '/', 'w', 'o', 'r', 'k', '/', 'l', 'i', 'b', 'c', '/', 't', 'e', 's', 't', '/', 's', 'r', 'c', '/', 'f', 'u', 'n', 'c', 't', 'i', 'o', 'n', 'a', 'l', '/', 's', 't', 'r', 'i', 'n', 'g', '_', 's', 't', 'r', 'c', 'h', 'r', '.', 'c', ':', '5', '8', ':', ' ', 's', 't', 'r', 'c', 'h', 'r', '(', '%', 's', ',', '%', 's', ')', ' ', 'w', 'i', 't', 'h', ' ', 'a', 'l', 'i', 'g', 'n', '=', '%', 'd', ' ', 'r', 'e', 't', 'u', 'r', 'n', 'e', 'd', ' ', '0', ',', ' ', 'w', 'a', 'n', 't', 'e', 'd', ' ', 's', 't', 'r', '+', '%', 'd', '\n', '\x00'})), (*int8)(unsafe.Pointer(&[4]int8{'"', 'a', '"', '\x00'})), (*int8)(unsafe.Pointer(&[4]int8{'\'', 'a', '\'', '\x00'})), align, int32(0))
			} else if int64(uintptr(unsafe.Pointer(q))-uintptr(unsafe.Pointer(p))) != int64(0) {
				common.T_printf((*int8)(unsafe.Pointer(&[126]int8{'/', 'U', 's', 'e', 'r', 's', '/', 'x', 'u', 's', 'h', 'i', 'w', 'e', 'i', '/', 'w', 'o', 'r', 'k', '/', 'l', 'i', 'b', 'c', '/', 't', 'e', 's', 't', '/', 's', 'r', 'c', '/', 'f', 'u', 'n', 'c', 't', 'i', 'o', 'n', 'a', 'l', '/', 's', 't', 'r', 'i', 'n', 'g', '_', 's', 't', 'r', 'c', 'h', 'r', '.', 'c', ':', '5', '8', ':', ' ', 's', 't', 'r', 'c', 'h', 'r', '(', '%', 's', ',', '%', 's', ')', ' ', 'w', 'i', 't', 'h', ' ', 'a', 'l', 'i', 'g', 'n', '=', '%', 'd', ' ', 'r', 'e', 't', 'u', 'r', 'n', 'e', 'd', ' ', 's', 't', 'r', '+', '%', 'd', ',', ' ', 'w', 'a', 'n', 't', 'e', 'd', ' ', 's', 't', 'r', '+', '%', 'd', '\n', '\x00'})), (*int8)(unsafe.Pointer(&[4]int8{'"', 'a', '"', '\x00'})), (*int8)(unsafe.Pointer(&[4]int8{'\'', 'a', '\'', '\x00'})), align, uintptr(unsafe.Pointer(q))-uintptr(unsafe.Pointer(p)), int32(0))
			}
		}
	}
	{
		var align int32
		for align = int32(0); align < int32(8); align++ {
			var p *int8 = (*int8)(_cgos_aligncpy_string_strchr(unsafe.Pointer((*int8)(unsafe.Pointer(&[2]int8{'a', '\x00'}))), 2, uint64(align)))
			var q *int8 = libc.Strchr(p, 353)
			if uintptr(unsafe.Pointer(q)) == uintptr(unsafe.Pointer(nil)) {
				common.T_printf((*int8)(unsafe.Pointer(&[121]int8{'/', 'U', 's', 'e', 'r', 's', '/', 'x', 'u', 's', 'h', 'i', 'w', 'e', 'i', '/', 'w', 'o', 'r', 'k', '/', 'l', 'i', 'b', 'c', '/', 't', 'e', 's', 't', '/', 's', 'r', 'c', '/', 'f', 'u', 'n', 'c', 't', 'i', 'o', 'n', 'a', 'l', '/', 's', 't', 'r', 'i', 'n', 'g', '_', 's', 't', 'r', 'c', 'h', 'r', '.', 'c', ':', '5', '9', ':', ' ', 's', 't', 'r', 'c', 'h', 'r', '(', '%', 's', ',', '%', 's', ')', ' ', 'w', 'i', 't', 'h', ' ', 'a', 'l', 'i', 'g', 'n', '=', '%', 'd', ' ', 'r', 'e', 't', 'u', 'r', 'n', 'e', 'd', ' ', '0', ',', ' ', 'w', 'a', 'n', 't', 'e', 'd', ' ', 's', 't', 'r', '+', '%', 'd', '\n', '\x00'})), (*int8)(unsafe.Pointer(&[4]int8{'"', 'a', '"', '\x00'})), (*int8)(unsafe.Pointer(&[8]int8{'\'', 'a', '\'', '+', '2', '5', '6', '\x00'})), align, int32(0))
			} else if int64(uintptr(unsafe.Pointer(q))-uintptr(unsafe.Pointer(p))) != int64(0) {
				common.T_printf((*int8)(unsafe.Pointer(&[126]int8{'/', 'U', 's', 'e', 'r', 's', '/', 'x', 'u', 's', 'h', 'i', 'w', 'e', 'i', '/', 'w', 'o', 'r', 'k', '/', 'l', 'i', 'b', 'c', '/', 't', 'e', 's', 't', '/', 's', 'r', 'c', '/', 'f', 'u', 'n', 'c', 't', 'i', 'o', 'n', 'a', 'l', '/', 's', 't', 'r', 'i', 'n', 'g', '_', 's', 't', 'r', 'c', 'h', 'r', '.', 'c', ':', '5', '9', ':', ' ', 's', 't', 'r', 'c', 'h', 'r', '(', '%', 's', ',', '%', 's', ')', ' ', 'w', 'i', 't', 'h', ' ', 'a', 'l', 'i', 'g', 'n', '=', '%', 'd', ' ', 'r', 'e', 't', 'u', 'r', 'n', 'e', 'd', ' ', 's', 't', 'r', '+', '%', 'd', ',', ' ', 'w', 'a', 'n', 't', 'e', 'd', ' ', 's', 't', 'r', '+', '%', 'd', '\n', '\x00'})), (*int8)(unsafe.Pointer(&[4]int8{'"', 'a', '"', '\x00'})), (*int8)(unsafe.Pointer(&[8]int8{'\'', 'a', '\'', '+', '2', '5', '6', '\x00'})), align, uintptr(unsafe.Pointer(q))-uintptr(unsafe.Pointer(p)), int32(0))
			}
		}
	}
	{
		var align int32
		for align = int32(0); align < int32(8); align++ {
			var p *int8 = (*int8)(_cgos_aligncpy_string_strchr(unsafe.Pointer((*int8)(unsafe.Pointer(&[2]int8{'a', '\x00'}))), 2, uint64(align)))
			var q *int8 = libc.Strchr(p, int32(0))
			if uintptr(unsafe.Pointer(q)) == uintptr(unsafe.Pointer(nil)) {
				common.T_printf((*int8)(unsafe.Pointer(&[121]int8{'/', 'U', 's', 'e', 'r', 's', '/', 'x', 'u', 's', 'h', 'i', 'w', 'e', 'i', '/', 'w', 'o', 'r', 'k', '/', 'l', 'i', 'b', 'c', '/', 't', 'e', 's', 't', '/', 's', 'r', 'c', '/', 'f', 'u', 'n', 'c', 't', 'i', 'o', 'n', 'a', 'l', '/', 's', 't', 'r', 'i', 'n', 'g', '_', 's', 't', 'r', 'c', 'h', 'r', '.', 'c', ':', '6', '0', ':', ' ', 's', 't', 'r', 'c', 'h', 'r', '(', '%', 's', ',', '%', 's', ')', ' ', 'w', 'i', 't', 'h', ' ', 'a', 'l', 'i', 'g', 'n', '=', '%', 'd', ' ', 'r', 'e', 't', 'u', 'r', 'n', 'e', 'd', ' ', '0', ',', ' ', 'w', 'a', 'n', 't', 'e', 'd', ' ', 's', 't', 'r', '+', '%', 'd', '\n', '\x00'})), (*int8)(unsafe.Pointer(&[4]int8{'"', 'a', '"', '\x00'})), (*int8)(unsafe.Pointer(&[2]int8{'0', '\x00'})), align, int32(1))
			} else if int64(uintptr(unsafe.Pointer(q))-uintptr(unsafe.Pointer(p))) != int64(1) {
				common.T_printf((*int8)(unsafe.Pointer(&[126]int8{'/', 'U', 's', 'e', 'r', 's', '/', 'x', 'u', 's', 'h', 'i', 'w', 'e', 'i', '/', 'w', 'o', 'r', 'k', '/', 'l', 'i', 'b', 'c', '/', 't', 'e', 's', 't', '/', 's', 'r', 'c', '/', 'f', 'u', 'n', 'c', 't', 'i', 'o', 'n', 'a', 'l', '/', 's', 't', 'r', 'i', 'n', 'g', '_', 's', 't', 'r', 'c', 'h', 'r', '.', 'c', ':', '6', '0', ':', ' ', 's', 't', 'r', 'c', 'h', 'r', '(', '%', 's', ',', '%', 's', ')', ' ', 'w', 'i', 't', 'h', ' ', 'a', 'l', 'i', 'g', 'n', '=', '%', 'd', ' ', 'r', 'e', 't', 'u', 'r', 'n', 'e', 'd', ' ', 's', 't', 'r', '+', '%', 'd', ',', ' ', 'w', 'a', 'n', 't', 'e', 'd', ' ', 's', 't', 'r', '+', '%', 'd', '\n', '\x00'})), (*int8)(unsafe.Pointer(&[4]int8{'"', 'a', '"', '\x00'})), (*int8)(unsafe.Pointer(&[2]int8{'0', '\x00'})), align, uintptr(unsafe.Pointer(q))-uintptr(unsafe.Pointer(p)), int32(1))
			}
		}
	}
	{
		var align int32
		for align = int32(0); align < int32(8); align++ {
			var p *int8 = (*int8)(_cgos_aligncpy_string_strchr(unsafe.Pointer((*int8)(unsafe.Pointer(&[4]int8{'a', 'b', 'b', '\x00'}))), 4, uint64(align)))
			var q *int8 = libc.Strchr(p, 'b')
			if uintptr(unsafe.Pointer(q)) == uintptr(unsafe.Pointer(nil)) {
				common.T_printf((*int8)(unsafe.Pointer(&[121]int8{'/', 'U', 's', 'e', 'r', 's', '/', 'x', 'u', 's', 'h', 'i', 'w', 'e', 'i', '/', 'w', 'o', 'r', 'k', '/', 'l', 'i', 'b', 'c', '/', 't', 'e', 's', 't', '/', 's', 'r', 'c', '/', 'f', 'u', 'n', 'c', 't', 'i', 'o', 'n', 'a', 'l', '/', 's', 't', 'r', 'i', 'n', 'g', '_', 's', 't', 'r', 'c', 'h', 'r', '.', 'c', ':', '6', '1', ':', ' ', 's', 't', 'r', 'c', 'h', 'r', '(', '%', 's', ',', '%', 's', ')', ' ', 'w', 'i', 't', 'h', ' ', 'a', 'l', 'i', 'g', 'n', '=', '%', 'd', ' ', 'r', 'e', 't', 'u', 'r', 'n', 'e', 'd', ' ', '0', ',', ' ', 'w', 'a', 'n', 't', 'e', 'd', ' ', 's', 't', 'r', '+', '%', 'd', '\n', '\x00'})), (*int8)(unsafe.Pointer(&[6]int8{'"', 'a', 'b', 'b', '"', '\x00'})), (*int8)(unsafe.Pointer(&[4]int8{'\'', 'b', '\'', '\x00'})), align, int32(1))
			} else if int64(uintptr(unsafe.Pointer(q))-uintptr(unsafe.Pointer(p))) != int64(1) {
				common.T_printf((*int8)(unsafe.Pointer(&[126]int8{'/', 'U', 's', 'e', 'r', 's', '/', 'x', 'u', 's', 'h', 'i', 'w', 'e', 'i', '/', 'w', 'o', 'r', 'k', '/', 'l', 'i', 'b', 'c', '/', 't', 'e', 's', 't', '/', 's', 'r', 'c', '/', 'f', 'u', 'n', 'c', 't', 'i', 'o', 'n', 'a', 'l', '/', 's', 't', 'r', 'i', 'n', 'g', '_', 's', 't', 'r', 'c', 'h', 'r', '.', 'c', ':', '6', '1', ':', ' ', 's', 't', 'r', 'c', 'h', 'r', '(', '%', 's', ',', '%', 's', ')', ' ', 'w', 'i', 't', 'h', ' ', 'a', 'l', 'i', 'g', 'n', '=', '%', 'd', ' ', 'r', 'e', 't', 'u', 'r', 'n', 'e', 'd', ' ', 's', 't', 'r', '+', '%', 'd', ',', ' ', 'w', 'a', 'n', 't', 'e', 'd', ' ', 's', 't', 'r', '+', '%', 'd', '\n', '\x00'})), (*int8)(unsafe.Pointer(&[6]int8{'"', 'a', 'b', 'b', '"', '\x00'})), (*int8)(unsafe.Pointer(&[4]int8{'\'', 'b', '\'', '\x00'})), align, uintptr(unsafe.Pointer(q))-uintptr(unsafe.Pointer(p)), int32(1))
			}
		}
	}
	{
		var align int32
		for align = int32(0); align < int32(8); align++ {
			var p *int8 = (*int8)(_cgos_aligncpy_string_strchr(unsafe.Pointer((*int8)(unsafe.Pointer(&[5]int8{'a', 'a', 'b', 'b', '\x00'}))), 5, uint64(align)))
			var q *int8 = libc.Strchr(p, 'b')
			if uintptr(unsafe.Pointer(q)) == uintptr(unsafe.Pointer(nil)) {
				common.T_printf((*int8)(unsafe.Pointer(&[121]int8{'/', 'U', 's', 'e', 'r', 's', '/', 'x', 'u', 's', 'h', 'i', 'w', 'e', 'i', '/', 'w', 'o', 'r', 'k', '/', 'l', 'i', 'b', 'c', '/', 't', 'e', 's', 't', '/', 's', 'r', 'c', '/', 'f', 'u', 'n', 'c', 't', 'i', 'o', 'n', 'a', 'l', '/', 's', 't', 'r', 'i', 'n', 'g', '_', 's', 't', 'r', 'c', 'h', 'r', '.', 'c', ':', '6', '2', ':', ' ', 's', 't', 'r', 'c', 'h', 'r', '(', '%', 's', ',', '%', 's', ')', ' ', 'w', 'i', 't', 'h', ' ', 'a', 'l', 'i', 'g', 'n', '=', '%', 'd', ' ', 'r', 'e', 't', 'u', 'r', 'n', 'e', 'd', ' ', '0', ',', ' ', 'w', 'a', 'n', 't', 'e', 'd', ' ', 's', 't', 'r', '+', '%', 'd', '\n', '\x00'})), (*int8)(unsafe.Pointer(&[7]int8{'"', 'a', 'a', 'b', 'b', '"', '\x00'})), (*int8)(unsafe.Pointer(&[4]int8{'\'', 'b', '\'', '\x00'})), align, int32(2))
			} else if int64(uintptr(unsafe.Pointer(q))-uintptr(unsafe.Pointer(p))) != int64(2) {
				common.T_printf((*int8)(unsafe.Pointer(&[126]int8{'/', 'U', 's', 'e', 'r', 's', '/', 'x', 'u', 's', 'h', 'i', 'w', 'e', 'i', '/', 'w', 'o', 'r', 'k', '/', 'l', 'i', 'b', 'c', '/', 't', 'e', 's', 't', '/', 's', 'r', 'c', '/', 'f', 'u', 'n', 'c', 't', 'i', 'o', 'n', 'a', 'l', '/', 's', 't', 'r', 'i', 'n', 'g', '_', 's', 't', 'r', 'c', 'h', 'r', '.', 'c', ':', '6', '2', ':', ' ', 's', 't', 'r', 'c', 'h', 'r', '(', '%', 's', ',', '%', 's', ')', ' ', 'w', 'i', 't', 'h', ' ', 'a', 'l', 'i', 'g', 'n', '=', '%', 'd', ' ', 'r', 'e', 't', 'u', 'r', 'n', 'e', 'd', ' ', 's', 't', 'r', '+', '%', 'd', ',', ' ', 'w', 'a', 'n', 't', 'e', 'd', ' ', 's', 't', 'r', '+', '%', 'd', '\n', '\x00'})), (*int8)(unsafe.Pointer(&[7]int8{'"', 'a', 'a', 'b', 'b', '"', '\x00'})), (*int8)(unsafe.Pointer(&[4]int8{'\'', 'b', '\'', '\x00'})), align, uintptr(unsafe.Pointer(q))-uintptr(unsafe.Pointer(p)), int32(2))
			}
		}
	}
	{
		var align int32
		for align = int32(0); align < int32(8); align++ {
			var p *int8 = (*int8)(_cgos_aligncpy_string_strchr(unsafe.Pointer((*int8)(unsafe.Pointer(&[6]int8{'a', 'a', 'a', 'b', 'b', '\x00'}))), 6, uint64(align)))
			var q *int8 = libc.Strchr(p, 'b')
			if uintptr(unsafe.Pointer(q)) == uintptr(unsafe.Pointer(nil)) {
				common.T_printf((*int8)(unsafe.Pointer(&[121]int8{'/', 'U', 's', 'e', 'r', 's', '/', 'x', 'u', 's', 'h', 'i', 'w', 'e', 'i', '/', 'w', 'o', 'r', 'k', '/', 'l', 'i', 'b', 'c', '/', 't', 'e', 's', 't', '/', 's', 'r', 'c', '/', 'f', 'u', 'n', 'c', 't', 'i', 'o', 'n', 'a', 'l', '/', 's', 't', 'r', 'i', 'n', 'g', '_', 's', 't', 'r', 'c', 'h', 'r', '.', 'c', ':', '6', '3', ':', ' ', 's', 't', 'r', 'c', 'h', 'r', '(', '%', 's', ',', '%', 's', ')', ' ', 'w', 'i', 't', 'h', ' ', 'a', 'l', 'i', 'g', 'n', '=', '%', 'd', ' ', 'r', 'e', 't', 'u', 'r', 'n', 'e', 'd', ' ', '0', ',', ' ', 'w', 'a', 'n', 't', 'e', 'd', ' ', 's', 't', 'r', '+', '%', 'd', '\n', '\x00'})), (*int8)(unsafe.Pointer(&[8]int8{'"', 'a', 'a', 'a', 'b', 'b', '"', '\x00'})), (*int8)(unsafe.Pointer(&[4]int8{'\'', 'b', '\'', '\x00'})), align, int32(3))
			} else if int64(uintptr(unsafe.Pointer(q))-uintptr(unsafe.Pointer(p))) != int64(3) {
				common.T_printf((*int8)(unsafe.Pointer(&[126]int8{'/', 'U', 's', 'e', 'r', 's', '/', 'x', 'u', 's', 'h', 'i', 'w', 'e', 'i', '/', 'w', 'o', 'r', 'k', '/', 'l', 'i', 'b', 'c', '/', 't', 'e', 's', 't', '/', 's', 'r', 'c', '/', 'f', 'u', 'n', 'c', 't', 'i', 'o', 'n', 'a', 'l', '/', 's', 't', 'r', 'i', 'n', 'g', '_', 's', 't', 'r', 'c', 'h', 'r', '.', 'c', ':', '6', '3', ':', ' ', 's', 't', 'r', 'c', 'h', 'r', '(', '%', 's', ',', '%', 's', ')', ' ', 'w', 'i', 't', 'h', ' ', 'a', 'l', 'i', 'g', 'n', '=', '%', 'd', ' ', 'r', 'e', 't', 'u', 'r', 'n', 'e', 'd', ' ', 's', 't', 'r', '+', '%', 'd', ',', ' ', 'w', 'a', 'n', 't', 'e', 'd', ' ', 's', 't', 'r', '+', '%', 'd', '\n', '\x00'})), (*int8)(unsafe.Pointer(&[8]int8{'"', 'a', 'a', 'a', 'b', 'b', '"', '\x00'})), (*int8)(unsafe.Pointer(&[4]int8{'\'', 'b', '\'', '\x00'})), align, uintptr(unsafe.Pointer(q))-uintptr(unsafe.Pointer(p)), int32(3))
			}
		}
	}
	{
		var align int32
		for align = int32(0); align < int32(8); align++ {
			var p *int8 = (*int8)(_cgos_aligncpy_string_strchr(unsafe.Pointer((*int8)(unsafe.Pointer(&[7]int8{'a', 'a', 'a', 'a', 'b', 'b', '\x00'}))), 7, uint64(align)))
			var q *int8 = libc.Strchr(p, 'b')
			if uintptr(unsafe.Pointer(q)) == uintptr(unsafe.Pointer(nil)) {
				common.T_printf((*int8)(unsafe.Pointer(&[121]int8{'/', 'U', 's', 'e', 'r', 's', '/', 'x', 'u', 's', 'h', 'i', 'w', 'e', 'i', '/', 'w', 'o', 'r', 'k', '/', 'l', 'i', 'b', 'c', '/', 't', 'e', 's', 't', '/', 's', 'r', 'c', '/', 'f', 'u', 'n', 'c', 't', 'i', 'o', 'n', 'a', 'l', '/', 's', 't', 'r', 'i', 'n', 'g', '_', 's', 't', 'r', 'c', 'h', 'r', '.', 'c', ':', '6', '4', ':', ' ', 's', 't', 'r', 'c', 'h', 'r', '(', '%', 's', ',', '%', 's', ')', ' ', 'w', 'i', 't', 'h', ' ', 'a', 'l', 'i', 'g', 'n', '=', '%', 'd', ' ', 'r', 'e', 't', 'u', 'r', 'n', 'e', 'd', ' ', '0', ',', ' ', 'w', 'a', 'n', 't', 'e', 'd', ' ', 's', 't', 'r', '+', '%', 'd', '\n', '\x00'})), (*int8)(unsafe.Pointer(&[9]int8{'"', 'a', 'a', 'a', 'a', 'b', 'b', '"', '\x00'})), (*int8)(unsafe.Pointer(&[4]int8{'\'', 'b', '\'', '\x00'})), align, int32(4))
			} else if int64(uintptr(unsafe.Pointer(q))-uintptr(unsafe.Pointer(p))) != int64(4) {
				common.T_printf((*int8)(unsafe.Pointer(&[126]int8{'/', 'U', 's', 'e', 'r', 's', '/', 'x', 'u', 's', 'h', 'i', 'w', 'e', 'i', '/', 'w', 'o', 'r', 'k', '/', 'l', 'i', 'b', 'c', '/', 't', 'e', 's', 't', '/', 's', 'r', 'c', '/', 'f', 'u', 'n', 'c', 't', 'i', 'o', 'n', 'a', 'l', '/', 's', 't', 'r', 'i', 'n', 'g', '_', 's', 't', 'r', 'c', 'h', 'r', '.', 'c', ':', '6', '4', ':', ' ', 's', 't', 'r', 'c', 'h', 'r', '(', '%', 's', ',', '%', 's', ')', ' ', 'w', 'i', 't', 'h', ' ', 'a', 'l', 'i', 'g', 'n', '=', '%', 'd', ' ', 'r', 'e', 't', 'u', 'r', 'n', 'e', 'd', ' ', 's', 't', 'r', '+', '%', 'd', ',', ' ', 'w', 'a', 'n', 't', 'e', 'd', ' ', 's', 't', 'r', '+', '%', 'd', '\n', '\x00'})), (*int8)(unsafe.Pointer(&[9]int8{'"', 'a', 'a', 'a', 'a', 'b', 'b', '"', '\x00'})), (*int8)(unsafe.Pointer(&[4]int8{'\'', 'b', '\'', '\x00'})), align, uintptr(unsafe.Pointer(q))-uintptr(unsafe.Pointer(p)), int32(4))
			}
		}
	}
	{
		var align int32
		for align = int32(0); align < int32(8); align++ {
			var p *int8 = (*int8)(_cgos_aligncpy_string_strchr(unsafe.Pointer((*int8)(unsafe.Pointer(&[8]int8{'a', 'a', 'a', 'a', 'a', 'b', 'b', '\x00'}))), 8, uint64(align)))
			var q *int8 = libc.Strchr(p, 'b')
			if uintptr(unsafe.Pointer(q)) == uintptr(unsafe.Pointer(nil)) {
				common.T_printf((*int8)(unsafe.Pointer(&[121]int8{'/', 'U', 's', 'e', 'r', 's', '/', 'x', 'u', 's', 'h', 'i', 'w', 'e', 'i', '/', 'w', 'o', 'r', 'k', '/', 'l', 'i', 'b', 'c', '/', 't', 'e', 's', 't', '/', 's', 'r', 'c', '/', 'f', 'u', 'n', 'c', 't', 'i', 'o', 'n', 'a', 'l', '/', 's', 't', 'r', 'i', 'n', 'g', '_', 's', 't', 'r', 'c', 'h', 'r', '.', 'c', ':', '6', '5', ':', ' ', 's', 't', 'r', 'c', 'h', 'r', '(', '%', 's', ',', '%', 's', ')', ' ', 'w', 'i', 't', 'h', ' ', 'a', 'l', 'i', 'g', 'n', '=', '%', 'd', ' ', 'r', 'e', 't', 'u', 'r', 'n', 'e', 'd', ' ', '0', ',', ' ', 'w', 'a', 'n', 't', 'e', 'd', ' ', 's', 't', 'r', '+', '%', 'd', '\n', '\x00'})), (*int8)(unsafe.Pointer(&[10]int8{'"', 'a', 'a', 'a', 'a', 'a', 'b', 'b', '"', '\x00'})), (*int8)(unsafe.Pointer(&[4]int8{'\'', 'b', '\'', '\x00'})), align, int32(5))
			} else if int64(uintptr(unsafe.Pointer(q))-uintptr(unsafe.Pointer(p))) != int64(5) {
				common.T_printf((*int8)(unsafe.Pointer(&[126]int8{'/', 'U', 's', 'e', 'r', 's', '/', 'x', 'u', 's', 'h', 'i', 'w', 'e', 'i', '/', 'w', 'o', 'r', 'k', '/', 'l', 'i', 'b', 'c', '/', 't', 'e', 's', 't', '/', 's', 'r', 'c', '/', 'f', 'u', 'n', 'c', 't', 'i', 'o', 'n', 'a', 'l', '/', 's', 't', 'r', 'i', 'n', 'g', '_', 's', 't', 'r', 'c', 'h', 'r', '.', 'c', ':', '6', '5', ':', ' ', 's', 't', 'r', 'c', 'h', 'r', '(', '%', 's', ',', '%', 's', ')', ' ', 'w', 'i', 't', 'h', ' ', 'a', 'l', 'i', 'g', 'n', '=', '%', 'd', ' ', 'r', 'e', 't', 'u', 'r', 'n', 'e', 'd', ' ', 's', 't', 'r', '+', '%', 'd', ',', ' ', 'w', 'a', 'n', 't', 'e', 'd', ' ', 's', 't', 'r', '+', '%', 'd', '\n', '\x00'})), (*int8)(unsafe.Pointer(&[10]int8{'"', 'a', 'a', 'a', 'a', 'a', 'b', 'b', '"', '\x00'})), (*int8)(unsafe.Pointer(&[4]int8{'\'', 'b', '\'', '\x00'})), align, uintptr(unsafe.Pointer(q))-uintptr(unsafe.Pointer(p)), int32(5))
			}
		}
	}
	{
		var align int32
		for align = int32(0); align < int32(8); align++ {
			var p *int8 = (*int8)(_cgos_aligncpy_string_strchr(unsafe.Pointer((*int8)(unsafe.Pointer(&[9]int8{'a', 'a', 'a', 'a', 'a', 'a', 'b', 'b', '\x00'}))), 9, uint64(align)))
			var q *int8 = libc.Strchr(p, 'b')
			if uintptr(unsafe.Pointer(q)) == uintptr(unsafe.Pointer(nil)) {
				common.T_printf((*int8)(unsafe.Pointer(&[121]int8{'/', 'U', 's', 'e', 'r', 's', '/', 'x', 'u', 's', 'h', 'i', 'w', 'e', 'i', '/', 'w', 'o', 'r', 'k', '/', 'l', 'i', 'b', 'c', '/', 't', 'e', 's', 't', '/', 's', 'r', 'c', '/', 'f', 'u', 'n', 'c', 't', 'i', 'o', 'n', 'a', 'l', '/', 's', 't', 'r', 'i', 'n', 'g', '_', 's', 't', 'r', 'c', 'h', 'r', '.', 'c', ':', '6', '6', ':', ' ', 's', 't', 'r', 'c', 'h', 'r', '(', '%', 's', ',', '%', 's', ')', ' ', 'w', 'i', 't', 'h', ' ', 'a', 'l', 'i', 'g', 'n', '=', '%', 'd', ' ', 'r', 'e', 't', 'u', 'r', 'n', 'e', 'd', ' ', '0', ',', ' ', 'w', 'a', 'n', 't', 'e', 'd', ' ', 's', 't', 'r', '+', '%', 'd', '\n', '\x00'})), (*int8)(unsafe.Pointer(&[11]int8{'"', 'a', 'a', 'a', 'a', 'a', 'a', 'b', 'b', '"', '\x00'})), (*int8)(unsafe.Pointer(&[4]int8{'\'', 'b', '\'', '\x00'})), align, int32(6))
			} else if int64(uintptr(unsafe.Pointer(q))-uintptr(unsafe.Pointer(p))) != int64(6) {
				common.T_printf((*int8)(unsafe.Pointer(&[126]int8{'/', 'U', 's', 'e', 'r', 's', '/', 'x', 'u', 's', 'h', 'i', 'w', 'e', 'i', '/', 'w', 'o', 'r', 'k', '/', 'l', 'i', 'b', 'c', '/', 't', 'e', 's', 't', '/', 's', 'r', 'c', '/', 'f', 'u', 'n', 'c', 't', 'i', 'o', 'n', 'a', 'l', '/', 's', 't', 'r', 'i', 'n', 'g', '_', 's', 't', 'r', 'c', 'h', 'r', '.', 'c', ':', '6', '6', ':', ' ', 's', 't', 'r', 'c', 'h', 'r', '(', '%', 's', ',', '%', 's', ')', ' ', 'w', 'i', 't', 'h', ' ', 'a', 'l', 'i', 'g', 'n', '=', '%', 'd', ' ', 'r', 'e', 't', 'u', 'r', 'n', 'e', 'd', ' ', 's', 't', 'r', '+', '%', 'd', ',', ' ', 'w', 'a', 'n', 't', 'e', 'd', ' ', 's', 't', 'r', '+', '%', 'd', '\n', '\x00'})), (*int8)(unsafe.Pointer(&[11]int8{'"', 'a', 'a', 'a', 'a', 'a', 'a', 'b', 'b', '"', '\x00'})), (*int8)(unsafe.Pointer(&[4]int8{'\'', 'b', '\'', '\x00'})), align, uintptr(unsafe.Pointer(q))-uintptr(unsafe.Pointer(p)), int32(6))
			}
		}
	}
	{
		var align int32
		for align = int32(0); align < int32(8); align++ {
			var p *int8 = (*int8)(_cgos_aligncpy_string_strchr(unsafe.Pointer((*int8)(unsafe.Pointer(&[8]int8{'a', 'b', 'c', ' ', 'a', 'b', 'c', '\x00'}))), 8, uint64(align)))
			var q *int8 = libc.Strchr(p, 'c')
			if uintptr(unsafe.Pointer(q)) == uintptr(unsafe.Pointer(nil)) {
				common.T_printf((*int8)(unsafe.Pointer(&[121]int8{'/', 'U', 's', 'e', 'r', 's', '/', 'x', 'u', 's', 'h', 'i', 'w', 'e', 'i', '/', 'w', 'o', 'r', 'k', '/', 'l', 'i', 'b', 'c', '/', 't', 'e', 's', 't', '/', 's', 'r', 'c', '/', 'f', 'u', 'n', 'c', 't', 'i', 'o', 'n', 'a', 'l', '/', 's', 't', 'r', 'i', 'n', 'g', '_', 's', 't', 'r', 'c', 'h', 'r', '.', 'c', ':', '6', '7', ':', ' ', 's', 't', 'r', 'c', 'h', 'r', '(', '%', 's', ',', '%', 's', ')', ' ', 'w', 'i', 't', 'h', ' ', 'a', 'l', 'i', 'g', 'n', '=', '%', 'd', ' ', 'r', 'e', 't', 'u', 'r', 'n', 'e', 'd', ' ', '0', ',', ' ', 'w', 'a', 'n', 't', 'e', 'd', ' ', 's', 't', 'r', '+', '%', 'd', '\n', '\x00'})), (*int8)(unsafe.Pointer(&[10]int8{'"', 'a', 'b', 'c', ' ', 'a', 'b', 'c', '"', '\x00'})), (*int8)(unsafe.Pointer(&[4]int8{'\'', 'c', '\'', '\x00'})), align, int32(2))
			} else if int64(uintptr(unsafe.Pointer(q))-uintptr(unsafe.Pointer(p))) != int64(2) {
				common.T_printf((*int8)(unsafe.Pointer(&[126]int8{'/', 'U', 's', 'e', 'r', 's', '/', 'x', 'u', 's', 'h', 'i', 'w', 'e', 'i', '/', 'w', 'o', 'r', 'k', '/', 'l', 'i', 'b', 'c', '/', 't', 'e', 's', 't', '/', 's', 'r', 'c', '/', 'f', 'u', 'n', 'c', 't', 'i', 'o', 'n', 'a', 'l', '/', 's', 't', 'r', 'i', 'n', 'g', '_', 's', 't', 'r', 'c', 'h', 'r', '.', 'c', ':', '6', '7', ':', ' ', 's', 't', 'r', 'c', 'h', 'r', '(', '%', 's', ',', '%', 's', ')', ' ', 'w', 'i', 't', 'h', ' ', 'a', 'l', 'i', 'g', 'n', '=', '%', 'd', ' ', 'r', 'e', 't', 'u', 'r', 'n', 'e', 'd', ' ', 's', 't', 'r', '+', '%', 'd', ',', ' ', 'w', 'a', 'n', 't', 'e', 'd', ' ', 's', 't', 'r', '+', '%', 'd', '\n', '\x00'})), (*int8)(unsafe.Pointer(&[10]int8{'"', 'a', 'b', 'c', ' ', 'a', 'b', 'c', '"', '\x00'})), (*int8)(unsafe.Pointer(&[4]int8{'\'', 'c', '\'', '\x00'})), align, uintptr(unsafe.Pointer(q))-uintptr(unsafe.Pointer(p)), int32(2))
			}
		}
	}
	{
		var align int32
		for align = int32(0); align < int32(8); align++ {
			var p *int8 = (*int8)(_cgos_aligncpy_string_strchr(unsafe.Pointer((*int8)(unsafe.Pointer(&s))), 256, uint64(align)))
			var q *int8 = libc.Strchr(p, int32(1))
			if uintptr(unsafe.Pointer(q)) == uintptr(unsafe.Pointer(nil)) {
				common.T_printf((*int8)(unsafe.Pointer(&[121]int8{'/', 'U', 's', 'e', 'r', 's', '/', 'x', 'u', 's', 'h', 'i', 'w', 'e', 'i', '/', 'w', 'o', 'r', 'k', '/', 'l', 'i', 'b', 'c', '/', 't', 'e', 's', 't', '/', 's', 'r', 'c', '/', 'f', 'u', 'n', 'c', 't', 'i', 'o', 'n', 'a', 'l', '/', 's', 't', 'r', 'i', 'n', 'g', '_', 's', 't', 'r', 'c', 'h', 'r', '.', 'c', ':', '6', '8', ':', ' ', 's', 't', 'r', 'c', 'h', 'r', '(', '%', 's', ',', '%', 's', ')', ' ', 'w', 'i', 't', 'h', ' ', 'a', 'l', 'i', 'g', 'n', '=', '%', 'd', ' ', 'r', 'e', 't', 'u', 'r', 'n', 'e', 'd', ' ', '0', ',', ' ', 'w', 'a', 'n', 't', 'e', 'd', ' ', 's', 't', 'r', '+', '%', 'd', '\n', '\x00'})), (*int8)(unsafe.Pointer(&[2]int8{'s', '\x00'})), (*int8)(unsafe.Pointer(&[2]int8{'1', '\x00'})), align, int32(0))
			} else if int64(uintptr(unsafe.Pointer(q))-uintptr(unsafe.Pointer(p))) != int64(0) {
				common.T_printf((*int8)(unsafe.Pointer(&[126]int8{'/', 'U', 's', 'e', 'r', 's', '/', 'x', 'u', 's', 'h', 'i', 'w', 'e', 'i', '/', 'w', 'o', 'r', 'k', '/', 'l', 'i', 'b', 'c', '/', 't', 'e', 's', 't', '/', 's', 'r', 'c', '/', 'f', 'u', 'n', 'c', 't', 'i', 'o', 'n', 'a', 'l', '/', 's', 't', 'r', 'i', 'n', 'g', '_', 's', 't', 'r', 'c', 'h', 'r', '.', 'c', ':', '6', '8', ':', ' ', 's', 't', 'r', 'c', 'h', 'r', '(', '%', 's', ',', '%', 's', ')', ' ', 'w', 'i', 't', 'h', ' ', 'a', 'l', 'i', 'g', 'n', '=', '%', 'd', ' ', 'r', 'e', 't', 'u', 'r', 'n', 'e', 'd', ' ', 's', 't', 'r', '+', '%', 'd', ',', ' ', 'w', 'a', 'n', 't', 'e', 'd', ' ', 's', 't', 'r', '+', '%', 'd', '\n', '\x00'})), (*int8)(unsafe.Pointer(&[2]int8{'s', '\x00'})), (*int8)(unsafe.Pointer(&[2]int8{'1', '\x00'})), align, uintptr(unsafe.Pointer(q))-uintptr(unsafe.Pointer(p)), int32(0))
			}
		}
	}
	{
		var align int32
		for align = int32(0); align < int32(8); align++ {
			var p *int8 = (*int8)(_cgos_aligncpy_string_strchr(unsafe.Pointer((*int8)(unsafe.Pointer(&s))), 256, uint64(align)))
			var q *int8 = libc.Strchr(p, int32(2))
			if uintptr(unsafe.Pointer(q)) == uintptr(unsafe.Pointer(nil)) {
				common.T_printf((*int8)(unsafe.Pointer(&[121]int8{'/', 'U', 's', 'e', 'r', 's', '/', 'x', 'u', 's', 'h', 'i', 'w', 'e', 'i', '/', 'w', 'o', 'r', 'k', '/', 'l', 'i', 'b', 'c', '/', 't', 'e', 's', 't', '/', 's', 'r', 'c', '/', 'f', 'u', 'n', 'c', 't', 'i', 'o', 'n', 'a', 'l', '/', 's', 't', 'r', 'i', 'n', 'g', '_', 's', 't', 'r', 'c', 'h', 'r', '.', 'c', ':', '6', '9', ':', ' ', 's', 't', 'r', 'c', 'h', 'r', '(', '%', 's', ',', '%', 's', ')', ' ', 'w', 'i', 't', 'h', ' ', 'a', 'l', 'i', 'g', 'n', '=', '%', 'd', ' ', 'r', 'e', 't', 'u', 'r', 'n', 'e', 'd', ' ', '0', ',', ' ', 'w', 'a', 'n', 't', 'e', 'd', ' ', 's', 't', 'r', '+', '%', 'd', '\n', '\x00'})), (*int8)(unsafe.Pointer(&[2]int8{'s', '\x00'})), (*int8)(unsafe.Pointer(&[2]int8{'2', '\x00'})), align, int32(1))
			} else if int64(uintptr(unsafe.Pointer(q))-uintptr(unsafe.Pointer(p))) != int64(1) {
				common.T_printf((*int8)(unsafe.Pointer(&[126]int8{'/', 'U', 's', 'e', 'r', 's', '/', 'x', 'u', 's', 'h', 'i', 'w', 'e', 'i', '/', 'w', 'o', 'r', 'k', '/', 'l', 'i', 'b', 'c', '/', 't', 'e', 's', 't', '/', 's', 'r', 'c', '/', 'f', 'u', 'n', 'c', 't', 'i', 'o', 'n', 'a', 'l', '/', 's', 't', 'r', 'i', 'n', 'g', '_', 's', 't', 'r', 'c', 'h', 'r', '.', 'c', ':', '6', '9', ':', ' ', 's', 't', 'r', 'c', 'h', 'r', '(', '%', 's', ',', '%', 's', ')', ' ', 'w', 'i', 't', 'h', ' ', 'a', 'l', 'i', 'g', 'n', '=', '%', 'd', ' ', 'r', 'e', 't', 'u', 'r', 'n', 'e', 'd', ' ', 's', 't', 'r', '+', '%', 'd', ',', ' ', 'w', 'a', 'n', 't', 'e', 'd', ' ', 's', 't', 'r', '+', '%', 'd', '\n', '\x00'})), (*int8)(unsafe.Pointer(&[2]int8{'s', '\x00'})), (*int8)(unsafe.Pointer(&[2]int8{'2', '\x00'})), align, uintptr(unsafe.Pointer(q))-uintptr(unsafe.Pointer(p)), int32(1))
			}
		}
	}
	{
		var align int32
		for align = int32(0); align < int32(8); align++ {
			var p *int8 = (*int8)(_cgos_aligncpy_string_strchr(unsafe.Pointer((*int8)(unsafe.Pointer(&s))), 256, uint64(align)))
			var q *int8 = libc.Strchr(p, int32(10))
			if uintptr(unsafe.Pointer(q)) == uintptr(unsafe.Pointer(nil)) {
				common.T_printf((*int8)(unsafe.Pointer(&[121]int8{'/', 'U', 's', 'e', 'r', 's', '/', 'x', 'u', 's', 'h', 'i', 'w', 'e', 'i', '/', 'w', 'o', 'r', 'k', '/', 'l', 'i', 'b', 'c', '/', 't', 'e', 's', 't', '/', 's', 'r', 'c', '/', 'f', 'u', 'n', 'c', 't', 'i', 'o', 'n', 'a', 'l', '/', 's', 't', 'r', 'i', 'n', 'g', '_', 's', 't', 'r', 'c', 'h', 'r', '.', 'c', ':', '7', '0', ':', ' ', 's', 't', 'r', 'c', 'h', 'r', '(', '%', 's', ',', '%', 's', ')', ' ', 'w', 'i', 't', 'h', ' ', 'a', 'l', 'i', 'g', 'n', '=', '%', 'd', ' ', 'r', 'e', 't', 'u', 'r', 'n', 'e', 'd', ' ', '0', ',', ' ', 'w', 'a', 'n', 't', 'e', 'd', ' ', 's', 't', 'r', '+', '%', 'd', '\n', '\x00'})), (*int8)(unsafe.Pointer(&[2]int8{'s', '\x00'})), (*int8)(unsafe.Pointer(&[3]int8{'1', '0', '\x00'})), align, int32(9))
			} else if int64(uintptr(unsafe.Pointer(q))-uintptr(unsafe.Pointer(p))) != int64(9) {
				common.T_printf((*int8)(unsafe.Pointer(&[126]int8{'/', 'U', 's', 'e', 'r', 's', '/', 'x', 'u', 's', 'h', 'i', 'w', 'e', 'i', '/', 'w', 'o', 'r', 'k', '/', 'l', 'i', 'b', 'c', '/', 't', 'e', 's', 't', '/', 's', 'r', 'c', '/', 'f', 'u', 'n', 'c', 't', 'i', 'o', 'n', 'a', 'l', '/', 's', 't', 'r', 'i', 'n', 'g', '_', 's', 't', 'r', 'c', 'h', 'r', '.', 'c', ':', '7', '0', ':', ' ', 's', 't', 'r', 'c', 'h', 'r', '(', '%', 's', ',', '%', 's', ')', ' ', 'w', 'i', 't', 'h', ' ', 'a', 'l', 'i', 'g', 'n', '=', '%', 'd', ' ', 'r', 'e', 't', 'u', 'r', 'n', 'e', 'd', ' ', 's', 't', 'r', '+', '%', 'd', ',', ' ', 'w', 'a', 'n', 't', 'e', 'd', ' ', 's', 't', 'r', '+', '%', 'd', '\n', '\x00'})), (*int8)(unsafe.Pointer(&[2]int8{'s', '\x00'})), (*int8)(unsafe.Pointer(&[3]int8{'1', '0', '\x00'})), align, uintptr(unsafe.Pointer(q))-uintptr(unsafe.Pointer(p)), int32(9))
			}
		}
	}
	{
		var align int32
		for align = int32(0); align < int32(8); align++ {
			var p *int8 = (*int8)(_cgos_aligncpy_string_strchr(unsafe.Pointer((*int8)(unsafe.Pointer(&s))), 256, uint64(align)))
			var q *int8 = libc.Strchr(p, int32(11))
			if uintptr(unsafe.Pointer(q)) == uintptr(unsafe.Pointer(nil)) {
				common.T_printf((*int8)(unsafe.Pointer(&[121]int8{'/', 'U', 's', 'e', 'r', 's', '/', 'x', 'u', 's', 'h', 'i', 'w', 'e', 'i', '/', 'w', 'o', 'r', 'k', '/', 'l', 'i', 'b', 'c', '/', 't', 'e', 's', 't', '/', 's', 'r', 'c', '/', 'f', 'u', 'n', 'c', 't', 'i', 'o', 'n', 'a', 'l', '/', 's', 't', 'r', 'i', 'n', 'g', '_', 's', 't', 'r', 'c', 'h', 'r', '.', 'c', ':', '7', '1', ':', ' ', 's', 't', 'r', 'c', 'h', 'r', '(', '%', 's', ',', '%', 's', ')', ' ', 'w', 'i', 't', 'h', ' ', 'a', 'l', 'i', 'g', 'n', '=', '%', 'd', ' ', 'r', 'e', 't', 'u', 'r', 'n', 'e', 'd', ' ', '0', ',', ' ', 'w', 'a', 'n', 't', 'e', 'd', ' ', 's', 't', 'r', '+', '%', 'd', '\n', '\x00'})), (*int8)(unsafe.Pointer(&[2]int8{'s', '\x00'})), (*int8)(unsafe.Pointer(&[3]int8{'1', '1', '\x00'})), align, int32(10))
			} else if int64(uintptr(unsafe.Pointer(q))-uintptr(unsafe.Pointer(p))) != int64(10) {
				common.T_printf((*int8)(unsafe.Pointer(&[126]int8{'/', 'U', 's', 'e', 'r', 's', '/', 'x', 'u', 's', 'h', 'i', 'w', 'e', 'i', '/', 'w', 'o', 'r', 'k', '/', 'l', 'i', 'b', 'c', '/', 't', 'e', 's', 't', '/', 's', 'r', 'c', '/', 'f', 'u', 'n', 'c', 't', 'i', 'o', 'n', 'a', 'l', '/', 's', 't', 'r', 'i', 'n', 'g', '_', 's', 't', 'r', 'c', 'h', 'r', '.', 'c', ':', '7', '1', ':', ' ', 's', 't', 'r', 'c', 'h', 'r', '(', '%', 's', ',', '%', 's', ')', ' ', 'w', 'i', 't', 'h', ' ', 'a', 'l', 'i', 'g', 'n', '=', '%', 'd', ' ', 'r', 'e', 't', 'u', 'r', 'n', 'e', 'd', ' ', 's', 't', 'r', '+', '%', 'd', ',', ' ', 'w', 'a', 'n', 't', 'e', 'd', ' ', 's', 't', 'r', '+', '%', 'd', '\n', '\x00'})), (*int8)(unsafe.Pointer(&[2]int8{'s', '\x00'})), (*int8)(unsafe.Pointer(&[3]int8{'1', '1', '\x00'})), align, uintptr(unsafe.Pointer(q))-uintptr(unsafe.Pointer(p)), int32(10))
			}
		}
	}
	{
		var align int32
		for align = int32(0); align < int32(8); align++ {
			var p *int8 = (*int8)(_cgos_aligncpy_string_strchr(unsafe.Pointer((*int8)(unsafe.Pointer(&s))), 256, uint64(align)))
			var q *int8 = libc.Strchr(p, int32(127))
			if uintptr(unsafe.Pointer(q)) == uintptr(unsafe.Pointer(nil)) {
				common.T_printf((*int8)(unsafe.Pointer(&[121]int8{'/', 'U', 's', 'e', 'r', 's', '/', 'x', 'u', 's', 'h', 'i', 'w', 'e', 'i', '/', 'w', 'o', 'r', 'k', '/', 'l', 'i', 'b', 'c', '/', 't', 'e', 's', 't', '/', 's', 'r', 'c', '/', 'f', 'u', 'n', 'c', 't', 'i', 'o', 'n', 'a', 'l', '/', 's', 't', 'r', 'i', 'n', 'g', '_', 's', 't', 'r', 'c', 'h', 'r', '.', 'c', ':', '7', '2', ':', ' ', 's', 't', 'r', 'c', 'h', 'r', '(', '%', 's', ',', '%', 's', ')', ' ', 'w', 'i', 't', 'h', ' ', 'a', 'l', 'i', 'g', 'n', '=', '%', 'd', ' ', 'r', 'e', 't', 'u', 'r', 'n', 'e', 'd', ' ', '0', ',', ' ', 'w', 'a', 'n', 't', 'e', 'd', ' ', 's', 't', 'r', '+', '%', 'd', '\n', '\x00'})), (*int8)(unsafe.Pointer(&[2]int8{'s', '\x00'})), (*int8)(unsafe.Pointer(&[4]int8{'1', '2', '7', '\x00'})), align, int32(126))
			} else if int64(uintptr(unsafe.Pointer(q))-uintptr(unsafe.Pointer(p))) != int64(126) {
				common.T_printf((*int8)(unsafe.Pointer(&[126]int8{'/', 'U', 's', 'e', 'r', 's', '/', 'x', 'u', 's', 'h', 'i', 'w', 'e', 'i', '/', 'w', 'o', 'r', 'k', '/', 'l', 'i', 'b', 'c', '/', 't', 'e', 's', 't', '/', 's', 'r', 'c', '/', 'f', 'u', 'n', 'c', 't', 'i', 'o', 'n', 'a', 'l', '/', 's', 't', 'r', 'i', 'n', 'g', '_', 's', 't', 'r', 'c', 'h', 'r', '.', 'c', ':', '7', '2', ':', ' ', 's', 't', 'r', 'c', 'h', 'r', '(', '%', 's', ',', '%', 's', ')', ' ', 'w', 'i', 't', 'h', ' ', 'a', 'l', 'i', 'g', 'n', '=', '%', 'd', ' ', 'r', 'e', 't', 'u', 'r', 'n', 'e', 'd', ' ', 's', 't', 'r', '+', '%', 'd', ',', ' ', 'w', 'a', 'n', 't', 'e', 'd', ' ', 's', 't', 'r', '+', '%', 'd', '\n', '\x00'})), (*int8)(unsafe.Pointer(&[2]int8{'s', '\x00'})), (*int8)(unsafe.Pointer(&[4]int8{'1', '2', '7', '\x00'})), align, uintptr(unsafe.Pointer(q))-uintptr(unsafe.Pointer(p)), int32(126))
			}
		}
	}
	{
		var align int32
		for align = int32(0); align < int32(8); align++ {
			var p *int8 = (*int8)(_cgos_aligncpy_string_strchr(unsafe.Pointer((*int8)(unsafe.Pointer(&s))), 256, uint64(align)))
			var q *int8 = libc.Strchr(p, int32(128))
			if uintptr(unsafe.Pointer(q)) == uintptr(unsafe.Pointer(nil)) {
				common.T_printf((*int8)(unsafe.Pointer(&[121]int8{'/', 'U', 's', 'e', 'r', 's', '/', 'x', 'u', 's', 'h', 'i', 'w', 'e', 'i', '/', 'w', 'o', 'r', 'k', '/', 'l', 'i', 'b', 'c', '/', 't', 'e', 's', 't', '/', 's', 'r', 'c', '/', 'f', 'u', 'n', 'c', 't', 'i', 'o', 'n', 'a', 'l', '/', 's', 't', 'r', 'i', 'n', 'g', '_', 's', 't', 'r', 'c', 'h', 'r', '.', 'c', ':', '7', '3', ':', ' ', 's', 't', 'r', 'c', 'h', 'r', '(', '%', 's', ',', '%', 's', ')', ' ', 'w', 'i', 't', 'h', ' ', 'a', 'l', 'i', 'g', 'n', '=', '%', 'd', ' ', 'r', 'e', 't', 'u', 'r', 'n', 'e', 'd', ' ', '0', ',', ' ', 'w', 'a', 'n', 't', 'e', 'd', ' ', 's', 't', 'r', '+', '%', 'd', '\n', '\x00'})), (*int8)(unsafe.Pointer(&[2]int8{'s', '\x00'})), (*int8)(unsafe.Pointer(&[4]int8{'1', '2', '8', '\x00'})), align, int32(127))
			} else if int64(uintptr(unsafe.Pointer(q))-uintptr(unsafe.Pointer(p))) != int64(127) {
				common.T_printf((*int8)(unsafe.Pointer(&[126]int8{'/', 'U', 's', 'e', 'r', 's', '/', 'x', 'u', 's', 'h', 'i', 'w', 'e', 'i', '/', 'w', 'o', 'r', 'k', '/', 'l', 'i', 'b', 'c', '/', 't', 'e', 's', 't', '/', 's', 'r', 'c', '/', 'f', 'u', 'n', 'c', 't', 'i', 'o', 'n', 'a', 'l', '/', 's', 't', 'r', 'i', 'n', 'g', '_', 's', 't', 'r', 'c', 'h', 'r', '.', 'c', ':', '7', '3', ':', ' ', 's', 't', 'r', 'c', 'h', 'r', '(', '%', 's', ',', '%', 's', ')', ' ', 'w', 'i', 't', 'h', ' ', 'a', 'l', 'i', 'g', 'n', '=', '%', 'd', ' ', 'r', 'e', 't', 'u', 'r', 'n', 'e', 'd', ' ', 's', 't', 'r', '+', '%', 'd', ',', ' ', 'w', 'a', 'n', 't', 'e', 'd', ' ', 's', 't', 'r', '+', '%', 'd', '\n', '\x00'})), (*int8)(unsafe.Pointer(&[2]int8{'s', '\x00'})), (*int8)(unsafe.Pointer(&[4]int8{'1', '2', '8', '\x00'})), align, uintptr(unsafe.Pointer(q))-uintptr(unsafe.Pointer(p)), int32(127))
			}
		}
	}
	{
		var align int32
		for align = int32(0); align < int32(8); align++ {
			var p *int8 = (*int8)(_cgos_aligncpy_string_strchr(unsafe.Pointer((*int8)(unsafe.Pointer(&s))), 256, uint64(align)))
			var q *int8 = libc.Strchr(p, int32(255))
			if uintptr(unsafe.Pointer(q)) == uintptr(unsafe.Pointer(nil)) {
				common.T_printf((*int8)(unsafe.Pointer(&[121]int8{'/', 'U', 's', 'e', 'r', 's', '/', 'x', 'u', 's', 'h', 'i', 'w', 'e', 'i', '/', 'w', 'o', 'r', 'k', '/', 'l', 'i', 'b', 'c', '/', 't', 'e', 's', 't', '/', 's', 'r', 'c', '/', 'f', 'u', 'n', 'c', 't', 'i', 'o', 'n', 'a', 'l', '/', 's', 't', 'r', 'i', 'n', 'g', '_', 's', 't', 'r', 'c', 'h', 'r', '.', 'c', ':', '7', '4', ':', ' ', 's', 't', 'r', 'c', 'h', 'r', '(', '%', 's', ',', '%', 's', ')', ' ', 'w', 'i', 't', 'h', ' ', 'a', 'l', 'i', 'g', 'n', '=', '%', 'd', ' ', 'r', 'e', 't', 'u', 'r', 'n', 'e', 'd', ' ', '0', ',', ' ', 'w', 'a', 'n', 't', 'e', 'd', ' ', 's', 't', 'r', '+', '%', 'd', '\n', '\x00'})), (*int8)(unsafe.Pointer(&[2]int8{'s', '\x00'})), (*int8)(unsafe.Pointer(&[4]int8{'2', '5', '5', '\x00'})), align, int32(254))
			} else if int64(uintptr(unsafe.Pointer(q))-uintptr(unsafe.Pointer(p))) != int64(254) {
				common.T_printf((*int8)(unsafe.Pointer(&[126]int8{'/', 'U', 's', 'e', 'r', 's', '/', 'x', 'u', 's', 'h', 'i', 'w', 'e', 'i', '/', 'w', 'o', 'r', 'k', '/', 'l', 'i', 'b', 'c', '/', 't', 'e', 's', 't', '/', 's', 'r', 'c', '/', 'f', 'u', 'n', 'c', 't', 'i', 'o', 'n', 'a', 'l', '/', 's', 't', 'r', 'i', 'n', 'g', '_', 's', 't', 'r', 'c', 'h', 'r', '.', 'c', ':', '7', '4', ':', ' ', 's', 't', 'r', 'c', 'h', 'r', '(', '%', 's', ',', '%', 's', ')', ' ', 'w', 'i', 't', 'h', ' ', 'a', 'l', 'i', 'g', 'n', '=', '%', 'd', ' ', 'r', 'e', 't', 'u', 'r', 'n', 'e', 'd', ' ', 's', 't', 'r', '+', '%', 'd', ',', ' ', 'w', 'a', 'n', 't', 'e', 'd', ' ', 's', 't', 'r', '+', '%', 'd', '\n', '\x00'})), (*int8)(unsafe.Pointer(&[2]int8{'s', '\x00'})), (*int8)(unsafe.Pointer(&[4]int8{'2', '5', '5', '\x00'})), align, uintptr(unsafe.Pointer(q))-uintptr(unsafe.Pointer(p)), int32(254))
			}
		}
	}
	{
		var align int32
		for align = int32(0); align < int32(8); align++ {
			var p *int8 = (*int8)(_cgos_aligncpy_string_strchr(unsafe.Pointer((*int8)(unsafe.Pointer(&s))), 256, uint64(align)))
			var q *int8 = libc.Strchr(p, int32(0))
			if uintptr(unsafe.Pointer(q)) == uintptr(unsafe.Pointer(nil)) {
				common.T_printf((*int8)(unsafe.Pointer(&[121]int8{'/', 'U', 's', 'e', 'r', 's', '/', 'x', 'u', 's', 'h', 'i', 'w', 'e', 'i', '/', 'w', 'o', 'r', 'k', '/', 'l', 'i', 'b', 'c', '/', 't', 'e', 's', 't', '/', 's', 'r', 'c', '/', 'f', 'u', 'n', 'c', 't', 'i', 'o', 'n', 'a', 'l', '/', 's', 't', 'r', 'i', 'n', 'g', '_', 's', 't', 'r', 'c', 'h', 'r', '.', 'c', ':', '7', '5', ':', ' ', 's', 't', 'r', 'c', 'h', 'r', '(', '%', 's', ',', '%', 's', ')', ' ', 'w', 'i', 't', 'h', ' ', 'a', 'l', 'i', 'g', 'n', '=', '%', 'd', ' ', 'r', 'e', 't', 'u', 'r', 'n', 'e', 'd', ' ', '0', ',', ' ', 'w', 'a', 'n', 't', 'e', 'd', ' ', 's', 't', 'r', '+', '%', 'd', '\n', '\x00'})), (*int8)(unsafe.Pointer(&[2]int8{'s', '\x00'})), (*int8)(unsafe.Pointer(&[2]int8{'0', '\x00'})), align, int32(255))
			} else if int64(uintptr(unsafe.Pointer(q))-uintptr(unsafe.Pointer(p))) != int64(255) {
				common.T_printf((*int8)(unsafe.Pointer(&[126]int8{'/', 'U', 's', 'e', 'r', 's', '/', 'x', 'u', 's', 'h', 'i', 'w', 'e', 'i', '/', 'w', 'o', 'r', 'k', '/', 'l', 'i', 'b', 'c', '/', 't', 'e', 's', 't', '/', 's', 'r', 'c', '/', 'f', 'u', 'n', 'c', 't', 'i', 'o', 'n', 'a', 'l', '/', 's', 't', 'r', 'i', 'n', 'g', '_', 's', 't', 'r', 'c', 'h', 'r', '.', 'c', ':', '7', '5', ':', ' ', 's', 't', 'r', 'c', 'h', 'r', '(', '%', 's', ',', '%', 's', ')', ' ', 'w', 'i', 't', 'h', ' ', 'a', 'l', 'i', 'g', 'n', '=', '%', 'd', ' ', 'r', 'e', 't', 'u', 'r', 'n', 'e', 'd', ' ', 's', 't', 'r', '+', '%', 'd', ',', ' ', 'w', 'a', 'n', 't', 'e', 'd', ' ', 's', 't', 'r', '+', '%', 'd', '\n', '\x00'})), (*int8)(unsafe.Pointer(&[2]int8{'s', '\x00'})), (*int8)(unsafe.Pointer(&[2]int8{'0', '\x00'})), align, uintptr(unsafe.Pointer(q))-uintptr(unsafe.Pointer(p)), int32(255))
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
