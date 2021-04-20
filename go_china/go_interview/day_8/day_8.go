package main

import "fmt"

func test1() {
	fmt.Println(`test1：关于init函数，下面说法正确的是： AB
  A. 一个包中，可以包含多个 init 函数；
  B. 程序编译时，先执行依赖包的 init 函数，再执行 main 包内的 init 函数；
  C. main 包中，不能有 init 函数；
  D. init 函数可以被其他函数调用；

解析：
  1.init() 函数是用于程序执行前做包的初始化的函数，比如初始化包里的变量等;
  2.一个包可以出线多个 init() 函数,一个源文件也可以包含多个 init() 函数；
  3.同一个包中多个 init() 函数的执行顺序没有明确定义，但是不同包的init函数是根据包导入的依赖关系决定的;
  4.init() 函数在代码中不能被显示调用、不能被引用（赋值给函数变量），否则出现编译错误;
  5.一个包被引用多次，如 A import B,C import B,A import C，B 被引用多次，但 B 包只会初始化一次；
  6.引入包，不可出现死循坏。即 A import B,B import A，这种情况编译失败；
--------------------------------- 分界线 ---------------------------------
  `)
}


func hello() []string {
	return nil
}

func test2() {
	fmt.Println(
		`test2考点：
1.用函数赋值时，注意返回的事函数还是函数的返回值.
  h = hello 返回的是函数;
  h := hello() 返回的是函数返回值
2.虽然slice和func都是不可比较类型，但是可以和nil进行比较`)
	h := hello   // 此处是把hello函数赋值给h，而不是返回hello()的返回值nil
	if h == nil {	// 若此处为h()，则 h() == nil
		fmt.Println("nil")
	} else {
		fmt.Println("not nil")
	}
	
	fmt.Println("--------------------------------- 分界线 ---------------------------------")
}

func GetValue() int {
	return 1
}


func test3() {
	fmt.Print("\ntest3考点：只有interface类型可以使用type switch（类型选择）\n")
	i := GetValue()
	j := interface{}(i)  // 修正：将int类型转换成interface类型才能使用type switch
	switch j.(type) {  // cannot type switch on non-interface value i (type int)
	case int:
		println("int")
	case string:
		println("string")
	case interface{}:
		println("interface")
	default:
		println("unknown")
	}
}

func main() {
	test1()
	test2()
	test3()
}