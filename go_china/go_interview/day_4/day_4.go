package main

import "fmt"


// 1.问：以下代码是否能通过编译，不能的话原因是什么？如果能，输出什么？
// 考点：new的返回类型，append的参数类型
func test1() {
	fmt.Println("test1:")
	list := new([]int)
	// list = append(list, 1)  // 错误：此处list为指针对象(*[]int)，只有slice([]int)才可以使用append
	*list = append(*list, 1)	// 修正：使用*传入slice
	fmt.Println(*list)
}
/*
补充：一般slice、map、channel的初始化建议使用make
*/

// 2.问：以下代码是否能编译通过，如何可以，输出什么？
// 考点： append的参数类型
func test2() {
	fmt.Println("test2:")
	s1 := []int{1, 2, 3}
	s2 := []int{4, 5}
	// s1 = append(s1, s2)  // 错误：append从只有第一个元素师切片，剩下的都是元素。
	s1 = append(s1, s2...)	// 修正：使用...打散切片传入参数。注：...的另一个作用是用于接收任意人数的函数参数。
	fmt.Println(s1)
}

// 3.问：以下代码是否能编译通过，如何可以，输出什么？
// 考点：声明变量的简短模式
// var size := 1024  
// 首先，没有 var size := 1024 这种形式，要么 size := 1024	, 不能 var size := 1024；
// !!!!! 并且， := 简短模式只能在函数内部使用!!!!!
var size = 1024
// 如果使用 size := 1024 会有语法错误：syntax error: non-declaration statement outside function body
var max_size = 2 * 1024 
func test3() {
	fmt.Println("test3:")
	fmt.Println(size, max_size)
}
/*
常见的声明方式：
var foo int // declaration without initialization
var foo int = 42 // declaration with initialization
var foo, bar int = 42, 1302 // declare/init multiple vars at once
var foo = 42 // type omitted, will be inferred
foo := 42 // shorthand, only in func bodies, implicit type
const constant = "This is a constant"

简短声明模式 := 的限制：
1.必须使用显示初始化；
2.不能提供数据类型，编译器会自动推导；
3.只能在函数内部使用简短模式
*/


func main() {
	test1()
	test2()
	test3()
}