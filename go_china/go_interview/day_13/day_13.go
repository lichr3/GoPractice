package main

import "fmt"


// var str string   // 可
// str := "test1"  	// 不可  
// str = "test1"	// 不可
// 以上两者会报错 syntax error: non-declaration statement outside function
// 即非声明语句不能出现在函数外部。
// var str = "test1"	// 可，感觉比较特殊，这里虽然赋值了，但是应该相当于初始化

func test1() {
	fmt.Println("test1考点：简短模式 := 不可用于全局变量。因为非声明语句不能出现在函数外部。\n可以理解成:=实际上会拆分为声明语句和赋值语句")
}

func hello(i int) {
	fmt.Println("i in defer:", i)
}

func test2() {
	fmt.Println("\ntest2考点：defer函数中传入的参数，是调用时刻的参数值，而不是执行时刻的参数值")
	i := 5
	defer hello(i)
	i = i + 10
	fmt.Println("i in test2:", i)
}

func main() {
	test1()
	test2()
}