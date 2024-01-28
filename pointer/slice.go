package pointer

import (
	"strings"
)

// RemoveMatchedText 함수에서 사용할 매칭 변수
const (
	REMOVE_KEY = "pjw"
)

func CopySlice(s []string) *[]string {
	c := &[]string{}
	for _, t := range s {
		*c = append(*c, t)
	}

	return c
}

// slice 엘리먼트 제거
// 제대로 동작하지 않는다.
func RemoveMatchedText(s []string) {
	for i, text := range s {
		if strings.HasPrefix(text, REMOVE_KEY) {
			s = append(s[:i], s[i+1:]...)
		}
	}
}

// 포인터 기반 slice element 제거
func RemoveMatchedTextByPointer(s *[]string) {
	for i, text := range *s {
		if strings.HasPrefix(text, REMOVE_KEY) {
			*s = append((*s)[:i], (*s)[i+1:]...)
		}
	}
}
