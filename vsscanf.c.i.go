package libc

import unsafe "unsafe"

func _cgos_string_read_vsscanf(f *Struct__IO_FILE, buf *uint8, len uint64) uint64 {
	var src *int8 = (*int8)(f.Cookie)
	var k uint64 = len + uint64(256)
	var end *int8 = (*int8)(Memchr(unsafe.Pointer(src), int32(0), k))
	if end != nil {
		k = uint64(uintptr(unsafe.Pointer(end)) - uintptr(unsafe.Pointer(src)))
	}
	if k < len {
		len = k
	}
	Memcpy(unsafe.Pointer(buf), unsafe.Pointer(src), len)
	f.Rpos = (*uint8)(unsafe.Pointer((*int8)(unsafe.Pointer(uintptr(unsafe.Pointer(src)) + uintptr(len)))))
	f.Rend = (*uint8)(unsafe.Pointer((*int8)(unsafe.Pointer(uintptr(unsafe.Pointer(src)) + uintptr(k)))))
	f.Cookie = unsafe.Pointer((*int8)(unsafe.Pointer(uintptr(unsafe.Pointer(src)) + uintptr(k))))
	return len
}
func Vsscanf(s *int8, fmt *int8, ap []interface {
}) int32 {
	var f Struct__IO_FILE = Struct__IO_FILE{0, nil, nil, nil, nil, nil, nil, nil, _cgos_string_read_vsscanf, nil, nil, (*uint8)(unsafe.Pointer(s)), 0, nil, nil, 0, 0, 0, 0, -1, 0, unsafe.Pointer(s), 0, nil, nil, nil, 0, 0, nil, nil, nil}
	return Vfscanf(&f, fmt, ap)
}
