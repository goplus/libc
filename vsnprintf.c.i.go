package libc

import unsafe "unsafe"

type _cgos_cookie_vsnprintf struct {
	s *int8
	n uint64
}

func _cgos_sn_write_vsnprintf(f *Struct__IO_FILE, s *uint8, l uint64) uint64 {
	var c *_cgos_cookie_vsnprintf = (*_cgos_cookie_vsnprintf)(f.Cookie)
	var k uint64 = func() uint64 {
		if c.n < uint64(uintptr(unsafe.Pointer(f.Wpos))-uintptr(unsafe.Pointer(f.Wbase))) {
			return c.n
		} else {
			return uint64(uintptr(unsafe.Pointer(f.Wpos)) - uintptr(unsafe.Pointer(f.Wbase)))
		}
	}()
	if k != 0 {
		Memcpy(unsafe.Pointer(c.s), unsafe.Pointer(f.Wbase), k)
		*(*uintptr)(unsafe.Pointer(&c.s)) += uintptr(k)
		c.n -= k
	}
	k = func() uint64 {
		if c.n < l {
			return c.n
		} else {
			return l
		}
	}()
	if k != 0 {
		Memcpy(unsafe.Pointer(c.s), unsafe.Pointer(s), k)
		*(*uintptr)(unsafe.Pointer(&c.s)) += uintptr(k)
		c.n -= k
	}
	*c.s = int8(0)
	f.Wpos = func() (_cgo_ret *uint8) {
		_cgo_addr := &f.Wbase
		*_cgo_addr = f.Buf
		return *_cgo_addr
	}()
	return l
}
func Vsnprintf(s *int8, n uint64, fmt *int8, ap []interface {
}) int32 {
	var buf [1]uint8
	var dummy [1]int8
	var c _cgos_cookie_vsnprintf = _cgos_cookie_vsnprintf{func() *int8 {
		if n != 0 {
			return s
		} else {
			return (*int8)(unsafe.Pointer(&dummy))
		}
	}(), func() uint64 {
		if n != 0 {
			return n - uint64(1)
		} else {
			return uint64(0)
		}
	}()}
	var f Struct__IO_FILE = Struct__IO_FILE{0, nil, nil, nil, nil, nil, nil, nil, nil, _cgos_sn_write_vsnprintf, nil, (*uint8)(unsafe.Pointer(&buf)), 0, nil, nil, 0, 0, 0, 0, -1, -1, unsafe.Pointer(&c), 0, nil, nil, nil, 0, 0, nil, nil, nil}
	if n > uint64(2147483647) {
		*__errno_location() = int32(75)
		return -1
	}
	*c.s = int8(0)
	return Vfprintf(&f, fmt, ap)
}
