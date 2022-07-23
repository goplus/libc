package libc

import unsafe "unsafe"

var _cgos_init_random [32]uint32 = [32]uint32{uint32(0), uint32(1481765933), uint32(3232861391), uint32(3417699910), uint32(3338875177), uint32(812669700), uint32(553475508), uint32(2592833400), uint32(1344887256), uint32(2877900904), uint32(1812158119), uint32(2295183359), uint32(3027751999), uint32(1889772843), uint32(2833562353), uint32(4253237756), uint32(2330030041), uint32(1949118330), uint32(220137366), uint32(1979932169), uint32(1089957932), uint32(1873226917), uint32(2863153495), uint32(1486937972), uint32(3343516516), uint32(2924690628), uint32(68706223), uint32(1843638549), uint32(212567592), uint32(4030971812), uint32(964776169), uint32(928126551)}
var _cgos_n_random int32 = int32(31)
var _cgos_i_random int32 = int32(3)
var _cgos_j_random int32 = int32(0)
var _cgos_x_random *uint32 = (*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer((*uint32)(unsafe.Pointer(&_cgos_init_random)))) + uintptr(int32(1))*4))
var _cgos_lock_random [1]int32
var __random_lockptr *int32 = (*int32)(unsafe.Pointer(&_cgos_lock_random))

func _cgos_lcg31_random(x uint32) uint32 {
	return (uint32(1103515245)*x + uint32(12345)) & uint32(2147483647)
}
func _cgos_lcg64_random(x uint64) uint64 {
	return uint64(6364136223846793005)*x + uint64(1)
}
func _cgos_savestate_random() unsafe.Pointer {
	*(*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer(_cgos_x_random)) - uintptr(1)*4)) = uint32(_cgos_n_random<<int32(16) | _cgos_i_random<<int32(8) | _cgos_j_random)
	return unsafe.Pointer((*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer(_cgos_x_random)) - uintptr(int32(1))*4)))
}
func _cgos_loadstate_random(state *uint32) {
	_cgos_x_random = (*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer(state)) + uintptr(int32(1))*4))
	_cgos_n_random = int32(*(*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer(_cgos_x_random)) - uintptr(1)*4)) >> int32(16))
	_cgos_i_random = int32(*(*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer(_cgos_x_random)) - uintptr(1)*4)) >> int32(8) & uint32(255))
	_cgos_j_random = int32(*(*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer(_cgos_x_random)) - uintptr(1)*4)) & uint32(255))
}
func _cgos___srandom_random(seed uint32) {
	var k int32
	var s uint64 = uint64(seed)
	if _cgos_n_random == int32(0) {
		*(*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer(_cgos_x_random)) + uintptr(int32(0))*4)) = uint32(s)
		return
	}
	_cgos_i_random = func() int32 {
		if _cgos_n_random == int32(31) || _cgos_n_random == int32(7) {
			return int32(3)
		} else {
			return int32(1)
		}
	}()
	_cgos_j_random = int32(0)
	for k = int32(0); k < _cgos_n_random; k++ {
		s = _cgos_lcg64_random(s)
		*(*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer(_cgos_x_random)) + uintptr(k)*4)) = uint32(s >> int32(32))
	}
	*(*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer(_cgos_x_random)) + uintptr(int32(0))*4)) |= uint32(1)
}
func Srandom(seed uint32) {
	__lock((*int32)(unsafe.Pointer(&_cgos_lock_random)))
	_cgos___srandom_random(seed)
	__unlock((*int32)(unsafe.Pointer(&_cgos_lock_random)))
}
func Initstate(seed uint32, state *int8, size uint64) *int8 {
	var old unsafe.Pointer
	if size < uint64(8) {
		return (*int8)(nil)
	}
	__lock((*int32)(unsafe.Pointer(&_cgos_lock_random)))
	old = _cgos_savestate_random()
	if size < uint64(32) {
		_cgos_n_random = int32(0)
	} else if size < uint64(64) {
		_cgos_n_random = int32(7)
	} else if size < uint64(128) {
		_cgos_n_random = int32(15)
	} else if size < uint64(256) {
		_cgos_n_random = int32(31)
	} else {
		_cgos_n_random = int32(63)
	}
	_cgos_x_random = (*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer((*uint32)(unsafe.Pointer(state)))) + uintptr(int32(1))*4))
	_cgos___srandom_random(seed)
	_cgos_savestate_random()
	__unlock((*int32)(unsafe.Pointer(&_cgos_lock_random)))
	return (*int8)(old)
}
func Setstate(state *int8) *int8 {
	var old unsafe.Pointer
	__lock((*int32)(unsafe.Pointer(&_cgos_lock_random)))
	old = _cgos_savestate_random()
	_cgos_loadstate_random((*uint32)(unsafe.Pointer(state)))
	__unlock((*int32)(unsafe.Pointer(&_cgos_lock_random)))
	return (*int8)(old)
}
func Random() int64 {
	var k int64
	__lock((*int32)(unsafe.Pointer(&_cgos_lock_random)))
	if _cgos_n_random == int32(0) {
		k = int64(func() (_cgo_ret uint32) {
			_cgo_addr := &*(*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer(_cgos_x_random)) + uintptr(int32(0))*4))
			*_cgo_addr = _cgos_lcg31_random(*(*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer(_cgos_x_random)) + uintptr(int32(0))*4)))
			return *_cgo_addr
		}())
		goto end
	}
	*(*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer(_cgos_x_random)) + uintptr(_cgos_i_random)*4)) += *(*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer(_cgos_x_random)) + uintptr(_cgos_j_random)*4))
	k = int64(*(*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer(_cgos_x_random)) + uintptr(_cgos_i_random)*4)) >> int32(1))
	if func() (_cgo_ret int32) {
		_cgo_addr := &_cgos_i_random
		*_cgo_addr++
		return *_cgo_addr
	}() == _cgos_n_random {
		_cgos_i_random = int32(0)
	}
	if func() (_cgo_ret int32) {
		_cgo_addr := &_cgos_j_random
		*_cgo_addr++
		return *_cgo_addr
	}() == _cgos_n_random {
		_cgos_j_random = int32(0)
	}
end:
	__unlock((*int32)(unsafe.Pointer(&_cgos_lock_random)))
	return k
}
