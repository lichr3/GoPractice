package main

import "fmt"

func increaseA() int {
	var i int
	defer func() {
		i++
	}()
	return i
}

func increaseB() (r int) {
	defer func() {
		r++
	}()
	return r
}

func test1() {
	fmt.Println("test1考点：理解defer和return的执行顺序；以及return的机制，从而理解为什么不带命名返回值函数和带命名返回值函数会有差异")
	fmt.Println(increaseA())
	fmt.Println(increaseB())
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

func test2() {
	fmt.Println("\ntest2考点：类型断言")
    var a A = Work{3}
    s, ok := a.(Work)
	fmt.Printf("%T, %v\n", s, s)
	fmt.Println(ok)
    fmt.Println(s.ShowA())
    fmt.Println(s.ShowB())
}

func test3() {
	fmt.Println("\ntest3考点：当多值赋值时，:= 左边的变量无论声明与否都可以")
	var x int
	// x, _ := 1, 2   // wrong
	x, _ = 1, 2
	x, y := 1, 2  // right!
	fmt.Println(x, y)
}

func main() {
	test1()
	test2()
	test3()
}