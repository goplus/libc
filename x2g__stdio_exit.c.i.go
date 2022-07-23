package libc

import unsafe "unsafe"

var _cgos_dummy_file____stdio_exit *struct__IO_FILE = nil

func _cgos_close_file____stdio_exit(f *struct__IO_FILE) {
	if !(f != nil) {
		return
	}
	func() int32 {
		if f.lock >= int32(0) {
			return __lockfile(f)
		} else {
			return int32(0)
		}
	}()
	if uintptr(unsafe.Pointer(f.wpos)) != uintptr(unsafe.Pointer(f.wbase)) {
		f.write(f, nil, uint64(0))
	}
	if uintptr(unsafe.Pointer(f.rpos)) != uintptr(unsafe.Pointer(f.rend)) {
		f.seek(f, int64(uintptr(unsafe.Pointer(f.rpos))-uintptr(unsafe.Pointer(f.rend))), int32(1))
	}
}
func __stdio_exit() {
	var f *struct__IO_FILE
	for f = *__ofl_lock(); f != nil; f = f.next {
		_cgos_close_file____stdio_exit(f)
	}
	_cgos_close_file____stdio_exit(__stdin_used)
	_cgos_close_file____stdio_exit(__stdout_used)
	_cgos_close_file____stdio_exit(__stderr_used)
}
