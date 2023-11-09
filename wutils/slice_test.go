package wutils

import (
	"fmt"
	"testing"
)

func TestWUtilSlice_StringIntersect(t *testing.T) {
	for i := 0; i < 100; i++ {
		go aa()
	}
}

func aa() {
	w := &WUtilSlice{}
	s1 := []string{"a", "b", "c"}
	s2 := []string{"a", "c"}
	intersect := w.StringsInter(s1, s2)
	fmt.Println(intersect)
}
