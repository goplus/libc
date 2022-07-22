package libc

import unsafe "unsafe"

var init_jk_cgo18___rem_pio2_large [4]int32 = [4]int32{int32(3), int32(4), int32(4), int32(6)}
var ipio2_cgo19___rem_pio2_large [66]int32 = [66]int32{int32(10680707), int32(7228996), int32(1387004), int32(2578385), int32(16069853), int32(12639074), int32(9804092), int32(4427841), int32(16666979), int32(11263675), int32(12935607), int32(2387514), int32(4345298), int32(14681673), int32(3074569), int32(13734428), int32(16653803), int32(1880361), int32(10960616), int32(8533493), int32(3062596), int32(8710556), int32(7349940), int32(6258241), int32(3772886), int32(3769171), int32(3798172), int32(8675211), int32(12450088), int32(3874808), int32(9961438), int32(366607), int32(15675153), int32(9132554), int32(7151469), int32(3571407), int32(2607881), int32(12013382), int32(4155038), int32(6285869), int32(7677882), int32(13102053), int32(15825725), int32(473591), int32(9065106), int32(15363067), int32(6271263), int32(9264392), int32(5636912), int32(4652155), int32(7056368), int32(13614112), int32(10155062), int32(1944035), int32(9527646), int32(15080200), int32(6658437), int32(6231200), int32(6832269), int32(16767104), int32(5075751), int32(3212806), int32(1398474), int32(7579849), int32(6349435), int32(12618859)}
var PIo2_cgo20___rem_pio2_large [8]float64 = [8]float64{1.5707962512969971, 7.5497894158615964e-8, 5.3903025299577648e-15, 3.2820034158079129e-22, 1.2706557530806761e-29, 1.2293330898111133e-36, 2.7337005381646456e-44, 2.1674168387780482e-51}

func __rem_pio2_large(x *float64, y *float64, e0 int32, nx int32, prec int32) int32 {
	var jz int32
	var jx int32
	var jv int32
	var jp int32
	var jk int32
	var carry int32
	var n int32
	var iq [20]int32
	var i int32
	var j int32
	var k int32
	var m int32
	var q0 int32
	var ih int32
	var z float64
	var fw float64
	var f [20]float64
	var fq [20]float64
	var q [20]float64
	jk = *(*int32)(unsafe.Pointer(uintptr(unsafe.Pointer((*int32)(unsafe.Pointer(&init_jk_cgo18___rem_pio2_large)))) + uintptr(prec)*4))
	jp = jk
	jx = nx - int32(1)
	jv = (e0 - int32(3)) / int32(24)
	if jv < int32(0) {
		jv = int32(0)
	}
	q0 = e0 - int32(24)*(jv+int32(1))
	j = jv - jx
	m = jx + jk
	for i = int32(0); i <= m; func() int32 {
		i++
		return func() (_cgo_ret int32) {
			_cgo_addr := &j
			_cgo_ret = *_cgo_addr
			*_cgo_addr++
			return
		}()
	}() {
		*(*float64)(unsafe.Pointer(uintptr(unsafe.Pointer((*float64)(unsafe.Pointer(&f)))) + uintptr(i)*8)) = func() float64 {
			if j < int32(0) {
				return 0
			} else {
				return float64(*(*int32)(unsafe.Pointer(uintptr(unsafe.Pointer((*int32)(unsafe.Pointer(&ipio2_cgo19___rem_pio2_large)))) + uintptr(j)*4)))
			}
		}()
	}
	for i = int32(0); i <= jk; i++ {
		for func() float64 {
			j = int32(0)
			return func() (_cgo_ret float64) {
				_cgo_addr := &fw
				*_cgo_addr = float64(0)
				return *_cgo_addr
			}()
		}(); j <= jx; j++ {
			fw += *(*float64)(unsafe.Pointer(uintptr(unsafe.Pointer(x)) + uintptr(j)*8)) * *(*float64)(unsafe.Pointer(uintptr(unsafe.Pointer((*float64)(unsafe.Pointer(&f)))) + uintptr(jx+i-j)*8))
		}
		*(*float64)(unsafe.Pointer(uintptr(unsafe.Pointer((*float64)(unsafe.Pointer(&q)))) + uintptr(i)*8)) = fw
	}
	jz = jk
recompute:
	for func() float64 {
		func() int32 {
			i = int32(0)
			return func() (_cgo_ret int32) {
				_cgo_addr := &j
				*_cgo_addr = jz
				return *_cgo_addr
			}()
		}()
		return func() (_cgo_ret float64) {
			_cgo_addr := &z
			*_cgo_addr = *(*float64)(unsafe.Pointer(uintptr(unsafe.Pointer((*float64)(unsafe.Pointer(&q)))) + uintptr(jz)*8))
			return *_cgo_addr
		}()
	}(); j > int32(0); func() int32 {
		i++
		return func() (_cgo_ret int32) {
			_cgo_addr := &j
			_cgo_ret = *_cgo_addr
			*_cgo_addr--
			return
		}()
	}() {
		fw = float64(int32(5.9604644775390625e-8 * z))
		*(*int32)(unsafe.Pointer(uintptr(unsafe.Pointer((*int32)(unsafe.Pointer(&iq)))) + uintptr(i)*4)) = int32(z - 16777216*fw)
		z = *(*float64)(unsafe.Pointer(uintptr(unsafe.Pointer((*float64)(unsafe.Pointer(&q)))) + uintptr(j-int32(1))*8)) + fw
	}
	z = Scalbn(z, q0)
	z -= 8 * Floor(z*0.125)
	n = int32(z)
	z -= float64(n)
	ih = int32(0)
	if q0 > int32(0) {
		i = *(*int32)(unsafe.Pointer(uintptr(unsafe.Pointer((*int32)(unsafe.Pointer(&iq)))) + uintptr(jz-int32(1))*4)) >> (int32(24) - q0)
		n += i
		*(*int32)(unsafe.Pointer(uintptr(unsafe.Pointer((*int32)(unsafe.Pointer(&iq)))) + uintptr(jz-int32(1))*4)) -= i << (int32(24) - q0)
		ih = *(*int32)(unsafe.Pointer(uintptr(unsafe.Pointer((*int32)(unsafe.Pointer(&iq)))) + uintptr(jz-int32(1))*4)) >> (int32(23) - q0)
	} else if q0 == int32(0) {
		ih = *(*int32)(unsafe.Pointer(uintptr(unsafe.Pointer((*int32)(unsafe.Pointer(&iq)))) + uintptr(jz-int32(1))*4)) >> int32(23)
	} else if z >= 0.5 {
		ih = int32(2)
	}
	if ih > int32(0) {
		n += int32(1)
		carry = int32(0)
		for i = int32(0); i < jz; i++ {
			j = *(*int32)(unsafe.Pointer(uintptr(unsafe.Pointer((*int32)(unsafe.Pointer(&iq)))) + uintptr(i)*4))
			if carry == int32(0) {
				if j != int32(0) {
					carry = int32(1)
					*(*int32)(unsafe.Pointer(uintptr(unsafe.Pointer((*int32)(unsafe.Pointer(&iq)))) + uintptr(i)*4)) = int32(16777216) - j
				}
			} else {
				*(*int32)(unsafe.Pointer(uintptr(unsafe.Pointer((*int32)(unsafe.Pointer(&iq)))) + uintptr(i)*4)) = int32(16777215) - j
			}
		}
		if q0 > int32(0) {
			switch q0 {
			case int32(1):
				*(*int32)(unsafe.Pointer(uintptr(unsafe.Pointer((*int32)(unsafe.Pointer(&iq)))) + uintptr(jz-int32(1))*4)) &= int32(8388607)
				break
			case int32(2):
				*(*int32)(unsafe.Pointer(uintptr(unsafe.Pointer((*int32)(unsafe.Pointer(&iq)))) + uintptr(jz-int32(1))*4)) &= int32(4194303)
				break
			}
		}
		if ih == int32(2) {
			z = 1 - z
			if carry != int32(0) {
				z -= Scalbn(1, q0)
			}
		}
	}
	if z == 0 {
		j = int32(0)
		for i = jz - int32(1); i >= jk; i-- {
			j |= *(*int32)(unsafe.Pointer(uintptr(unsafe.Pointer((*int32)(unsafe.Pointer(&iq)))) + uintptr(i)*4))
		}
		if j == int32(0) {
			for k = int32(1); *(*int32)(unsafe.Pointer(uintptr(unsafe.Pointer((*int32)(unsafe.Pointer(&iq)))) + uintptr(jk-k)*4)) == int32(0); k++ {
			}
			for i = jz + int32(1); i <= jz+k; i++ {
				*(*float64)(unsafe.Pointer(uintptr(unsafe.Pointer((*float64)(unsafe.Pointer(&f)))) + uintptr(jx+i)*8)) = float64(*(*int32)(unsafe.Pointer(uintptr(unsafe.Pointer((*int32)(unsafe.Pointer(&ipio2_cgo19___rem_pio2_large)))) + uintptr(jv+i)*4)))
				for func() float64 {
					j = int32(0)
					return func() (_cgo_ret float64) {
						_cgo_addr := &fw
						*_cgo_addr = float64(0)
						return *_cgo_addr
					}()
				}(); j <= jx; j++ {
					fw += *(*float64)(unsafe.Pointer(uintptr(unsafe.Pointer(x)) + uintptr(j)*8)) * *(*float64)(unsafe.Pointer(uintptr(unsafe.Pointer((*float64)(unsafe.Pointer(&f)))) + uintptr(jx+i-j)*8))
				}
				*(*float64)(unsafe.Pointer(uintptr(unsafe.Pointer((*float64)(unsafe.Pointer(&q)))) + uintptr(i)*8)) = fw
			}
			jz += k
			goto recompute
		}
	}
	if z == 0 {
		jz -= int32(1)
		q0 -= int32(24)
		for *(*int32)(unsafe.Pointer(uintptr(unsafe.Pointer((*int32)(unsafe.Pointer(&iq)))) + uintptr(jz)*4)) == int32(0) {
			jz--
			q0 -= int32(24)
		}
	} else {
		z = Scalbn(z, -q0)
		if z >= 16777216 {
			fw = float64(int32(5.9604644775390625e-8 * z))
			*(*int32)(unsafe.Pointer(uintptr(unsafe.Pointer((*int32)(unsafe.Pointer(&iq)))) + uintptr(jz)*4)) = int32(z - 16777216*fw)
			jz += int32(1)
			q0 += int32(24)
			*(*int32)(unsafe.Pointer(uintptr(unsafe.Pointer((*int32)(unsafe.Pointer(&iq)))) + uintptr(jz)*4)) = int32(fw)
		} else {
			*(*int32)(unsafe.Pointer(uintptr(unsafe.Pointer((*int32)(unsafe.Pointer(&iq)))) + uintptr(jz)*4)) = int32(z)
		}
	}
	fw = Scalbn(1, q0)
	for i = jz; i >= int32(0); i-- {
		*(*float64)(unsafe.Pointer(uintptr(unsafe.Pointer((*float64)(unsafe.Pointer(&q)))) + uintptr(i)*8)) = fw * float64(*(*int32)(unsafe.Pointer(uintptr(unsafe.Pointer((*int32)(unsafe.Pointer(&iq)))) + uintptr(i)*4)))
		fw *= float64(5.9604644775390625e-8)
	}
	for i = jz; i >= int32(0); i-- {
		for func() int32 {
			fw = float64(0)
			return func() (_cgo_ret int32) {
				_cgo_addr := &k
				*_cgo_addr = int32(0)
				return *_cgo_addr
			}()
		}(); k <= jp && k <= jz-i; k++ {
			fw += *(*float64)(unsafe.Pointer(uintptr(unsafe.Pointer((*float64)(unsafe.Pointer(&PIo2_cgo20___rem_pio2_large)))) + uintptr(k)*8)) * *(*float64)(unsafe.Pointer(uintptr(unsafe.Pointer((*float64)(unsafe.Pointer(&q)))) + uintptr(i+k)*8))
		}
		*(*float64)(unsafe.Pointer(uintptr(unsafe.Pointer((*float64)(unsafe.Pointer(&fq)))) + uintptr(jz-i)*8)) = fw
	}
	switch prec {
	case int32(0):
		fw = float64(0)
		for i = jz; i >= int32(0); i-- {
			fw += *(*float64)(unsafe.Pointer(uintptr(unsafe.Pointer((*float64)(unsafe.Pointer(&fq)))) + uintptr(i)*8))
		}
		*(*float64)(unsafe.Pointer(uintptr(unsafe.Pointer(y)) + uintptr(int32(0))*8)) = func() float64 {
			if ih == int32(0) {
				return fw
			} else {
				return -fw
			}
		}()
		break
	case int32(1):
		fallthrough
	case int32(2):
		fw = float64(0)
		for i = jz; i >= int32(0); i-- {
			fw += *(*float64)(unsafe.Pointer(uintptr(unsafe.Pointer((*float64)(unsafe.Pointer(&fq)))) + uintptr(i)*8))
		}
		fw = float64(fw)
		*(*float64)(unsafe.Pointer(uintptr(unsafe.Pointer(y)) + uintptr(int32(0))*8)) = func() float64 {
			if ih == int32(0) {
				return fw
			} else {
				return -fw
			}
		}()
		fw = *(*float64)(unsafe.Pointer(uintptr(unsafe.Pointer((*float64)(unsafe.Pointer(&fq)))) + uintptr(int32(0))*8)) - fw
		for i = int32(1); i <= jz; i++ {
			fw += *(*float64)(unsafe.Pointer(uintptr(unsafe.Pointer((*float64)(unsafe.Pointer(&fq)))) + uintptr(i)*8))
		}
		*(*float64)(unsafe.Pointer(uintptr(unsafe.Pointer(y)) + uintptr(int32(1))*8)) = func() float64 {
			if ih == int32(0) {
				return fw
			} else {
				return -fw
			}
		}()
		break
	case int32(3):
		for i = jz; i > int32(0); i-- {
			fw = *(*float64)(unsafe.Pointer(uintptr(unsafe.Pointer((*float64)(unsafe.Pointer(&fq)))) + uintptr(i-int32(1))*8)) + *(*float64)(unsafe.Pointer(uintptr(unsafe.Pointer((*float64)(unsafe.Pointer(&fq)))) + uintptr(i)*8))
			*(*float64)(unsafe.Pointer(uintptr(unsafe.Pointer((*float64)(unsafe.Pointer(&fq)))) + uintptr(i)*8)) += *(*float64)(unsafe.Pointer(uintptr(unsafe.Pointer((*float64)(unsafe.Pointer(&fq)))) + uintptr(i-int32(1))*8)) - fw
			*(*float64)(unsafe.Pointer(uintptr(unsafe.Pointer((*float64)(unsafe.Pointer(&fq)))) + uintptr(i-int32(1))*8)) = fw
		}
		for i = jz; i > int32(1); i-- {
			fw = *(*float64)(unsafe.Pointer(uintptr(unsafe.Pointer((*float64)(unsafe.Pointer(&fq)))) + uintptr(i-int32(1))*8)) + *(*float64)(unsafe.Pointer(uintptr(unsafe.Pointer((*float64)(unsafe.Pointer(&fq)))) + uintptr(i)*8))
			*(*float64)(unsafe.Pointer(uintptr(unsafe.Pointer((*float64)(unsafe.Pointer(&fq)))) + uintptr(i)*8)) += *(*float64)(unsafe.Pointer(uintptr(unsafe.Pointer((*float64)(unsafe.Pointer(&fq)))) + uintptr(i-int32(1))*8)) - fw
			*(*float64)(unsafe.Pointer(uintptr(unsafe.Pointer((*float64)(unsafe.Pointer(&fq)))) + uintptr(i-int32(1))*8)) = fw
		}
		for func() int32 {
			fw = float64(0)
			return func() (_cgo_ret int32) {
				_cgo_addr := &i
				*_cgo_addr = jz
				return *_cgo_addr
			}()
		}(); i >= int32(2); i-- {
			fw += *(*float64)(unsafe.Pointer(uintptr(unsafe.Pointer((*float64)(unsafe.Pointer(&fq)))) + uintptr(i)*8))
		}
		if ih == int32(0) {
			*(*float64)(unsafe.Pointer(uintptr(unsafe.Pointer(y)) + uintptr(int32(0))*8)) = *(*float64)(unsafe.Pointer(uintptr(unsafe.Pointer((*float64)(unsafe.Pointer(&fq)))) + uintptr(int32(0))*8))
			*(*float64)(unsafe.Pointer(uintptr(unsafe.Pointer(y)) + uintptr(int32(1))*8)) = *(*float64)(unsafe.Pointer(uintptr(unsafe.Pointer((*float64)(unsafe.Pointer(&fq)))) + uintptr(int32(1))*8))
			*(*float64)(unsafe.Pointer(uintptr(unsafe.Pointer(y)) + uintptr(int32(2))*8)) = fw
		} else {
			*(*float64)(unsafe.Pointer(uintptr(unsafe.Pointer(y)) + uintptr(int32(0))*8)) = -*(*float64)(unsafe.Pointer(uintptr(unsafe.Pointer((*float64)(unsafe.Pointer(&fq)))) + uintptr(int32(0))*8))
			*(*float64)(unsafe.Pointer(uintptr(unsafe.Pointer(y)) + uintptr(int32(1))*8)) = -*(*float64)(unsafe.Pointer(uintptr(unsafe.Pointer((*float64)(unsafe.Pointer(&fq)))) + uintptr(int32(1))*8))
			*(*float64)(unsafe.Pointer(uintptr(unsafe.Pointer(y)) + uintptr(int32(2))*8)) = -fw
		}
	}
	return n & int32(7)
}
