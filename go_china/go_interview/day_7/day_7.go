package main

import "fmt"

const (
	a = iota   //0
	b          //1
	c          //2
	d = "ha"   //独立值，iota += 1
	e          //"ha"   iota += 1
	f = 100    //iota +=1
	g          //100  iota +=1
	h = iota   //7,恢复计数
	i          //8
)

const (
	a1 = 1 << iota  // a == 1  (iota == 0)
	b1 = 1 << iota  // b == 2  (iota == 1)
	c1 = 3          // c == 3  (iota == 2, unused)
	d1 = 1 << iota  // d == 8  (iota == 3)
)

func main() {
	fmt.Println("考点1:iota用于给常量计数，在下次一const关键字出现时被重置为0")
	fmt.Println(a,b,c,d,e,f,g,h,i)
	fmt.Println("iota可用于移位运算")
	fmt.Println(a1, b1, c1, d1)
	fmt.Println("\n 考点2:nil只能赋值给指针、chan、func、interface、map、slice")
}