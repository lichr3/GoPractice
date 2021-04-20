package employee

import (
	"fmt"
)

// 结构体首字母小写，变成不可引用，防止其他包对它对访问
type employee struct {
	FirstName   string
	LastName    string
	TotalLeaves int
	LeavesTaken int
}

func New(firstName string, lastName string, totalLeaves int, leavesTaken int) employee {
	e := employee{firstName, lastName, totalLeaves, leavesTaken}
	return e
}

func (e employee) LeavesRemaining() {
	fmt.Printf("%s %s has %d leaves remaining", e.FirstName, e.LastName, (e.TotalLeaves - e.LeavesTaken))
}
