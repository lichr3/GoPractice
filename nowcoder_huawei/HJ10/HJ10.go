// HJ10.字符个数统计

package main

import "fmt"

func main() {
	var str string
	m := make(map[rune]bool)
	fmt.Scanln(&str)
	for _, v := range str {
		if int(v) >= 0 && int(v) <= 127 {
			m[v] = true
		}
	}
	fmt.Println(len(m))
}