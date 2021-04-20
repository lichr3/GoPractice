/* HJ20.密码验证合格程序
密码要求:
1.长度超过8位
2.包括大小写字母.数字.其它符号,以上四种至少三种
3.不能有相同长度大于2的子串重复（暴力求解长度为3的重复子串）

注意：
	1.写代码的时候记得看题，做错了两句，竟然是应为没把"密码长度超过8"这个条件加上去

问：长度为3的重复子串除了暴力有没有更好的解法？

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
		str := scanner.Text()
		if len(str) > 8 && getCharKinds(str) >= 3 && !isSubarrayRepeat(str) {
			fmt.Println("OK")
		} else {
			fmt.Println("NG")
		}
	}
}


// 计算密码中有几种符号类型。总共有四种类型，大写字母、小写字母、数字、特殊符号
func getCharKinds(pw string) int {
	var lower, upper, num, char int
	for i := 0; i < len(pw); i++ {
		if pw[i] >= 'a' && pw[i] <= 'z' {
			lower = 1
		} else if pw[i] >= 'A' && pw[i] <= 'Z' {
			upper = 1
		} else if pw[i] >= '0' && pw[i] <= '9' {
			num = 1
		} else {
			char = 1
		}
	}
	return lower + upper + num + char
}

// 判断是否有长度为3的重复子串
func isSubarrayRepeat(str string) bool {
	if len(str) < 6 { return true }
	for i := 0; i < len(str)-6; i++ {
		for k := 3+i; k < len(str)-3; k++ {
			if str[i:i+3] == str[k:k+3] { return true }
		}
	}
	return false
}