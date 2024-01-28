package pmath

import "math"

func LineAngle(a [][]int) int {
	if math.Abs(float64(a[0][0]-a[1][0]))-math.Abs(float64(a[2][0]-a[3][0])) == 0 &&
		math.Abs(float64(a[0][1]-a[1][1]))-math.Abs(float64(a[2][1]-a[3][1])) == 0 {
		return 1
	}

	if math.Abs(float64(a[0][0]-a[2][0]))-math.Abs(float64(a[1][0]-a[3][0])) == 0 &&
		math.Abs(float64(a[0][1]-a[2][1]))-math.Abs(float64(a[1][1]-a[3][1])) == 0 {
		return 1
	}

	if math.Abs(float64(a[0][0]-a[3][0]))-math.Abs(float64(a[1][0]-a[2][0])) == 0 &&
		math.Abs(float64(a[0][1]-a[3][1]))-math.Abs(float64(a[1][1]-a[2][1])) == 0 {
		return 1
	}

	return 0
}
