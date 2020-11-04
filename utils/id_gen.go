package utils

import "math/rand"

func RangeIn(lo, hi int) int {
	return lo + rand.Intn(hi-lo)
}
