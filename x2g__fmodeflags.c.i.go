package libc

func __fmodeflags(mode *int8) int32 {
	var flags int32
	if Strchr(mode, '+') != nil {
		flags = int32(2)
	} else if int32(*mode) == 'r' {
		flags = int32(0)
	} else {
		flags = int32(1)
	}
	if Strchr(mode, 'x') != nil {
		flags |= int32(128)
	}
	if Strchr(mode, 'e') != nil {
		flags |= int32(524288)
	}
	if int32(*mode) != 'r' {
		flags |= int32(64)
	}
	if int32(*mode) == 'w' {
		flags |= int32(512)
	}
	if int32(*mode) == 'a' {
		flags |= int32(1024)
	}
	return flags
}
