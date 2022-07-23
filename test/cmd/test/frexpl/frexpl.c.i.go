package frexpl

import (
	unsafe "unsafe"
	libc "github.com/goplus/libc"
	testing "testing"
)

var _cgos_t_frexpl [21]struct_l_li = [21]struct_l_li{struct_l_li{(*int8)(unsafe.Pointer(&[24]int8{'s', 'r', 'c', '/', 'm', 'a', 't', 'h', '/', 's', 'a', 'n', 'i', 't', 'y', '/', 'f', 'r', 'e', 'x', 'p', '.', 'h', '\x00'})), int32(1), int32(0), float64(-8.0668483905796808), float64(-0.50417802441123005), float32(0), int64(4), int32(0)}, struct_l_li{(*int8)(unsafe.Pointer(&[24]int8{'s', 'r', 'c', '/', 'm', 'a', 't', 'h', '/', 's', 'a', 'n', 'i', 't', 'y', '/', 'f', 'r', 'e', 'x', 'p', '.', 'h', '\x00'})), int32(2), int32(0), float64(4.3452398493383049), float64(0.54315498116728811), float32(0), int64(3), int32(0)}, struct_l_li{(*int8)(unsafe.Pointer(&[24]int8{'s', 'r', 'c', '/', 'm', 'a', 't', 'h', '/', 's', 'a', 'n', 'i', 't', 'y', '/', 'f', 'r', 'e', 'x', 'p', '.', 'h', '\x00'})), int32(3), int32(0), float64(-8.3814334275552493), float64(-0.52383958922220308), float32(0), int64(4), int32(0)}, struct_l_li{(*int8)(unsafe.Pointer(&[24]int8{'s', 'r', 'c', '/', 'm', 'a', 't', 'h', '/', 's', 'a', 'n', 'i', 't', 'y', '/', 'f', 'r', 'e', 'x', 'p', '.', 'h', '\x00'})), int32(4), int32(0), float64(-6.5316735819134841), float64(-0.81645919773918551), float32(0), int64(3), int32(0)}, struct_l_li{(*int8)(unsafe.Pointer(&[24]int8{'s', 'r', 'c', '/', 'm', 'a', 't', 'h', '/', 's', 'a', 'n', 'i', 't', 'y', '/', 'f', 'r', 'e', 'x', 'p', '.', 'h', '\x00'})), int32(5), int32(0), float64(9.2670569669725857), float64(0.57919106043578661), float32(0), int64(4), int32(0)}, struct_l_li{(*int8)(unsafe.Pointer(&[24]int8{'s', 'r', 'c', '/', 'm', 'a', 't', 'h', '/', 's', 'a', 'n', 'i', 't', 'y', '/', 'f', 'r', 'e', 'x', 'p', '.', 'h', '\x00'})), int32(6), int32(0), float64(0.66198589809950448), float64(0.66198589809950448), float32(0), int64(0), int32(0)}, struct_l_li{(*int8)(unsafe.Pointer(&[24]int8{'s', 'r', 'c', '/', 'm', 'a', 't', 'h', '/', 's', 'a', 'n', 'i', 't', 'y', '/', 'f', 'r', 'e', 'x', 'p', '.', 'h', '\x00'})), int32(7), int32(0), float64(-0.40660392238535531), float64(-0.81320784477071062), float32(0), int64(-1), int32(0)}, struct_l_li{(*int8)(unsafe.Pointer(&[24]int8{'s', 'r', 'c', '/', 'm', 'a', 't', 'h', '/', 's', 'a', 'n', 'i', 't', 'y', '/', 'f', 'r', 'e', 'x', 'p', '.', 'h', '\x00'})), int32(8), int32(0), float64(0.56175974622072411), float64(0.56175974622072411), float32(0), int64(0), int32(0)}, struct_l_li{(*int8)(unsafe.Pointer(&[24]int8{'s', 'r', 'c', '/', 'm', 'a', 't', 'h', '/', 's', 'a', 'n', 'i', 't', 'y', '/', 'f', 'r', 'e', 'x', 'p', '.', 'h', '\x00'})), int32(9), int32(0), float64(0.77415229659130369), float64(0.77415229659130369), float32(0), int64(0), int32(0)}, struct_l_li{(*int8)(unsafe.Pointer(&[24]int8{'s', 'r', 'c', '/', 'm', 'a', 't', 'h', '/', 's', 'a', 'n', 'i', 't', 'y', '/', 'f', 'r', 'e', 'x', 'p', '.', 'h', '\x00'})), int32(10), int32(0), float64(-0.67876370263940244), float64(-0.67876370263940244), float32(0), int64(0), int32(0)}, struct_l_li{(*int8)(unsafe.Pointer(&[25]int8{'s', 'r', 'c', '/', 'm', 'a', 't', 'h', '/', 's', 'p', 'e', 'c', 'i', 'a', 'l', '/', 'f', 'r', 'e', 'x', 'p', '.', 'h', '\x00'})), int32(1), int32(0), float64(0), float64(0), float32(0), int64(0), int32(0)}, struct_l_li{(*int8)(unsafe.Pointer(&[25]int8{'s', 'r', 'c', '/', 'm', 'a', 't', 'h', '/', 's', 'p', 'e', 'c', 'i', 'a', 'l', '/', 'f', 'r', 'e', 'x', 'p', '.', 'h', '\x00'})), int32(2), int32(0), float64(-0), float64(-0), float32(0), int64(0), int32(0)}, struct_l_li{(*int8)(unsafe.Pointer(&[25]int8{'s', 'r', 'c', '/', 'm', 'a', 't', 'h', '/', 's', 'p', 'e', 'c', 'i', 'a', 'l', '/', 'f', 'r', 'e', 'x', 'p', '.', 'h', '\x00'})), int32(3), int32(0), float64(0.5), float64(0.5), float32(0), int64(0), int32(0)}, struct_l_li{(*int8)(unsafe.Pointer(&[25]int8{'s', 'r', 'c', '/', 'm', 'a', 't', 'h', '/', 's', 'p', 'e', 'c', 'i', 'a', 'l', '/', 'f', 'r', 'e', 'x', 'p', '.', 'h', '\x00'})), int32(4), int32(0), float64(-0.5), float64(-0.5), float32(0), int64(0), int32(0)}, struct_l_li{(*int8)(unsafe.Pointer(&[25]int8{'s', 'r', 'c', '/', 'm', 'a', 't', 'h', '/', 's', 'p', 'e', 'c', 'i', 'a', 'l', '/', 'f', 'r', 'e', 'x', 'p', '.', 'h', '\x00'})), int32(5), int32(0), float64(1), float64(0.5), float32(0), int64(1), int32(0)}, struct_l_li{(*int8)(unsafe.Pointer(&[25]int8{'s', 'r', 'c', '/', 'm', 'a', 't', 'h', '/', 's', 'p', 'e', 'c', 'i', 'a', 'l', '/', 'f', 'r', 'e', 'x', 'p', '.', 'h', '\x00'})), int32(6), int32(0), float64(-1), float64(-0.5), float32(0), int64(1), int32(0)}, struct_l_li{(*int8)(unsafe.Pointer(&[25]int8{'s', 'r', 'c', '/', 'm', 'a', 't', 'h', '/', 's', 'p', 'e', 'c', 'i', 'a', 'l', '/', 'f', 'r', 'e', 'x', 'p', '.', 'h', '\x00'})), int32(7), int32(0), float64(2), float64(0.5), float32(0), int64(2), int32(0)}, struct_l_li{(*int8)(unsafe.Pointer(&[25]int8{'s', 'r', 'c', '/', 'm', 'a', 't', 'h', '/', 's', 'p', 'e', 'c', 'i', 'a', 'l', '/', 'f', 'r', 'e', 'x', 'p', '.', 'h', '\x00'})), int32(8), int32(0), float64(-2), float64(-0.5), float32(0), int64(2), int32(0)}, struct_l_li{(*int8)(unsafe.Pointer(&[25]int8{'s', 'r', 'c', '/', 'm', 'a', 't', 'h', '/', 's', 'p', 'e', 'c', 'i', 'a', 'l', '/', 'f', 'r', 'e', 'x', 'p', '.', 'h', '\x00'})), int32(10), int32(0), float64(libc.X__builtin_inff()), float64(libc.X__builtin_inff()), float32(0), int64(0), int32(0)}, struct_l_li{(*int8)(unsafe.Pointer(&[25]int8{'s', 'r', 'c', '/', 'm', 'a', 't', 'h', '/', 's', 'p', 'e', 'c', 'i', 'a', 'l', '/', 'f', 'r', 'e', 'x', 'p', '.', 'h', '\x00'})), int32(11), int32(0), float64(-libc.X__builtin_inff()), float64(-libc.X__builtin_inff()), float32(0), int64(0), int32(0)}, struct_l_li{(*int8)(unsafe.Pointer(&[25]int8{'s', 'r', 'c', '/', 'm', 'a', 't', 'h', '/', 's', 'p', 'e', 'c', 'i', 'a', 'l', '/', 'f', 'r', 'e', 'x', 'p', '.', 'h', '\x00'})), int32(12), int32(0), float64(libc.X__builtin_nanf((*int8)(unsafe.Pointer(&[1]int8{'\x00'})))), float64(libc.X__builtin_nanf((*int8)(unsafe.Pointer(&[1]int8{'\x00'})))), float32(0), int64(0), int32(0)}}

func _cgo_main() int32 {
	var yi int32
	var y float64
	var d float32
	var e int32
	var i int32
	var err int32 = int32(0)
	var p *struct_l_li
	for i = int32(0); uint64(i) < 22; i++ {
		p = (*struct_l_li)(unsafe.Pointer(uintptr(unsafe.Pointer((*struct_l_li)(unsafe.Pointer(&_cgos_t_frexpl)))) + uintptr(i)*52))
		if p.r < int32(0) {
			continue
		}
		libc.Fesetround(p.r)
		libc.Feclearexcept(int32(0))
		y = libc.Frexpl(p.x, &yi)
		e = libc.Fetestexcept(0)
		if !(checkexceptall(e, p.e, p.r) != 0) {
			libc.Printf((*int8)(unsafe.Pointer(&[58]int8{'%', 's', ':', '%', 'd', ':', ' ', 'b', 'a', 'd', ' ', 'f', 'p', ' ', 'e', 'x', 'c', 'e', 'p', 't', 'i', 'o', 'n', ':', ' ', '%', 's', ' ', 'f', 'r', 'e', 'x', 'p', 'l', '(', '%', 'L', 'a', ')', '=', '%', 'L', 'a', ',', '%', 'l', 'l', 'd', ',', ' ', 'w', 'a', 'n', 't', ' ', '%', 's', '\x00'})), p.file, p.line, rstr(p.r), p.x, p.y, p.i, estr(p.e))
			libc.Printf((*int8)(unsafe.Pointer(&[9]int8{' ', 'g', 'o', 't', ' ', '%', 's', '\n', '\x00'})), estr(e))
			err++
		}
		d = ulperrl(y, p.y, p.dy)
		if !(checkcr(y, p.y, p.r) != 0) || func() int32 {
			if 8 == 4 {
				return func() int32 {
					if libc.X__FLOAT_BITS(float32(p.x))&uint32(2147483647) < uint32(2139095040) {
						return 1
					} else {
						return 0
					}
				}()
			} else {
				return func() int32 {
					if 8 == 8 {
						return func() int32 {
							if libc.X__DOUBLE_BITS(float64(p.x))&9223372036854775807 < 9218868437227405312 {
								return 1
							} else {
								return 0
							}
						}()
					} else {
						return func() int32 {
							if libc.X__fpclassifyl(p.x) > int32(1) {
								return 1
							} else {
								return 0
							}
						}()
					}
				}()
			}
		}() != 0 && int64(yi) != p.i {
			libc.Printf((*int8)(unsafe.Pointer(&[70]int8{'%', 's', ':', '%', 'd', ':', ' ', '%', 's', ' ', 'f', 'r', 'e', 'x', 'p', 'l', '(', '%', 'L', 'a', ')', ' ', 'w', 'a', 'n', 't', ' ', '%', 'L', 'a', ',', '%', 'l', 'l', 'd', ' ', 'g', 'o', 't', ' ', '%', 'L', 'a', ',', '%', 'd', ' ', 'u', 'l', 'p', 'e', 'r', 'r', ' ', '%', '.', '3', 'f', ' ', '=', ' ', '%', 'a', ' ', '+', ' ', '%', 'a', '\n', '\x00'})), p.file, p.line, rstr(p.r), p.x, p.y, p.i, y, yi, float64(d), float64(d-p.dy), float64(p.dy))
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
