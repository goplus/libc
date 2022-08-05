package strtod_long

import (
	libc "github.com/goplus/libc"
	unsafe "unsafe"
	common "github.com/goplus/libc/test/common"
	testing "testing"
)

func _cgo_main() int32 {
	var x float64
	var want float64 = 0.1111111111111111
	var buf [40000]int8
	libc.Memset(unsafe.Pointer((*int8)(unsafe.Pointer(&buf))), '1', 40000)
	*(*int8)(unsafe.Pointer(uintptr(unsafe.Pointer((*int8)(unsafe.Pointer(&buf)))) + uintptr(int32(0)))) = int8('.')
	*(*int8)(unsafe.Pointer(uintptr(unsafe.Pointer((*int8)(unsafe.Pointer(&buf)))) + uintptr(39999))) = int8(0)
	if func() (_cgo_ret float64) {
		_cgo_addr := &x
		*_cgo_addr = libc.Strtod((*int8)(unsafe.Pointer(&buf)), nil)
		return *_cgo_addr
	}() != want {
		common.T_printf((*int8)(unsafe.Pointer(&[98]int8{'/', 'U', 's', 'e', 'r', 's', '/', 'x', 'u', 's', 'h', 'i', 'w', 'e', 'i', '/', 'w', 'o', 'r', 'k', '/', 'l', 'i', 'b', 'c', '/', 't', 'e', 's', 't', '/', 's', 'r', 'c', '/', 'f', 'u', 'n', 'c', 't', 'i', 'o', 'n', 'a', 'l', '/', 's', 't', 'r', 't', 'o', 'd', '_', 'l', 'o', 'n', 'g', '.', 'c', ':', '1', '6', ':', ' ', 's', 't', 'r', 't', 'o', 'd', '(', '.', '1', '1', '[', '.', '.', '.', ']', '1', ')', ' ', 'g', 'o', 't', ' ', '%', 'a', ' ', 'w', 'a', 'n', 't', ' ', '%', 'a', '\n', '\x00'})), x, want)
	}
	return common.T_status
}
func TestMain(t *testing.T) {
	if _cgo_ret := _cgo_main(); _cgo_ret != 0 {
		t.Fatal("exit status", _cgo_ret)
	}
}
