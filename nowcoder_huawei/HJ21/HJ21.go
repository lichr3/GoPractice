/* HJ21. 简单密码破解 （考察代码能力）
映射关系： 1--1， abc--2, def--3, ghi--4, jkl--5, mno--6, pqrs--7, tuv--8 wxyz--9, 0--0
若是数字，则不表，若是大写字母，则变成小写字母往后诺一位，例如X->y, Z->a
例子：YUANzhi1987->zvbo9441987
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
		fmt.Println(translate(line))
	}
}

func translate(str string) string {
	res := make([]byte, len(str))
	for i := 0; i < len(str); i++ {
		// abc--2, def--3, ghi--4, jkl--5, mno--6, pqrs--7, tuv--8 wxyz--9, 0--0
		if str[i] >= '0' && str[i] <= '9' {
			res[i] = str[i]
		} else if str[i] >= 'A' && str[i] <= 'Y' {
			res[i] = 'b' + str[i] - 'A'
		} else if str[i] == 'Z' {
			res[i] = 'a'
		}else if str[i] >= 'a' && str[i] <= 'c' {
			res[i] = '2'
		} else if str[i] >= 'd' && str[i] <= 'f' {
			res[i] = '3'
		} else if str[i] >= 'g' && str[i] <= 'i' {
			res[i] = '4'
		} else if str[i] >= 'j' && str[i] <= 'l' {
			res[i] = '5'
		} else if str[i] >= 'm' && str[i] <= 'o' {
			res[i] = '6'
		} else if str[i] >= 'p' && str[i] <= 's' {
			res[i] = '7'
		} else if str[i] >= 't' && str[i] <= 'v' {
			res[i] = '8'
		} else if str[i] >= 'w' && str[i] <= 'z' {
			res[i] = '9'
		}  else {
			res[i] = str[i]
		}
	}
	return string(res)
}