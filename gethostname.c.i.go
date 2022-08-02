package libc

import unsafe "unsafe"

func gethostname(name *int8, len uint64) int32 {
	var i uint64
	var uts struct_utsname
	if uname(&uts) != 0 {
		return -1
	}
	if len > 65 {
		len = uint64(65)
	}
	for i = uint64(0); i < len && int32(func() (_cgo_ret int8) {
		_cgo_addr := &*(*int8)(unsafe.Pointer(uintptr(unsafe.Pointer(name)) + uintptr(i)))
		*_cgo_addr = *(*int8)(unsafe.Pointer(uintptr(unsafe.Pointer((*int8)(unsafe.Pointer(&uts.nodename)))) + uintptr(i)))
		return *_cgo_addr
	}()) != 0; i++ {
	}
	if i != 0 && i == len {
		*(*int8)(unsafe.Pointer(uintptr(unsafe.Pointer(name)) + uintptr(i-uint64(1)))) = int8(0)
	}
	return int32(0)
}
