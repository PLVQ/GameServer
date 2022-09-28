package comm

import "math"

func GetDiffDay(tTime1, tTime2 int64) uint {
	return uint(math.Abs(float64(tTime1-tTime2))/86400 + 1)
}
