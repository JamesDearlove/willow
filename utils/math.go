package utils

import "math"

func EaseOutQuart(x float64) float64 {
	return 1 - math.Pow(1-x, 4)
}

func MinInt(a int, b int) int {
	return int(math.Min(float64(a), float64(b)))
}

func MaxInt(a int, b int) int {
	return int(math.Max(float64(a), float64(b)))
}