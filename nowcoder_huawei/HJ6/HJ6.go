/* HJ6.质数因子
算法：
	1.将要分解的数放入一个列表，不断进行因数分解（从2到N/2）
	2.分解成功后，继续对后面的数进行分解
180
2 90
2 2 45
2 2 3 15
2 2 3 3 5

*/

package main

import (
	"fmt"
	"math"
)

func help(num int) []int {
	res := []int{}
	for {
		flag := true
		for i := 2; float64(i) <= math.Sqrt(float64(num)); i++ {
			if num % i == 0 {
				res = append(res, i)
				num = num / i
				flag = false
				break
			}
		}
		if flag { break }
	}
	res = append(res, num)
	return res
}

func my_print(res []int) {
	for _, v := range res {
		fmt.Printf("%d ", v)
	}
	fmt.Println()
}

func main() {
	var num int
	var res []int
	fmt.Scanln(&num)
	res = help(num)
	my_print(res)
}