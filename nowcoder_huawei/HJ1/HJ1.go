/*《字符串最后一个单词的长度》
题目描述
计算字符串最后一个单词的长度，单词以空格隔开。

输入描述:
输入一行，代表要计算的字符串，非空，长度小于5000。

输出描述:
输出一个整数，表示输入字符串最后一个单词的长度。
*/

package main

import (
	"fmt"
	"bufio"
	"os"
	"strings"
)

func method1(input string, err error) {
	var cnt int = 0
	if err == nil {
		for i := len(input)-2; i >= 0; i-- {
			if input[i] == ' ' && cnt != 0 {
				fmt.Println(cnt)
				return
			} else if input[i] == ' ' && cnt == 0 {
				continue
			} else {
				cnt++
			}
		}
	}
	fmt.Println(cnt)
}

func method2(input string, err error) {
	if err == nil {
		s := strings.Split(input[:len(input)-1], " ")
		last_ele := s[len(s)-1]
		fmt.Println(len(last_ele))
	}
}


func main() {
	inputReader := bufio.NewReader(os.Stdin)
	input, err := inputReader.ReadString('\n')
	method1(input, err)
	method2(input, err)

	// method3
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	if err2 := scanner.Err(); err2 != nil {
		fmt.Println(err2)
		return
	}
	line := scanner.Text()
	strs := strings.Fields(line)  // strings.Fields函数按空格进行分割字符串
	fmt.Println(len(strs[len(strs)-1]))

}