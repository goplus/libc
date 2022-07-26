package libc

import unsafe "unsafe"

func _cgos_scanexp_floatscan(f *Struct__IO_FILE, pok int32) int64 {
	var c int32
	var x int32
	var y int64
	var neg int32 = int32(0)
	c = func() int32 {
		if uintptr(unsafe.Pointer(f.Rpos)) != uintptr(unsafe.Pointer(f.Shend)) {
			return int32(*func() (_cgo_ret *uint8) {
				_cgo_addr := &f.Rpos
				_cgo_ret = *_cgo_addr
				*(*uintptr)(unsafe.Pointer(_cgo_addr))++
				return
			}())
		} else {
			return __shgetc(f)
		}
	}()
	if c == '+' || c == '-' {
		neg = func() int32 {
			if c == '-' {
				return 1
			} else {
				return 0
			}
		}()
		c = func() int32 {
			if uintptr(unsafe.Pointer(f.Rpos)) != uintptr(unsafe.Pointer(f.Shend)) {
				return int32(*func() (_cgo_ret *uint8) {
					_cgo_addr := &f.Rpos
					_cgo_ret = *_cgo_addr
					*(*uintptr)(unsafe.Pointer(_cgo_addr))++
					return
				}())
			} else {
				return __shgetc(f)
			}
		}()
		if uint32(c-'0') >= uint32(10) && pok != 0 {
			if f.Shlim >= int64(0) {
				func() int {
					_ = func() (_cgo_ret *uint8) {
						_cgo_addr := &f.Rpos
						_cgo_ret = *_cgo_addr
						*(*uintptr)(unsafe.Pointer(_cgo_addr))--
						return
					}()
					return 0
				}()
			} else {
				func() int {
					_ = int32(0)
					return 0
				}()
			}
		}
	}
	if uint32(c-'0') >= uint32(10) {
		if f.Shlim >= int64(0) {
			func() int {
				_ = func() (_cgo_ret *uint8) {
					_cgo_addr := &f.Rpos
					_cgo_ret = *_cgo_addr
					*(*uintptr)(unsafe.Pointer(_cgo_addr))--
					return
				}()
				return 0
			}()
		} else {
			func() int {
				_ = int32(0)
				return 0
			}()
		}
		return -9223372036854775808
	}
	for x = int32(0); uint32(c-'0') < uint32(10) && x < 214748364; c = func() int32 {
		if uintptr(unsafe.Pointer(f.Rpos)) != uintptr(unsafe.Pointer(f.Shend)) {
			return int32(*func() (_cgo_ret *uint8) {
				_cgo_addr := &f.Rpos
				_cgo_ret = *_cgo_addr
				*(*uintptr)(unsafe.Pointer(_cgo_addr))++
				return
			}())
		} else {
			return __shgetc(f)
		}
	}() {
		x = int32(10)*x + c - '0'
	}
	for y = int64(x); uint32(c-'0') < uint32(10) && y < 92233720368547758; c = func() int32 {
		if uintptr(unsafe.Pointer(f.Rpos)) != uintptr(unsafe.Pointer(f.Shend)) {
			return int32(*func() (_cgo_ret *uint8) {
				_cgo_addr := &f.Rpos
				_cgo_ret = *_cgo_addr
				*(*uintptr)(unsafe.Pointer(_cgo_addr))++
				return
			}())
		} else {
			return __shgetc(f)
		}
	}() {
		y = int64(10)*y + int64(c) - int64('0')
	}
	for ; uint32(c-'0') < uint32(10); c = func() int32 {
		if uintptr(unsafe.Pointer(f.Rpos)) != uintptr(unsafe.Pointer(f.Shend)) {
			return int32(*func() (_cgo_ret *uint8) {
				_cgo_addr := &f.Rpos
				_cgo_ret = *_cgo_addr
				*(*uintptr)(unsafe.Pointer(_cgo_addr))++
				return
			}())
		} else {
			return __shgetc(f)
		}
	}() {
	}
	if f.Shlim >= int64(0) {
		func() int {
			_ = func() (_cgo_ret *uint8) {
				_cgo_addr := &f.Rpos
				_cgo_ret = *_cgo_addr
				*(*uintptr)(unsafe.Pointer(_cgo_addr))--
				return
			}()
			return 0
		}()
	} else {
		func() int {
			_ = int32(0)
			return 0
		}()
	}
	return func() int64 {
		if neg != 0 {
			return -y
		} else {
			return y
		}
	}()
}
func _cgos_decfloat_floatscan(f *Struct__IO_FILE, c int32, bits int32, emin int32, sign int32, pok int32) float64 {
	var x [128]uint32
	var i int32
	var j int32
	var k int32
	var a int32
	var z int32
	var lrp int64 = int64(0)
	var dc int64 = int64(0)
	var e10 int64 = int64(0)
	var lnz int32 = int32(0)
	var gotdig int32 = int32(0)
	var gotrad int32 = int32(0)
	var rp int32
	var e2 int32
	var emax int32 = -emin - bits + int32(3)
	var denormal int32 = int32(0)
	var y float64
	var frac float64 = float64(int32(0))
	var bias float64 = float64(int32(0))
	j = int32(0)
	k = int32(0)
	for ; c == '0'; c = func() int32 {
		if uintptr(unsafe.Pointer(f.Rpos)) != uintptr(unsafe.Pointer(f.Shend)) {
			return int32(*func() (_cgo_ret *uint8) {
				_cgo_addr := &f.Rpos
				_cgo_ret = *_cgo_addr
				*(*uintptr)(unsafe.Pointer(_cgo_addr))++
				return
			}())
		} else {
			return __shgetc(f)
		}
	}() {
		gotdig = int32(1)
	}
	if c == '.' {
		gotrad = int32(1)
		for c = func() int32 {
			if uintptr(unsafe.Pointer(f.Rpos)) != uintptr(unsafe.Pointer(f.Shend)) {
				return int32(*func() (_cgo_ret *uint8) {
					_cgo_addr := &f.Rpos
					_cgo_ret = *_cgo_addr
					*(*uintptr)(unsafe.Pointer(_cgo_addr))++
					return
				}())
			} else {
				return __shgetc(f)
			}
		}(); c == '0'; c = func() int32 {
			if uintptr(unsafe.Pointer(f.Rpos)) != uintptr(unsafe.Pointer(f.Shend)) {
				return int32(*func() (_cgo_ret *uint8) {
					_cgo_addr := &f.Rpos
					_cgo_ret = *_cgo_addr
					*(*uintptr)(unsafe.Pointer(_cgo_addr))++
					return
				}())
			} else {
				return __shgetc(f)
			}
		}() {
			func() int64 {
				gotdig = int32(1)
				return func() (_cgo_ret int64) {
					_cgo_addr := &lrp
					_cgo_ret = *_cgo_addr
					*_cgo_addr--
					return
				}()
			}()
		}
	}
	*(*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer((*uint32)(unsafe.Pointer(&x)))) + uintptr(int32(0))*4)) = uint32(0)
	for ; uint32(c-'0') < uint32(10) || c == '.'; c = func() int32 {
		if uintptr(unsafe.Pointer(f.Rpos)) != uintptr(unsafe.Pointer(f.Shend)) {
			return int32(*func() (_cgo_ret *uint8) {
				_cgo_addr := &f.Rpos
				_cgo_ret = *_cgo_addr
				*(*uintptr)(unsafe.Pointer(_cgo_addr))++
				return
			}())
		} else {
			return __shgetc(f)
		}
	}() {
		if c == '.' {
			if gotrad != 0 {
				break
			}
			gotrad = int32(1)
			lrp = dc
		} else if k < 125 {
			dc++
			if c != '0' {
				lnz = int32(dc)
			}
			if j != 0 {
				*(*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer((*uint32)(unsafe.Pointer(&x)))) + uintptr(k)*4)) = *(*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer((*uint32)(unsafe.Pointer(&x)))) + uintptr(k)*4))*uint32(10) + uint32(c) - uint32('0')
			} else {
				*(*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer((*uint32)(unsafe.Pointer(&x)))) + uintptr(k)*4)) = uint32(c - '0')
			}
			if func() (_cgo_ret int32) {
				_cgo_addr := &j
				*_cgo_addr++
				return *_cgo_addr
			}() == int32(9) {
				k++
				j = int32(0)
			}
			gotdig = int32(1)
		} else {
			dc++
			if c != '0' {
				lnz = 1116
				*(*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer((*uint32)(unsafe.Pointer(&x)))) + uintptr(124)*4)) |= uint32(1)
			}
		}
	}
	if !(gotrad != 0) {
		lrp = dc
	}
	if gotdig != 0 && c|int32(32) == 'e' {
		e10 = _cgos_scanexp_floatscan(f, pok)
		if e10 == -9223372036854775808 {
			if pok != 0 {
				if f.Shlim >= int64(0) {
					func() int {
						_ = func() (_cgo_ret *uint8) {
							_cgo_addr := &f.Rpos
							_cgo_ret = *_cgo_addr
							*(*uintptr)(unsafe.Pointer(_cgo_addr))--
							return
						}()
						return 0
					}()
				} else {
					func() int {
						_ = int32(0)
						return 0
					}()
				}
			} else {
				__shlim(f, int64(0))
				return float64(int32(0))
			}
			e10 = int64(0)
		}
		lrp += e10
	} else if c >= int32(0) {
		if f.Shlim >= int64(0) {
			func() int {
				_ = func() (_cgo_ret *uint8) {
					_cgo_addr := &f.Rpos
					_cgo_ret = *_cgo_addr
					*(*uintptr)(unsafe.Pointer(_cgo_addr))--
					return
				}()
				return 0
			}()
		} else {
			func() int {
				_ = int32(0)
				return 0
			}()
		}
	}
	if !(gotdig != 0) {
		*__errno_location() = int32(22)
		__shlim(f, int64(0))
		return float64(int32(0))
	}
	if !(*(*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer((*uint32)(unsafe.Pointer(&x)))) + uintptr(int32(0))*4)) != 0) {
		return float64(float64(sign) * 0)
	}
	if lrp == dc && dc < int64(10) && (bits > int32(30) || *(*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer((*uint32)(unsafe.Pointer(&x)))) + uintptr(int32(0))*4))>>bits == uint32(0)) {
		return float64(sign) * float64(*(*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer((*uint32)(unsafe.Pointer(&x)))) + uintptr(int32(0))*4)))
	}
	if lrp > int64(-emin/int32(2)) {
		*__errno_location() = int32(34)
		return float64(sign) * 1.79769313486231570815e+308 * 1.79769313486231570815e+308
	}
	if lrp < int64(emin-106) {
		*__errno_location() = int32(34)
		return float64(sign) * 2.22507385850720138309e-308 * 2.22507385850720138309e-308
	}
	if j != 0 {
		for ; j < int32(9); j++ {
			*(*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer((*uint32)(unsafe.Pointer(&x)))) + uintptr(k)*4)) *= uint32(10)
		}
		k++
		j = int32(0)
	}
	a = int32(0)
	z = k
	e2 = int32(0)
	rp = int32(lrp)
	if lnz < int32(9) && lnz <= rp && rp < int32(18) {
		if rp == int32(9) {
			return float64(sign) * float64(*(*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer((*uint32)(unsafe.Pointer(&x)))) + uintptr(int32(0))*4)))
		}
		if rp < int32(9) {
			return float64(sign) * float64(*(*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer((*uint32)(unsafe.Pointer(&x)))) + uintptr(int32(0))*4))) / float64(*(*int32)(unsafe.Pointer(uintptr(unsafe.Pointer((*int32)(unsafe.Pointer(&_cgos_p10s_floatscan)))) + uintptr(int32(8)-rp)*4)))
		}
		var bitlim int32 = bits - int32(3)*int32(rp-int32(9))
		if bitlim > int32(30) || *(*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer((*uint32)(unsafe.Pointer(&x)))) + uintptr(int32(0))*4))>>bitlim == uint32(0) {
			return float64(sign) * float64(*(*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer((*uint32)(unsafe.Pointer(&x)))) + uintptr(int32(0))*4))) * float64(*(*int32)(unsafe.Pointer(uintptr(unsafe.Pointer((*int32)(unsafe.Pointer(&_cgos_p10s_floatscan)))) + uintptr(rp-int32(10))*4)))
		}
	}
	for ; !(*(*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer((*uint32)(unsafe.Pointer(&x)))) + uintptr(z-int32(1))*4)) != 0); z-- {
	}
	if rp%int32(9) != 0 {
		var rpm9 int32 = func() int32 {
			if rp >= int32(0) {
				return rp % int32(9)
			} else {
				return rp%int32(9) + int32(9)
			}
		}()
		var p10 int32 = *(*int32)(unsafe.Pointer(uintptr(unsafe.Pointer((*int32)(unsafe.Pointer(&_cgos_p10s_floatscan)))) + uintptr(int32(8)-rpm9)*4))
		var carry uint32 = uint32(0)
		for k = a; k != z; k++ {
			var tmp uint32 = *(*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer((*uint32)(unsafe.Pointer(&x)))) + uintptr(k)*4)) % uint32(p10)
			*(*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer((*uint32)(unsafe.Pointer(&x)))) + uintptr(k)*4)) = *(*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer((*uint32)(unsafe.Pointer(&x)))) + uintptr(k)*4))/uint32(p10) + carry
			carry = uint32(int32(1000000000)/p10) * tmp
			if k == a && !(*(*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer((*uint32)(unsafe.Pointer(&x)))) + uintptr(k)*4)) != 0) {
				a = (a + int32(1)) & 127
				rp -= int32(9)
			}
		}
		if carry != 0 {
			*(*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer((*uint32)(unsafe.Pointer(&x)))) + uintptr(func() (_cgo_ret int32) {
				_cgo_addr := &z
				_cgo_ret = *_cgo_addr
				*_cgo_addr++
				return
			}())*4)) = carry
		}
		rp += int32(9) - rpm9
	}
	for rp < 18 || rp == 18 && *(*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer((*uint32)(unsafe.Pointer(&x)))) + uintptr(a)*4)) < *(*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer((*uint32)(unsafe.Pointer(&_cgos_th_floatscan)))) + uintptr(int32(0))*4)) {
		var carry uint32 = uint32(0)
		e2 -= int32(29)
		for k = (z - int32(1)) & 127; ; k = (k - int32(1)) & 127 {
			var tmp uint64 = uint64(*(*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer((*uint32)(unsafe.Pointer(&x)))) + uintptr(k)*4)))<<int32(29) + uint64(carry)
			if tmp > uint64(1000000000) {
				carry = uint32(tmp / uint64(1000000000))
				*(*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer((*uint32)(unsafe.Pointer(&x)))) + uintptr(k)*4)) = uint32(tmp % uint64(1000000000))
			} else {
				carry = uint32(0)
				*(*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer((*uint32)(unsafe.Pointer(&x)))) + uintptr(k)*4)) = uint32(tmp)
			}
			if k == (z-int32(1))&127 && k != a && !(*(*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer((*uint32)(unsafe.Pointer(&x)))) + uintptr(k)*4)) != 0) {
				z = k
			}
			if k == a {
				break
			}
		}
		if carry != 0 {
			rp += int32(9)
			a = (a - int32(1)) & 127
			if a == z {
				z = (z - int32(1)) & 127
				*(*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer((*uint32)(unsafe.Pointer(&x)))) + uintptr((z-int32(1))&127)*4)) |= *(*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer((*uint32)(unsafe.Pointer(&x)))) + uintptr(z)*4))
			}
			*(*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer((*uint32)(unsafe.Pointer(&x)))) + uintptr(a)*4)) = carry
		}
	}
	for {
		var carry uint32 = uint32(0)
		var sh int32 = int32(1)
		for i = int32(0); i < int32(2); i++ {
			k = (a + i) & 127
			if k == z || *(*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer((*uint32)(unsafe.Pointer(&x)))) + uintptr(k)*4)) < *(*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer((*uint32)(unsafe.Pointer(&_cgos_th_floatscan)))) + uintptr(i)*4)) {
				i = int32(2)
				break
			}
			if *(*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer((*uint32)(unsafe.Pointer(&x)))) + uintptr((a+i)&127)*4)) > *(*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer((*uint32)(unsafe.Pointer(&_cgos_th_floatscan)))) + uintptr(i)*4)) {
				break
			}
		}
		if i == int32(2) && rp == 18 {
			break
		}
		if rp > 27 {
			sh = int32(9)
		}
		e2 += sh
		for k = a; k != z; k = (k + int32(1)) & 127 {
			var tmp uint32 = *(*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer((*uint32)(unsafe.Pointer(&x)))) + uintptr(k)*4)) & uint32(int32(1)<<sh-int32(1))
			*(*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer((*uint32)(unsafe.Pointer(&x)))) + uintptr(k)*4)) = *(*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer((*uint32)(unsafe.Pointer(&x)))) + uintptr(k)*4))>>sh + carry
			carry = uint32(int32(1000000000)>>sh) * tmp
			if k == a && !(*(*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer((*uint32)(unsafe.Pointer(&x)))) + uintptr(k)*4)) != 0) {
				a = (a + int32(1)) & 127
				i--
				rp -= int32(9)
			}
		}
		if carry != 0 {
			if (z+int32(1))&127 != a {
				*(*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer((*uint32)(unsafe.Pointer(&x)))) + uintptr(z)*4)) = carry
				z = (z + int32(1)) & 127
			} else {
				*(*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer((*uint32)(unsafe.Pointer(&x)))) + uintptr((z-int32(1))&127)*4)) |= uint32(1)
			}
		}
	}
	for y = float64(func() (_cgo_ret int32) {
		_cgo_addr := &i
		*_cgo_addr = int32(0)
		return *_cgo_addr
	}()); i < int32(2); i++ {
		if (a+i)&127 == z {
			*(*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer((*uint32)(unsafe.Pointer(&x)))) + uintptr(func() (_cgo_ret int32) {
				_cgo_addr := &z
				*_cgo_addr = (z + int32(1)) & 127
				return *_cgo_addr
			}()-int32(1))*4)) = uint32(0)
		}
		y = 1.0e+9*y + float64(*(*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer((*uint32)(unsafe.Pointer(&x)))) + uintptr((a+i)&127)*4)))
	}
	y *= float64(sign)
	if bits > int32(53)+e2-emin {
		bits = int32(53) + e2 - emin
		if bits < int32(0) {
			bits = int32(0)
		}
		denormal = int32(1)
	}
	if bits < int32(53) {
		bias = Copysignl(float64(Scalbn(float64(int32(1)), 106-bits-int32(1))), y)
		frac = Fmodl(y, float64(Scalbn(float64(int32(1)), int32(53)-bits)))
		y -= frac
		y += bias
	}
	if (a+i)&127 != z {
		var t uint32 = *(*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer((*uint32)(unsafe.Pointer(&x)))) + uintptr((a+i)&127)*4))
		if t < uint32(500000000) && (t != 0 || (a+i+int32(1))&127 != z) {
			frac += float64(0.25 * float64(sign))
		} else if t > uint32(500000000) {
			frac += float64(0.75 * float64(sign))
		} else if t == uint32(500000000) {
			if (a+i+int32(1))&127 == z {
				frac += float64(0.5 * float64(sign))
			} else {
				frac += float64(0.75 * float64(sign))
			}
		}
		if int32(53)-bits >= int32(2) && !(Fmodl(frac, float64(int32(1))) != 0) {
			frac++
		}
	}
	y += frac
	y -= bias
	if (e2+int32(53))&int32(2147483647) > emax-int32(5) {
		if Fabsl(y) >= float64(int32(2))/2.22044604925031308085e-16 {
			if denormal != 0 && bits == int32(53)+e2-emin {
				denormal = int32(0)
			}
			y *= float64(0.5)
			e2++
		}
		if e2+int32(53) > emax || denormal != 0 && frac != 0 {
			*__errno_location() = int32(34)
		}
	}
	return Scalbnl(y, e2)
}

var _cgos_th_floatscan [2]uint32 = [2]uint32{uint32(9007199), uint32(254740991)}
var _cgos_p10s_floatscan [8]int32 = [8]int32{int32(10), int32(100), int32(1000), int32(10000), int32(100000), int32(1000000), int32(10000000), int32(100000000)}

func _cgos_hexfloat_floatscan(f *Struct__IO_FILE, bits int32, emin int32, sign int32, pok int32) float64 {
	var x uint32 = uint32(0)
	var y float64 = float64(int32(0))
	var scale float64 = float64(int32(1))
	var bias float64 = float64(int32(0))
	var gottail int32 = int32(0)
	var gotrad int32 = int32(0)
	var gotdig int32 = int32(0)
	var rp int64 = int64(0)
	var dc int64 = int64(0)
	var e2 int64 = int64(0)
	var d int32
	var c int32
	c = func() int32 {
		if uintptr(unsafe.Pointer(f.Rpos)) != uintptr(unsafe.Pointer(f.Shend)) {
			return int32(*func() (_cgo_ret *uint8) {
				_cgo_addr := &f.Rpos
				_cgo_ret = *_cgo_addr
				*(*uintptr)(unsafe.Pointer(_cgo_addr))++
				return
			}())
		} else {
			return __shgetc(f)
		}
	}()
	for ; c == '0'; c = func() int32 {
		if uintptr(unsafe.Pointer(f.Rpos)) != uintptr(unsafe.Pointer(f.Shend)) {
			return int32(*func() (_cgo_ret *uint8) {
				_cgo_addr := &f.Rpos
				_cgo_ret = *_cgo_addr
				*(*uintptr)(unsafe.Pointer(_cgo_addr))++
				return
			}())
		} else {
			return __shgetc(f)
		}
	}() {
		gotdig = int32(1)
	}
	if c == '.' {
		gotrad = int32(1)
		c = func() int32 {
			if uintptr(unsafe.Pointer(f.Rpos)) != uintptr(unsafe.Pointer(f.Shend)) {
				return int32(*func() (_cgo_ret *uint8) {
					_cgo_addr := &f.Rpos
					_cgo_ret = *_cgo_addr
					*(*uintptr)(unsafe.Pointer(_cgo_addr))++
					return
				}())
			} else {
				return __shgetc(f)
			}
		}()
		for rp = int64(0); c == '0'; func() int64 {
			c = func() int32 {
				if uintptr(unsafe.Pointer(f.Rpos)) != uintptr(unsafe.Pointer(f.Shend)) {
					return int32(*func() (_cgo_ret *uint8) {
						_cgo_addr := &f.Rpos
						_cgo_ret = *_cgo_addr
						*(*uintptr)(unsafe.Pointer(_cgo_addr))++
						return
					}())
				} else {
					return __shgetc(f)
				}
			}()
			return func() (_cgo_ret int64) {
				_cgo_addr := &rp
				_cgo_ret = *_cgo_addr
				*_cgo_addr--
				return
			}()
		}() {
			gotdig = int32(1)
		}
	}
	for ; uint32(c-'0') < uint32(10) || uint32(c|int32(32)-'a') < uint32(6) || c == '.'; c = func() int32 {
		if uintptr(unsafe.Pointer(f.Rpos)) != uintptr(unsafe.Pointer(f.Shend)) {
			return int32(*func() (_cgo_ret *uint8) {
				_cgo_addr := &f.Rpos
				_cgo_ret = *_cgo_addr
				*(*uintptr)(unsafe.Pointer(_cgo_addr))++
				return
			}())
		} else {
			return __shgetc(f)
		}
	}() {
		if c == '.' {
			if gotrad != 0 {
				break
			}
			rp = dc
			gotrad = int32(1)
		} else {
			gotdig = int32(1)
			if c > '9' {
				d = c | int32(32) + int32(10) - 'a'
			} else {
				d = c - '0'
			}
			if dc < int64(8) {
				x = x*uint32(16) + uint32(d)
			} else if dc < int64(14) {
				y += float64(d) * func() (_cgo_ret float64) {
					_cgo_addr := &scale
					*_cgo_addr /= float64(int32(16))
					return *_cgo_addr
				}()
			} else if d != 0 && !(gottail != 0) {
				y += float64(0.5) * scale
				gottail = int32(1)
			}
			dc++
		}
	}
	if !(gotdig != 0) {
		if f.Shlim >= int64(0) {
			func() int {
				_ = func() (_cgo_ret *uint8) {
					_cgo_addr := &f.Rpos
					_cgo_ret = *_cgo_addr
					*(*uintptr)(unsafe.Pointer(_cgo_addr))--
					return
				}()
				return 0
			}()
		} else {
			func() int {
				_ = int32(0)
				return 0
			}()
		}
		if pok != 0 {
			if f.Shlim >= int64(0) {
				func() int {
					_ = func() (_cgo_ret *uint8) {
						_cgo_addr := &f.Rpos
						_cgo_ret = *_cgo_addr
						*(*uintptr)(unsafe.Pointer(_cgo_addr))--
						return
					}()
					return 0
				}()
			} else {
				func() int {
					_ = int32(0)
					return 0
				}()
			}
			if gotrad != 0 {
				if f.Shlim >= int64(0) {
					func() int {
						_ = func() (_cgo_ret *uint8) {
							_cgo_addr := &f.Rpos
							_cgo_ret = *_cgo_addr
							*(*uintptr)(unsafe.Pointer(_cgo_addr))--
							return
						}()
						return 0
					}()
				} else {
					func() int {
						_ = int32(0)
						return 0
					}()
				}
			}
		} else {
			__shlim(f, int64(0))
		}
		return float64(float64(sign) * 0)
	}
	if !(gotrad != 0) {
		rp = dc
	}
	for dc < int64(8) {
		func() int64 {
			x *= uint32(16)
			return func() (_cgo_ret int64) {
				_cgo_addr := &dc
				_cgo_ret = *_cgo_addr
				*_cgo_addr++
				return
			}()
		}()
	}
	if c|int32(32) == 'p' {
		e2 = _cgos_scanexp_floatscan(f, pok)
		if e2 == -9223372036854775808 {
			if pok != 0 {
				if f.Shlim >= int64(0) {
					func() int {
						_ = func() (_cgo_ret *uint8) {
							_cgo_addr := &f.Rpos
							_cgo_ret = *_cgo_addr
							*(*uintptr)(unsafe.Pointer(_cgo_addr))--
							return
						}()
						return 0
					}()
				} else {
					func() int {
						_ = int32(0)
						return 0
					}()
				}
			} else {
				__shlim(f, int64(0))
				return float64(int32(0))
			}
			e2 = int64(0)
		}
	} else if f.Shlim >= int64(0) {
		func() int {
			_ = func() (_cgo_ret *uint8) {
				_cgo_addr := &f.Rpos
				_cgo_ret = *_cgo_addr
				*(*uintptr)(unsafe.Pointer(_cgo_addr))--
				return
			}()
			return 0
		}()
	} else {
		func() int {
			_ = int32(0)
			return 0
		}()
	}
	e2 += int64(4)*rp - int64(32)
	if !(x != 0) {
		return float64(float64(sign) * 0)
	}
	if e2 > int64(-emin) {
		*__errno_location() = int32(34)
		return float64(sign) * 1.79769313486231570815e+308 * 1.79769313486231570815e+308
	}
	if e2 < int64(emin-106) {
		*__errno_location() = int32(34)
		return float64(sign) * 2.22507385850720138309e-308 * 2.22507385850720138309e-308
	}
	for x < uint32(2147483648) {
		if y >= float64(0.5) {
			x += x + uint32(1)
			y += y - float64(int32(1))
		} else {
			x += x
			y += y
		}
		e2--
	}
	if int64(bits) > int64(32)+e2-int64(emin) {
		bits = int32(int64(32) + e2 - int64(emin))
		if bits < int32(0) {
			bits = int32(0)
		}
	}
	if bits < int32(53) {
		bias = Copysignl(float64(Scalbn(float64(int32(1)), 85-bits-int32(1))), float64(sign))
	}
	if bits < int32(32) && y != 0 && !(x&uint32(1) != 0) {
		func() float64 {
			x++
			return func() (_cgo_ret float64) {
				_cgo_addr := &y
				*_cgo_addr = float64(int32(0))
				return *_cgo_addr
			}()
		}()
	}
	y = bias + float64(sign)*float64(x) + float64(sign)*y
	y -= bias
	if !(y != 0) {
		*__errno_location() = int32(34)
	}
	return Scalbnl(y, int32(e2))
}
func __floatscan(f *Struct__IO_FILE, prec int32, pok int32) float64 {
	var sign int32 = int32(1)
	var i uint64
	var bits int32
	var emin int32
	var c int32
	switch prec {
	case int32(0):
		bits = int32(24)
		emin = -125 - bits
		break
	case int32(1):
		bits = int32(53)
		emin = -1021 - bits
		break
	case int32(2):
		bits = int32(53)
		emin = -1021 - bits
		break
	default:
		return float64(int32(0))
	}
	for X__isspace(func() (_cgo_ret int32) {
		_cgo_addr := &c
		*_cgo_addr = func() int32 {
			if uintptr(unsafe.Pointer(f.Rpos)) != uintptr(unsafe.Pointer(f.Shend)) {
				return int32(*func() (_cgo_ret *uint8) {
					_cgo_addr := &f.Rpos
					_cgo_ret = *_cgo_addr
					*(*uintptr)(unsafe.Pointer(_cgo_addr))++
					return
				}())
			} else {
				return __shgetc(f)
			}
		}()
		return *_cgo_addr
	}()) != 0 {
	}
	if c == '+' || c == '-' {
		sign -= int32(2) * func() int32 {
			if c == '-' {
				return 1
			} else {
				return 0
			}
		}()
		c = func() int32 {
			if uintptr(unsafe.Pointer(f.Rpos)) != uintptr(unsafe.Pointer(f.Shend)) {
				return int32(*func() (_cgo_ret *uint8) {
					_cgo_addr := &f.Rpos
					_cgo_ret = *_cgo_addr
					*(*uintptr)(unsafe.Pointer(_cgo_addr))++
					return
				}())
			} else {
				return __shgetc(f)
			}
		}()
	}
	for i = uint64(0); i < uint64(8) && c|int32(32) == int32(*(*int8)(unsafe.Pointer(uintptr(unsafe.Pointer((*int8)(unsafe.Pointer(&[9]int8{'i', 'n', 'f', 'i', 'n', 'i', 't', 'y', '\x00'})))) + uintptr(i)))); i++ {
		if i < uint64(7) {
			c = func() int32 {
				if uintptr(unsafe.Pointer(f.Rpos)) != uintptr(unsafe.Pointer(f.Shend)) {
					return int32(*func() (_cgo_ret *uint8) {
						_cgo_addr := &f.Rpos
						_cgo_ret = *_cgo_addr
						*(*uintptr)(unsafe.Pointer(_cgo_addr))++
						return
					}())
				} else {
					return __shgetc(f)
				}
			}()
		}
	}
	if i == uint64(3) || i == uint64(8) || i > uint64(3) && pok != 0 {
		if i != uint64(8) {
			if f.Shlim >= int64(0) {
				func() int {
					_ = func() (_cgo_ret *uint8) {
						_cgo_addr := &f.Rpos
						_cgo_ret = *_cgo_addr
						*(*uintptr)(unsafe.Pointer(_cgo_addr))--
						return
					}()
					return 0
				}()
			} else {
				func() int {
					_ = int32(0)
					return 0
				}()
			}
			if pok != 0 {
				for ; i > uint64(3); i-- {
					if f.Shlim >= int64(0) {
						func() int {
							_ = func() (_cgo_ret *uint8) {
								_cgo_addr := &f.Rpos
								_cgo_ret = *_cgo_addr
								*(*uintptr)(unsafe.Pointer(_cgo_addr))--
								return
							}()
							return 0
						}()
					} else {
						func() int {
							_ = int32(0)
							return 0
						}()
					}
				}
			}
		}
		return float64(float32(sign) * X__builtin_inff())
	}
	if !(i != 0) {
		for i = uint64(0); i < uint64(3) && c|int32(32) == int32(*(*int8)(unsafe.Pointer(uintptr(unsafe.Pointer((*int8)(unsafe.Pointer(&[4]int8{'n', 'a', 'n', '\x00'})))) + uintptr(i)))); i++ {
			if i < uint64(2) {
				c = func() int32 {
					if uintptr(unsafe.Pointer(f.Rpos)) != uintptr(unsafe.Pointer(f.Shend)) {
						return int32(*func() (_cgo_ret *uint8) {
							_cgo_addr := &f.Rpos
							_cgo_ret = *_cgo_addr
							*(*uintptr)(unsafe.Pointer(_cgo_addr))++
							return
						}())
					} else {
						return __shgetc(f)
					}
				}()
			}
		}
	}
	if i == uint64(3) {
		if func() int32 {
			if uintptr(unsafe.Pointer(f.Rpos)) != uintptr(unsafe.Pointer(f.Shend)) {
				return int32(*func() (_cgo_ret *uint8) {
					_cgo_addr := &f.Rpos
					_cgo_ret = *_cgo_addr
					*(*uintptr)(unsafe.Pointer(_cgo_addr))++
					return
				}())
			} else {
				return __shgetc(f)
			}
		}() != '(' {
			if f.Shlim >= int64(0) {
				func() int {
					_ = func() (_cgo_ret *uint8) {
						_cgo_addr := &f.Rpos
						_cgo_ret = *_cgo_addr
						*(*uintptr)(unsafe.Pointer(_cgo_addr))--
						return
					}()
					return 0
				}()
			} else {
				func() int {
					_ = int32(0)
					return 0
				}()
			}
			return float64(X__builtin_nanf((*int8)(unsafe.Pointer(&[1]int8{'\x00'}))))
		}
		for i = uint64(1); ; i++ {
			c = func() int32 {
				if uintptr(unsafe.Pointer(f.Rpos)) != uintptr(unsafe.Pointer(f.Shend)) {
					return int32(*func() (_cgo_ret *uint8) {
						_cgo_addr := &f.Rpos
						_cgo_ret = *_cgo_addr
						*(*uintptr)(unsafe.Pointer(_cgo_addr))++
						return
					}())
				} else {
					return __shgetc(f)
				}
			}()
			if uint32(c-'0') < uint32(10) || uint32(c-'A') < uint32(26) || uint32(c-'a') < uint32(26) || c == '_' {
				continue
			}
			if c == ')' {
				return float64(X__builtin_nanf((*int8)(unsafe.Pointer(&[1]int8{'\x00'}))))
			}
			if f.Shlim >= int64(0) {
				func() int {
					_ = func() (_cgo_ret *uint8) {
						_cgo_addr := &f.Rpos
						_cgo_ret = *_cgo_addr
						*(*uintptr)(unsafe.Pointer(_cgo_addr))--
						return
					}()
					return 0
				}()
			} else {
				func() int {
					_ = int32(0)
					return 0
				}()
			}
			if !(pok != 0) {
				*__errno_location() = int32(22)
				__shlim(f, int64(0))
				return float64(int32(0))
			}
			for func() (_cgo_ret uint64) {
				_cgo_addr := &i
				_cgo_ret = *_cgo_addr
				*_cgo_addr--
				return
			}() != 0 {
				if f.Shlim >= int64(0) {
					func() int {
						_ = func() (_cgo_ret *uint8) {
							_cgo_addr := &f.Rpos
							_cgo_ret = *_cgo_addr
							*(*uintptr)(unsafe.Pointer(_cgo_addr))--
							return
						}()
						return 0
					}()
				} else {
					func() int {
						_ = int32(0)
						return 0
					}()
				}
			}
			return float64(X__builtin_nanf((*int8)(unsafe.Pointer(&[1]int8{'\x00'}))))
		}
		return float64(X__builtin_nanf((*int8)(unsafe.Pointer(&[1]int8{'\x00'}))))
	}
	if i != 0 {
		if f.Shlim >= int64(0) {
			func() int {
				_ = func() (_cgo_ret *uint8) {
					_cgo_addr := &f.Rpos
					_cgo_ret = *_cgo_addr
					*(*uintptr)(unsafe.Pointer(_cgo_addr))--
					return
				}()
				return 0
			}()
		} else {
			func() int {
				_ = int32(0)
				return 0
			}()
		}
		*__errno_location() = int32(22)
		__shlim(f, int64(0))
		return float64(int32(0))
	}
	if c == '0' {
		c = func() int32 {
			if uintptr(unsafe.Pointer(f.Rpos)) != uintptr(unsafe.Pointer(f.Shend)) {
				return int32(*func() (_cgo_ret *uint8) {
					_cgo_addr := &f.Rpos
					_cgo_ret = *_cgo_addr
					*(*uintptr)(unsafe.Pointer(_cgo_addr))++
					return
				}())
			} else {
				return __shgetc(f)
			}
		}()
		if c|int32(32) == 'x' {
			return _cgos_hexfloat_floatscan(f, bits, emin, sign, pok)
		}
		if f.Shlim >= int64(0) {
			func() int {
				_ = func() (_cgo_ret *uint8) {
					_cgo_addr := &f.Rpos
					_cgo_ret = *_cgo_addr
					*(*uintptr)(unsafe.Pointer(_cgo_addr))--
					return
				}()
				return 0
			}()
		} else {
			func() int {
				_ = int32(0)
				return 0
			}()
		}
		c = int32('0')
	}
	return _cgos_decfloat_floatscan(f, c, bits, emin, sign, pok)
}
