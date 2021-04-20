/* 《计算某字母出现次数》
题目描述
写出一个程序，接受一个由字母、数字和空格组成的字符串，和一个字母，然后输出输入字符串中该字母的出现次数。不区分大小写。

输入描述:
第一行输入一个由字母和数字以及空格组成的字符串，第二行输入一个字母。

输出描述:
输出输入字符串中含有该字符的个数。
*/

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// func count1(line, c string) {
// 	line = strings.ToUpper(line)
// 	c = strings.ToUpper(c)
// 	cnt := 0
// 	for _, v := range line {
// 		// fmt.Println(v, string(v), c)
// 		if v == rune(c[0]) {
// 			cnt++
// 		}
// 	}
// 	fmt.Println(cnt)
// }

func count2(line, c string) {
	line = strings.ToUpper(line)
	c = strings.ToUpper(c)
	hash := make(map[rune]int)
	for _, v := range line {
		hash[v] += 1
	}
	fmt.Println(hash[rune(c[0])])
}

func main() {
	input := bufio.NewReader(os.Stdin)
	line, err1 := input.ReadString('\n')
	c, err2 := input.ReadString('\n')
	if err1 != nil || err2 != nil{
		fmt.Println(err1)
		fmt.Println(err2)
		return
	}
	count2(line, c)
}