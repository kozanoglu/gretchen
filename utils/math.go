package utils

import "math"

func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func PercentageDiff(a, b float64) float64 {
	return math.Round((a-b)/b*100*100) / 100
}
