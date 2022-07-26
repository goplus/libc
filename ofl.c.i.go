package libc

import unsafe "unsafe"

var _cgos_ofl_head_ofl *Struct__IO_FILE
var _cgos_ofl_lock_ofl [1]int32
var __stdio_ofl_lockptr *int32 = (*int32)(unsafe.Pointer(&_cgos_ofl_lock_ofl))

func __ofl_lock() **Struct__IO_FILE {
	__lock((*int32)(unsafe.Pointer(&_cgos_ofl_lock_ofl)))
	return &_cgos_ofl_head_ofl
}
func __ofl_unlock() {
	__unlock((*int32)(unsafe.Pointer(&_cgos_ofl_lock_ofl)))
}
