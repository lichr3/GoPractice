package main

import "fmt"

func test1() {
	fmt.Println("test1:class不是go语言的关键字，go有25个关键字：")
	fmt.Println(`
	break 		default		func	interface	select
	case 		defer		go	map		struct
	chan		else		goto	package		switch
	const		fallthrough	if	range		type
	continue	for		import	return		var		
	`)
}

func test2() {
	fmt.Printf("\ntest2考点：%%+d占位符的作用，打印带符号整型\n")
	i := -5
	j := +5
	k := 5
	fmt.Printf("%+d %+d %+d\n\n", i, j, k)
}

type People struct{}

func (p *People) ShowA() {
	fmt.Println("showA")
	p.ShowB()
}

func (p *People) ShowB() {
	fmt.Println("showB")
}

type Teacher struct {
	People
}

func (t *Teacher) ShowB() {
	fmt.Println("teacher showB")
}

func test3() {
	fmt.Println("test3考点：结构体嵌套，外部类型和内部类型有同名方法时，会调用外部类型的方法，而不是内部类型的方法")
	t := Teacher{}
	t.ShowB()
	t.ShowA()
}



func main() {
	test1()
	test2()
	test3()
}