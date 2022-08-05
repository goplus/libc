package libc

import unsafe "unsafe"

type _cgos_cookie_vswprintf struct {
	ws *uint32
	l  uint64
}

func _cgos_sw_write_vswprintf(f *Struct__IO_FILE, s *uint8, l uint64) uint64 {
	var l0 uint64 = l
	var i int32 = int32(0)
	var c *_cgos_cookie_vswprintf = (*_cgos_cookie_vswprintf)(f.Cookie)
	if uintptr(unsafe.Pointer(s)) != uintptr(unsafe.Pointer((*uint8)(unsafe.Pointer(f.Wbase)))) && _cgos_sw_write_vswprintf(f, f.Wbase, uint64(uintptr(unsafe.Pointer(f.Wpos))-uintptr(unsafe.Pointer(f.Wbase)))) == uint64(18446744073709551615) {
		return uint64(18446744073709551615)
	}
	for c.l != 0 && l != 0 && func() (_cgo_ret int32) {
		_cgo_addr := &i
		*_cgo_addr = Mbtowc(c.ws, (*int8)(unsafe.Pointer(s)), l)
		return *_cgo_addr
	}() >= int32(0) {
		*(*uintptr)(unsafe.Pointer(&s)) += uintptr(i)
		l -= uint64(i)
		c.l--
		*(*uintptr)(unsafe.Pointer(&c.ws)) += 4
	}
	*c.ws = uint32(0)
	if i < int32(0) {
		f.Wpos = func() (_cgo_ret *uint8) {
			_cgo_addr := &f.Wbase
			*_cgo_addr = func() (_cgo_ret *uint8) {
				_cgo_addr := &f.Wend
				*_cgo_addr = (*uint8)(nil)
				return *_cgo_addr
			}()
			return *_cgo_addr
		}()
		f.Flags |= uint32(32)
		return uint64(i)
	}
	f.Wend = (*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(f.Buf)) + uintptr(f.Buf_size)))
	f.Wpos = func() (_cgo_ret *uint8) {
		_cgo_addr := &f.Wbase
		*_cgo_addr = f.Buf
		return *_cgo_addr
	}()
	return l0
}
func vswprintf(s *uint32, n uint64, fmt *uint32, ap []interface {
}) int32 {
	var r int32
	var buf [256]uint8
	var c _cgos_cookie_vswprintf = _cgos_cookie_vswprintf{s, n - uint64(1)}
	var f Struct__IO_FILE = Struct__IO_FILE{0, nil, nil, nil, nil, nil, nil, nil, nil, _cgos_sw_write_vswprintf, nil, (*uint8)(unsafe.Pointer(&buf)), 256, nil, nil, 0, 0, 0, 0, -1, -1, unsafe.Pointer(&c), 0, nil, nil, nil, 0, 0, nil, nil, nil}
	if !(n != 0) {
		return -1
	} else if n > uint64(2147483647) {
		*X__errno_location() = int32(75)
		return -1
	}
	r = vfwprintf(&f, fmt, ap)
	_cgos_sw_write_vswprintf(&f, nil, uint64(0))
	return func() int32 {
		if uint64(r) >= n {
			return -1
		} else {
			return r
		}
	}()
}
