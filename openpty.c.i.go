package libc

import unsafe "unsafe"

func openpty(pm *int32, ps *int32, name *int8, tio *struct_termios, ws *Struct_winsize) int32 {
	var m int32
	var s int32
	var n int32 = int32(0)
	var cs int32
	var buf [20]int8
	m = open((*int8)(unsafe.Pointer(&[10]int8{'/', 'd', 'e', 'v', '/', 'p', 't', 'm', 'x', '\x00'})), 258)
	if m < int32(0) {
		return -1
	}
	pthread_setcancelstate(int32(1), &cs)
	if Ioctl(m, int32(1074025521), &n) != 0 || Ioctl(m, int32(-2147199952), &n) != 0 {
		goto fail
	}
	if !(name != nil) {
		name = (*int8)(unsafe.Pointer(&buf))
	}
	Snprintf(name, 20, (*int8)(unsafe.Pointer(&[12]int8{'/', 'd', 'e', 'v', '/', 'p', 't', 's', '/', '%', 'd', '\x00'})), n)
	if func() (_cgo_ret int32) {
		_cgo_addr := &s
		*_cgo_addr = open(name, 258)
		return *_cgo_addr
	}() < int32(0) {
		goto fail
	}
	if tio != nil {
		tcsetattr(s, int32(0), tio)
	}
	if ws != nil {
		Ioctl(s, int32(21524), ws)
	}
	*pm = m
	*ps = s
	pthread_setcancelstate(cs, nil)
	return int32(0)
fail:
	Close(m)
	pthread_setcancelstate(cs, nil)
	return -1
}
