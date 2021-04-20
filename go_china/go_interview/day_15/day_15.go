package main

import "fmt"

func test1() {
	fmt.Println("test1考点：声明但没有赋值的pointer、map、slice、channel、interface、func类型为nil")
	var s1 []int
	var s2 = []int{}
	s3 := make([]int, 0)
	help1(s1)
	help1(s2)
	help1(s3)
}

func help1(s []int) {
	if s == nil {
		fmt.Println("yes nil")
	} else {
		fmt.Println("no nil")
	}
}

type A interface {
	ShowA() int
}
type B interface {
	ShowB() int
}
type Work struct {
	i int
}
func (w Work) ShowA() int {
	return w.i + 10
}
func (w Work) ShowB() int {
	return w.i + 20
}
func test3() {
	fmt.Println("\ntest3考点：interface的使用")
	c := Work{3}
	var a A = c
	var b B = c
	fmt.Println(a.ShowA())
	fmt.Println(b.ShowB())
}

func test2() {
	fmt.Println("\ntest2考点：将int类型通过string(int)转换为string类型会转换成该以该数字作为编码对应的字符")
	fmt.Println("UTF-8编码中，十进制数65对应的符号是A")
	i := 65
	fmt.Println(string(i))
}

func main() {
	test1()
	test2()
	test3()
	func () {
		fmt.Println("11")
	}()
}

