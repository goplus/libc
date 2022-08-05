package libc

import unsafe "unsafe"

type _cgos_sha256_crypt_sha256 struct {
	len uint64
	h   [8]uint32
	buf [64]uint8
}

func _cgos_ror_crypt_sha256(n uint32, k int32) uint32 {
	return n>>k | n<<(int32(32)-k)
}

var _cgos_K_crypt_sha256 [64]uint32 = [64]uint32{uint32(1116352408), uint32(1899447441), uint32(3049323471), uint32(3921009573), uint32(961987163), uint32(1508970993), uint32(2453635748), uint32(2870763221), uint32(3624381080), uint32(310598401), uint32(607225278), uint32(1426881987), uint32(1925078388), uint32(2162078206), uint32(2614888103), uint32(3248222580), uint32(3835390401), uint32(4022224774), uint32(264347078), uint32(604807628), uint32(770255983), uint32(1249150122), uint32(1555081692), uint32(1996064986), uint32(2554220882), uint32(2821834349), uint32(2952996808), uint32(3210313671), uint32(3336571891), uint32(3584528711), uint32(113926993), uint32(338241895), uint32(666307205), uint32(773529912), uint32(1294757372), uint32(1396182291), uint32(1695183700), uint32(1986661051), uint32(2177026350), uint32(2456956037), uint32(2730485921), uint32(2820302411), uint32(3259730800), uint32(3345764771), uint32(3516065817), uint32(3600352804), uint32(4094571909), uint32(275423344), uint32(430227734), uint32(506948616), uint32(659060556), uint32(883997877), uint32(958139571), uint32(1322822218), uint32(1537002063), uint32(1747873779), uint32(1955562222), uint32(2024104815), uint32(2227730452), uint32(2361852424), uint32(2428436474), uint32(2756734187), uint32(3204031479), uint32(3329325298)}

func _cgos_processblock_crypt_sha256(s *_cgos_sha256_crypt_sha256, buf *uint8) {
	var W [64]uint32
	var t1 uint32
	var t2 uint32
	var a uint32
	var b uint32
	var c uint32
	var d uint32
	var e uint32
	var f uint32
	var g uint32
	var h uint32
	var i int32
	for i = int32(0); i < int32(16); i++ {
		*(*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer((*uint32)(unsafe.Pointer(&W)))) + uintptr(i)*4)) = uint32(*(*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(buf)) + uintptr(int32(4)*i)))) << int32(24)
		*(*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer((*uint32)(unsafe.Pointer(&W)))) + uintptr(i)*4)) |= uint32(*(*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(buf)) + uintptr(int32(4)*i+int32(1))))) << int32(16)
		*(*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer((*uint32)(unsafe.Pointer(&W)))) + uintptr(i)*4)) |= uint32(*(*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(buf)) + uintptr(int32(4)*i+int32(2))))) << int32(8)
		*(*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer((*uint32)(unsafe.Pointer(&W)))) + uintptr(i)*4)) |= uint32(*(*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(buf)) + uintptr(int32(4)*i+int32(3)))))
	}
	for ; i < int32(64); i++ {
		*(*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer((*uint32)(unsafe.Pointer(&W)))) + uintptr(i)*4)) = _cgos_ror_crypt_sha256(*(*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer((*uint32)(unsafe.Pointer(&W)))) + uintptr(i-int32(2))*4)), int32(17)) ^ _cgos_ror_crypt_sha256(*(*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer((*uint32)(unsafe.Pointer(&W)))) + uintptr(i-int32(2))*4)), int32(19)) ^ *(*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer((*uint32)(unsafe.Pointer(&W)))) + uintptr(i-int32(2))*4))>>int32(10) + *(*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer((*uint32)(unsafe.Pointer(&W)))) + uintptr(i-int32(7))*4)) + (_cgos_ror_crypt_sha256(*(*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer((*uint32)(unsafe.Pointer(&W)))) + uintptr(i-int32(15))*4)), int32(7)) ^ _cgos_ror_crypt_sha256(*(*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer((*uint32)(unsafe.Pointer(&W)))) + uintptr(i-int32(15))*4)), int32(18)) ^ *(*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer((*uint32)(unsafe.Pointer(&W)))) + uintptr(i-int32(15))*4))>>int32(3)) + *(*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer((*uint32)(unsafe.Pointer(&W)))) + uintptr(i-int32(16))*4))
	}
	a = *(*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer((*uint32)(unsafe.Pointer(&s.h)))) + uintptr(int32(0))*4))
	b = *(*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer((*uint32)(unsafe.Pointer(&s.h)))) + uintptr(int32(1))*4))
	c = *(*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer((*uint32)(unsafe.Pointer(&s.h)))) + uintptr(int32(2))*4))
	d = *(*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer((*uint32)(unsafe.Pointer(&s.h)))) + uintptr(int32(3))*4))
	e = *(*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer((*uint32)(unsafe.Pointer(&s.h)))) + uintptr(int32(4))*4))
	f = *(*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer((*uint32)(unsafe.Pointer(&s.h)))) + uintptr(int32(5))*4))
	g = *(*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer((*uint32)(unsafe.Pointer(&s.h)))) + uintptr(int32(6))*4))
	h = *(*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer((*uint32)(unsafe.Pointer(&s.h)))) + uintptr(int32(7))*4))
	for i = int32(0); i < int32(64); i++ {
		t1 = h + (_cgos_ror_crypt_sha256(e, int32(6)) ^ _cgos_ror_crypt_sha256(e, int32(11)) ^ _cgos_ror_crypt_sha256(e, int32(25))) + (g ^ e&(f^g)) + *(*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer((*uint32)(unsafe.Pointer(&_cgos_K_crypt_sha256)))) + uintptr(i)*4)) + *(*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer((*uint32)(unsafe.Pointer(&W)))) + uintptr(i)*4))
		t2 = _cgos_ror_crypt_sha256(a, int32(2)) ^ _cgos_ror_crypt_sha256(a, int32(13)) ^ _cgos_ror_crypt_sha256(a, int32(22)) + (a&b | c&(a|b))
		h = g
		g = f
		f = e
		e = d + t1
		d = c
		c = b
		b = a
		a = t1 + t2
	}
	*(*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer((*uint32)(unsafe.Pointer(&s.h)))) + uintptr(int32(0))*4)) += a
	*(*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer((*uint32)(unsafe.Pointer(&s.h)))) + uintptr(int32(1))*4)) += b
	*(*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer((*uint32)(unsafe.Pointer(&s.h)))) + uintptr(int32(2))*4)) += c
	*(*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer((*uint32)(unsafe.Pointer(&s.h)))) + uintptr(int32(3))*4)) += d
	*(*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer((*uint32)(unsafe.Pointer(&s.h)))) + uintptr(int32(4))*4)) += e
	*(*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer((*uint32)(unsafe.Pointer(&s.h)))) + uintptr(int32(5))*4)) += f
	*(*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer((*uint32)(unsafe.Pointer(&s.h)))) + uintptr(int32(6))*4)) += g
	*(*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer((*uint32)(unsafe.Pointer(&s.h)))) + uintptr(int32(7))*4)) += h
}
func _cgos_pad_crypt_sha256(s *_cgos_sha256_crypt_sha256) {
	var r uint32 = uint32(s.len % uint64(64))
	*(*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer((*uint8)(unsafe.Pointer(&s.buf)))) + uintptr(func() (_cgo_ret uint32) {
		_cgo_addr := &r
		_cgo_ret = *_cgo_addr
		*_cgo_addr++
		return
	}()))) = uint8(128)
	if r > uint32(56) {
		Memset(unsafe.Pointer((*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer((*uint8)(unsafe.Pointer(&s.buf))))+uintptr(r)))), int32(0), uint64(uint32(64)-r))
		r = uint32(0)
		_cgos_processblock_crypt_sha256(s, (*uint8)(unsafe.Pointer(&s.buf)))
	}
	Memset(unsafe.Pointer((*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer((*uint8)(unsafe.Pointer(&s.buf))))+uintptr(r)))), int32(0), uint64(uint32(56)-r))
	s.len *= uint64(8)
	*(*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer((*uint8)(unsafe.Pointer(&s.buf)))) + uintptr(int32(56)))) = uint8(s.len >> int32(56))
	*(*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer((*uint8)(unsafe.Pointer(&s.buf)))) + uintptr(int32(57)))) = uint8(s.len >> int32(48))
	*(*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer((*uint8)(unsafe.Pointer(&s.buf)))) + uintptr(int32(58)))) = uint8(s.len >> int32(40))
	*(*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer((*uint8)(unsafe.Pointer(&s.buf)))) + uintptr(int32(59)))) = uint8(s.len >> int32(32))
	*(*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer((*uint8)(unsafe.Pointer(&s.buf)))) + uintptr(int32(60)))) = uint8(s.len >> int32(24))
	*(*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer((*uint8)(unsafe.Pointer(&s.buf)))) + uintptr(int32(61)))) = uint8(s.len >> int32(16))
	*(*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer((*uint8)(unsafe.Pointer(&s.buf)))) + uintptr(int32(62)))) = uint8(s.len >> int32(8))
	*(*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer((*uint8)(unsafe.Pointer(&s.buf)))) + uintptr(int32(63)))) = uint8(s.len)
	_cgos_processblock_crypt_sha256(s, (*uint8)(unsafe.Pointer(&s.buf)))
}
func _cgos_sha256_init_crypt_sha256(s *_cgos_sha256_crypt_sha256) {
	s.len = uint64(0)
	*(*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer((*uint32)(unsafe.Pointer(&s.h)))) + uintptr(int32(0))*4)) = uint32(1779033703)
	*(*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer((*uint32)(unsafe.Pointer(&s.h)))) + uintptr(int32(1))*4)) = uint32(3144134277)
	*(*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer((*uint32)(unsafe.Pointer(&s.h)))) + uintptr(int32(2))*4)) = uint32(1013904242)
	*(*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer((*uint32)(unsafe.Pointer(&s.h)))) + uintptr(int32(3))*4)) = uint32(2773480762)
	*(*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer((*uint32)(unsafe.Pointer(&s.h)))) + uintptr(int32(4))*4)) = uint32(1359893119)
	*(*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer((*uint32)(unsafe.Pointer(&s.h)))) + uintptr(int32(5))*4)) = uint32(2600822924)
	*(*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer((*uint32)(unsafe.Pointer(&s.h)))) + uintptr(int32(6))*4)) = uint32(528734635)
	*(*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer((*uint32)(unsafe.Pointer(&s.h)))) + uintptr(int32(7))*4)) = uint32(1541459225)
}
func _cgos_sha256_sum_crypt_sha256(s *_cgos_sha256_crypt_sha256, md *uint8) {
	var i int32
	_cgos_pad_crypt_sha256(s)
	for i = int32(0); i < int32(8); i++ {
		*(*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(md)) + uintptr(int32(4)*i))) = uint8(*(*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer((*uint32)(unsafe.Pointer(&s.h)))) + uintptr(i)*4)) >> int32(24))
		*(*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(md)) + uintptr(int32(4)*i+int32(1)))) = uint8(*(*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer((*uint32)(unsafe.Pointer(&s.h)))) + uintptr(i)*4)) >> int32(16))
		*(*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(md)) + uintptr(int32(4)*i+int32(2)))) = uint8(*(*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer((*uint32)(unsafe.Pointer(&s.h)))) + uintptr(i)*4)) >> int32(8))
		*(*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(md)) + uintptr(int32(4)*i+int32(3)))) = uint8(*(*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer((*uint32)(unsafe.Pointer(&s.h)))) + uintptr(i)*4)))
	}
}
func _cgos_sha256_update_crypt_sha256(s *_cgos_sha256_crypt_sha256, m unsafe.Pointer, len uint64) {
	var p *uint8 = (*uint8)(m)
	var r uint32 = uint32(s.len % uint64(64))
	s.len += uint64(len)
	if r != 0 {
		if len < uint64(uint32(64)-r) {
			Memcpy(unsafe.Pointer((*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer((*uint8)(unsafe.Pointer(&s.buf))))+uintptr(r)))), unsafe.Pointer(p), len)
			return
		}
		Memcpy(unsafe.Pointer((*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer((*uint8)(unsafe.Pointer(&s.buf))))+uintptr(r)))), unsafe.Pointer(p), uint64(uint32(64)-r))
		len -= uint64(uint32(64) - r)
		*(*uintptr)(unsafe.Pointer(&p)) += uintptr(uint32(64) - r)
		_cgos_processblock_crypt_sha256(s, (*uint8)(unsafe.Pointer(&s.buf)))
	}
	for ; len >= uint64(64); func() *uint8 {
		len -= uint64(64)
		return func() (_cgo_ret *uint8) {
			_cgo_addr := &p
			*(*uintptr)(unsafe.Pointer(&*_cgo_addr)) += uintptr(int32(64))
			return *_cgo_addr
		}()
	}() {
		_cgos_processblock_crypt_sha256(s, p)
	}
	Memcpy(unsafe.Pointer((*uint8)(unsafe.Pointer(&s.buf))), unsafe.Pointer(p), len)
}

var _cgos_b64_crypt_sha256 [65]uint8 = [65]uint8{'.', '/', '0', '1', '2', '3', '4', '5', '6', '7', '8', '9', 'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J', 'K', 'L', 'M', 'N', 'O', 'P', 'Q', 'R', 'S', 'T', 'U', 'V', 'W', 'X', 'Y', 'Z', 'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z', '\x00'}

func _cgos_to64_crypt_sha256(s *int8, u uint32, n int32) *int8 {
	for func() (_cgo_ret int32) {
		_cgo_addr := &n
		*_cgo_addr--
		return *_cgo_addr
	}() >= int32(0) {
		*func() (_cgo_ret *int8) {
			_cgo_addr := &s
			_cgo_ret = *_cgo_addr
			*(*uintptr)(unsafe.Pointer(_cgo_addr))++
			return
		}() = int8(*(*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer((*uint8)(unsafe.Pointer(&_cgos_b64_crypt_sha256)))) + uintptr(u%uint32(64)))))
		u /= uint32(64)
	}
	return s
}
func _cgos_hashmd_crypt_sha256(s *_cgos_sha256_crypt_sha256, n uint32, md unsafe.Pointer) {
	var i uint32
	for i = n; i > uint32(32); i -= uint32(32) {
		_cgos_sha256_update_crypt_sha256(s, md, uint64(32))
	}
	_cgos_sha256_update_crypt_sha256(s, md, uint64(i))
}
func _cgos_sha256crypt_crypt_sha256(key *int8, setting *int8, output *int8) *int8 {
	var ctx _cgos_sha256_crypt_sha256
	var md [32]uint8
	var kmd [32]uint8
	var smd [32]uint8
	var i uint32
	var r uint32
	var klen uint32
	var slen uint32
	var rounds [20]int8 = [20]int8{'\x00'}
	var salt *int8
	var p *int8
	klen = uint32(Strnlen(key, uint64(257)))
	if klen > uint32(256) {
		return (*int8)(nil)
	}
	if Strncmp(setting, (*int8)(unsafe.Pointer(&[4]int8{'$', '5', '$', '\x00'})), uint64(3)) != int32(0) {
		return (*int8)(nil)
	}
	salt = (*int8)(unsafe.Pointer(uintptr(unsafe.Pointer(setting)) + uintptr(int32(3))))
	r = uint32(5000)
	if Strncmp(salt, (*int8)(unsafe.Pointer(&[8]int8{'r', 'o', 'u', 'n', 'd', 's', '=', '\x00'})), 7) == int32(0) {
		var u uint64
		var end *int8
		*(*uintptr)(unsafe.Pointer(&salt)) += uintptr(7)
		if !(func() int32 {
			if int32(0) != 0 {
				return Isdigit(int32(*salt))
			} else {
				return func() int32 {
					if uint32(*salt)-uint32('0') < uint32(10) {
						return 1
					} else {
						return 0
					}
				}()
			}
		}() != 0) {
			return (*int8)(nil)
		}
		u = Strtoul(salt, &end, int32(10))
		if int32(*end) != '$' {
			return (*int8)(nil)
		}
		salt = (*int8)(unsafe.Pointer(uintptr(unsafe.Pointer(end)) + uintptr(int32(1))))
		if u < uint64(1000) {
			r = uint32(1000)
		} else if u > uint64(9999999) {
			return (*int8)(nil)
		} else {
			r = uint32(u)
		}
		Sprintf((*int8)(unsafe.Pointer(&rounds)), (*int8)(unsafe.Pointer(&[11]int8{'r', 'o', 'u', 'n', 'd', 's', '=', '%', 'u', '$', '\x00'})), r)
	}
	for i = uint32(0); i < uint32(16) && int32(*(*int8)(unsafe.Pointer(uintptr(unsafe.Pointer(salt)) + uintptr(i)))) != 0 && int32(*(*int8)(unsafe.Pointer(uintptr(unsafe.Pointer(salt)) + uintptr(i)))) != '$'; i++ {
		if int32(*(*int8)(unsafe.Pointer(uintptr(unsafe.Pointer(salt)) + uintptr(i)))) == '\n' || int32(*(*int8)(unsafe.Pointer(uintptr(unsafe.Pointer(salt)) + uintptr(i)))) == ':' {
			return (*int8)(nil)
		}
	}
	slen = i
	_cgos_sha256_init_crypt_sha256(&ctx)
	_cgos_sha256_update_crypt_sha256(&ctx, unsafe.Pointer(key), uint64(klen))
	_cgos_sha256_update_crypt_sha256(&ctx, unsafe.Pointer(salt), uint64(slen))
	_cgos_sha256_update_crypt_sha256(&ctx, unsafe.Pointer(key), uint64(klen))
	_cgos_sha256_sum_crypt_sha256(&ctx, (*uint8)(unsafe.Pointer(&md)))
	_cgos_sha256_init_crypt_sha256(&ctx)
	_cgos_sha256_update_crypt_sha256(&ctx, unsafe.Pointer(key), uint64(klen))
	_cgos_sha256_update_crypt_sha256(&ctx, unsafe.Pointer(salt), uint64(slen))
	_cgos_hashmd_crypt_sha256(&ctx, klen, unsafe.Pointer((*uint8)(unsafe.Pointer(&md))))
	for i = klen; i > uint32(0); i >>= int32(1) {
		if i&uint32(1) != 0 {
			_cgos_sha256_update_crypt_sha256(&ctx, unsafe.Pointer((*uint8)(unsafe.Pointer(&md))), 32)
		} else {
			_cgos_sha256_update_crypt_sha256(&ctx, unsafe.Pointer(key), uint64(klen))
		}
	}
	_cgos_sha256_sum_crypt_sha256(&ctx, (*uint8)(unsafe.Pointer(&md)))
	_cgos_sha256_init_crypt_sha256(&ctx)
	for i = uint32(0); i < klen; i++ {
		_cgos_sha256_update_crypt_sha256(&ctx, unsafe.Pointer(key), uint64(klen))
	}
	_cgos_sha256_sum_crypt_sha256(&ctx, (*uint8)(unsafe.Pointer(&kmd)))
	_cgos_sha256_init_crypt_sha256(&ctx)
	for i = uint32(0); i < uint32(int32(16)+int32(*(*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer((*uint8)(unsafe.Pointer(&md)))) + uintptr(int32(0)))))); i++ {
		_cgos_sha256_update_crypt_sha256(&ctx, unsafe.Pointer(salt), uint64(slen))
	}
	_cgos_sha256_sum_crypt_sha256(&ctx, (*uint8)(unsafe.Pointer(&smd)))
	for i = uint32(0); i < r; i++ {
		_cgos_sha256_init_crypt_sha256(&ctx)
		if i%uint32(2) != 0 {
			_cgos_hashmd_crypt_sha256(&ctx, klen, unsafe.Pointer((*uint8)(unsafe.Pointer(&kmd))))
		} else {
			_cgos_sha256_update_crypt_sha256(&ctx, unsafe.Pointer((*uint8)(unsafe.Pointer(&md))), 32)
		}
		if i%uint32(3) != 0 {
			_cgos_sha256_update_crypt_sha256(&ctx, unsafe.Pointer((*uint8)(unsafe.Pointer(&smd))), uint64(slen))
		}
		if i%uint32(7) != 0 {
			_cgos_hashmd_crypt_sha256(&ctx, klen, unsafe.Pointer((*uint8)(unsafe.Pointer(&kmd))))
		}
		if i%uint32(2) != 0 {
			_cgos_sha256_update_crypt_sha256(&ctx, unsafe.Pointer((*uint8)(unsafe.Pointer(&md))), 32)
		} else {
			_cgos_hashmd_crypt_sha256(&ctx, klen, unsafe.Pointer((*uint8)(unsafe.Pointer(&kmd))))
		}
		_cgos_sha256_sum_crypt_sha256(&ctx, (*uint8)(unsafe.Pointer(&md)))
	}
	p = output
	*(*uintptr)(unsafe.Pointer(&p)) += uintptr(Sprintf(p, (*int8)(unsafe.Pointer(&[11]int8{'$', '5', '$', '%', 's', '%', '.', '*', 's', '$', '\x00'})), (*int8)(unsafe.Pointer(&rounds)), slen, salt))
	for i = uint32(0); i < uint32(10); i++ {
		p = _cgos_to64_crypt_sha256(p, uint32(int32(*(*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer((*uint8)(unsafe.Pointer(&md)))) + uintptr(*(*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer((*uint8)(unsafe.Pointer(&*(*[3]uint8)(unsafe.Pointer(uintptr(unsafe.Pointer((*[3]uint8)(unsafe.Pointer(&_cgos_perm_crypt_sha256)))) + uintptr(i)*3)))))) + uintptr(int32(0))))))))<<int32(16)|int32(*(*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer((*uint8)(unsafe.Pointer(&md)))) + uintptr(*(*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer((*uint8)(unsafe.Pointer(&*(*[3]uint8)(unsafe.Pointer(uintptr(unsafe.Pointer((*[3]uint8)(unsafe.Pointer(&_cgos_perm_crypt_sha256)))) + uintptr(i)*3)))))) + uintptr(int32(1))))))))<<int32(8)|int32(*(*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer((*uint8)(unsafe.Pointer(&md)))) + uintptr(*(*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer((*uint8)(unsafe.Pointer(&*(*[3]uint8)(unsafe.Pointer(uintptr(unsafe.Pointer((*[3]uint8)(unsafe.Pointer(&_cgos_perm_crypt_sha256)))) + uintptr(i)*3)))))) + uintptr(int32(2))))))))), int32(4))
	}
	p = _cgos_to64_crypt_sha256(p, uint32(int32(*(*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer((*uint8)(unsafe.Pointer(&md)))) + uintptr(int32(31)))))<<int32(8)|int32(*(*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer((*uint8)(unsafe.Pointer(&md)))) + uintptr(int32(30)))))), int32(3))
	*p = int8(0)
	return output
}

var _cgos_perm_crypt_sha256 [10][3]uint8 = [10][3]uint8{[3]uint8{uint8(0), uint8(10), uint8(20)}, [3]uint8{uint8(21), uint8(1), uint8(11)}, [3]uint8{uint8(12), uint8(22), uint8(2)}, [3]uint8{uint8(3), uint8(13), uint8(23)}, [3]uint8{uint8(24), uint8(4), uint8(14)}, [3]uint8{uint8(15), uint8(25), uint8(5)}, [3]uint8{uint8(6), uint8(16), uint8(26)}, [3]uint8{uint8(27), uint8(7), uint8(17)}, [3]uint8{uint8(18), uint8(28), uint8(8)}, [3]uint8{uint8(9), uint8(19), uint8(29)}}

func __crypt_sha256(key *int8, setting *int8, output *int8) *int8 {
	var testbuf [128]int8
	var p *int8
	var q *int8
	p = _cgos_sha256crypt_crypt_sha256(key, setting, output)
	q = _cgos_sha256crypt_crypt_sha256((*int8)(unsafe.Pointer(&_cgos_testkey_crypt_sha256)), (*int8)(unsafe.Pointer(&_cgos_testsetting_crypt_sha256)), (*int8)(unsafe.Pointer(&testbuf)))
	if !(p != nil) || uintptr(unsafe.Pointer(q)) != uintptr(unsafe.Pointer((*int8)(unsafe.Pointer(&testbuf)))) || Memcmp(unsafe.Pointer((*int8)(unsafe.Pointer(&testbuf))), unsafe.Pointer((*int8)(unsafe.Pointer(&_cgos_testhash_crypt_sha256))), 73) != 0 {
		return (*int8)(unsafe.Pointer(&[2]int8{'*', '\x00'}))
	}
	return p
}

var _cgos_testkey_crypt_sha256 [18]int8 = [18]int8{'X', 'y', '0', '1', '@', '#', '\x01', '\x02', -128, '\u007f', -1, '\r', '\n', -127, '\t', ' ', '!', '\x00'}
var _cgos_testsetting_crypt_sha256 [30]int8 = [30]int8{'$', '5', '$', 'r', 'o', 'u', 'n', 'd', 's', '=', '1', '2', '3', '4', '$', 'a', 'b', 'c', '0', '1', '2', '3', '4', '5', '6', '7', '8', '9', '$', '\x00'}
var _cgos_testhash_crypt_sha256 [73]int8 = [73]int8{'$', '5', '$', 'r', 'o', 'u', 'n', 'd', 's', '=', '1', '2', '3', '4', '$', 'a', 'b', 'c', '0', '1', '2', '3', '4', '5', '6', '7', '8', '9', '$', '3', 'V', 'f', 'D', 'j', 'P', 't', '0', '5', 'V', 'H', 'F', 'n', '4', '7', 'C', '/', 'o', 'j', 'F', 'Z', '6', 'K', 'R', 'P', 'Y', 'r', 'O', 'j', 'j', '1', 'l', 'L', 'b', 'H', '.', 'd', 'k', 'F', '3', 'b', 'Z', '6', '\x00'}
