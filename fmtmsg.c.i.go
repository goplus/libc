package libc

import unsafe "unsafe"

func _cgos__strcolcmp_fmtmsg(lstr *int8, bstr *int8) int32 {
	var i uint64 = uint64(0)
	for int32(*(*int8)(unsafe.Pointer(uintptr(unsafe.Pointer(lstr)) + uintptr(i)))) != 0 && int32(*(*int8)(unsafe.Pointer(uintptr(unsafe.Pointer(bstr)) + uintptr(i)))) != 0 && int32(*(*int8)(unsafe.Pointer(uintptr(unsafe.Pointer(bstr)) + uintptr(i)))) == int32(*(*int8)(unsafe.Pointer(uintptr(unsafe.Pointer(lstr)) + uintptr(i)))) {
		i++
	}
	if int32(*(*int8)(unsafe.Pointer(uintptr(unsafe.Pointer(lstr)) + uintptr(i)))) != 0 || int32(*(*int8)(unsafe.Pointer(uintptr(unsafe.Pointer(bstr)) + uintptr(i)))) != 0 && int32(*(*int8)(unsafe.Pointer(uintptr(unsafe.Pointer(bstr)) + uintptr(i)))) != ':' {
		return int32(1)
	}
	return int32(0)
}
func fmtmsg(classification int64, label *int8, severity int32, text *int8, action *int8, tag *int8) int32 {
	var ret int32 = int32(0)
	var i int32
	var consolefd int32
	var verb int32 = int32(0)
	var errstring *int8 = nil
	var cmsg *int8 = Getenv((*int8)(unsafe.Pointer(&[8]int8{'M', 'S', 'G', 'V', 'E', 'R', 'B', '\x00'})))
	var msgs [6]*int8 = [6]*int8{(*int8)(unsafe.Pointer(&[6]int8{'l', 'a', 'b', 'e', 'l', '\x00'})), (*int8)(unsafe.Pointer(&[9]int8{'s', 'e', 'v', 'e', 'r', 'i', 't', 'y', '\x00'})), (*int8)(unsafe.Pointer(&[5]int8{'t', 'e', 'x', 't', '\x00'})), (*int8)(unsafe.Pointer(&[7]int8{'a', 'c', 't', 'i', 'o', 'n', '\x00'})), (*int8)(unsafe.Pointer(&[4]int8{'t', 'a', 'g', '\x00'})), nil}
	var cs int32
	pthread_setcancelstate(int32(1), &cs)
	if severity == int32(1) {
		errstring = (*int8)(unsafe.Pointer(&[7]int8{'H', 'A', 'L', 'T', ':', ' ', '\x00'}))
	} else if severity == int32(2) {
		errstring = (*int8)(unsafe.Pointer(&[8]int8{'E', 'R', 'R', 'O', 'R', ':', ' ', '\x00'}))
	} else if severity == int32(3) {
		errstring = (*int8)(unsafe.Pointer(&[10]int8{'W', 'A', 'R', 'N', 'I', 'N', 'G', ':', ' ', '\x00'}))
	} else if severity == int32(4) {
		errstring = (*int8)(unsafe.Pointer(&[7]int8{'I', 'N', 'F', 'O', ':', ' ', '\x00'}))
	}
	if classification&int64(512) != 0 {
		consolefd = open((*int8)(unsafe.Pointer(&[13]int8{'/', 'd', 'e', 'v', '/', 'c', 'o', 'n', 's', 'o', 'l', 'e', '\x00'})), int32(1))
		if consolefd < int32(0) {
			ret = int32(4)
		} else {
			if Dprintf(consolefd, (*int8)(unsafe.Pointer(&[18]int8{'%', 's', '%', 's', '%', 's', '%', 's', '%', 's', '%', 's', '%', 's', '%', 's', '\n', '\x00'})), func() *int8 {
				if label != nil {
					return label
				} else {
					return (*int8)(unsafe.Pointer((*int8)(unsafe.Pointer(&[1]int8{'\x00'}))))
				}
			}(), func() *int8 {
				if label != nil {
					return (*int8)(unsafe.Pointer(&[3]int8{':', ' ', '\x00'}))
				} else {
					return (*int8)(unsafe.Pointer(&[1]int8{'\x00'}))
				}
			}(), func() *int8 {
				if severity != 0 {
					return errstring
				} else {
					return (*int8)(unsafe.Pointer(&[1]int8{'\x00'}))
				}
			}(), func() *int8 {
				if text != nil {
					return text
				} else {
					return (*int8)(unsafe.Pointer((*int8)(unsafe.Pointer(&[1]int8{'\x00'}))))
				}
			}(), func() *int8 {
				if action != nil {
					return (*int8)(unsafe.Pointer(&[10]int8{'\n', 'T', 'O', ' ', 'F', 'I', 'X', ':', ' ', '\x00'}))
				} else {
					return (*int8)(unsafe.Pointer(&[1]int8{'\x00'}))
				}
			}(), func() *int8 {
				if action != nil {
					return action
				} else {
					return (*int8)(unsafe.Pointer((*int8)(unsafe.Pointer(&[1]int8{'\x00'}))))
				}
			}(), func() *int8 {
				if action != nil {
					return (*int8)(unsafe.Pointer(&[2]int8{' ', '\x00'}))
				} else {
					return (*int8)(unsafe.Pointer(&[1]int8{'\x00'}))
				}
			}(), func() *int8 {
				if tag != nil {
					return tag
				} else {
					return (*int8)(unsafe.Pointer((*int8)(unsafe.Pointer(&[1]int8{'\x00'}))))
				}
			}()) < int32(1) {
				ret = int32(4)
			}
			close(consolefd)
		}
	}
	if classification&int64(256) != 0 {
		for cmsg != nil && int32(*(*int8)(unsafe.Pointer(uintptr(unsafe.Pointer(cmsg)) + uintptr(int32(0))))) != 0 {
			for i = int32(0); *(**int8)(unsafe.Pointer(uintptr(unsafe.Pointer((**int8)(unsafe.Pointer(&msgs)))) + uintptr(i)*8)) != nil; i++ {
				if !(_cgos__strcolcmp_fmtmsg(*(**int8)(unsafe.Pointer(uintptr(unsafe.Pointer((**int8)(unsafe.Pointer(&msgs)))) + uintptr(i)*8)), cmsg) != 0) {
					break
				}
			}
			if uintptr(unsafe.Pointer(*(**int8)(unsafe.Pointer(uintptr(unsafe.Pointer((**int8)(unsafe.Pointer(&msgs)))) + uintptr(i)*8)))) == uintptr(unsafe.Pointer((*int8)(unsafe.Pointer(nil)))) {
				verb = int32(255)
				break
			} else {
				verb |= int32(1) << i
				cmsg = Strchr(cmsg, ':')
				if cmsg != nil {
					*(*uintptr)(unsafe.Pointer(&cmsg))++
				}
			}
		}
		if !(verb != 0) {
			verb = int32(255)
		}
		if Dprintf(int32(2), (*int8)(unsafe.Pointer(&[18]int8{'%', 's', '%', 's', '%', 's', '%', 's', '%', 's', '%', 's', '%', 's', '%', 's', '\n', '\x00'})), func() *int8 {
			if verb&int32(1) != 0 && label != nil {
				return label
			} else {
				return (*int8)(unsafe.Pointer((*int8)(unsafe.Pointer(&[1]int8{'\x00'}))))
			}
		}(), func() *int8 {
			if verb&int32(1) != 0 && label != nil {
				return (*int8)(unsafe.Pointer(&[3]int8{':', ' ', '\x00'}))
			} else {
				return (*int8)(unsafe.Pointer(&[1]int8{'\x00'}))
			}
		}(), func() *int8 {
			if verb&int32(2) != 0 && severity != 0 {
				return errstring
			} else {
				return (*int8)(unsafe.Pointer(&[1]int8{'\x00'}))
			}
		}(), func() *int8 {
			if verb&int32(4) != 0 && text != nil {
				return text
			} else {
				return (*int8)(unsafe.Pointer((*int8)(unsafe.Pointer(&[1]int8{'\x00'}))))
			}
		}(), func() *int8 {
			if verb&int32(8) != 0 && action != nil {
				return (*int8)(unsafe.Pointer(&[10]int8{'\n', 'T', 'O', ' ', 'F', 'I', 'X', ':', ' ', '\x00'}))
			} else {
				return (*int8)(unsafe.Pointer(&[1]int8{'\x00'}))
			}
		}(), func() *int8 {
			if verb&int32(8) != 0 && action != nil {
				return action
			} else {
				return (*int8)(unsafe.Pointer((*int8)(unsafe.Pointer(&[1]int8{'\x00'}))))
			}
		}(), func() *int8 {
			if verb&int32(8) != 0 && action != nil {
				return (*int8)(unsafe.Pointer(&[2]int8{' ', '\x00'}))
			} else {
				return (*int8)(unsafe.Pointer(&[1]int8{'\x00'}))
			}
		}(), func() *int8 {
			if verb&int32(16) != 0 && tag != nil {
				return tag
			} else {
				return (*int8)(unsafe.Pointer((*int8)(unsafe.Pointer(&[1]int8{'\x00'}))))
			}
		}()) < int32(1) {
			ret |= int32(1)
		}
	}
	if ret&5 == 5 {
		ret = -1
	}
	pthread_setcancelstate(cs, nil)
	return ret
}
