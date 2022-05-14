package main

import (
	unsafe "unsafe"
	libc "github.com/goplus/libc"
	os "os"
)

type _cgoa_1 struct {
	fmt    *int8
	i      int32
	expect *int8
}

var int_tests [40]_cgoa_1 = [40]_cgoa_1{_cgoa_1{(*int8)(unsafe.Pointer(&[5]int8{'%', '0', '4', 'd', '\x00'})), int32(12), (*int8)(unsafe.Pointer(&[5]int8{'0', '0', '1', '2', '\x00'}))}, _cgoa_1{(*int8)(unsafe.Pointer(&[5]int8{'%', '.', '3', 'd', '\x00'})), int32(12), (*int8)(unsafe.Pointer(&[4]int8{'0', '1', '2', '\x00'}))}, _cgoa_1{(*int8)(unsafe.Pointer(&[4]int8{'%', '3', 'd', '\x00'})), int32(12), (*int8)(unsafe.Pointer(&[4]int8{' ', '1', '2', '\x00'}))}, _cgoa_1{(*int8)(unsafe.Pointer(&[5]int8{'%', '-', '3', 'd', '\x00'})), int32(12), (*int8)(unsafe.Pointer(&[4]int8{'1', '2', ' ', '\x00'}))}, _cgoa_1{(*int8)(unsafe.Pointer(&[5]int8{'%', '+', '3', 'd', '\x00'})), int32(12), (*int8)(unsafe.Pointer(&[4]int8{'+', '1', '2', '\x00'}))}, _cgoa_1{(*int8)(unsafe.Pointer(&[6]int8{'%', '+', '-', '5', 'd', '\x00'})), int32(12), (*int8)(unsafe.Pointer(&[6]int8{'+', '1', '2', ' ', ' ', '\x00'}))}, _cgoa_1{(*int8)(unsafe.Pointer(&[7]int8{'%', '+', '-', ' ', '5', 'd', '\x00'})), int32(12), (*int8)(unsafe.Pointer(&[6]int8{'+', '1', '2', ' ', ' ', '\x00'}))}, _cgoa_1{(*int8)(unsafe.Pointer(&[6]int8{'%', '-', ' ', '5', 'd', '\x00'})), int32(12), (*int8)(unsafe.Pointer(&[6]int8{' ', '1', '2', ' ', ' ', '\x00'}))}, _cgoa_1{(*int8)(unsafe.Pointer(&[4]int8{'%', ' ', 'd', '\x00'})), int32(12), (*int8)(unsafe.Pointer(&[4]int8{' ', '1', '2', '\x00'}))}, _cgoa_1{(*int8)(unsafe.Pointer(&[6]int8{'%', '0', '-', '5', 'd', '\x00'})), int32(12), (*int8)(unsafe.Pointer(&[6]int8{'1', '2', ' ', ' ', ' ', '\x00'}))}, _cgoa_1{(*int8)(unsafe.Pointer(&[6]int8{'%', '-', '0', '5', 'd', '\x00'})), int32(12), (*int8)(unsafe.Pointer(&[6]int8{'1', '2', ' ', ' ', ' ', '\x00'}))}, _cgoa_1{(*int8)(unsafe.Pointer(&[5]int8{'%', '.', '0', 'd', '\x00'})), int32(0), (*int8)(unsafe.Pointer(&[1]int8{'\x00'}))}, _cgoa_1{(*int8)(unsafe.Pointer(&[5]int8{'%', '.', '0', 'o', '\x00'})), int32(0), (*int8)(unsafe.Pointer(&[1]int8{'\x00'}))}, _cgoa_1{(*int8)(unsafe.Pointer(&[6]int8{'%', '#', '.', '0', 'd', '\x00'})), int32(0), (*int8)(unsafe.Pointer(&[1]int8{'\x00'}))}, _cgoa_1{(*int8)(unsafe.Pointer(&[6]int8{'%', '#', '.', '0', 'o', '\x00'})), int32(0), (*int8)(unsafe.Pointer(&[2]int8{'0', '\x00'}))}, _cgoa_1{(*int8)(unsafe.Pointer(&[6]int8{'%', '#', '.', '0', 'x', '\x00'})), int32(0), (*int8)(unsafe.Pointer(&[1]int8{'\x00'}))}, _cgoa_1{(*int8)(unsafe.Pointer(&[6]int8{'%', '2', '.', '0', 'u', '\x00'})), int32(0), (*int8)(unsafe.Pointer(&[3]int8{' ', ' ', '\x00'}))}, _cgoa_1{(*int8)(unsafe.Pointer(&[7]int8{'%', '0', '2', '.', '0', 'u', '\x00'})), int32(0), (*int8)(unsafe.Pointer(&[3]int8{' ', ' ', '\x00'}))}, _cgoa_1{(*int8)(unsafe.Pointer(&[6]int8{'%', '2', '.', '0', 'd', '\x00'})), int32(0), (*int8)(unsafe.Pointer(&[3]int8{' ', ' ', '\x00'}))}, _cgoa_1{(*int8)(unsafe.Pointer(&[7]int8{'%', '0', '2', '.', '0', 'd', '\x00'})), int32(0), (*int8)(unsafe.Pointer(&[3]int8{' ', ' ', '\x00'}))}, _cgoa_1{(*int8)(unsafe.Pointer(&[6]int8{'%', ' ', '.', '0', 'd', '\x00'})), int32(0), (*int8)(unsafe.Pointer(&[2]int8{' ', '\x00'}))}, _cgoa_1{(*int8)(unsafe.Pointer(&[6]int8{'%', '+', '.', '0', 'd', '\x00'})), int32(0), (*int8)(unsafe.Pointer(&[2]int8{'+', '\x00'}))}, _cgoa_1{(*int8)(unsafe.Pointer(&[3]int8{'%', 'x', '\x00'})), int32(63), (*int8)(unsafe.Pointer(&[3]int8{'3', 'f', '\x00'}))}, _cgoa_1{(*int8)(unsafe.Pointer(&[4]int8{'%', '#', 'x', '\x00'})), int32(63), (*int8)(unsafe.Pointer(&[5]int8{'0', 'x', '3', 'f', '\x00'}))}, _cgoa_1{(*int8)(unsafe.Pointer(&[3]int8{'%', 'X', '\x00'})), int32(63), (*int8)(unsafe.Pointer(&[3]int8{'3', 'F', '\x00'}))}, _cgoa_1{(*int8)(unsafe.Pointer(&[3]int8{'%', 'o', '\x00'})), int32(15), (*int8)(unsafe.Pointer(&[3]int8{'1', '7', '\x00'}))}, _cgoa_1{(*int8)(unsafe.Pointer(&[4]int8{'%', '#', 'o', '\x00'})), int32(15), (*int8)(unsafe.Pointer(&[4]int8{'0', '1', '7', '\x00'}))}, _cgoa_1{(*int8)(unsafe.Pointer(&[4]int8{'%', '#', 'o', '\x00'})), int32(0), (*int8)(unsafe.Pointer(&[2]int8{'0', '\x00'}))}, _cgoa_1{(*int8)(unsafe.Pointer(&[6]int8{'%', '#', '.', '0', 'o', '\x00'})), int32(0), (*int8)(unsafe.Pointer(&[2]int8{'0', '\x00'}))}, _cgoa_1{(*int8)(unsafe.Pointer(&[6]int8{'%', '#', '.', '1', 'o', '\x00'})), int32(0), (*int8)(unsafe.Pointer(&[2]int8{'0', '\x00'}))}, _cgoa_1{(*int8)(unsafe.Pointer(&[4]int8{'%', '#', 'o', '\x00'})), int32(1), (*int8)(unsafe.Pointer(&[3]int8{'0', '1', '\x00'}))}, _cgoa_1{(*int8)(unsafe.Pointer(&[6]int8{'%', '#', '.', '0', 'o', '\x00'})), int32(1), (*int8)(unsafe.Pointer(&[3]int8{'0', '1', '\x00'}))}, _cgoa_1{(*int8)(unsafe.Pointer(&[6]int8{'%', '#', '.', '1', 'o', '\x00'})), int32(1), (*int8)(unsafe.Pointer(&[3]int8{'0', '1', '\x00'}))}, _cgoa_1{(*int8)(unsafe.Pointer(&[6]int8{'%', '#', '0', '4', 'o', '\x00'})), int32(1), (*int8)(unsafe.Pointer(&[5]int8{'0', '0', '0', '1', '\x00'}))}, _cgoa_1{(*int8)(unsafe.Pointer(&[8]int8{'%', '#', '0', '4', '.', '0', 'o', '\x00'})), int32(1), (*int8)(unsafe.Pointer(&[5]int8{' ', ' ', '0', '1', '\x00'}))}, _cgoa_1{(*int8)(unsafe.Pointer(&[8]int8{'%', '#', '0', '4', '.', '1', 'o', '\x00'})), int32(1), (*int8)(unsafe.Pointer(&[5]int8{' ', ' ', '0', '1', '\x00'}))}, _cgoa_1{(*int8)(unsafe.Pointer(&[5]int8{'%', '0', '4', 'o', '\x00'})), int32(1), (*int8)(unsafe.Pointer(&[5]int8{'0', '0', '0', '1', '\x00'}))}, _cgoa_1{(*int8)(unsafe.Pointer(&[7]int8{'%', '0', '4', '.', '0', 'o', '\x00'})), int32(1), (*int8)(unsafe.Pointer(&[5]int8{' ', ' ', ' ', '1', '\x00'}))}, _cgoa_1{(*int8)(unsafe.Pointer(&[7]int8{'%', '0', '4', '.', '1', 'o', '\x00'})), int32(1), (*int8)(unsafe.Pointer(&[5]int8{' ', ' ', ' ', '1', '\x00'}))}, _cgoa_1{nil, int32(0), nil}}

type _cgoa_2 struct {
	fmt    *int8
	f      float64
	expect *int8
}

var fp_tests [43]_cgoa_2 = [43]_cgoa_2{_cgoa_2{(*int8)(unsafe.Pointer(&[3]int8{'%', 'a', '\x00'})), 0, (*int8)(unsafe.Pointer(&[7]int8{'0', 'x', '0', 'p', '+', '0', '\x00'}))}, _cgoa_2{(*int8)(unsafe.Pointer(&[3]int8{'%', 'e', '\x00'})), 0, (*int8)(unsafe.Pointer(&[13]int8{'0', '.', '0', '0', '0', '0', '0', '0', 'e', '+', '0', '0', '\x00'}))}, _cgoa_2{(*int8)(unsafe.Pointer(&[3]int8{'%', 'f', '\x00'})), 0, (*int8)(unsafe.Pointer(&[9]int8{'0', '.', '0', '0', '0', '0', '0', '0', '\x00'}))}, _cgoa_2{(*int8)(unsafe.Pointer(&[3]int8{'%', 'g', '\x00'})), 0, (*int8)(unsafe.Pointer(&[2]int8{'0', '\x00'}))}, _cgoa_2{(*int8)(unsafe.Pointer(&[4]int8{'%', '#', 'g', '\x00'})), 0, (*int8)(unsafe.Pointer(&[8]int8{'0', '.', '0', '0', '0', '0', '0', '\x00'}))}, _cgoa_2{(*int8)(unsafe.Pointer(&[4]int8{'%', 'l', 'a', '\x00'})), 0, (*int8)(unsafe.Pointer(&[7]int8{'0', 'x', '0', 'p', '+', '0', '\x00'}))}, _cgoa_2{(*int8)(unsafe.Pointer(&[4]int8{'%', 'l', 'e', '\x00'})), 0, (*int8)(unsafe.Pointer(&[13]int8{'0', '.', '0', '0', '0', '0', '0', '0', 'e', '+', '0', '0', '\x00'}))}, _cgoa_2{(*int8)(unsafe.Pointer(&[4]int8{'%', 'l', 'f', '\x00'})), 0, (*int8)(unsafe.Pointer(&[9]int8{'0', '.', '0', '0', '0', '0', '0', '0', '\x00'}))}, _cgoa_2{(*int8)(unsafe.Pointer(&[4]int8{'%', 'l', 'g', '\x00'})), 0, (*int8)(unsafe.Pointer(&[2]int8{'0', '\x00'}))}, _cgoa_2{(*int8)(unsafe.Pointer(&[5]int8{'%', '#', 'l', 'g', '\x00'})), 0, (*int8)(unsafe.Pointer(&[8]int8{'0', '.', '0', '0', '0', '0', '0', '\x00'}))}, _cgoa_2{(*int8)(unsafe.Pointer(&[3]int8{'%', 'f', '\x00'})), 1.1000000000000001, (*int8)(unsafe.Pointer(&[9]int8{'1', '.', '1', '0', '0', '0', '0', '0', '\x00'}))}, _cgoa_2{(*int8)(unsafe.Pointer(&[3]int8{'%', 'f', '\x00'})), 1.2, (*int8)(unsafe.Pointer(&[9]int8{'1', '.', '2', '0', '0', '0', '0', '0', '\x00'}))}, _cgoa_2{(*int8)(unsafe.Pointer(&[3]int8{'%', 'f', '\x00'})), 1.3, (*int8)(unsafe.Pointer(&[9]int8{'1', '.', '3', '0', '0', '0', '0', '0', '\x00'}))}, _cgoa_2{(*int8)(unsafe.Pointer(&[3]int8{'%', 'f', '\x00'})), 1.3999999999999999, (*int8)(unsafe.Pointer(&[9]int8{'1', '.', '4', '0', '0', '0', '0', '0', '\x00'}))}, _cgoa_2{(*int8)(unsafe.Pointer(&[3]int8{'%', 'f', '\x00'})), 1.5, (*int8)(unsafe.Pointer(&[9]int8{'1', '.', '5', '0', '0', '0', '0', '0', '\x00'}))}, _cgoa_2{(*int8)(unsafe.Pointer(&[5]int8{'%', '.', '4', 'f', '\x00'})), 1.06125, (*int8)(unsafe.Pointer(&[7]int8{'1', '.', '0', '6', '1', '3', '\x00'}))}, _cgoa_2{(*int8)(unsafe.Pointer(&[5]int8{'%', '.', '4', 'f', '\x00'})), 1.03125, (*int8)(unsafe.Pointer(&[7]int8{'1', '.', '0', '3', '1', '2', '\x00'}))}, _cgoa_2{(*int8)(unsafe.Pointer(&[5]int8{'%', '.', '2', 'f', '\x00'})), 1.375, (*int8)(unsafe.Pointer(&[5]int8{'1', '.', '3', '8', '\x00'}))}, _cgoa_2{(*int8)(unsafe.Pointer(&[5]int8{'%', '.', '1', 'f', '\x00'})), 1.375, (*int8)(unsafe.Pointer(&[4]int8{'1', '.', '4', '\x00'}))}, _cgoa_2{(*int8)(unsafe.Pointer(&[6]int8{'%', '.', '1', 'l', 'f', '\x00'})), 1.375, (*int8)(unsafe.Pointer(&[4]int8{'1', '.', '4', '\x00'}))}, _cgoa_2{(*int8)(unsafe.Pointer(&[6]int8{'%', '.', '1', '5', 'f', '\x00'})), 1.1000000000000001, (*int8)(unsafe.Pointer(&[18]int8{'1', '.', '1', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '\x00'}))}, _cgoa_2{(*int8)(unsafe.Pointer(&[6]int8{'%', '.', '1', '6', 'f', '\x00'})), 1.1000000000000001, (*int8)(unsafe.Pointer(&[19]int8{'1', '.', '1', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '1', '\x00'}))}, _cgoa_2{(*int8)(unsafe.Pointer(&[6]int8{'%', '.', '1', '7', 'f', '\x00'})), 1.1000000000000001, (*int8)(unsafe.Pointer(&[20]int8{'1', '.', '1', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '9', '\x00'}))}, _cgoa_2{(*int8)(unsafe.Pointer(&[5]int8{'%', '.', '2', 'e', '\x00'})), 1500001, (*int8)(unsafe.Pointer(&[9]int8{'1', '.', '5', '0', 'e', '+', '0', '6', '\x00'}))}, _cgoa_2{(*int8)(unsafe.Pointer(&[5]int8{'%', '.', '2', 'e', '\x00'})), 1505000, (*int8)(unsafe.Pointer(&[9]int8{'1', '.', '5', '0', 'e', '+', '0', '6', '\x00'}))}, _cgoa_2{(*int8)(unsafe.Pointer(&[5]int8{'%', '.', '2', 'e', '\x00'})), 1505000.0000009537, (*int8)(unsafe.Pointer(&[9]int8{'1', '.', '5', '1', 'e', '+', '0', '6', '\x00'}))}, _cgoa_2{(*int8)(unsafe.Pointer(&[5]int8{'%', '.', '2', 'e', '\x00'})), 1505001, (*int8)(unsafe.Pointer(&[9]int8{'1', '.', '5', '1', 'e', '+', '0', '6', '\x00'}))}, _cgoa_2{(*int8)(unsafe.Pointer(&[5]int8{'%', '.', '2', 'e', '\x00'})), 1506000, (*int8)(unsafe.Pointer(&[9]int8{'1', '.', '5', '1', 'e', '+', '0', '6', '\x00'}))}, _cgoa_2{(*int8)(unsafe.Pointer(&[6]int8{'%', '.', '1', '5', 'g', '\x00'})), 1.23456789012345, (*int8)(unsafe.Pointer(&[17]int8{'1', '.', '2', '3', '4', '5', '6', '7', '8', '9', '0', '1', '2', '3', '4', '5', '\x00'}))}, _cgoa_2{(*int8)(unsafe.Pointer(&[3]int8{'%', 'g', '\x00'})), 1.0e-4, (*int8)(unsafe.Pointer(&[7]int8{'0', '.', '0', '0', '0', '1', '\x00'}))}, _cgoa_2{(*int8)(unsafe.Pointer(&[3]int8{'%', 'g', '\x00'})), 1.0000000000000001e-5, (*int8)(unsafe.Pointer(&[6]int8{'1', 'e', '-', '0', '5', '\x00'}))}, _cgoa_2{(*int8)(unsafe.Pointer(&[3]int8{'%', 'g', '\x00'})), float64(int32(123456)), (*int8)(unsafe.Pointer(&[7]int8{'1', '2', '3', '4', '5', '6', '\x00'}))}, _cgoa_2{(*int8)(unsafe.Pointer(&[3]int8{'%', 'g', '\x00'})), float64(int32(1234567)), (*int8)(unsafe.Pointer(&[12]int8{'1', '.', '2', '3', '4', '5', '7', 'e', '+', '0', '6', '\x00'}))}, _cgoa_2{(*int8)(unsafe.Pointer(&[5]int8{'%', '.', '7', 'g', '\x00'})), float64(int32(1234567)), (*int8)(unsafe.Pointer(&[8]int8{'1', '2', '3', '4', '5', '6', '7', '\x00'}))}, _cgoa_2{(*int8)(unsafe.Pointer(&[5]int8{'%', '.', '7', 'g', '\x00'})), float64(int32(12345678)), (*int8)(unsafe.Pointer(&[13]int8{'1', '.', '2', '3', '4', '5', '6', '8', 'e', '+', '0', '7', '\x00'}))}, _cgoa_2{(*int8)(unsafe.Pointer(&[5]int8{'%', '.', '8', 'g', '\x00'})), 0.10000000000000001, (*int8)(unsafe.Pointer(&[4]int8{'0', '.', '1', '\x00'}))}, _cgoa_2{(*int8)(unsafe.Pointer(&[5]int8{'%', '.', '9', 'g', '\x00'})), 0.10000000000000001, (*int8)(unsafe.Pointer(&[4]int8{'0', '.', '1', '\x00'}))}, _cgoa_2{(*int8)(unsafe.Pointer(&[6]int8{'%', '.', '1', '0', 'g', '\x00'})), 0.10000000000000001, (*int8)(unsafe.Pointer(&[4]int8{'0', '.', '1', '\x00'}))}, _cgoa_2{(*int8)(unsafe.Pointer(&[6]int8{'%', '.', '1', '1', 'g', '\x00'})), 0.10000000000000001, (*int8)(unsafe.Pointer(&[4]int8{'0', '.', '1', '\x00'}))}, _cgoa_2{(*int8)(unsafe.Pointer(&[6]int8{'%', '.', '1', '5', 'f', '\x00'})), 3.1415926535897931, (*int8)(unsafe.Pointer(&[18]int8{'3', '.', '1', '4', '1', '5', '9', '2', '6', '5', '3', '5', '8', '9', '7', '9', '3', '\x00'}))}, _cgoa_2{(*int8)(unsafe.Pointer(&[6]int8{'%', '.', '1', '8', 'f', '\x00'})), 3.1415926535897931, (*int8)(unsafe.Pointer(&[21]int8{'3', '.', '1', '4', '1', '5', '9', '2', '6', '5', '3', '5', '8', '9', '7', '9', '3', '1', '1', '6', '\x00'}))}, _cgoa_2{(*int8)(unsafe.Pointer(&[5]int8{'%', '.', '0', 'f', '\x00'})), 3.4028236692093846e+38, (*int8)(unsafe.Pointer(&[40]int8{'3', '4', '0', '2', '8', '2', '3', '6', '6', '9', '2', '0', '9', '3', '8', '4', '6', '3', '4', '6', '3', '3', '7', '4', '6', '0', '7', '4', '3', '1', '7', '6', '8', '2', '1', '1', '4', '5', '6', '\x00'}))}, _cgoa_2{nil, 0, nil}}

func _cgo_main() int32 {
	var i int32
	var j int32
	var k int32
	var b [2000]int8
	if !(func() (_cgo_ret int32) {
		_cgo_addr := &i
		*_cgo_addr = libc.Snprintf(nil, uint64(int32(0)), (*int8)(unsafe.Pointer(&[3]int8{'%', 'd', '\x00'})), int32(123456))
		return *_cgo_addr
	}() == int32(6)) {
		func() int32 {
			t_printf((*int8)(unsafe.Pointer(&[69]int8{'s', 'r', 'c', '/', 'f', 'u', 'n', 'c', 't', 'i', 'o', 'n', 'a', 'l', '/', 's', 'n', 'p', 'r', 'i', 'n', 't', 'f', '.', 'c', ':', '1', '4', '8', ':', ' ', '%', 's', ' ', 'f', 'a', 'i', 'l', 'e', 'd', ' ', '(', 'l', 'e', 'n', 'g', 't', 'h', ' ', 'r', 'e', 't', 'u', 'r', 'n', 'e', 'd', ' ', '%', 'd', ' ', '!', '=', ' ', '%', 'd', ')', '\n', '\x00'})), (*int8)(unsafe.Pointer(&[29]int8{'s', 'n', 'p', 'r', 'i', 'n', 't', 'f', '(', '0', ',', ' ', '0', ',', ' ', '"', '%', 'd', '"', ',', ' ', '1', '2', '3', '4', '5', '6', ')', '\x00'})), i, int32(6))
			return int32(0)
		}()
	}
	if !(func() (_cgo_ret int32) {
		_cgo_addr := &i
		*_cgo_addr = libc.Snprintf(nil, uint64(int32(0)), (*int8)(unsafe.Pointer(&[5]int8{'%', '.', '4', 's', '\x00'})), (*int8)(unsafe.Pointer(&[6]int8{'h', 'e', 'l', 'l', 'o', '\x00'})))
		return *_cgo_addr
	}() == int32(4)) {
		func() int32 {
			t_printf((*int8)(unsafe.Pointer(&[69]int8{'s', 'r', 'c', '/', 'f', 'u', 'n', 'c', 't', 'i', 'o', 'n', 'a', 'l', '/', 's', 'n', 'p', 'r', 'i', 'n', 't', 'f', '.', 'c', ':', '1', '4', '9', ':', ' ', '%', 's', ' ', 'f', 'a', 'i', 'l', 'e', 'd', ' ', '(', 'l', 'e', 'n', 'g', 't', 'h', ' ', 'r', 'e', 't', 'u', 'r', 'n', 'e', 'd', ' ', '%', 'd', ' ', '!', '=', ' ', '%', 'd', ')', '\n', '\x00'})), (*int8)(unsafe.Pointer(&[32]int8{'s', 'n', 'p', 'r', 'i', 'n', 't', 'f', '(', '0', ',', ' ', '0', ',', ' ', '"', '%', '.', '4', 's', '"', ',', ' ', '"', 'h', 'e', 'l', 'l', 'o', '"', ')', '\x00'})), i, int32(4))
			return int32(0)
		}()
	}
	if !(func() (_cgo_ret int32) {
		_cgo_addr := &i
		*_cgo_addr = libc.Snprintf((*int8)(unsafe.Pointer(&b)), uint64(int32(0)), (*int8)(unsafe.Pointer(&[5]int8{'%', '.', '0', 's', '\x00'})), (*int8)(unsafe.Pointer(&[8]int8{'g', 'o', 'o', 'd', 'b', 'y', 'e', '\x00'})))
		return *_cgo_addr
	}() == int32(0)) {
		func() int32 {
			t_printf((*int8)(unsafe.Pointer(&[69]int8{'s', 'r', 'c', '/', 'f', 'u', 'n', 'c', 't', 'i', 'o', 'n', 'a', 'l', '/', 's', 'n', 'p', 'r', 'i', 'n', 't', 'f', '.', 'c', ':', '1', '5', '0', ':', ' ', '%', 's', ' ', 'f', 'a', 'i', 'l', 'e', 'd', ' ', '(', 'l', 'e', 'n', 'g', 't', 'h', ' ', 'r', 'e', 't', 'u', 'r', 'n', 'e', 'd', ' ', '%', 'd', ' ', '!', '=', ' ', '%', 'd', ')', '\n', '\x00'})), (*int8)(unsafe.Pointer(&[34]int8{'s', 'n', 'p', 'r', 'i', 'n', 't', 'f', '(', 'b', ',', ' ', '0', ',', ' ', '"', '%', '.', '0', 's', '"', ',', ' ', '"', 'g', 'o', 'o', 'd', 'b', 'y', 'e', '"', ')', '\x00'})), i, int32(0))
			return int32(0)
		}()
	}
	libc.Strcpy((*int8)(unsafe.Pointer(&b)), (*int8)(unsafe.Pointer(&[9]int8{'x', 'x', 'x', 'x', 'x', 'x', 'x', 'x', '\x00'})))
	if !(func() (_cgo_ret int32) {
		_cgo_addr := &i
		*_cgo_addr = libc.Snprintf((*int8)(unsafe.Pointer(&b)), uint64(int32(4)), (*int8)(unsafe.Pointer(&[3]int8{'%', 'd', '\x00'})), int32(123456))
		return *_cgo_addr
	}() == int32(6)) {
		func() int32 {
			t_printf((*int8)(unsafe.Pointer(&[69]int8{'s', 'r', 'c', '/', 'f', 'u', 'n', 'c', 't', 'i', 'o', 'n', 'a', 'l', '/', 's', 'n', 'p', 'r', 'i', 'n', 't', 'f', '.', 'c', ':', '1', '5', '3', ':', ' ', '%', 's', ' ', 'f', 'a', 'i', 'l', 'e', 'd', ' ', '(', 'l', 'e', 'n', 'g', 't', 'h', ' ', 'r', 'e', 't', 'u', 'r', 'n', 'e', 'd', ' ', '%', 'd', ' ', '!', '=', ' ', '%', 'd', ')', '\n', '\x00'})), (*int8)(unsafe.Pointer(&[29]int8{'s', 'n', 'p', 'r', 'i', 'n', 't', 'f', '(', 'b', ',', ' ', '4', ',', ' ', '"', '%', 'd', '"', ',', ' ', '1', '2', '3', '4', '5', '6', ')', '\x00'})), i, int32(6))
			return int32(0)
		}()
	}
	if !!(libc.Strcmp((*int8)(unsafe.Pointer(&b)), (*int8)(unsafe.Pointer(&[4]int8{'1', '2', '3', '\x00'}))) != 0) {
		func() int32 {
			t_printf((*int8)(unsafe.Pointer(&[50]int8{'s', 'r', 'c', '/', 'f', 'u', 'n', 'c', 't', 'i', 'o', 'n', 'a', 'l', '/', 's', 'n', 'p', 'r', 'i', 'n', 't', 'f', '.', 'c', ':', '1', '5', '4', ':', ' ', '[', '%', 's', ']', ' ', '!', '=', ' ', '[', '%', 's', ']', ' ', '(', '%', 's', ')', '\n', '\x00'})), (*int8)(unsafe.Pointer(&b)), (*int8)(unsafe.Pointer(&[4]int8{'1', '2', '3', '\x00'})), (*int8)(unsafe.Pointer(&[17]int8{'i', 'n', 'c', 'o', 'r', 'r', 'e', 'c', 't', ' ', 'o', 'u', 't', 'p', 'u', 't', '\x00'})))
			return int32(0)
		}()
	}
	if !(func() (_cgo_ret int32) {
		_cgo_addr := &i
		*_cgo_addr = int32(*(*int8)(unsafe.Pointer(uintptr(unsafe.Pointer((*int8)(unsafe.Pointer(&b)))) + uintptr(int32(5)))))
		return *_cgo_addr
	}() == 'x') {
		func() int32 {
			t_printf((*int8)(unsafe.Pointer(&[59]int8{'s', 'r', 'c', '/', 'f', 'u', 'n', 'c', 't', 'i', 'o', 'n', 'a', 'l', '/', 's', 'n', 'p', 'r', 'i', 'n', 't', 'f', '.', 'c', ':', '1', '5', '5', ':', ' ', '%', 's', ' ', 'f', 'a', 'i', 'l', 'e', 'd', ' ', '(', 'b', 'u', 'f', 'f', 'e', 'r', ' ', 'o', 'v', 'e', 'r', 'r', 'u', 'n', ')', '\n', '\x00'})), (*int8)(unsafe.Pointer(&[5]int8{'b', '[', '5', ']', '\x00'})), i, 'x')
			return int32(0)
		}()
	}
	if !(func() (_cgo_ret int32) {
		_cgo_addr := &i
		*_cgo_addr = libc.Snprintf((*int8)(unsafe.Pointer(&b)), 2000, (*int8)(unsafe.Pointer(&[8]int8{'%', '.', '1', '0', '2', '2', 'f', '\x00'})), 4.4501477170144028e-308)
		return *_cgo_addr
	}() == int32(1024)) {
		func() int32 {
			t_printf((*int8)(unsafe.Pointer(&[53]int8{'s', 'r', 'c', '/', 'f', 'u', 'n', 'c', 't', 'i', 'o', 'n', 'a', 'l', '/', 's', 'n', 'p', 'r', 'i', 'n', 't', 'f', '.', 'c', ':', '1', '5', '8', ':', ' ', '%', 's', ' ', 'f', 'a', 'i', 'l', 'e', 'd', ' ', '(', '%', 'd', ' ', '!', '=', ' ', '%', 'd', ')', '\n', '\x00'})), (*int8)(unsafe.Pointer(&[44]int8{'s', 'n', 'p', 'r', 'i', 'n', 't', 'f', '(', 'b', ',', ' ', 's', 'i', 'z', 'e', 'o', 'f', ' ', 'b', ',', ' ', '"', '%', '.', '1', '0', '2', '2', 'f', '"', ',', ' ', '0', 'x', '1', 'p', '-', '1', '0', '2', '1', ')', '\x00'})), i, int32(1024))
			return int32(0)
		}()
	}
	*(*int8)(unsafe.Pointer(uintptr(unsafe.Pointer((*int8)(unsafe.Pointer(&b)))) + uintptr(int32(1)))) = int8('0')
	for i = int32(0); i < int32(1021); i++ {
		for func() int32 {
			k = int32(0)
			return func() (_cgo_ret int32) {
				_cgo_addr := &j
				*_cgo_addr = int32(1023)
				return *_cgo_addr
			}()
		}(); j > int32(0); j-- {
			if int32(*(*int8)(unsafe.Pointer(uintptr(unsafe.Pointer((*int8)(unsafe.Pointer(&b)))) + uintptr(j)))) < '5' {
				func() int32 {
					*(*int8)(unsafe.Pointer(uintptr(unsafe.Pointer((*int8)(unsafe.Pointer(&b)))) + uintptr(j))) += int8(int32(*(*int8)(unsafe.Pointer(uintptr(unsafe.Pointer((*int8)(unsafe.Pointer(&b)))) + uintptr(j)))) - '0' + k)
					return func() (_cgo_ret int32) {
						_cgo_addr := &k
						*_cgo_addr = int32(0)
						return *_cgo_addr
					}()
				}()
			} else {
				func() int32 {
					*(*int8)(unsafe.Pointer(uintptr(unsafe.Pointer((*int8)(unsafe.Pointer(&b)))) + uintptr(j))) += int8(int32(*(*int8)(unsafe.Pointer(uintptr(unsafe.Pointer((*int8)(unsafe.Pointer(&b)))) + uintptr(j)))) - '0' - int32(10) + k)
					return func() (_cgo_ret int32) {
						_cgo_addr := &k
						*_cgo_addr = int32(1)
						return *_cgo_addr
					}()
				}()
			}
		}
	}
	if !(func() (_cgo_ret int32) {
		_cgo_addr := &i
		*_cgo_addr = int32(*(*int8)(unsafe.Pointer(uintptr(unsafe.Pointer((*int8)(unsafe.Pointer(&b)))) + uintptr(int32(1)))))
		return *_cgo_addr
	}() == '1') {
		func() int32 {
			t_printf((*int8)(unsafe.Pointer(&[57]int8{'s', 'r', 'c', '/', 'f', 'u', 'n', 'c', 't', 'i', 'o', 'n', 'a', 'l', '/', 's', 'n', 'p', 'r', 'i', 'n', 't', 'f', '.', 'c', ':', '1', '6', '6', ':', ' ', '%', 's', ' ', 'f', 'a', 'i', 'l', 'e', 'd', ' ', '(', '\'', '%', 'c', '\'', ' ', '!', '=', ' ', '\'', '%', 'c', '\'', ')', '\n', '\x00'})), (*int8)(unsafe.Pointer(&[5]int8{'b', '[', '1', ']', '\x00'})), i, '1')
			return int32(0)
		}()
	}
	for j = int32(2); int32(*(*int8)(unsafe.Pointer(uintptr(unsafe.Pointer((*int8)(unsafe.Pointer(&b)))) + uintptr(j)))) == '0'; j++ {
	}
	if !(func() (_cgo_ret int32) {
		_cgo_addr := &i
		*_cgo_addr = j
		return *_cgo_addr
	}() == int32(1024)) {
		func() int32 {
			t_printf((*int8)(unsafe.Pointer(&[53]int8{'s', 'r', 'c', '/', 'f', 'u', 'n', 'c', 't', 'i', 'o', 'n', 'a', 'l', '/', 's', 'n', 'p', 'r', 'i', 'n', 't', 'f', '.', 'c', ':', '1', '6', '8', ':', ' ', '%', 's', ' ', 'f', 'a', 'i', 'l', 'e', 'd', ' ', '(', '%', 'd', ' ', '!', '=', ' ', '%', 'd', ')', '\n', '\x00'})), (*int8)(unsafe.Pointer(&[2]int8{'j', '\x00'})), i, int32(1024))
			return int32(0)
		}()
	}
	for j = int32(0); (*(*_cgoa_1)(unsafe.Pointer(uintptr(unsafe.Pointer((*_cgoa_1)(unsafe.Pointer(&int_tests)))) + uintptr(j)*24))).fmt != nil; j++ {
		i = libc.Snprintf((*int8)(unsafe.Pointer(&b)), 2000, (*(*_cgoa_1)(unsafe.Pointer(uintptr(unsafe.Pointer((*_cgoa_1)(unsafe.Pointer(&int_tests)))) + uintptr(j)*24))).fmt, (*(*_cgoa_1)(unsafe.Pointer(uintptr(unsafe.Pointer((*_cgoa_1)(unsafe.Pointer(&int_tests)))) + uintptr(j)*24))).i)
		if uint64(i) != libc.Strlen((*(*_cgoa_1)(unsafe.Pointer(uintptr(unsafe.Pointer((*_cgoa_1)(unsafe.Pointer(&int_tests)))) + uintptr(j)*24))).expect) {
			t_printf((*int8)(unsafe.Pointer(&[86]int8{'s', 'r', 'c', '/', 'f', 'u', 'n', 'c', 't', 'i', 'o', 'n', 'a', 'l', '/', 's', 'n', 'p', 'r', 'i', 'n', 't', 'f', '.', 'c', ':', '1', '8', '1', ':', ' ', 's', 'n', 'p', 'r', 'i', 'n', 't', 'f', '(', 'b', ',', ' ', 's', 'i', 'z', 'e', 'o', 'f', ' ', 'b', ',', ' ', '"', '%', 's', '"', ',', ' ', '%', 'd', ')', ' ', 'r', 'e', 't', 'u', 'r', 'n', 'e', 'd', ' ', '%', 'd', ' ', 'w', 'a', 'n', 't', 'e', 'd', ' ', '%', 'd', '\n', '\x00'})), (*(*_cgoa_1)(unsafe.Pointer(uintptr(unsafe.Pointer((*_cgoa_1)(unsafe.Pointer(&int_tests)))) + uintptr(j)*24))).fmt, (*(*_cgoa_1)(unsafe.Pointer(uintptr(unsafe.Pointer((*_cgoa_1)(unsafe.Pointer(&int_tests)))) + uintptr(j)*24))).i, i, libc.Strlen((*(*_cgoa_1)(unsafe.Pointer(uintptr(unsafe.Pointer((*_cgoa_1)(unsafe.Pointer(&int_tests)))) + uintptr(j)*24))).expect))
		}
		if libc.Strcmp((*int8)(unsafe.Pointer(&b)), (*(*_cgoa_1)(unsafe.Pointer(uintptr(unsafe.Pointer((*_cgoa_1)(unsafe.Pointer(&int_tests)))) + uintptr(j)*24))).expect) != int32(0) {
			t_printf((*int8)(unsafe.Pointer(&[76]int8{'s', 'r', 'c', '/', 'f', 'u', 'n', 'c', 't', 'i', 'o', 'n', 'a', 'l', '/', 's', 'n', 'p', 'r', 'i', 'n', 't', 'f', '.', 'c', ':', '1', '8', '4', ':', ' ', 'b', 'a', 'd', ' ', 'i', 'n', 't', 'e', 'g', 'e', 'r', ' ', 'c', 'o', 'n', 'v', 'e', 'r', 's', 'i', 'o', 'n', ':', ' ', 'g', 'o', 't', ' ', '"', '%', 's', '"', ',', ' ', 'w', 'a', 'n', 't', ' ', '"', '%', 's', '"', '\n', '\x00'})), (*int8)(unsafe.Pointer(&b)), (*(*_cgoa_1)(unsafe.Pointer(uintptr(unsafe.Pointer((*_cgoa_1)(unsafe.Pointer(&int_tests)))) + uintptr(j)*24))).expect)
		}
	}
	for j = int32(0); (*(*_cgoa_2)(unsafe.Pointer(uintptr(unsafe.Pointer((*_cgoa_2)(unsafe.Pointer(&fp_tests)))) + uintptr(j)*24))).fmt != nil; j++ {
		i = libc.Snprintf((*int8)(unsafe.Pointer(&b)), 2000, (*(*_cgoa_2)(unsafe.Pointer(uintptr(unsafe.Pointer((*_cgoa_2)(unsafe.Pointer(&fp_tests)))) + uintptr(j)*24))).fmt, (*(*_cgoa_2)(unsafe.Pointer(uintptr(unsafe.Pointer((*_cgoa_2)(unsafe.Pointer(&fp_tests)))) + uintptr(j)*24))).f)
		if uint64(i) != libc.Strlen((*(*_cgoa_2)(unsafe.Pointer(uintptr(unsafe.Pointer((*_cgoa_2)(unsafe.Pointer(&fp_tests)))) + uintptr(j)*24))).expect) {
			t_printf((*int8)(unsafe.Pointer(&[86]int8{'s', 'r', 'c', '/', 'f', 'u', 'n', 'c', 't', 'i', 'o', 'n', 'a', 'l', '/', 's', 'n', 'p', 'r', 'i', 'n', 't', 'f', '.', 'c', ':', '1', '9', '1', ':', ' ', 's', 'n', 'p', 'r', 'i', 'n', 't', 'f', '(', 'b', ',', ' ', 's', 'i', 'z', 'e', 'o', 'f', ' ', 'b', ',', ' ', '"', '%', 's', '"', ',', ' ', '%', 'f', ')', ' ', 'r', 'e', 't', 'u', 'r', 'n', 'e', 'd', ' ', '%', 'd', ' ', 'w', 'a', 'n', 't', 'e', 'd', ' ', '%', 'd', '\n', '\x00'})), (*(*_cgoa_2)(unsafe.Pointer(uintptr(unsafe.Pointer((*_cgoa_2)(unsafe.Pointer(&fp_tests)))) + uintptr(j)*24))).fmt, (*(*_cgoa_2)(unsafe.Pointer(uintptr(unsafe.Pointer((*_cgoa_2)(unsafe.Pointer(&fp_tests)))) + uintptr(j)*24))).f, i, libc.Strlen((*(*_cgoa_2)(unsafe.Pointer(uintptr(unsafe.Pointer((*_cgoa_2)(unsafe.Pointer(&fp_tests)))) + uintptr(j)*24))).expect))
		}
		if libc.Strcmp((*int8)(unsafe.Pointer(&b)), (*(*_cgoa_2)(unsafe.Pointer(uintptr(unsafe.Pointer((*_cgoa_2)(unsafe.Pointer(&fp_tests)))) + uintptr(j)*24))).expect) != int32(0) {
			t_printf((*int8)(unsafe.Pointer(&[83]int8{'s', 'r', 'c', '/', 'f', 'u', 'n', 'c', 't', 'i', 'o', 'n', 'a', 'l', '/', 's', 'n', 'p', 'r', 'i', 'n', 't', 'f', '.', 'c', ':', '1', '9', '4', ':', ' ', 'b', 'a', 'd', ' ', 'f', 'l', 'o', 'a', 't', 'i', 'n', 'g', '-', 'p', 'o', 'i', 'n', 't', ' ', 'c', 'o', 'n', 'v', 'e', 'r', 's', 'i', 'o', 'n', ':', ' ', 'g', 'o', 't', ' ', '"', '%', 's', '"', ',', ' ', 'w', 'a', 'n', 't', ' ', '"', '%', 's', '"', '\n', '\x00'})), (*int8)(unsafe.Pointer(&b)), (*(*_cgoa_2)(unsafe.Pointer(uintptr(unsafe.Pointer((*_cgoa_2)(unsafe.Pointer(&fp_tests)))) + uintptr(j)*24))).expect)
		}
	}
	if !(func() (_cgo_ret int32) {
		_cgo_addr := &i
		*_cgo_addr = libc.Snprintf(nil, uint64(int32(0)), (*int8)(unsafe.Pointer(&[5]int8{'%', '.', '4', 'a', '\x00'})), 1)
		return *_cgo_addr
	}() == int32(11)) {
		func() int32 {
			t_printf((*int8)(unsafe.Pointer(&[53]int8{'s', 'r', 'c', '/', 'f', 'u', 'n', 'c', 't', 'i', 'o', 'n', 'a', 'l', '/', 's', 'n', 'p', 'r', 'i', 'n', 't', 'f', '.', 'c', ':', '1', '9', '7', ':', ' ', '%', 's', ' ', 'f', 'a', 'i', 'l', 'e', 'd', ' ', '(', '%', 'd', ' ', '!', '=', ' ', '%', 'd', ')', '\n', '\x00'})), (*int8)(unsafe.Pointer(&[28]int8{'s', 'n', 'p', 'r', 'i', 'n', 't', 'f', '(', '0', ',', ' ', '0', ',', ' ', '"', '%', '.', '4', 'a', '"', ',', ' ', '1', '.', '0', ')', '\x00'})), i, int32(11))
			return int32(0)
		}()
	}
	return t_status
}
func main() {
	os.Exit(int(_cgo_main()))
}
