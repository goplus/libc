package libc

func Secure_getenv(name *int8) *int8 {
	return func() *int8 {
		if int32(__libc.secure) != 0 {
			return nil
		} else {
			return Getenv(name)
		}
	}()
}
