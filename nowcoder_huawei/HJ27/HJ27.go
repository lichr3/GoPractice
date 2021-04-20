/* HJ27.查找兄弟单词
思路：
	1.需要一个判断是否为兄弟单词的函数，把是兄弟单词的放入一个切片
	2.对切片中的字符串进行排序，然后选择第k个单词即可
如何判断是否为兄弟单词：
	方式1：直接排序后看两者是否完全相等
	方式2：使用map，但是也很麻烦
*/

package main

import (
	"fmt"
	"sort"
)

func main() {
	var n, k int
	words := make([]string, 0)		// 存放输入的n个单词
	var target string 				// 最后输入的目标单词
	brothers := make([]string, 0) 	// 存放n个单词中目标单词的兄弟单词
	// 读取输入
	fmt.Scan(&n)
	for i := 0; i < n; i++ {
		var word string
		fmt.Scan(&word)
		words = append(words, word)
	}
	fmt.Scan(&target, &k)

	// 判读哪些单词是兄弟单词
	for _, word := range words {
		if isBrother(word, target) {
			brothers = append(brothers, word)
		}
	}

	// 对兄弟单词进行排序
	sort.Strings(brothers)
	// 输出查找到的兄弟单词个数
	fmt.Println(len(brothers))
	// 取排序后第k个兄弟单词，没有则不取
	if len(brothers) >= k {
		fmt.Println(brothers[k-1])
	}
}

func isBrother(strA, strB string) bool {
	// 若两个字符串完全相同或者长度不同，则不是兄弟单词
	if strA == strB || len(strA) != len(strB) {
		return false
	}
	// 若排序后两者完全相同，则是兄弟单词
	tmpA := []byte(strA)
	tmpB := []byte(strB)
	sort.Slice(tmpA, func(i, j int) bool {
		return tmpA[i] < tmpA[j]
	})
	sort.Slice(tmpB, func(i, j int) bool {
		return tmpB[i] < tmpB[j]
	})
	return string(tmpA) == string(tmpB)
}