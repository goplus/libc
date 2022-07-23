package libc

import unsafe "unsafe"

func _cgos_wstring_read_vswscanf(f *struct__IO_FILE, buf *uint8, len uint64) uint64 {
	var src *uint32 = (*uint32)(f.cookie)
	var k uint64
	if !(src != nil) {
		return uint64(0)
	}
	k = wcsrtombs((*int8)(unsafe.Pointer(f.buf)), &src, f.buf_size, nil)
	if k == uint64(18446744073709551615) {
		f.rpos = func() (_cgo_ret *uint8) {
			_cgo_addr := &f.rend
			*_cgo_addr = (*uint8)(nil)
			return *_cgo_addr
		}()
		return uint64(0)
	}
	f.rpos = f.buf
	f.rend = (*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(f.buf)) + uintptr(k)))
	f.cookie = unsafe.Pointer(src)
	if !(len != 0) || !(k != 0) {
		return uint64(0)
	}
	*buf = *func() (_cgo_ret *uint8) {
		_cgo_addr := &f.rpos
		_cgo_ret = *_cgo_addr
		*(*uintptr)(unsafe.Pointer(_cgo_addr))++
		return
	}()
	return uint64(1)
}
func vswscanf(s *uint32, fmt *uint32, ap []interface {
}) int32 {
	var buf [256]uint8
	var f struct__IO_FILE = struct__IO_FILE{0, nil, nil, nil, nil, nil, nil, nil, _cgos_wstring_read_vswscanf, nil, nil, (*uint8)(unsafe.Pointer(&buf)), 256, nil, nil, 0, 0, 0, 0, -1, 0, unsafe.Pointer(s), 0, nil, nil, nil, 0, 0, nil, nil, nil}
	return vfwscanf(&f, fmt, ap)
}
