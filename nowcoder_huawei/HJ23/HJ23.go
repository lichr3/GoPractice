/* HJ23.删除字符串中出现次数最少的字符
题目描述：
	1.实现删除字符串中出现次数最少的字符，若多个字符出现次数一样，则都删除。输出删除这些单词后的字符串，字符串中其它字符保持原来的顺序。
	2.注意每个输入文件有多组输入，即多个字符串用回车隔开.
输入描述: 字符串只包含小写英文字母, 不考虑非法输入，输入的字符串长度小于等于20个字节。
输出描述: 删除字符串中出现次数最少的字符后的字符串。
注意事项：
	1.若有多个字符出现的次数相同（都最少），则都删除
	2.字符串只包含小写英文字母，且长度小于等于20个字节（限制了输入就更好处理，例如可以用一个长26的切片来存储每个字母的出现次数）
算法：
	1.创建一个长为26的int型slice，存放每个小写字母的出现次数
	2.遍历后得到除0以外的最小值，凡事等于这个值的，就不放入新strin
*/

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		letters := make([]int, 26)
		// 对出现的每个小写字母进行计数
		for i := 0; i < len(line); i++ {
			letters[line[i]-'a'] += 1
		}
		// 获取出现次数最少的具体数目
		minNum := len(line)
		for i := 0; i < len(letters); i++ {
			if letters[i] > 0 {
				minNum = min(minNum, letters[i])
			}
		}
		// 删除字符串中出现次数最少的字符，把剩下的字符放进新的slice
		res := make([]byte, 0)
		for i := 0; i < len(line); i++ {
			if letters[line[i]-'a'] != minNum && letters[line[i]-'a'] != 0 {
				res = append(res, line[i])
			}
		}
		fmt.Println(string(res))
	}
}

func min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}