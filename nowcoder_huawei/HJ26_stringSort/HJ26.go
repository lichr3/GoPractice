/* HJ26_stringSort.字符串排序
思路：建立一个等长数组，往里面插数据即可。若是遇到非字母，就跳过.
伪代码：
	for i in [0:26]:	// 在字符串中寻找每一个字母a～z和A～Z,0表示a或A，25表示Z或Z
		for k in [0:len(str)]:	// 对于当前字母letter[i]，若出现在str中，则将其插入新数组res
			if str[k] == letter[i]:
				若当前所在位置curIdx不为字母，则不能插入，curIdx自增。
				否则res[curIdx] = str[k]，并且curIdx自增
*/

package main

import (
	"bufio"
	"fmt"
	"os"
	"unicode"
)

func main() {
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			input := scanner.Text()
			res := []byte(input)
			curIdx := 0
			for i := 0; i < 26; i++ {	// 26个字母
				for k := 0; k < len(input); k++ {  // 不可用若使用for a, b :=  range要小心，因为其中的变量b始终都是一个地址，不能用b来赋值
					if input[k] == uint8(i + 'a') || input[k] == uint8(i + 'A') {	// 判断是否为当前寻找的字母
						for !unicode.IsLetter(rune(input[curIdx])) && curIdx < len(input) {	// 判断是否当前位置是否可以放字母
							curIdx++
						}
						res[curIdx] = input[k]	// 放字母
						curIdx++				// 增idx
					}
				}
			}
			fmt.Println(string(res))
		}
}

//func method1() {  // 错误方法：因为交换位置会导致顺序不对
//	scanner := bufio.NewScanner(os.Stdin)
//	for scanner.Scan() {
//		input := scanner.Text()
//		line := []byte(input)
//		curIdx := 0					// 表示处理到哪个位置了
//		for i := 0; i < 26; i++ {  	// 表示26个字母依次遍历，比如现在处理Aa，就寻找str中的Aa，然后和str[currIdx]交换，就起到了排序的效果
//			for k := curIdx; k < len(input); k++ {
//				//fmt.Println(i, k, string(line),curIdx,string(line[curIdx]), string(line[k]))
//				// 什么情况跳过curIdx：如果目前位置是非字母，就保持原样不处理
//				if !unicode.IsLetter(rune(line[curIdx])) {
//					curIdx++
//				}
//				// 什么情况交换：找到字母，进行交换，从而排序
//				if line[k] == byte(i + 'a') || line[k] == byte(i + 'A') {
//					t := line[k]
//					line[k] = line[curIdx]
//					line[curIdx] = t
//					curIdx++
//				}
//			}
//		}
//		fmt.Println(string(line))
//	}
//}

