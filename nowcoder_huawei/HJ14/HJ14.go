// HJ14.字符串排序
package main

import (
	"fmt"
	"sort"
)

func main() {
	var num int
	var word string
	words := []string{}
	fmt.Scanln(&num)
	for i := 0; i < num; i++ {
		fmt.Scanln(&word)
		words = append(words, word)
	}
	sort.Strings(words)
	for _, s := range words {
		fmt.Println(s)
	}
}