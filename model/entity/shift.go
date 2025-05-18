package entity

import "time"

type (
	ShiftSchedule struct {
		ID        string
		StartTime time.Time
		EndTime   time.Time
		User      *User
	}

	ShiftRequest struct {
		ID     string
		User   User
		Status StatusShiftRequest
	}

	StatusShiftRequest int
)

const (
	StatusRequestPending StatusShiftRequest = iota
	StatusRequestApproved
	StatusRequestRejected
)

func (s StatusShiftRequest) String() string {
	return [...]string{"PENDING", "APPROVED", "REJECTED"}[s]
}
