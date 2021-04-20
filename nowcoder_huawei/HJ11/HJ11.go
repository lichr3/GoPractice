// HJ11.数字颠倒

package main

import "fmt"

func main() {
    var input, output string
    fmt.Scanf("%s", &input)
    for i := len(input)-1 ; i >= 0; i-- {
        output += string(input[i])
    }
    fmt.Println(output)
}