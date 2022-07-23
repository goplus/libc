package libc

import unsafe "unsafe"

func casemap_cgos__towctrans(c uint32, dir int32) int32 {
	var b uint32
	var x uint32
	var y uint32
	var v uint32
	var rt uint32
	var xb uint32
	var xn uint32
	var r int32
	var rd int32
	var c0 int32 = int32(c)
	if c >= uint32(131072) {
		return int32(c)
	}
	b = c >> int32(8)
	c &= uint32(255)
	x = c / uint32(3)
	y = c % uint32(3)
	v = uint32(*(*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer((*uint8)(unsafe.Pointer(&tab_cgos__towctrans)))) + uintptr(uint32(int32(*(*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer((*uint8)(unsafe.Pointer(&tab_cgos__towctrans)))) + uintptr(b))))*int32(86))+x))))
	v = v * uint32(*(*int32)(unsafe.Pointer(uintptr(unsafe.Pointer((*int32)(unsafe.Pointer(&mt_cgos__towctrans)))) + uintptr(y)*4))) >> int32(11) % uint32(6)
	r = *(*int32)(unsafe.Pointer(uintptr(unsafe.Pointer((*int32)(unsafe.Pointer(&rules_cgos__towctrans)))) + uintptr(uint32(*(*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer((*uint8)(unsafe.Pointer(&rulebases_cgos__towctrans)))) + uintptr(b))))+v)*4))
	rt = uint32(r & int32(255))
	rd = r >> int32(8)
	if rt < uint32(2) {
		return int32(uint32(c0) + uint32(rd)&-(rt^uint32(dir)))
	}
	xn = uint32(rd & int32(255))
	xb = uint32(rd) >> int32(8)
	for xn != 0 {
		var try uint32 = uint32(*(*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer((*uint8)(unsafe.Pointer(&*(*[2]uint8)(unsafe.Pointer(uintptr(unsafe.Pointer((*[2]uint8)(unsafe.Pointer(&exceptions_cgos__towctrans)))) + uintptr(xb+xn/uint32(2))*2)))))) + uintptr(int32(0)))))
		if try == c {
			r = *(*int32)(unsafe.Pointer(uintptr(unsafe.Pointer((*int32)(unsafe.Pointer(&rules_cgos__towctrans)))) + uintptr(*(*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer((*uint8)(unsafe.Pointer(&*(*[2]uint8)(unsafe.Pointer(uintptr(unsafe.Pointer((*[2]uint8)(unsafe.Pointer(&exceptions_cgos__towctrans)))) + uintptr(xb+xn/uint32(2))*2)))))) + uintptr(int32(1)))))*4))
			rt = uint32(r & int32(255))
			rd = r >> int32(8)
			if rt < uint32(2) {
				return int32(uint32(c0) + uint32(rd)&-(rt^uint32(dir)))
			}
			return c0 + func() int32 {
				if dir != 0 {
					return -1
				} else {
					return int32(1)
				}
			}()
		} else if try > c {
			xn /= uint32(2)
		} else {
			xb += xn / uint32(2)
			xn -= xn / uint32(2)
		}
	}
	return c0
}

var mt_cgos__towctrans [3]int32 = [3]int32{int32(2048), int32(342), int32(57)}

func towlower(wc uint32) uint32 {
	return uint32(casemap_cgos__towctrans(wc, int32(0)))
}
func towupper(wc uint32) uint32 {
	return uint32(casemap_cgos__towctrans(wc, int32(1)))
}
func __towupper_l(c uint32, l *struct___locale_struct) uint32 {
	return towupper(c)
}
func __towlower_l(c uint32, l *struct___locale_struct) uint32 {
	return towlower(c)
}
