package main

import "oop/employee"

func main() {
	e := employee.New("sam", "Adolf", 30, 20)
	e.LeavesRemaining()
}
