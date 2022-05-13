package libc

func Strcpy(dest *int8, src *int8) *int8 {
	__stpcpy(dest, src)
	return dest
}
