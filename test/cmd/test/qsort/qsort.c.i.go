package qsort

import (
	libc "github.com/goplus/libc"
	unsafe "unsafe"
	common "github.com/goplus/libc/test/common"
	testing "testing"
)

func _cgos_scmp_qsort(a unsafe.Pointer, b unsafe.Pointer) int32 {
	return libc.Strcmp(*(**int8)(a), *(**int8)(b))
}
func _cgos_icmp_qsort(a unsafe.Pointer, b unsafe.Pointer) int32 {
	return *(*int32)(a) - *(*int32)(b)
}
func _cgos_ccmp_qsort(a unsafe.Pointer, b unsafe.Pointer) int32 {
	return int32(*(*int8)(a)) - int32(*(*int8)(b))
}
func _cgos_cmp64_qsort(a unsafe.Pointer, b unsafe.Pointer) int32 {
	var ua *uint64 = (*uint64)(a)
	var ub *uint64 = (*uint64)(b)
	return func() int32 {
		if *ua < *ub {
			return -1
		} else {
			return func() int32 {
				if *ua != *ub {
					return 1
				} else {
					return 0
				}
			}()
		}
	}()
}

var _cgos_s_qsort [26]*int8 = [26]*int8{(*int8)(unsafe.Pointer(&[4]int8{'B', 'o', 'b', '\x00'})), (*int8)(unsafe.Pointer(&[6]int8{'A', 'l', 'i', 'c', 'e', '\x00'})), (*int8)(unsafe.Pointer(&[5]int8{'J', 'o', 'h', 'n', '\x00'})), (*int8)(unsafe.Pointer(&[6]int8{'C', 'e', 'r', 'e', 's', '\x00'})), (*int8)(unsafe.Pointer(&[6]int8{'H', 'e', 'l', 'g', 'a', '\x00'})), (*int8)(unsafe.Pointer(&[8]int8{'D', 'r', 'e', 'p', 'p', 'e', 'r', '\x00'})), (*int8)(unsafe.Pointer(&[9]int8{'E', 'm', 'e', 'r', 'a', 'l', 'd', 'a', '\x00'})), (*int8)(unsafe.Pointer(&[6]int8{'Z', 'o', 'r', 'a', 'n', '\x00'})), (*int8)(unsafe.Pointer(&[5]int8{'M', 'o', 'm', 'o', '\x00'})), (*int8)(unsafe.Pointer(&[6]int8{'F', 'r', 'a', 'n', 'k', '\x00'})), (*int8)(unsafe.Pointer(&[5]int8{'P', 'e', 'm', 'a', '\x00'})), (*int8)(unsafe.Pointer(&[7]int8{'X', 'a', 'v', 'i', 'e', 'r', '\x00'})), (*int8)(unsafe.Pointer(&[5]int8{'Y', 'e', 'v', 'a', '\x00'})), (*int8)(unsafe.Pointer(&[6]int8{'G', 'e', 'd', 'u', 'n', '\x00'})), (*int8)(unsafe.Pointer(&[6]int8{'I', 'r', 'i', 'n', 'a', '\x00'})), (*int8)(unsafe.Pointer(&[5]int8{'N', 'o', 'n', 'o', '\x00'})), (*int8)(unsafe.Pointer(&[7]int8{'W', 'i', 'e', 'n', 'e', 'r', '\x00'})), (*int8)(unsafe.Pointer(&[8]int8{'V', 'i', 'n', 'c', 'e', 'n', 't', '\x00'})), (*int8)(unsafe.Pointer(&[8]int8{'T', 's', 'e', 'r', 'i', 'n', 'g', '\x00'})), (*int8)(unsafe.Pointer(&[8]int8{'K', 'a', 'r', 'n', 'i', 'c', 'a', '\x00'})), (*int8)(unsafe.Pointer(&[5]int8{'L', 'u', 'l', 'u', '\x00'})), (*int8)(unsafe.Pointer(&[7]int8{'Q', 'u', 'i', 'n', 'c', 'y', '\x00'})), (*int8)(unsafe.Pointer(&[6]int8{'O', 's', 'a', 'm', 'a', '\x00'})), (*int8)(unsafe.Pointer(&[6]int8{'R', 'i', 'l', 'e', 'y', '\x00'})), (*int8)(unsafe.Pointer(&[7]int8{'U', 'r', 's', 'u', 'l', 'a', '\x00'})), (*int8)(unsafe.Pointer(&[4]int8{'S', 'a', 'm', '\x00'}))}
var _cgos_s_sorted_qsort [26]*int8 = [26]*int8{(*int8)(unsafe.Pointer(&[6]int8{'A', 'l', 'i', 'c', 'e', '\x00'})), (*int8)(unsafe.Pointer(&[4]int8{'B', 'o', 'b', '\x00'})), (*int8)(unsafe.Pointer(&[6]int8{'C', 'e', 'r', 'e', 's', '\x00'})), (*int8)(unsafe.Pointer(&[8]int8{'D', 'r', 'e', 'p', 'p', 'e', 'r', '\x00'})), (*int8)(unsafe.Pointer(&[9]int8{'E', 'm', 'e', 'r', 'a', 'l', 'd', 'a', '\x00'})), (*int8)(unsafe.Pointer(&[6]int8{'F', 'r', 'a', 'n', 'k', '\x00'})), (*int8)(unsafe.Pointer(&[6]int8{'G', 'e', 'd', 'u', 'n', '\x00'})), (*int8)(unsafe.Pointer(&[6]int8{'H', 'e', 'l', 'g', 'a', '\x00'})), (*int8)(unsafe.Pointer(&[6]int8{'I', 'r', 'i', 'n', 'a', '\x00'})), (*int8)(unsafe.Pointer(&[5]int8{'J', 'o', 'h', 'n', '\x00'})), (*int8)(unsafe.Pointer(&[8]int8{'K', 'a', 'r', 'n', 'i', 'c', 'a', '\x00'})), (*int8)(unsafe.Pointer(&[5]int8{'L', 'u', 'l', 'u', '\x00'})), (*int8)(unsafe.Pointer(&[5]int8{'M', 'o', 'm', 'o', '\x00'})), (*int8)(unsafe.Pointer(&[5]int8{'N', 'o', 'n', 'o', '\x00'})), (*int8)(unsafe.Pointer(&[6]int8{'O', 's', 'a', 'm', 'a', '\x00'})), (*int8)(unsafe.Pointer(&[5]int8{'P', 'e', 'm', 'a', '\x00'})), (*int8)(unsafe.Pointer(&[7]int8{'Q', 'u', 'i', 'n', 'c', 'y', '\x00'})), (*int8)(unsafe.Pointer(&[6]int8{'R', 'i', 'l', 'e', 'y', '\x00'})), (*int8)(unsafe.Pointer(&[4]int8{'S', 'a', 'm', '\x00'})), (*int8)(unsafe.Pointer(&[8]int8{'T', 's', 'e', 'r', 'i', 'n', 'g', '\x00'})), (*int8)(unsafe.Pointer(&[7]int8{'U', 'r', 's', 'u', 'l', 'a', '\x00'})), (*int8)(unsafe.Pointer(&[8]int8{'V', 'i', 'n', 'c', 'e', 'n', 't', '\x00'})), (*int8)(unsafe.Pointer(&[7]int8{'W', 'i', 'e', 'n', 'e', 'r', '\x00'})), (*int8)(unsafe.Pointer(&[7]int8{'X', 'a', 'v', 'i', 'e', 'r', '\x00'})), (*int8)(unsafe.Pointer(&[5]int8{'Y', 'e', 'v', 'a', '\x00'})), (*int8)(unsafe.Pointer(&[6]int8{'Z', 'o', 'r', 'a', 'n', '\x00'}))}
var _cgos_n_qsort [23]int32 = [23]int32{int32(879045), int32(394), int32(99405644), int32(33434), int32(232323), int32(4334), int32(5454), int32(343), int32(45545), int32(454), int32(324), int32(22), int32(34344), int32(233), int32(45345), int32(343), int32(848405), int32(3434), int32(3434344), int32(3535), int32(93994), int32(2230404), int32(4334)}
var _cgos_n_sorted_qsort [23]int32 = [23]int32{int32(22), int32(233), int32(324), int32(343), int32(343), int32(394), int32(454), int32(3434), int32(3535), int32(4334), int32(4334), int32(5454), int32(33434), int32(34344), int32(45345), int32(45545), int32(93994), int32(232323), int32(848405), int32(879045), int32(2230404), int32(3434344), int32(99405644)}

func _cgos_str_test_qsort(a **int8, a_sorted **int8, len int32) {
	var i int32
	libc.Qsort(unsafe.Pointer(a), uint64(len), 8, _cgos_scmp_qsort)
	for i = int32(0); i < len; i++ {
		if libc.Strcmp(*(**int8)(unsafe.Pointer(uintptr(unsafe.Pointer(a)) + uintptr(i)*8)), *(**int8)(unsafe.Pointer(uintptr(unsafe.Pointer(a_sorted)) + uintptr(i)*8))) != int32(0) {
			common.T_printf((*int8)(unsafe.Pointer(&[87]int8{'/', 'U', 's', 'e', 'r', 's', '/', 'v', 'f', 'c', '/', 'g', 'o', 'p', 'l', 'u', 's', '/', 'l', 'i', 'b', 'c', '/', 't', 'e', 's', 't', '/', 's', 'r', 'c', '/', 'f', 'u', 'n', 'c', 't', 'i', 'o', 'n', 'a', 'l', '/', 'q', 's', 'o', 'r', 't', '.', 'c', ':', '6', '6', ':', ' ', 's', 't', 'r', 'i', 'n', 'g', ' ', 's', 'o', 'r', 't', ' ', 'f', 'a', 'i', 'l', 'e', 'd', ' ', 'a', 't', ' ', 'i', 'n', 'd', 'e', 'x', ' ', '%', 'd', '\n', '\x00'})), i)
			common.T_printf((*int8)(unsafe.Pointer(&[13]int8{'\t', 'i', '\t', 'g', 'o', 't', '\t', 'w', 'a', 'n', 't', '\n', '\x00'})))
			for i = int32(0); i < len; i++ {
				common.T_printf((*int8)(unsafe.Pointer(&[11]int8{'\t', '%', 'd', '\t', '%', 's', '\t', '%', 's', '\n', '\x00'})), i, *(**int8)(unsafe.Pointer(uintptr(unsafe.Pointer(a)) + uintptr(i)*8)), *(**int8)(unsafe.Pointer(uintptr(unsafe.Pointer(a_sorted)) + uintptr(i)*8)))
			}
			break
		}
	}
}
func _cgos_int_test_qsort(a *int32, a_sorted *int32, len int32) {
	var i int32
	libc.Qsort(unsafe.Pointer(a), uint64(len), 4, _cgos_icmp_qsort)
	for i = int32(0); i < len; i++ {
		if *(*int32)(unsafe.Pointer(uintptr(unsafe.Pointer(a)) + uintptr(i)*4)) != *(*int32)(unsafe.Pointer(uintptr(unsafe.Pointer(a_sorted)) + uintptr(i)*4)) {
			common.T_printf((*int8)(unsafe.Pointer(&[88]int8{'/', 'U', 's', 'e', 'r', 's', '/', 'v', 'f', 'c', '/', 'g', 'o', 'p', 'l', 'u', 's', '/', 'l', 'i', 'b', 'c', '/', 't', 'e', 's', 't', '/', 's', 'r', 'c', '/', 'f', 'u', 'n', 'c', 't', 'i', 'o', 'n', 'a', 'l', '/', 'q', 's', 'o', 'r', 't', '.', 'c', ':', '8', '1', ':', ' ', 'i', 'n', 't', 'e', 'g', 'e', 'r', ' ', 's', 'o', 'r', 't', ' ', 'f', 'a', 'i', 'l', 'e', 'd', ' ', 'a', 't', ' ', 'i', 'n', 'd', 'e', 'x', ' ', '%', 'd', '\n', '\x00'})), i)
			common.T_printf((*int8)(unsafe.Pointer(&[13]int8{'\t', 'i', '\t', 'g', 'o', 't', '\t', 'w', 'a', 'n', 't', '\n', '\x00'})))
			for i = int32(0); i < len; i++ {
				common.T_printf((*int8)(unsafe.Pointer(&[11]int8{'\t', '%', 'd', '\t', '%', 'd', '\t', '%', 'd', '\n', '\x00'})), i, *(*int32)(unsafe.Pointer(uintptr(unsafe.Pointer(a)) + uintptr(i)*4)), *(*int32)(unsafe.Pointer(uintptr(unsafe.Pointer(a_sorted)) + uintptr(i)*4)))
			}
			break
		}
	}
}
func _cgos_uint64_gen_qsort(p *uint64, p_sorted *uint64, n int32) {
	var i int32
	var r uint64 = uint64(0)
	common.T_randseed(uint64(n))
	for i = int32(0); i < n; i++ {
		r += common.T_randn(uint64(20))
		*(*uint64)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + uintptr(i)*8)) = r
	}
	libc.Memcpy(unsafe.Pointer(p_sorted), unsafe.Pointer(p), uint64(n)*8)
	common.T_shuffle(p, uint64(n))
}
func _cgos_uint64_test_qsort(a *uint64, a_sorted *uint64, len int32) {
	var i int32
	libc.Qsort(unsafe.Pointer(a), uint64(len), 8, _cgos_cmp64_qsort)
	for i = int32(0); i < len; i++ {
		if *(*uint64)(unsafe.Pointer(uintptr(unsafe.Pointer(a)) + uintptr(i)*8)) != *(*uint64)(unsafe.Pointer(uintptr(unsafe.Pointer(a_sorted)) + uintptr(i)*8)) {
			common.T_printf((*int8)(unsafe.Pointer(&[88]int8{'/', 'U', 's', 'e', 'r', 's', '/', 'v', 'f', 'c', '/', 'g', 'o', 'p', 'l', 'u', 's', '/', 'l', 'i', 'b', 'c', '/', 't', 'e', 's', 't', '/', 's', 'r', 'c', '/', 'f', 'u', 'n', 'c', 't', 'i', 'o', 'n', 'a', 'l', '/', 'q', 's', 'o', 'r', 't', '.', 'c', ':', '1', '0', '9', ':', ' ', 'u', 'i', 'n', 't', '6', '4', ' ', 's', 'o', 'r', 't', ' ', 'f', 'a', 'i', 'l', 'e', 'd', ' ', 'a', 't', ' ', 'i', 'n', 'd', 'e', 'x', ' ', '%', 'd', '\n', '\x00'})), i)
			common.T_printf((*int8)(unsafe.Pointer(&[13]int8{'\t', 'i', '\t', 'g', 'o', 't', '\t', 'w', 'a', 'n', 't', '\n', '\x00'})))
			for i = int32(0); i < len; i++ {
				common.T_printf((*int8)(unsafe.Pointer(&[15]int8{'\t', '%', 'd', '\t', '%', 'l', 'l', 'u', '\t', '%', 'l', 'l', 'u', '\n', '\x00'})), i, *(*uint64)(unsafe.Pointer(uintptr(unsafe.Pointer(a)) + uintptr(i)*8)), *(*uint64)(unsafe.Pointer(uintptr(unsafe.Pointer(a_sorted)) + uintptr(i)*8)))
			}
			break
		}
	}
}
func _cgos_char_test_qsort() {
	for {
		var p [1]int8 = [1]int8{'\x00'}
		libc.Qsort(unsafe.Pointer((*int8)(unsafe.Pointer(&p))), 0, uint64(1), _cgos_ccmp_qsort)
		if libc.Memcmp(unsafe.Pointer((*int8)(unsafe.Pointer(&p))), unsafe.Pointer((*int8)(unsafe.Pointer(&[1]int8{'\x00'}))), 1) != int32(0) {
			common.T_printf((*int8)(unsafe.Pointer(&[79]int8{'/', 'U', 's', 'e', 'r', 's', '/', 'v', 'f', 'c', '/', 'g', 'o', 'p', 'l', 'u', 's', '/', 'l', 'i', 'b', 'c', '/', 't', 'e', 's', 't', '/', 's', 'r', 'c', '/', 'f', 'u', 'n', 'c', 't', 'i', 'o', 'n', 'a', 'l', '/', 'q', 's', 'o', 'r', 't', '.', 'c', ':', '1', '3', '0', ':', ' ', 'c', 'h', 'a', 'r', 'a', 'c', 't', 'e', 'r', ' ', 's', 'o', 'r', 't', ' ', 'f', 'a', 'i', 'l', 'e', 'd', '\n', '\x00'})))
			common.T_printf((*int8)(unsafe.Pointer(&[13]int8{'\t', 'g', 'o', 't', ':', ' ', ' ', '"', '%', 's', '"', '\n', '\x00'})), (*int8)(unsafe.Pointer(&p)))
			common.T_printf((*int8)(unsafe.Pointer(&[13]int8{'\t', 'w', 'a', 'n', 't', ':', ' ', '"', '%', 's', '"', '\n', '\x00'})), (*int8)(unsafe.Pointer(&[1]int8{'\x00'})))
		}
		if true {
			break
		}
	}
	for {
		var p [2]int8 = [2]int8{'1', '\x00'}
		libc.Qsort(unsafe.Pointer((*int8)(unsafe.Pointer(&p))), 1, uint64(1), _cgos_ccmp_qsort)
		if libc.Memcmp(unsafe.Pointer((*int8)(unsafe.Pointer(&p))), unsafe.Pointer((*int8)(unsafe.Pointer(&[2]int8{'1', '\x00'}))), 2) != int32(0) {
			common.T_printf((*int8)(unsafe.Pointer(&[79]int8{'/', 'U', 's', 'e', 'r', 's', '/', 'v', 'f', 'c', '/', 'g', 'o', 'p', 'l', 'u', 's', '/', 'l', 'i', 'b', 'c', '/', 't', 'e', 's', 't', '/', 's', 'r', 'c', '/', 'f', 'u', 'n', 'c', 't', 'i', 'o', 'n', 'a', 'l', '/', 'q', 's', 'o', 'r', 't', '.', 'c', ':', '1', '3', '1', ':', ' ', 'c', 'h', 'a', 'r', 'a', 'c', 't', 'e', 'r', ' ', 's', 'o', 'r', 't', ' ', 'f', 'a', 'i', 'l', 'e', 'd', '\n', '\x00'})))
			common.T_printf((*int8)(unsafe.Pointer(&[13]int8{'\t', 'g', 'o', 't', ':', ' ', ' ', '"', '%', 's', '"', '\n', '\x00'})), (*int8)(unsafe.Pointer(&p)))
			common.T_printf((*int8)(unsafe.Pointer(&[13]int8{'\t', 'w', 'a', 'n', 't', ':', ' ', '"', '%', 's', '"', '\n', '\x00'})), (*int8)(unsafe.Pointer(&[2]int8{'1', '\x00'})))
		}
		if true {
			break
		}
	}
	for {
		var p [3]int8 = [3]int8{'1', '1', '\x00'}
		libc.Qsort(unsafe.Pointer((*int8)(unsafe.Pointer(&p))), 2, uint64(1), _cgos_ccmp_qsort)
		if libc.Memcmp(unsafe.Pointer((*int8)(unsafe.Pointer(&p))), unsafe.Pointer((*int8)(unsafe.Pointer(&[3]int8{'1', '1', '\x00'}))), 3) != int32(0) {
			common.T_printf((*int8)(unsafe.Pointer(&[79]int8{'/', 'U', 's', 'e', 'r', 's', '/', 'v', 'f', 'c', '/', 'g', 'o', 'p', 'l', 'u', 's', '/', 'l', 'i', 'b', 'c', '/', 't', 'e', 's', 't', '/', 's', 'r', 'c', '/', 'f', 'u', 'n', 'c', 't', 'i', 'o', 'n', 'a', 'l', '/', 'q', 's', 'o', 'r', 't', '.', 'c', ':', '1', '3', '2', ':', ' ', 'c', 'h', 'a', 'r', 'a', 'c', 't', 'e', 'r', ' ', 's', 'o', 'r', 't', ' ', 'f', 'a', 'i', 'l', 'e', 'd', '\n', '\x00'})))
			common.T_printf((*int8)(unsafe.Pointer(&[13]int8{'\t', 'g', 'o', 't', ':', ' ', ' ', '"', '%', 's', '"', '\n', '\x00'})), (*int8)(unsafe.Pointer(&p)))
			common.T_printf((*int8)(unsafe.Pointer(&[13]int8{'\t', 'w', 'a', 'n', 't', ':', ' ', '"', '%', 's', '"', '\n', '\x00'})), (*int8)(unsafe.Pointer(&[3]int8{'1', '1', '\x00'})))
		}
		if true {
			break
		}
	}
	for {
		var p [3]int8 = [3]int8{'1', '2', '\x00'}
		libc.Qsort(unsafe.Pointer((*int8)(unsafe.Pointer(&p))), 2, uint64(1), _cgos_ccmp_qsort)
		if libc.Memcmp(unsafe.Pointer((*int8)(unsafe.Pointer(&p))), unsafe.Pointer((*int8)(unsafe.Pointer(&[3]int8{'1', '2', '\x00'}))), 3) != int32(0) {
			common.T_printf((*int8)(unsafe.Pointer(&[79]int8{'/', 'U', 's', 'e', 'r', 's', '/', 'v', 'f', 'c', '/', 'g', 'o', 'p', 'l', 'u', 's', '/', 'l', 'i', 'b', 'c', '/', 't', 'e', 's', 't', '/', 's', 'r', 'c', '/', 'f', 'u', 'n', 'c', 't', 'i', 'o', 'n', 'a', 'l', '/', 'q', 's', 'o', 'r', 't', '.', 'c', ':', '1', '3', '3', ':', ' ', 'c', 'h', 'a', 'r', 'a', 'c', 't', 'e', 'r', ' ', 's', 'o', 'r', 't', ' ', 'f', 'a', 'i', 'l', 'e', 'd', '\n', '\x00'})))
			common.T_printf((*int8)(unsafe.Pointer(&[13]int8{'\t', 'g', 'o', 't', ':', ' ', ' ', '"', '%', 's', '"', '\n', '\x00'})), (*int8)(unsafe.Pointer(&p)))
			common.T_printf((*int8)(unsafe.Pointer(&[13]int8{'\t', 'w', 'a', 'n', 't', ':', ' ', '"', '%', 's', '"', '\n', '\x00'})), (*int8)(unsafe.Pointer(&[3]int8{'1', '2', '\x00'})))
		}
		if true {
			break
		}
	}
	for {
		var p [3]int8 = [3]int8{'2', '1', '\x00'}
		libc.Qsort(unsafe.Pointer((*int8)(unsafe.Pointer(&p))), 2, uint64(1), _cgos_ccmp_qsort)
		if libc.Memcmp(unsafe.Pointer((*int8)(unsafe.Pointer(&p))), unsafe.Pointer((*int8)(unsafe.Pointer(&[3]int8{'1', '2', '\x00'}))), 3) != int32(0) {
			common.T_printf((*int8)(unsafe.Pointer(&[79]int8{'/', 'U', 's', 'e', 'r', 's', '/', 'v', 'f', 'c', '/', 'g', 'o', 'p', 'l', 'u', 's', '/', 'l', 'i', 'b', 'c', '/', 't', 'e', 's', 't', '/', 's', 'r', 'c', '/', 'f', 'u', 'n', 'c', 't', 'i', 'o', 'n', 'a', 'l', '/', 'q', 's', 'o', 'r', 't', '.', 'c', ':', '1', '3', '4', ':', ' ', 'c', 'h', 'a', 'r', 'a', 'c', 't', 'e', 'r', ' ', 's', 'o', 'r', 't', ' ', 'f', 'a', 'i', 'l', 'e', 'd', '\n', '\x00'})))
			common.T_printf((*int8)(unsafe.Pointer(&[13]int8{'\t', 'g', 'o', 't', ':', ' ', ' ', '"', '%', 's', '"', '\n', '\x00'})), (*int8)(unsafe.Pointer(&p)))
			common.T_printf((*int8)(unsafe.Pointer(&[13]int8{'\t', 'w', 'a', 'n', 't', ':', ' ', '"', '%', 's', '"', '\n', '\x00'})), (*int8)(unsafe.Pointer(&[3]int8{'1', '2', '\x00'})))
		}
		if true {
			break
		}
	}
	for {
		var p [4]int8 = [4]int8{'1', '1', '1', '\x00'}
		libc.Qsort(unsafe.Pointer((*int8)(unsafe.Pointer(&p))), 3, uint64(1), _cgos_ccmp_qsort)
		if libc.Memcmp(unsafe.Pointer((*int8)(unsafe.Pointer(&p))), unsafe.Pointer((*int8)(unsafe.Pointer(&[4]int8{'1', '1', '1', '\x00'}))), 4) != int32(0) {
			common.T_printf((*int8)(unsafe.Pointer(&[79]int8{'/', 'U', 's', 'e', 'r', 's', '/', 'v', 'f', 'c', '/', 'g', 'o', 'p', 'l', 'u', 's', '/', 'l', 'i', 'b', 'c', '/', 't', 'e', 's', 't', '/', 's', 'r', 'c', '/', 'f', 'u', 'n', 'c', 't', 'i', 'o', 'n', 'a', 'l', '/', 'q', 's', 'o', 'r', 't', '.', 'c', ':', '1', '3', '5', ':', ' ', 'c', 'h', 'a', 'r', 'a', 'c', 't', 'e', 'r', ' ', 's', 'o', 'r', 't', ' ', 'f', 'a', 'i', 'l', 'e', 'd', '\n', '\x00'})))
			common.T_printf((*int8)(unsafe.Pointer(&[13]int8{'\t', 'g', 'o', 't', ':', ' ', ' ', '"', '%', 's', '"', '\n', '\x00'})), (*int8)(unsafe.Pointer(&p)))
			common.T_printf((*int8)(unsafe.Pointer(&[13]int8{'\t', 'w', 'a', 'n', 't', ':', ' ', '"', '%', 's', '"', '\n', '\x00'})), (*int8)(unsafe.Pointer(&[4]int8{'1', '1', '1', '\x00'})))
		}
		if true {
			break
		}
	}
	for {
		var p [4]int8 = [4]int8{'2', '1', '1', '\x00'}
		libc.Qsort(unsafe.Pointer((*int8)(unsafe.Pointer(&p))), 3, uint64(1), _cgos_ccmp_qsort)
		if libc.Memcmp(unsafe.Pointer((*int8)(unsafe.Pointer(&p))), unsafe.Pointer((*int8)(unsafe.Pointer(&[4]int8{'1', '1', '2', '\x00'}))), 4) != int32(0) {
			common.T_printf((*int8)(unsafe.Pointer(&[79]int8{'/', 'U', 's', 'e', 'r', 's', '/', 'v', 'f', 'c', '/', 'g', 'o', 'p', 'l', 'u', 's', '/', 'l', 'i', 'b', 'c', '/', 't', 'e', 's', 't', '/', 's', 'r', 'c', '/', 'f', 'u', 'n', 'c', 't', 'i', 'o', 'n', 'a', 'l', '/', 'q', 's', 'o', 'r', 't', '.', 'c', ':', '1', '3', '6', ':', ' ', 'c', 'h', 'a', 'r', 'a', 'c', 't', 'e', 'r', ' ', 's', 'o', 'r', 't', ' ', 'f', 'a', 'i', 'l', 'e', 'd', '\n', '\x00'})))
			common.T_printf((*int8)(unsafe.Pointer(&[13]int8{'\t', 'g', 'o', 't', ':', ' ', ' ', '"', '%', 's', '"', '\n', '\x00'})), (*int8)(unsafe.Pointer(&p)))
			common.T_printf((*int8)(unsafe.Pointer(&[13]int8{'\t', 'w', 'a', 'n', 't', ':', ' ', '"', '%', 's', '"', '\n', '\x00'})), (*int8)(unsafe.Pointer(&[4]int8{'1', '1', '2', '\x00'})))
		}
		if true {
			break
		}
	}
	for {
		var p [4]int8 = [4]int8{'1', '2', '1', '\x00'}
		libc.Qsort(unsafe.Pointer((*int8)(unsafe.Pointer(&p))), 3, uint64(1), _cgos_ccmp_qsort)
		if libc.Memcmp(unsafe.Pointer((*int8)(unsafe.Pointer(&p))), unsafe.Pointer((*int8)(unsafe.Pointer(&[4]int8{'1', '1', '2', '\x00'}))), 4) != int32(0) {
			common.T_printf((*int8)(unsafe.Pointer(&[79]int8{'/', 'U', 's', 'e', 'r', 's', '/', 'v', 'f', 'c', '/', 'g', 'o', 'p', 'l', 'u', 's', '/', 'l', 'i', 'b', 'c', '/', 't', 'e', 's', 't', '/', 's', 'r', 'c', '/', 'f', 'u', 'n', 'c', 't', 'i', 'o', 'n', 'a', 'l', '/', 'q', 's', 'o', 'r', 't', '.', 'c', ':', '1', '3', '7', ':', ' ', 'c', 'h', 'a', 'r', 'a', 'c', 't', 'e', 'r', ' ', 's', 'o', 'r', 't', ' ', 'f', 'a', 'i', 'l', 'e', 'd', '\n', '\x00'})))
			common.T_printf((*int8)(unsafe.Pointer(&[13]int8{'\t', 'g', 'o', 't', ':', ' ', ' ', '"', '%', 's', '"', '\n', '\x00'})), (*int8)(unsafe.Pointer(&p)))
			common.T_printf((*int8)(unsafe.Pointer(&[13]int8{'\t', 'w', 'a', 'n', 't', ':', ' ', '"', '%', 's', '"', '\n', '\x00'})), (*int8)(unsafe.Pointer(&[4]int8{'1', '1', '2', '\x00'})))
		}
		if true {
			break
		}
	}
	for {
		var p [4]int8 = [4]int8{'1', '1', '2', '\x00'}
		libc.Qsort(unsafe.Pointer((*int8)(unsafe.Pointer(&p))), 3, uint64(1), _cgos_ccmp_qsort)
		if libc.Memcmp(unsafe.Pointer((*int8)(unsafe.Pointer(&p))), unsafe.Pointer((*int8)(unsafe.Pointer(&[4]int8{'1', '1', '2', '\x00'}))), 4) != int32(0) {
			common.T_printf((*int8)(unsafe.Pointer(&[79]int8{'/', 'U', 's', 'e', 'r', 's', '/', 'v', 'f', 'c', '/', 'g', 'o', 'p', 'l', 'u', 's', '/', 'l', 'i', 'b', 'c', '/', 't', 'e', 's', 't', '/', 's', 'r', 'c', '/', 'f', 'u', 'n', 'c', 't', 'i', 'o', 'n', 'a', 'l', '/', 'q', 's', 'o', 'r', 't', '.', 'c', ':', '1', '3', '8', ':', ' ', 'c', 'h', 'a', 'r', 'a', 'c', 't', 'e', 'r', ' ', 's', 'o', 'r', 't', ' ', 'f', 'a', 'i', 'l', 'e', 'd', '\n', '\x00'})))
			common.T_printf((*int8)(unsafe.Pointer(&[13]int8{'\t', 'g', 'o', 't', ':', ' ', ' ', '"', '%', 's', '"', '\n', '\x00'})), (*int8)(unsafe.Pointer(&p)))
			common.T_printf((*int8)(unsafe.Pointer(&[13]int8{'\t', 'w', 'a', 'n', 't', ':', ' ', '"', '%', 's', '"', '\n', '\x00'})), (*int8)(unsafe.Pointer(&[4]int8{'1', '1', '2', '\x00'})))
		}
		if true {
			break
		}
	}
	for {
		var p [4]int8 = [4]int8{'2', '2', '1', '\x00'}
		libc.Qsort(unsafe.Pointer((*int8)(unsafe.Pointer(&p))), 3, uint64(1), _cgos_ccmp_qsort)
		if libc.Memcmp(unsafe.Pointer((*int8)(unsafe.Pointer(&p))), unsafe.Pointer((*int8)(unsafe.Pointer(&[4]int8{'1', '2', '2', '\x00'}))), 4) != int32(0) {
			common.T_printf((*int8)(unsafe.Pointer(&[79]int8{'/', 'U', 's', 'e', 'r', 's', '/', 'v', 'f', 'c', '/', 'g', 'o', 'p', 'l', 'u', 's', '/', 'l', 'i', 'b', 'c', '/', 't', 'e', 's', 't', '/', 's', 'r', 'c', '/', 'f', 'u', 'n', 'c', 't', 'i', 'o', 'n', 'a', 'l', '/', 'q', 's', 'o', 'r', 't', '.', 'c', ':', '1', '3', '9', ':', ' ', 'c', 'h', 'a', 'r', 'a', 'c', 't', 'e', 'r', ' ', 's', 'o', 'r', 't', ' ', 'f', 'a', 'i', 'l', 'e', 'd', '\n', '\x00'})))
			common.T_printf((*int8)(unsafe.Pointer(&[13]int8{'\t', 'g', 'o', 't', ':', ' ', ' ', '"', '%', 's', '"', '\n', '\x00'})), (*int8)(unsafe.Pointer(&p)))
			common.T_printf((*int8)(unsafe.Pointer(&[13]int8{'\t', 'w', 'a', 'n', 't', ':', ' ', '"', '%', 's', '"', '\n', '\x00'})), (*int8)(unsafe.Pointer(&[4]int8{'1', '2', '2', '\x00'})))
		}
		if true {
			break
		}
	}
	for {
		var p [4]int8 = [4]int8{'2', '1', '2', '\x00'}
		libc.Qsort(unsafe.Pointer((*int8)(unsafe.Pointer(&p))), 3, uint64(1), _cgos_ccmp_qsort)
		if libc.Memcmp(unsafe.Pointer((*int8)(unsafe.Pointer(&p))), unsafe.Pointer((*int8)(unsafe.Pointer(&[4]int8{'1', '2', '2', '\x00'}))), 4) != int32(0) {
			common.T_printf((*int8)(unsafe.Pointer(&[79]int8{'/', 'U', 's', 'e', 'r', 's', '/', 'v', 'f', 'c', '/', 'g', 'o', 'p', 'l', 'u', 's', '/', 'l', 'i', 'b', 'c', '/', 't', 'e', 's', 't', '/', 's', 'r', 'c', '/', 'f', 'u', 'n', 'c', 't', 'i', 'o', 'n', 'a', 'l', '/', 'q', 's', 'o', 'r', 't', '.', 'c', ':', '1', '4', '0', ':', ' ', 'c', 'h', 'a', 'r', 'a', 'c', 't', 'e', 'r', ' ', 's', 'o', 'r', 't', ' ', 'f', 'a', 'i', 'l', 'e', 'd', '\n', '\x00'})))
			common.T_printf((*int8)(unsafe.Pointer(&[13]int8{'\t', 'g', 'o', 't', ':', ' ', ' ', '"', '%', 's', '"', '\n', '\x00'})), (*int8)(unsafe.Pointer(&p)))
			common.T_printf((*int8)(unsafe.Pointer(&[13]int8{'\t', 'w', 'a', 'n', 't', ':', ' ', '"', '%', 's', '"', '\n', '\x00'})), (*int8)(unsafe.Pointer(&[4]int8{'1', '2', '2', '\x00'})))
		}
		if true {
			break
		}
	}
	for {
		var p [4]int8 = [4]int8{'1', '2', '2', '\x00'}
		libc.Qsort(unsafe.Pointer((*int8)(unsafe.Pointer(&p))), 3, uint64(1), _cgos_ccmp_qsort)
		if libc.Memcmp(unsafe.Pointer((*int8)(unsafe.Pointer(&p))), unsafe.Pointer((*int8)(unsafe.Pointer(&[4]int8{'1', '2', '2', '\x00'}))), 4) != int32(0) {
			common.T_printf((*int8)(unsafe.Pointer(&[79]int8{'/', 'U', 's', 'e', 'r', 's', '/', 'v', 'f', 'c', '/', 'g', 'o', 'p', 'l', 'u', 's', '/', 'l', 'i', 'b', 'c', '/', 't', 'e', 's', 't', '/', 's', 'r', 'c', '/', 'f', 'u', 'n', 'c', 't', 'i', 'o', 'n', 'a', 'l', '/', 'q', 's', 'o', 'r', 't', '.', 'c', ':', '1', '4', '1', ':', ' ', 'c', 'h', 'a', 'r', 'a', 'c', 't', 'e', 'r', ' ', 's', 'o', 'r', 't', ' ', 'f', 'a', 'i', 'l', 'e', 'd', '\n', '\x00'})))
			common.T_printf((*int8)(unsafe.Pointer(&[13]int8{'\t', 'g', 'o', 't', ':', ' ', ' ', '"', '%', 's', '"', '\n', '\x00'})), (*int8)(unsafe.Pointer(&p)))
			common.T_printf((*int8)(unsafe.Pointer(&[13]int8{'\t', 'w', 'a', 'n', 't', ':', ' ', '"', '%', 's', '"', '\n', '\x00'})), (*int8)(unsafe.Pointer(&[4]int8{'1', '2', '2', '\x00'})))
		}
		if true {
			break
		}
	}
	for {
		var p [4]int8 = [4]int8{'1', '2', '3', '\x00'}
		libc.Qsort(unsafe.Pointer((*int8)(unsafe.Pointer(&p))), 3, uint64(1), _cgos_ccmp_qsort)
		if libc.Memcmp(unsafe.Pointer((*int8)(unsafe.Pointer(&p))), unsafe.Pointer((*int8)(unsafe.Pointer(&[4]int8{'1', '2', '3', '\x00'}))), 4) != int32(0) {
			common.T_printf((*int8)(unsafe.Pointer(&[79]int8{'/', 'U', 's', 'e', 'r', 's', '/', 'v', 'f', 'c', '/', 'g', 'o', 'p', 'l', 'u', 's', '/', 'l', 'i', 'b', 'c', '/', 't', 'e', 's', 't', '/', 's', 'r', 'c', '/', 'f', 'u', 'n', 'c', 't', 'i', 'o', 'n', 'a', 'l', '/', 'q', 's', 'o', 'r', 't', '.', 'c', ':', '1', '4', '2', ':', ' ', 'c', 'h', 'a', 'r', 'a', 'c', 't', 'e', 'r', ' ', 's', 'o', 'r', 't', ' ', 'f', 'a', 'i', 'l', 'e', 'd', '\n', '\x00'})))
			common.T_printf((*int8)(unsafe.Pointer(&[13]int8{'\t', 'g', 'o', 't', ':', ' ', ' ', '"', '%', 's', '"', '\n', '\x00'})), (*int8)(unsafe.Pointer(&p)))
			common.T_printf((*int8)(unsafe.Pointer(&[13]int8{'\t', 'w', 'a', 'n', 't', ':', ' ', '"', '%', 's', '"', '\n', '\x00'})), (*int8)(unsafe.Pointer(&[4]int8{'1', '2', '3', '\x00'})))
		}
		if true {
			break
		}
	}
	for {
		var p [4]int8 = [4]int8{'1', '3', '2', '\x00'}
		libc.Qsort(unsafe.Pointer((*int8)(unsafe.Pointer(&p))), 3, uint64(1), _cgos_ccmp_qsort)
		if libc.Memcmp(unsafe.Pointer((*int8)(unsafe.Pointer(&p))), unsafe.Pointer((*int8)(unsafe.Pointer(&[4]int8{'1', '2', '3', '\x00'}))), 4) != int32(0) {
			common.T_printf((*int8)(unsafe.Pointer(&[79]int8{'/', 'U', 's', 'e', 'r', 's', '/', 'v', 'f', 'c', '/', 'g', 'o', 'p', 'l', 'u', 's', '/', 'l', 'i', 'b', 'c', '/', 't', 'e', 's', 't', '/', 's', 'r', 'c', '/', 'f', 'u', 'n', 'c', 't', 'i', 'o', 'n', 'a', 'l', '/', 'q', 's', 'o', 'r', 't', '.', 'c', ':', '1', '4', '3', ':', ' ', 'c', 'h', 'a', 'r', 'a', 'c', 't', 'e', 'r', ' ', 's', 'o', 'r', 't', ' ', 'f', 'a', 'i', 'l', 'e', 'd', '\n', '\x00'})))
			common.T_printf((*int8)(unsafe.Pointer(&[13]int8{'\t', 'g', 'o', 't', ':', ' ', ' ', '"', '%', 's', '"', '\n', '\x00'})), (*int8)(unsafe.Pointer(&p)))
			common.T_printf((*int8)(unsafe.Pointer(&[13]int8{'\t', 'w', 'a', 'n', 't', ':', ' ', '"', '%', 's', '"', '\n', '\x00'})), (*int8)(unsafe.Pointer(&[4]int8{'1', '2', '3', '\x00'})))
		}
		if true {
			break
		}
	}
	for {
		var p [4]int8 = [4]int8{'2', '1', '3', '\x00'}
		libc.Qsort(unsafe.Pointer((*int8)(unsafe.Pointer(&p))), 3, uint64(1), _cgos_ccmp_qsort)
		if libc.Memcmp(unsafe.Pointer((*int8)(unsafe.Pointer(&p))), unsafe.Pointer((*int8)(unsafe.Pointer(&[4]int8{'1', '2', '3', '\x00'}))), 4) != int32(0) {
			common.T_printf((*int8)(unsafe.Pointer(&[79]int8{'/', 'U', 's', 'e', 'r', 's', '/', 'v', 'f', 'c', '/', 'g', 'o', 'p', 'l', 'u', 's', '/', 'l', 'i', 'b', 'c', '/', 't', 'e', 's', 't', '/', 's', 'r', 'c', '/', 'f', 'u', 'n', 'c', 't', 'i', 'o', 'n', 'a', 'l', '/', 'q', 's', 'o', 'r', 't', '.', 'c', ':', '1', '4', '4', ':', ' ', 'c', 'h', 'a', 'r', 'a', 'c', 't', 'e', 'r', ' ', 's', 'o', 'r', 't', ' ', 'f', 'a', 'i', 'l', 'e', 'd', '\n', '\x00'})))
			common.T_printf((*int8)(unsafe.Pointer(&[13]int8{'\t', 'g', 'o', 't', ':', ' ', ' ', '"', '%', 's', '"', '\n', '\x00'})), (*int8)(unsafe.Pointer(&p)))
			common.T_printf((*int8)(unsafe.Pointer(&[13]int8{'\t', 'w', 'a', 'n', 't', ':', ' ', '"', '%', 's', '"', '\n', '\x00'})), (*int8)(unsafe.Pointer(&[4]int8{'1', '2', '3', '\x00'})))
		}
		if true {
			break
		}
	}
	for {
		var p [4]int8 = [4]int8{'2', '3', '1', '\x00'}
		libc.Qsort(unsafe.Pointer((*int8)(unsafe.Pointer(&p))), 3, uint64(1), _cgos_ccmp_qsort)
		if libc.Memcmp(unsafe.Pointer((*int8)(unsafe.Pointer(&p))), unsafe.Pointer((*int8)(unsafe.Pointer(&[4]int8{'1', '2', '3', '\x00'}))), 4) != int32(0) {
			common.T_printf((*int8)(unsafe.Pointer(&[79]int8{'/', 'U', 's', 'e', 'r', 's', '/', 'v', 'f', 'c', '/', 'g', 'o', 'p', 'l', 'u', 's', '/', 'l', 'i', 'b', 'c', '/', 't', 'e', 's', 't', '/', 's', 'r', 'c', '/', 'f', 'u', 'n', 'c', 't', 'i', 'o', 'n', 'a', 'l', '/', 'q', 's', 'o', 'r', 't', '.', 'c', ':', '1', '4', '5', ':', ' ', 'c', 'h', 'a', 'r', 'a', 'c', 't', 'e', 'r', ' ', 's', 'o', 'r', 't', ' ', 'f', 'a', 'i', 'l', 'e', 'd', '\n', '\x00'})))
			common.T_printf((*int8)(unsafe.Pointer(&[13]int8{'\t', 'g', 'o', 't', ':', ' ', ' ', '"', '%', 's', '"', '\n', '\x00'})), (*int8)(unsafe.Pointer(&p)))
			common.T_printf((*int8)(unsafe.Pointer(&[13]int8{'\t', 'w', 'a', 'n', 't', ':', ' ', '"', '%', 's', '"', '\n', '\x00'})), (*int8)(unsafe.Pointer(&[4]int8{'1', '2', '3', '\x00'})))
		}
		if true {
			break
		}
	}
	for {
		var p [4]int8 = [4]int8{'3', '2', '1', '\x00'}
		libc.Qsort(unsafe.Pointer((*int8)(unsafe.Pointer(&p))), 3, uint64(1), _cgos_ccmp_qsort)
		if libc.Memcmp(unsafe.Pointer((*int8)(unsafe.Pointer(&p))), unsafe.Pointer((*int8)(unsafe.Pointer(&[4]int8{'1', '2', '3', '\x00'}))), 4) != int32(0) {
			common.T_printf((*int8)(unsafe.Pointer(&[79]int8{'/', 'U', 's', 'e', 'r', 's', '/', 'v', 'f', 'c', '/', 'g', 'o', 'p', 'l', 'u', 's', '/', 'l', 'i', 'b', 'c', '/', 't', 'e', 's', 't', '/', 's', 'r', 'c', '/', 'f', 'u', 'n', 'c', 't', 'i', 'o', 'n', 'a', 'l', '/', 'q', 's', 'o', 'r', 't', '.', 'c', ':', '1', '4', '6', ':', ' ', 'c', 'h', 'a', 'r', 'a', 'c', 't', 'e', 'r', ' ', 's', 'o', 'r', 't', ' ', 'f', 'a', 'i', 'l', 'e', 'd', '\n', '\x00'})))
			common.T_printf((*int8)(unsafe.Pointer(&[13]int8{'\t', 'g', 'o', 't', ':', ' ', ' ', '"', '%', 's', '"', '\n', '\x00'})), (*int8)(unsafe.Pointer(&p)))
			common.T_printf((*int8)(unsafe.Pointer(&[13]int8{'\t', 'w', 'a', 'n', 't', ':', ' ', '"', '%', 's', '"', '\n', '\x00'})), (*int8)(unsafe.Pointer(&[4]int8{'1', '2', '3', '\x00'})))
		}
		if true {
			break
		}
	}
	for {
		var p [4]int8 = [4]int8{'3', '1', '2', '\x00'}
		libc.Qsort(unsafe.Pointer((*int8)(unsafe.Pointer(&p))), 3, uint64(1), _cgos_ccmp_qsort)
		if libc.Memcmp(unsafe.Pointer((*int8)(unsafe.Pointer(&p))), unsafe.Pointer((*int8)(unsafe.Pointer(&[4]int8{'1', '2', '3', '\x00'}))), 4) != int32(0) {
			common.T_printf((*int8)(unsafe.Pointer(&[79]int8{'/', 'U', 's', 'e', 'r', 's', '/', 'v', 'f', 'c', '/', 'g', 'o', 'p', 'l', 'u', 's', '/', 'l', 'i', 'b', 'c', '/', 't', 'e', 's', 't', '/', 's', 'r', 'c', '/', 'f', 'u', 'n', 'c', 't', 'i', 'o', 'n', 'a', 'l', '/', 'q', 's', 'o', 'r', 't', '.', 'c', ':', '1', '4', '7', ':', ' ', 'c', 'h', 'a', 'r', 'a', 'c', 't', 'e', 'r', ' ', 's', 'o', 'r', 't', ' ', 'f', 'a', 'i', 'l', 'e', 'd', '\n', '\x00'})))
			common.T_printf((*int8)(unsafe.Pointer(&[13]int8{'\t', 'g', 'o', 't', ':', ' ', ' ', '"', '%', 's', '"', '\n', '\x00'})), (*int8)(unsafe.Pointer(&p)))
			common.T_printf((*int8)(unsafe.Pointer(&[13]int8{'\t', 'w', 'a', 'n', 't', ':', ' ', '"', '%', 's', '"', '\n', '\x00'})), (*int8)(unsafe.Pointer(&[4]int8{'1', '2', '3', '\x00'})))
		}
		if true {
			break
		}
	}
	for {
		var p [5]int8 = [5]int8{'1', '4', '2', '3', '\x00'}
		libc.Qsort(unsafe.Pointer((*int8)(unsafe.Pointer(&p))), 4, uint64(1), _cgos_ccmp_qsort)
		if libc.Memcmp(unsafe.Pointer((*int8)(unsafe.Pointer(&p))), unsafe.Pointer((*int8)(unsafe.Pointer(&[5]int8{'1', '2', '3', '4', '\x00'}))), 5) != int32(0) {
			common.T_printf((*int8)(unsafe.Pointer(&[79]int8{'/', 'U', 's', 'e', 'r', 's', '/', 'v', 'f', 'c', '/', 'g', 'o', 'p', 'l', 'u', 's', '/', 'l', 'i', 'b', 'c', '/', 't', 'e', 's', 't', '/', 's', 'r', 'c', '/', 'f', 'u', 'n', 'c', 't', 'i', 'o', 'n', 'a', 'l', '/', 'q', 's', 'o', 'r', 't', '.', 'c', ':', '1', '4', '8', ':', ' ', 'c', 'h', 'a', 'r', 'a', 'c', 't', 'e', 'r', ' ', 's', 'o', 'r', 't', ' ', 'f', 'a', 'i', 'l', 'e', 'd', '\n', '\x00'})))
			common.T_printf((*int8)(unsafe.Pointer(&[13]int8{'\t', 'g', 'o', 't', ':', ' ', ' ', '"', '%', 's', '"', '\n', '\x00'})), (*int8)(unsafe.Pointer(&p)))
			common.T_printf((*int8)(unsafe.Pointer(&[13]int8{'\t', 'w', 'a', 'n', 't', ':', ' ', '"', '%', 's', '"', '\n', '\x00'})), (*int8)(unsafe.Pointer(&[5]int8{'1', '2', '3', '4', '\x00'})))
		}
		if true {
			break
		}
	}
	for {
		var p [6]int8 = [6]int8{'5', '1', '3', '4', '2', '\x00'}
		libc.Qsort(unsafe.Pointer((*int8)(unsafe.Pointer(&p))), 5, uint64(1), _cgos_ccmp_qsort)
		if libc.Memcmp(unsafe.Pointer((*int8)(unsafe.Pointer(&p))), unsafe.Pointer((*int8)(unsafe.Pointer(&[6]int8{'1', '2', '3', '4', '5', '\x00'}))), 6) != int32(0) {
			common.T_printf((*int8)(unsafe.Pointer(&[79]int8{'/', 'U', 's', 'e', 'r', 's', '/', 'v', 'f', 'c', '/', 'g', 'o', 'p', 'l', 'u', 's', '/', 'l', 'i', 'b', 'c', '/', 't', 'e', 's', 't', '/', 's', 'r', 'c', '/', 'f', 'u', 'n', 'c', 't', 'i', 'o', 'n', 'a', 'l', '/', 'q', 's', 'o', 'r', 't', '.', 'c', ':', '1', '4', '9', ':', ' ', 'c', 'h', 'a', 'r', 'a', 'c', 't', 'e', 'r', ' ', 's', 'o', 'r', 't', ' ', 'f', 'a', 'i', 'l', 'e', 'd', '\n', '\x00'})))
			common.T_printf((*int8)(unsafe.Pointer(&[13]int8{'\t', 'g', 'o', 't', ':', ' ', ' ', '"', '%', 's', '"', '\n', '\x00'})), (*int8)(unsafe.Pointer(&p)))
			common.T_printf((*int8)(unsafe.Pointer(&[13]int8{'\t', 'w', 'a', 'n', 't', ':', ' ', '"', '%', 's', '"', '\n', '\x00'})), (*int8)(unsafe.Pointer(&[6]int8{'1', '2', '3', '4', '5', '\x00'})))
		}
		if true {
			break
		}
	}
	for {
		var p [7]int8 = [7]int8{'2', '6', '1', '4', '3', '5', '\x00'}
		libc.Qsort(unsafe.Pointer((*int8)(unsafe.Pointer(&p))), 6, uint64(1), _cgos_ccmp_qsort)
		if libc.Memcmp(unsafe.Pointer((*int8)(unsafe.Pointer(&p))), unsafe.Pointer((*int8)(unsafe.Pointer(&[7]int8{'1', '2', '3', '4', '5', '6', '\x00'}))), 7) != int32(0) {
			common.T_printf((*int8)(unsafe.Pointer(&[79]int8{'/', 'U', 's', 'e', 'r', 's', '/', 'v', 'f', 'c', '/', 'g', 'o', 'p', 'l', 'u', 's', '/', 'l', 'i', 'b', 'c', '/', 't', 'e', 's', 't', '/', 's', 'r', 'c', '/', 'f', 'u', 'n', 'c', 't', 'i', 'o', 'n', 'a', 'l', '/', 'q', 's', 'o', 'r', 't', '.', 'c', ':', '1', '5', '0', ':', ' ', 'c', 'h', 'a', 'r', 'a', 'c', 't', 'e', 'r', ' ', 's', 'o', 'r', 't', ' ', 'f', 'a', 'i', 'l', 'e', 'd', '\n', '\x00'})))
			common.T_printf((*int8)(unsafe.Pointer(&[13]int8{'\t', 'g', 'o', 't', ':', ' ', ' ', '"', '%', 's', '"', '\n', '\x00'})), (*int8)(unsafe.Pointer(&p)))
			common.T_printf((*int8)(unsafe.Pointer(&[13]int8{'\t', 'w', 'a', 'n', 't', ':', ' ', '"', '%', 's', '"', '\n', '\x00'})), (*int8)(unsafe.Pointer(&[7]int8{'1', '2', '3', '4', '5', '6', '\x00'})))
		}
		if true {
			break
		}
	}
	for {
		var p [8]int8 = [8]int8{'4', '5', '1', '7', '2', '6', '3', '\x00'}
		libc.Qsort(unsafe.Pointer((*int8)(unsafe.Pointer(&p))), 7, uint64(1), _cgos_ccmp_qsort)
		if libc.Memcmp(unsafe.Pointer((*int8)(unsafe.Pointer(&p))), unsafe.Pointer((*int8)(unsafe.Pointer(&[8]int8{'1', '2', '3', '4', '5', '6', '7', '\x00'}))), 8) != int32(0) {
			common.T_printf((*int8)(unsafe.Pointer(&[79]int8{'/', 'U', 's', 'e', 'r', 's', '/', 'v', 'f', 'c', '/', 'g', 'o', 'p', 'l', 'u', 's', '/', 'l', 'i', 'b', 'c', '/', 't', 'e', 's', 't', '/', 's', 'r', 'c', '/', 'f', 'u', 'n', 'c', 't', 'i', 'o', 'n', 'a', 'l', '/', 'q', 's', 'o', 'r', 't', '.', 'c', ':', '1', '5', '1', ':', ' ', 'c', 'h', 'a', 'r', 'a', 'c', 't', 'e', 'r', ' ', 's', 'o', 'r', 't', ' ', 'f', 'a', 'i', 'l', 'e', 'd', '\n', '\x00'})))
			common.T_printf((*int8)(unsafe.Pointer(&[13]int8{'\t', 'g', 'o', 't', ':', ' ', ' ', '"', '%', 's', '"', '\n', '\x00'})), (*int8)(unsafe.Pointer(&p)))
			common.T_printf((*int8)(unsafe.Pointer(&[13]int8{'\t', 'w', 'a', 'n', 't', ':', ' ', '"', '%', 's', '"', '\n', '\x00'})), (*int8)(unsafe.Pointer(&[8]int8{'1', '2', '3', '4', '5', '6', '7', '\x00'})))
		}
		if true {
			break
		}
	}
	for {
		var p [9]int8 = [9]int8{'3', '7', '2', '4', '5', '6', '1', '8', '\x00'}
		libc.Qsort(unsafe.Pointer((*int8)(unsafe.Pointer(&p))), 8, uint64(1), _cgos_ccmp_qsort)
		if libc.Memcmp(unsafe.Pointer((*int8)(unsafe.Pointer(&p))), unsafe.Pointer((*int8)(unsafe.Pointer(&[9]int8{'1', '2', '3', '4', '5', '6', '7', '8', '\x00'}))), 9) != int32(0) {
			common.T_printf((*int8)(unsafe.Pointer(&[79]int8{'/', 'U', 's', 'e', 'r', 's', '/', 'v', 'f', 'c', '/', 'g', 'o', 'p', 'l', 'u', 's', '/', 'l', 'i', 'b', 'c', '/', 't', 'e', 's', 't', '/', 's', 'r', 'c', '/', 'f', 'u', 'n', 'c', 't', 'i', 'o', 'n', 'a', 'l', '/', 'q', 's', 'o', 'r', 't', '.', 'c', ':', '1', '5', '2', ':', ' ', 'c', 'h', 'a', 'r', 'a', 'c', 't', 'e', 'r', ' ', 's', 'o', 'r', 't', ' ', 'f', 'a', 'i', 'l', 'e', 'd', '\n', '\x00'})))
			common.T_printf((*int8)(unsafe.Pointer(&[13]int8{'\t', 'g', 'o', 't', ':', ' ', ' ', '"', '%', 's', '"', '\n', '\x00'})), (*int8)(unsafe.Pointer(&p)))
			common.T_printf((*int8)(unsafe.Pointer(&[13]int8{'\t', 'w', 'a', 'n', 't', ':', ' ', '"', '%', 's', '"', '\n', '\x00'})), (*int8)(unsafe.Pointer(&[9]int8{'1', '2', '3', '4', '5', '6', '7', '8', '\x00'})))
		}
		if true {
			break
		}
	}
	for {
		var p [10]int8 = [10]int8{'8', '1', '2', '4', '3', '6', '5', '9', '7', '\x00'}
		libc.Qsort(unsafe.Pointer((*int8)(unsafe.Pointer(&p))), 9, uint64(1), _cgos_ccmp_qsort)
		if libc.Memcmp(unsafe.Pointer((*int8)(unsafe.Pointer(&p))), unsafe.Pointer((*int8)(unsafe.Pointer(&[10]int8{'1', '2', '3', '4', '5', '6', '7', '8', '9', '\x00'}))), 10) != int32(0) {
			common.T_printf((*int8)(unsafe.Pointer(&[79]int8{'/', 'U', 's', 'e', 'r', 's', '/', 'v', 'f', 'c', '/', 'g', 'o', 'p', 'l', 'u', 's', '/', 'l', 'i', 'b', 'c', '/', 't', 'e', 's', 't', '/', 's', 'r', 'c', '/', 'f', 'u', 'n', 'c', 't', 'i', 'o', 'n', 'a', 'l', '/', 'q', 's', 'o', 'r', 't', '.', 'c', ':', '1', '5', '3', ':', ' ', 'c', 'h', 'a', 'r', 'a', 'c', 't', 'e', 'r', ' ', 's', 'o', 'r', 't', ' ', 'f', 'a', 'i', 'l', 'e', 'd', '\n', '\x00'})))
			common.T_printf((*int8)(unsafe.Pointer(&[13]int8{'\t', 'g', 'o', 't', ':', ' ', ' ', '"', '%', 's', '"', '\n', '\x00'})), (*int8)(unsafe.Pointer(&p)))
			common.T_printf((*int8)(unsafe.Pointer(&[13]int8{'\t', 'w', 'a', 'n', 't', ':', ' ', '"', '%', 's', '"', '\n', '\x00'})), (*int8)(unsafe.Pointer(&[10]int8{'1', '2', '3', '4', '5', '6', '7', '8', '9', '\x00'})))
		}
		if true {
			break
		}
	}
	for {
		var p [10]int8 = [10]int8{'9', '8', '7', '6', '5', '4', '3', '2', '1', '\x00'}
		libc.Qsort(unsafe.Pointer((*int8)(unsafe.Pointer(&p))), 9, uint64(1), _cgos_ccmp_qsort)
		if libc.Memcmp(unsafe.Pointer((*int8)(unsafe.Pointer(&p))), unsafe.Pointer((*int8)(unsafe.Pointer(&[10]int8{'1', '2', '3', '4', '5', '6', '7', '8', '9', '\x00'}))), 10) != int32(0) {
			common.T_printf((*int8)(unsafe.Pointer(&[79]int8{'/', 'U', 's', 'e', 'r', 's', '/', 'v', 'f', 'c', '/', 'g', 'o', 'p', 'l', 'u', 's', '/', 'l', 'i', 'b', 'c', '/', 't', 'e', 's', 't', '/', 's', 'r', 'c', '/', 'f', 'u', 'n', 'c', 't', 'i', 'o', 'n', 'a', 'l', '/', 'q', 's', 'o', 'r', 't', '.', 'c', ':', '1', '5', '4', ':', ' ', 'c', 'h', 'a', 'r', 'a', 'c', 't', 'e', 'r', ' ', 's', 'o', 'r', 't', ' ', 'f', 'a', 'i', 'l', 'e', 'd', '\n', '\x00'})))
			common.T_printf((*int8)(unsafe.Pointer(&[13]int8{'\t', 'g', 'o', 't', ':', ' ', ' ', '"', '%', 's', '"', '\n', '\x00'})), (*int8)(unsafe.Pointer(&p)))
			common.T_printf((*int8)(unsafe.Pointer(&[13]int8{'\t', 'w', 'a', 'n', 't', ':', ' ', '"', '%', 's', '"', '\n', '\x00'})), (*int8)(unsafe.Pointer(&[10]int8{'1', '2', '3', '4', '5', '6', '7', '8', '9', '\x00'})))
		}
		if true {
			break
		}
	}
	for {
		var p [10]int8 = [10]int8{'3', '2', '1', '3', '2', '1', '3', '2', '1', '\x00'}
		libc.Qsort(unsafe.Pointer((*int8)(unsafe.Pointer(&p))), 9, uint64(1), _cgos_ccmp_qsort)
		if libc.Memcmp(unsafe.Pointer((*int8)(unsafe.Pointer(&p))), unsafe.Pointer((*int8)(unsafe.Pointer(&[10]int8{'1', '1', '1', '2', '2', '2', '3', '3', '3', '\x00'}))), 10) != int32(0) {
			common.T_printf((*int8)(unsafe.Pointer(&[79]int8{'/', 'U', 's', 'e', 'r', 's', '/', 'v', 'f', 'c', '/', 'g', 'o', 'p', 'l', 'u', 's', '/', 'l', 'i', 'b', 'c', '/', 't', 'e', 's', 't', '/', 's', 'r', 'c', '/', 'f', 'u', 'n', 'c', 't', 'i', 'o', 'n', 'a', 'l', '/', 'q', 's', 'o', 'r', 't', '.', 'c', ':', '1', '5', '5', ':', ' ', 'c', 'h', 'a', 'r', 'a', 'c', 't', 'e', 'r', ' ', 's', 'o', 'r', 't', ' ', 'f', 'a', 'i', 'l', 'e', 'd', '\n', '\x00'})))
			common.T_printf((*int8)(unsafe.Pointer(&[13]int8{'\t', 'g', 'o', 't', ':', ' ', ' ', '"', '%', 's', '"', '\n', '\x00'})), (*int8)(unsafe.Pointer(&p)))
			common.T_printf((*int8)(unsafe.Pointer(&[13]int8{'\t', 'w', 'a', 'n', 't', ':', ' ', '"', '%', 's', '"', '\n', '\x00'})), (*int8)(unsafe.Pointer(&[10]int8{'1', '1', '1', '2', '2', '2', '3', '3', '3', '\x00'})))
		}
		if true {
			break
		}
	}
	for {
		var p [18]int8 = [18]int8{'4', '9', '7', '3', '5', '8', '6', '2', '1', '8', '5', '2', '3', '6', '1', '7', '4', '\x00'}
		libc.Qsort(unsafe.Pointer((*int8)(unsafe.Pointer(&p))), 17, uint64(1), _cgos_ccmp_qsort)
		if libc.Memcmp(unsafe.Pointer((*int8)(unsafe.Pointer(&p))), unsafe.Pointer((*int8)(unsafe.Pointer(&[18]int8{'1', '1', '2', '2', '3', '3', '4', '4', '5', '5', '6', '6', '7', '7', '8', '8', '9', '\x00'}))), 18) != int32(0) {
			common.T_printf((*int8)(unsafe.Pointer(&[79]int8{'/', 'U', 's', 'e', 'r', 's', '/', 'v', 'f', 'c', '/', 'g', 'o', 'p', 'l', 'u', 's', '/', 'l', 'i', 'b', 'c', '/', 't', 'e', 's', 't', '/', 's', 'r', 'c', '/', 'f', 'u', 'n', 'c', 't', 'i', 'o', 'n', 'a', 'l', '/', 'q', 's', 'o', 'r', 't', '.', 'c', ':', '1', '5', '6', ':', ' ', 'c', 'h', 'a', 'r', 'a', 'c', 't', 'e', 'r', ' ', 's', 'o', 'r', 't', ' ', 'f', 'a', 'i', 'l', 'e', 'd', '\n', '\x00'})))
			common.T_printf((*int8)(unsafe.Pointer(&[13]int8{'\t', 'g', 'o', 't', ':', ' ', ' ', '"', '%', 's', '"', '\n', '\x00'})), (*int8)(unsafe.Pointer(&p)))
			common.T_printf((*int8)(unsafe.Pointer(&[13]int8{'\t', 'w', 'a', 'n', 't', ':', ' ', '"', '%', 's', '"', '\n', '\x00'})), (*int8)(unsafe.Pointer(&[18]int8{'1', '1', '2', '2', '3', '3', '4', '4', '5', '5', '6', '6', '7', '7', '8', '8', '9', '\x00'})))
		}
		if true {
			break
		}
	}
}
func _cgo_main() int32 {
	var i int32
	_cgos_str_test_qsort((**int8)(unsafe.Pointer(&_cgos_s_qsort)), (**int8)(unsafe.Pointer(&_cgos_s_sorted_qsort)), int32(26))
	_cgos_int_test_qsort((*int32)(unsafe.Pointer(&_cgos_n_qsort)), (*int32)(unsafe.Pointer(&_cgos_n_sorted_qsort)), int32(23))
	_cgos_char_test_qsort()
	for i = int32(1023); i <= int32(1026); i++ {
		var p [1026]uint64
		var p_sorted [1026]uint64
		_cgos_uint64_gen_qsort((*uint64)(unsafe.Pointer(&p)), (*uint64)(unsafe.Pointer(&p_sorted)), i)
		_cgos_uint64_test_qsort((*uint64)(unsafe.Pointer(&p)), (*uint64)(unsafe.Pointer(&p_sorted)), i)
	}
	return common.T_status
}
func TestMain(t *testing.T) {
	if _cgo_ret := _cgo_main(); _cgo_ret != 0 {
		t.Fatal("exit status", _cgo_ret)
	}
}
