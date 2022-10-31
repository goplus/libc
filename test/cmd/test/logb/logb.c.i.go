package logb

import (
	common "github.com/goplus/libc/test/common"
	unsafe "unsafe"
	libc "github.com/goplus/libc"
	testing "testing"
)

var _cgos_t_logb [18]common.Struct_d_d = [18]common.Struct_d_d{common.Struct_d_d{(*int8)(unsafe.Pointer(&[51]int8{'/', 'U', 's', 'e', 'r', 's', '/', 'v', 'f', 'c', '/', 'g', 'o', 'p', 'l', 'u', 's', '/', 'l', 'i', 'b', 'c', '/', 't', 'e', 's', 't', '/', 's', 'r', 'c', '/', 'm', 'a', 't', 'h', '/', 's', 'a', 'n', 'i', 't', 'y', '/', 'l', 'o', 'g', 'b', '.', 'h', '\x00'})), int32(1), int32(0), -8.0668483905796808, 3.0, float32(0.0), int32(0)}, common.Struct_d_d{(*int8)(unsafe.Pointer(&[51]int8{'/', 'U', 's', 'e', 'r', 's', '/', 'v', 'f', 'c', '/', 'g', 'o', 'p', 'l', 'u', 's', '/', 'l', 'i', 'b', 'c', '/', 't', 'e', 's', 't', '/', 's', 'r', 'c', '/', 'm', 'a', 't', 'h', '/', 's', 'a', 'n', 'i', 't', 'y', '/', 'l', 'o', 'g', 'b', '.', 'h', '\x00'})), int32(2), int32(0), 4.3452398493383049, 2.0, float32(0.0), int32(0)}, common.Struct_d_d{(*int8)(unsafe.Pointer(&[51]int8{'/', 'U', 's', 'e', 'r', 's', '/', 'v', 'f', 'c', '/', 'g', 'o', 'p', 'l', 'u', 's', '/', 'l', 'i', 'b', 'c', '/', 't', 'e', 's', 't', '/', 's', 'r', 'c', '/', 'm', 'a', 't', 'h', '/', 's', 'a', 'n', 'i', 't', 'y', '/', 'l', 'o', 'g', 'b', '.', 'h', '\x00'})), int32(3), int32(0), -8.3814334275552493, 3.0, float32(0.0), int32(0)}, common.Struct_d_d{(*int8)(unsafe.Pointer(&[51]int8{'/', 'U', 's', 'e', 'r', 's', '/', 'v', 'f', 'c', '/', 'g', 'o', 'p', 'l', 'u', 's', '/', 'l', 'i', 'b', 'c', '/', 't', 'e', 's', 't', '/', 's', 'r', 'c', '/', 'm', 'a', 't', 'h', '/', 's', 'a', 'n', 'i', 't', 'y', '/', 'l', 'o', 'g', 'b', '.', 'h', '\x00'})), int32(4), int32(0), -6.5316735819134841, 2.0, float32(0.0), int32(0)}, common.Struct_d_d{(*int8)(unsafe.Pointer(&[51]int8{'/', 'U', 's', 'e', 'r', 's', '/', 'v', 'f', 'c', '/', 'g', 'o', 'p', 'l', 'u', 's', '/', 'l', 'i', 'b', 'c', '/', 't', 'e', 's', 't', '/', 's', 'r', 'c', '/', 'm', 'a', 't', 'h', '/', 's', 'a', 'n', 'i', 't', 'y', '/', 'l', 'o', 'g', 'b', '.', 'h', '\x00'})), int32(5), int32(0), 9.2670569669725857, 3.0, float32(0.0), int32(0)}, common.Struct_d_d{(*int8)(unsafe.Pointer(&[51]int8{'/', 'U', 's', 'e', 'r', 's', '/', 'v', 'f', 'c', '/', 'g', 'o', 'p', 'l', 'u', 's', '/', 'l', 'i', 'b', 'c', '/', 't', 'e', 's', 't', '/', 's', 'r', 'c', '/', 'm', 'a', 't', 'h', '/', 's', 'a', 'n', 'i', 't', 'y', '/', 'l', 'o', 'g', 'b', '.', 'h', '\x00'})), int32(6), int32(0), 0.66198589809950448, -1.0, float32(0.0), int32(0)}, common.Struct_d_d{(*int8)(unsafe.Pointer(&[51]int8{'/', 'U', 's', 'e', 'r', 's', '/', 'v', 'f', 'c', '/', 'g', 'o', 'p', 'l', 'u', 's', '/', 'l', 'i', 'b', 'c', '/', 't', 'e', 's', 't', '/', 's', 'r', 'c', '/', 'm', 'a', 't', 'h', '/', 's', 'a', 'n', 'i', 't', 'y', '/', 'l', 'o', 'g', 'b', '.', 'h', '\x00'})), int32(7), int32(0), -0.40660392238535531, -2.0, float32(0.0), int32(0)}, common.Struct_d_d{(*int8)(unsafe.Pointer(&[51]int8{'/', 'U', 's', 'e', 'r', 's', '/', 'v', 'f', 'c', '/', 'g', 'o', 'p', 'l', 'u', 's', '/', 'l', 'i', 'b', 'c', '/', 't', 'e', 's', 't', '/', 's', 'r', 'c', '/', 'm', 'a', 't', 'h', '/', 's', 'a', 'n', 'i', 't', 'y', '/', 'l', 'o', 'g', 'b', '.', 'h', '\x00'})), int32(8), int32(0), 0.56175974622072411, -1.0, float32(0.0), int32(0)}, common.Struct_d_d{(*int8)(unsafe.Pointer(&[51]int8{'/', 'U', 's', 'e', 'r', 's', '/', 'v', 'f', 'c', '/', 'g', 'o', 'p', 'l', 'u', 's', '/', 'l', 'i', 'b', 'c', '/', 't', 'e', 's', 't', '/', 's', 'r', 'c', '/', 'm', 'a', 't', 'h', '/', 's', 'a', 'n', 'i', 't', 'y', '/', 'l', 'o', 'g', 'b', '.', 'h', '\x00'})), int32(9), int32(0), 0.77415229659130369, -1.0, float32(0.0), int32(0)}, common.Struct_d_d{(*int8)(unsafe.Pointer(&[51]int8{'/', 'U', 's', 'e', 'r', 's', '/', 'v', 'f', 'c', '/', 'g', 'o', 'p', 'l', 'u', 's', '/', 'l', 'i', 'b', 'c', '/', 't', 'e', 's', 't', '/', 's', 'r', 'c', '/', 'm', 'a', 't', 'h', '/', 's', 'a', 'n', 'i', 't', 'y', '/', 'l', 'o', 'g', 'b', '.', 'h', '\x00'})), int32(10), int32(0), -0.67876370263940244, -1.0, float32(0.0), int32(0)}, common.Struct_d_d{(*int8)(unsafe.Pointer(&[52]int8{'/', 'U', 's', 'e', 'r', 's', '/', 'v', 'f', 'c', '/', 'g', 'o', 'p', 'l', 'u', 's', '/', 'l', 'i', 'b', 'c', '/', 't', 'e', 's', 't', '/', 's', 'r', 'c', '/', 'm', 'a', 't', 'h', '/', 's', 'p', 'e', 'c', 'i', 'a', 'l', '/', 'l', 'o', 'g', 'b', '.', 'h', '\x00'})), int32(1), int32(0), 0.0, float64(-libc.X__builtin_inff()), float32(0.0), int32(0)}, common.Struct_d_d{(*int8)(unsafe.Pointer(&[52]int8{'/', 'U', 's', 'e', 'r', 's', '/', 'v', 'f', 'c', '/', 'g', 'o', 'p', 'l', 'u', 's', '/', 'l', 'i', 'b', 'c', '/', 't', 'e', 's', 't', '/', 's', 'r', 'c', '/', 'm', 'a', 't', 'h', '/', 's', 'p', 'e', 'c', 'i', 'a', 'l', '/', 'l', 'o', 'g', 'b', '.', 'h', '\x00'})), int32(2), int32(0), -0.0, float64(-libc.X__builtin_inff()), float32(0.0), int32(0)}, common.Struct_d_d{(*int8)(unsafe.Pointer(&[52]int8{'/', 'U', 's', 'e', 'r', 's', '/', 'v', 'f', 'c', '/', 'g', 'o', 'p', 'l', 'u', 's', '/', 'l', 'i', 'b', 'c', '/', 't', 'e', 's', 't', '/', 's', 'r', 'c', '/', 'm', 'a', 't', 'h', '/', 's', 'p', 'e', 'c', 'i', 'a', 'l', '/', 'l', 'o', 'g', 'b', '.', 'h', '\x00'})), int32(3), int32(0), -7.8886090522101181e-31, -100.0, float32(0.0), int32(0)}, common.Struct_d_d{(*int8)(unsafe.Pointer(&[52]int8{'/', 'U', 's', 'e', 'r', 's', '/', 'v', 'f', 'c', '/', 'g', 'o', 'p', 'l', 'u', 's', '/', 'l', 'i', 'b', 'c', '/', 't', 'e', 's', 't', '/', 's', 'r', 'c', '/', 'm', 'a', 't', 'h', '/', 's', 'p', 'e', 'c', 'i', 'a', 'l', '/', 'l', 'o', 'g', 'b', '.', 'h', '\x00'})), int32(4), int32(0), 1.0, 0.0, float32(0.0), int32(0)}, common.Struct_d_d{(*int8)(unsafe.Pointer(&[52]int8{'/', 'U', 's', 'e', 'r', 's', '/', 'v', 'f', 'c', '/', 'g', 'o', 'p', 'l', 'u', 's', '/', 'l', 'i', 'b', 'c', '/', 't', 'e', 's', 't', '/', 's', 'r', 'c', '/', 'm', 'a', 't', 'h', '/', 's', 'p', 'e', 'c', 'i', 'a', 'l', '/', 'l', 'o', 'g', 'b', '.', 'h', '\x00'})), int32(5), int32(0), -1.0, 0.0, float32(0.0), int32(0)}, common.Struct_d_d{(*int8)(unsafe.Pointer(&[52]int8{'/', 'U', 's', 'e', 'r', 's', '/', 'v', 'f', 'c', '/', 'g', 'o', 'p', 'l', 'u', 's', '/', 'l', 'i', 'b', 'c', '/', 't', 'e', 's', 't', '/', 's', 'r', 'c', '/', 'm', 'a', 't', 'h', '/', 's', 'p', 'e', 'c', 'i', 'a', 'l', '/', 'l', 'o', 'g', 'b', '.', 'h', '\x00'})), int32(6), int32(0), float64(libc.X__builtin_inff()), float64(libc.X__builtin_inff()), float32(0.0), int32(0)}, common.Struct_d_d{(*int8)(unsafe.Pointer(&[52]int8{'/', 'U', 's', 'e', 'r', 's', '/', 'v', 'f', 'c', '/', 'g', 'o', 'p', 'l', 'u', 's', '/', 'l', 'i', 'b', 'c', '/', 't', 'e', 's', 't', '/', 's', 'r', 'c', '/', 'm', 'a', 't', 'h', '/', 's', 'p', 'e', 'c', 'i', 'a', 'l', '/', 'l', 'o', 'g', 'b', '.', 'h', '\x00'})), int32(7), int32(0), float64(-libc.X__builtin_inff()), float64(libc.X__builtin_inff()), float32(0.0), int32(0)}, common.Struct_d_d{(*int8)(unsafe.Pointer(&[52]int8{'/', 'U', 's', 'e', 'r', 's', '/', 'v', 'f', 'c', '/', 'g', 'o', 'p', 'l', 'u', 's', '/', 'l', 'i', 'b', 'c', '/', 't', 'e', 's', 't', '/', 's', 'r', 'c', '/', 'm', 'a', 't', 'h', '/', 's', 'p', 'e', 'c', 'i', 'a', 'l', '/', 'l', 'o', 'g', 'b', '.', 'h', '\x00'})), int32(8), int32(0), float64(libc.X__builtin_nanf((*int8)(unsafe.Pointer(&[1]int8{'\x00'})))), float64(libc.X__builtin_nanf((*int8)(unsafe.Pointer(&[1]int8{'\x00'})))), float32(0.0), int32(0)}}

func _cgo_main() int32 {
	var y float64
	var d float32
	var e int32
	var i int32
	var err int32 = int32(0)
	var p *common.Struct_d_d
	for i = int32(0); uint64(i) < 18; i++ {
		p = (*common.Struct_d_d)(unsafe.Pointer(uintptr(unsafe.Pointer((*common.Struct_d_d)(unsafe.Pointer(&_cgos_t_logb)))) + uintptr(i)*40))
		if p.R < int32(0) {
			continue
		}
		libc.Fesetround(p.R)
		libc.Feclearexcept(int32(0))
		y = libc.Logb(p.X)
		e = libc.Fetestexcept(0)
		if !(common.Checkexceptall(e, p.E, p.R) != 0) {
			libc.Printf((*int8)(unsafe.Pointer(&[49]int8{'%', 's', ':', '%', 'd', ':', ' ', 'b', 'a', 'd', ' ', 'f', 'p', ' ', 'e', 'x', 'c', 'e', 'p', 't', 'i', 'o', 'n', ':', ' ', '%', 's', ' ', 'l', 'o', 'g', 'b', '(', '%', 'a', ')', '=', '%', 'a', ',', ' ', 'w', 'a', 'n', 't', ' ', '%', 's', '\x00'})), p.File, p.Line, common.Rstr(p.R), p.X, p.Y, common.Estr(p.E))
			libc.Printf((*int8)(unsafe.Pointer(&[9]int8{' ', 'g', 'o', 't', ' ', '%', 's', '\n', '\x00'})), common.Estr(e))
			err++
		}
		d = common.Ulperr(y, p.Y, p.Dy)
		if !(common.Checkcr(float64(y), float64(p.Y), p.R) != 0) {
			libc.Printf((*int8)(unsafe.Pointer(&[57]int8{'%', 's', ':', '%', 'd', ':', ' ', '%', 's', ' ', 'l', 'o', 'g', 'b', '(', '%', 'a', ')', ' ', 'w', 'a', 'n', 't', ' ', '%', 'a', ' ', 'g', 'o', 't', ' ', '%', 'a', ' ', 'u', 'l', 'p', 'e', 'r', 'r', ' ', '%', '.', '3', 'f', ' ', '=', ' ', '%', 'a', ' ', '+', ' ', '%', 'a', '\n', '\x00'})), p.File, p.Line, common.Rstr(p.R), p.X, p.Y, y, float64(d), float64(d-p.Dy), float64(p.Dy))
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
