package libc

func nice(inc int32) int32 {
	var prio int32 = inc
	if inc > -40 && inc < 40 {
		prio += getpriority(int32(0), uint32(0))
	}
	if prio > 19 {
		prio = 19
	}
	if prio < -20 {
		prio = -20
	}
	return func() int32 {
		if setpriority(int32(0), uint32(0), prio) != 0 {
			return -1
		} else {
			return prio
		}
	}()
}
