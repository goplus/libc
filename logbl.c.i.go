package libc

func Logbl(x float64) float64 {
	if !(func() int32 {
		if 8 == 4 {
			return func() int32 {
				if X__FLOAT_BITS(float32(x))&uint32(2147483647) < uint32(2139095040) {
					return 1
				} else {
					return 0
				}
			}()
		} else {
			return func() int32 {
				if 8 == 8 {
					return func() int32 {
						if X__DOUBLE_BITS(float64(x))&9223372036854775807 < 9218868437227405312 {
							return 1
						} else {
							return 0
						}
					}()
				} else {
					return func() int32 {
						if X__fpclassifyl(x) > int32(1) {
							return 1
						} else {
							return 0
						}
					}()
				}
			}()
		}
	}() != 0) {
		return x * x
	}
	if x == float64(int32(0)) {
		return float64(-1) / (x * x)
	}
	return float64(Ilogbl(x))
}
