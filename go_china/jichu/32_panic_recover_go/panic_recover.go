package main

import (
	"fmt"
	"runtime/debug"
)

func r() {
	if r := recover(); r != nil {
		fmt.Println("Recovered", r)
		debug.PrintStack() // recover后打印堆栈跟踪
	}
}

func a() {
	defer r()
	n := []int{5, 7, 4}
	fmt.Println(n[3])
	fmt.Println("normally returned from a")
}

func main() {
	a()
	fmt.Println("normally returned from main")
}
