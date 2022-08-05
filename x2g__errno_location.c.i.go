package libc

func X__errno_location() *int32 {
	return &__pthread_self().Errno_val
}
