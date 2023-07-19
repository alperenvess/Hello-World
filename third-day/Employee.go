package employee

func New() *employee {
	return new(employee)
}

type employee struct {
	FirstName     string
	LastName      string
	LeavesAllowed int
	Leavestaken   int
}

func (empl *employee) TakeHolidays(days int) bool {
	if empl.LeavesTaken+days <= empl.LeavesAllowed {
		empl.LeavesTaken += days
	}
	return false
}

func (empl *employee) LimitExceeded(days int) bool {
	return empl.LeavesTaken+days > empl.LeavesAllowed
}
