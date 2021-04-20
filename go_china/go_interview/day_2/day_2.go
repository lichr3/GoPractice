/*
问：test1()输出结果是什么？
考点：for range 中的变量是值传递（创建每个元素的副本，而非引用），所以使用指针的话，
	最后指向的值都是最终那个值，之前的值都会被逐一覆盖。
解析：创建一个新的变量来赋值
*/

package main

import "fmt"

func test1() {

	fmt.Println("test1:")

	slice := []int{0, 1, 2, 3}
	m := make(map[int]*int)

	for key, val := range slice {
		m[key] = &val
	}

	for k, v := range m {
		fmt.Println(k, "->", *v)
	}

	fmt.Println("结果都是3")
}

func answer1() {
	
	fmt.Println("answer1:")

	slice := []int{0, 1, 2, 3}
	m := make(map[int]*int)

	for key, val := range slice {
		value := val	// 修改的地方
		m[key] = &value
	}

	for k, v := range m {
		fmt.Println(k, "->", *v)
	}
}

func main() {
	test1()
	answer1()
}