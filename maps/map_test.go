package maps

import (
	"fmt"
	"testing"
)

func TestMergeMap(t *testing.T) {
	dest := make(map[interface{}]interface{})
	dest["a"] = "fsdsdfasd"
	dest["b"] = 124

	src := make(map[interface{}]interface{})
	src["a"] = "fsdsdfasd"
	src["b"] = 12
	src["g"] = []int{1, 2, 3}
	src["f"] = []string{"aaa", "bb", "ccc"}

	mergeMap := MergeMap(dest, src)
	fmt.Println(mergeMap)
}
