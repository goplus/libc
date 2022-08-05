package libc

import unsafe "unsafe"

type _cgos_md5_crypt_md5 struct {
	len uint64
	h   [4]uint32
	buf [64]uint8
}

func _cgos_rol_crypt_md5(n uint32, k int32) uint32 {
	return n<<k | n>>(int32(32)-k)
}

var _cgos_tab_crypt_md5 [64]uint32 = [64]uint32{uint32(3614090360), uint32(3905402710), uint32(606105819), uint32(3250441966), uint32(4118548399), uint32(1200080426), uint32(2821735955), uint32(4249261313), uint32(1770035416), uint32(2336552879), uint32(4294925233), uint32(2304563134), uint32(1804603682), uint32(4254626195), uint32(2792965006), uint32(1236535329), uint32(4129170786), uint32(3225465664), uint32(643717713), uint32(3921069994), uint32(3593408605), uint32(38016083), uint32(3634488961), uint32(3889429448), uint32(568446438), uint32(3275163606), uint32(4107603335), uint32(1163531501), uint32(2850285829), uint32(4243563512), uint32(1735328473), uint32(2368359562), uint32(4294588738), uint32(2272392833), uint32(1839030562), uint32(4259657740), uint32(2763975236), uint32(1272893353), uint32(4139469664), uint32(3200236656), uint32(681279174), uint32(3936430074), uint32(3572445317), uint32(76029189), uint32(3654602809), uint32(3873151461), uint32(530742520), uint32(3299628645), uint32(4096336452), uint32(1126891415), uint32(2878612391), uint32(4237533241), uint32(1700485571), uint32(2399980690), uint32(4293915773), uint32(2240044497), uint32(1873313359), uint32(4264355552), uint32(2734768916), uint32(1309151649), uint32(4149444226), uint32(3174756917), uint32(718787259), uint32(3951481745)}

func _cgos_processblock_crypt_md5(s *_cgos_md5_crypt_md5, buf *uint8) {
	var i uint32
	var W [16]uint32
	var a uint32
	var b uint32
	var c uint32
	var d uint32
	for i = uint32(0); i < uint32(16); i++ {
		*(*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer((*uint32)(unsafe.Pointer(&W)))) + uintptr(i)*4)) = uint32(*(*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(buf)) + uintptr(uint32(4)*i))))
		*(*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer((*uint32)(unsafe.Pointer(&W)))) + uintptr(i)*4)) |= uint32(*(*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(buf)) + uintptr(uint32(4)*i+uint32(1))))) << int32(8)
		*(*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer((*uint32)(unsafe.Pointer(&W)))) + uintptr(i)*4)) |= uint32(*(*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(buf)) + uintptr(uint32(4)*i+uint32(2))))) << int32(16)
		*(*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer((*uint32)(unsafe.Pointer(&W)))) + uintptr(i)*4)) |= uint32(*(*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(buf)) + uintptr(uint32(4)*i+uint32(3))))) << int32(24)
	}
	a = *(*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer((*uint32)(unsafe.Pointer(&s.h)))) + uintptr(int32(0))*4))
	b = *(*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer((*uint32)(unsafe.Pointer(&s.h)))) + uintptr(int32(1))*4))
	c = *(*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer((*uint32)(unsafe.Pointer(&s.h)))) + uintptr(int32(2))*4))
	d = *(*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer((*uint32)(unsafe.Pointer(&s.h)))) + uintptr(int32(3))*4))
	i = uint32(0)
	for i < uint32(16) {
		a += d ^ b&(c^d) + *(*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer((*uint32)(unsafe.Pointer(&W)))) + uintptr(i)*4)) + *(*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer((*uint32)(unsafe.Pointer(&_cgos_tab_crypt_md5)))) + uintptr(i)*4))
		a = _cgos_rol_crypt_md5(a, int32(7)) + b
		i++
		d += c ^ a&(b^c) + *(*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer((*uint32)(unsafe.Pointer(&W)))) + uintptr(i)*4)) + *(*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer((*uint32)(unsafe.Pointer(&_cgos_tab_crypt_md5)))) + uintptr(i)*4))
		d = _cgos_rol_crypt_md5(d, int32(12)) + a
		i++
		c += b ^ d&(a^b) + *(*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer((*uint32)(unsafe.Pointer(&W)))) + uintptr(i)*4)) + *(*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer((*uint32)(unsafe.Pointer(&_cgos_tab_crypt_md5)))) + uintptr(i)*4))
		c = _cgos_rol_crypt_md5(c, int32(17)) + d
		i++
		b += a ^ c&(d^a) + *(*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer((*uint32)(unsafe.Pointer(&W)))) + uintptr(i)*4)) + *(*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer((*uint32)(unsafe.Pointer(&_cgos_tab_crypt_md5)))) + uintptr(i)*4))
		b = _cgos_rol_crypt_md5(b, int32(22)) + c
		i++
	}
	for i < uint32(32) {
		a += c ^ d&(c^b) + *(*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer((*uint32)(unsafe.Pointer(&W)))) + uintptr((uint32(5)*i+uint32(1))%uint32(16))*4)) + *(*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer((*uint32)(unsafe.Pointer(&_cgos_tab_crypt_md5)))) + uintptr(i)*4))
		a = _cgos_rol_crypt_md5(a, int32(5)) + b
		i++
		d += b ^ c&(b^a) + *(*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer((*uint32)(unsafe.Pointer(&W)))) + uintptr((uint32(5)*i+uint32(1))%uint32(16))*4)) + *(*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer((*uint32)(unsafe.Pointer(&_cgos_tab_crypt_md5)))) + uintptr(i)*4))
		d = _cgos_rol_crypt_md5(d, int32(9)) + a
		i++
		c += a ^ b&(a^d) + *(*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer((*uint32)(unsafe.Pointer(&W)))) + uintptr((uint32(5)*i+uint32(1))%uint32(16))*4)) + *(*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer((*uint32)(unsafe.Pointer(&_cgos_tab_crypt_md5)))) + uintptr(i)*4))
		c = _cgos_rol_crypt_md5(c, int32(14)) + d
		i++
		b += d ^ a&(d^c) + *(*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer((*uint32)(unsafe.Pointer(&W)))) + uintptr((uint32(5)*i+uint32(1))%uint32(16))*4)) + *(*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer((*uint32)(unsafe.Pointer(&_cgos_tab_crypt_md5)))) + uintptr(i)*4))
		b = _cgos_rol_crypt_md5(b, int32(20)) + c
		i++
	}
	for i < uint32(48) {
		a += b ^ c ^ d + *(*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer((*uint32)(unsafe.Pointer(&W)))) + uintptr((uint32(3)*i+uint32(5))%uint32(16))*4)) + *(*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer((*uint32)(unsafe.Pointer(&_cgos_tab_crypt_md5)))) + uintptr(i)*4))
		a = _cgos_rol_crypt_md5(a, int32(4)) + b
		i++
		d += a ^ b ^ c + *(*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer((*uint32)(unsafe.Pointer(&W)))) + uintptr((uint32(3)*i+uint32(5))%uint32(16))*4)) + *(*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer((*uint32)(unsafe.Pointer(&_cgos_tab_crypt_md5)))) + uintptr(i)*4))
		d = _cgos_rol_crypt_md5(d, int32(11)) + a
		i++
		c += d ^ a ^ b + *(*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer((*uint32)(unsafe.Pointer(&W)))) + uintptr((uint32(3)*i+uint32(5))%uint32(16))*4)) + *(*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer((*uint32)(unsafe.Pointer(&_cgos_tab_crypt_md5)))) + uintptr(i)*4))
		c = _cgos_rol_crypt_md5(c, int32(16)) + d
		i++
		b += c ^ d ^ a + *(*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer((*uint32)(unsafe.Pointer(&W)))) + uintptr((uint32(3)*i+uint32(5))%uint32(16))*4)) + *(*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer((*uint32)(unsafe.Pointer(&_cgos_tab_crypt_md5)))) + uintptr(i)*4))
		b = _cgos_rol_crypt_md5(b, int32(23)) + c
		i++
	}
	for i < uint32(64) {
		a += c ^ (b | ^d) + *(*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer((*uint32)(unsafe.Pointer(&W)))) + uintptr(uint32(7)*i%uint32(16))*4)) + *(*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer((*uint32)(unsafe.Pointer(&_cgos_tab_crypt_md5)))) + uintptr(i)*4))
		a = _cgos_rol_crypt_md5(a, int32(6)) + b
		i++
		d += b ^ (a | ^c) + *(*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer((*uint32)(unsafe.Pointer(&W)))) + uintptr(uint32(7)*i%uint32(16))*4)) + *(*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer((*uint32)(unsafe.Pointer(&_cgos_tab_crypt_md5)))) + uintptr(i)*4))
		d = _cgos_rol_crypt_md5(d, int32(10)) + a
		i++
		c += a ^ (d | ^b) + *(*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer((*uint32)(unsafe.Pointer(&W)))) + uintptr(uint32(7)*i%uint32(16))*4)) + *(*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer((*uint32)(unsafe.Pointer(&_cgos_tab_crypt_md5)))) + uintptr(i)*4))
		c = _cgos_rol_crypt_md5(c, int32(15)) + d
		i++
		b += d ^ (c | ^a) + *(*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer((*uint32)(unsafe.Pointer(&W)))) + uintptr(uint32(7)*i%uint32(16))*4)) + *(*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer((*uint32)(unsafe.Pointer(&_cgos_tab_crypt_md5)))) + uintptr(i)*4))
		b = _cgos_rol_crypt_md5(b, int32(21)) + c
		i++
	}
	*(*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer((*uint32)(unsafe.Pointer(&s.h)))) + uintptr(int32(0))*4)) += a
	*(*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer((*uint32)(unsafe.Pointer(&s.h)))) + uintptr(int32(1))*4)) += b
	*(*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer((*uint32)(unsafe.Pointer(&s.h)))) + uintptr(int32(2))*4)) += c
	*(*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer((*uint32)(unsafe.Pointer(&s.h)))) + uintptr(int32(3))*4)) += d
}
func _cgos_pad_crypt_md5(s *_cgos_md5_crypt_md5) {
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
		_cgos_processblock_crypt_md5(s, (*uint8)(unsafe.Pointer(&s.buf)))
	}
	Memset(unsafe.Pointer((*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer((*uint8)(unsafe.Pointer(&s.buf))))+uintptr(r)))), int32(0), uint64(uint32(56)-r))
	s.len *= uint64(8)
	*(*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer((*uint8)(unsafe.Pointer(&s.buf)))) + uintptr(int32(56)))) = uint8(s.len)
	*(*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer((*uint8)(unsafe.Pointer(&s.buf)))) + uintptr(int32(57)))) = uint8(s.len >> int32(8))
	*(*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer((*uint8)(unsafe.Pointer(&s.buf)))) + uintptr(int32(58)))) = uint8(s.len >> int32(16))
	*(*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer((*uint8)(unsafe.Pointer(&s.buf)))) + uintptr(int32(59)))) = uint8(s.len >> int32(24))
	*(*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer((*uint8)(unsafe.Pointer(&s.buf)))) + uintptr(int32(60)))) = uint8(s.len >> int32(32))
	*(*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer((*uint8)(unsafe.Pointer(&s.buf)))) + uintptr(int32(61)))) = uint8(s.len >> int32(40))
	*(*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer((*uint8)(unsafe.Pointer(&s.buf)))) + uintptr(int32(62)))) = uint8(s.len >> int32(48))
	*(*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer((*uint8)(unsafe.Pointer(&s.buf)))) + uintptr(int32(63)))) = uint8(s.len >> int32(56))
	_cgos_processblock_crypt_md5(s, (*uint8)(unsafe.Pointer(&s.buf)))
}
func _cgos_md5_init_crypt_md5(s *_cgos_md5_crypt_md5) {
	s.len = uint64(0)
	*(*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer((*uint32)(unsafe.Pointer(&s.h)))) + uintptr(int32(0))*4)) = uint32(1732584193)
	*(*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer((*uint32)(unsafe.Pointer(&s.h)))) + uintptr(int32(1))*4)) = uint32(4023233417)
	*(*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer((*uint32)(unsafe.Pointer(&s.h)))) + uintptr(int32(2))*4)) = uint32(2562383102)
	*(*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer((*uint32)(unsafe.Pointer(&s.h)))) + uintptr(int32(3))*4)) = uint32(271733878)
}
func _cgos_md5_sum_crypt_md5(s *_cgos_md5_crypt_md5, md *uint8) {
	var i int32
	_cgos_pad_crypt_md5(s)
	for i = int32(0); i < int32(4); i++ {
		*(*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(md)) + uintptr(int32(4)*i))) = uint8(*(*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer((*uint32)(unsafe.Pointer(&s.h)))) + uintptr(i)*4)))
		*(*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(md)) + uintptr(int32(4)*i+int32(1)))) = uint8(*(*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer((*uint32)(unsafe.Pointer(&s.h)))) + uintptr(i)*4)) >> int32(8))
		*(*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(md)) + uintptr(int32(4)*i+int32(2)))) = uint8(*(*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer((*uint32)(unsafe.Pointer(&s.h)))) + uintptr(i)*4)) >> int32(16))
		*(*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(md)) + uintptr(int32(4)*i+int32(3)))) = uint8(*(*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer((*uint32)(unsafe.Pointer(&s.h)))) + uintptr(i)*4)) >> int32(24))
	}
}
func _cgos_md5_update_crypt_md5(s *_cgos_md5_crypt_md5, m unsafe.Pointer, len uint64) {
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
		_cgos_processblock_crypt_md5(s, (*uint8)(unsafe.Pointer(&s.buf)))
	}
	for ; len >= uint64(64); func() *uint8 {
		len -= uint64(64)
		return func() (_cgo_ret *uint8) {
			_cgo_addr := &p
			*(*uintptr)(unsafe.Pointer(&*_cgo_addr)) += uintptr(int32(64))
			return *_cgo_addr
		}()
	}() {
		_cgos_processblock_crypt_md5(s, p)
	}
	Memcpy(unsafe.Pointer((*uint8)(unsafe.Pointer(&s.buf))), unsafe.Pointer(p), len)
}

var _cgos_b64_crypt_md5 [65]uint8 = [65]uint8{'.', '/', '0', '1', '2', '3', '4', '5', '6', '7', '8', '9', 'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J', 'K', 'L', 'M', 'N', 'O', 'P', 'Q', 'R', 'S', 'T', 'U', 'V', 'W', 'X', 'Y', 'Z', 'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z', '\x00'}

func _cgos_to64_crypt_md5(s *int8, u uint32, n int32) *int8 {
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
		}() = int8(*(*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer((*uint8)(unsafe.Pointer(&_cgos_b64_crypt_md5)))) + uintptr(u%uint32(64)))))
		u /= uint32(64)
	}
	return s
}
func _cgos_md5crypt_crypt_md5(key *int8, setting *int8, output *int8) *int8 {
	var ctx _cgos_md5_crypt_md5
	var md [16]uint8
	var i uint32
	var klen uint32
	var slen uint32
	var salt *int8
	var p *int8
	klen = uint32(Strnlen(key, uint64(30001)))
	if klen > uint32(30000) {
		return (*int8)(nil)
	}
	if Strncmp(setting, (*int8)(unsafe.Pointer(&[4]int8{'$', '1', '$', '\x00'})), uint64(3)) != int32(0) {
		return (*int8)(nil)
	}
	salt = (*int8)(unsafe.Pointer(uintptr(unsafe.Pointer(setting)) + uintptr(int32(3))))
	for i = uint32(0); i < uint32(8) && int32(*(*int8)(unsafe.Pointer(uintptr(unsafe.Pointer(salt)) + uintptr(i)))) != 0 && int32(*(*int8)(unsafe.Pointer(uintptr(unsafe.Pointer(salt)) + uintptr(i)))) != '$'; i++ {
	}
	slen = i
	_cgos_md5_init_crypt_md5(&ctx)
	_cgos_md5_update_crypt_md5(&ctx, unsafe.Pointer(key), uint64(klen))
	_cgos_md5_update_crypt_md5(&ctx, unsafe.Pointer(salt), uint64(slen))
	_cgos_md5_update_crypt_md5(&ctx, unsafe.Pointer(key), uint64(klen))
	_cgos_md5_sum_crypt_md5(&ctx, (*uint8)(unsafe.Pointer(&md)))
	_cgos_md5_init_crypt_md5(&ctx)
	_cgos_md5_update_crypt_md5(&ctx, unsafe.Pointer(key), uint64(klen))
	_cgos_md5_update_crypt_md5(&ctx, unsafe.Pointer(setting), uint64(uint32(3)+slen))
	for i = klen; uint64(i) > 16; i -= uint32(16) {
		_cgos_md5_update_crypt_md5(&ctx, unsafe.Pointer((*uint8)(unsafe.Pointer(&md))), 16)
	}
	_cgos_md5_update_crypt_md5(&ctx, unsafe.Pointer((*uint8)(unsafe.Pointer(&md))), uint64(i))
	*(*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer((*uint8)(unsafe.Pointer(&md)))) + uintptr(int32(0)))) = uint8(0)
	for i = klen; i != 0; i >>= int32(1) {
		if i&uint32(1) != 0 {
			_cgos_md5_update_crypt_md5(&ctx, unsafe.Pointer((*uint8)(unsafe.Pointer(&md))), uint64(1))
		} else {
			_cgos_md5_update_crypt_md5(&ctx, unsafe.Pointer(key), uint64(1))
		}
	}
	_cgos_md5_sum_crypt_md5(&ctx, (*uint8)(unsafe.Pointer(&md)))
	for i = uint32(0); i < uint32(1000); i++ {
		_cgos_md5_init_crypt_md5(&ctx)
		if i%uint32(2) != 0 {
			_cgos_md5_update_crypt_md5(&ctx, unsafe.Pointer(key), uint64(klen))
		} else {
			_cgos_md5_update_crypt_md5(&ctx, unsafe.Pointer((*uint8)(unsafe.Pointer(&md))), 16)
		}
		if i%uint32(3) != 0 {
			_cgos_md5_update_crypt_md5(&ctx, unsafe.Pointer(salt), uint64(slen))
		}
		if i%uint32(7) != 0 {
			_cgos_md5_update_crypt_md5(&ctx, unsafe.Pointer(key), uint64(klen))
		}
		if i%uint32(2) != 0 {
			_cgos_md5_update_crypt_md5(&ctx, unsafe.Pointer((*uint8)(unsafe.Pointer(&md))), 16)
		} else {
			_cgos_md5_update_crypt_md5(&ctx, unsafe.Pointer(key), uint64(klen))
		}
		_cgos_md5_sum_crypt_md5(&ctx, (*uint8)(unsafe.Pointer(&md)))
	}
	Memcpy(unsafe.Pointer(output), unsafe.Pointer(setting), uint64(uint32(3)+slen))
	p = (*int8)(unsafe.Pointer(uintptr(unsafe.Pointer((*int8)(unsafe.Pointer(uintptr(unsafe.Pointer(output))+uintptr(int32(3)))))) + uintptr(slen)))
	*func() (_cgo_ret *int8) {
		_cgo_addr := &p
		_cgo_ret = *_cgo_addr
		*(*uintptr)(unsafe.Pointer(_cgo_addr))++
		return
	}() = int8('$')
	for i = uint32(0); i < uint32(5); i++ {
		p = _cgos_to64_crypt_md5(p, uint32(int32(*(*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer((*uint8)(unsafe.Pointer(&md)))) + uintptr(*(*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer((*uint8)(unsafe.Pointer(&*(*[3]uint8)(unsafe.Pointer(uintptr(unsafe.Pointer((*[3]uint8)(unsafe.Pointer(&_cgos_perm_crypt_md5)))) + uintptr(i)*3)))))) + uintptr(int32(0))))))))<<int32(16)|int32(*(*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer((*uint8)(unsafe.Pointer(&md)))) + uintptr(*(*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer((*uint8)(unsafe.Pointer(&*(*[3]uint8)(unsafe.Pointer(uintptr(unsafe.Pointer((*[3]uint8)(unsafe.Pointer(&_cgos_perm_crypt_md5)))) + uintptr(i)*3)))))) + uintptr(int32(1))))))))<<int32(8)|int32(*(*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer((*uint8)(unsafe.Pointer(&md)))) + uintptr(*(*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer((*uint8)(unsafe.Pointer(&*(*[3]uint8)(unsafe.Pointer(uintptr(unsafe.Pointer((*[3]uint8)(unsafe.Pointer(&_cgos_perm_crypt_md5)))) + uintptr(i)*3)))))) + uintptr(int32(2))))))))), int32(4))
	}
	p = _cgos_to64_crypt_md5(p, uint32(*(*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer((*uint8)(unsafe.Pointer(&md)))) + uintptr(int32(11))))), int32(2))
	*p = int8(0)
	return output
}

var _cgos_perm_crypt_md5 [5][3]uint8 = [5][3]uint8{[3]uint8{uint8(0), uint8(6), uint8(12)}, [3]uint8{uint8(1), uint8(7), uint8(13)}, [3]uint8{uint8(2), uint8(8), uint8(14)}, [3]uint8{uint8(3), uint8(9), uint8(15)}, [3]uint8{uint8(4), uint8(10), uint8(5)}}

func __crypt_md5(key *int8, setting *int8, output *int8) *int8 {
	var testbuf [64]int8
	var p *int8
	var q *int8
	p = _cgos_md5crypt_crypt_md5(key, setting, output)
	q = _cgos_md5crypt_crypt_md5((*int8)(unsafe.Pointer(&_cgos_testkey_crypt_md5)), (*int8)(unsafe.Pointer(&_cgos_testsetting_crypt_md5)), (*int8)(unsafe.Pointer(&testbuf)))
	if !(p != nil) || uintptr(unsafe.Pointer(q)) != uintptr(unsafe.Pointer((*int8)(unsafe.Pointer(&testbuf)))) || Memcmp(unsafe.Pointer((*int8)(unsafe.Pointer(&testbuf))), unsafe.Pointer((*int8)(unsafe.Pointer(&_cgos_testhash_crypt_md5))), 35) != 0 {
		return (*int8)(unsafe.Pointer(&[2]int8{'*', '\x00'}))
	}
	return p
}

var _cgos_testkey_crypt_md5 [18]int8 = [18]int8{'X', 'y', '0', '1', '@', '#', '\x01', '\x02', -128, '\u007f', -1, '\r', '\n', -127, '\t', ' ', '!', '\x00'}
var _cgos_testsetting_crypt_md5 [13]int8 = [13]int8{'$', '1', '$', 'a', 'b', 'c', 'd', '0', '1', '2', '3', '$', '\x00'}
var _cgos_testhash_crypt_md5 [35]int8 = [35]int8{'$', '1', '$', 'a', 'b', 'c', 'd', '0', '1', '2', '3', '$', '9', 'Q', 'c', 'g', '8', 'D', 'y', 'v', 'i', 'e', 'k', 'V', '3', 't', 'D', 'G', 'M', 'Z', 'y', 'n', 'J', '1', '\x00'}
