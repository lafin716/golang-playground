package pmath

import (
	"fmt"
	"testing"
)

func TestLineAngle(t *testing.T) {
	a := [][]int{
		[]int{1, 4},
		[]int{9, 2},
		[]int{3, 8},
		[]int{11, 6},
	}

	result := LineAngle(a)
	if result == 1 {
		fmt.Println("정상")
	} else {
		fmt.Println("오류")
	}
}
