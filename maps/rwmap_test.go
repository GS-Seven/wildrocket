package maps

import (
	"fmt"
	"testing"
	"time"
)

// 测试并发 写 - 读 - 删
func TestConcurrencyRwMap(t *testing.T) {
	// 写
	rmap := NewConcurrencyRwMap(1000)
	for i := 0; i < 300; i++ {
		go func(i int) {
			rmap.Set(fmt.Sprintf("%v", i), i*2)
		}(i)
	}
	time.Sleep(time.Second * 4)
	fmt.Println(len(rmap.Map))
	fmt.Println(rmap.Map)

	// 读
	for i := 0; i < 300; i++ {
		go func(i int) {
			res := rmap.Get(fmt.Sprintf("%v", i))
			fmt.Printf("res:  key:%d => value: %d\n", i, res)
		}(i)
	}
	time.Sleep(time.Second * 4)

	// 删
	for i := 0; i < 30; i++ {
		go func(i int) {
			rmap.Delete(fmt.Sprintf("%v", i))
		}(i)
	}
}
