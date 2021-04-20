package main

import "fmt"

func test1() {
	fmt.Println("test1考点，不可通过下标改变string类型，会报错：\ncannot assign to str[0] (strings are immutable)")
	str := "hello"
	//str[0] = 'x'  // cannot assign to str[0] (strings are immutable)
	fmt.Println("str =",str)
}

func incr(p *int) int {
	*p++
	return *p
}

func test2() {
	fmt.Println("\ntest2考点：传入指针修改整型值")
	p := 1
	incr(&p)
	fmt.Println(p)
}

func test3() {
	fmt.Println("\ntest3考点（自制）：string是8-bit bytes的集合，通常（但不一定都是）UTF-8编码的文本。string可为空（长度为0）但不会是nil，string是不可修改类型（immutable）")
	s := "hello你好"
	for a, b := range s{
		fmt.Printf("%T, %d--%v--%c\n",b, a, b, b)
	}
}

func main() {
	test1()
	test2()
	test3()
}