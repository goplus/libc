package libc

func getlogin_r(name *int8, size uint64) int32 {
	var logname *int8 = Getlogin()
	if !(logname != nil) {
		return int32(6)
	}
	if Strlen(logname) >= size {
		return int32(34)
	}
	Strcpy(name, logname)
	return int32(0)
}
