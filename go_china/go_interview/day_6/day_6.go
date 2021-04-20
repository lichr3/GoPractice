package main

import "fmt"

type person struct {
	name string
	age int
}

type MyInt1 int  // 类型定义
type MyInt2 = int	// 类型别名

func test2() {
	var i int = 0
	// var i1 MyInt1 = i  // MyInt1是一个由int类型定义的一个新类型，但不是int
	var i1 MyInt1 =  MyInt1(i)  // 修正：强制转换一下
	var i2 MyInt2 = i
	fmt.Println(i1, i2)
}

func test3() {
	a := []int{7, 8, 9}
	fmt.Printf("%+v\n", a)
	ap(a)
	fmt.Printf("期望得到[7, 8, 9, 10], 结果得到 %+v\n", a)
	app(a)
	fmt.Printf("%+v\n", a)
}

func ap(a []int) {
	a = append(a, 10)
}

func app(a []int) {
	a[0] = 1
}

func main () {
	fmt.Println("test1考点: 由于go的语法糖，指针类型的strunt可以不解引用直接访问其成员")
	p1 := &person{"lichr3", 24}
	fmt.Printf("type of p1: %T\n", p1)
	fmt.Printf("p1.name = %v, (*p1).name = %v\n\n", p1.name, (*p1).name)

	fmt.Println("test2考点: 类型定义和类型别名的区别")
	test2()
	fmt.Println()

	fmt.Println("test3考点: append会使得slice的底层数组重新分配内存，扩容后的地址不再是原来数组的指针地址")
	test3()

	fmt.Printf("abc%d", 123)
}