package libc

import unsafe "unsafe"

func fwide(f *struct__IO_FILE, mode int32) int32 {
	var __need_unlock int32 = func() int32 {
		if f.lock >= int32(0) {
			return __lockfile(f)
		} else {
			return int32(0)
		}
	}()
	if mode != 0 {
		if !(f.locale != nil) {
			f.locale = func() *struct___locale_struct {
				if func() int32 {
					if !!(*(**struct___locale_map)(unsafe.Pointer(uintptr(unsafe.Pointer((**struct___locale_map)(unsafe.Pointer(&__pthread_self().locale.cat)))) + uintptr(int32(0))*8)) != nil) {
						return int32(4)
					} else {
						return int32(1)
					}
				}() == int32(1) {
					return (*struct___locale_struct)(unsafe.Pointer(&__c_locale))
				} else {
					return (*struct___locale_struct)(unsafe.Pointer(&__c_dot_utf8_locale))
				}
			}()
		}
		if !(f.mode != 0) {
			f.mode = func() int32 {
				if mode > int32(0) {
					return int32(1)
				} else {
					return -1
				}
			}()
		}
	}
	mode = f.mode
	for {
		if __need_unlock != 0 {
			__unlockfile(f)
		}
		if true {
			break
		}
	}
	return mode
}
