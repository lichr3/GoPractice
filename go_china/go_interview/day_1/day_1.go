/*
问：以下代码的输出内容是？
考点：panic和defer的执行顺序
解析：
当函数发生 panic 时，它会终止运行，在执行完所有的延迟函数后，程序控制返回到该函数的调用方。然后程序会打印出 panic 信息，接着打印出堆栈跟踪（Stack Trace），最后程序终止。
*/

package main

import "fmt"

func main() {
	defer_call()
}

func defer_call() {
	defer func() { fmt.Println("打印前") }()
	defer func() { fmt.Println("打印中") }()
	defer func() { fmt.Println("打印后") }()

	panic("触发异常")
}