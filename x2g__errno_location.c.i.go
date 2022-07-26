package libc

func __errno_location() *int32 {
	return &__pthread_self().Errno_val
}
