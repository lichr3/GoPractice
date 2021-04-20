/* HJ17.坐标移动
算法：
	1.使用bufio.scanner读取一整行数据，然后用用分号作为分隔符分开得到一个slice
	2.创建一个move函数，返回每次移动的结果：
		2.1 首先使用TrimSpace去掉首位空格
		2.2 然后判断是否有非法字符
*/

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input := scanner.Text()
	operations := strings.Split(input, ";")
	x, y := 0, 0
	for _, op := range operations {
		a, b := move(op)
		x += a
		y += b
	}
	fmt.Printf("%d,%d\n", x, y)
}

func move(m string) (int, int) {
	m = strings.TrimSpace(m)	// 去掉空字符
	if len(m) < 2 || len(m) > 3 {	// 合法坐标只能是2到3位
		return 0, 0
	}
	for i := 0; i < len(m); i++ {	// 判断坐标内容是否合法
		if (m[i] >= 'A' && m[i] <= 'Z') || (m[i] >= '0' && m[i] <= '9') {
			continue
		} else {
			return 0, 0
		}
	}
	if m[0] == 'A' {
		n, _ := strconv.Atoi(m[1:])
		return -n, 0
	}
	if m[0] == 'D' {
		n, _ := strconv.Atoi(m[1:])
		return n, 0
	}
	if m[0] == 'W' {
		n, _ := strconv.Atoi(m[1:])
		return 0, n
	} else {  // S
		n, _ := strconv.Atoi(m[1:])
		return 0, -n
	}

}