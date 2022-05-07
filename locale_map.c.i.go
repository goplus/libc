package libc

import unsafe "unsafe"

func __lctrans_impl(msg *int8, lm *struct___locale_map) *int8 {
	var trans *int8 = nil
	if lm != nil {
		trans = __mo_lookup(lm.map_, lm.map_size, msg)
	}
	return func() *int8 {
		if trans != nil {
			return trans
		} else {
			return msg
		}
	}()
}

var envvars_cgo574 [6][12]int8 = [6][12]int8{[12]int8{'L', 'C', '_', 'C', 'T', 'Y', 'P', 'E', '\x00'}, [12]int8{'L', 'C', '_', 'N', 'U', 'M', 'E', 'R', 'I', 'C', '\x00'}, [12]int8{'L', 'C', '_', 'T', 'I', 'M', 'E', '\x00'}, [12]int8{'L', 'C', '_', 'C', 'O', 'L', 'L', 'A', 'T', 'E', '\x00'}, [12]int8{'L', 'C', '_', 'M', 'O', 'N', 'E', 'T', 'A', 'R', 'Y', '\x00'}, [12]int8{'L', 'C', '_', 'M', 'E', 'S', 'S', 'A', 'G', 'E', 'S', '\x00'}}
var __locale_lock [1]int32
var __locale_lockptr *int32 = (*int32)(unsafe.Pointer(&__locale_lock))

func __get_locale(cat int32, val *int8) *struct___locale_map {
	var loc_head unsafe.Pointer
	_ = loc_head
	var p *struct___locale_map
	var new *struct___locale_map = nil
	var path *int8 = nil
	var z *int8
	var buf [256]int8
	var l uint64
	var n uint64
	if !(*val != 0) {
		if !(func() (_cgo_ret *int8) {
			_cgo_addr := &val
			*_cgo_addr = Getenv((*int8)(unsafe.Pointer(&[7]int8{'L', 'C', '_', 'A', 'L', 'L', '\x00'})))
			return *_cgo_addr
		}() != nil && int32(*val) != 0 || func() (_cgo_ret *int8) {
			_cgo_addr := &val
			*_cgo_addr = Getenv((*int8)(unsafe.Pointer(&*(*[12]int8)(unsafe.Pointer(uintptr(unsafe.Pointer((*[12]int8)(unsafe.Pointer(&envvars_cgo574)))) + uintptr(cat)*12)))))
			return *_cgo_addr
		}() != nil && int32(*val) != 0 || func() (_cgo_ret *int8) {
			_cgo_addr := &val
			*_cgo_addr = Getenv((*int8)(unsafe.Pointer(&[5]int8{'L', 'A', 'N', 'G', '\x00'})))
			return *_cgo_addr
		}() != nil && int32(*val) != 0) {
			func() (_cgo_ret *int8) {
				_cgo_addr := &val
				*_cgo_addr = (*int8)(unsafe.Pointer(&[8]int8{'C', '.', 'U', 'T', 'F', '-', '8', '\x00'}))
				return *_cgo_addr
			}()
		}
	}
	for n = uint64(0); n < uint64(23) && int32(*(*int8)(unsafe.Pointer(uintptr(unsafe.Pointer(val)) + uintptr(n)))) != 0 && int32(*(*int8)(unsafe.Pointer(uintptr(unsafe.Pointer(val)) + uintptr(n)))) != '/'; n++ {
	}
	if int32(*(*int8)(unsafe.Pointer(uintptr(unsafe.Pointer(val)) + uintptr(0)))) == '.' || int32(*(*int8)(unsafe.Pointer(uintptr(unsafe.Pointer(val)) + uintptr(n)))) != 0 {
		val = (*int8)(unsafe.Pointer(&[8]int8{'C', '.', 'U', 'T', 'F', '-', '8', '\x00'}))
	}
	var builtin int32 = func() int32 {
		if int32(*(*int8)(unsafe.Pointer(uintptr(unsafe.Pointer(val)) + uintptr(0)))) == 'C' && !(*(*int8)(unsafe.Pointer(uintptr(unsafe.Pointer(val)) + uintptr(1))) != 0) || !(Strcmp(val, (*int8)(unsafe.Pointer(&[8]int8{'C', '.', 'U', 'T', 'F', '-', '8', '\x00'}))) != 0) || !(Strcmp(val, (*int8)(unsafe.Pointer(&[6]int8{'P', 'O', 'S', 'I', 'X', '\x00'}))) != 0) {
			return 1
		} else {
			return 0
		}
	}()
	if builtin != 0 {
		if cat == 0 && int32(*(*int8)(unsafe.Pointer(uintptr(unsafe.Pointer(val)) + uintptr(1)))) == '.' {
			return (*struct___locale_map)(unsafe.Pointer(&__c_dot_utf8))
		}
		return (*struct___locale_map)(nil)
	}
	for p = (*struct___locale_map)(loc_head); p != nil; p = p.next {
		if !(Strcmp(val, (*int8)(unsafe.Pointer(&p.name))) != 0) {
			return p
		}
	}
	if !(__libc.secure != 0) {
		path = Getenv((*int8)(unsafe.Pointer(&[13]int8{'M', 'U', 'S', 'L', '_', 'L', 'O', 'C', 'P', 'A', 'T', 'H', '\x00'})))
	}
	if path != nil {
		for ; *path != 0; path = (*int8)(unsafe.Pointer(uintptr(unsafe.Pointer(z)) + func() uintptr {
			if !!(*z != 0) {
				return 1
			} else {
				return 0
			}
		}())) {
			z = __strchrnul(path, ':')
			l = uint64(uintptr(unsafe.Pointer(z)) - uintptr(unsafe.Pointer(path)))
			if l >= 256-n-uint64(2) {
				continue
			}
			Memcpy(unsafe.Pointer((*int8)(unsafe.Pointer(&buf))), unsafe.Pointer(path), l)
			*(*int8)(unsafe.Pointer(uintptr(unsafe.Pointer((*int8)(unsafe.Pointer(&buf)))) + uintptr(l))) = int8('/')
			Memcpy(unsafe.Pointer((*int8)(unsafe.Pointer(uintptr(unsafe.Pointer((*int8)(unsafe.Pointer(uintptr(unsafe.Pointer((*int8)(unsafe.Pointer(&buf))))+uintptr(l)))))+uintptr(1)))), unsafe.Pointer(val), n)
			*(*int8)(unsafe.Pointer(uintptr(unsafe.Pointer((*int8)(unsafe.Pointer(&buf)))) + uintptr(l+uint64(1)+n))) = int8(0)
			var map_size uint64
			var map_ unsafe.Pointer = unsafe.Pointer(__map_file((*int8)(unsafe.Pointer(&buf)), &map_size))
			if map_ != nil {
				new = (*struct___locale_map)(__libc_malloc(48))
				if !(new != nil) {
					__munmap(unsafe.Pointer(map_), map_size)
					break
				}
				new.map_ = map_
				new.map_size = map_size
				Memcpy(unsafe.Pointer((*int8)(unsafe.Pointer(&new.name))), unsafe.Pointer(val), n)
				*(*int8)(unsafe.Pointer(uintptr(unsafe.Pointer((*int8)(unsafe.Pointer(&new.name)))) + uintptr(n))) = int8(0)
				new.next = (*struct___locale_map)(loc_head)
				loc_head = unsafe.Pointer(new)
				break
			}
		}
	}
	if !(new != nil) && func() (_cgo_ret *struct___locale_map) {
		_cgo_addr := &new
		*_cgo_addr = (*struct___locale_map)(__libc_malloc(48))
		return *_cgo_addr
	}() != nil {
		new.map_ = __c_dot_utf8.map_
		new.map_size = __c_dot_utf8.map_size
		Memcpy(unsafe.Pointer((*int8)(unsafe.Pointer(&new.name))), unsafe.Pointer(val), n)
		*(*int8)(unsafe.Pointer(uintptr(unsafe.Pointer((*int8)(unsafe.Pointer(&new.name)))) + uintptr(n))) = int8(0)
		new.next = (*struct___locale_map)(loc_head)
		loc_head = unsafe.Pointer(new)
	}
	if !(new != nil) && cat == 0 {
		new = (*struct___locale_map)(unsafe.Pointer(&__c_dot_utf8))
	}
	return new
}
