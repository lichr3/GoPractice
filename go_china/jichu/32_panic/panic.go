package main

import (
	"fmt"
)

func fullName(firstName *string, lastName *string) {
	defer fmt.Println("deferred call in fullName")
	if firstName == nil {
		panic("runtime error: first name cannot be nil")
	}
	if lastName == nil {
		panic("runtime error: last name cannot be nil")
	}
	fmt.Printf("%s %s\n", *firstName, *lastName)
	fmt.Println("returned normally from fullName")
}

func helper(fullName func(firstName *string, lastName *string), firstName *string, lastName *string) {
	defer fmt.Println("deferred call in helper")
	fullName(firstName, lastName)
}

func main() {
	defer fmt.Println("deferred call in main")
	firstName := "Elon"
	// 触发panic时，先自底向上的运行defer函数，到达函数顶层时，打印panic信息
	helper(fullName, &firstName, nil)
	fmt.Println("returned normally from main")
}
