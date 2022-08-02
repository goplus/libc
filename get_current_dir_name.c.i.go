package libc

import unsafe "unsafe"

func get_current_dir_name() *int8 {
	var a struct_stat
	var b struct_stat
	var res *int8 = Getenv((*int8)(unsafe.Pointer(&[4]int8{'P', 'W', 'D', '\x00'})))
	if res != nil && int32(*res) != 0 && !(stat(res, &a) != 0) && !(stat((*int8)(unsafe.Pointer(&[2]int8{'.', '\x00'})), &b) != 0) && a.st_dev == b.st_dev && a.st_ino == b.st_ino {
		return Strdup(res)
	}
	return Getcwd(nil, uint64(0))
}
