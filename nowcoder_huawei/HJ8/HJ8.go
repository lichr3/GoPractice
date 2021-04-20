package main

import (
	"fmt"
	"sort"
)


func main() {
	var size, key, value int
	m := make(map[int]int)
	to_sort := []int{} // 用于对索引排序
	fmt.Scanln(&size)
	for size > 0 {
		fmt.Scanln(&key, &value)
		m[key] += value
		size--
	}
	// 排序
	for k := range m {
		to_sort = append(to_sort, k)
	}
	sort.Slice(to_sort, func(i, j int) bool {
		return to_sort[i] < to_sort[j]
	})
	// 输出
	for _, i := range to_sort {
		fmt.Println(i, m[i])
	}
}