package usecase

import "github.com/mahdifr17/ScheduleManagement/repository"

type AdminShiftUsecaseV1 struct {
	ShiftRP        repository.ShiftRepository
	ShiftRequestRP repository.ShiftRequestRepository
}

func NewAdminShiftUsecaseV1(
	shiftRP repository.ShiftRepository,
	shiftRequestRP repository.ShiftRequestRepository,
) AdminShiftUsecase {
	return &AdminShiftUsecaseV1{
		ShiftRP:        shiftRP,
		ShiftRequestRP: shiftRequestRP,
	}
}

func (u *AdminShiftUsecaseV1) CreateShift() {

}

func (u *AdminShiftUsecaseV1) EditShift() {

}

func (u *AdminShiftUsecaseV1) GetAllShiftRequest() {

}

func (u *AdminShiftUsecaseV1) ApproveShiftRequest() {

}

func (u *AdminShiftUsecaseV1) RejectShiftRequest() {

}

func (u *AdminShiftUsecaseV1) GetAllFinalizedShift() {

}
