// HJ13.句子逆序
// 思路：用scanner获取一整行数据，然后用string.Fields拆分字符串为words（即以空格为中间符拆分）

package main

import (
	"fmt"
	"bufio"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	line := scanner.Text()
	words := strings.Fields(line)
	for i := len(words)-1; i >= 0; i-- {
		fmt.Printf("%s ", words[i])
	}
	fmt.Println()
}