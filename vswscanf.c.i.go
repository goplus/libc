package libc

import unsafe "unsafe"

func _cgos_wstring_read_vswscanf(f *Struct__IO_FILE, buf *uint8, len uint64) uint64 {
	var src *uint32 = (*uint32)(f.Cookie)
	var k uint64
	if !(src != nil) {
		return uint64(0)
	}
	k = wcsrtombs((*int8)(unsafe.Pointer(f.Buf)), &src, f.Buf_size, nil)
	if k == uint64(18446744073709551615) {
		f.Rpos = func() (_cgo_ret *uint8) {
			_cgo_addr := &f.Rend
			*_cgo_addr = (*uint8)(nil)
			return *_cgo_addr
		}()
		return uint64(0)
	}
	f.Rpos = f.Buf
	f.Rend = (*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(f.Buf)) + uintptr(k)))
	f.Cookie = unsafe.Pointer(src)
	if !(len != 0) || !(k != 0) {
		return uint64(0)
	}
	*buf = *func() (_cgo_ret *uint8) {
		_cgo_addr := &f.Rpos
		_cgo_ret = *_cgo_addr
		*(*uintptr)(unsafe.Pointer(_cgo_addr))++
		return
	}()
	return uint64(1)
}
func vswscanf(s *uint32, fmt *uint32, ap []interface {
}) int32 {
	var buf [256]uint8
	var f Struct__IO_FILE = Struct__IO_FILE{0, nil, nil, nil, nil, nil, nil, nil, _cgos_wstring_read_vswscanf, nil, nil, (*uint8)(unsafe.Pointer(&buf)), 256, nil, nil, 0, 0, 0, 0, -1, 0, unsafe.Pointer(s), 0, nil, nil, nil, 0, 0, nil, nil, nil}
	return vfwscanf(&f, fmt, ap)
}
