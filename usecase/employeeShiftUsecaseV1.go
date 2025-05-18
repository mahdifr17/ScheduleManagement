package usecase

import "github.com/mahdifr17/ScheduleManagement/repository"

type EmployeeShiftUsecaseV1 struct {
	ShiftRP        repository.ShiftRepository
	ShiftRequestRP repository.ShiftRequestRepository
}

func NewEmployeeShiftUsecaseV1(
	shiftRP repository.ShiftRepository,
	shiftRequestRP repository.ShiftRequestRepository,
) EmployeeShiftUsecase {
	return &EmployeeShiftUsecaseV1{
		ShiftRP:        shiftRP,
		ShiftRequestRP: shiftRequestRP,
	}
}

func (u *EmployeeShiftUsecaseV1) GetAssignedShift() {

}

func (u *EmployeeShiftUsecaseV1) GetAvailableShift() {

}

func (u *EmployeeShiftUsecaseV1) CreateShiftRequest() {

}

func (u *EmployeeShiftUsecaseV1) GetShiftRequest() {

}
