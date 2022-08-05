package fabsl

import (
	common "github.com/goplus/libc/test/common"
	unsafe "unsafe"
	libc "github.com/goplus/libc"
	testing "testing"
)

var _cgos_t_fabsl [85]common.Struct_l_l = [85]common.Struct_l_l{common.Struct_l_l{(*int8)(unsafe.Pointer(&[51]int8{'/', 'U', 's', 'e', 'r', 's', '/', 'x', 'u', 's', 'h', 'i', 'w', 'e', 'i', '/', 'w', 'o', 'r', 'k', '/', 'l', 'i', 'b', 'c', '/', 't', 'e', 's', 't', '/', 's', 'r', 'c', '/', 'm', 'a', 't', 'h', '/', 'u', 'c', 'b', '/', 'f', 'a', 'b', 's', '.', 'h', '\x00'})), int32(37), int32(0), float64(0), float64(0), float32(0), int32(0)}, common.Struct_l_l{(*int8)(unsafe.Pointer(&[51]int8{'/', 'U', 's', 'e', 'r', 's', '/', 'x', 'u', 's', 'h', 'i', 'w', 'e', 'i', '/', 'w', 'o', 'r', 'k', '/', 'l', 'i', 'b', 'c', '/', 't', 'e', 's', 't', '/', 's', 'r', 'c', '/', 'm', 'a', 't', 'h', '/', 'u', 'c', 'b', '/', 'f', 'a', 'b', 's', '.', 'h', '\x00'})), int32(38), int32(0), float64(4.9406564584124654e-324), float64(4.9406564584124654e-324), float32(0), int32(0)}, common.Struct_l_l{(*int8)(unsafe.Pointer(&[51]int8{'/', 'U', 's', 'e', 'r', 's', '/', 'x', 'u', 's', 'h', 'i', 'w', 'e', 'i', '/', 'w', 'o', 'r', 'k', '/', 'l', 'i', 'b', 'c', '/', 't', 'e', 's', 't', '/', 's', 'r', 'c', '/', 'm', 'a', 't', 'h', '/', 'u', 'c', 'b', '/', 'f', 'a', 'b', 's', '.', 'h', '\x00'})), int32(39), int32(0), float64(2.1219957909652723e-314), float64(2.1219957909652723e-314), float32(0), int32(0)}, common.Struct_l_l{(*int8)(unsafe.Pointer(&[51]int8{'/', 'U', 's', 'e', 'r', 's', '/', 'x', 'u', 's', 'h', 'i', 'w', 'e', 'i', '/', 'w', 'o', 'r', 'k', '/', 'l', 'i', 'b', 'c', '/', 't', 'e', 's', 't', '/', 's', 'r', 'c', '/', 'm', 'a', 't', 'h', '/', 'u', 'c', 'b', '/', 'f', 'a', 'b', 's', '.', 'h', '\x00'})), int32(40), int32(0), float64(2.2250738585072014e-308), float64(2.2250738585072014e-308), float32(0), int32(0)}, common.Struct_l_l{(*int8)(unsafe.Pointer(&[51]int8{'/', 'U', 's', 'e', 'r', 's', '/', 'x', 'u', 's', 'h', 'i', 'w', 'e', 'i', '/', 'w', 'o', 'r', 'k', '/', 'l', 'i', 'b', 'c', '/', 't', 'e', 's', 't', '/', 's', 'r', 'c', '/', 'm', 'a', 't', 'h', '/', 'u', 'c', 'b', '/', 'f', 'a', 'b', 's', '.', 'h', '\x00'})), int32(41), int32(0), float64(1), float64(1), float32(0), int32(0)}, common.Struct_l_l{(*int8)(unsafe.Pointer(&[51]int8{'/', 'U', 's', 'e', 'r', 's', '/', 'x', 'u', 's', 'h', 'i', 'w', 'e', 'i', '/', 'w', 'o', 'r', 'k', '/', 'l', 'i', 'b', 'c', '/', 't', 'e', 's', 't', '/', 's', 'r', 'c', '/', 'm', 'a', 't', 'h', '/', 'u', 'c', 'b', '/', 'f', 'a', 'b', 's', '.', 'h', '\x00'})), int32(42), int32(0), float64(1.7976931348623157e+308), float64(1.7976931348623157e+308), float32(0), int32(0)}, common.Struct_l_l{(*int8)(unsafe.Pointer(&[51]int8{'/', 'U', 's', 'e', 'r', 's', '/', 'x', 'u', 's', 'h', 'i', 'w', 'e', 'i', '/', 'w', 'o', 'r', 'k', '/', 'l', 'i', 'b', 'c', '/', 't', 'e', 's', 't', '/', 's', 'r', 'c', '/', 'm', 'a', 't', 'h', '/', 'u', 'c', 'b', '/', 'f', 'a', 'b', 's', '.', 'h', '\x00'})), int32(43), int32(0), float64(libc.X__builtin_inff()), float64(libc.X__builtin_inff()), float32(0), int32(0)}, common.Struct_l_l{(*int8)(unsafe.Pointer(&[51]int8{'/', 'U', 's', 'e', 'r', 's', '/', 'x', 'u', 's', 'h', 'i', 'w', 'e', 'i', '/', 'w', 'o', 'r', 'k', '/', 'l', 'i', 'b', 'c', '/', 't', 'e', 's', 't', '/', 's', 'r', 'c', '/', 'm', 'a', 't', 'h', '/', 'u', 'c', 'b', '/', 'f', 'a', 'b', 's', '.', 'h', '\x00'})), int32(44), int32(0), float64(libc.X__builtin_nanf((*int8)(unsafe.Pointer(&[1]int8{'\x00'})))), float64(libc.X__builtin_nanf((*int8)(unsafe.Pointer(&[1]int8{'\x00'})))), float32(0), int32(0)}, common.Struct_l_l{(*int8)(unsafe.Pointer(&[51]int8{'/', 'U', 's', 'e', 'r', 's', '/', 'x', 'u', 's', 'h', 'i', 'w', 'e', 'i', '/', 'w', 'o', 'r', 'k', '/', 'l', 'i', 'b', 'c', '/', 't', 'e', 's', 't', '/', 's', 'r', 'c', '/', 'm', 'a', 't', 'h', '/', 'u', 'c', 'b', '/', 'f', 'a', 'b', 's', '.', 'h', '\x00'})), int32(45), int32(0), float64(-0), float64(0), float32(0), int32(0)}, common.Struct_l_l{(*int8)(unsafe.Pointer(&[51]int8{'/', 'U', 's', 'e', 'r', 's', '/', 'x', 'u', 's', 'h', 'i', 'w', 'e', 'i', '/', 'w', 'o', 'r', 'k', '/', 'l', 'i', 'b', 'c', '/', 't', 'e', 's', 't', '/', 's', 'r', 'c', '/', 'm', 'a', 't', 'h', '/', 'u', 'c', 'b', '/', 'f', 'a', 'b', 's', '.', 'h', '\x00'})), int32(46), int32(0), float64(-4.9406564584124654e-324), float64(4.9406564584124654e-324), float32(0), int32(0)}, common.Struct_l_l{(*int8)(unsafe.Pointer(&[51]int8{'/', 'U', 's', 'e', 'r', 's', '/', 'x', 'u', 's', 'h', 'i', 'w', 'e', 'i', '/', 'w', 'o', 'r', 'k', '/', 'l', 'i', 'b', 'c', '/', 't', 'e', 's', 't', '/', 's', 'r', 'c', '/', 'm', 'a', 't', 'h', '/', 'u', 'c', 'b', '/', 'f', 'a', 'b', 's', '.', 'h', '\x00'})), int32(47), int32(0), float64(-2.1219957909652723e-314), float64(2.1219957909652723e-314), float32(0), int32(0)}, common.Struct_l_l{(*int8)(unsafe.Pointer(&[51]int8{'/', 'U', 's', 'e', 'r', 's', '/', 'x', 'u', 's', 'h', 'i', 'w', 'e', 'i', '/', 'w', 'o', 'r', 'k', '/', 'l', 'i', 'b', 'c', '/', 't', 'e', 's', 't', '/', 's', 'r', 'c', '/', 'm', 'a', 't', 'h', '/', 'u', 'c', 'b', '/', 'f', 'a', 'b', 's', '.', 'h', '\x00'})), int32(48), int32(0), float64(-2.2250738585072014e-308), float64(2.2250738585072014e-308), float32(0), int32(0)}, common.Struct_l_l{(*int8)(unsafe.Pointer(&[51]int8{'/', 'U', 's', 'e', 'r', 's', '/', 'x', 'u', 's', 'h', 'i', 'w', 'e', 'i', '/', 'w', 'o', 'r', 'k', '/', 'l', 'i', 'b', 'c', '/', 't', 'e', 's', 't', '/', 's', 'r', 'c', '/', 'm', 'a', 't', 'h', '/', 'u', 'c', 'b', '/', 'f', 'a', 'b', 's', '.', 'h', '\x00'})), int32(49), int32(0), float64(-1.0000000000000033), float64(1.0000000000000033), float32(0), int32(0)}, common.Struct_l_l{(*int8)(unsafe.Pointer(&[51]int8{'/', 'U', 's', 'e', 'r', 's', '/', 'x', 'u', 's', 'h', 'i', 'w', 'e', 'i', '/', 'w', 'o', 'r', 'k', '/', 'l', 'i', 'b', 'c', '/', 't', 'e', 's', 't', '/', 's', 'r', 'c', '/', 'm', 'a', 't', 'h', '/', 'u', 'c', 'b', '/', 'f', 'a', 'b', 's', '.', 'h', '\x00'})), int32(50), int32(0), float64(-1.7976931348623157e+308), float64(1.7976931348623157e+308), float32(0), int32(0)}, common.Struct_l_l{(*int8)(unsafe.Pointer(&[51]int8{'/', 'U', 's', 'e', 'r', 's', '/', 'x', 'u', 's', 'h', 'i', 'w', 'e', 'i', '/', 'w', 'o', 'r', 'k', '/', 'l', 'i', 'b', 'c', '/', 't', 'e', 's', 't', '/', 's', 'r', 'c', '/', 'm', 'a', 't', 'h', '/', 'u', 'c', 'b', '/', 'f', 'a', 'b', 's', '.', 'h', '\x00'})), int32(51), int32(0), float64(-libc.X__builtin_inff()), float64(libc.X__builtin_inff()), float32(0), int32(0)}, common.Struct_l_l{(*int8)(unsafe.Pointer(&[51]int8{'/', 'U', 's', 'e', 'r', 's', '/', 'x', 'u', 's', 'h', 'i', 'w', 'e', 'i', '/', 'w', 'o', 'r', 'k', '/', 'l', 'i', 'b', 'c', '/', 't', 'e', 's', 't', '/', 's', 'r', 'c', '/', 'm', 'a', 't', 'h', '/', 'u', 'c', 'b', '/', 'f', 'a', 'b', 's', '.', 'h', '\x00'})), int32(52), int32(0), float64(libc.X__builtin_nanf((*int8)(unsafe.Pointer(&[1]int8{'\x00'})))), float64(libc.X__builtin_nanf((*int8)(unsafe.Pointer(&[1]int8{'\x00'})))), float32(0), int32(0)}, common.Struct_l_l{(*int8)(unsafe.Pointer(&[51]int8{'/', 'U', 's', 'e', 'r', 's', '/', 'x', 'u', 's', 'h', 'i', 'w', 'e', 'i', '/', 'w', 'o', 'r', 'k', '/', 'l', 'i', 'b', 'c', '/', 't', 'e', 's', 't', '/', 's', 'r', 'c', '/', 'm', 'a', 't', 'h', '/', 'u', 'c', 'b', '/', 'f', 'a', 'b', 's', '.', 'h', '\x00'})), int32(53), -1, float64(0), float64(0), float32(0), int32(0)}, common.Struct_l_l{(*int8)(unsafe.Pointer(&[51]int8{'/', 'U', 's', 'e', 'r', 's', '/', 'x', 'u', 's', 'h', 'i', 'w', 'e', 'i', '/', 'w', 'o', 'r', 'k', '/', 'l', 'i', 'b', 'c', '/', 't', 'e', 's', 't', '/', 's', 'r', 'c', '/', 'm', 'a', 't', 'h', '/', 'u', 'c', 'b', '/', 'f', 'a', 'b', 's', '.', 'h', '\x00'})), int32(54), -1, float64(4.9406564584124654e-324), float64(4.9406564584124654e-324), float32(0), int32(0)}, common.Struct_l_l{(*int8)(unsafe.Pointer(&[51]int8{'/', 'U', 's', 'e', 'r', 's', '/', 'x', 'u', 's', 'h', 'i', 'w', 'e', 'i', '/', 'w', 'o', 'r', 'k', '/', 'l', 'i', 'b', 'c', '/', 't', 'e', 's', 't', '/', 's', 'r', 'c', '/', 'm', 'a', 't', 'h', '/', 'u', 'c', 'b', '/', 'f', 'a', 'b', 's', '.', 'h', '\x00'})), int32(55), -1, float64(2.2250738585072009e-308), float64(2.2250738585072009e-308), float32(0), int32(0)}, common.Struct_l_l{(*int8)(unsafe.Pointer(&[51]int8{'/', 'U', 's', 'e', 'r', 's', '/', 'x', 'u', 's', 'h', 'i', 'w', 'e', 'i', '/', 'w', 'o', 'r', 'k', '/', 'l', 'i', 'b', 'c', '/', 't', 'e', 's', 't', '/', 's', 'r', 'c', '/', 'm', 'a', 't', 'h', '/', 'u', 'c', 'b', '/', 'f', 'a', 'b', 's', '.', 'h', '\x00'})), int32(56), -1, float64(8.9884656743115795e+307), float64(8.9884656743115795e+307), float32(0), int32(0)}, common.Struct_l_l{(*int8)(unsafe.Pointer(&[51]int8{'/', 'U', 's', 'e', 'r', 's', '/', 'x', 'u', 's', 'h', 'i', 'w', 'e', 'i', '/', 'w', 'o', 'r', 'k', '/', 'l', 'i', 'b', 'c', '/', 't', 'e', 's', 't', '/', 's', 'r', 'c', '/', 'm', 'a', 't', 'h', '/', 'u', 'c', 'b', '/', 'f', 'a', 'b', 's', '.', 'h', '\x00'})), int32(57), -1, float64(1.7976931348623157e+308), float64(1.7976931348623157e+308), float32(0), int32(0)}, common.Struct_l_l{(*int8)(unsafe.Pointer(&[51]int8{'/', 'U', 's', 'e', 'r', 's', '/', 'x', 'u', 's', 'h', 'i', 'w', 'e', 'i', '/', 'w', 'o', 'r', 'k', '/', 'l', 'i', 'b', 'c', '/', 't', 'e', 's', 't', '/', 's', 'r', 'c', '/', 'm', 'a', 't', 'h', '/', 'u', 'c', 'b', '/', 'f', 'a', 'b', 's', '.', 'h', '\x00'})), int32(58), -1, float64(libc.X__builtin_inff()), float64(libc.X__builtin_inff()), float32(0), int32(0)}, common.Struct_l_l{(*int8)(unsafe.Pointer(&[51]int8{'/', 'U', 's', 'e', 'r', 's', '/', 'x', 'u', 's', 'h', 'i', 'w', 'e', 'i', '/', 'w', 'o', 'r', 'k', '/', 'l', 'i', 'b', 'c', '/', 't', 'e', 's', 't', '/', 's', 'r', 'c', '/', 'm', 'a', 't', 'h', '/', 'u', 'c', 'b', '/', 'f', 'a', 'b', 's', '.', 'h', '\x00'})), int32(59), -1, float64(-0), float64(0), float32(0), int32(0)}, common.Struct_l_l{(*int8)(unsafe.Pointer(&[51]int8{'/', 'U', 's', 'e', 'r', 's', '/', 'x', 'u', 's', 'h', 'i', 'w', 'e', 'i', '/', 'w', 'o', 'r', 'k', '/', 'l', 'i', 'b', 'c', '/', 't', 'e', 's', 't', '/', 's', 'r', 'c', '/', 'm', 'a', 't', 'h', '/', 'u', 'c', 'b', '/', 'f', 'a', 'b', 's', '.', 'h', '\x00'})), int32(60), -1, float64(-4.9406564584124654e-324), float64(4.9406564584124654e-324), float32(0), int32(0)}, common.Struct_l_l{(*int8)(unsafe.Pointer(&[51]int8{'/', 'U', 's', 'e', 'r', 's', '/', 'x', 'u', 's', 'h', 'i', 'w', 'e', 'i', '/', 'w', 'o', 'r', 'k', '/', 'l', 'i', 'b', 'c', '/', 't', 'e', 's', 't', '/', 's', 'r', 'c', '/', 'm', 'a', 't', 'h', '/', 'u', 'c', 'b', '/', 'f', 'a', 'b', 's', '.', 'h', '\x00'})), int32(61), -1, float64(-2.2250738585072009e-308), float64(2.2250738585072009e-308), float32(0), int32(0)}, common.Struct_l_l{(*int8)(unsafe.Pointer(&[51]int8{'/', 'U', 's', 'e', 'r', 's', '/', 'x', 'u', 's', 'h', 'i', 'w', 'e', 'i', '/', 'w', 'o', 'r', 'k', '/', 'l', 'i', 'b', 'c', '/', 't', 'e', 's', 't', '/', 's', 'r', 'c', '/', 'm', 'a', 't', 'h', '/', 'u', 'c', 'b', '/', 'f', 'a', 'b', 's', '.', 'h', '\x00'})), int32(62), -1, float64(-1), float64(1), float32(0), int32(0)}, common.Struct_l_l{(*int8)(unsafe.Pointer(&[51]int8{'/', 'U', 's', 'e', 'r', 's', '/', 'x', 'u', 's', 'h', 'i', 'w', 'e', 'i', '/', 'w', 'o', 'r', 'k', '/', 'l', 'i', 'b', 'c', '/', 't', 'e', 's', 't', '/', 's', 'r', 'c', '/', 'm', 'a', 't', 'h', '/', 'u', 'c', 'b', '/', 'f', 'a', 'b', 's', '.', 'h', '\x00'})), int32(63), -1, float64(-8.9884656743115795e+307), float64(8.9884656743115795e+307), float32(0), int32(0)}, common.Struct_l_l{(*int8)(unsafe.Pointer(&[51]int8{'/', 'U', 's', 'e', 'r', 's', '/', 'x', 'u', 's', 'h', 'i', 'w', 'e', 'i', '/', 'w', 'o', 'r', 'k', '/', 'l', 'i', 'b', 'c', '/', 't', 'e', 's', 't', '/', 's', 'r', 'c', '/', 'm', 'a', 't', 'h', '/', 'u', 'c', 'b', '/', 'f', 'a', 'b', 's', '.', 'h', '\x00'})), int32(64), -1, float64(-1.7976931348623157e+308), float64(1.7976931348623157e+308), float32(0), int32(0)}, common.Struct_l_l{(*int8)(unsafe.Pointer(&[51]int8{'/', 'U', 's', 'e', 'r', 's', '/', 'x', 'u', 's', 'h', 'i', 'w', 'e', 'i', '/', 'w', 'o', 'r', 'k', '/', 'l', 'i', 'b', 'c', '/', 't', 'e', 's', 't', '/', 's', 'r', 'c', '/', 'm', 'a', 't', 'h', '/', 'u', 'c', 'b', '/', 'f', 'a', 'b', 's', '.', 'h', '\x00'})), int32(65), -1, float64(-libc.X__builtin_inff()), float64(libc.X__builtin_inff()), float32(0), int32(0)}, common.Struct_l_l{(*int8)(unsafe.Pointer(&[51]int8{'/', 'U', 's', 'e', 'r', 's', '/', 'x', 'u', 's', 'h', 'i', 'w', 'e', 'i', '/', 'w', 'o', 'r', 'k', '/', 'l', 'i', 'b', 'c', '/', 't', 'e', 's', 't', '/', 's', 'r', 'c', '/', 'm', 'a', 't', 'h', '/', 'u', 'c', 'b', '/', 'f', 'a', 'b', 's', '.', 'h', '\x00'})), int32(66), -1, float64(libc.X__builtin_nanf((*int8)(unsafe.Pointer(&[1]int8{'\x00'})))), float64(libc.X__builtin_nanf((*int8)(unsafe.Pointer(&[1]int8{'\x00'})))), float32(0), int32(0)}, common.Struct_l_l{(*int8)(unsafe.Pointer(&[51]int8{'/', 'U', 's', 'e', 'r', 's', '/', 'x', 'u', 's', 'h', 'i', 'w', 'e', 'i', '/', 'w', 'o', 'r', 'k', '/', 'l', 'i', 'b', 'c', '/', 't', 'e', 's', 't', '/', 's', 'r', 'c', '/', 'm', 'a', 't', 'h', '/', 'u', 'c', 'b', '/', 'f', 'a', 'b', 's', '.', 'h', '\x00'})), int32(67), -1, float64(libc.X__builtin_nanf((*int8)(unsafe.Pointer(&[1]int8{'\x00'})))), float64(libc.X__builtin_nanf((*int8)(unsafe.Pointer(&[1]int8{'\x00'})))), float32(0), int32(0)}, common.Struct_l_l{(*int8)(unsafe.Pointer(&[51]int8{'/', 'U', 's', 'e', 'r', 's', '/', 'x', 'u', 's', 'h', 'i', 'w', 'e', 'i', '/', 'w', 'o', 'r', 'k', '/', 'l', 'i', 'b', 'c', '/', 't', 'e', 's', 't', '/', 's', 'r', 'c', '/', 'm', 'a', 't', 'h', '/', 'u', 'c', 'b', '/', 'f', 'a', 'b', 's', '.', 'h', '\x00'})), int32(68), int32(0), float64(2.2250738585072009e-308), float64(2.2250738585072009e-308), float32(0), int32(0)}, common.Struct_l_l{(*int8)(unsafe.Pointer(&[51]int8{'/', 'U', 's', 'e', 'r', 's', '/', 'x', 'u', 's', 'h', 'i', 'w', 'e', 'i', '/', 'w', 'o', 'r', 'k', '/', 'l', 'i', 'b', 'c', '/', 't', 'e', 's', 't', '/', 's', 'r', 'c', '/', 'm', 'a', 't', 'h', '/', 'u', 'c', 'b', '/', 'f', 'a', 'b', 's', '.', 'h', '\x00'})), int32(69), int32(0), float64(8.9884656743115795e+307), float64(8.9884656743115795e+307), float32(0), int32(0)}, common.Struct_l_l{(*int8)(unsafe.Pointer(&[51]int8{'/', 'U', 's', 'e', 'r', 's', '/', 'x', 'u', 's', 'h', 'i', 'w', 'e', 'i', '/', 'w', 'o', 'r', 'k', '/', 'l', 'i', 'b', 'c', '/', 't', 'e', 's', 't', '/', 's', 'r', 'c', '/', 'm', 'a', 't', 'h', '/', 'u', 'c', 'b', '/', 'f', 'a', 'b', 's', '.', 'h', '\x00'})), int32(70), int32(0), float64(-2.2250738585072009e-308), float64(2.2250738585072009e-308), float32(0), int32(0)}, common.Struct_l_l{(*int8)(unsafe.Pointer(&[51]int8{'/', 'U', 's', 'e', 'r', 's', '/', 'x', 'u', 's', 'h', 'i', 'w', 'e', 'i', '/', 'w', 'o', 'r', 'k', '/', 'l', 'i', 'b', 'c', '/', 't', 'e', 's', 't', '/', 's', 'r', 'c', '/', 'm', 'a', 't', 'h', '/', 'u', 'c', 'b', '/', 'f', 'a', 'b', 's', '.', 'h', '\x00'})), int32(71), int32(0), float64(-1), float64(1), float32(0), int32(0)}, common.Struct_l_l{(*int8)(unsafe.Pointer(&[51]int8{'/', 'U', 's', 'e', 'r', 's', '/', 'x', 'u', 's', 'h', 'i', 'w', 'e', 'i', '/', 'w', 'o', 'r', 'k', '/', 'l', 'i', 'b', 'c', '/', 't', 'e', 's', 't', '/', 's', 'r', 'c', '/', 'm', 'a', 't', 'h', '/', 'u', 'c', 'b', '/', 'f', 'a', 'b', 's', '.', 'h', '\x00'})), int32(72), int32(0), float64(-8.9884656743115795e+307), float64(8.9884656743115795e+307), float32(0), int32(0)}, common.Struct_l_l{(*int8)(unsafe.Pointer(&[51]int8{'/', 'U', 's', 'e', 'r', 's', '/', 'x', 'u', 's', 'h', 'i', 'w', 'e', 'i', '/', 'w', 'o', 'r', 'k', '/', 'l', 'i', 'b', 'c', '/', 't', 'e', 's', 't', '/', 's', 'r', 'c', '/', 'm', 'a', 't', 'h', '/', 'u', 'c', 'b', '/', 'f', 'a', 'b', 's', '.', 'h', '\x00'})), int32(73), -1, float64(0), float64(0), float32(0), int32(0)}, common.Struct_l_l{(*int8)(unsafe.Pointer(&[51]int8{'/', 'U', 's', 'e', 'r', 's', '/', 'x', 'u', 's', 'h', 'i', 'w', 'e', 'i', '/', 'w', 'o', 'r', 'k', '/', 'l', 'i', 'b', 'c', '/', 't', 'e', 's', 't', '/', 's', 'r', 'c', '/', 'm', 'a', 't', 'h', '/', 'u', 'c', 'b', '/', 'f', 'a', 'b', 's', '.', 'h', '\x00'})), int32(74), -1, float64(4.9406564584124654e-324), float64(4.9406564584124654e-324), float32(0), int32(0)}, common.Struct_l_l{(*int8)(unsafe.Pointer(&[51]int8{'/', 'U', 's', 'e', 'r', 's', '/', 'x', 'u', 's', 'h', 'i', 'w', 'e', 'i', '/', 'w', 'o', 'r', 'k', '/', 'l', 'i', 'b', 'c', '/', 't', 'e', 's', 't', '/', 's', 'r', 'c', '/', 'm', 'a', 't', 'h', '/', 'u', 'c', 'b', '/', 'f', 'a', 'b', 's', '.', 'h', '\x00'})), int32(75), -1, float64(2.2250738585072009e-308), float64(2.2250738585072009e-308), float32(0), int32(0)}, common.Struct_l_l{(*int8)(unsafe.Pointer(&[51]int8{'/', 'U', 's', 'e', 'r', 's', '/', 'x', 'u', 's', 'h', 'i', 'w', 'e', 'i', '/', 'w', 'o', 'r', 'k', '/', 'l', 'i', 'b', 'c', '/', 't', 'e', 's', 't', '/', 's', 'r', 'c', '/', 'm', 'a', 't', 'h', '/', 'u', 'c', 'b', '/', 'f', 'a', 'b', 's', '.', 'h', '\x00'})), int32(76), -1, float64(1), float64(1), float32(0), int32(0)}, common.Struct_l_l{(*int8)(unsafe.Pointer(&[51]int8{'/', 'U', 's', 'e', 'r', 's', '/', 'x', 'u', 's', 'h', 'i', 'w', 'e', 'i', '/', 'w', 'o', 'r', 'k', '/', 'l', 'i', 'b', 'c', '/', 't', 'e', 's', 't', '/', 's', 'r', 'c', '/', 'm', 'a', 't', 'h', '/', 'u', 'c', 'b', '/', 'f', 'a', 'b', 's', '.', 'h', '\x00'})), int32(77), -1, float64(8.9884656743115795e+307), float64(8.9884656743115795e+307), float32(0), int32(0)}, common.Struct_l_l{(*int8)(unsafe.Pointer(&[51]int8{'/', 'U', 's', 'e', 'r', 's', '/', 'x', 'u', 's', 'h', 'i', 'w', 'e', 'i', '/', 'w', 'o', 'r', 'k', '/', 'l', 'i', 'b', 'c', '/', 't', 'e', 's', 't', '/', 's', 'r', 'c', '/', 'm', 'a', 't', 'h', '/', 'u', 'c', 'b', '/', 'f', 'a', 'b', 's', '.', 'h', '\x00'})), int32(78), -1, float64(1.7976931348623157e+308), float64(1.7976931348623157e+308), float32(0), int32(0)}, common.Struct_l_l{(*int8)(unsafe.Pointer(&[51]int8{'/', 'U', 's', 'e', 'r', 's', '/', 'x', 'u', 's', 'h', 'i', 'w', 'e', 'i', '/', 'w', 'o', 'r', 'k', '/', 'l', 'i', 'b', 'c', '/', 't', 'e', 's', 't', '/', 's', 'r', 'c', '/', 'm', 'a', 't', 'h', '/', 'u', 'c', 'b', '/', 'f', 'a', 'b', 's', '.', 'h', '\x00'})), int32(79), -1, float64(libc.X__builtin_inff()), float64(libc.X__builtin_inff()), float32(0), int32(0)}, common.Struct_l_l{(*int8)(unsafe.Pointer(&[51]int8{'/', 'U', 's', 'e', 'r', 's', '/', 'x', 'u', 's', 'h', 'i', 'w', 'e', 'i', '/', 'w', 'o', 'r', 'k', '/', 'l', 'i', 'b', 'c', '/', 't', 'e', 's', 't', '/', 's', 'r', 'c', '/', 'm', 'a', 't', 'h', '/', 'u', 'c', 'b', '/', 'f', 'a', 'b', 's', '.', 'h', '\x00'})), int32(80), -1, float64(-0), float64(0), float32(0), int32(0)}, common.Struct_l_l{(*int8)(unsafe.Pointer(&[51]int8{'/', 'U', 's', 'e', 'r', 's', '/', 'x', 'u', 's', 'h', 'i', 'w', 'e', 'i', '/', 'w', 'o', 'r', 'k', '/', 'l', 'i', 'b', 'c', '/', 't', 'e', 's', 't', '/', 's', 'r', 'c', '/', 'm', 'a', 't', 'h', '/', 'u', 'c', 'b', '/', 'f', 'a', 'b', 's', '.', 'h', '\x00'})), int32(81), -1, float64(-4.9406564584124654e-324), float64(4.9406564584124654e-324), float32(0), int32(0)}, common.Struct_l_l{(*int8)(unsafe.Pointer(&[51]int8{'/', 'U', 's', 'e', 'r', 's', '/', 'x', 'u', 's', 'h', 'i', 'w', 'e', 'i', '/', 'w', 'o', 'r', 'k', '/', 'l', 'i', 'b', 'c', '/', 't', 'e', 's', 't', '/', 's', 'r', 'c', '/', 'm', 'a', 't', 'h', '/', 'u', 'c', 'b', '/', 'f', 'a', 'b', 's', '.', 'h', '\x00'})), int32(82), -1, float64(-2.2250738585072009e-308), float64(2.2250738585072009e-308), float32(0), int32(0)}, common.Struct_l_l{(*int8)(unsafe.Pointer(&[51]int8{'/', 'U', 's', 'e', 'r', 's', '/', 'x', 'u', 's', 'h', 'i', 'w', 'e', 'i', '/', 'w', 'o', 'r', 'k', '/', 'l', 'i', 'b', 'c', '/', 't', 'e', 's', 't', '/', 's', 'r', 'c', '/', 'm', 'a', 't', 'h', '/', 'u', 'c', 'b', '/', 'f', 'a', 'b', 's', '.', 'h', '\x00'})), int32(83), -1, float64(-1), float64(1), float32(0), int32(0)}, common.Struct_l_l{(*int8)(unsafe.Pointer(&[51]int8{'/', 'U', 's', 'e', 'r', 's', '/', 'x', 'u', 's', 'h', 'i', 'w', 'e', 'i', '/', 'w', 'o', 'r', 'k', '/', 'l', 'i', 'b', 'c', '/', 't', 'e', 's', 't', '/', 's', 'r', 'c', '/', 'm', 'a', 't', 'h', '/', 'u', 'c', 'b', '/', 'f', 'a', 'b', 's', '.', 'h', '\x00'})), int32(84), -1, float64(-8.9884656743115795e+307), float64(8.9884656743115795e+307), float32(0), int32(0)}, common.Struct_l_l{(*int8)(unsafe.Pointer(&[51]int8{'/', 'U', 's', 'e', 'r', 's', '/', 'x', 'u', 's', 'h', 'i', 'w', 'e', 'i', '/', 'w', 'o', 'r', 'k', '/', 'l', 'i', 'b', 'c', '/', 't', 'e', 's', 't', '/', 's', 'r', 'c', '/', 'm', 'a', 't', 'h', '/', 'u', 'c', 'b', '/', 'f', 'a', 'b', 's', '.', 'h', '\x00'})), int32(85), -1, float64(-1.7976931348623157e+308), float64(1.7976931348623157e+308), float32(0), int32(0)}, common.Struct_l_l{(*int8)(unsafe.Pointer(&[51]int8{'/', 'U', 's', 'e', 'r', 's', '/', 'x', 'u', 's', 'h', 'i', 'w', 'e', 'i', '/', 'w', 'o', 'r', 'k', '/', 'l', 'i', 'b', 'c', '/', 't', 'e', 's', 't', '/', 's', 'r', 'c', '/', 'm', 'a', 't', 'h', '/', 'u', 'c', 'b', '/', 'f', 'a', 'b', 's', '.', 'h', '\x00'})), int32(86), -1, float64(-libc.X__builtin_inff()), float64(libc.X__builtin_inff()), float32(0), int32(0)}, common.Struct_l_l{(*int8)(unsafe.Pointer(&[51]int8{'/', 'U', 's', 'e', 'r', 's', '/', 'x', 'u', 's', 'h', 'i', 'w', 'e', 'i', '/', 'w', 'o', 'r', 'k', '/', 'l', 'i', 'b', 'c', '/', 't', 'e', 's', 't', '/', 's', 'r', 'c', '/', 'm', 'a', 't', 'h', '/', 'u', 'c', 'b', '/', 'f', 'a', 'b', 's', '.', 'h', '\x00'})), int32(87), -1, float64(libc.X__builtin_nanf((*int8)(unsafe.Pointer(&[1]int8{'\x00'})))), float64(libc.X__builtin_nanf((*int8)(unsafe.Pointer(&[1]int8{'\x00'})))), float32(0), int32(0)}, common.Struct_l_l{(*int8)(unsafe.Pointer(&[51]int8{'/', 'U', 's', 'e', 'r', 's', '/', 'x', 'u', 's', 'h', 'i', 'w', 'e', 'i', '/', 'w', 'o', 'r', 'k', '/', 'l', 'i', 'b', 'c', '/', 't', 'e', 's', 't', '/', 's', 'r', 'c', '/', 'm', 'a', 't', 'h', '/', 'u', 'c', 'b', '/', 'f', 'a', 'b', 's', '.', 'h', '\x00'})), int32(88), -1, float64(libc.X__builtin_nanf((*int8)(unsafe.Pointer(&[1]int8{'\x00'})))), float64(libc.X__builtin_nanf((*int8)(unsafe.Pointer(&[1]int8{'\x00'})))), float32(0), int32(0)}, common.Struct_l_l{(*int8)(unsafe.Pointer(&[51]int8{'/', 'U', 's', 'e', 'r', 's', '/', 'x', 'u', 's', 'h', 'i', 'w', 'e', 'i', '/', 'w', 'o', 'r', 'k', '/', 'l', 'i', 'b', 'c', '/', 't', 'e', 's', 't', '/', 's', 'r', 'c', '/', 'm', 'a', 't', 'h', '/', 'u', 'c', 'b', '/', 'f', 'a', 'b', 's', '.', 'h', '\x00'})), int32(89), -1, float64(0), float64(0), float32(0), int32(0)}, common.Struct_l_l{(*int8)(unsafe.Pointer(&[51]int8{'/', 'U', 's', 'e', 'r', 's', '/', 'x', 'u', 's', 'h', 'i', 'w', 'e', 'i', '/', 'w', 'o', 'r', 'k', '/', 'l', 'i', 'b', 'c', '/', 't', 'e', 's', 't', '/', 's', 'r', 'c', '/', 'm', 'a', 't', 'h', '/', 'u', 'c', 'b', '/', 'f', 'a', 'b', 's', '.', 'h', '\x00'})), int32(90), -1, float64(4.9406564584124654e-324), float64(4.9406564584124654e-324), float32(0), int32(0)}, common.Struct_l_l{(*int8)(unsafe.Pointer(&[51]int8{'/', 'U', 's', 'e', 'r', 's', '/', 'x', 'u', 's', 'h', 'i', 'w', 'e', 'i', '/', 'w', 'o', 'r', 'k', '/', 'l', 'i', 'b', 'c', '/', 't', 'e', 's', 't', '/', 's', 'r', 'c', '/', 'm', 'a', 't', 'h', '/', 'u', 'c', 'b', '/', 'f', 'a', 'b', 's', '.', 'h', '\x00'})), int32(91), -1, float64(2.2250738585072009e-308), float64(2.2250738585072009e-308), float32(0), int32(0)}, common.Struct_l_l{(*int8)(unsafe.Pointer(&[51]int8{'/', 'U', 's', 'e', 'r', 's', '/', 'x', 'u', 's', 'h', 'i', 'w', 'e', 'i', '/', 'w', 'o', 'r', 'k', '/', 'l', 'i', 'b', 'c', '/', 't', 'e', 's', 't', '/', 's', 'r', 'c', '/', 'm', 'a', 't', 'h', '/', 'u', 'c', 'b', '/', 'f', 'a', 'b', 's', '.', 'h', '\x00'})), int32(92), -1, float64(1), float64(1), float32(0), int32(0)}, common.Struct_l_l{(*int8)(unsafe.Pointer(&[51]int8{'/', 'U', 's', 'e', 'r', 's', '/', 'x', 'u', 's', 'h', 'i', 'w', 'e', 'i', '/', 'w', 'o', 'r', 'k', '/', 'l', 'i', 'b', 'c', '/', 't', 'e', 's', 't', '/', 's', 'r', 'c', '/', 'm', 'a', 't', 'h', '/', 'u', 'c', 'b', '/', 'f', 'a', 'b', 's', '.', 'h', '\x00'})), int32(93), -1, float64(8.9884656743115795e+307), float64(8.9884656743115795e+307), float32(0), int32(0)}, common.Struct_l_l{(*int8)(unsafe.Pointer(&[51]int8{'/', 'U', 's', 'e', 'r', 's', '/', 'x', 'u', 's', 'h', 'i', 'w', 'e', 'i', '/', 'w', 'o', 'r', 'k', '/', 'l', 'i', 'b', 'c', '/', 't', 'e', 's', 't', '/', 's', 'r', 'c', '/', 'm', 'a', 't', 'h', '/', 'u', 'c', 'b', '/', 'f', 'a', 'b', 's', '.', 'h', '\x00'})), int32(94), -1, float64(1.7976931348623157e+308), float64(1.7976931348623157e+308), float32(0), int32(0)}, common.Struct_l_l{(*int8)(unsafe.Pointer(&[51]int8{'/', 'U', 's', 'e', 'r', 's', '/', 'x', 'u', 's', 'h', 'i', 'w', 'e', 'i', '/', 'w', 'o', 'r', 'k', '/', 'l', 'i', 'b', 'c', '/', 't', 'e', 's', 't', '/', 's', 'r', 'c', '/', 'm', 'a', 't', 'h', '/', 'u', 'c', 'b', '/', 'f', 'a', 'b', 's', '.', 'h', '\x00'})), int32(95), -1, float64(libc.X__builtin_inff()), float64(libc.X__builtin_inff()), float32(0), int32(0)}, common.Struct_l_l{(*int8)(unsafe.Pointer(&[51]int8{'/', 'U', 's', 'e', 'r', 's', '/', 'x', 'u', 's', 'h', 'i', 'w', 'e', 'i', '/', 'w', 'o', 'r', 'k', '/', 'l', 'i', 'b', 'c', '/', 't', 'e', 's', 't', '/', 's', 'r', 'c', '/', 'm', 'a', 't', 'h', '/', 'u', 'c', 'b', '/', 'f', 'a', 'b', 's', '.', 'h', '\x00'})), int32(96), -1, float64(-0), float64(0), float32(0), int32(0)}, common.Struct_l_l{(*int8)(unsafe.Pointer(&[51]int8{'/', 'U', 's', 'e', 'r', 's', '/', 'x', 'u', 's', 'h', 'i', 'w', 'e', 'i', '/', 'w', 'o', 'r', 'k', '/', 'l', 'i', 'b', 'c', '/', 't', 'e', 's', 't', '/', 's', 'r', 'c', '/', 'm', 'a', 't', 'h', '/', 'u', 'c', 'b', '/', 'f', 'a', 'b', 's', '.', 'h', '\x00'})), int32(97), -1, float64(-4.9406564584124654e-324), float64(4.9406564584124654e-324), float32(0), int32(0)}, common.Struct_l_l{(*int8)(unsafe.Pointer(&[51]int8{'/', 'U', 's', 'e', 'r', 's', '/', 'x', 'u', 's', 'h', 'i', 'w', 'e', 'i', '/', 'w', 'o', 'r', 'k', '/', 'l', 'i', 'b', 'c', '/', 't', 'e', 's', 't', '/', 's', 'r', 'c', '/', 'm', 'a', 't', 'h', '/', 'u', 'c', 'b', '/', 'f', 'a', 'b', 's', '.', 'h', '\x00'})), int32(98), -1, float64(-2.2250738585072009e-308), float64(2.2250738585072009e-308), float32(0), int32(0)}, common.Struct_l_l{(*int8)(unsafe.Pointer(&[51]int8{'/', 'U', 's', 'e', 'r', 's', '/', 'x', 'u', 's', 'h', 'i', 'w', 'e', 'i', '/', 'w', 'o', 'r', 'k', '/', 'l', 'i', 'b', 'c', '/', 't', 'e', 's', 't', '/', 's', 'r', 'c', '/', 'm', 'a', 't', 'h', '/', 'u', 'c', 'b', '/', 'f', 'a', 'b', 's', '.', 'h', '\x00'})), int32(99), -1, float64(-1), float64(1), float32(0), int32(0)}, common.Struct_l_l{(*int8)(unsafe.Pointer(&[51]int8{'/', 'U', 's', 'e', 'r', 's', '/', 'x', 'u', 's', 'h', 'i', 'w', 'e', 'i', '/', 'w', 'o', 'r', 'k', '/', 'l', 'i', 'b', 'c', '/', 't', 'e', 's', 't', '/', 's', 'r', 'c', '/', 'm', 'a', 't', 'h', '/', 'u', 'c', 'b', '/', 'f', 'a', 'b', 's', '.', 'h', '\x00'})), int32(100), -1, float64(-8.9884656743115795e+307), float64(8.9884656743115795e+307), float32(0), int32(0)}, common.Struct_l_l{(*int8)(unsafe.Pointer(&[51]int8{'/', 'U', 's', 'e', 'r', 's', '/', 'x', 'u', 's', 'h', 'i', 'w', 'e', 'i', '/', 'w', 'o', 'r', 'k', '/', 'l', 'i', 'b', 'c', '/', 't', 'e', 's', 't', '/', 's', 'r', 'c', '/', 'm', 'a', 't', 'h', '/', 'u', 'c', 'b', '/', 'f', 'a', 'b', 's', '.', 'h', '\x00'})), int32(101), -1, float64(-1.7976931348623157e+308), float64(1.7976931348623157e+308), float32(0), int32(0)}, common.Struct_l_l{(*int8)(unsafe.Pointer(&[51]int8{'/', 'U', 's', 'e', 'r', 's', '/', 'x', 'u', 's', 'h', 'i', 'w', 'e', 'i', '/', 'w', 'o', 'r', 'k', '/', 'l', 'i', 'b', 'c', '/', 't', 'e', 's', 't', '/', 's', 'r', 'c', '/', 'm', 'a', 't', 'h', '/', 'u', 'c', 'b', '/', 'f', 'a', 'b', 's', '.', 'h', '\x00'})), int32(102), -1, float64(-libc.X__builtin_inff()), float64(libc.X__builtin_inff()), float32(0), int32(0)}, common.Struct_l_l{(*int8)(unsafe.Pointer(&[51]int8{'/', 'U', 's', 'e', 'r', 's', '/', 'x', 'u', 's', 'h', 'i', 'w', 'e', 'i', '/', 'w', 'o', 'r', 'k', '/', 'l', 'i', 'b', 'c', '/', 't', 'e', 's', 't', '/', 's', 'r', 'c', '/', 'm', 'a', 't', 'h', '/', 'u', 'c', 'b', '/', 'f', 'a', 'b', 's', '.', 'h', '\x00'})), int32(103), -1, float64(libc.X__builtin_nanf((*int8)(unsafe.Pointer(&[1]int8{'\x00'})))), float64(libc.X__builtin_nanf((*int8)(unsafe.Pointer(&[1]int8{'\x00'})))), float32(0), int32(0)}, common.Struct_l_l{(*int8)(unsafe.Pointer(&[51]int8{'/', 'U', 's', 'e', 'r', 's', '/', 'x', 'u', 's', 'h', 'i', 'w', 'e', 'i', '/', 'w', 'o', 'r', 'k', '/', 'l', 'i', 'b', 'c', '/', 't', 'e', 's', 't', '/', 's', 'r', 'c', '/', 'm', 'a', 't', 'h', '/', 'u', 'c', 'b', '/', 'f', 'a', 'b', 's', '.', 'h', '\x00'})), int32(104), -1, float64(libc.X__builtin_nanf((*int8)(unsafe.Pointer(&[1]int8{'\x00'})))), float64(libc.X__builtin_nanf((*int8)(unsafe.Pointer(&[1]int8{'\x00'})))), float32(0), int32(0)}, common.Struct_l_l{(*int8)(unsafe.Pointer(&[54]int8{'/', 'U', 's', 'e', 'r', 's', '/', 'x', 'u', 's', 'h', 'i', 'w', 'e', 'i', '/', 'w', 'o', 'r', 'k', '/', 'l', 'i', 'b', 'c', '/', 't', 'e', 's', 't', '/', 's', 'r', 'c', '/', 'm', 'a', 't', 'h', '/', 's', 'a', 'n', 'i', 't', 'y', '/', 'f', 'a', 'b', 's', '.', 'h', '\x00'})), int32(1), int32(0), float64(-8.0668483905796808), float64(8.0668483905796808), float32(0), int32(0)}, common.Struct_l_l{(*int8)(unsafe.Pointer(&[54]int8{'/', 'U', 's', 'e', 'r', 's', '/', 'x', 'u', 's', 'h', 'i', 'w', 'e', 'i', '/', 'w', 'o', 'r', 'k', '/', 'l', 'i', 'b', 'c', '/', 't', 'e', 's', 't', '/', 's', 'r', 'c', '/', 'm', 'a', 't', 'h', '/', 's', 'a', 'n', 'i', 't', 'y', '/', 'f', 'a', 'b', 's', '.', 'h', '\x00'})), int32(2), int32(0), float64(4.3452398493383049), float64(4.3452398493383049), float32(0), int32(0)}, common.Struct_l_l{(*int8)(unsafe.Pointer(&[54]int8{'/', 'U', 's', 'e', 'r', 's', '/', 'x', 'u', 's', 'h', 'i', 'w', 'e', 'i', '/', 'w', 'o', 'r', 'k', '/', 'l', 'i', 'b', 'c', '/', 't', 'e', 's', 't', '/', 's', 'r', 'c', '/', 'm', 'a', 't', 'h', '/', 's', 'a', 'n', 'i', 't', 'y', '/', 'f', 'a', 'b', 's', '.', 'h', '\x00'})), int32(3), int32(0), float64(-8.3814334275552493), float64(8.3814334275552493), float32(0), int32(0)}, common.Struct_l_l{(*int8)(unsafe.Pointer(&[54]int8{'/', 'U', 's', 'e', 'r', 's', '/', 'x', 'u', 's', 'h', 'i', 'w', 'e', 'i', '/', 'w', 'o', 'r', 'k', '/', 'l', 'i', 'b', 'c', '/', 't', 'e', 's', 't', '/', 's', 'r', 'c', '/', 'm', 'a', 't', 'h', '/', 's', 'a', 'n', 'i', 't', 'y', '/', 'f', 'a', 'b', 's', '.', 'h', '\x00'})), int32(4), int32(0), float64(-6.5316735819134841), float64(6.5316735819134841), float32(0), int32(0)}, common.Struct_l_l{(*int8)(unsafe.Pointer(&[54]int8{'/', 'U', 's', 'e', 'r', 's', '/', 'x', 'u', 's', 'h', 'i', 'w', 'e', 'i', '/', 'w', 'o', 'r', 'k', '/', 'l', 'i', 'b', 'c', '/', 't', 'e', 's', 't', '/', 's', 'r', 'c', '/', 'm', 'a', 't', 'h', '/', 's', 'a', 'n', 'i', 't', 'y', '/', 'f', 'a', 'b', 's', '.', 'h', '\x00'})), int32(5), int32(0), float64(9.2670569669725857), float64(9.2670569669725857), float32(0), int32(0)}, common.Struct_l_l{(*int8)(unsafe.Pointer(&[54]int8{'/', 'U', 's', 'e', 'r', 's', '/', 'x', 'u', 's', 'h', 'i', 'w', 'e', 'i', '/', 'w', 'o', 'r', 'k', '/', 'l', 'i', 'b', 'c', '/', 't', 'e', 's', 't', '/', 's', 'r', 'c', '/', 'm', 'a', 't', 'h', '/', 's', 'a', 'n', 'i', 't', 'y', '/', 'f', 'a', 'b', 's', '.', 'h', '\x00'})), int32(6), int32(0), float64(0.66198589809950448), float64(0.66198589809950448), float32(0), int32(0)}, common.Struct_l_l{(*int8)(unsafe.Pointer(&[54]int8{'/', 'U', 's', 'e', 'r', 's', '/', 'x', 'u', 's', 'h', 'i', 'w', 'e', 'i', '/', 'w', 'o', 'r', 'k', '/', 'l', 'i', 'b', 'c', '/', 't', 'e', 's', 't', '/', 's', 'r', 'c', '/', 'm', 'a', 't', 'h', '/', 's', 'a', 'n', 'i', 't', 'y', '/', 'f', 'a', 'b', 's', '.', 'h', '\x00'})), int32(7), int32(0), float64(-0.40660392238535531), float64(0.40660392238535531), float32(0), int32(0)}, common.Struct_l_l{(*int8)(unsafe.Pointer(&[54]int8{'/', 'U', 's', 'e', 'r', 's', '/', 'x', 'u', 's', 'h', 'i', 'w', 'e', 'i', '/', 'w', 'o', 'r', 'k', '/', 'l', 'i', 'b', 'c', '/', 't', 'e', 's', 't', '/', 's', 'r', 'c', '/', 'm', 'a', 't', 'h', '/', 's', 'a', 'n', 'i', 't', 'y', '/', 'f', 'a', 'b', 's', '.', 'h', '\x00'})), int32(8), int32(0), float64(0.56175974622072411), float64(0.56175974622072411), float32(0), int32(0)}, common.Struct_l_l{(*int8)(unsafe.Pointer(&[54]int8{'/', 'U', 's', 'e', 'r', 's', '/', 'x', 'u', 's', 'h', 'i', 'w', 'e', 'i', '/', 'w', 'o', 'r', 'k', '/', 'l', 'i', 'b', 'c', '/', 't', 'e', 's', 't', '/', 's', 'r', 'c', '/', 'm', 'a', 't', 'h', '/', 's', 'a', 'n', 'i', 't', 'y', '/', 'f', 'a', 'b', 's', '.', 'h', '\x00'})), int32(9), int32(0), float64(0.77415229659130369), float64(0.77415229659130369), float32(0), int32(0)}, common.Struct_l_l{(*int8)(unsafe.Pointer(&[54]int8{'/', 'U', 's', 'e', 'r', 's', '/', 'x', 'u', 's', 'h', 'i', 'w', 'e', 'i', '/', 'w', 'o', 'r', 'k', '/', 'l', 'i', 'b', 'c', '/', 't', 'e', 's', 't', '/', 's', 'r', 'c', '/', 'm', 'a', 't', 'h', '/', 's', 'a', 'n', 'i', 't', 'y', '/', 'f', 'a', 'b', 's', '.', 'h', '\x00'})), int32(10), int32(0), float64(-0.67876370263940244), float64(0.67876370263940244), float32(0), int32(0)}, common.Struct_l_l{(*int8)(unsafe.Pointer(&[55]int8{'/', 'U', 's', 'e', 'r', 's', '/', 'x', 'u', 's', 'h', 'i', 'w', 'e', 'i', '/', 'w', 'o', 'r', 'k', '/', 'l', 'i', 'b', 'c', '/', 't', 'e', 's', 't', '/', 's', 'r', 'c', '/', 'm', 'a', 't', 'h', '/', 's', 'p', 'e', 'c', 'i', 'a', 'l', '/', 'f', 'a', 'b', 's', '.', 'h', '\x00'})), int32(1), int32(0), float64(0), float64(0), float32(0), int32(0)}, common.Struct_l_l{(*int8)(unsafe.Pointer(&[55]int8{'/', 'U', 's', 'e', 'r', 's', '/', 'x', 'u', 's', 'h', 'i', 'w', 'e', 'i', '/', 'w', 'o', 'r', 'k', '/', 'l', 'i', 'b', 'c', '/', 't', 'e', 's', 't', '/', 's', 'r', 'c', '/', 'm', 'a', 't', 'h', '/', 's', 'p', 'e', 'c', 'i', 'a', 'l', '/', 'f', 'a', 'b', 's', '.', 'h', '\x00'})), int32(2), int32(0), float64(-0), float64(0), float32(0), int32(0)}, common.Struct_l_l{(*int8)(unsafe.Pointer(&[55]int8{'/', 'U', 's', 'e', 'r', 's', '/', 'x', 'u', 's', 'h', 'i', 'w', 'e', 'i', '/', 'w', 'o', 'r', 'k', '/', 'l', 'i', 'b', 'c', '/', 't', 'e', 's', 't', '/', 's', 'r', 'c', '/', 'm', 'a', 't', 'h', '/', 's', 'p', 'e', 'c', 'i', 'a', 'l', '/', 'f', 'a', 'b', 's', '.', 'h', '\x00'})), int32(3), int32(0), float64(1), float64(1), float32(0), int32(0)}, common.Struct_l_l{(*int8)(unsafe.Pointer(&[55]int8{'/', 'U', 's', 'e', 'r', 's', '/', 'x', 'u', 's', 'h', 'i', 'w', 'e', 'i', '/', 'w', 'o', 'r', 'k', '/', 'l', 'i', 'b', 'c', '/', 't', 'e', 's', 't', '/', 's', 'r', 'c', '/', 'm', 'a', 't', 'h', '/', 's', 'p', 'e', 'c', 'i', 'a', 'l', '/', 'f', 'a', 'b', 's', '.', 'h', '\x00'})), int32(4), int32(0), float64(-1), float64(1), float32(0), int32(0)}, common.Struct_l_l{(*int8)(unsafe.Pointer(&[55]int8{'/', 'U', 's', 'e', 'r', 's', '/', 'x', 'u', 's', 'h', 'i', 'w', 'e', 'i', '/', 'w', 'o', 'r', 'k', '/', 'l', 'i', 'b', 'c', '/', 't', 'e', 's', 't', '/', 's', 'r', 'c', '/', 'm', 'a', 't', 'h', '/', 's', 'p', 'e', 'c', 'i', 'a', 'l', '/', 'f', 'a', 'b', 's', '.', 'h', '\x00'})), int32(5), int32(0), float64(libc.X__builtin_inff()), float64(libc.X__builtin_inff()), float32(0), int32(0)}, common.Struct_l_l{(*int8)(unsafe.Pointer(&[55]int8{'/', 'U', 's', 'e', 'r', 's', '/', 'x', 'u', 's', 'h', 'i', 'w', 'e', 'i', '/', 'w', 'o', 'r', 'k', '/', 'l', 'i', 'b', 'c', '/', 't', 'e', 's', 't', '/', 's', 'r', 'c', '/', 'm', 'a', 't', 'h', '/', 's', 'p', 'e', 'c', 'i', 'a', 'l', '/', 'f', 'a', 'b', 's', '.', 'h', '\x00'})), int32(6), int32(0), float64(-libc.X__builtin_inff()), float64(libc.X__builtin_inff()), float32(0), int32(0)}, common.Struct_l_l{(*int8)(unsafe.Pointer(&[55]int8{'/', 'U', 's', 'e', 'r', 's', '/', 'x', 'u', 's', 'h', 'i', 'w', 'e', 'i', '/', 'w', 'o', 'r', 'k', '/', 'l', 'i', 'b', 'c', '/', 't', 'e', 's', 't', '/', 's', 'r', 'c', '/', 'm', 'a', 't', 'h', '/', 's', 'p', 'e', 'c', 'i', 'a', 'l', '/', 'f', 'a', 'b', 's', '.', 'h', '\x00'})), int32(7), int32(0), float64(libc.X__builtin_nanf((*int8)(unsafe.Pointer(&[1]int8{'\x00'})))), float64(libc.X__builtin_nanf((*int8)(unsafe.Pointer(&[1]int8{'\x00'})))), float32(0), int32(0)}}

func _cgo_main() int32 {
	var y float64
	var d float32
	var e int32
	var i int32
	var err int32 = int32(0)
	var p *common.Struct_l_l
	for i = int32(0); uint64(i) < 85; i++ {
		p = (*common.Struct_l_l)(unsafe.Pointer(uintptr(unsafe.Pointer((*common.Struct_l_l)(unsafe.Pointer(&_cgos_t_fabsl)))) + uintptr(i)*40))
		if p.R < int32(0) {
			continue
		}
		libc.Fesetround(p.R)
		libc.Feclearexcept(int32(0))
		y = libc.Fabsl(p.X)
		e = libc.Fetestexcept(0)
		if !(common.Checkexceptall(e, p.E, p.R) != 0) {
			libc.Printf((*int8)(unsafe.Pointer(&[52]int8{'%', 's', ':', '%', 'd', ':', ' ', 'b', 'a', 'd', ' ', 'f', 'p', ' ', 'e', 'x', 'c', 'e', 'p', 't', 'i', 'o', 'n', ':', ' ', '%', 's', ' ', 'f', 'a', 'b', 's', 'l', '(', '%', 'L', 'a', ')', '=', '%', 'L', 'a', ',', ' ', 'w', 'a', 'n', 't', ' ', '%', 's', '\x00'})), p.File, p.Line, common.Rstr(p.R), p.X, p.Y, common.Estr(p.E))
			libc.Printf((*int8)(unsafe.Pointer(&[9]int8{' ', 'g', 'o', 't', ' ', '%', 's', '\n', '\x00'})), common.Estr(e))
			err++
		}
		d = common.Ulperrl(y, p.Y, p.Dy)
		if !(common.Checkcr(y, p.Y, p.R) != 0) {
			libc.Printf((*int8)(unsafe.Pointer(&[61]int8{'%', 's', ':', '%', 'd', ':', ' ', '%', 's', ' ', 'f', 'a', 'b', 's', 'l', '(', '%', 'L', 'a', ')', ' ', 'w', 'a', 'n', 't', ' ', '%', 'L', 'a', ' ', 'g', 'o', 't', ' ', '%', 'L', 'a', ' ', 'u', 'l', 'p', 'e', 'r', 'r', ' ', '%', '.', '3', 'f', ' ', '=', ' ', '%', 'a', ' ', '+', ' ', '%', 'a', '\n', '\x00'})), p.File, p.Line, common.Rstr(p.R), p.X, p.Y, y, float64(d), float64(d-p.Dy), float64(p.Dy))
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