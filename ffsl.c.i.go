package libc

func ffsl(i int64) int32 {
	return func() int32 {
		if i != 0 {
			return a_ctz_l(uint64(i)) + int32(1)
		} else {
			return int32(0)
		}
	}()
}
