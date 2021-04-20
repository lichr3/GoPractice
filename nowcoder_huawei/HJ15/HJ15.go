// HJ15.求int型正整数在内存中存储时1的个数
// 考点：10进制转2进制算法

package main

import "fmt"

func main() {
	var num int
	fmt.Scan(&num)
	DecToBin(num)
}

func DecToBin(num int) {
	var cnt = 0
	for num != 0 {
		if num%2 == 1 {
			cnt++
		}
		num /= 2
	}
	fmt.Println(cnt)
}
