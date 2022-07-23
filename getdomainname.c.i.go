package libc

import unsafe "unsafe"

func getdomainname(name *int8, len uint64) int32 {
	var temp struct_utsname
	uname(&temp)
	if !(len != 0) || Strlen((*int8)(unsafe.Pointer(&temp.domainname))) >= len {
		*__errno_location() = int32(22)
		return -1
	}
	Strcpy(name, (*int8)(unsafe.Pointer(&temp.domainname)))
	return int32(0)
}
