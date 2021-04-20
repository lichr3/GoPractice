package main

import (
	"fmt"
	// "time"
)

func test1() {
	fmt.Println("test1考点：var在声明chan、map后要再make一下分配空间并初始化，否则无法使用。并且，var声明指针后也要用new分配空间并初始化后才能使用")
	// var ch chan int	  // 无法使用，没有分配空间
	ch := make(chan int) // 修正：make才会给channel分配空间并初始化
	go func() {
		ch <- 1
	}()
	i := <-ch
	// time.Sleep(1 * time.Second)
	fmt.Println(i)
}

type person struct {
	name string
}

func test2() {
	fmt.Println("test2考点：打印一个 map 中不存在的值时，返回元素类型的零值")
	var m map[person]int
	p := person{"mike"}
	fmt.Println(m[p])
}

func hello(num ...int) {
	fmt.Printf("切片在可变函数中的地址：%p\n", num)
	num[0] = 18
}

func test3() {
	fmt.Println("test3考点：把一个切片用...解序列后传入可变参数函数，本质上就是传入的切片，在函数中修改切片会影响原切片")
	i := []int{5, 6, 7}
	fmt.Printf("切片的地址：%p\n", i)
	hello(i...)
	fmt.Println(i[0])
}

func main() {
	test1()
	fmt.Print("\n-------------------- 分界线 --------------------\n\n")
	test2()
	fmt.Print("\n-------------------- 分界线 --------------------\n\n")
	test3()
}