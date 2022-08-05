package libc

import unsafe "unsafe"

type _cgos_fcookie_fopencookie struct {
	cookie  unsafe.Pointer
	iofuncs Struct__IO_cookie_io_functions_t
}
type _cgos_cookie_FILE_fopencookie struct {
	f   Struct__IO_FILE
	fc  _cgos_fcookie_fopencookie
	buf [1032]uint8
}

func _cgos_cookieread_fopencookie(f *Struct__IO_FILE, buf *uint8, len uint64) uint64 {
	var fc *_cgos_fcookie_fopencookie = (*_cgos_fcookie_fopencookie)(f.Cookie)
	var ret int64 = int64(-1)
	var remain uint64 = len
	var readlen uint64 = uint64(0)
	var len2 uint64 = len - func() uint64 {
		if !!(f.Buf_size != 0) {
			return 1
		} else {
			return 0
		}
	}()
	if !(fc.iofuncs.Read != nil) {
		goto bail
	}
	if len2 != 0 {
		ret = fc.iofuncs.Read(fc.cookie, (*int8)(unsafe.Pointer(buf)), len2)
		if ret <= int64(0) {
			goto bail
		}
		readlen += uint64(ret)
		remain -= uint64(ret)
	}
	if !(f.Buf_size != 0) || remain > func() uint64 {
		if !!(f.Buf_size != 0) {
			return 1
		} else {
			return 0
		}
	}() {
		return readlen
	}
	f.Rpos = f.Buf
	ret = fc.iofuncs.Read(fc.cookie, (*int8)(unsafe.Pointer(f.Rpos)), f.Buf_size)
	if ret <= int64(0) {
		goto bail
	}
	f.Rend = (*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(f.Rpos)) + uintptr(ret)))
	*(*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(buf)) + uintptr(func() (_cgo_ret uint64) {
		_cgo_addr := &readlen
		_cgo_ret = *_cgo_addr
		*_cgo_addr++
		return
	}()))) = *func() (_cgo_ret *uint8) {
		_cgo_addr := &f.Rpos
		_cgo_ret = *_cgo_addr
		*(*uintptr)(unsafe.Pointer(_cgo_addr))++
		return
	}()
	return readlen
bail:
	f.Flags |= uint32(func() int32 {
		if ret == int64(0) {
			return int32(16)
		} else {
			return int32(32)
		}
	}())
	f.Rpos = func() (_cgo_ret *uint8) {
		_cgo_addr := &f.Rend
		*_cgo_addr = f.Buf
		return *_cgo_addr
	}()
	return readlen
}
func _cgos_cookiewrite_fopencookie(f *Struct__IO_FILE, buf *uint8, len uint64) uint64 {
	var fc *_cgos_fcookie_fopencookie = (*_cgos_fcookie_fopencookie)(f.Cookie)
	var ret int64
	var len2 uint64 = uint64(uintptr(unsafe.Pointer(f.Wpos)) - uintptr(unsafe.Pointer(f.Wbase)))
	if !(fc.iofuncs.Write != nil) {
		return len
	}
	if len2 != 0 {
		f.Wpos = f.Wbase
		if _cgos_cookiewrite_fopencookie(f, f.Wpos, len2) < len2 {
			return uint64(0)
		}
	}
	ret = fc.iofuncs.Write(fc.cookie, (*int8)(unsafe.Pointer(buf)), len)
	if ret < int64(0) {
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
		return uint64(0)
	}
	return uint64(ret)
}
func _cgos_cookieseek_fopencookie(f *Struct__IO_FILE, off int64, whence int32) int64 {
	var fc *_cgos_fcookie_fopencookie = (*_cgos_fcookie_fopencookie)(f.Cookie)
	var res int32
	if uint32(whence) > uint32(2) {
		*X__errno_location() = int32(22)
		return int64(-1)
	}
	if !(fc.iofuncs.Seek != nil) {
		*X__errno_location() = int32(95)
		return int64(-1)
	}
	res = fc.iofuncs.Seek(fc.cookie, &off, whence)
	if res < int32(0) {
		return int64(res)
	}
	return off
}
func _cgos_cookieclose_fopencookie(f *Struct__IO_FILE) int32 {
	var fc *_cgos_fcookie_fopencookie = (*_cgos_fcookie_fopencookie)(f.Cookie)
	if fc.iofuncs.Close != nil {
		return fc.iofuncs.Close(fc.cookie)
	}
	return int32(0)
}
func Fopencookie(cookie unsafe.Pointer, mode *int8, iofuncs Struct__IO_cookie_io_functions_t) *Struct__IO_FILE {
	var f *_cgos_cookie_FILE_fopencookie
	if !(Strchr((*int8)(unsafe.Pointer(&[4]int8{'r', 'w', 'a', '\x00'})), int32(*mode)) != nil) {
		*X__errno_location() = int32(22)
		return (*Struct__IO_FILE)(nil)
	}
	if !(func() (_cgo_ret *_cgos_cookie_FILE_fopencookie) {
		_cgo_addr := &f
		*_cgo_addr = (*_cgos_cookie_FILE_fopencookie)(Malloc(1304))
		return *_cgo_addr
	}() != nil) {
		return (*Struct__IO_FILE)(nil)
	}
	Memset(unsafe.Pointer(&f.f), int32(0), 232)
	if !(Strchr(mode, '+') != nil) {
		f.f.Flags = uint32(func() int32 {
			if int32(*mode) == 'r' {
				return int32(8)
			} else {
				return int32(4)
			}
		}())
	}
	f.fc.cookie = cookie
	f.fc.iofuncs = iofuncs
	f.f.Fd = -1
	f.f.Cookie = unsafe.Pointer(&f.fc)
	f.f.Buf = (*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer((*uint8)(unsafe.Pointer(&f.buf)))) + uintptr(int32(8))))
	f.f.Buf_size = 1024
	f.f.Lbf = -1
	f.f.Read = _cgos_cookieread_fopencookie
	f.f.Write = _cgos_cookiewrite_fopencookie
	f.f.Seek = _cgos_cookieseek_fopencookie
	f.f.Close = _cgos_cookieclose_fopencookie
	return __ofl_add(&f.f)
}
