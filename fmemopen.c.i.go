package libc

import unsafe "unsafe"

type _cgos_cookie_fmemopen struct {
	pos  uint64
	len  uint64
	size uint64
	buf  *uint8
	mode int32
}
type _cgos_mem_FILE_fmemopen struct {
	f    Struct__IO_FILE
	c    _cgos_cookie_fmemopen
	buf  [1032]uint8
	buf2 [0]uint8
}

func _cgos_mseek_fmemopen(f *Struct__IO_FILE, off int64, whence int32) int64 {

	var base int64
	var c *_cgos_cookie_fmemopen = (*_cgos_cookie_fmemopen)(f.Cookie)
	if !(uint32(whence) > uint32(2)) {
		goto _cgol_1
	}
fail:
	*X__errno_location() = int32(22)
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
func _cgos_mread_fmemopen(f *Struct__IO_FILE, buf *uint8, len uint64) uint64 {
	var c *_cgos_cookie_fmemopen = (*_cgos_cookie_fmemopen)(f.Cookie)
	var rem uint64 = c.len - c.pos
	if c.pos > c.len {
		rem = uint64(0)
	}
	if len > rem {
		len = rem
		f.Flags |= uint32(16)
	}
	Memcpy(unsafe.Pointer(buf), unsafe.Pointer((*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(c.buf))+uintptr(c.pos)))), len)
	c.pos += len
	rem -= len
	if rem > f.Buf_size {
		rem = f.Buf_size
	}
	f.Rpos = f.Buf
	f.Rend = (*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(f.Buf)) + uintptr(rem)))
	Memcpy(unsafe.Pointer(f.Rpos), unsafe.Pointer((*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(c.buf))+uintptr(c.pos)))), rem)
	c.pos += rem
	return len
}
func _cgos_mwrite_fmemopen(f *Struct__IO_FILE, buf *uint8, len uint64) uint64 {
	var c *_cgos_cookie_fmemopen = (*_cgos_cookie_fmemopen)(f.Cookie)
	var rem uint64
	var len2 uint64 = uint64(uintptr(unsafe.Pointer(f.Wpos)) - uintptr(unsafe.Pointer(f.Wbase)))
	if len2 != 0 {
		f.Wpos = f.Wbase
		if _cgos_mwrite_fmemopen(f, f.Wpos, len2) < len2 {
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
		} else if f.Flags&uint32(4) != 0 && c.size != 0 {
			*(*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(c.buf)) + uintptr(c.size-uint64(1)))) = uint8(0)
		}
	}
	return len
}
func _cgos_mclose_fmemopen(m *Struct__IO_FILE) int32 {
	return int32(0)
}
func Fmemopen(buf unsafe.Pointer, size uint64, mode *int8) *Struct__IO_FILE {
	var f *_cgos_mem_FILE_fmemopen
	var plus int32 = func() int32 {
		if !!(Strchr(mode, '+') != nil) {
			return 1
		} else {
			return 0
		}
	}()
	if !(Strchr((*int8)(unsafe.Pointer(&[4]int8{'r', 'w', 'a', '\x00'})), int32(*mode)) != nil) {
		*X__errno_location() = int32(22)
		return (*Struct__IO_FILE)(nil)
	}
	if !(buf != nil) && size > uint64(2147483647) {
		*X__errno_location() = int32(12)
		return (*Struct__IO_FILE)(nil)
	}
	f = (*_cgos_mem_FILE_fmemopen)(Malloc(1300 + func() uint64 {
		if buf != nil {
			return uint64(0)
		} else {
			return size
		}
	}()))
	if !(f != nil) {
		return (*Struct__IO_FILE)(nil)
	}
	Memset(unsafe.Pointer(f), int32(0), 268)
	f.f.Cookie = unsafe.Pointer(&f.c)
	f.f.Fd = -1
	f.f.Lbf = -1
	f.f.Buf = (*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer((*uint8)(unsafe.Pointer(&f.buf)))) + uintptr(int32(8))))
	f.f.Buf_size = 1024
	if !(buf != nil) {
		buf = unsafe.Pointer((*uint8)(unsafe.Pointer(&f.buf2)))
		Memset(buf, int32(0), size)
	}
	f.c.buf = (*uint8)(buf)
	f.c.size = size
	f.c.mode = int32(*mode)
	if !(plus != 0) {
		f.f.Flags = uint32(func() int32 {
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
	f.f.Read = _cgos_mread_fmemopen
	f.f.Write = _cgos_mwrite_fmemopen
	f.f.Seek = _cgos_mseek_fmemopen
	f.f.Close = _cgos_mclose_fmemopen
	if !(__libc.threaded != 0) {
		f.f.Lock = -1
	}
	return __ofl_add(&f.f)
}
