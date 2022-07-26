package libc

import unsafe "unsafe"

var _cgos_dummy_file___stdio_exit *Struct__IO_FILE = nil

func _cgos_close_file___stdio_exit(f *Struct__IO_FILE) {
	if !(f != nil) {
		return
	}
	func() int32 {
		if f.Lock >= int32(0) {
			return __lockfile(f)
		} else {
			return int32(0)
		}
	}()
	if uintptr(unsafe.Pointer(f.Wpos)) != uintptr(unsafe.Pointer(f.Wbase)) {
		f.Write(f, nil, uint64(0))
	}
	if uintptr(unsafe.Pointer(f.Rpos)) != uintptr(unsafe.Pointer(f.Rend)) {
		f.Seek(f, int64(uintptr(unsafe.Pointer(f.Rpos))-uintptr(unsafe.Pointer(f.Rend))), int32(1))
	}
}
func __stdio_exit() {
	var f *Struct__IO_FILE
	for f = *__ofl_lock(); f != nil; f = f.Next {
		_cgos_close_file___stdio_exit(f)
	}
	_cgos_close_file___stdio_exit(__stdin_used)
	_cgos_close_file___stdio_exit(__stdout_used)
	_cgos_close_file___stdio_exit(__stderr_used)
}
