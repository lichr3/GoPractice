package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("begin")
	file, _ := os.Open("/Users/muzi/go/src/leetcode/test.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}