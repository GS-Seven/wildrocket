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
	intersect := w.Intersect(s1, s2)
	fmt.Println(intersect)
}

func TestWUtilSlice_DifferentSet(t *testing.T) {
	w := &WUtilSlice{}
	s1 := []string{"a", "b"}
	s2 := []string{"a", "b", "c"}
	intersect := w.DifferentSet(s1, s2)
	fmt.Println(intersect)
}

func TestWUtilSlice_Merge(t *testing.T) {
	w := &WUtilSlice{}
	s1 := []string{"a", "b"}
	s2 := []string{"a", "b", "c"}
	merge := w.Merge(s1, s2)
	fmt.Println(merge)
}

func TestWUtilSlice_MergeRepeatedElement(t *testing.T) {
	w := &WUtilSlice{}
	s1 := []string{"a", "b"}
	s2 := []string{"a", "b", "c"}
	merge := w.MergeRepeatedElement(s1, s2)
	fmt.Println(merge)
}

func TestWUtilSlice_RemoveRepeatedElement(t *testing.T) {
	//s1 := []string{"a", "b", "c", "a", "b", "1", "1"}
	//w := &WUtilSlice{}
	//arr := w.RemoveRepeatedElement(s1)
	//fmt.Println(arr)
	for i := 0; i < 1000; i++ {
		go bb()
	}
}

func bb() {
	s1 := []string{"a", "b", "c", "a", "b", "1", "1"}
	w := &WUtilSlice{}
	arr := w.RemoveRepeatedElement(s1)
	fmt.Println(arr)
}
