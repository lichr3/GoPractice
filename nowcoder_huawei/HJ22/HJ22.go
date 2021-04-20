// HJ22.空汽水瓶
package main

import "fmt"

func main() {
	for i := 0; i < 10; i++ {
		var n int
		fmt.Scanln(&n)
		if n == 0 {
			return
		}
		fmt.Println(n/2)
	}
}
