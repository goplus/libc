package libc

import unsafe "unsafe"

type fcookie_cgo18_fopencookie struct {
	cookie  unsafe.Pointer
	iofuncs struct__IO_cookie_io_functions_t
}
type cookie_FILE_cgo19_fopencookie struct {
	f   struct__IO_FILE
	fc  fcookie_cgo18_fopencookie
	buf [1032]uint8
}

func cookieread_cgo20_fopencookie(f *struct__IO_FILE, buf *uint8, len uint64) uint64 {
	var fc *fcookie_cgo18_fopencookie = (*fcookie_cgo18_fopencookie)(f.cookie)
	var ret int64 = int64(-1)
	var remain uint64 = len
	var readlen uint64 = uint64(0)
	var len2 uint64 = len - func() uint64 {
		if !!(f.buf_size != 0) {
			return 1
		} else {
			return 0
		}
	}()
	if !(fc.iofuncs.read != nil) {
		goto bail
	}
	if len2 != 0 {
		ret = fc.iofuncs.read(fc.cookie, (*int8)(unsafe.Pointer(buf)), len2)
		if ret <= int64(0) {
			goto bail
		}
		readlen += uint64(ret)
		remain -= uint64(ret)
	}
	if !(f.buf_size != 0) || remain > func() uint64 {
		if !!(f.buf_size != 0) {
			return 1
		} else {
			return 0
		}
	}() {
		return readlen
	}
	f.rpos = f.buf
	ret = fc.iofuncs.read(fc.cookie, (*int8)(unsafe.Pointer(f.rpos)), f.buf_size)
	if ret <= int64(0) {
		goto bail
	}
	f.rend = (*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(f.rpos)) + uintptr(ret)))
	*(*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(buf)) + uintptr(func() (_cgo_ret uint64) {
		_cgo_addr := &readlen
		_cgo_ret = *_cgo_addr
		*_cgo_addr++
		return
	}()))) = *func() (_cgo_ret *uint8) {
		_cgo_addr := &f.rpos
		_cgo_ret = *_cgo_addr
		*(*uintptr)(unsafe.Pointer(_cgo_addr))++
		return
	}()
	return readlen
bail:
	f.flags |= uint32(func() int32 {
		if ret == int64(0) {
			return int32(16)
		} else {
			return int32(32)
		}
	}())
	f.rpos = func() (_cgo_ret *uint8) {
		_cgo_addr := &f.rend
		*_cgo_addr = f.buf
		return *_cgo_addr
	}()
	return readlen
}
func cookiewrite_cgo21_fopencookie(f *struct__IO_FILE, buf *uint8, len uint64) uint64 {
	var fc *fcookie_cgo18_fopencookie = (*fcookie_cgo18_fopencookie)(f.cookie)
	var ret int64
	var len2 uint64 = uint64(uintptr(unsafe.Pointer(f.wpos)) - uintptr(unsafe.Pointer(f.wbase)))
	if !(fc.iofuncs.write != nil) {
		return len
	}
	if len2 != 0 {
		f.wpos = f.wbase
		if cookiewrite_cgo21_fopencookie(f, f.wpos, len2) < len2 {
			return uint64(0)
		}
	}
	ret = fc.iofuncs.write(fc.cookie, (*int8)(unsafe.Pointer(buf)), len)
	if ret < int64(0) {
		f.wpos = func() (_cgo_ret *uint8) {
			_cgo_addr := &f.wbase
			*_cgo_addr = func() (_cgo_ret *uint8) {
				_cgo_addr := &f.wend
				*_cgo_addr = (*uint8)(nil)
				return *_cgo_addr
			}()
			return *_cgo_addr
		}()
		f.flags |= uint32(32)
		return uint64(0)
	}
	return uint64(ret)
}
func cookieseek_cgo22_fopencookie(f *struct__IO_FILE, off int64, whence int32) int64 {
	var fc *fcookie_cgo18_fopencookie = (*fcookie_cgo18_fopencookie)(f.cookie)
	var res int32
	if uint32(whence) > uint32(2) {
		*__errno_location() = int32(22)
		return int64(-1)
	}
	if !(fc.iofuncs.seek != nil) {
		*__errno_location() = int32(95)
		return int64(-1)
	}
	res = fc.iofuncs.seek(fc.cookie, &off, whence)
	if res < int32(0) {
		return int64(res)
	}
	return off
}
func cookieclose_cgo23_fopencookie(f *struct__IO_FILE) int32 {
	var fc *fcookie_cgo18_fopencookie = (*fcookie_cgo18_fopencookie)(f.cookie)
	if fc.iofuncs.close != nil {
		return fc.iofuncs.close(fc.cookie)
	}
	return int32(0)
}
func Fopencookie(cookie unsafe.Pointer, mode *int8, iofuncs struct__IO_cookie_io_functions_t) *struct__IO_FILE {
	var f *cookie_FILE_cgo19_fopencookie
	if !(Strchr((*int8)(unsafe.Pointer(&[4]int8{'r', 'w', 'a', '\x00'})), int32(*mode)) != nil) {
		*__errno_location() = int32(22)
		return (*struct__IO_FILE)(nil)
	}
	if !(func() (_cgo_ret *cookie_FILE_cgo19_fopencookie) {
		_cgo_addr := &f
		*_cgo_addr = (*cookie_FILE_cgo19_fopencookie)(Malloc(1304))
		return *_cgo_addr
	}() != nil) {
		return (*struct__IO_FILE)(nil)
	}
	Memset(unsafe.Pointer(&f.f), int32(0), 232)
	if !(Strchr(mode, '+') != nil) {
		f.f.flags = uint32(func() int32 {
			if int32(*mode) == 'r' {
				return int32(8)
			} else {
				return int32(4)
			}
		}())
	}
	f.fc.cookie = cookie
	f.fc.iofuncs = iofuncs
	f.f.fd = -1
	f.f.cookie = unsafe.Pointer(&f.fc)
	f.f.buf = (*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer((*uint8)(unsafe.Pointer(&f.buf)))) + uintptr(int32(8))))
	f.f.buf_size = 1024
	f.f.lbf = -1
	f.f.read = cookieread_cgo20_fopencookie
	f.f.write = cookiewrite_cgo21_fopencookie
	f.f.seek = cookieseek_cgo22_fopencookie
	f.f.close = cookieclose_cgo23_fopencookie
	return __ofl_add(&f.f)
}
