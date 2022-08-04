package erf

import (
	common "github.com/goplus/libc/test/common"
	unsafe "unsafe"
	libc "github.com/goplus/libc"
	testing "testing"
)

var _cgos_t_erf [15]common.Struct_d_d = [15]common.Struct_d_d{common.Struct_d_d{(*int8)(unsafe.Pointer(&[53]int8{'/', 'U', 's', 'e', 'r', 's', '/', 'x', 'u', 's', 'h', 'i', 'w', 'e', 'i', '/', 'w', 'o', 'r', 'k', '/', 'l', 'i', 'b', 'c', '/', 't', 'e', 's', 't', '/', 's', 'r', 'c', '/', 'm', 'a', 't', 'h', '/', 's', 'a', 'n', 'i', 't', 'y', '/', 'e', 'r', 'f', '.', 'h', '\x00'})), int32(1), int32(0), -8.0668483905796808, -1, float32(-1.7128054820270654e-14), int32(0)}, common.Struct_d_d{(*int8)(unsafe.Pointer(&[53]int8{'/', 'U', 's', 'e', 'r', 's', '/', 'x', 'u', 's', 'h', 'i', 'w', 'e', 'i', '/', 'w', 'o', 'r', 'k', '/', 'l', 'i', 'b', 'c', '/', 't', 'e', 's', 't', '/', 's', 'r', 'c', '/', 'm', 'a', 't', 'h', '/', 's', 'a', 'n', 'i', 't', 'y', '/', 'e', 'r', 'f', '.', 'h', '\x00'})), int32(2), int32(0), 4.3452398493383049, 0.99999999920085402, float32(0.18850083649158478), int32(0)}, common.Struct_d_d{(*int8)(unsafe.Pointer(&[53]int8{'/', 'U', 's', 'e', 'r', 's', '/', 'x', 'u', 's', 'h', 'i', 'w', 'e', 'i', '/', 'w', 'o', 'r', 'k', '/', 'l', 'i', 'b', 'c', '/', 't', 'e', 's', 't', '/', 's', 'r', 'c', '/', 'm', 'a', 't', 'h', '/', 's', 'a', 'n', 'i', 't', 'y', '/', 'e', 'r', 'f', '.', 'h', '\x00'})), int32(3), int32(0), -8.3814334275552493, -1, float32(-9.3352229035835487e-17), int32(0)}, common.Struct_d_d{(*int8)(unsafe.Pointer(&[53]int8{'/', 'U', 's', 'e', 'r', 's', '/', 'x', 'u', 's', 'h', 'i', 'w', 'e', 'i', '/', 'w', 'o', 'r', 'k', '/', 'l', 'i', 'b', 'c', '/', 't', 'e', 's', 't', '/', 's', 'r', 'c', '/', 'm', 'a', 't', 'h', '/', 's', 'a', 'n', 'i', 't', 'y', '/', 'e', 'r', 'f', '.', 'h', '\x00'})), int32(4), int32(0), -6.5316735819134841, -1, float32(-1.1397524212952703e-4), int32(0)}, common.Struct_d_d{(*int8)(unsafe.Pointer(&[53]int8{'/', 'U', 's', 'e', 'r', 's', '/', 'x', 'u', 's', 'h', 'i', 'w', 'e', 'i', '/', 'w', 'o', 'r', 'k', '/', 'l', 'i', 'b', 'c', '/', 't', 'e', 's', 't', '/', 's', 'r', 'c', '/', 'm', 'a', 't', 'h', '/', 's', 'a', 'n', 'i', 't', 'y', '/', 'e', 'r', 'f', '.', 'h', '\x00'})), int32(5), int32(0), 9.2670569669725857, 1, float32(1.3234889800848443e-23), int32(0)}, common.Struct_d_d{(*int8)(unsafe.Pointer(&[53]int8{'/', 'U', 's', 'e', 'r', 's', '/', 'x', 'u', 's', 'h', 'i', 'w', 'e', 'i', '/', 'w', 'o', 'r', 'k', '/', 'l', 'i', 'b', 'c', '/', 't', 'e', 's', 't', '/', 's', 'r', 'c', '/', 'm', 'a', 't', 'h', '/', 's', 'a', 'n', 'i', 't', 'y', '/', 'e', 'r', 'f', '.', 'h', '\x00'})), int32(6), int32(0), 0.66198589809950448, 0.65082433828191166, float32(0.35772353410720825), int32(0)}, common.Struct_d_d{(*int8)(unsafe.Pointer(&[53]int8{'/', 'U', 's', 'e', 'r', 's', '/', 'x', 'u', 's', 'h', 'i', 'w', 'e', 'i', '/', 'w', 'o', 'r', 'k', '/', 'l', 'i', 'b', 'c', '/', 't', 'e', 's', 't', '/', 's', 'r', 'c', '/', 'm', 'a', 't', 'h', '/', 's', 'a', 'n', 'i', 't', 'y', '/', 'e', 'r', 'f', '.', 'h', '\x00'})), int32(7), int32(0), -0.40660392238535531, -0.43472546289517394, float32(-0.2345312237739563), int32(0)}, common.Struct_d_d{(*int8)(unsafe.Pointer(&[53]int8{'/', 'U', 's', 'e', 'r', 's', '/', 'x', 'u', 's', 'h', 'i', 'w', 'e', 'i', '/', 'w', 'o', 'r', 'k', '/', 'l', 'i', 'b', 'c', '/', 't', 'e', 's', 't', '/', 's', 'r', 'c', '/', 'm', 'a', 't', 'h', '/', 's', 'a', 'n', 'i', 't', 'y', '/', 'e', 'r', 'f', '.', 'h', '\x00'})), int32(8), int32(0), 0.56175974622072411, 0.5730654766932296, float32(0.050273332744836807), int32(0)}, common.Struct_d_d{(*int8)(unsafe.Pointer(&[53]int8{'/', 'U', 's', 'e', 'r', 's', '/', 'x', 'u', 's', 'h', 'i', 'w', 'e', 'i', '/', 'w', 'o', 'r', 'k', '/', 'l', 'i', 'b', 'c', '/', 't', 'e', 's', 't', '/', 's', 'r', 'c', '/', 'm', 'a', 't', 'h', '/', 's', 'a', 'n', 'i', 't', 'y', '/', 'e', 'r', 'f', '.', 'h', '\x00'})), int32(9), int32(0), 0.77415229659130369, 0.72640304139050882, float32(0.022778626531362534), int32(0)}, common.Struct_d_d{(*int8)(unsafe.Pointer(&[53]int8{'/', 'U', 's', 'e', 'r', 's', '/', 'x', 'u', 's', 'h', 'i', 'w', 'e', 'i', '/', 'w', 'o', 'r', 'k', '/', 'l', 'i', 'b', 'c', '/', 't', 'e', 's', 't', '/', 's', 'r', 'c', '/', 'm', 'a', 't', 'h', '/', 's', 'a', 'n', 'i', 't', 'y', '/', 'e', 'r', 'f', '.', 'h', '\x00'})), int32(10), int32(0), -0.67876370263940244, -0.66290292684822816, float32(-0.21418938040733337), int32(0)}, common.Struct_d_d{(*int8)(unsafe.Pointer(&[54]int8{'/', 'U', 's', 'e', 'r', 's', '/', 'x', 'u', 's', 'h', 'i', 'w', 'e', 'i', '/', 'w', 'o', 'r', 'k', '/', 'l', 'i', 'b', 'c', '/', 't', 'e', 's', 't', '/', 's', 'r', 'c', '/', 'm', 'a', 't', 'h', '/', 's', 'p', 'e', 'c', 'i', 'a', 'l', '/', 'e', 'r', 'f', '.', 'h', '\x00'})), int32(1), int32(0), 0, 0, float32(0), int32(0)}, common.Struct_d_d{(*int8)(unsafe.Pointer(&[54]int8{'/', 'U', 's', 'e', 'r', 's', '/', 'x', 'u', 's', 'h', 'i', 'w', 'e', 'i', '/', 'w', 'o', 'r', 'k', '/', 'l', 'i', 'b', 'c', '/', 't', 'e', 's', 't', '/', 's', 'r', 'c', '/', 'm', 'a', 't', 'h', '/', 's', 'p', 'e', 'c', 'i', 'a', 'l', '/', 'e', 'r', 'f', '.', 'h', '\x00'})), int32(2), int32(0), -0, -0, float32(0), int32(0)}, common.Struct_d_d{(*int8)(unsafe.Pointer(&[54]int8{'/', 'U', 's', 'e', 'r', 's', '/', 'x', 'u', 's', 'h', 'i', 'w', 'e', 'i', '/', 'w', 'o', 'r', 'k', '/', 'l', 'i', 'b', 'c', '/', 't', 'e', 's', 't', '/', 's', 'r', 'c', '/', 'm', 'a', 't', 'h', '/', 's', 'p', 'e', 'c', 'i', 'a', 'l', '/', 'e', 'r', 'f', '.', 'h', '\x00'})), int32(3), int32(0), float64(libc.X__builtin_inff()), 1, float32(0), int32(0)}, common.Struct_d_d{(*int8)(unsafe.Pointer(&[54]int8{'/', 'U', 's', 'e', 'r', 's', '/', 'x', 'u', 's', 'h', 'i', 'w', 'e', 'i', '/', 'w', 'o', 'r', 'k', '/', 'l', 'i', 'b', 'c', '/', 't', 'e', 's', 't', '/', 's', 'r', 'c', '/', 'm', 'a', 't', 'h', '/', 's', 'p', 'e', 'c', 'i', 'a', 'l', '/', 'e', 'r', 'f', '.', 'h', '\x00'})), int32(4), int32(0), float64(-libc.X__builtin_inff()), -1, float32(0), int32(0)}, common.Struct_d_d{(*int8)(unsafe.Pointer(&[54]int8{'/', 'U', 's', 'e', 'r', 's', '/', 'x', 'u', 's', 'h', 'i', 'w', 'e', 'i', '/', 'w', 'o', 'r', 'k', '/', 'l', 'i', 'b', 'c', '/', 't', 'e', 's', 't', '/', 's', 'r', 'c', '/', 'm', 'a', 't', 'h', '/', 's', 'p', 'e', 'c', 'i', 'a', 'l', '/', 'e', 'r', 'f', '.', 'h', '\x00'})), int32(5), int32(0), float64(libc.X__builtin_nanf((*int8)(unsafe.Pointer(&[1]int8{'\x00'})))), float64(libc.X__builtin_nanf((*int8)(unsafe.Pointer(&[1]int8{'\x00'})))), float32(0), int32(0)}}

func _cgo_main() int32 {
	var y float64
	var d float32
	var e int32
	var i int32
	var err int32 = int32(0)
	var p *common.Struct_d_d
	for i = int32(0); uint64(i) < 15; i++ {
		p = (*common.Struct_d_d)(unsafe.Pointer(uintptr(unsafe.Pointer((*common.Struct_d_d)(unsafe.Pointer(&_cgos_t_erf)))) + uintptr(i)*40))
		if p.R < int32(0) {
			continue
		}
		libc.Fesetround(p.R)
		libc.Feclearexcept(int32(0))
		y = libc.Erf(p.X)
		e = libc.Fetestexcept(0)
		if !(common.Checkexcept(e, p.E, p.R) != 0) {
			libc.Printf((*int8)(unsafe.Pointer(&[48]int8{'%', 's', ':', '%', 'd', ':', ' ', 'b', 'a', 'd', ' ', 'f', 'p', ' ', 'e', 'x', 'c', 'e', 'p', 't', 'i', 'o', 'n', ':', ' ', '%', 's', ' ', 'e', 'r', 'f', '(', '%', 'a', ')', '=', '%', 'a', ',', ' ', 'w', 'a', 'n', 't', ' ', '%', 's', '\x00'})), p.File, p.Line, common.Rstr(p.R), p.X, p.Y, common.Estr(p.E))
			libc.Printf((*int8)(unsafe.Pointer(&[9]int8{' ', 'g', 'o', 't', ' ', '%', 's', '\n', '\x00'})), common.Estr(e))
			err++
		}
		d = common.Ulperr(y, p.Y, p.Dy)
		if !(common.Checkulp(d, p.R) != 0) {
			if libc.Fabs(float64(d)) < float64(4) {
				libc.Printf((*int8)(unsafe.Pointer(&[3]int8{'X', ' ', '\x00'})))
			} else {
				err++
			}
			libc.Printf((*int8)(unsafe.Pointer(&[56]int8{'%', 's', ':', '%', 'd', ':', ' ', '%', 's', ' ', 'e', 'r', 'f', '(', '%', 'a', ')', ' ', 'w', 'a', 'n', 't', ' ', '%', 'a', ' ', 'g', 'o', 't', ' ', '%', 'a', ' ', 'u', 'l', 'p', 'e', 'r', 'r', ' ', '%', '.', '3', 'f', ' ', '=', ' ', '%', 'a', ' ', '+', ' ', '%', 'a', '\n', '\x00'})), p.File, p.Line, common.Rstr(p.R), p.X, p.Y, y, float64(d), float64(d-p.Dy), float64(p.Dy))
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
