package libc

import unsafe "unsafe"

type _cgos_sha512_crypt_sha512 struct {
	len uint64
	h   [8]uint64
	buf [128]uint8
}

func _cgos_ror_crypt_sha512(n uint64, k int32) uint64 {
	return n>>k | n<<(int32(64)-k)
}

var _cgos_K_crypt_sha512 [80]uint64 = [80]uint64{uint64(4794697086780616226), uint64(8158064640168781261), uint64(13096744586834688815), uint64(16840607885511220156), uint64(4131703408338449720), uint64(6480981068601479193), uint64(10538285296894168987), uint64(12329834152419229976), uint64(15566598209576043074), uint64(1334009975649890238), uint64(2608012711638119052), uint64(6128411473006802146), uint64(8268148722764581231), uint64(9286055187155687089), uint64(11230858885718282805), uint64(13951009754708518548), uint64(16472876342353939154), uint64(17275323862435702243), uint64(1135362057144423861), uint64(2597628984639134821), uint64(3308224258029322869), uint64(5365058923640841347), uint64(6679025012923562964), uint64(8573033837759648693), uint64(10970295158949994411), uint64(12119686244451234320), uint64(12683024718118986047), uint64(13788192230050041572), uint64(14330467153632333762), uint64(15395433587784984357), uint64(489312712824947311), uint64(1452737877330783856), uint64(2861767655752347644), uint64(3322285676063803686), uint64(5560940570517711597), uint64(5996557281743188959), uint64(7280758554555802590), uint64(8532644243296465576), uint64(9350256976987008742), uint64(10552545826968843579), uint64(11727347734174303076), uint64(12113106623233404929), uint64(14000437183269869457), uint64(14369950271660146224), uint64(15101387698204529176), uint64(15463397548674623760), uint64(17586052441742319658), uint64(1182934255886127544), uint64(1847814050463011016), uint64(2177327727835720531), uint64(2830643537854262169), uint64(3796741975233480872), uint64(4115178125766777443), uint64(5681478168544905931), uint64(6601373596472566643), uint64(7507060721942968483), uint64(8399075790359081724), uint64(8693463985226723168), uint64(9568029438360202098), uint64(10144078919501101548), uint64(10430055236837252648), uint64(11840083180663258601), uint64(13761210420658862357), uint64(14299343276471374635), uint64(14566680578165727644), uint64(15097957966210449927), uint64(16922976911328602910), uint64(17689382322260857208), uint64(500013540394364858), uint64(748580250866718886), uint64(1242879168328830382), uint64(1977374033974150939), uint64(2944078676154940804), uint64(3659926193048069267), uint64(4368137639120453308), uint64(4836135668995329356), uint64(5532061633213252278), uint64(6448918945643986474), uint64(6902733635092675308), uint64(7801388544844847127)}

func _cgos_processblock_crypt_sha512(s *_cgos_sha512_crypt_sha512, buf *uint8) {
	var W [80]uint64
	var t1 uint64
	var t2 uint64
	var a uint64
	var b uint64
	var c uint64
	var d uint64
	var e uint64
	var f uint64
	var g uint64
	var h uint64
	var i int32
	for i = int32(0); i < int32(16); i++ {
		*(*uint64)(unsafe.Pointer(uintptr(unsafe.Pointer((*uint64)(unsafe.Pointer(&W)))) + uintptr(i)*8)) = uint64(*(*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(buf)) + uintptr(int32(8)*i)))) << int32(56)
		*(*uint64)(unsafe.Pointer(uintptr(unsafe.Pointer((*uint64)(unsafe.Pointer(&W)))) + uintptr(i)*8)) |= uint64(*(*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(buf)) + uintptr(int32(8)*i+int32(1))))) << int32(48)
		*(*uint64)(unsafe.Pointer(uintptr(unsafe.Pointer((*uint64)(unsafe.Pointer(&W)))) + uintptr(i)*8)) |= uint64(*(*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(buf)) + uintptr(int32(8)*i+int32(2))))) << int32(40)
		*(*uint64)(unsafe.Pointer(uintptr(unsafe.Pointer((*uint64)(unsafe.Pointer(&W)))) + uintptr(i)*8)) |= uint64(*(*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(buf)) + uintptr(int32(8)*i+int32(3))))) << int32(32)
		*(*uint64)(unsafe.Pointer(uintptr(unsafe.Pointer((*uint64)(unsafe.Pointer(&W)))) + uintptr(i)*8)) |= uint64(*(*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(buf)) + uintptr(int32(8)*i+int32(4))))) << int32(24)
		*(*uint64)(unsafe.Pointer(uintptr(unsafe.Pointer((*uint64)(unsafe.Pointer(&W)))) + uintptr(i)*8)) |= uint64(*(*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(buf)) + uintptr(int32(8)*i+int32(5))))) << int32(16)
		*(*uint64)(unsafe.Pointer(uintptr(unsafe.Pointer((*uint64)(unsafe.Pointer(&W)))) + uintptr(i)*8)) |= uint64(*(*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(buf)) + uintptr(int32(8)*i+int32(6))))) << int32(8)
		*(*uint64)(unsafe.Pointer(uintptr(unsafe.Pointer((*uint64)(unsafe.Pointer(&W)))) + uintptr(i)*8)) |= uint64(*(*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(buf)) + uintptr(int32(8)*i+int32(7)))))
	}
	for ; i < int32(80); i++ {
		*(*uint64)(unsafe.Pointer(uintptr(unsafe.Pointer((*uint64)(unsafe.Pointer(&W)))) + uintptr(i)*8)) = _cgos_ror_crypt_sha512(*(*uint64)(unsafe.Pointer(uintptr(unsafe.Pointer((*uint64)(unsafe.Pointer(&W)))) + uintptr(i-int32(2))*8)), int32(19)) ^ _cgos_ror_crypt_sha512(*(*uint64)(unsafe.Pointer(uintptr(unsafe.Pointer((*uint64)(unsafe.Pointer(&W)))) + uintptr(i-int32(2))*8)), int32(61)) ^ *(*uint64)(unsafe.Pointer(uintptr(unsafe.Pointer((*uint64)(unsafe.Pointer(&W)))) + uintptr(i-int32(2))*8))>>int32(6) + *(*uint64)(unsafe.Pointer(uintptr(unsafe.Pointer((*uint64)(unsafe.Pointer(&W)))) + uintptr(i-int32(7))*8)) + (_cgos_ror_crypt_sha512(*(*uint64)(unsafe.Pointer(uintptr(unsafe.Pointer((*uint64)(unsafe.Pointer(&W)))) + uintptr(i-int32(15))*8)), int32(1)) ^ _cgos_ror_crypt_sha512(*(*uint64)(unsafe.Pointer(uintptr(unsafe.Pointer((*uint64)(unsafe.Pointer(&W)))) + uintptr(i-int32(15))*8)), int32(8)) ^ *(*uint64)(unsafe.Pointer(uintptr(unsafe.Pointer((*uint64)(unsafe.Pointer(&W)))) + uintptr(i-int32(15))*8))>>int32(7)) + *(*uint64)(unsafe.Pointer(uintptr(unsafe.Pointer((*uint64)(unsafe.Pointer(&W)))) + uintptr(i-int32(16))*8))
	}
	a = *(*uint64)(unsafe.Pointer(uintptr(unsafe.Pointer((*uint64)(unsafe.Pointer(&s.h)))) + uintptr(int32(0))*8))
	b = *(*uint64)(unsafe.Pointer(uintptr(unsafe.Pointer((*uint64)(unsafe.Pointer(&s.h)))) + uintptr(int32(1))*8))
	c = *(*uint64)(unsafe.Pointer(uintptr(unsafe.Pointer((*uint64)(unsafe.Pointer(&s.h)))) + uintptr(int32(2))*8))
	d = *(*uint64)(unsafe.Pointer(uintptr(unsafe.Pointer((*uint64)(unsafe.Pointer(&s.h)))) + uintptr(int32(3))*8))
	e = *(*uint64)(unsafe.Pointer(uintptr(unsafe.Pointer((*uint64)(unsafe.Pointer(&s.h)))) + uintptr(int32(4))*8))
	f = *(*uint64)(unsafe.Pointer(uintptr(unsafe.Pointer((*uint64)(unsafe.Pointer(&s.h)))) + uintptr(int32(5))*8))
	g = *(*uint64)(unsafe.Pointer(uintptr(unsafe.Pointer((*uint64)(unsafe.Pointer(&s.h)))) + uintptr(int32(6))*8))
	h = *(*uint64)(unsafe.Pointer(uintptr(unsafe.Pointer((*uint64)(unsafe.Pointer(&s.h)))) + uintptr(int32(7))*8))
	for i = int32(0); i < int32(80); i++ {
		t1 = h + (_cgos_ror_crypt_sha512(e, int32(14)) ^ _cgos_ror_crypt_sha512(e, int32(18)) ^ _cgos_ror_crypt_sha512(e, int32(41))) + (g ^ e&(f^g)) + *(*uint64)(unsafe.Pointer(uintptr(unsafe.Pointer((*uint64)(unsafe.Pointer(&_cgos_K_crypt_sha512)))) + uintptr(i)*8)) + *(*uint64)(unsafe.Pointer(uintptr(unsafe.Pointer((*uint64)(unsafe.Pointer(&W)))) + uintptr(i)*8))
		t2 = _cgos_ror_crypt_sha512(a, int32(28)) ^ _cgos_ror_crypt_sha512(a, int32(34)) ^ _cgos_ror_crypt_sha512(a, int32(39)) + (a&b | c&(a|b))
		h = g
		g = f
		f = e
		e = d + t1
		d = c
		c = b
		b = a
		a = t1 + t2
	}
	*(*uint64)(unsafe.Pointer(uintptr(unsafe.Pointer((*uint64)(unsafe.Pointer(&s.h)))) + uintptr(int32(0))*8)) += a
	*(*uint64)(unsafe.Pointer(uintptr(unsafe.Pointer((*uint64)(unsafe.Pointer(&s.h)))) + uintptr(int32(1))*8)) += b
	*(*uint64)(unsafe.Pointer(uintptr(unsafe.Pointer((*uint64)(unsafe.Pointer(&s.h)))) + uintptr(int32(2))*8)) += c
	*(*uint64)(unsafe.Pointer(uintptr(unsafe.Pointer((*uint64)(unsafe.Pointer(&s.h)))) + uintptr(int32(3))*8)) += d
	*(*uint64)(unsafe.Pointer(uintptr(unsafe.Pointer((*uint64)(unsafe.Pointer(&s.h)))) + uintptr(int32(4))*8)) += e
	*(*uint64)(unsafe.Pointer(uintptr(unsafe.Pointer((*uint64)(unsafe.Pointer(&s.h)))) + uintptr(int32(5))*8)) += f
	*(*uint64)(unsafe.Pointer(uintptr(unsafe.Pointer((*uint64)(unsafe.Pointer(&s.h)))) + uintptr(int32(6))*8)) += g
	*(*uint64)(unsafe.Pointer(uintptr(unsafe.Pointer((*uint64)(unsafe.Pointer(&s.h)))) + uintptr(int32(7))*8)) += h
}
func _cgos_pad_crypt_sha512(s *_cgos_sha512_crypt_sha512) {
	var r uint32 = uint32(s.len % uint64(128))
	*(*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer((*uint8)(unsafe.Pointer(&s.buf)))) + uintptr(func() (_cgo_ret uint32) {
		_cgo_addr := &r
		_cgo_ret = *_cgo_addr
		*_cgo_addr++
		return
	}()))) = uint8(128)
	if r > uint32(112) {
		Memset(unsafe.Pointer((*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer((*uint8)(unsafe.Pointer(&s.buf))))+uintptr(r)))), int32(0), uint64(uint32(128)-r))
		r = uint32(0)
		_cgos_processblock_crypt_sha512(s, (*uint8)(unsafe.Pointer(&s.buf)))
	}
	Memset(unsafe.Pointer((*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer((*uint8)(unsafe.Pointer(&s.buf))))+uintptr(r)))), int32(0), uint64(uint32(120)-r))
	s.len *= uint64(8)
	*(*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer((*uint8)(unsafe.Pointer(&s.buf)))) + uintptr(int32(120)))) = uint8(s.len >> int32(56))
	*(*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer((*uint8)(unsafe.Pointer(&s.buf)))) + uintptr(int32(121)))) = uint8(s.len >> int32(48))
	*(*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer((*uint8)(unsafe.Pointer(&s.buf)))) + uintptr(int32(122)))) = uint8(s.len >> int32(40))
	*(*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer((*uint8)(unsafe.Pointer(&s.buf)))) + uintptr(int32(123)))) = uint8(s.len >> int32(32))
	*(*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer((*uint8)(unsafe.Pointer(&s.buf)))) + uintptr(int32(124)))) = uint8(s.len >> int32(24))
	*(*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer((*uint8)(unsafe.Pointer(&s.buf)))) + uintptr(int32(125)))) = uint8(s.len >> int32(16))
	*(*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer((*uint8)(unsafe.Pointer(&s.buf)))) + uintptr(int32(126)))) = uint8(s.len >> int32(8))
	*(*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer((*uint8)(unsafe.Pointer(&s.buf)))) + uintptr(int32(127)))) = uint8(s.len)
	_cgos_processblock_crypt_sha512(s, (*uint8)(unsafe.Pointer(&s.buf)))
}
func _cgos_sha512_init_crypt_sha512(s *_cgos_sha512_crypt_sha512) {
	s.len = uint64(0)
	*(*uint64)(unsafe.Pointer(uintptr(unsafe.Pointer((*uint64)(unsafe.Pointer(&s.h)))) + uintptr(int32(0))*8)) = uint64(7640891576956012808)
	*(*uint64)(unsafe.Pointer(uintptr(unsafe.Pointer((*uint64)(unsafe.Pointer(&s.h)))) + uintptr(int32(1))*8)) = uint64(13503953896175478587)
	*(*uint64)(unsafe.Pointer(uintptr(unsafe.Pointer((*uint64)(unsafe.Pointer(&s.h)))) + uintptr(int32(2))*8)) = uint64(4354685564936845355)
	*(*uint64)(unsafe.Pointer(uintptr(unsafe.Pointer((*uint64)(unsafe.Pointer(&s.h)))) + uintptr(int32(3))*8)) = uint64(11912009170470909681)
	*(*uint64)(unsafe.Pointer(uintptr(unsafe.Pointer((*uint64)(unsafe.Pointer(&s.h)))) + uintptr(int32(4))*8)) = uint64(5840696475078001361)
	*(*uint64)(unsafe.Pointer(uintptr(unsafe.Pointer((*uint64)(unsafe.Pointer(&s.h)))) + uintptr(int32(5))*8)) = uint64(11170449401992604703)
	*(*uint64)(unsafe.Pointer(uintptr(unsafe.Pointer((*uint64)(unsafe.Pointer(&s.h)))) + uintptr(int32(6))*8)) = uint64(2270897969802886507)
	*(*uint64)(unsafe.Pointer(uintptr(unsafe.Pointer((*uint64)(unsafe.Pointer(&s.h)))) + uintptr(int32(7))*8)) = uint64(6620516959819538809)
}
func _cgos_sha512_sum_crypt_sha512(s *_cgos_sha512_crypt_sha512, md *uint8) {
	var i int32
	_cgos_pad_crypt_sha512(s)
	for i = int32(0); i < int32(8); i++ {
		*(*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(md)) + uintptr(int32(8)*i))) = uint8(*(*uint64)(unsafe.Pointer(uintptr(unsafe.Pointer((*uint64)(unsafe.Pointer(&s.h)))) + uintptr(i)*8)) >> int32(56))
		*(*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(md)) + uintptr(int32(8)*i+int32(1)))) = uint8(*(*uint64)(unsafe.Pointer(uintptr(unsafe.Pointer((*uint64)(unsafe.Pointer(&s.h)))) + uintptr(i)*8)) >> int32(48))
		*(*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(md)) + uintptr(int32(8)*i+int32(2)))) = uint8(*(*uint64)(unsafe.Pointer(uintptr(unsafe.Pointer((*uint64)(unsafe.Pointer(&s.h)))) + uintptr(i)*8)) >> int32(40))
		*(*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(md)) + uintptr(int32(8)*i+int32(3)))) = uint8(*(*uint64)(unsafe.Pointer(uintptr(unsafe.Pointer((*uint64)(unsafe.Pointer(&s.h)))) + uintptr(i)*8)) >> int32(32))
		*(*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(md)) + uintptr(int32(8)*i+int32(4)))) = uint8(*(*uint64)(unsafe.Pointer(uintptr(unsafe.Pointer((*uint64)(unsafe.Pointer(&s.h)))) + uintptr(i)*8)) >> int32(24))
		*(*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(md)) + uintptr(int32(8)*i+int32(5)))) = uint8(*(*uint64)(unsafe.Pointer(uintptr(unsafe.Pointer((*uint64)(unsafe.Pointer(&s.h)))) + uintptr(i)*8)) >> int32(16))
		*(*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(md)) + uintptr(int32(8)*i+int32(6)))) = uint8(*(*uint64)(unsafe.Pointer(uintptr(unsafe.Pointer((*uint64)(unsafe.Pointer(&s.h)))) + uintptr(i)*8)) >> int32(8))
		*(*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(md)) + uintptr(int32(8)*i+int32(7)))) = uint8(*(*uint64)(unsafe.Pointer(uintptr(unsafe.Pointer((*uint64)(unsafe.Pointer(&s.h)))) + uintptr(i)*8)))
	}
}
func _cgos_sha512_update_crypt_sha512(s *_cgos_sha512_crypt_sha512, m unsafe.Pointer, len uint64) {
	var p *uint8 = (*uint8)(m)
	var r uint32 = uint32(s.len % uint64(128))
	s.len += uint64(len)
	if r != 0 {
		if len < uint64(uint32(128)-r) {
			Memcpy(unsafe.Pointer((*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer((*uint8)(unsafe.Pointer(&s.buf))))+uintptr(r)))), unsafe.Pointer(p), len)
			return
		}
		Memcpy(unsafe.Pointer((*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer((*uint8)(unsafe.Pointer(&s.buf))))+uintptr(r)))), unsafe.Pointer(p), uint64(uint32(128)-r))
		len -= uint64(uint32(128) - r)
		*(*uintptr)(unsafe.Pointer(&p)) += uintptr(uint32(128) - r)
		_cgos_processblock_crypt_sha512(s, (*uint8)(unsafe.Pointer(&s.buf)))
	}
	for ; len >= uint64(128); func() *uint8 {
		len -= uint64(128)
		return func() (_cgo_ret *uint8) {
			_cgo_addr := &p
			*(*uintptr)(unsafe.Pointer(&*_cgo_addr)) += uintptr(int32(128))
			return *_cgo_addr
		}()
	}() {
		_cgos_processblock_crypt_sha512(s, p)
	}
	Memcpy(unsafe.Pointer((*uint8)(unsafe.Pointer(&s.buf))), unsafe.Pointer(p), len)
}

var _cgos_b64_crypt_sha512 [65]uint8 = [65]uint8{'.', '/', '0', '1', '2', '3', '4', '5', '6', '7', '8', '9', 'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J', 'K', 'L', 'M', 'N', 'O', 'P', 'Q', 'R', 'S', 'T', 'U', 'V', 'W', 'X', 'Y', 'Z', 'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z', '\x00'}

func _cgos_to64_crypt_sha512(s *int8, u uint32, n int32) *int8 {
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
		}() = int8(*(*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer((*uint8)(unsafe.Pointer(&_cgos_b64_crypt_sha512)))) + uintptr(u%uint32(64)))))
		u /= uint32(64)
	}
	return s
}
func _cgos_hashmd_crypt_sha512(s *_cgos_sha512_crypt_sha512, n uint32, md unsafe.Pointer) {
	var i uint32
	for i = n; i > uint32(64); i -= uint32(64) {
		_cgos_sha512_update_crypt_sha512(s, md, uint64(64))
	}
	_cgos_sha512_update_crypt_sha512(s, md, uint64(i))
}
func _cgos_sha512crypt_crypt_sha512(key *int8, setting *int8, output *int8) *int8 {
	var ctx _cgos_sha512_crypt_sha512
	var md [64]uint8
	var kmd [64]uint8
	var smd [64]uint8
	var i uint32
	var r uint32
	var klen uint32
	var slen uint32
	var rounds [20]int8 = [20]int8{'\x00'}
	var salt *int8
	var p *int8
	for i = uint32(0); i <= uint32(256) && int32(*(*int8)(unsafe.Pointer(uintptr(unsafe.Pointer(key)) + uintptr(i)))) != 0; i++ {
	}
	if i > uint32(256) {
		return (*int8)(nil)
	}
	klen = i
	if Strncmp(setting, (*int8)(unsafe.Pointer(&[4]int8{'$', '6', '$', '\x00'})), uint64(3)) != int32(0) {
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
	_cgos_sha512_init_crypt_sha512(&ctx)
	_cgos_sha512_update_crypt_sha512(&ctx, unsafe.Pointer(key), uint64(klen))
	_cgos_sha512_update_crypt_sha512(&ctx, unsafe.Pointer(salt), uint64(slen))
	_cgos_sha512_update_crypt_sha512(&ctx, unsafe.Pointer(key), uint64(klen))
	_cgos_sha512_sum_crypt_sha512(&ctx, (*uint8)(unsafe.Pointer(&md)))
	_cgos_sha512_init_crypt_sha512(&ctx)
	_cgos_sha512_update_crypt_sha512(&ctx, unsafe.Pointer(key), uint64(klen))
	_cgos_sha512_update_crypt_sha512(&ctx, unsafe.Pointer(salt), uint64(slen))
	_cgos_hashmd_crypt_sha512(&ctx, klen, unsafe.Pointer((*uint8)(unsafe.Pointer(&md))))
	for i = klen; i > uint32(0); i >>= int32(1) {
		if i&uint32(1) != 0 {
			_cgos_sha512_update_crypt_sha512(&ctx, unsafe.Pointer((*uint8)(unsafe.Pointer(&md))), 64)
		} else {
			_cgos_sha512_update_crypt_sha512(&ctx, unsafe.Pointer(key), uint64(klen))
		}
	}
	_cgos_sha512_sum_crypt_sha512(&ctx, (*uint8)(unsafe.Pointer(&md)))
	_cgos_sha512_init_crypt_sha512(&ctx)
	for i = uint32(0); i < klen; i++ {
		_cgos_sha512_update_crypt_sha512(&ctx, unsafe.Pointer(key), uint64(klen))
	}
	_cgos_sha512_sum_crypt_sha512(&ctx, (*uint8)(unsafe.Pointer(&kmd)))
	_cgos_sha512_init_crypt_sha512(&ctx)
	for i = uint32(0); i < uint32(int32(16)+int32(*(*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer((*uint8)(unsafe.Pointer(&md)))) + uintptr(int32(0)))))); i++ {
		_cgos_sha512_update_crypt_sha512(&ctx, unsafe.Pointer(salt), uint64(slen))
	}
	_cgos_sha512_sum_crypt_sha512(&ctx, (*uint8)(unsafe.Pointer(&smd)))
	for i = uint32(0); i < r; i++ {
		_cgos_sha512_init_crypt_sha512(&ctx)
		if i%uint32(2) != 0 {
			_cgos_hashmd_crypt_sha512(&ctx, klen, unsafe.Pointer((*uint8)(unsafe.Pointer(&kmd))))
		} else {
			_cgos_sha512_update_crypt_sha512(&ctx, unsafe.Pointer((*uint8)(unsafe.Pointer(&md))), 64)
		}
		if i%uint32(3) != 0 {
			_cgos_sha512_update_crypt_sha512(&ctx, unsafe.Pointer((*uint8)(unsafe.Pointer(&smd))), uint64(slen))
		}
		if i%uint32(7) != 0 {
			_cgos_hashmd_crypt_sha512(&ctx, klen, unsafe.Pointer((*uint8)(unsafe.Pointer(&kmd))))
		}
		if i%uint32(2) != 0 {
			_cgos_sha512_update_crypt_sha512(&ctx, unsafe.Pointer((*uint8)(unsafe.Pointer(&md))), 64)
		} else {
			_cgos_hashmd_crypt_sha512(&ctx, klen, unsafe.Pointer((*uint8)(unsafe.Pointer(&kmd))))
		}
		_cgos_sha512_sum_crypt_sha512(&ctx, (*uint8)(unsafe.Pointer(&md)))
	}
	p = output
	*(*uintptr)(unsafe.Pointer(&p)) += uintptr(Sprintf(p, (*int8)(unsafe.Pointer(&[11]int8{'$', '6', '$', '%', 's', '%', '.', '*', 's', '$', '\x00'})), (*int8)(unsafe.Pointer(&rounds)), slen, salt))
	for i = uint32(0); i < uint32(21); i++ {
		p = _cgos_to64_crypt_sha512(p, uint32(int32(*(*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer((*uint8)(unsafe.Pointer(&md)))) + uintptr(*(*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer((*uint8)(unsafe.Pointer(&*(*[3]uint8)(unsafe.Pointer(uintptr(unsafe.Pointer((*[3]uint8)(unsafe.Pointer(&_cgos_perm_crypt_sha512)))) + uintptr(i)*3)))))) + uintptr(int32(0))))))))<<int32(16)|int32(*(*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer((*uint8)(unsafe.Pointer(&md)))) + uintptr(*(*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer((*uint8)(unsafe.Pointer(&*(*[3]uint8)(unsafe.Pointer(uintptr(unsafe.Pointer((*[3]uint8)(unsafe.Pointer(&_cgos_perm_crypt_sha512)))) + uintptr(i)*3)))))) + uintptr(int32(1))))))))<<int32(8)|int32(*(*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer((*uint8)(unsafe.Pointer(&md)))) + uintptr(*(*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer((*uint8)(unsafe.Pointer(&*(*[3]uint8)(unsafe.Pointer(uintptr(unsafe.Pointer((*[3]uint8)(unsafe.Pointer(&_cgos_perm_crypt_sha512)))) + uintptr(i)*3)))))) + uintptr(int32(2))))))))), int32(4))
	}
	p = _cgos_to64_crypt_sha512(p, uint32(*(*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer((*uint8)(unsafe.Pointer(&md)))) + uintptr(int32(63))))), int32(2))
	*p = int8(0)
	return output
}

var _cgos_perm_crypt_sha512 [21][3]uint8 = [21][3]uint8{[3]uint8{uint8(0), uint8(21), uint8(42)}, [3]uint8{uint8(22), uint8(43), uint8(1)}, [3]uint8{uint8(44), uint8(2), uint8(23)}, [3]uint8{uint8(3), uint8(24), uint8(45)}, [3]uint8{uint8(25), uint8(46), uint8(4)}, [3]uint8{uint8(47), uint8(5), uint8(26)}, [3]uint8{uint8(6), uint8(27), uint8(48)}, [3]uint8{uint8(28), uint8(49), uint8(7)}, [3]uint8{uint8(50), uint8(8), uint8(29)}, [3]uint8{uint8(9), uint8(30), uint8(51)}, [3]uint8{uint8(31), uint8(52), uint8(10)}, [3]uint8{uint8(53), uint8(11), uint8(32)}, [3]uint8{uint8(12), uint8(33), uint8(54)}, [3]uint8{uint8(34), uint8(55), uint8(13)}, [3]uint8{uint8(56), uint8(14), uint8(35)}, [3]uint8{uint8(15), uint8(36), uint8(57)}, [3]uint8{uint8(37), uint8(58), uint8(16)}, [3]uint8{uint8(59), uint8(17), uint8(38)}, [3]uint8{uint8(18), uint8(39), uint8(60)}, [3]uint8{uint8(40), uint8(61), uint8(19)}, [3]uint8{uint8(62), uint8(20), uint8(41)}}

func __crypt_sha512(key *int8, setting *int8, output *int8) *int8 {
	var testbuf [128]int8
	var p *int8
	var q *int8
	p = _cgos_sha512crypt_crypt_sha512(key, setting, output)
	q = _cgos_sha512crypt_crypt_sha512((*int8)(unsafe.Pointer(&_cgos_testkey_crypt_sha512)), (*int8)(unsafe.Pointer(&_cgos_testsetting_crypt_sha512)), (*int8)(unsafe.Pointer(&testbuf)))
	if !(p != nil) || uintptr(unsafe.Pointer(q)) != uintptr(unsafe.Pointer((*int8)(unsafe.Pointer(&testbuf)))) || Memcmp(unsafe.Pointer((*int8)(unsafe.Pointer(&testbuf))), unsafe.Pointer((*int8)(unsafe.Pointer(&_cgos_testhash_crypt_sha512))), 116) != 0 {
		return (*int8)(unsafe.Pointer(&[2]int8{'*', '\x00'}))
	}
	return p
}

var _cgos_testkey_crypt_sha512 [18]int8 = [18]int8{'X', 'y', '0', '1', '@', '#', '\x01', '\x02', -128, '\u007f', -1, '\r', '\n', -127, '\t', ' ', '!', '\x00'}
var _cgos_testsetting_crypt_sha512 [30]int8 = [30]int8{'$', '6', '$', 'r', 'o', 'u', 'n', 'd', 's', '=', '1', '2', '3', '4', '$', 'a', 'b', 'c', '0', '1', '2', '3', '4', '5', '6', '7', '8', '9', '$', '\x00'}
var _cgos_testhash_crypt_sha512 [116]int8 = [116]int8{'$', '6', '$', 'r', 'o', 'u', 'n', 'd', 's', '=', '1', '2', '3', '4', '$', 'a', 'b', 'c', '0', '1', '2', '3', '4', '5', '6', '7', '8', '9', '$', 'B', 'C', 'p', 't', '8', 'z', 'L', 'r', 'c', '/', 'R', 'c', 'y', 'u', 'X', 'm', 'C', 'D', 'O', 'E', '1', 'A', 'L', 'q', 'M', 'X', 'B', '2', 'M', 'H', '6', 'n', '1', 'g', '8', '9', '1', 'H', 'h', 'F', 'j', '8', '.', 'w', '7', 'L', 'x', 'G', 'v', '.', 'F', 'T', 'k', 'q', 'q', '6', 'V', 'x', 'c', '/', 'k', 'm', '3', 'Y', '0', 'j', 'E', '0', 'j', '2', '4', 'j', 'Y', '5', 'P', 'I', 'v', '/', 'o', 'O', 'u', '6', 'r', 'e', 'g', '1', '\x00'}
