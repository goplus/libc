package random

import (
	unsafe "unsafe"
	libc "github.com/goplus/libc"
	common "github.com/goplus/libc/test/common"
	testing "testing"
)

func chkmissing_cgo1_random(x *int64) int32 {
	var d [8]int32 = [8]int32{int32(0)}
	var i int32
	for i = int32(0); i < int32(100); i++ {
		*(*int32)(unsafe.Pointer(uintptr(unsafe.Pointer((*int32)(unsafe.Pointer(&d)))) + uintptr(*(*int64)(unsafe.Pointer(uintptr(unsafe.Pointer(x)) + uintptr(i)*8))%int64(8))*4))++
	}
	for i = int32(0); i < int32(8); i++ {
		if *(*int32)(unsafe.Pointer(uintptr(unsafe.Pointer((*int32)(unsafe.Pointer(&d)))) + uintptr(i)*4)) == int32(0) {
			return int32(1)
		}
	}
	return int32(0)
}
func chkrepeat_cgo2_random(x *int64) int32 {
	var i int32
	var j int32
	for i = int32(0); i < int32(100); i++ {
		for j = int32(0); j < i; j++ {
			if *(*int64)(unsafe.Pointer(uintptr(unsafe.Pointer(x)) + uintptr(i)*8)) == *(*int64)(unsafe.Pointer(uintptr(unsafe.Pointer(x)) + uintptr(j)*8)) {
				return int32(1)
			}
		}
	}
	return int32(0)
}

var orx_cgo3_random uint32

func chkones_cgo4_random(x *int64) int32 {
	var i int32
	orx_cgo3_random = uint32(0)
	for i = int32(0); i < int32(20); i++ {
		orx_cgo3_random |= uint32(*(*int64)(unsafe.Pointer(uintptr(unsafe.Pointer(x)) + uintptr(i)*8)))
	}
	return func() int32 {
		if orx_cgo3_random != uint32(2147483647) {
			return 1
		} else {
			return 0
		}
	}()
}
func checkseed(seed uint32, x *int64) {
	var i int32
	libc.Srandom(seed)
	for i = int32(0); i < int32(100); i++ {
		*(*int64)(unsafe.Pointer(uintptr(unsafe.Pointer(x)) + uintptr(i)*8)) = libc.Random()
	}
	if chkmissing_cgo1_random(x) != 0 {
		common.T_printf((*int8)(unsafe.Pointer(&[71]int8{'s', 'r', 'c', '/', 'f', 'u', 'n', 'c', 't', 'i', 'o', 'n', 'a', 'l', '/', 'r', 'a', 'n', 'd', 'o', 'm', '.', 'c', ':', '5', '0', ':', ' ', 'w', 'e', 'a', 'k', ' ', 's', 'e', 'e', 'd', ' ', '%', 'd', ',', ' ', 'm', 'i', 's', 's', 'i', 'n', 'g', ' ', 'p', 'a', 't', 't', 'e', 'r', 'n', ' ', 'i', 'n', ' ', 'l', 'o', 'w', ' ', 'b', 'i', 't', 's', '\n', '\x00'})), seed)
	}
	if chkrepeat_cgo2_random(x) != 0 {
		common.T_printf((*int8)(unsafe.Pointer(&[57]int8{'s', 'r', 'c', '/', 'f', 'u', 'n', 'c', 't', 'i', 'o', 'n', 'a', 'l', '/', 'r', 'a', 'n', 'd', 'o', 'm', '.', 'c', ':', '5', '2', ':', ' ', 'w', 'e', 'a', 'k', ' ', 's', 'e', 'e', 'd', ' ', '%', 'd', ',', ' ', 'e', 'x', 'a', 'c', 't', ' ', 'r', 'e', 'p', 'e', 'a', 't', 's', '\n', '\x00'})), seed)
	}
	if chkones_cgo4_random(x) != 0 {
		common.T_printf((*int8)(unsafe.Pointer(&[62]int8{'s', 'r', 'c', '/', 'f', 'u', 'n', 'c', 't', 'i', 'o', 'n', 'a', 'l', '/', 'r', 'a', 'n', 'd', 'o', 'm', '.', 'c', ':', '5', '4', ':', ' ', 'w', 'e', 'a', 'k', ' ', 's', 'e', 'e', 'd', ' ', '%', 'd', ',', ' ', 'o', 'r', ' ', 'p', 'a', 't', 't', 'e', 'r', 'n', ':', ' ', '0', 'x', '%', '0', '8', 'x', '\n', '\x00'})), seed, orx_cgo3_random)
	}
}
func _cgo_main() int32 {
	var x [100]int64
	var y int64
	var z int64
	var i int32
	var state [128]int8
	var p *int8
	var q *int8
	for i = int32(0); i < int32(100); i++ {
		*(*int64)(unsafe.Pointer(uintptr(unsafe.Pointer((*int64)(unsafe.Pointer(&x)))) + uintptr(i)*8)) = libc.Random()
	}
	p = libc.Initstate(uint32(1), (*int8)(unsafe.Pointer(&state)), 128)
	for i = int32(0); i < int32(100); i++ {
		if *(*int64)(unsafe.Pointer(uintptr(unsafe.Pointer((*int64)(unsafe.Pointer(&x)))) + uintptr(i)*8)) != func() (_cgo_ret int64) {
			_cgo_addr := &y
			*_cgo_addr = libc.Random()
			return *_cgo_addr
		}() {
			common.T_printf((*int8)(unsafe.Pointer(&[88]int8{'s', 'r', 'c', '/', 'f', 'u', 'n', 'c', 't', 'i', 'o', 'n', 'a', 'l', '/', 'r', 'a', 'n', 'd', 'o', 'm', '.', 'c', ':', '7', '1', ':', ' ', 'i', 'n', 'i', 't', 's', 't', 'a', 't', 'e', '(', '1', ')', ' ', 'i', 's', ' ', 'n', 'o', 't', ' ', 'd', 'e', 'f', 'a', 'u', 'l', 't', ':', ' ', '(', '%', 'd', ')', ' ', 'd', 'e', 'f', 'a', 'u', 'l', 't', ':', ' ', '%', 'l', 'd', ',', ' ', 's', 'e', 'e', 'd', '1', ':', ' ', '%', 'l', 'd', '\n', '\x00'})), i, *(*int64)(unsafe.Pointer(uintptr(unsafe.Pointer((*int64)(unsafe.Pointer(&x)))) + uintptr(i)*8)), y)
		}
	}
	for i = int32(0); i < int32(10); i++ {
		z = libc.Random()
		q = libc.Setstate(p)
		if z != func() (_cgo_ret int64) {
			_cgo_addr := &y
			*_cgo_addr = libc.Random()
			return *_cgo_addr
		}() {
			common.T_printf((*int8)(unsafe.Pointer(&[72]int8{'s', 'r', 'c', '/', 'f', 'u', 'n', 'c', 't', 'i', 'o', 'n', 'a', 'l', '/', 'r', 'a', 'n', 'd', 'o', 'm', '.', 'c', ':', '7', '6', ':', ' ', 's', 'e', 't', 's', 't', 'a', 't', 'e', ' ', 'f', 'a', 'i', 'l', 'e', 'd', ' ', '(', '%', 'd', ')', ' ', 'o', 'r', 'i', 'g', ':', ' ', '%', 'l', 'd', ',', ' ', 'r', 'e', 's', 'e', 't', ':', ' ', '%', 'l', 'd', '\n', '\x00'})), i, z, y)
		}
		p = libc.Setstate(q)
	}
	libc.Srandom(uint32(1))
	for i = int32(0); i < int32(100); i++ {
		if *(*int64)(unsafe.Pointer(uintptr(unsafe.Pointer((*int64)(unsafe.Pointer(&x)))) + uintptr(i)*8)) != func() (_cgo_ret int64) {
			_cgo_addr := &y
			*_cgo_addr = libc.Random()
			return *_cgo_addr
		}() {
			common.T_printf((*int8)(unsafe.Pointer(&[86]int8{'s', 'r', 'c', '/', 'f', 'u', 'n', 'c', 't', 'i', 'o', 'n', 'a', 'l', '/', 'r', 'a', 'n', 'd', 'o', 'm', '.', 'c', ':', '8', '2', ':', ' ', 's', 'r', 'a', 'n', 'd', 'o', 'm', '(', '1', ')', ' ', 'i', 's', ' ', 'n', 'o', 't', ' ', 'd', 'e', 'f', 'a', 'u', 'l', 't', ':', ' ', '(', '%', 'd', ')', ' ', 'd', 'e', 'f', 'a', 'u', 'l', 't', ':', ' ', '%', 'l', 'd', ',', ' ', 's', 'e', 'e', 'd', '1', ':', ' ', '%', 'l', 'd', '\n', '\x00'})), i, *(*int64)(unsafe.Pointer(uintptr(unsafe.Pointer((*int64)(unsafe.Pointer(&x)))) + uintptr(i)*8)), y)
		}
	}
	checkseed(uint32(2147483647), (*int64)(unsafe.Pointer(&x)))
	for i = int32(0); i < int32(10); i++ {
		checkseed(uint32(i), (*int64)(unsafe.Pointer(&x)))
	}
	return common.T_status
}
func TestMain(t *testing.T) {
	if _cgo_ret := _cgo_main(); _cgo_ret != 0 {
		t.Fatal("exit status", _cgo_ret)
	}
}
