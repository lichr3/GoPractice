package main

import "fmt"


func main() {
	defer func() {
		fmt.Println(`
考点：
1. 结构体只能比较是否相等，不能比较大小
2. 结构体内的所有成员都可以比较的时候，这个结构体才有可能可以比较
3. 不可比较的类型有：slice、map，func
4. 两个结构体的成员类型都可以比较，且成员排序也相同时，才可比较
		`)
	}()

	sn1 := struct {
		age int
		name string
	} {age: 11, name: "qq"}

	sn2 := struct {
		age int
		name string
	} {age: 11, name: "qq"}

	if sn1 == sn2 {
		fmt.Println("sn1 == sn2, 因为结构体内的成员都可以比较，且值和顺序都相同")
	} else {
		fmt.Println("sn1 != sn2")
	}

	fmt.Println("map 类型不可比较")
	// sm1 := struct {
	// 	age int
	// 	m map[string]string
	// } {age: 11, m: map[string]string{"a": "1"}}

	// sm2 := struct {
	// 	age int
	// 	m map[string]string
	// } {age: 11, m: map[string]string{"a": "1"}}

	// if sm1 == sm2 {
	// 	fmt.Println("sm1 == sm2")
	// } else {
	// 	fmt.Println("sm1 != sm2")
	// }
}