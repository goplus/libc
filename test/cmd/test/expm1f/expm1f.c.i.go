package expm1f

import (
	common "github.com/goplus/libc/test/common"
	unsafe "unsafe"
	libc "github.com/goplus/libc"
	testing "testing"
)

var _cgos_t_expm1f [17]common.Struct_f_f = [17]common.Struct_f_f{common.Struct_f_f{(*int8)(unsafe.Pointer(&[53]int8{'/', 'U', 's', 'e', 'r', 's', '/', 'v', 'f', 'c', '/', 'g', 'o', 'p', 'l', 'u', 's', '/', 'l', 'i', 'b', 'c', '/', 't', 'e', 's', 't', '/', 's', 'r', 'c', '/', 'm', 'a', 't', 'h', '/', 's', 'a', 'n', 'i', 't', 'y', '/', 'e', 'x', 'p', 'm', '1', 'f', '.', 'h', '\x00'})), int32(1), int32(0), float32(-8.0668487548828125), float32(-0.99968624114990234), float32(-0.19532723724842072), int32(0)}, common.Struct_f_f{(*int8)(unsafe.Pointer(&[53]int8{'/', 'U', 's', 'e', 'r', 's', '/', 'v', 'f', 'c', '/', 'g', 'o', 'p', 'l', 'u', 's', '/', 'l', 'i', 'b', 'c', '/', 't', 'e', 's', 't', '/', 's', 'r', 'c', '/', 'm', 'a', 't', 'h', '/', 's', 'a', 'n', 'i', 't', 'y', '/', 'e', 'x', 'p', 'm', '1', 'f', '.', 'h', '\x00'})), int32(2), int32(0), float32(4.3452396392822266), float32(76.110511779785156), float32(-0.2875460684299469), int32(0)}, common.Struct_f_f{(*int8)(unsafe.Pointer(&[53]int8{'/', 'U', 's', 'e', 'r', 's', '/', 'v', 'f', 'c', '/', 'g', 'o', 'p', 'l', 'u', 's', '/', 'l', 'i', 'b', 'c', '/', 't', 'e', 's', 't', '/', 's', 'r', 'c', '/', 'm', 'a', 't', 'h', '/', 's', 'a', 'n', 'i', 't', 'y', '/', 'e', 'x', 'p', 'm', '1', 'f', '.', 'h', '\x00'})), int32(3), int32(0), float32(-8.3814334869384765), float32(-0.99977093935012817), float32(-0.34686920046806335), int32(0)}, common.Struct_f_f{(*int8)(unsafe.Pointer(&[53]int8{'/', 'U', 's', 'e', 'r', 's', '/', 'v', 'f', 'c', '/', 'g', 'o', 'p', 'l', 'u', 's', '/', 'l', 'i', 'b', 'c', '/', 't', 'e', 's', 't', '/', 's', 'r', 'c', '/', 'm', 'a', 't', 'h', '/', 's', 'a', 'n', 'i', 't', 'y', '/', 'e', 'x', 'p', 'm', '1', 'f', '.', 'h', '\x00'})), int32(4), int32(0), float32(-6.5316734313964844), float32(-0.99854344129562378), float32(-0.12819394469261169), int32(0)}, common.Struct_f_f{(*int8)(unsafe.Pointer(&[53]int8{'/', 'U', 's', 'e', 'r', 's', '/', 'v', 'f', 'c', '/', 'g', 'o', 'p', 'l', 'u', 's', '/', 'l', 'i', 'b', 'c', '/', 't', 'e', 's', 't', '/', 's', 'r', 'c', '/', 'm', 'a', 't', 'h', '/', 's', 'a', 'n', 'i', 't', 'y', '/', 'e', 'x', 'p', 'm', '1', 'f', '.', 'h', '\x00'})), int32(5), int32(0), float32(9.2670574188232421), float32(10582.5634765625), float32(0.45962104201316833), int32(0)}, common.Struct_f_f{(*int8)(unsafe.Pointer(&[53]int8{'/', 'U', 's', 'e', 'r', 's', '/', 'v', 'f', 'c', '/', 'g', 'o', 'p', 'l', 'u', 's', '/', 'l', 'i', 'b', 'c', '/', 't', 'e', 's', 't', '/', 's', 'r', 'c', '/', 'm', 'a', 't', 'h', '/', 's', 'a', 'n', 'i', 't', 'y', '/', 'e', 'x', 'p', 'm', '1', 'f', '.', 'h', '\x00'})), int32(6), int32(0), float32(0.66198587417602539), float32(0.93863838911056519), float32(-0.28634780645370483), int32(0)}, common.Struct_f_f{(*int8)(unsafe.Pointer(&[53]int8{'/', 'U', 's', 'e', 'r', 's', '/', 'v', 'f', 'c', '/', 'g', 'o', 'p', 'l', 'u', 's', '/', 'l', 'i', 'b', 'c', '/', 't', 'e', 's', 't', '/', 's', 'r', 'c', '/', 'm', 'a', 't', 'h', '/', 's', 'a', 'n', 'i', 't', 'y', '/', 'e', 'x', 'p', 'm', '1', 'f', '.', 'h', '\x00'})), int32(7), int32(0), float32(-0.40660393238067627), float32(-0.33409211039543152), float32(0.23410017788410187), int32(0)}, common.Struct_f_f{(*int8)(unsafe.Pointer(&[53]int8{'/', 'U', 's', 'e', 'r', 's', '/', 'v', 'f', 'c', '/', 'g', 'o', 'p', 'l', 'u', 's', '/', 'l', 'i', 'b', 'c', '/', 't', 'e', 's', 't', '/', 's', 'r', 'c', '/', 'm', 'a', 't', 'h', '/', 's', 'a', 'n', 'i', 't', 'y', '/', 'e', 'x', 'p', 'm', '1', 'f', '.', 'h', '\x00'})), int32(8), int32(0), float32(0.56175976991653442), float32(0.75375598669052124), float32(-0.11289017647504807), int32(0)}, common.Struct_f_f{(*int8)(unsafe.Pointer(&[53]int8{'/', 'U', 's', 'e', 'r', 's', '/', 'v', 'f', 'c', '/', 'g', 'o', 'p', 'l', 'u', 's', '/', 'l', 'i', 'b', 'c', '/', 't', 'e', 's', 't', '/', 's', 'r', 'c', '/', 'm', 'a', 't', 'h', '/', 's', 'a', 'n', 'i', 't', 'y', '/', 'e', 'x', 'p', 'm', '1', 'f', '.', 'h', '\x00'})), int32(9), int32(0), float32(0.77415227890014648), float32(1.168752908706665), float32(0.49124938249588013), int32(0)}, common.Struct_f_f{(*int8)(unsafe.Pointer(&[53]int8{'/', 'U', 's', 'e', 'r', 's', '/', 'v', 'f', 'c', '/', 'g', 'o', 'p', 'l', 'u', 's', '/', 'l', 'i', 'b', 'c', '/', 't', 'e', 's', 't', '/', 's', 'r', 'c', '/', 'm', 'a', 't', 'h', '/', 's', 'a', 'n', 'i', 't', 'y', '/', 'e', 'x', 'p', 'm', '1', 'f', '.', 'h', '\x00'})), int32(10), int32(0), float32(-0.67876368761062622), float32(-0.49275627732276917), float32(0.20514154434204102), int32(0)}, common.Struct_f_f{(*int8)(unsafe.Pointer(&[54]int8{'/', 'U', 's', 'e', 'r', 's', '/', 'v', 'f', 'c', '/', 'g', 'o', 'p', 'l', 'u', 's', '/', 'l', 'i', 'b', 'c', '/', 't', 'e', 's', 't', '/', 's', 'r', 'c', '/', 'm', 'a', 't', 'h', '/', 's', 'p', 'e', 'c', 'i', 'a', 'l', '/', 'e', 'x', 'p', 'm', '1', 'f', '.', 'h', '\x00'})), int32(1), int32(0), float32(0.0), float32(0.0), float32(0.0), int32(0)}, common.Struct_f_f{(*int8)(unsafe.Pointer(&[54]int8{'/', 'U', 's', 'e', 'r', 's', '/', 'v', 'f', 'c', '/', 'g', 'o', 'p', 'l', 'u', 's', '/', 'l', 'i', 'b', 'c', '/', 't', 'e', 's', 't', '/', 's', 'r', 'c', '/', 'm', 'a', 't', 'h', '/', 's', 'p', 'e', 'c', 'i', 'a', 'l', '/', 'e', 'x', 'p', 'm', '1', 'f', '.', 'h', '\x00'})), int32(2), int32(0), float32(-0.0), float32(-0.0), float32(0.0), int32(0)}, common.Struct_f_f{(*int8)(unsafe.Pointer(&[54]int8{'/', 'U', 's', 'e', 'r', 's', '/', 'v', 'f', 'c', '/', 'g', 'o', 'p', 'l', 'u', 's', '/', 'l', 'i', 'b', 'c', '/', 't', 'e', 's', 't', '/', 's', 'r', 'c', '/', 'm', 'a', 't', 'h', '/', 's', 'p', 'e', 'c', 'i', 'a', 'l', '/', 'e', 'x', 'p', 'm', '1', 'f', '.', 'h', '\x00'})), int32(3), int32(0), float32(1.0), float32(1.7182818651199341), float32(0.30753383040428162), int32(0)}, common.Struct_f_f{(*int8)(unsafe.Pointer(&[54]int8{'/', 'U', 's', 'e', 'r', 's', '/', 'v', 'f', 'c', '/', 'g', 'o', 'p', 'l', 'u', 's', '/', 'l', 'i', 'b', 'c', '/', 't', 'e', 's', 't', '/', 's', 'r', 'c', '/', 'm', 'a', 't', 'h', '/', 's', 'p', 'e', 'c', 'i', 'a', 'l', '/', 'e', 'x', 'p', 'm', '1', 'f', '.', 'h', '\x00'})), int32(4), int32(0), float32(-1.0), float32(-0.63212054967880249), float32(0.15350742638111115), int32(0)}, common.Struct_f_f{(*int8)(unsafe.Pointer(&[54]int8{'/', 'U', 's', 'e', 'r', 's', '/', 'v', 'f', 'c', '/', 'g', 'o', 'p', 'l', 'u', 's', '/', 'l', 'i', 'b', 'c', '/', 't', 'e', 's', 't', '/', 's', 'r', 'c', '/', 'm', 'a', 't', 'h', '/', 's', 'p', 'e', 'c', 'i', 'a', 'l', '/', 'e', 'x', 'p', 'm', '1', 'f', '.', 'h', '\x00'})), int32(5), int32(0), libc.X__builtin_inff(), libc.X__builtin_inff(), float32(0.0), int32(0)}, common.Struct_f_f{(*int8)(unsafe.Pointer(&[54]int8{'/', 'U', 's', 'e', 'r', 's', '/', 'v', 'f', 'c', '/', 'g', 'o', 'p', 'l', 'u', 's', '/', 'l', 'i', 'b', 'c', '/', 't', 'e', 's', 't', '/', 's', 'r', 'c', '/', 'm', 'a', 't', 'h', '/', 's', 'p', 'e', 'c', 'i', 'a', 'l', '/', 'e', 'x', 'p', 'm', '1', 'f', '.', 'h', '\x00'})), int32(6), int32(0), -libc.X__builtin_inff(), float32(-1.0), float32(0.0), int32(0)}, common.Struct_f_f{(*int8)(unsafe.Pointer(&[54]int8{'/', 'U', 's', 'e', 'r', 's', '/', 'v', 'f', 'c', '/', 'g', 'o', 'p', 'l', 'u', 's', '/', 'l', 'i', 'b', 'c', '/', 't', 'e', 's', 't', '/', 's', 'r', 'c', '/', 'm', 'a', 't', 'h', '/', 's', 'p', 'e', 'c', 'i', 'a', 'l', '/', 'e', 'x', 'p', 'm', '1', 'f', '.', 'h', '\x00'})), int32(7), int32(0), libc.X__builtin_nanf((*int8)(unsafe.Pointer(&[1]int8{'\x00'}))), libc.X__builtin_nanf((*int8)(unsafe.Pointer(&[1]int8{'\x00'}))), float32(0.0), int32(0)}}

func _cgo_main() int32 {
	var y float32
	var d float32
	var e int32
	var i int32
	var err int32 = int32(0)
	var p *common.Struct_f_f
	for i = int32(0); uint64(i) < 17; i++ {
		p = (*common.Struct_f_f)(unsafe.Pointer(uintptr(unsafe.Pointer((*common.Struct_f_f)(unsafe.Pointer(&_cgos_t_expm1f)))) + uintptr(i)*32))
		if p.R < int32(0) {
			continue
		}
		libc.Fesetround(p.R)
		libc.Feclearexcept(int32(0))
		y = libc.Expm1f(p.X)
		e = libc.Fetestexcept(0)
		if !(common.Checkexcept(e, p.E, p.R) != 0) {
			libc.Printf((*int8)(unsafe.Pointer(&[51]int8{'%', 's', ':', '%', 'd', ':', ' ', 'b', 'a', 'd', ' ', 'f', 'p', ' ', 'e', 'x', 'c', 'e', 'p', 't', 'i', 'o', 'n', ':', ' ', '%', 's', ' ', 'e', 'x', 'p', 'm', '1', 'f', '(', '%', 'a', ')', '=', '%', 'a', ',', ' ', 'w', 'a', 'n', 't', ' ', '%', 's', '\x00'})), p.File, p.Line, common.Rstr(p.R), float64(p.X), float64(p.Y), common.Estr(p.E))
			libc.Printf((*int8)(unsafe.Pointer(&[9]int8{' ', 'g', 'o', 't', ' ', '%', 's', '\n', '\x00'})), common.Estr(e))
			err++
		}
		d = common.Ulperrf(y, p.Y, p.Dy)
		if !(common.Checkulp(d, p.R) != 0) {
			libc.Printf((*int8)(unsafe.Pointer(&[59]int8{'%', 's', ':', '%', 'd', ':', ' ', '%', 's', ' ', 'e', 'x', 'p', 'm', '1', 'f', '(', '%', 'a', ')', ' ', 'w', 'a', 'n', 't', ' ', '%', 'a', ' ', 'g', 'o', 't', ' ', '%', 'a', ' ', 'u', 'l', 'p', 'e', 'r', 'r', ' ', '%', '.', '3', 'f', ' ', '=', ' ', '%', 'a', ' ', '+', ' ', '%', 'a', '\n', '\x00'})), p.File, p.Line, common.Rstr(p.R), float64(p.X), float64(p.Y), float64(y), float64(d), float64(d-p.Dy), float64(p.Dy))
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
