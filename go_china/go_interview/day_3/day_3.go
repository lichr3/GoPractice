package main

import "fmt"

// 1.考点：s1中初始化长度为5，append 3个元素后，一共有8个元素
func test1() {
	fmt.Println("test1_1:")
	s1 := make ([]int, 5)
	s1 = append(s1, 1, 2, 3)
	fmt.Println(s1)

	s2 := make([]int, 0)
	s2 = append(s2, 1, 2, 3, 4)
	fmt.Println(s2)
}

// 2.考点：命名函数参数 和 未命名函数参数不能混用
// syntax error: mixed named and unnamed function parameters
// func funcMui(x,y int) (sum int, error) {
// 	return x + y, nil
// }


func main () {
	test1()
	// fmt.Println(funcMui(1, 2))
}

/*
3. new() 与 make() 的区别
解析：
1）适用对象不同：make只用于slice、map、channel
2）返回结果不同：new(T)分配内存后，返回一个指针，类型为T的零值。适用于值类型，如数组、结构体等；   make(T, args)分配内存后，返回初始化的T类型的值，不是零值也不是指针，而是引用
*/