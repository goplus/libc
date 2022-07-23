package libc

import unsafe "unsafe"

type cookie_cgos__fmemopen struct {
	pos  uint64
	len  uint64
	size uint64
	buf  *uint8
	mode int32
}
type mem_FILE_cgos__fmemopen struct {
	f    struct__IO_FILE
	c    cookie_cgos__fmemopen
	buf  [1032]uint8
	buf2 [0]uint8
}

func mseek_cgos__fmemopen(f *struct__IO_FILE, off int64, whence int32) int64 {

	var base int64
	var c *cookie_cgos__fmemopen = (*cookie_cgos__fmemopen)(f.cookie)
	if !(uint32(whence) > uint32(2)) {
		goto _cgol_1
	}
fail:
	*__errno_location() = int32(22)
	return int64(-1)
_cgol_1:
	base = int64(*(*uint64)(unsafe.Pointer(uintptr(unsafe.Pointer((*uint64)(unsafe.Pointer(&[3]uint64{uint64(0), c.pos, c.len})))) + uintptr(whence)*8)))
	if !(off < int64(-base) || off > int64(int64(c.size)-base)) {
		goto _cgol_2
	}
	goto fail
_cgol_2:
	return int64(func() (_cgo_ret uint64) {
		_cgo_addr := &c.pos
		*_cgo_addr = uint64(int64(base) + off)
		return *_cgo_addr
	}())
}
func mread_cgos__fmemopen(f *struct__IO_FILE, buf *uint8, len uint64) uint64 {
	var c *cookie_cgos__fmemopen = (*cookie_cgos__fmemopen)(f.cookie)
	var rem uint64 = c.len - c.pos
	if c.pos > c.len {
		rem = uint64(0)
	}
	if len > rem {
		len = rem
		f.flags |= uint32(16)
	}
	Memcpy(unsafe.Pointer(buf), unsafe.Pointer((*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(c.buf))+uintptr(c.pos)))), len)
	c.pos += len
	rem -= len
	if rem > f.buf_size {
		rem = f.buf_size
	}
	f.rpos = f.buf
	f.rend = (*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(f.buf)) + uintptr(rem)))
	Memcpy(unsafe.Pointer(f.rpos), unsafe.Pointer((*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(c.buf))+uintptr(c.pos)))), rem)
	c.pos += rem
	return len
}
func mwrite_cgos__fmemopen(f *struct__IO_FILE, buf *uint8, len uint64) uint64 {
	var c *cookie_cgos__fmemopen = (*cookie_cgos__fmemopen)(f.cookie)
	var rem uint64
	var len2 uint64 = uint64(uintptr(unsafe.Pointer(f.wpos)) - uintptr(unsafe.Pointer(f.wbase)))
	if len2 != 0 {
		f.wpos = f.wbase
		if mwrite_cgos__fmemopen(f, f.wpos, len2) < len2 {
			return uint64(0)
		}
	}
	if c.mode == 'a' {
		c.pos = c.len
	}
	rem = c.size - c.pos
	if len > rem {
		len = rem
	}
	Memcpy(unsafe.Pointer((*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(c.buf))+uintptr(c.pos)))), unsafe.Pointer(buf), len)
	c.pos += len
	if c.pos > c.len {
		c.len = c.pos
		if c.len < c.size {
			*(*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(c.buf)) + uintptr(c.len))) = uint8(0)
		} else if f.flags&uint32(4) != 0 && c.size != 0 {
			*(*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(c.buf)) + uintptr(c.size-uint64(1)))) = uint8(0)
		}
	}
	return len
}
func mclose_cgos__fmemopen(m *struct__IO_FILE) int32 {
	return int32(0)
}
func Fmemopen(buf unsafe.Pointer, size uint64, mode *int8) *struct__IO_FILE {
	var f *mem_FILE_cgos__fmemopen
	var plus int32 = func() int32 {
		if !!(Strchr(mode, '+') != nil) {
			return 1
		} else {
			return 0
		}
	}()
	if !(Strchr((*int8)(unsafe.Pointer(&[4]int8{'r', 'w', 'a', '\x00'})), int32(*mode)) != nil) {
		*__errno_location() = int32(22)
		return (*struct__IO_FILE)(nil)
	}
	if !(buf != nil) && size > uint64(2147483647) {
		*__errno_location() = int32(12)
		return (*struct__IO_FILE)(nil)
	}
	f = (*mem_FILE_cgos__fmemopen)(Malloc(1300 + func() uint64 {
		if buf != nil {
			return uint64(0)
		} else {
			return size
		}
	}()))
	if !(f != nil) {
		return (*struct__IO_FILE)(nil)
	}
	Memset(unsafe.Pointer(f), int32(0), 268)
	f.f.cookie = unsafe.Pointer(&f.c)
	f.f.fd = -1
	f.f.lbf = -1
	f.f.buf = (*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer((*uint8)(unsafe.Pointer(&f.buf)))) + uintptr(int32(8))))
	f.f.buf_size = 1024
	if !(buf != nil) {
		buf = unsafe.Pointer((*uint8)(unsafe.Pointer(&f.buf2)))
		Memset(buf, int32(0), size)
	}
	f.c.buf = (*uint8)(buf)
	f.c.size = size
	f.c.mode = int32(*mode)
	if !(plus != 0) {
		f.f.flags = uint32(func() int32 {
			if int32(*mode) == 'r' {
				return int32(8)
			} else {
				return int32(4)
			}
		}())
	}
	if int32(*mode) == 'r' {
		f.c.len = size
	} else if int32(*mode) == 'a' {
		f.c.len = func() (_cgo_ret uint64) {
			_cgo_addr := &f.c.pos
			*_cgo_addr = Strnlen((*int8)(buf), size)
			return *_cgo_addr
		}()
	} else if plus != 0 {
		*f.c.buf = uint8(0)
	}
	f.f.read = mread_cgos__fmemopen
	f.f.write = mwrite_cgos__fmemopen
	f.f.seek = mseek_cgos__fmemopen
	f.f.close = mclose_cgos__fmemopen
	if !(__libc.threaded != 0) {
		f.f.lock = -1
	}
	return __ofl_add(&f.f)
}
