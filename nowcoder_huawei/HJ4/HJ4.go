/* HJ4.字符串分隔
算法：
	1.分行读取字符串
	2.判断字符串长度，小于8补0，等于8不处理，大于8进行分割。
	3.进行输出：判断是否为空字符串，若不是则输出
*/
package main

import (
	"fmt"
	"bufio"
	"os"
)

func my_print(str string) {
	length := len(str)
	if length == 0 {
		return
	} else if length < 8 {
		for i := length; i < 8; i++ {
			str += "0"
		}
	}
	fmt.Println(str)
}

func splitStr(str string) []string {
	res := []string{}
	for len(str) - 8 > 0 {
		res = append(res, str[:8])
		str = str[8:]
	}
	res = append(res, str)
	return res
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		tmp := scanner.Text()
		if len(tmp) <= 8 {
			my_print(tmp)
		} else {
			strs := splitStr(tmp)
			for _, s := range strs {
				my_print(s)
			}
		}
	}
}