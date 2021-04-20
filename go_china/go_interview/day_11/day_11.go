package main

import "fmt"

func test1() {
	fmt.Println("test1考点：cap()适用于array, slize, channel ，但不适用于map")
	my_arr := [...]int{1, 2, 3, 4, 5}
	my_slice := my_arr[0:3:4]
	// my_map := make(map[string]int)
	my_chan := make(chan int, 3)
	fmt.Println("cap(my_arr):", cap(my_arr))
	fmt.Println("cap(my_slice):", cap(my_slice))
	// fmt.Println("cap(my_map):", cap(my_map))
	fmt.Println("cap(my_chan):", cap(my_chan))
	fmt.Println()
}

func test2() {
	fmt.Println("test2考点：当且仅当接口的动态值和动态类型都为nil时，接口类型值才为nil")
	var i interface{}
	if i == nil {
		fmt.Println("nil")
		return
	}
	fmt.Println("not nil")
}

func test3() {
	fmt.Println("\ntest3考点：delete在map中的运用：\n" +
		"1.delete删除amp不存在的键值对时不会报错，相当于没有任何作用\n" +
		"2.map获取不存在的键值对时候，返回值类型对应的零值")
	s := make(map[string]interface{})
	delete(s, "h")
	fmt.Println(s["h"])
}

func main() {
	test1()
	test2()
	test3()
}