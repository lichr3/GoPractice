/* HJ9.提取不重复的蒸熟
思路：逆序读取字符串，用map来去重，首次出现的数字可以加入
*/

package main

import (
	"fmt"
)


func main() {
	var input string
	m := make(map[byte]bool)
	r := []byte{}
	fmt.Scanln(&input)
	for i := len(input)-1; i >= 0; i-- {
		if !m[input[i]]{
			m[input[i]] = true
			r = append(r, input[i])
		}
	}
	fmt.Printf("%s\n", r)
}