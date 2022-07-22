package libc

func Nearbyintf(x float32) float32 {
	x = Rintf(x)
	return x
}
