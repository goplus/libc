package fminf

import (
	common "github.com/goplus/libc/test/common"
	unsafe "unsafe"
	libc "github.com/goplus/libc"
	testing "testing"
)

var _cgos_t_fminf [68]common.Struct_ff_f = [68]common.Struct_ff_f{common.Struct_ff_f{(*int8)(unsafe.Pointer(&[55]int8{'/', 'U', 's', 'e', 'r', 's', '/', 'x', 'u', 's', 'h', 'i', 'w', 'e', 'i', '/', 'w', 'o', 'r', 'k', '/', 'l', 'i', 'b', 'c', '/', 't', 'e', 's', 't', '/', 's', 'r', 'c', '/', 'm', 'a', 't', 'h', '/', 's', 'a', 'n', 'i', 't', 'y', '/', 'f', 'm', 'i', 'n', 'f', '.', 'h', '\x00'})), int32(1), int32(0), float32(-8.0668487548828125), float32(4.5356626510620117), float32(-8.0668487548828125), float32(0), int32(0)}, common.Struct_ff_f{(*int8)(unsafe.Pointer(&[55]int8{'/', 'U', 's', 'e', 'r', 's', '/', 'x', 'u', 's', 'h', 'i', 'w', 'e', 'i', '/', 'w', 'o', 'r', 'k', '/', 'l', 'i', 'b', 'c', '/', 't', 'e', 's', 't', '/', 's', 'r', 'c', '/', 'm', 'a', 't', 'h', '/', 's', 'a', 'n', 'i', 't', 'y', '/', 'f', 'm', 'i', 'n', 'f', '.', 'h', '\x00'})), int32(2), int32(0), float32(4.3452396392822266), float32(-8.8879909515380859), float32(-8.8879909515380859), float32(0), int32(0)}, common.Struct_ff_f{(*int8)(unsafe.Pointer(&[55]int8{'/', 'U', 's', 'e', 'r', 's', '/', 'x', 'u', 's', 'h', 'i', 'w', 'e', 'i', '/', 'w', 'o', 'r', 'k', '/', 'l', 'i', 'b', 'c', '/', 't', 'e', 's', 't', '/', 's', 'r', 'c', '/', 'm', 'a', 't', 'h', '/', 's', 'a', 'n', 'i', 't', 'y', '/', 'f', 'm', 'i', 'n', 'f', '.', 'h', '\x00'})), int32(3), int32(0), float32(-8.3814334869384765), float32(-2.7636072635650635), float32(-8.3814334869384765), float32(0), int32(0)}, common.Struct_ff_f{(*int8)(unsafe.Pointer(&[55]int8{'/', 'U', 's', 'e', 'r', 's', '/', 'x', 'u', 's', 'h', 'i', 'w', 'e', 'i', '/', 'w', 'o', 'r', 'k', '/', 'l', 'i', 'b', 'c', '/', 't', 'e', 's', 't', '/', 's', 'r', 'c', '/', 'm', 'a', 't', 'h', '/', 's', 'a', 'n', 'i', 't', 'y', '/', 'f', 'm', 'i', 'n', 'f', '.', 'h', '\x00'})), int32(4), int32(0), float32(-6.5316734313964844), float32(4.567535400390625), float32(-6.5316734313964844), float32(0), int32(0)}, common.Struct_ff_f{(*int8)(unsafe.Pointer(&[55]int8{'/', 'U', 's', 'e', 'r', 's', '/', 'x', 'u', 's', 'h', 'i', 'w', 'e', 'i', '/', 'w', 'o', 'r', 'k', '/', 'l', 'i', 'b', 'c', '/', 't', 'e', 's', 't', '/', 's', 'r', 'c', '/', 'm', 'a', 't', 'h', '/', 's', 'a', 'n', 'i', 't', 'y', '/', 'f', 'm', 'i', 'n', 'f', '.', 'h', '\x00'})), int32(5), int32(0), float32(9.2670574188232421), float32(4.8113923072814941), float32(4.8113923072814941), float32(0), int32(0)}, common.Struct_ff_f{(*int8)(unsafe.Pointer(&[55]int8{'/', 'U', 's', 'e', 'r', 's', '/', 'x', 'u', 's', 'h', 'i', 'w', 'e', 'i', '/', 'w', 'o', 'r', 'k', '/', 'l', 'i', 'b', 'c', '/', 't', 'e', 's', 't', '/', 's', 'r', 'c', '/', 'm', 'a', 't', 'h', '/', 's', 'a', 'n', 'i', 't', 'y', '/', 'f', 'm', 'i', 'n', 'f', '.', 'h', '\x00'})), int32(6), int32(0), float32(-6.4500455856323242), float32(0.66207176446914673), float32(-6.4500455856323242), float32(0), int32(0)}, common.Struct_ff_f{(*int8)(unsafe.Pointer(&[55]int8{'/', 'U', 's', 'e', 'r', 's', '/', 'x', 'u', 's', 'h', 'i', 'w', 'e', 'i', '/', 'w', 'o', 'r', 'k', '/', 'l', 'i', 'b', 'c', '/', 't', 'e', 's', 't', '/', 's', 'r', 'c', '/', 'm', 'a', 't', 'h', '/', 's', 'a', 'n', 'i', 't', 'y', '/', 'f', 'm', 'i', 'n', 'f', '.', 'h', '\x00'})), int32(7), int32(0), float32(7.8588900566101074), float32(0.052154526114463806), float32(0.052154526114463806), float32(0), int32(0)}, common.Struct_ff_f{(*int8)(unsafe.Pointer(&[55]int8{'/', 'U', 's', 'e', 'r', 's', '/', 'x', 'u', 's', 'h', 'i', 'w', 'e', 'i', '/', 'w', 'o', 'r', 'k', '/', 'l', 'i', 'b', 'c', '/', 't', 'e', 's', 't', '/', 's', 'r', 'c', '/', 'm', 'a', 't', 'h', '/', 's', 'a', 'n', 'i', 't', 'y', '/', 'f', 'm', 'i', 'n', 'f', '.', 'h', '\x00'})), int32(8), int32(0), float32(-0.79205453395843505), float32(7.6764025688171387), float32(-0.79205453395843505), float32(0), int32(0)}, common.Struct_ff_f{(*int8)(unsafe.Pointer(&[55]int8{'/', 'U', 's', 'e', 'r', 's', '/', 'x', 'u', 's', 'h', 'i', 'w', 'e', 'i', '/', 'w', 'o', 'r', 'k', '/', 'l', 'i', 'b', 'c', '/', 't', 'e', 's', 't', '/', 's', 'r', 'c', '/', 'm', 'a', 't', 'h', '/', 's', 'a', 'n', 'i', 't', 'y', '/', 'f', 'm', 'i', 'n', 'f', '.', 'h', '\x00'})), int32(9), int32(0), float32(0.61570268869400024), float32(2.0119025707244873), float32(0.61570268869400024), float32(0), int32(0)}, common.Struct_ff_f{(*int8)(unsafe.Pointer(&[55]int8{'/', 'U', 's', 'e', 'r', 's', '/', 'x', 'u', 's', 'h', 'i', 'w', 'e', 'i', '/', 'w', 'o', 'r', 'k', '/', 'l', 'i', 'b', 'c', '/', 't', 'e', 's', 't', '/', 's', 'r', 'c', '/', 'm', 'a', 't', 'h', '/', 's', 'a', 'n', 'i', 't', 'y', '/', 'f', 'm', 'i', 'n', 'f', '.', 'h', '\x00'})), int32(10), int32(0), float32(-0.55875867605209351), float32(0.032239831984043121), float32(-0.55875867605209351), float32(0), int32(0)}, common.Struct_ff_f{(*int8)(unsafe.Pointer(&[56]int8{'/', 'U', 's', 'e', 'r', 's', '/', 'x', 'u', 's', 'h', 'i', 'w', 'e', 'i', '/', 'w', 'o', 'r', 'k', '/', 'l', 'i', 'b', 'c', '/', 't', 'e', 's', 't', '/', 's', 'r', 'c', '/', 'm', 'a', 't', 'h', '/', 's', 'p', 'e', 'c', 'i', 'a', 'l', '/', 'f', 'm', 'i', 'n', 'f', '.', 'h', '\x00'})), int32(1), int32(0), float32(0), float32(1), float32(0), float32(0), int32(0)}, common.Struct_ff_f{(*int8)(unsafe.Pointer(&[56]int8{'/', 'U', 's', 'e', 'r', 's', '/', 'x', 'u', 's', 'h', 'i', 'w', 'e', 'i', '/', 'w', 'o', 'r', 'k', '/', 'l', 'i', 'b', 'c', '/', 't', 'e', 's', 't', '/', 's', 'r', 'c', '/', 'm', 'a', 't', 'h', '/', 's', 'p', 'e', 'c', 'i', 'a', 'l', '/', 'f', 'm', 'i', 'n', 'f', '.', 'h', '\x00'})), int32(2), int32(0), float32(-0), float32(1), float32(-0), float32(0), int32(0)}, common.Struct_ff_f{(*int8)(unsafe.Pointer(&[56]int8{'/', 'U', 's', 'e', 'r', 's', '/', 'x', 'u', 's', 'h', 'i', 'w', 'e', 'i', '/', 'w', 'o', 'r', 'k', '/', 'l', 'i', 'b', 'c', '/', 't', 'e', 's', 't', '/', 's', 'r', 'c', '/', 'm', 'a', 't', 'h', '/', 's', 'p', 'e', 'c', 'i', 'a', 'l', '/', 'f', 'm', 'i', 'n', 'f', '.', 'h', '\x00'})), int32(3), int32(0), float32(0.5), float32(1), float32(0.5), float32(0), int32(0)}, common.Struct_ff_f{(*int8)(unsafe.Pointer(&[56]int8{'/', 'U', 's', 'e', 'r', 's', '/', 'x', 'u', 's', 'h', 'i', 'w', 'e', 'i', '/', 'w', 'o', 'r', 'k', '/', 'l', 'i', 'b', 'c', '/', 't', 'e', 's', 't', '/', 's', 'r', 'c', '/', 'm', 'a', 't', 'h', '/', 's', 'p', 'e', 'c', 'i', 'a', 'l', '/', 'f', 'm', 'i', 'n', 'f', '.', 'h', '\x00'})), int32(4), int32(0), float32(-0.5), float32(1), float32(-0.5), float32(0), int32(0)}, common.Struct_ff_f{(*int8)(unsafe.Pointer(&[56]int8{'/', 'U', 's', 'e', 'r', 's', '/', 'x', 'u', 's', 'h', 'i', 'w', 'e', 'i', '/', 'w', 'o', 'r', 'k', '/', 'l', 'i', 'b', 'c', '/', 't', 'e', 's', 't', '/', 's', 'r', 'c', '/', 'm', 'a', 't', 'h', '/', 's', 'p', 'e', 'c', 'i', 'a', 'l', '/', 'f', 'm', 'i', 'n', 'f', '.', 'h', '\x00'})), int32(5), int32(0), float32(1), float32(1), float32(1), float32(0), int32(0)}, common.Struct_ff_f{(*int8)(unsafe.Pointer(&[56]int8{'/', 'U', 's', 'e', 'r', 's', '/', 'x', 'u', 's', 'h', 'i', 'w', 'e', 'i', '/', 'w', 'o', 'r', 'k', '/', 'l', 'i', 'b', 'c', '/', 't', 'e', 's', 't', '/', 's', 'r', 'c', '/', 'm', 'a', 't', 'h', '/', 's', 'p', 'e', 'c', 'i', 'a', 'l', '/', 'f', 'm', 'i', 'n', 'f', '.', 'h', '\x00'})), int32(6), int32(0), float32(-1), float32(1), float32(-1), float32(0), int32(0)}, common.Struct_ff_f{(*int8)(unsafe.Pointer(&[56]int8{'/', 'U', 's', 'e', 'r', 's', '/', 'x', 'u', 's', 'h', 'i', 'w', 'e', 'i', '/', 'w', 'o', 'r', 'k', '/', 'l', 'i', 'b', 'c', '/', 't', 'e', 's', 't', '/', 's', 'r', 'c', '/', 'm', 'a', 't', 'h', '/', 's', 'p', 'e', 'c', 'i', 'a', 'l', '/', 'f', 'm', 'i', 'n', 'f', '.', 'h', '\x00'})), int32(7), int32(0), libc.X__builtin_inff(), float32(1), float32(1), float32(0), int32(0)}, common.Struct_ff_f{(*int8)(unsafe.Pointer(&[56]int8{'/', 'U', 's', 'e', 'r', 's', '/', 'x', 'u', 's', 'h', 'i', 'w', 'e', 'i', '/', 'w', 'o', 'r', 'k', '/', 'l', 'i', 'b', 'c', '/', 't', 'e', 's', 't', '/', 's', 'r', 'c', '/', 'm', 'a', 't', 'h', '/', 's', 'p', 'e', 'c', 'i', 'a', 'l', '/', 'f', 'm', 'i', 'n', 'f', '.', 'h', '\x00'})), int32(8), int32(0), -libc.X__builtin_inff(), float32(1), -libc.X__builtin_inff(), float32(0), int32(0)}, common.Struct_ff_f{(*int8)(unsafe.Pointer(&[56]int8{'/', 'U', 's', 'e', 'r', 's', '/', 'x', 'u', 's', 'h', 'i', 'w', 'e', 'i', '/', 'w', 'o', 'r', 'k', '/', 'l', 'i', 'b', 'c', '/', 't', 'e', 's', 't', '/', 's', 'r', 'c', '/', 'm', 'a', 't', 'h', '/', 's', 'p', 'e', 'c', 'i', 'a', 'l', '/', 'f', 'm', 'i', 'n', 'f', '.', 'h', '\x00'})), int32(9), int32(0), libc.X__builtin_nanf((*int8)(unsafe.Pointer(&[1]int8{'\x00'}))), float32(1), float32(1), float32(0), int32(0)}, common.Struct_ff_f{(*int8)(unsafe.Pointer(&[56]int8{'/', 'U', 's', 'e', 'r', 's', '/', 'x', 'u', 's', 'h', 'i', 'w', 'e', 'i', '/', 'w', 'o', 'r', 'k', '/', 'l', 'i', 'b', 'c', '/', 't', 'e', 's', 't', '/', 's', 'r', 'c', '/', 'm', 'a', 't', 'h', '/', 's', 'p', 'e', 'c', 'i', 'a', 'l', '/', 'f', 'm', 'i', 'n', 'f', '.', 'h', '\x00'})), int32(10), int32(0), float32(0), float32(-1), float32(-1), float32(0), int32(0)}, common.Struct_ff_f{(*int8)(unsafe.Pointer(&[56]int8{'/', 'U', 's', 'e', 'r', 's', '/', 'x', 'u', 's', 'h', 'i', 'w', 'e', 'i', '/', 'w', 'o', 'r', 'k', '/', 'l', 'i', 'b', 'c', '/', 't', 'e', 's', 't', '/', 's', 'r', 'c', '/', 'm', 'a', 't', 'h', '/', 's', 'p', 'e', 'c', 'i', 'a', 'l', '/', 'f', 'm', 'i', 'n', 'f', '.', 'h', '\x00'})), int32(11), int32(0), float32(-0), float32(-1), float32(-1), float32(0), int32(0)}, common.Struct_ff_f{(*int8)(unsafe.Pointer(&[56]int8{'/', 'U', 's', 'e', 'r', 's', '/', 'x', 'u', 's', 'h', 'i', 'w', 'e', 'i', '/', 'w', 'o', 'r', 'k', '/', 'l', 'i', 'b', 'c', '/', 't', 'e', 's', 't', '/', 's', 'r', 'c', '/', 'm', 'a', 't', 'h', '/', 's', 'p', 'e', 'c', 'i', 'a', 'l', '/', 'f', 'm', 'i', 'n', 'f', '.', 'h', '\x00'})), int32(12), int32(0), float32(0.5), float32(-1), float32(-1), float32(0), int32(0)}, common.Struct_ff_f{(*int8)(unsafe.Pointer(&[56]int8{'/', 'U', 's', 'e', 'r', 's', '/', 'x', 'u', 's', 'h', 'i', 'w', 'e', 'i', '/', 'w', 'o', 'r', 'k', '/', 'l', 'i', 'b', 'c', '/', 't', 'e', 's', 't', '/', 's', 'r', 'c', '/', 'm', 'a', 't', 'h', '/', 's', 'p', 'e', 'c', 'i', 'a', 'l', '/', 'f', 'm', 'i', 'n', 'f', '.', 'h', '\x00'})), int32(13), int32(0), float32(-0.5), float32(-1), float32(-1), float32(0), int32(0)}, common.Struct_ff_f{(*int8)(unsafe.Pointer(&[56]int8{'/', 'U', 's', 'e', 'r', 's', '/', 'x', 'u', 's', 'h', 'i', 'w', 'e', 'i', '/', 'w', 'o', 'r', 'k', '/', 'l', 'i', 'b', 'c', '/', 't', 'e', 's', 't', '/', 's', 'r', 'c', '/', 'm', 'a', 't', 'h', '/', 's', 'p', 'e', 'c', 'i', 'a', 'l', '/', 'f', 'm', 'i', 'n', 'f', '.', 'h', '\x00'})), int32(14), int32(0), float32(1), float32(-1), float32(-1), float32(0), int32(0)}, common.Struct_ff_f{(*int8)(unsafe.Pointer(&[56]int8{'/', 'U', 's', 'e', 'r', 's', '/', 'x', 'u', 's', 'h', 'i', 'w', 'e', 'i', '/', 'w', 'o', 'r', 'k', '/', 'l', 'i', 'b', 'c', '/', 't', 'e', 's', 't', '/', 's', 'r', 'c', '/', 'm', 'a', 't', 'h', '/', 's', 'p', 'e', 'c', 'i', 'a', 'l', '/', 'f', 'm', 'i', 'n', 'f', '.', 'h', '\x00'})), int32(15), int32(0), float32(-1), float32(-1), float32(-1), float32(0), int32(0)}, common.Struct_ff_f{(*int8)(unsafe.Pointer(&[56]int8{'/', 'U', 's', 'e', 'r', 's', '/', 'x', 'u', 's', 'h', 'i', 'w', 'e', 'i', '/', 'w', 'o', 'r', 'k', '/', 'l', 'i', 'b', 'c', '/', 't', 'e', 's', 't', '/', 's', 'r', 'c', '/', 'm', 'a', 't', 'h', '/', 's', 'p', 'e', 'c', 'i', 'a', 'l', '/', 'f', 'm', 'i', 'n', 'f', '.', 'h', '\x00'})), int32(16), int32(0), libc.X__builtin_inff(), float32(-1), float32(-1), float32(0), int32(0)}, common.Struct_ff_f{(*int8)(unsafe.Pointer(&[56]int8{'/', 'U', 's', 'e', 'r', 's', '/', 'x', 'u', 's', 'h', 'i', 'w', 'e', 'i', '/', 'w', 'o', 'r', 'k', '/', 'l', 'i', 'b', 'c', '/', 't', 'e', 's', 't', '/', 's', 'r', 'c', '/', 'm', 'a', 't', 'h', '/', 's', 'p', 'e', 'c', 'i', 'a', 'l', '/', 'f', 'm', 'i', 'n', 'f', '.', 'h', '\x00'})), int32(17), int32(0), -libc.X__builtin_inff(), float32(-1), -libc.X__builtin_inff(), float32(0), int32(0)}, common.Struct_ff_f{(*int8)(unsafe.Pointer(&[56]int8{'/', 'U', 's', 'e', 'r', 's', '/', 'x', 'u', 's', 'h', 'i', 'w', 'e', 'i', '/', 'w', 'o', 'r', 'k', '/', 'l', 'i', 'b', 'c', '/', 't', 'e', 's', 't', '/', 's', 'r', 'c', '/', 'm', 'a', 't', 'h', '/', 's', 'p', 'e', 'c', 'i', 'a', 'l', '/', 'f', 'm', 'i', 'n', 'f', '.', 'h', '\x00'})), int32(18), int32(0), libc.X__builtin_nanf((*int8)(unsafe.Pointer(&[1]int8{'\x00'}))), float32(-1), float32(-1), float32(0), int32(0)}, common.Struct_ff_f{(*int8)(unsafe.Pointer(&[56]int8{'/', 'U', 's', 'e', 'r', 's', '/', 'x', 'u', 's', 'h', 'i', 'w', 'e', 'i', '/', 'w', 'o', 'r', 'k', '/', 'l', 'i', 'b', 'c', '/', 't', 'e', 's', 't', '/', 's', 'r', 'c', '/', 'm', 'a', 't', 'h', '/', 's', 'p', 'e', 'c', 'i', 'a', 'l', '/', 'f', 'm', 'i', 'n', 'f', '.', 'h', '\x00'})), int32(19), int32(0), float32(0), float32(0), float32(0), float32(0), int32(0)}, common.Struct_ff_f{(*int8)(unsafe.Pointer(&[56]int8{'/', 'U', 's', 'e', 'r', 's', '/', 'x', 'u', 's', 'h', 'i', 'w', 'e', 'i', '/', 'w', 'o', 'r', 'k', '/', 'l', 'i', 'b', 'c', '/', 't', 'e', 's', 't', '/', 's', 'r', 'c', '/', 'm', 'a', 't', 'h', '/', 's', 'p', 'e', 'c', 'i', 'a', 'l', '/', 'f', 'm', 'i', 'n', 'f', '.', 'h', '\x00'})), int32(20), int32(0), float32(0), float32(-0), float32(-0), float32(0), int32(0)}, common.Struct_ff_f{(*int8)(unsafe.Pointer(&[56]int8{'/', 'U', 's', 'e', 'r', 's', '/', 'x', 'u', 's', 'h', 'i', 'w', 'e', 'i', '/', 'w', 'o', 'r', 'k', '/', 'l', 'i', 'b', 'c', '/', 't', 'e', 's', 't', '/', 's', 'r', 'c', '/', 'm', 'a', 't', 'h', '/', 's', 'p', 'e', 'c', 'i', 'a', 'l', '/', 'f', 'm', 'i', 'n', 'f', '.', 'h', '\x00'})), int32(21), int32(0), float32(0), libc.X__builtin_inff(), float32(0), float32(0), int32(0)}, common.Struct_ff_f{(*int8)(unsafe.Pointer(&[56]int8{'/', 'U', 's', 'e', 'r', 's', '/', 'x', 'u', 's', 'h', 'i', 'w', 'e', 'i', '/', 'w', 'o', 'r', 'k', '/', 'l', 'i', 'b', 'c', '/', 't', 'e', 's', 't', '/', 's', 'r', 'c', '/', 'm', 'a', 't', 'h', '/', 's', 'p', 'e', 'c', 'i', 'a', 'l', '/', 'f', 'm', 'i', 'n', 'f', '.', 'h', '\x00'})), int32(22), int32(0), float32(0), -libc.X__builtin_inff(), -libc.X__builtin_inff(), float32(0), int32(0)}, common.Struct_ff_f{(*int8)(unsafe.Pointer(&[56]int8{'/', 'U', 's', 'e', 'r', 's', '/', 'x', 'u', 's', 'h', 'i', 'w', 'e', 'i', '/', 'w', 'o', 'r', 'k', '/', 'l', 'i', 'b', 'c', '/', 't', 'e', 's', 't', '/', 's', 'r', 'c', '/', 'm', 'a', 't', 'h', '/', 's', 'p', 'e', 'c', 'i', 'a', 'l', '/', 'f', 'm', 'i', 'n', 'f', '.', 'h', '\x00'})), int32(23), int32(0), float32(0), libc.X__builtin_nanf((*int8)(unsafe.Pointer(&[1]int8{'\x00'}))), float32(0), float32(0), int32(0)}, common.Struct_ff_f{(*int8)(unsafe.Pointer(&[56]int8{'/', 'U', 's', 'e', 'r', 's', '/', 'x', 'u', 's', 'h', 'i', 'w', 'e', 'i', '/', 'w', 'o', 'r', 'k', '/', 'l', 'i', 'b', 'c', '/', 't', 'e', 's', 't', '/', 's', 'r', 'c', '/', 'm', 'a', 't', 'h', '/', 's', 'p', 'e', 'c', 'i', 'a', 'l', '/', 'f', 'm', 'i', 'n', 'f', '.', 'h', '\x00'})), int32(24), int32(0), float32(-0), float32(0), float32(-0), float32(0), int32(0)}, common.Struct_ff_f{(*int8)(unsafe.Pointer(&[56]int8{'/', 'U', 's', 'e', 'r', 's', '/', 'x', 'u', 's', 'h', 'i', 'w', 'e', 'i', '/', 'w', 'o', 'r', 'k', '/', 'l', 'i', 'b', 'c', '/', 't', 'e', 's', 't', '/', 's', 'r', 'c', '/', 'm', 'a', 't', 'h', '/', 's', 'p', 'e', 'c', 'i', 'a', 'l', '/', 'f', 'm', 'i', 'n', 'f', '.', 'h', '\x00'})), int32(25), int32(0), float32(-0), float32(-0), float32(-0), float32(0), int32(0)}, common.Struct_ff_f{(*int8)(unsafe.Pointer(&[56]int8{'/', 'U', 's', 'e', 'r', 's', '/', 'x', 'u', 's', 'h', 'i', 'w', 'e', 'i', '/', 'w', 'o', 'r', 'k', '/', 'l', 'i', 'b', 'c', '/', 't', 'e', 's', 't', '/', 's', 'r', 'c', '/', 'm', 'a', 't', 'h', '/', 's', 'p', 'e', 'c', 'i', 'a', 'l', '/', 'f', 'm', 'i', 'n', 'f', '.', 'h', '\x00'})), int32(26), int32(0), float32(-0), libc.X__builtin_inff(), float32(-0), float32(0), int32(0)}, common.Struct_ff_f{(*int8)(unsafe.Pointer(&[56]int8{'/', 'U', 's', 'e', 'r', 's', '/', 'x', 'u', 's', 'h', 'i', 'w', 'e', 'i', '/', 'w', 'o', 'r', 'k', '/', 'l', 'i', 'b', 'c', '/', 't', 'e', 's', 't', '/', 's', 'r', 'c', '/', 'm', 'a', 't', 'h', '/', 's', 'p', 'e', 'c', 'i', 'a', 'l', '/', 'f', 'm', 'i', 'n', 'f', '.', 'h', '\x00'})), int32(27), int32(0), float32(-0), -libc.X__builtin_inff(), -libc.X__builtin_inff(), float32(0), int32(0)}, common.Struct_ff_f{(*int8)(unsafe.Pointer(&[56]int8{'/', 'U', 's', 'e', 'r', 's', '/', 'x', 'u', 's', 'h', 'i', 'w', 'e', 'i', '/', 'w', 'o', 'r', 'k', '/', 'l', 'i', 'b', 'c', '/', 't', 'e', 's', 't', '/', 's', 'r', 'c', '/', 'm', 'a', 't', 'h', '/', 's', 'p', 'e', 'c', 'i', 'a', 'l', '/', 'f', 'm', 'i', 'n', 'f', '.', 'h', '\x00'})), int32(28), int32(0), float32(-0), libc.X__builtin_nanf((*int8)(unsafe.Pointer(&[1]int8{'\x00'}))), float32(-0), float32(0), int32(0)}, common.Struct_ff_f{(*int8)(unsafe.Pointer(&[56]int8{'/', 'U', 's', 'e', 'r', 's', '/', 'x', 'u', 's', 'h', 'i', 'w', 'e', 'i', '/', 'w', 'o', 'r', 'k', '/', 'l', 'i', 'b', 'c', '/', 't', 'e', 's', 't', '/', 's', 'r', 'c', '/', 'm', 'a', 't', 'h', '/', 's', 'p', 'e', 'c', 'i', 'a', 'l', '/', 'f', 'm', 'i', 'n', 'f', '.', 'h', '\x00'})), int32(29), int32(0), float32(1), float32(0), float32(0), float32(0), int32(0)}, common.Struct_ff_f{(*int8)(unsafe.Pointer(&[56]int8{'/', 'U', 's', 'e', 'r', 's', '/', 'x', 'u', 's', 'h', 'i', 'w', 'e', 'i', '/', 'w', 'o', 'r', 'k', '/', 'l', 'i', 'b', 'c', '/', 't', 'e', 's', 't', '/', 's', 'r', 'c', '/', 'm', 'a', 't', 'h', '/', 's', 'p', 'e', 'c', 'i', 'a', 'l', '/', 'f', 'm', 'i', 'n', 'f', '.', 'h', '\x00'})), int32(30), int32(0), float32(-1), float32(0), float32(-1), float32(0), int32(0)}, common.Struct_ff_f{(*int8)(unsafe.Pointer(&[56]int8{'/', 'U', 's', 'e', 'r', 's', '/', 'x', 'u', 's', 'h', 'i', 'w', 'e', 'i', '/', 'w', 'o', 'r', 'k', '/', 'l', 'i', 'b', 'c', '/', 't', 'e', 's', 't', '/', 's', 'r', 'c', '/', 'm', 'a', 't', 'h', '/', 's', 'p', 'e', 'c', 'i', 'a', 'l', '/', 'f', 'm', 'i', 'n', 'f', '.', 'h', '\x00'})), int32(31), int32(0), libc.X__builtin_inff(), float32(0), float32(0), float32(0), int32(0)}, common.Struct_ff_f{(*int8)(unsafe.Pointer(&[56]int8{'/', 'U', 's', 'e', 'r', 's', '/', 'x', 'u', 's', 'h', 'i', 'w', 'e', 'i', '/', 'w', 'o', 'r', 'k', '/', 'l', 'i', 'b', 'c', '/', 't', 'e', 's', 't', '/', 's', 'r', 'c', '/', 'm', 'a', 't', 'h', '/', 's', 'p', 'e', 'c', 'i', 'a', 'l', '/', 'f', 'm', 'i', 'n', 'f', '.', 'h', '\x00'})), int32(32), int32(0), -libc.X__builtin_inff(), float32(0), -libc.X__builtin_inff(), float32(0), int32(0)}, common.Struct_ff_f{(*int8)(unsafe.Pointer(&[56]int8{'/', 'U', 's', 'e', 'r', 's', '/', 'x', 'u', 's', 'h', 'i', 'w', 'e', 'i', '/', 'w', 'o', 'r', 'k', '/', 'l', 'i', 'b', 'c', '/', 't', 'e', 's', 't', '/', 's', 'r', 'c', '/', 'm', 'a', 't', 'h', '/', 's', 'p', 'e', 'c', 'i', 'a', 'l', '/', 'f', 'm', 'i', 'n', 'f', '.', 'h', '\x00'})), int32(33), int32(0), libc.X__builtin_nanf((*int8)(unsafe.Pointer(&[1]int8{'\x00'}))), float32(0), float32(0), float32(0), int32(0)}, common.Struct_ff_f{(*int8)(unsafe.Pointer(&[56]int8{'/', 'U', 's', 'e', 'r', 's', '/', 'x', 'u', 's', 'h', 'i', 'w', 'e', 'i', '/', 'w', 'o', 'r', 'k', '/', 'l', 'i', 'b', 'c', '/', 't', 'e', 's', 't', '/', 's', 'r', 'c', '/', 'm', 'a', 't', 'h', '/', 's', 'p', 'e', 'c', 'i', 'a', 'l', '/', 'f', 'm', 'i', 'n', 'f', '.', 'h', '\x00'})), int32(34), int32(0), float32(-1), float32(-0), float32(-1), float32(0), int32(0)}, common.Struct_ff_f{(*int8)(unsafe.Pointer(&[56]int8{'/', 'U', 's', 'e', 'r', 's', '/', 'x', 'u', 's', 'h', 'i', 'w', 'e', 'i', '/', 'w', 'o', 'r', 'k', '/', 'l', 'i', 'b', 'c', '/', 't', 'e', 's', 't', '/', 's', 'r', 'c', '/', 'm', 'a', 't', 'h', '/', 's', 'p', 'e', 'c', 'i', 'a', 'l', '/', 'f', 'm', 'i', 'n', 'f', '.', 'h', '\x00'})), int32(35), int32(0), libc.X__builtin_inff(), float32(-0), float32(-0), float32(0), int32(0)}, common.Struct_ff_f{(*int8)(unsafe.Pointer(&[56]int8{'/', 'U', 's', 'e', 'r', 's', '/', 'x', 'u', 's', 'h', 'i', 'w', 'e', 'i', '/', 'w', 'o', 'r', 'k', '/', 'l', 'i', 'b', 'c', '/', 't', 'e', 's', 't', '/', 's', 'r', 'c', '/', 'm', 'a', 't', 'h', '/', 's', 'p', 'e', 'c', 'i', 'a', 'l', '/', 'f', 'm', 'i', 'n', 'f', '.', 'h', '\x00'})), int32(36), int32(0), -libc.X__builtin_inff(), float32(-0), -libc.X__builtin_inff(), float32(0), int32(0)}, common.Struct_ff_f{(*int8)(unsafe.Pointer(&[56]int8{'/', 'U', 's', 'e', 'r', 's', '/', 'x', 'u', 's', 'h', 'i', 'w', 'e', 'i', '/', 'w', 'o', 'r', 'k', '/', 'l', 'i', 'b', 'c', '/', 't', 'e', 's', 't', '/', 's', 'r', 'c', '/', 'm', 'a', 't', 'h', '/', 's', 'p', 'e', 'c', 'i', 'a', 'l', '/', 'f', 'm', 'i', 'n', 'f', '.', 'h', '\x00'})), int32(37), int32(0), libc.X__builtin_nanf((*int8)(unsafe.Pointer(&[1]int8{'\x00'}))), float32(-0), float32(-0), float32(0), int32(0)}, common.Struct_ff_f{(*int8)(unsafe.Pointer(&[56]int8{'/', 'U', 's', 'e', 'r', 's', '/', 'x', 'u', 's', 'h', 'i', 'w', 'e', 'i', '/', 'w', 'o', 'r', 'k', '/', 'l', 'i', 'b', 'c', '/', 't', 'e', 's', 't', '/', 's', 'r', 'c', '/', 'm', 'a', 't', 'h', '/', 's', 'p', 'e', 'c', 'i', 'a', 'l', '/', 'f', 'm', 'i', 'n', 'f', '.', 'h', '\x00'})), int32(38), int32(0), libc.X__builtin_inff(), float32(2), float32(2), float32(0), int32(0)}, common.Struct_ff_f{(*int8)(unsafe.Pointer(&[56]int8{'/', 'U', 's', 'e', 'r', 's', '/', 'x', 'u', 's', 'h', 'i', 'w', 'e', 'i', '/', 'w', 'o', 'r', 'k', '/', 'l', 'i', 'b', 'c', '/', 't', 'e', 's', 't', '/', 's', 'r', 'c', '/', 'm', 'a', 't', 'h', '/', 's', 'p', 'e', 'c', 'i', 'a', 'l', '/', 'f', 'm', 'i', 'n', 'f', '.', 'h', '\x00'})), int32(39), int32(0), libc.X__builtin_inff(), float32(-0.5), float32(-0.5), float32(0), int32(0)}, common.Struct_ff_f{(*int8)(unsafe.Pointer(&[56]int8{'/', 'U', 's', 'e', 'r', 's', '/', 'x', 'u', 's', 'h', 'i', 'w', 'e', 'i', '/', 'w', 'o', 'r', 'k', '/', 'l', 'i', 'b', 'c', '/', 't', 'e', 's', 't', '/', 's', 'r', 'c', '/', 'm', 'a', 't', 'h', '/', 's', 'p', 'e', 'c', 'i', 'a', 'l', '/', 'f', 'm', 'i', 'n', 'f', '.', 'h', '\x00'})), int32(40), int32(0), libc.X__builtin_inff(), libc.X__builtin_nanf((*int8)(unsafe.Pointer(&[1]int8{'\x00'}))), libc.X__builtin_inff(), float32(0), int32(0)}, common.Struct_ff_f{(*int8)(unsafe.Pointer(&[56]int8{'/', 'U', 's', 'e', 'r', 's', '/', 'x', 'u', 's', 'h', 'i', 'w', 'e', 'i', '/', 'w', 'o', 'r', 'k', '/', 'l', 'i', 'b', 'c', '/', 't', 'e', 's', 't', '/', 's', 'r', 'c', '/', 'm', 'a', 't', 'h', '/', 's', 'p', 'e', 'c', 'i', 'a', 'l', '/', 'f', 'm', 'i', 'n', 'f', '.', 'h', '\x00'})), int32(41), int32(0), -libc.X__builtin_inff(), float32(2), -libc.X__builtin_inff(), float32(0), int32(0)}, common.Struct_ff_f{(*int8)(unsafe.Pointer(&[56]int8{'/', 'U', 's', 'e', 'r', 's', '/', 'x', 'u', 's', 'h', 'i', 'w', 'e', 'i', '/', 'w', 'o', 'r', 'k', '/', 'l', 'i', 'b', 'c', '/', 't', 'e', 's', 't', '/', 's', 'r', 'c', '/', 'm', 'a', 't', 'h', '/', 's', 'p', 'e', 'c', 'i', 'a', 'l', '/', 'f', 'm', 'i', 'n', 'f', '.', 'h', '\x00'})), int32(42), int32(0), -libc.X__builtin_inff(), float32(-0.5), -libc.X__builtin_inff(), float32(0), int32(0)}, common.Struct_ff_f{(*int8)(unsafe.Pointer(&[56]int8{'/', 'U', 's', 'e', 'r', 's', '/', 'x', 'u', 's', 'h', 'i', 'w', 'e', 'i', '/', 'w', 'o', 'r', 'k', '/', 'l', 'i', 'b', 'c', '/', 't', 'e', 's', 't', '/', 's', 'r', 'c', '/', 'm', 'a', 't', 'h', '/', 's', 'p', 'e', 'c', 'i', 'a', 'l', '/', 'f', 'm', 'i', 'n', 'f', '.', 'h', '\x00'})), int32(43), int32(0), -libc.X__builtin_inff(), libc.X__builtin_nanf((*int8)(unsafe.Pointer(&[1]int8{'\x00'}))), -libc.X__builtin_inff(), float32(0), int32(0)}, common.Struct_ff_f{(*int8)(unsafe.Pointer(&[56]int8{'/', 'U', 's', 'e', 'r', 's', '/', 'x', 'u', 's', 'h', 'i', 'w', 'e', 'i', '/', 'w', 'o', 'r', 'k', '/', 'l', 'i', 'b', 'c', '/', 't', 'e', 's', 't', '/', 's', 'r', 'c', '/', 'm', 'a', 't', 'h', '/', 's', 'p', 'e', 'c', 'i', 'a', 'l', '/', 'f', 'm', 'i', 'n', 'f', '.', 'h', '\x00'})), int32(44), int32(0), libc.X__builtin_nanf((*int8)(unsafe.Pointer(&[1]int8{'\x00'}))), libc.X__builtin_nanf((*int8)(unsafe.Pointer(&[1]int8{'\x00'}))), libc.X__builtin_nanf((*int8)(unsafe.Pointer(&[1]int8{'\x00'}))), float32(0), int32(0)}, common.Struct_ff_f{(*int8)(unsafe.Pointer(&[56]int8{'/', 'U', 's', 'e', 'r', 's', '/', 'x', 'u', 's', 'h', 'i', 'w', 'e', 'i', '/', 'w', 'o', 'r', 'k', '/', 'l', 'i', 'b', 'c', '/', 't', 'e', 's', 't', '/', 's', 'r', 'c', '/', 'm', 'a', 't', 'h', '/', 's', 'p', 'e', 'c', 'i', 'a', 'l', '/', 'f', 'm', 'i', 'n', 'f', '.', 'h', '\x00'})), int32(45), int32(0), float32(1), libc.X__builtin_nanf((*int8)(unsafe.Pointer(&[1]int8{'\x00'}))), float32(1), float32(0), int32(0)}, common.Struct_ff_f{(*int8)(unsafe.Pointer(&[56]int8{'/', 'U', 's', 'e', 'r', 's', '/', 'x', 'u', 's', 'h', 'i', 'w', 'e', 'i', '/', 'w', 'o', 'r', 'k', '/', 'l', 'i', 'b', 'c', '/', 't', 'e', 's', 't', '/', 's', 'r', 'c', '/', 'm', 'a', 't', 'h', '/', 's', 'p', 'e', 'c', 'i', 'a', 'l', '/', 'f', 'm', 'i', 'n', 'f', '.', 'h', '\x00'})), int32(46), int32(0), float32(-1), libc.X__builtin_nanf((*int8)(unsafe.Pointer(&[1]int8{'\x00'}))), float32(-1), float32(0), int32(0)}, common.Struct_ff_f{(*int8)(unsafe.Pointer(&[56]int8{'/', 'U', 's', 'e', 'r', 's', '/', 'x', 'u', 's', 'h', 'i', 'w', 'e', 'i', '/', 'w', 'o', 'r', 'k', '/', 'l', 'i', 'b', 'c', '/', 't', 'e', 's', 't', '/', 's', 'r', 'c', '/', 'm', 'a', 't', 'h', '/', 's', 'p', 'e', 'c', 'i', 'a', 'l', '/', 'f', 'm', 'i', 'n', 'f', '.', 'h', '\x00'})), int32(47), int32(0), float32(1), libc.X__builtin_inff(), float32(1), float32(0), int32(0)}, common.Struct_ff_f{(*int8)(unsafe.Pointer(&[56]int8{'/', 'U', 's', 'e', 'r', 's', '/', 'x', 'u', 's', 'h', 'i', 'w', 'e', 'i', '/', 'w', 'o', 'r', 'k', '/', 'l', 'i', 'b', 'c', '/', 't', 'e', 's', 't', '/', 's', 'r', 'c', '/', 'm', 'a', 't', 'h', '/', 's', 'p', 'e', 'c', 'i', 'a', 'l', '/', 'f', 'm', 'i', 'n', 'f', '.', 'h', '\x00'})), int32(48), int32(0), float32(-1), libc.X__builtin_inff(), float32(-1), float32(0), int32(0)}, common.Struct_ff_f{(*int8)(unsafe.Pointer(&[56]int8{'/', 'U', 's', 'e', 'r', 's', '/', 'x', 'u', 's', 'h', 'i', 'w', 'e', 'i', '/', 'w', 'o', 'r', 'k', '/', 'l', 'i', 'b', 'c', '/', 't', 'e', 's', 't', '/', 's', 'r', 'c', '/', 'm', 'a', 't', 'h', '/', 's', 'p', 'e', 'c', 'i', 'a', 'l', '/', 'f', 'm', 'i', 'n', 'f', '.', 'h', '\x00'})), int32(49), int32(0), libc.X__builtin_inff(), libc.X__builtin_inff(), libc.X__builtin_inff(), float32(0), int32(0)}, common.Struct_ff_f{(*int8)(unsafe.Pointer(&[56]int8{'/', 'U', 's', 'e', 'r', 's', '/', 'x', 'u', 's', 'h', 'i', 'w', 'e', 'i', '/', 'w', 'o', 'r', 'k', '/', 'l', 'i', 'b', 'c', '/', 't', 'e', 's', 't', '/', 's', 'r', 'c', '/', 'm', 'a', 't', 'h', '/', 's', 'p', 'e', 'c', 'i', 'a', 'l', '/', 'f', 'm', 'i', 'n', 'f', '.', 'h', '\x00'})), int32(50), int32(0), -libc.X__builtin_inff(), libc.X__builtin_inff(), -libc.X__builtin_inff(), float32(0), int32(0)}, common.Struct_ff_f{(*int8)(unsafe.Pointer(&[56]int8{'/', 'U', 's', 'e', 'r', 's', '/', 'x', 'u', 's', 'h', 'i', 'w', 'e', 'i', '/', 'w', 'o', 'r', 'k', '/', 'l', 'i', 'b', 'c', '/', 't', 'e', 's', 't', '/', 's', 'r', 'c', '/', 'm', 'a', 't', 'h', '/', 's', 'p', 'e', 'c', 'i', 'a', 'l', '/', 'f', 'm', 'i', 'n', 'f', '.', 'h', '\x00'})), int32(51), int32(0), float32(1), -libc.X__builtin_inff(), -libc.X__builtin_inff(), float32(0), int32(0)}, common.Struct_ff_f{(*int8)(unsafe.Pointer(&[56]int8{'/', 'U', 's', 'e', 'r', 's', '/', 'x', 'u', 's', 'h', 'i', 'w', 'e', 'i', '/', 'w', 'o', 'r', 'k', '/', 'l', 'i', 'b', 'c', '/', 't', 'e', 's', 't', '/', 's', 'r', 'c', '/', 'm', 'a', 't', 'h', '/', 's', 'p', 'e', 'c', 'i', 'a', 'l', '/', 'f', 'm', 'i', 'n', 'f', '.', 'h', '\x00'})), int32(52), int32(0), float32(-1), -libc.X__builtin_inff(), -libc.X__builtin_inff(), float32(0), int32(0)}, common.Struct_ff_f{(*int8)(unsafe.Pointer(&[56]int8{'/', 'U', 's', 'e', 'r', 's', '/', 'x', 'u', 's', 'h', 'i', 'w', 'e', 'i', '/', 'w', 'o', 'r', 'k', '/', 'l', 'i', 'b', 'c', '/', 't', 'e', 's', 't', '/', 's', 'r', 'c', '/', 'm', 'a', 't', 'h', '/', 's', 'p', 'e', 'c', 'i', 'a', 'l', '/', 'f', 'm', 'i', 'n', 'f', '.', 'h', '\x00'})), int32(53), int32(0), libc.X__builtin_inff(), -libc.X__builtin_inff(), -libc.X__builtin_inff(), float32(0), int32(0)}, common.Struct_ff_f{(*int8)(unsafe.Pointer(&[56]int8{'/', 'U', 's', 'e', 'r', 's', '/', 'x', 'u', 's', 'h', 'i', 'w', 'e', 'i', '/', 'w', 'o', 'r', 'k', '/', 'l', 'i', 'b', 'c', '/', 't', 'e', 's', 't', '/', 's', 'r', 'c', '/', 'm', 'a', 't', 'h', '/', 's', 'p', 'e', 'c', 'i', 'a', 'l', '/', 'f', 'm', 'i', 'n', 'f', '.', 'h', '\x00'})), int32(54), int32(0), -libc.X__builtin_inff(), -libc.X__builtin_inff(), -libc.X__builtin_inff(), float32(0), int32(0)}, common.Struct_ff_f{(*int8)(unsafe.Pointer(&[56]int8{'/', 'U', 's', 'e', 'r', 's', '/', 'x', 'u', 's', 'h', 'i', 'w', 'e', 'i', '/', 'w', 'o', 'r', 'k', '/', 'l', 'i', 'b', 'c', '/', 't', 'e', 's', 't', '/', 's', 'r', 'c', '/', 'm', 'a', 't', 'h', '/', 's', 'p', 'e', 'c', 'i', 'a', 'l', '/', 'f', 'm', 'i', 'n', 'f', '.', 'h', '\x00'})), int32(55), int32(0), float32(1.75), float32(0.5), float32(0.5), float32(0), int32(0)}, common.Struct_ff_f{(*int8)(unsafe.Pointer(&[56]int8{'/', 'U', 's', 'e', 'r', 's', '/', 'x', 'u', 's', 'h', 'i', 'w', 'e', 'i', '/', 'w', 'o', 'r', 'k', '/', 'l', 'i', 'b', 'c', '/', 't', 'e', 's', 't', '/', 's', 'r', 'c', '/', 'm', 'a', 't', 'h', '/', 's', 'p', 'e', 'c', 'i', 'a', 'l', '/', 'f', 'm', 'i', 'n', 'f', '.', 'h', '\x00'})), int32(56), int32(0), float32(-1.75), float32(0.5), float32(-1.75), float32(0), int32(0)}, common.Struct_ff_f{(*int8)(unsafe.Pointer(&[56]int8{'/', 'U', 's', 'e', 'r', 's', '/', 'x', 'u', 's', 'h', 'i', 'w', 'e', 'i', '/', 'w', 'o', 'r', 'k', '/', 'l', 'i', 'b', 'c', '/', 't', 'e', 's', 't', '/', 's', 'r', 'c', '/', 'm', 'a', 't', 'h', '/', 's', 'p', 'e', 'c', 'i', 'a', 'l', '/', 'f', 'm', 'i', 'n', 'f', '.', 'h', '\x00'})), int32(57), int32(0), float32(1.75), float32(-0.5), float32(-0.5), float32(0), int32(0)}, common.Struct_ff_f{(*int8)(unsafe.Pointer(&[56]int8{'/', 'U', 's', 'e', 'r', 's', '/', 'x', 'u', 's', 'h', 'i', 'w', 'e', 'i', '/', 'w', 'o', 'r', 'k', '/', 'l', 'i', 'b', 'c', '/', 't', 'e', 's', 't', '/', 's', 'r', 'c', '/', 'm', 'a', 't', 'h', '/', 's', 'p', 'e', 'c', 'i', 'a', 'l', '/', 'f', 'm', 'i', 'n', 'f', '.', 'h', '\x00'})), int32(58), int32(0), float32(-1.75), float32(-0.5), float32(-1.75), float32(0), int32(0)}}

func _cgo_main() int32 {
	var y float32
	var d float32
	var e int32
	var i int32
	var err int32 = int32(0)
	var p *common.Struct_ff_f
	for i = int32(0); uint64(i) < 68; i++ {
		p = (*common.Struct_ff_f)(unsafe.Pointer(uintptr(unsafe.Pointer((*common.Struct_ff_f)(unsafe.Pointer(&_cgos_t_fminf)))) + uintptr(i)*40))
		if p.R < int32(0) {
			continue
		}
		libc.Fesetround(p.R)
		libc.Feclearexcept(int32(0))
		y = libc.Fminf(p.X, p.X2)
		e = libc.Fetestexcept(0)
		if !(common.Checkexceptall(e, p.E, p.R) != 0) {
			libc.Printf((*int8)(unsafe.Pointer(&[53]int8{'%', 's', ':', '%', 'd', ':', ' ', 'b', 'a', 'd', ' ', 'f', 'p', ' ', 'e', 'x', 'c', 'e', 'p', 't', 'i', 'o', 'n', ':', ' ', '%', 's', ' ', 'f', 'm', 'i', 'n', 'f', '(', '%', 'a', ',', '%', 'a', ')', '=', '%', 'a', ',', ' ', 'w', 'a', 'n', 't', ' ', '%', 's', '\x00'})), p.File, p.Line, common.Rstr(p.R), float64(p.X), float64(p.X2), float64(p.Y), common.Estr(p.E))
			libc.Printf((*int8)(unsafe.Pointer(&[9]int8{' ', 'g', 'o', 't', ' ', '%', 's', '\n', '\x00'})), common.Estr(e))
			err++
		}
		d = common.Ulperrf(y, p.Y, p.Dy)
		if !(common.Checkcr(float64(y), float64(p.Y), p.R) != 0) {
			libc.Printf((*int8)(unsafe.Pointer(&[61]int8{'%', 's', ':', '%', 'd', ':', ' ', '%', 's', ' ', 'f', 'm', 'i', 'n', 'f', '(', '%', 'a', ',', '%', 'a', ')', ' ', 'w', 'a', 'n', 't', ' ', '%', 'a', ' ', 'g', 'o', 't', ' ', '%', 'a', ' ', 'u', 'l', 'p', 'e', 'r', 'r', ' ', '%', '.', '3', 'f', ' ', '=', ' ', '%', 'a', ' ', '+', ' ', '%', 'a', '\n', '\x00'})), p.File, p.Line, common.Rstr(p.R), float64(p.X), float64(p.X2), float64(p.Y), float64(y), float64(d), float64(d-p.Dy), float64(p.Dy))
			err++
		}
	}
	return func() int32 {
		if !!(err != 0) {
			return 1
		} else {
			return 0
		}
	}()
}
func TestMain(t *testing.T) {
	if _cgo_ret := _cgo_main(); _cgo_ret != 0 {
		t.Fatal("exit status", _cgo_ret)
	}
}
