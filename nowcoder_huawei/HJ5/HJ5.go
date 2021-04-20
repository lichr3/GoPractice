/* HJ5.进制转换

*/

package main

import (
	"fmt"
	"io"
)

func main() {
	var num int
	var str string
	for {
		_, err := fmt.Scanln(&str)
		if err == io.EOF { break }
		fmt.Sscanf(str[2:] ,"%X", &num)
		fmt.Println(num)
	}
}