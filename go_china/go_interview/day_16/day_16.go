package main

import "fmt"

func test1() {
	fmt.Println("test1考点：如何确定切片的len和cap")
	s := [3]int{1, 2, 3}
	a := s[:0]
	b := s[:2]
	c := s[1:2:cap(s)]
	fmt.Printf("len: %d, cap: %d\n",len(s), cap(s))
	help1(a)
	help1(b)
	help1(c)
}

func help1(s []int) {
	fmt.Printf("len: %d, cap: %d\n", len(s), cap(s))
}

func test2() {
	fmt.Printf("\ntest2考点：\n1.int类型不可与nil比较。\n2.当key不存在时，返回值类型对应的零值。\n3.可用第二个返回值判断map键值对是否存在，该返回值为bool类型。")
	m := make(map[string]int)
	m["a"] = 1
	// if v := m["b"]; v != nil {  // nvalid operation: v != nil (mismatched types int and nil)
	if v, ok := m["b"]; ok {
		fmt.Println(v)
	}
}


func main() {
	test1()
	test2()
}