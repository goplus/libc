package fmax

import (
	common "github.com/goplus/libc/test/common"
	unsafe "unsafe"
	libc "github.com/goplus/libc"
	testing "testing"
)

var _cgos_t_fmax [68]common.Struct_dd_d = [68]common.Struct_dd_d{common.Struct_dd_d{(*int8)(unsafe.Pointer(&[54]int8{'/', 'U', 's', 'e', 'r', 's', '/', 'x', 'u', 's', 'h', 'i', 'w', 'e', 'i', '/', 'w', 'o', 'r', 'k', '/', 'l', 'i', 'b', 'c', '/', 't', 'e', 's', 't', '/', 's', 'r', 'c', '/', 'm', 'a', 't', 'h', '/', 's', 'a', 'n', 'i', 't', 'y', '/', 'f', 'm', 'a', 'x', '.', 'h', '\x00'})), int32(1), int32(0), -8.0668483905796808, 4.5356625606768688, 4.5356625606768688, float32(0), int32(0)}, common.Struct_dd_d{(*int8)(unsafe.Pointer(&[54]int8{'/', 'U', 's', 'e', 'r', 's', '/', 'x', 'u', 's', 'h', 'i', 'w', 'e', 'i', '/', 'w', 'o', 'r', 'k', '/', 'l', 'i', 'b', 'c', '/', 't', 'e', 's', 't', '/', 's', 'r', 'c', '/', 'm', 'a', 't', 'h', '/', 's', 'a', 'n', 'i', 't', 'y', '/', 'f', 'm', 'a', 'x', '.', 'h', '\x00'})), int32(2), int32(0), 4.3452398493383049, -8.8879913630034508, 4.3452398493383049, float32(0), int32(0)}, common.Struct_dd_d{(*int8)(unsafe.Pointer(&[54]int8{'/', 'U', 's', 'e', 'r', 's', '/', 'x', 'u', 's', 'h', 'i', 'w', 'e', 'i', '/', 'w', 'o', 'r', 'k', '/', 'l', 'i', 'b', 'c', '/', 't', 'e', 's', 't', '/', 's', 'r', 'c', '/', 'm', 'a', 't', 'h', '/', 's', 'a', 'n', 'i', 't', 'y', '/', 'f', 'm', 'a', 'x', '.', 'h', '\x00'})), int32(3), int32(0), -8.3814334275552493, -2.7636073373795882, -2.7636073373795882, float32(0), int32(0)}, common.Struct_dd_d{(*int8)(unsafe.Pointer(&[54]int8{'/', 'U', 's', 'e', 'r', 's', '/', 'x', 'u', 's', 'h', 'i', 'w', 'e', 'i', '/', 'w', 'o', 'r', 'k', '/', 'l', 'i', 'b', 'c', '/', 't', 'e', 's', 't', '/', 's', 'r', 'c', '/', 'm', 'a', 't', 'h', '/', 's', 'a', 'n', 'i', 't', 'y', '/', 'f', 'm', 'a', 'x', '.', 'h', '\x00'})), int32(4), int32(0), -6.5316735819134841, 4.5675352768427437, 4.5675352768427437, float32(0), int32(0)}, common.Struct_dd_d{(*int8)(unsafe.Pointer(&[54]int8{'/', 'U', 's', 'e', 'r', 's', '/', 'x', 'u', 's', 'h', 'i', 'w', 'e', 'i', '/', 'w', 'o', 'r', 'k', '/', 'l', 'i', 'b', 'c', '/', 't', 'e', 's', 't', '/', 's', 'r', 'c', '/', 'm', 'a', 't', 'h', '/', 's', 'a', 'n', 'i', 't', 'y', '/', 'f', 'm', 'a', 'x', '.', 'h', '\x00'})), int32(5), int32(0), 9.2670569669725857, 4.8113920843597962, 9.2670569669725857, float32(0), int32(0)}, common.Struct_dd_d{(*int8)(unsafe.Pointer(&[54]int8{'/', 'U', 's', 'e', 'r', 's', '/', 'x', 'u', 's', 'h', 'i', 'w', 'e', 'i', '/', 'w', 'o', 'r', 'k', '/', 'l', 'i', 'b', 'c', '/', 't', 'e', 's', 't', '/', 's', 'r', 'c', '/', 'm', 'a', 't', 'h', '/', 's', 'a', 'n', 'i', 't', 'y', '/', 'f', 'm', 'a', 'x', '.', 'h', '\x00'})), int32(6), int32(0), -6.4500455560602363, 0.66207179233767388, 0.66207179233767388, float32(0), int32(0)}, common.Struct_dd_d{(*int8)(unsafe.Pointer(&[54]int8{'/', 'U', 's', 'e', 'r', 's', '/', 'x', 'u', 's', 'h', 'i', 'w', 'e', 'i', '/', 'w', 'o', 'r', 'k', '/', 'l', 'i', 'b', 'c', '/', 't', 'e', 's', 't', '/', 's', 'r', 'c', '/', 'm', 'a', 't', 'h', '/', 's', 'a', 'n', 'i', 't', 'y', '/', 'f', 'm', 'a', 'x', '.', 'h', '\x00'})), int32(7), int32(0), 7.8588902530416966, 0.052154526750062248, 7.8588902530416966, float32(0), int32(0)}, common.Struct_dd_d{(*int8)(unsafe.Pointer(&[54]int8{'/', 'U', 's', 'e', 'r', 's', '/', 'x', 'u', 's', 'h', 'i', 'w', 'e', 'i', '/', 'w', 'o', 'r', 'k', '/', 'l', 'i', 'b', 'c', '/', 't', 'e', 's', 't', '/', 's', 'r', 'c', '/', 'm', 'a', 't', 'h', '/', 's', 'a', 'n', 'i', 't', 'y', '/', 'f', 'm', 'a', 'x', '.', 'h', '\x00'})), int32(8), int32(0), -0.79205451198489596, 7.6764026851175399, 7.6764026851175399, float32(0), int32(0)}, common.Struct_dd_d{(*int8)(unsafe.Pointer(&[54]int8{'/', 'U', 's', 'e', 'r', 's', '/', 'x', 'u', 's', 'h', 'i', 'w', 'e', 'i', '/', 'w', 'o', 'r', 'k', '/', 'l', 'i', 'b', 'c', '/', 't', 'e', 's', 't', '/', 's', 'r', 'c', '/', 'm', 'a', 't', 'h', '/', 's', 'a', 'n', 'i', 't', 'y', '/', 'f', 'm', 'a', 'x', '.', 'h', '\x00'})), int32(9), int32(0), 0.61570267319792404, 2.0119025790324803, 2.0119025790324803, float32(0), int32(0)}, common.Struct_dd_d{(*int8)(unsafe.Pointer(&[54]int8{'/', 'U', 's', 'e', 'r', 's', '/', 'x', 'u', 's', 'h', 'i', 'w', 'e', 'i', '/', 'w', 'o', 'r', 'k', '/', 'l', 'i', 'b', 'c', '/', 't', 'e', 's', 't', '/', 's', 'r', 'c', '/', 'm', 'a', 't', 'h', '/', 's', 'a', 'n', 'i', 't', 'y', '/', 'f', 'm', 'a', 'x', '.', 'h', '\x00'})), int32(10), int32(0), -0.55875868236091519, 0.032239830602638041, 0.032239830602638041, float32(0), int32(0)}, common.Struct_dd_d{(*int8)(unsafe.Pointer(&[55]int8{'/', 'U', 's', 'e', 'r', 's', '/', 'x', 'u', 's', 'h', 'i', 'w', 'e', 'i', '/', 'w', 'o', 'r', 'k', '/', 'l', 'i', 'b', 'c', '/', 't', 'e', 's', 't', '/', 's', 'r', 'c', '/', 'm', 'a', 't', 'h', '/', 's', 'p', 'e', 'c', 'i', 'a', 'l', '/', 'f', 'm', 'a', 'x', '.', 'h', '\x00'})), int32(1), int32(0), 0, 1, 1, float32(0), int32(0)}, common.Struct_dd_d{(*int8)(unsafe.Pointer(&[55]int8{'/', 'U', 's', 'e', 'r', 's', '/', 'x', 'u', 's', 'h', 'i', 'w', 'e', 'i', '/', 'w', 'o', 'r', 'k', '/', 'l', 'i', 'b', 'c', '/', 't', 'e', 's', 't', '/', 's', 'r', 'c', '/', 'm', 'a', 't', 'h', '/', 's', 'p', 'e', 'c', 'i', 'a', 'l', '/', 'f', 'm', 'a', 'x', '.', 'h', '\x00'})), int32(2), int32(0), -0, 1, 1, float32(0), int32(0)}, common.Struct_dd_d{(*int8)(unsafe.Pointer(&[55]int8{'/', 'U', 's', 'e', 'r', 's', '/', 'x', 'u', 's', 'h', 'i', 'w', 'e', 'i', '/', 'w', 'o', 'r', 'k', '/', 'l', 'i', 'b', 'c', '/', 't', 'e', 's', 't', '/', 's', 'r', 'c', '/', 'm', 'a', 't', 'h', '/', 's', 'p', 'e', 'c', 'i', 'a', 'l', '/', 'f', 'm', 'a', 'x', '.', 'h', '\x00'})), int32(3), int32(0), 0.5, 1, 1, float32(0), int32(0)}, common.Struct_dd_d{(*int8)(unsafe.Pointer(&[55]int8{'/', 'U', 's', 'e', 'r', 's', '/', 'x', 'u', 's', 'h', 'i', 'w', 'e', 'i', '/', 'w', 'o', 'r', 'k', '/', 'l', 'i', 'b', 'c', '/', 't', 'e', 's', 't', '/', 's', 'r', 'c', '/', 'm', 'a', 't', 'h', '/', 's', 'p', 'e', 'c', 'i', 'a', 'l', '/', 'f', 'm', 'a', 'x', '.', 'h', '\x00'})), int32(4), int32(0), -0.5, 1, 1, float32(0), int32(0)}, common.Struct_dd_d{(*int8)(unsafe.Pointer(&[55]int8{'/', 'U', 's', 'e', 'r', 's', '/', 'x', 'u', 's', 'h', 'i', 'w', 'e', 'i', '/', 'w', 'o', 'r', 'k', '/', 'l', 'i', 'b', 'c', '/', 't', 'e', 's', 't', '/', 's', 'r', 'c', '/', 'm', 'a', 't', 'h', '/', 's', 'p', 'e', 'c', 'i', 'a', 'l', '/', 'f', 'm', 'a', 'x', '.', 'h', '\x00'})), int32(5), int32(0), 1, 1, 1, float32(0), int32(0)}, common.Struct_dd_d{(*int8)(unsafe.Pointer(&[55]int8{'/', 'U', 's', 'e', 'r', 's', '/', 'x', 'u', 's', 'h', 'i', 'w', 'e', 'i', '/', 'w', 'o', 'r', 'k', '/', 'l', 'i', 'b', 'c', '/', 't', 'e', 's', 't', '/', 's', 'r', 'c', '/', 'm', 'a', 't', 'h', '/', 's', 'p', 'e', 'c', 'i', 'a', 'l', '/', 'f', 'm', 'a', 'x', '.', 'h', '\x00'})), int32(6), int32(0), -1, 1, 1, float32(0), int32(0)}, common.Struct_dd_d{(*int8)(unsafe.Pointer(&[55]int8{'/', 'U', 's', 'e', 'r', 's', '/', 'x', 'u', 's', 'h', 'i', 'w', 'e', 'i', '/', 'w', 'o', 'r', 'k', '/', 'l', 'i', 'b', 'c', '/', 't', 'e', 's', 't', '/', 's', 'r', 'c', '/', 'm', 'a', 't', 'h', '/', 's', 'p', 'e', 'c', 'i', 'a', 'l', '/', 'f', 'm', 'a', 'x', '.', 'h', '\x00'})), int32(7), int32(0), float64(libc.X__builtin_inff()), 1, float64(libc.X__builtin_inff()), float32(0), int32(0)}, common.Struct_dd_d{(*int8)(unsafe.Pointer(&[55]int8{'/', 'U', 's', 'e', 'r', 's', '/', 'x', 'u', 's', 'h', 'i', 'w', 'e', 'i', '/', 'w', 'o', 'r', 'k', '/', 'l', 'i', 'b', 'c', '/', 't', 'e', 's', 't', '/', 's', 'r', 'c', '/', 'm', 'a', 't', 'h', '/', 's', 'p', 'e', 'c', 'i', 'a', 'l', '/', 'f', 'm', 'a', 'x', '.', 'h', '\x00'})), int32(8), int32(0), float64(-libc.X__builtin_inff()), 1, 1, float32(0), int32(0)}, common.Struct_dd_d{(*int8)(unsafe.Pointer(&[55]int8{'/', 'U', 's', 'e', 'r', 's', '/', 'x', 'u', 's', 'h', 'i', 'w', 'e', 'i', '/', 'w', 'o', 'r', 'k', '/', 'l', 'i', 'b', 'c', '/', 't', 'e', 's', 't', '/', 's', 'r', 'c', '/', 'm', 'a', 't', 'h', '/', 's', 'p', 'e', 'c', 'i', 'a', 'l', '/', 'f', 'm', 'a', 'x', '.', 'h', '\x00'})), int32(9), int32(0), float64(libc.X__builtin_nanf((*int8)(unsafe.Pointer(&[1]int8{'\x00'})))), 1, 1, float32(0), int32(0)}, common.Struct_dd_d{(*int8)(unsafe.Pointer(&[55]int8{'/', 'U', 's', 'e', 'r', 's', '/', 'x', 'u', 's', 'h', 'i', 'w', 'e', 'i', '/', 'w', 'o', 'r', 'k', '/', 'l', 'i', 'b', 'c', '/', 't', 'e', 's', 't', '/', 's', 'r', 'c', '/', 'm', 'a', 't', 'h', '/', 's', 'p', 'e', 'c', 'i', 'a', 'l', '/', 'f', 'm', 'a', 'x', '.', 'h', '\x00'})), int32(10), int32(0), 0, -1, 0, float32(0), int32(0)}, common.Struct_dd_d{(*int8)(unsafe.Pointer(&[55]int8{'/', 'U', 's', 'e', 'r', 's', '/', 'x', 'u', 's', 'h', 'i', 'w', 'e', 'i', '/', 'w', 'o', 'r', 'k', '/', 'l', 'i', 'b', 'c', '/', 't', 'e', 's', 't', '/', 's', 'r', 'c', '/', 'm', 'a', 't', 'h', '/', 's', 'p', 'e', 'c', 'i', 'a', 'l', '/', 'f', 'm', 'a', 'x', '.', 'h', '\x00'})), int32(11), int32(0), -0, -1, -0, float32(0), int32(0)}, common.Struct_dd_d{(*int8)(unsafe.Pointer(&[55]int8{'/', 'U', 's', 'e', 'r', 's', '/', 'x', 'u', 's', 'h', 'i', 'w', 'e', 'i', '/', 'w', 'o', 'r', 'k', '/', 'l', 'i', 'b', 'c', '/', 't', 'e', 's', 't', '/', 's', 'r', 'c', '/', 'm', 'a', 't', 'h', '/', 's', 'p', 'e', 'c', 'i', 'a', 'l', '/', 'f', 'm', 'a', 'x', '.', 'h', '\x00'})), int32(12), int32(0), 0.5, -1, 0.5, float32(0), int32(0)}, common.Struct_dd_d{(*int8)(unsafe.Pointer(&[55]int8{'/', 'U', 's', 'e', 'r', 's', '/', 'x', 'u', 's', 'h', 'i', 'w', 'e', 'i', '/', 'w', 'o', 'r', 'k', '/', 'l', 'i', 'b', 'c', '/', 't', 'e', 's', 't', '/', 's', 'r', 'c', '/', 'm', 'a', 't', 'h', '/', 's', 'p', 'e', 'c', 'i', 'a', 'l', '/', 'f', 'm', 'a', 'x', '.', 'h', '\x00'})), int32(13), int32(0), -0.5, -1, -0.5, float32(0), int32(0)}, common.Struct_dd_d{(*int8)(unsafe.Pointer(&[55]int8{'/', 'U', 's', 'e', 'r', 's', '/', 'x', 'u', 's', 'h', 'i', 'w', 'e', 'i', '/', 'w', 'o', 'r', 'k', '/', 'l', 'i', 'b', 'c', '/', 't', 'e', 's', 't', '/', 's', 'r', 'c', '/', 'm', 'a', 't', 'h', '/', 's', 'p', 'e', 'c', 'i', 'a', 'l', '/', 'f', 'm', 'a', 'x', '.', 'h', '\x00'})), int32(14), int32(0), 1, -1, 1, float32(0), int32(0)}, common.Struct_dd_d{(*int8)(unsafe.Pointer(&[55]int8{'/', 'U', 's', 'e', 'r', 's', '/', 'x', 'u', 's', 'h', 'i', 'w', 'e', 'i', '/', 'w', 'o', 'r', 'k', '/', 'l', 'i', 'b', 'c', '/', 't', 'e', 's', 't', '/', 's', 'r', 'c', '/', 'm', 'a', 't', 'h', '/', 's', 'p', 'e', 'c', 'i', 'a', 'l', '/', 'f', 'm', 'a', 'x', '.', 'h', '\x00'})), int32(15), int32(0), -1, -1, -1, float32(0), int32(0)}, common.Struct_dd_d{(*int8)(unsafe.Pointer(&[55]int8{'/', 'U', 's', 'e', 'r', 's', '/', 'x', 'u', 's', 'h', 'i', 'w', 'e', 'i', '/', 'w', 'o', 'r', 'k', '/', 'l', 'i', 'b', 'c', '/', 't', 'e', 's', 't', '/', 's', 'r', 'c', '/', 'm', 'a', 't', 'h', '/', 's', 'p', 'e', 'c', 'i', 'a', 'l', '/', 'f', 'm', 'a', 'x', '.', 'h', '\x00'})), int32(16), int32(0), float64(libc.X__builtin_inff()), -1, float64(libc.X__builtin_inff()), float32(0), int32(0)}, common.Struct_dd_d{(*int8)(unsafe.Pointer(&[55]int8{'/', 'U', 's', 'e', 'r', 's', '/', 'x', 'u', 's', 'h', 'i', 'w', 'e', 'i', '/', 'w', 'o', 'r', 'k', '/', 'l', 'i', 'b', 'c', '/', 't', 'e', 's', 't', '/', 's', 'r', 'c', '/', 'm', 'a', 't', 'h', '/', 's', 'p', 'e', 'c', 'i', 'a', 'l', '/', 'f', 'm', 'a', 'x', '.', 'h', '\x00'})), int32(17), int32(0), float64(-libc.X__builtin_inff()), -1, -1, float32(0), int32(0)}, common.Struct_dd_d{(*int8)(unsafe.Pointer(&[55]int8{'/', 'U', 's', 'e', 'r', 's', '/', 'x', 'u', 's', 'h', 'i', 'w', 'e', 'i', '/', 'w', 'o', 'r', 'k', '/', 'l', 'i', 'b', 'c', '/', 't', 'e', 's', 't', '/', 's', 'r', 'c', '/', 'm', 'a', 't', 'h', '/', 's', 'p', 'e', 'c', 'i', 'a', 'l', '/', 'f', 'm', 'a', 'x', '.', 'h', '\x00'})), int32(18), int32(0), float64(libc.X__builtin_nanf((*int8)(unsafe.Pointer(&[1]int8{'\x00'})))), -1, -1, float32(0), int32(0)}, common.Struct_dd_d{(*int8)(unsafe.Pointer(&[55]int8{'/', 'U', 's', 'e', 'r', 's', '/', 'x', 'u', 's', 'h', 'i', 'w', 'e', 'i', '/', 'w', 'o', 'r', 'k', '/', 'l', 'i', 'b', 'c', '/', 't', 'e', 's', 't', '/', 's', 'r', 'c', '/', 'm', 'a', 't', 'h', '/', 's', 'p', 'e', 'c', 'i', 'a', 'l', '/', 'f', 'm', 'a', 'x', '.', 'h', '\x00'})), int32(19), int32(0), 0, 0, 0, float32(0), int32(0)}, common.Struct_dd_d{(*int8)(unsafe.Pointer(&[55]int8{'/', 'U', 's', 'e', 'r', 's', '/', 'x', 'u', 's', 'h', 'i', 'w', 'e', 'i', '/', 'w', 'o', 'r', 'k', '/', 'l', 'i', 'b', 'c', '/', 't', 'e', 's', 't', '/', 's', 'r', 'c', '/', 'm', 'a', 't', 'h', '/', 's', 'p', 'e', 'c', 'i', 'a', 'l', '/', 'f', 'm', 'a', 'x', '.', 'h', '\x00'})), int32(20), int32(0), 0, -0, 0, float32(0), int32(0)}, common.Struct_dd_d{(*int8)(unsafe.Pointer(&[55]int8{'/', 'U', 's', 'e', 'r', 's', '/', 'x', 'u', 's', 'h', 'i', 'w', 'e', 'i', '/', 'w', 'o', 'r', 'k', '/', 'l', 'i', 'b', 'c', '/', 't', 'e', 's', 't', '/', 's', 'r', 'c', '/', 'm', 'a', 't', 'h', '/', 's', 'p', 'e', 'c', 'i', 'a', 'l', '/', 'f', 'm', 'a', 'x', '.', 'h', '\x00'})), int32(21), int32(0), 0, float64(libc.X__builtin_inff()), float64(libc.X__builtin_inff()), float32(0), int32(0)}, common.Struct_dd_d{(*int8)(unsafe.Pointer(&[55]int8{'/', 'U', 's', 'e', 'r', 's', '/', 'x', 'u', 's', 'h', 'i', 'w', 'e', 'i', '/', 'w', 'o', 'r', 'k', '/', 'l', 'i', 'b', 'c', '/', 't', 'e', 's', 't', '/', 's', 'r', 'c', '/', 'm', 'a', 't', 'h', '/', 's', 'p', 'e', 'c', 'i', 'a', 'l', '/', 'f', 'm', 'a', 'x', '.', 'h', '\x00'})), int32(22), int32(0), 0, float64(-libc.X__builtin_inff()), 0, float32(0), int32(0)}, common.Struct_dd_d{(*int8)(unsafe.Pointer(&[55]int8{'/', 'U', 's', 'e', 'r', 's', '/', 'x', 'u', 's', 'h', 'i', 'w', 'e', 'i', '/', 'w', 'o', 'r', 'k', '/', 'l', 'i', 'b', 'c', '/', 't', 'e', 's', 't', '/', 's', 'r', 'c', '/', 'm', 'a', 't', 'h', '/', 's', 'p', 'e', 'c', 'i', 'a', 'l', '/', 'f', 'm', 'a', 'x', '.', 'h', '\x00'})), int32(23), int32(0), 0, float64(libc.X__builtin_nanf((*int8)(unsafe.Pointer(&[1]int8{'\x00'})))), 0, float32(0), int32(0)}, common.Struct_dd_d{(*int8)(unsafe.Pointer(&[55]int8{'/', 'U', 's', 'e', 'r', 's', '/', 'x', 'u', 's', 'h', 'i', 'w', 'e', 'i', '/', 'w', 'o', 'r', 'k', '/', 'l', 'i', 'b', 'c', '/', 't', 'e', 's', 't', '/', 's', 'r', 'c', '/', 'm', 'a', 't', 'h', '/', 's', 'p', 'e', 'c', 'i', 'a', 'l', '/', 'f', 'm', 'a', 'x', '.', 'h', '\x00'})), int32(24), int32(0), -0, 0, 0, float32(0), int32(0)}, common.Struct_dd_d{(*int8)(unsafe.Pointer(&[55]int8{'/', 'U', 's', 'e', 'r', 's', '/', 'x', 'u', 's', 'h', 'i', 'w', 'e', 'i', '/', 'w', 'o', 'r', 'k', '/', 'l', 'i', 'b', 'c', '/', 't', 'e', 's', 't', '/', 's', 'r', 'c', '/', 'm', 'a', 't', 'h', '/', 's', 'p', 'e', 'c', 'i', 'a', 'l', '/', 'f', 'm', 'a', 'x', '.', 'h', '\x00'})), int32(25), int32(0), -0, -0, -0, float32(0), int32(0)}, common.Struct_dd_d{(*int8)(unsafe.Pointer(&[55]int8{'/', 'U', 's', 'e', 'r', 's', '/', 'x', 'u', 's', 'h', 'i', 'w', 'e', 'i', '/', 'w', 'o', 'r', 'k', '/', 'l', 'i', 'b', 'c', '/', 't', 'e', 's', 't', '/', 's', 'r', 'c', '/', 'm', 'a', 't', 'h', '/', 's', 'p', 'e', 'c', 'i', 'a', 'l', '/', 'f', 'm', 'a', 'x', '.', 'h', '\x00'})), int32(26), int32(0), -0, float64(libc.X__builtin_inff()), float64(libc.X__builtin_inff()), float32(0), int32(0)}, common.Struct_dd_d{(*int8)(unsafe.Pointer(&[55]int8{'/', 'U', 's', 'e', 'r', 's', '/', 'x', 'u', 's', 'h', 'i', 'w', 'e', 'i', '/', 'w', 'o', 'r', 'k', '/', 'l', 'i', 'b', 'c', '/', 't', 'e', 's', 't', '/', 's', 'r', 'c', '/', 'm', 'a', 't', 'h', '/', 's', 'p', 'e', 'c', 'i', 'a', 'l', '/', 'f', 'm', 'a', 'x', '.', 'h', '\x00'})), int32(27), int32(0), -0, float64(-libc.X__builtin_inff()), -0, float32(0), int32(0)}, common.Struct_dd_d{(*int8)(unsafe.Pointer(&[55]int8{'/', 'U', 's', 'e', 'r', 's', '/', 'x', 'u', 's', 'h', 'i', 'w', 'e', 'i', '/', 'w', 'o', 'r', 'k', '/', 'l', 'i', 'b', 'c', '/', 't', 'e', 's', 't', '/', 's', 'r', 'c', '/', 'm', 'a', 't', 'h', '/', 's', 'p', 'e', 'c', 'i', 'a', 'l', '/', 'f', 'm', 'a', 'x', '.', 'h', '\x00'})), int32(28), int32(0), -0, float64(libc.X__builtin_nanf((*int8)(unsafe.Pointer(&[1]int8{'\x00'})))), -0, float32(0), int32(0)}, common.Struct_dd_d{(*int8)(unsafe.Pointer(&[55]int8{'/', 'U', 's', 'e', 'r', 's', '/', 'x', 'u', 's', 'h', 'i', 'w', 'e', 'i', '/', 'w', 'o', 'r', 'k', '/', 'l', 'i', 'b', 'c', '/', 't', 'e', 's', 't', '/', 's', 'r', 'c', '/', 'm', 'a', 't', 'h', '/', 's', 'p', 'e', 'c', 'i', 'a', 'l', '/', 'f', 'm', 'a', 'x', '.', 'h', '\x00'})), int32(29), int32(0), 1, 0, 1, float32(0), int32(0)}, common.Struct_dd_d{(*int8)(unsafe.Pointer(&[55]int8{'/', 'U', 's', 'e', 'r', 's', '/', 'x', 'u', 's', 'h', 'i', 'w', 'e', 'i', '/', 'w', 'o', 'r', 'k', '/', 'l', 'i', 'b', 'c', '/', 't', 'e', 's', 't', '/', 's', 'r', 'c', '/', 'm', 'a', 't', 'h', '/', 's', 'p', 'e', 'c', 'i', 'a', 'l', '/', 'f', 'm', 'a', 'x', '.', 'h', '\x00'})), int32(30), int32(0), -1, 0, 0, float32(0), int32(0)}, common.Struct_dd_d{(*int8)(unsafe.Pointer(&[55]int8{'/', 'U', 's', 'e', 'r', 's', '/', 'x', 'u', 's', 'h', 'i', 'w', 'e', 'i', '/', 'w', 'o', 'r', 'k', '/', 'l', 'i', 'b', 'c', '/', 't', 'e', 's', 't', '/', 's', 'r', 'c', '/', 'm', 'a', 't', 'h', '/', 's', 'p', 'e', 'c', 'i', 'a', 'l', '/', 'f', 'm', 'a', 'x', '.', 'h', '\x00'})), int32(31), int32(0), float64(libc.X__builtin_inff()), 0, float64(libc.X__builtin_inff()), float32(0), int32(0)}, common.Struct_dd_d{(*int8)(unsafe.Pointer(&[55]int8{'/', 'U', 's', 'e', 'r', 's', '/', 'x', 'u', 's', 'h', 'i', 'w', 'e', 'i', '/', 'w', 'o', 'r', 'k', '/', 'l', 'i', 'b', 'c', '/', 't', 'e', 's', 't', '/', 's', 'r', 'c', '/', 'm', 'a', 't', 'h', '/', 's', 'p', 'e', 'c', 'i', 'a', 'l', '/', 'f', 'm', 'a', 'x', '.', 'h', '\x00'})), int32(32), int32(0), float64(-libc.X__builtin_inff()), 0, 0, float32(0), int32(0)}, common.Struct_dd_d{(*int8)(unsafe.Pointer(&[55]int8{'/', 'U', 's', 'e', 'r', 's', '/', 'x', 'u', 's', 'h', 'i', 'w', 'e', 'i', '/', 'w', 'o', 'r', 'k', '/', 'l', 'i', 'b', 'c', '/', 't', 'e', 's', 't', '/', 's', 'r', 'c', '/', 'm', 'a', 't', 'h', '/', 's', 'p', 'e', 'c', 'i', 'a', 'l', '/', 'f', 'm', 'a', 'x', '.', 'h', '\x00'})), int32(33), int32(0), float64(libc.X__builtin_nanf((*int8)(unsafe.Pointer(&[1]int8{'\x00'})))), 0, 0, float32(0), int32(0)}, common.Struct_dd_d{(*int8)(unsafe.Pointer(&[55]int8{'/', 'U', 's', 'e', 'r', 's', '/', 'x', 'u', 's', 'h', 'i', 'w', 'e', 'i', '/', 'w', 'o', 'r', 'k', '/', 'l', 'i', 'b', 'c', '/', 't', 'e', 's', 't', '/', 's', 'r', 'c', '/', 'm', 'a', 't', 'h', '/', 's', 'p', 'e', 'c', 'i', 'a', 'l', '/', 'f', 'm', 'a', 'x', '.', 'h', '\x00'})), int32(34), int32(0), -1, -0, -0, float32(0), int32(0)}, common.Struct_dd_d{(*int8)(unsafe.Pointer(&[55]int8{'/', 'U', 's', 'e', 'r', 's', '/', 'x', 'u', 's', 'h', 'i', 'w', 'e', 'i', '/', 'w', 'o', 'r', 'k', '/', 'l', 'i', 'b', 'c', '/', 't', 'e', 's', 't', '/', 's', 'r', 'c', '/', 'm', 'a', 't', 'h', '/', 's', 'p', 'e', 'c', 'i', 'a', 'l', '/', 'f', 'm', 'a', 'x', '.', 'h', '\x00'})), int32(35), int32(0), float64(libc.X__builtin_inff()), -0, float64(libc.X__builtin_inff()), float32(0), int32(0)}, common.Struct_dd_d{(*int8)(unsafe.Pointer(&[55]int8{'/', 'U', 's', 'e', 'r', 's', '/', 'x', 'u', 's', 'h', 'i', 'w', 'e', 'i', '/', 'w', 'o', 'r', 'k', '/', 'l', 'i', 'b', 'c', '/', 't', 'e', 's', 't', '/', 's', 'r', 'c', '/', 'm', 'a', 't', 'h', '/', 's', 'p', 'e', 'c', 'i', 'a', 'l', '/', 'f', 'm', 'a', 'x', '.', 'h', '\x00'})), int32(36), int32(0), float64(-libc.X__builtin_inff()), -0, -0, float32(0), int32(0)}, common.Struct_dd_d{(*int8)(unsafe.Pointer(&[55]int8{'/', 'U', 's', 'e', 'r', 's', '/', 'x', 'u', 's', 'h', 'i', 'w', 'e', 'i', '/', 'w', 'o', 'r', 'k', '/', 'l', 'i', 'b', 'c', '/', 't', 'e', 's', 't', '/', 's', 'r', 'c', '/', 'm', 'a', 't', 'h', '/', 's', 'p', 'e', 'c', 'i', 'a', 'l', '/', 'f', 'm', 'a', 'x', '.', 'h', '\x00'})), int32(37), int32(0), float64(libc.X__builtin_nanf((*int8)(unsafe.Pointer(&[1]int8{'\x00'})))), -0, -0, float32(0), int32(0)}, common.Struct_dd_d{(*int8)(unsafe.Pointer(&[55]int8{'/', 'U', 's', 'e', 'r', 's', '/', 'x', 'u', 's', 'h', 'i', 'w', 'e', 'i', '/', 'w', 'o', 'r', 'k', '/', 'l', 'i', 'b', 'c', '/', 't', 'e', 's', 't', '/', 's', 'r', 'c', '/', 'm', 'a', 't', 'h', '/', 's', 'p', 'e', 'c', 'i', 'a', 'l', '/', 'f', 'm', 'a', 'x', '.', 'h', '\x00'})), int32(38), int32(0), float64(libc.X__builtin_inff()), 2, float64(libc.X__builtin_inff()), float32(0), int32(0)}, common.Struct_dd_d{(*int8)(unsafe.Pointer(&[55]int8{'/', 'U', 's', 'e', 'r', 's', '/', 'x', 'u', 's', 'h', 'i', 'w', 'e', 'i', '/', 'w', 'o', 'r', 'k', '/', 'l', 'i', 'b', 'c', '/', 't', 'e', 's', 't', '/', 's', 'r', 'c', '/', 'm', 'a', 't', 'h', '/', 's', 'p', 'e', 'c', 'i', 'a', 'l', '/', 'f', 'm', 'a', 'x', '.', 'h', '\x00'})), int32(39), int32(0), float64(libc.X__builtin_inff()), -0.5, float64(libc.X__builtin_inff()), float32(0), int32(0)}, common.Struct_dd_d{(*int8)(unsafe.Pointer(&[55]int8{'/', 'U', 's', 'e', 'r', 's', '/', 'x', 'u', 's', 'h', 'i', 'w', 'e', 'i', '/', 'w', 'o', 'r', 'k', '/', 'l', 'i', 'b', 'c', '/', 't', 'e', 's', 't', '/', 's', 'r', 'c', '/', 'm', 'a', 't', 'h', '/', 's', 'p', 'e', 'c', 'i', 'a', 'l', '/', 'f', 'm', 'a', 'x', '.', 'h', '\x00'})), int32(40), int32(0), float64(libc.X__builtin_inff()), float64(libc.X__builtin_nanf((*int8)(unsafe.Pointer(&[1]int8{'\x00'})))), float64(libc.X__builtin_inff()), float32(0), int32(0)}, common.Struct_dd_d{(*int8)(unsafe.Pointer(&[55]int8{'/', 'U', 's', 'e', 'r', 's', '/', 'x', 'u', 's', 'h', 'i', 'w', 'e', 'i', '/', 'w', 'o', 'r', 'k', '/', 'l', 'i', 'b', 'c', '/', 't', 'e', 's', 't', '/', 's', 'r', 'c', '/', 'm', 'a', 't', 'h', '/', 's', 'p', 'e', 'c', 'i', 'a', 'l', '/', 'f', 'm', 'a', 'x', '.', 'h', '\x00'})), int32(41), int32(0), float64(-libc.X__builtin_inff()), 2, 2, float32(0), int32(0)}, common.Struct_dd_d{(*int8)(unsafe.Pointer(&[55]int8{'/', 'U', 's', 'e', 'r', 's', '/', 'x', 'u', 's', 'h', 'i', 'w', 'e', 'i', '/', 'w', 'o', 'r', 'k', '/', 'l', 'i', 'b', 'c', '/', 't', 'e', 's', 't', '/', 's', 'r', 'c', '/', 'm', 'a', 't', 'h', '/', 's', 'p', 'e', 'c', 'i', 'a', 'l', '/', 'f', 'm', 'a', 'x', '.', 'h', '\x00'})), int32(42), int32(0), float64(-libc.X__builtin_inff()), -0.5, -0.5, float32(0), int32(0)}, common.Struct_dd_d{(*int8)(unsafe.Pointer(&[55]int8{'/', 'U', 's', 'e', 'r', 's', '/', 'x', 'u', 's', 'h', 'i', 'w', 'e', 'i', '/', 'w', 'o', 'r', 'k', '/', 'l', 'i', 'b', 'c', '/', 't', 'e', 's', 't', '/', 's', 'r', 'c', '/', 'm', 'a', 't', 'h', '/', 's', 'p', 'e', 'c', 'i', 'a', 'l', '/', 'f', 'm', 'a', 'x', '.', 'h', '\x00'})), int32(43), int32(0), float64(-libc.X__builtin_inff()), float64(libc.X__builtin_nanf((*int8)(unsafe.Pointer(&[1]int8{'\x00'})))), float64(-libc.X__builtin_inff()), float32(0), int32(0)}, common.Struct_dd_d{(*int8)(unsafe.Pointer(&[55]int8{'/', 'U', 's', 'e', 'r', 's', '/', 'x', 'u', 's', 'h', 'i', 'w', 'e', 'i', '/', 'w', 'o', 'r', 'k', '/', 'l', 'i', 'b', 'c', '/', 't', 'e', 's', 't', '/', 's', 'r', 'c', '/', 'm', 'a', 't', 'h', '/', 's', 'p', 'e', 'c', 'i', 'a', 'l', '/', 'f', 'm', 'a', 'x', '.', 'h', '\x00'})), int32(44), int32(0), float64(libc.X__builtin_nanf((*int8)(unsafe.Pointer(&[1]int8{'\x00'})))), float64(libc.X__builtin_nanf((*int8)(unsafe.Pointer(&[1]int8{'\x00'})))), float64(libc.X__builtin_nanf((*int8)(unsafe.Pointer(&[1]int8{'\x00'})))), float32(0), int32(0)}, common.Struct_dd_d{(*int8)(unsafe.Pointer(&[55]int8{'/', 'U', 's', 'e', 'r', 's', '/', 'x', 'u', 's', 'h', 'i', 'w', 'e', 'i', '/', 'w', 'o', 'r', 'k', '/', 'l', 'i', 'b', 'c', '/', 't', 'e', 's', 't', '/', 's', 'r', 'c', '/', 'm', 'a', 't', 'h', '/', 's', 'p', 'e', 'c', 'i', 'a', 'l', '/', 'f', 'm', 'a', 'x', '.', 'h', '\x00'})), int32(45), int32(0), 1, float64(libc.X__builtin_nanf((*int8)(unsafe.Pointer(&[1]int8{'\x00'})))), 1, float32(0), int32(0)}, common.Struct_dd_d{(*int8)(unsafe.Pointer(&[55]int8{'/', 'U', 's', 'e', 'r', 's', '/', 'x', 'u', 's', 'h', 'i', 'w', 'e', 'i', '/', 'w', 'o', 'r', 'k', '/', 'l', 'i', 'b', 'c', '/', 't', 'e', 's', 't', '/', 's', 'r', 'c', '/', 'm', 'a', 't', 'h', '/', 's', 'p', 'e', 'c', 'i', 'a', 'l', '/', 'f', 'm', 'a', 'x', '.', 'h', '\x00'})), int32(46), int32(0), -1, float64(libc.X__builtin_nanf((*int8)(unsafe.Pointer(&[1]int8{'\x00'})))), -1, float32(0), int32(0)}, common.Struct_dd_d{(*int8)(unsafe.Pointer(&[55]int8{'/', 'U', 's', 'e', 'r', 's', '/', 'x', 'u', 's', 'h', 'i', 'w', 'e', 'i', '/', 'w', 'o', 'r', 'k', '/', 'l', 'i', 'b', 'c', '/', 't', 'e', 's', 't', '/', 's', 'r', 'c', '/', 'm', 'a', 't', 'h', '/', 's', 'p', 'e', 'c', 'i', 'a', 'l', '/', 'f', 'm', 'a', 'x', '.', 'h', '\x00'})), int32(47), int32(0), 1, float64(libc.X__builtin_inff()), float64(libc.X__builtin_inff()), float32(0), int32(0)}, common.Struct_dd_d{(*int8)(unsafe.Pointer(&[55]int8{'/', 'U', 's', 'e', 'r', 's', '/', 'x', 'u', 's', 'h', 'i', 'w', 'e', 'i', '/', 'w', 'o', 'r', 'k', '/', 'l', 'i', 'b', 'c', '/', 't', 'e', 's', 't', '/', 's', 'r', 'c', '/', 'm', 'a', 't', 'h', '/', 's', 'p', 'e', 'c', 'i', 'a', 'l', '/', 'f', 'm', 'a', 'x', '.', 'h', '\x00'})), int32(48), int32(0), -1, float64(libc.X__builtin_inff()), float64(libc.X__builtin_inff()), float32(0), int32(0)}, common.Struct_dd_d{(*int8)(unsafe.Pointer(&[55]int8{'/', 'U', 's', 'e', 'r', 's', '/', 'x', 'u', 's', 'h', 'i', 'w', 'e', 'i', '/', 'w', 'o', 'r', 'k', '/', 'l', 'i', 'b', 'c', '/', 't', 'e', 's', 't', '/', 's', 'r', 'c', '/', 'm', 'a', 't', 'h', '/', 's', 'p', 'e', 'c', 'i', 'a', 'l', '/', 'f', 'm', 'a', 'x', '.', 'h', '\x00'})), int32(49), int32(0), float64(libc.X__builtin_inff()), float64(libc.X__builtin_inff()), float64(libc.X__builtin_inff()), float32(0), int32(0)}, common.Struct_dd_d{(*int8)(unsafe.Pointer(&[55]int8{'/', 'U', 's', 'e', 'r', 's', '/', 'x', 'u', 's', 'h', 'i', 'w', 'e', 'i', '/', 'w', 'o', 'r', 'k', '/', 'l', 'i', 'b', 'c', '/', 't', 'e', 's', 't', '/', 's', 'r', 'c', '/', 'm', 'a', 't', 'h', '/', 's', 'p', 'e', 'c', 'i', 'a', 'l', '/', 'f', 'm', 'a', 'x', '.', 'h', '\x00'})), int32(50), int32(0), float64(-libc.X__builtin_inff()), float64(libc.X__builtin_inff()), float64(libc.X__builtin_inff()), float32(0), int32(0)}, common.Struct_dd_d{(*int8)(unsafe.Pointer(&[55]int8{'/', 'U', 's', 'e', 'r', 's', '/', 'x', 'u', 's', 'h', 'i', 'w', 'e', 'i', '/', 'w', 'o', 'r', 'k', '/', 'l', 'i', 'b', 'c', '/', 't', 'e', 's', 't', '/', 's', 'r', 'c', '/', 'm', 'a', 't', 'h', '/', 's', 'p', 'e', 'c', 'i', 'a', 'l', '/', 'f', 'm', 'a', 'x', '.', 'h', '\x00'})), int32(51), int32(0), 1, float64(-libc.X__builtin_inff()), 1, float32(0), int32(0)}, common.Struct_dd_d{(*int8)(unsafe.Pointer(&[55]int8{'/', 'U', 's', 'e', 'r', 's', '/', 'x', 'u', 's', 'h', 'i', 'w', 'e', 'i', '/', 'w', 'o', 'r', 'k', '/', 'l', 'i', 'b', 'c', '/', 't', 'e', 's', 't', '/', 's', 'r', 'c', '/', 'm', 'a', 't', 'h', '/', 's', 'p', 'e', 'c', 'i', 'a', 'l', '/', 'f', 'm', 'a', 'x', '.', 'h', '\x00'})), int32(52), int32(0), -1, float64(-libc.X__builtin_inff()), -1, float32(0), int32(0)}, common.Struct_dd_d{(*int8)(unsafe.Pointer(&[55]int8{'/', 'U', 's', 'e', 'r', 's', '/', 'x', 'u', 's', 'h', 'i', 'w', 'e', 'i', '/', 'w', 'o', 'r', 'k', '/', 'l', 'i', 'b', 'c', '/', 't', 'e', 's', 't', '/', 's', 'r', 'c', '/', 'm', 'a', 't', 'h', '/', 's', 'p', 'e', 'c', 'i', 'a', 'l', '/', 'f', 'm', 'a', 'x', '.', 'h', '\x00'})), int32(53), int32(0), float64(libc.X__builtin_inff()), float64(-libc.X__builtin_inff()), float64(libc.X__builtin_inff()), float32(0), int32(0)}, common.Struct_dd_d{(*int8)(unsafe.Pointer(&[55]int8{'/', 'U', 's', 'e', 'r', 's', '/', 'x', 'u', 's', 'h', 'i', 'w', 'e', 'i', '/', 'w', 'o', 'r', 'k', '/', 'l', 'i', 'b', 'c', '/', 't', 'e', 's', 't', '/', 's', 'r', 'c', '/', 'm', 'a', 't', 'h', '/', 's', 'p', 'e', 'c', 'i', 'a', 'l', '/', 'f', 'm', 'a', 'x', '.', 'h', '\x00'})), int32(54), int32(0), float64(-libc.X__builtin_inff()), float64(-libc.X__builtin_inff()), float64(-libc.X__builtin_inff()), float32(0), int32(0)}, common.Struct_dd_d{(*int8)(unsafe.Pointer(&[55]int8{'/', 'U', 's', 'e', 'r', 's', '/', 'x', 'u', 's', 'h', 'i', 'w', 'e', 'i', '/', 'w', 'o', 'r', 'k', '/', 'l', 'i', 'b', 'c', '/', 't', 'e', 's', 't', '/', 's', 'r', 'c', '/', 'm', 'a', 't', 'h', '/', 's', 'p', 'e', 'c', 'i', 'a', 'l', '/', 'f', 'm', 'a', 'x', '.', 'h', '\x00'})), int32(55), int32(0), 1.75, 0.5, 1.75, float32(0), int32(0)}, common.Struct_dd_d{(*int8)(unsafe.Pointer(&[55]int8{'/', 'U', 's', 'e', 'r', 's', '/', 'x', 'u', 's', 'h', 'i', 'w', 'e', 'i', '/', 'w', 'o', 'r', 'k', '/', 'l', 'i', 'b', 'c', '/', 't', 'e', 's', 't', '/', 's', 'r', 'c', '/', 'm', 'a', 't', 'h', '/', 's', 'p', 'e', 'c', 'i', 'a', 'l', '/', 'f', 'm', 'a', 'x', '.', 'h', '\x00'})), int32(56), int32(0), -1.75, 0.5, 0.5, float32(0), int32(0)}, common.Struct_dd_d{(*int8)(unsafe.Pointer(&[55]int8{'/', 'U', 's', 'e', 'r', 's', '/', 'x', 'u', 's', 'h', 'i', 'w', 'e', 'i', '/', 'w', 'o', 'r', 'k', '/', 'l', 'i', 'b', 'c', '/', 't', 'e', 's', 't', '/', 's', 'r', 'c', '/', 'm', 'a', 't', 'h', '/', 's', 'p', 'e', 'c', 'i', 'a', 'l', '/', 'f', 'm', 'a', 'x', '.', 'h', '\x00'})), int32(57), int32(0), 1.75, -0.5, 1.75, float32(0), int32(0)}, common.Struct_dd_d{(*int8)(unsafe.Pointer(&[55]int8{'/', 'U', 's', 'e', 'r', 's', '/', 'x', 'u', 's', 'h', 'i', 'w', 'e', 'i', '/', 'w', 'o', 'r', 'k', '/', 'l', 'i', 'b', 'c', '/', 't', 'e', 's', 't', '/', 's', 'r', 'c', '/', 'm', 'a', 't', 'h', '/', 's', 'p', 'e', 'c', 'i', 'a', 'l', '/', 'f', 'm', 'a', 'x', '.', 'h', '\x00'})), int32(58), int32(0), -1.75, -0.5, -0.5, float32(0), int32(0)}}

func _cgo_main() int32 {
	var y float64
	var d float32
	var e int32
	var i int32
	var err int32 = int32(0)
	var p *common.Struct_dd_d
	for i = int32(0); uint64(i) < 68; i++ {
		p = (*common.Struct_dd_d)(unsafe.Pointer(uintptr(unsafe.Pointer((*common.Struct_dd_d)(unsafe.Pointer(&_cgos_t_fmax)))) + uintptr(i)*48))
		if p.R < int32(0) {
			continue
		}
		libc.Fesetround(p.R)
		libc.Feclearexcept(int32(0))
		y = libc.Fmax(p.X, p.X2)
		e = libc.Fetestexcept(0)
		if !(common.Checkexceptall(e, p.E, p.R) != 0) {
			libc.Printf((*int8)(unsafe.Pointer(&[52]int8{'%', 's', ':', '%', 'd', ':', ' ', 'b', 'a', 'd', ' ', 'f', 'p', ' ', 'e', 'x', 'c', 'e', 'p', 't', 'i', 'o', 'n', ':', ' ', '%', 's', ' ', 'f', 'm', 'a', 'x', '(', '%', 'a', ',', '%', 'a', ')', '=', '%', 'a', ',', ' ', 'w', 'a', 'n', 't', ' ', '%', 's', '\x00'})), p.File, p.Line, common.Rstr(p.R), p.X, p.X2, p.Y, common.Estr(p.E))
			libc.Printf((*int8)(unsafe.Pointer(&[9]int8{' ', 'g', 'o', 't', ' ', '%', 's', '\n', '\x00'})), common.Estr(e))
			err++
		}
		d = common.Ulperr(y, p.Y, p.Dy)
		if !(common.Checkcr(float64(y), float64(p.Y), p.R) != 0) {
			libc.Printf((*int8)(unsafe.Pointer(&[60]int8{'%', 's', ':', '%', 'd', ':', ' ', '%', 's', ' ', 'f', 'm', 'a', 'x', '(', '%', 'a', ',', '%', 'a', ')', ' ', 'w', 'a', 'n', 't', ' ', '%', 'a', ' ', 'g', 'o', 't', ' ', '%', 'a', ' ', 'u', 'l', 'p', 'e', 'r', 'r', ' ', '%', '.', '3', 'f', ' ', '=', ' ', '%', 'a', ' ', '+', ' ', '%', 'a', '\n', '\x00'})), p.File, p.Line, common.Rstr(p.R), p.X, p.X2, p.Y, y, float64(d), float64(d-p.Dy), float64(p.Dy))
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
