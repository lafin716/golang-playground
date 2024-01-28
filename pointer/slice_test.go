package pointer

import (
	"fmt"
	"testing"
)

func TestRemoveSlice(t *testing.T) {
	testSlice := []string{
		"11111",
		"pjwtest",
		"33333333",
		"444444444",
	}

	fmt.Printf("before: %+v \n", testSlice)
	RemoveMatchedText(testSlice)
	fmt.Printf("after: %+v \n", testSlice)
}

func TestRemoveSliceByPointer(t *testing.T) {
	testSlice := &[]string{
		"11111",
		"pjwtest",
		"33333333",
		"444444444",
	}

	fmt.Printf("before: %+v \n", testSlice)
	RemoveMatchedTextByPointer(testSlice)
	fmt.Printf("after: %+v \n", testSlice)
}
