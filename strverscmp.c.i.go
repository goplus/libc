package libc

import unsafe "unsafe"

func strverscmp(l0 *int8, r0 *int8) int32 {
	var l *uint8 = (*uint8)(unsafe.Pointer(l0))
	var r *uint8 = (*uint8)(unsafe.Pointer(r0))
	var i uint64
	var dp uint64
	var j uint64
	var z int32 = int32(1)
	for dp = func() (_cgo_ret uint64) {
		_cgo_addr := &i
		*_cgo_addr = uint64(0)
		return *_cgo_addr
	}(); int32(*(*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(l)) + uintptr(i)))) == int32(*(*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(r)) + uintptr(i)))); i++ {
		var c int32 = int32(*(*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(l)) + uintptr(i))))
		if !(c != 0) {
			return int32(0)
		}
		if !(func() int32 {
			if int32(0) != 0 {
				return Isdigit(c)
			} else {
				return func() int32 {
					if uint32(c)-uint32('0') < uint32(10) {
						return 1
					} else {
						return 0
					}
				}()
			}
		}() != 0) {
			func() int32 {
				dp = i + uint64(1)
				return func() (_cgo_ret int32) {
					_cgo_addr := &z
					*_cgo_addr = int32(1)
					return *_cgo_addr
				}()
			}()
		} else if c != '0' {
			z = int32(0)
		}
	}
	if int32(*(*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(l)) + uintptr(dp)))) != '0' && int32(*(*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(r)) + uintptr(dp)))) != '0' {
		for j = i; func() int32 {
			if int32(0) != 0 {
				return Isdigit(int32(*(*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(l)) + uintptr(j)))))
			} else {
				return func() int32 {
					if uint32(*(*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(l)) + uintptr(j))))-uint32('0') < uint32(10) {
						return 1
					} else {
						return 0
					}
				}()
			}
		}() != 0; j++ {
			if !(func() int32 {
				if int32(0) != 0 {
					return Isdigit(int32(*(*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(r)) + uintptr(j)))))
				} else {
					return func() int32 {
						if uint32(*(*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(r)) + uintptr(j))))-uint32('0') < uint32(10) {
							return 1
						} else {
							return 0
						}
					}()
				}
			}() != 0) {
				return int32(1)
			}
		}
		if func() int32 {
			if int32(0) != 0 {
				return Isdigit(int32(*(*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(r)) + uintptr(j)))))
			} else {
				return func() int32 {
					if uint32(*(*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(r)) + uintptr(j))))-uint32('0') < uint32(10) {
						return 1
					} else {
						return 0
					}
				}()
			}
		}() != 0 {
			return -1
		}
	} else if z != 0 && dp < i && (func() int32 {
		if int32(0) != 0 {
			return Isdigit(int32(*(*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(l)) + uintptr(i)))))
		} else {
			return func() int32 {
				if uint32(*(*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(l)) + uintptr(i))))-uint32('0') < uint32(10) {
					return 1
				} else {
					return 0
				}
			}()
		}
	}() != 0 || func() int32 {
		if int32(0) != 0 {
			return Isdigit(int32(*(*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(r)) + uintptr(i)))))
		} else {
			return func() int32 {
				if uint32(*(*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(r)) + uintptr(i))))-uint32('0') < uint32(10) {
					return 1
				} else {
					return 0
				}
			}()
		}
	}() != 0) {
		return int32(uint8(int32(*(*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(l)) + uintptr(i))))-'0')) - int32(uint8(int32(*(*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(r)) + uintptr(i))))-'0'))
	}
	return int32(*(*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(l)) + uintptr(i)))) - int32(*(*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(r)) + uintptr(i))))
}
