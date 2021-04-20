package main

import (
	"fmt"
	"time"
)

func main() {
	m := make(map[int]int)
	s := []int{1,2,3,4,5}

	// 错误写法
	for i, v := range s{
		go func(idx, val int) {
			m[idx] = val
		}(i, v)
	}
	time.Sleep(2 * time.Second)
	fmt.Println(m)
}