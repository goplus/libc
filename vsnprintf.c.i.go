package libc

import unsafe "unsafe"

type struct_cookie struct {
	s *int8
	n uint64
}

func sn_write_cgo455(f *struct__IO_FILE, s *uint8, l uint64) uint64 {
	var c *struct_cookie = (*struct_cookie)(f.cookie)
	var k uint64 = func() uint64 {
		if c.n < uint64(uintptr(unsafe.Pointer(f.wpos))-uintptr(unsafe.Pointer(f.wbase))) {
			return c.n
		} else {
			return uint64(uintptr(unsafe.Pointer(f.wpos)) - uintptr(unsafe.Pointer(f.wbase)))
		}
	}()
	if k != 0 {
		Memcpy(unsafe.Pointer(c.s), unsafe.Pointer(f.wbase), k)
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
	f.wpos = func() (_cgo_ret *uint8) {
		_cgo_addr := &f.wbase
		*_cgo_addr = f.buf
		return *_cgo_addr
	}()
	return l
}
func Vsnprintf(s *int8, n uint64, fmt *int8, ap []interface {
}) int32 {
	var buf [1]uint8
	var dummy [1]int8
	var c struct_cookie = struct_cookie{func() *int8 {
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
	var f struct__IO_FILE = struct__IO_FILE{0, nil, nil, nil, nil, nil, nil, nil, nil, sn_write_cgo455, nil, (*uint8)(unsafe.Pointer(&buf)), 0, nil, nil, 0, 0, 0, 0, -1, -1, unsafe.Pointer(&c), 0, nil, nil, nil, 0, 0, nil, nil, nil}
	if n > uint64(2147483647) {
		*__errno_location() = int32(75)
		return int32(-1)
	}
	*c.s = int8(0)
	return Vfprintf(&f, fmt, ap)
}
