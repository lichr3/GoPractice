/* HJ18_ipAddress.识别有效的IP地址和掩码并进行分类统计
思路：
	1.本质上是对IP地址进行分类，分为A,B,C,D,E,错误IP或错误掩码,私有IP。对于没种类别，写出相应函数进行判断，然后进行计数即可
	2.判断顺序：先判断是否错误IP或者掩码，错误则不再继续判断。然后判断是否私有IP，若是也要继续判断，随后对ABCDE符合一个则可不继续判断。
	3.注意事项：0.和127.开头的不属于任何一类
坑：
	1.判断掩码正确性的时候，分成四段分别转2进制时候，每段不会转成固定长度的八位。所以要考虑当某段对应的值不为0时，长度应该为8，否则这个掩码就是错的
	2.子网掩码 255.255.255.255 0.0.0.0 不合法
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
	// 每一行输入ip和掩码，用'～'符号隔开，分别用[]ip和[]mask存储
	ip := make([]string, 0)
	mask := make([]string, 0)
	res := [7]int{}	// 存储每个分类的个数，分别是A,B,C,D,E,错误IP或错误掩码,私有IP
	scanner := bufio.NewScanner(os.Stdin)
	// 每次读取一行数据
	for scanner.Scan() {
		input := scanner.Text()
		//if input == "EOF" {break}
		ip = append(ip, strings.Split(input, "~")[0])
		mask = append(mask, strings.Split(input, "~")[1])
	}
	// 数据读取完毕，判断每个ip和掩码分到哪一类
	for i := 0; i < len(ip); i++ {
		if !isIpRight(ip[i]) || !isMaskRight(mask[i]){ // class5：错误掩码或错误ip
			res[5] += 1; continue
		}
		if isPrivateIp(ip[i]) {	res[6] += 1	}
		if isClassA(ip[i]) { res[0] += 1; continue }
		if isClassB(ip[i]) { res[1] += 1; continue }
		if isClassC(ip[i]) { res[2] += 1; continue }
		if isClassD(ip[i]) { res[3] += 1; continue }
		if isClassE(ip[i]) { res[4] += 1}
	}
	// 输出结果
	for i, v := range res {
		if i == len(res)-1 {
			fmt.Println(v)
		} else {
			fmt.Printf("%d ", v)
		}
	}
}

func isMaskRight(mask string) bool {
	if mask == "255.255.255.255" || mask == "0.0.0.0" {return false}
	subMask := strings.Split(mask, ".")
	var binaryMask string
	// 将掩码转为二进制后进行判断即可
	for _, str := range subMask {
		if str == "" {return false }  // 为空返回false
		num, _ := strconv.Atoi(str)
		binaryNum := fmt.Sprintf("%b", num)
		if binaryNum != "0" && len(binaryNum) < 8 {return false} // 转换成2进制后不为0也不足8位，说明前几位为空
		binaryMask +=binaryNum
	}
	//fmt.Printf("binaryMask: %s\n", binaryMask)
	for i := 1; i < len(binaryMask); i++ {
		if binaryMask[i] == uint8('1') && binaryMask[i-1] == uint8('0') {
			return false
		}
	}
	return true
}

func isIpRight(ip string) bool {
	subIp := strings.Split(ip, ".")
	for _, str := range subIp {
		if str == "" { return false }
		num, _ := strconv.Atoi(str)
		if num < 0 || num > 255 {
			return false
		}
	}
	return true
}

// 查看ip是否在指定范围
func isInRange(ip, begin, end string) bool {
	subIp := strings.Split(ip, ".")
	subBeg := strings.Split(begin, ".")
	subEnd := strings.Split(end, ".")
	for i := 0; i < 4; i++ {
		numIp, _ := strconv.Atoi(subIp[i])
		numBeg, _ := strconv.Atoi(subBeg[i])
		numEnd, _ := strconv.Atoi(subEnd[i])
		if numIp < numBeg || numIp > numEnd	{
			return false
		}
	}
	return true
}

func isPrivateIp(ip string) bool {
/*	私网IP范围是：
	10.0.0.0～10.255.255.255
	172.16.0.0～172.31.255.255
	192.168.0.0～192.168.255.255*/
	if isInRange(ip, "10.0.0.0", "10.255.255.255") ||
		isInRange(ip, "172.16.0.0", "172.31.255.255") ||
		isInRange(ip, "192.168.0.0", "192.168.255.255") { return true }
	return false

}

func isClassA(ip string) bool {
	// A类地址1.0.0.0~126.255.255.255;
	if isInRange(ip, "1.0.0.0", "126.255.255.255") {
		//fmt.Println("A类：", ip)
		return true
	}
	return false
}

func isClassB(ip string) bool {
	// B类地址128.0.0.0~191.255.255.255;
	return isInRange(ip, "128.0.0.0", "191.255.255.255")
}

func isClassC(ip string) bool {
	// C类地址192.0.0.0~223.255.255.255;
	return isInRange(ip, "192.0.0.0", "223.255.255.255")
}

func isClassD(ip string) bool {
	// D类地址224.0.0.0~239.255.255.255；
	if isInRange(ip, "224.0.0.0", "239.255.255.255") {
		//fmt.Println("D类：", ip)
		return true
	}
	return false
}

func isClassE(ip string) bool {
	// E类地址240.0.0.0~255.255.255.255
	return isInRange(ip, "240.0.0.0", "255.255.255.255")
}