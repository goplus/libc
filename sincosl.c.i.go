package libc

func Sincosl(x float64, sin *float64, cos *float64) {
	var sind float64
	var cosd float64
	Sincos(float64(x), &sind, &cosd)
	*sin = float64(sind)
	*cos = float64(cosd)
}
