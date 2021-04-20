package main

import (
    "fmt"
    "net"
)

func main() {
    addr, err := net.LookupHost("baidu.com")
    if err, ok := err.(*net.DNSError); ok {  // 类型断言
        if err.Timeout() {
            fmt.Println("operation timed out")
        } else if err.Temporary() {
            fmt.Println("temporary error")
        } else {
            fmt.Println("generic error: ", err)
        }
        return
    }
    fmt.Println(addr)
}