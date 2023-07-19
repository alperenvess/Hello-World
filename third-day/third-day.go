package main

import (
	"errors"
	"fmt"
)

func New() *employee {
	return New(employee)
}

type

var ErrLimitExceeded = errors.New("Limit Exceeded")
var ErrEmployeeNotActive = errors.New("Employee status not active")



type employee struct {
	FirstName string

	LastName string

	leavesAllowed int

	LeavesTaken int
}

func (empl *employee) TakeHolidays(days int) bool {

	newTaken := empl.LeavesTaken + days

	if empl.LimitExceeded(newTaken) {

		return false

	}

	empl.LeavesTaken = newTaken

	return true

}

func (empl *employee) LimitExceeded(newDaysTaken int) bool {

	if newDaysTaken > empl.leavesAllowed {

		return true

	}

	return false

}

func main() {

	empl := new(employee)
	empl.FirstName = "Alperen"
	empl.leavesAllowed = 40
	fmt.Println(empl.FirstName)

	err := empl.TakeHolidays(50)

	if err != nil {

		var target = &employee.ErrLimitExceeded{}
		if errors.As(err, &target) {
			fmt.Printf("Cannot take holidays: %v", err)
		}	
	}
	fmt.Printf("total taken: %d \n",empl.LeavesTaken)

	err = empl.TakeHolidays(10)

}
