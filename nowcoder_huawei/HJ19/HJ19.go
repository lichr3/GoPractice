/* HJ19.简单错误记录
考察：代码实现能力

思路：
	1.错误存储：既然文件名和错误行数都相同才算是相同错误，那么可以用 文件名+空格+行号 作为key，错误次数错位value，建立一个map errLogMap
	2.只输出最近8条错误记录：用数组allErrs按出现顺序存储错误记录，重复出现的不再二次记录。输出最后八次记录即可
算法：
	1.每次读取一行输入，以"/"作为分隔符号，取最后一个子串作为错误记录curErrLog
	2.将curErrlog放入errLogMap； 判断后将当前curErrlog放入allErrs和recentErrs；
	3.以recentErrs为数据源，参考allErrs排序，然后去errLogMap映射得到错误出现次数
注意事项：
	1.文件名至多保留16个字符
*/

package main

import (
	"bufio"
	"fmt"
	"os"
	"reflect"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	allErrs := make([]string, 0)		// 存储所有错误记录，已出现过的不再计入
	errLogMap := make(map[string]int)	// 存储每个错误出现的次数
	for scanner.Scan() {
		line := scanner.Text()          // 读取一整行
		//if line == "EOF" {break}
		err := strings.Split(line, "\\") // 以"\"分割成子串
		curErrLog := err[len(err)-1]    // 取最后一个子串作为errLog
		curErrLog = get16(curErrLog)	// 超过16个字符的文件名称，只记录文件的最后有效16个字符
		errLogMap[curErrLog] += 1       // 记录某一错误的出现次数到errLogMap
		if !isExistItem(curErrLog, allErrs) { allErrs = append(allErrs, curErrLog) } // 放入记录所有错误的allErrs
	}
	// 按照出现次序输出err
	for _, err := range allErrs[max(len(allErrs)-8, 0):] {
		fmt.Printf("%s %d\n", get16(err), errLogMap[err])
	}
}

// 判断slice中是否存在某个item
func isExistItem(value interface{}, array interface{}) bool {
	switch reflect.TypeOf(array).Kind() {
	case reflect.Slice:
		s := reflect.ValueOf(array)
		for i := 0; i < s.Len(); i++ {
			if reflect.DeepEqual(value, s.Index(i).Interface()) {
				return true
			}
		}
	}
	return false
}

func get16(errLog string) string {
	tmp := strings.Split(errLog, " ")
	filename, column := tmp[0], tmp[1]
	if len(filename) <= 16 {
		return errLog
	} else {
		return filename[len(filename)-16:] + " " + column
	}
}

func max(a, b int) int {
	if a > b { return a}
	return b
}