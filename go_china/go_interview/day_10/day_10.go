package main

import "fmt"

func test1() {
	fmt.Println("test1考点：不同类型的数值不能直接相加，需要显示转换类型")
	a := 5
	b := 8.6
	// fmt.Println(a + b)  // invalid operation: a + b (mismatched types int and float64)
	fmt.Println(a + int(b))
}

func test2() {
	fmt.Println("test2考点：切片的使用。使用数组获取切片时如何定义len和cap")
	a := [5]int{1, 2, 3, 4, 5}
	t := a[3:4:4]
	fmt.Println("t[0] =", t[0], "len(t) =", len(t), ", cap(t) = ", cap(t))
}

func test3() {
	fmt.Println("test3考点：1.数组的缺省值以0表示;\n 2.数组的长度也是数组类型的组成部分")
	a := [2]int{5, 6}
	b := [3]int{5, 6}
	fmt.Println("a =", a, ", b =", b)
	// if a == b {   // nvalid operation: a == b (mismatched types [2]int and [3]int)
	// 	fmt.Println("equal")
	// } else {
	// 	fmt.Println("not equal")
	// }
}

func main() {
	test1()
	test2()
	test3()
}