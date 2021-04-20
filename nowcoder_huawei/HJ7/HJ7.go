package main

import (
	"fmt"
	"math"
)

func main() {
	var input float64
	fmt.Scanf("%f", &input)
	fmt.Println(math.Round(input))
}