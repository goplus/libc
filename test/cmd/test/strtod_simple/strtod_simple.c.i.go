package strtod_simple

import (
	libc "github.com/goplus/libc"
	unsafe "unsafe"
	common "github.com/goplus/libc/test/common"
	testing "testing"
)

func _cgo_main() int32 {
	var i int32
	var d float64
	var d2 float64
	var buf [1000]int8
	for i = int32(0); i < int32(100); i++ {
		d = libc.Sin(float64(i))
		libc.Snprintf((*int8)(unsafe.Pointer(&buf)), 1000, (*int8)(unsafe.Pointer(&[7]int8{'%', '.', '3', '0', '0', 'f', '\x00'})), d)
		if !(func() (_cgo_ret float64) {
			_cgo_addr := &d2
			*_cgo_addr = libc.Strtod((*int8)(unsafe.Pointer(&buf)), nil)
			return *_cgo_addr
		}() == d) {
			func() int32 {
				common.T_printf((*int8)(unsafe.Pointer(&[109]int8{'/', 'U', 's', 'e', 'r', 's', '/', 'x', 'u', 's', 'h', 'i', 'w', 'e', 'i', '/', 'w', 'o', 'r', 'k', '/', 'l', 'i', 'b', 'c', '/', 't', 'e', 's', 't', '/', 's', 'r', 'c', '/', 'f', 'u', 'n', 'c', 't', 'i', 'o', 'n', 'a', 'l', '/', 's', 't', 'r', 't', 'o', 'd', '_', 's', 'i', 'm', 'p', 'l', 'e', '.', 'c', ':', '2', '5', ':', ' ', '%', 's', ' ', 'f', 'a', 'i', 'l', 'e', 'd', ' ', '(', 'r', 'o', 'u', 'n', 'd', ' ', 't', 'r', 'i', 'p', ' ', 'f', 'a', 'i', 'l', ' ', '%', 'a', ' ', '!', '=', ' ', '%', 'a', ' ', '(', '%', 'a', ')', ')', '\n', '\x00'})), (*int8)(unsafe.Pointer(&[15]int8{'s', 't', 'r', 't', 'o', 'd', '(', 'b', 'u', 'f', ',', ' ', '0', ')', '\x00'})), d2, d, d2-d)
				return int32(0)
			}()
		}
	}
	if !(func() (_cgo_ret float64) {
		_cgo_addr := &d
		*_cgo_addr = libc.Strtod((*int8)(unsafe.Pointer(&[6]int8{'0', 'x', '1', 'p', '4', '\x00'})), nil)
		return *_cgo_addr
	}() == 16) {
		func() int32 {
			common.T_printf((*int8)(unsafe.Pointer(&[98]int8{'/', 'U', 's', 'e', 'r', 's', '/', 'x', 'u', 's', 'h', 'i', 'w', 'e', 'i', '/', 'w', 'o', 'r', 'k', '/', 'l', 'i', 'b', 'c', '/', 't', 'e', 's', 't', '/', 's', 'r', 'c', '/', 'f', 'u', 'n', 'c', 't', 'i', 'o', 'n', 'a', 'l', '/', 's', 't', 'r', 't', 'o', 'd', '_', 's', 'i', 'm', 'p', 'l', 'e', '.', 'c', ':', '2', '8', ':', ' ', '%', 's', ' ', 'f', 'a', 'i', 'l', 'e', 'd', ' ', '(', 'h', 'e', 'x', ' ', 'f', 'l', 'o', 'a', 't', ' ', '%', 'a', ' ', '!', '=', ' ', '%', 'a', ')', '\n', '\x00'})), (*int8)(unsafe.Pointer(&[19]int8{'s', 't', 'r', 't', 'o', 'd', '(', '"', '0', 'x', '1', 'p', '4', '"', ',', ' ', '0', ')', '\x00'})), d, 16, d-16)
			return int32(0)
		}()
	}
	if !(func() (_cgo_ret float64) {
		_cgo_addr := &d
		*_cgo_addr = libc.Strtod((*int8)(unsafe.Pointer(&[8]int8{'0', 'x', '1', '.', '1', 'p', '4', '\x00'})), nil)
		return *_cgo_addr
	}() == 17) {
		func() int32 {
			common.T_printf((*int8)(unsafe.Pointer(&[98]int8{'/', 'U', 's', 'e', 'r', 's', '/', 'x', 'u', 's', 'h', 'i', 'w', 'e', 'i', '/', 'w', 'o', 'r', 'k', '/', 'l', 'i', 'b', 'c', '/', 't', 'e', 's', 't', '/', 's', 'r', 'c', '/', 'f', 'u', 'n', 'c', 't', 'i', 'o', 'n', 'a', 'l', '/', 's', 't', 'r', 't', 'o', 'd', '_', 's', 'i', 'm', 'p', 'l', 'e', '.', 'c', ':', '2', '9', ':', ' ', '%', 's', ' ', 'f', 'a', 'i', 'l', 'e', 'd', ' ', '(', 'h', 'e', 'x', ' ', 'f', 'l', 'o', 'a', 't', ' ', '%', 'a', ' ', '!', '=', ' ', '%', 'a', ')', '\n', '\x00'})), (*int8)(unsafe.Pointer(&[21]int8{'s', 't', 'r', 't', 'o', 'd', '(', '"', '0', 'x', '1', '.', '1', 'p', '4', '"', ',', ' ', '0', ')', '\x00'})), d, 17, d-17)
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
