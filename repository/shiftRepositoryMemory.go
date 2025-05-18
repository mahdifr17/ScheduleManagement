package repository

import (
	"github.com/mahdifr17/ScheduleManagement/model"
	"github.com/mahdifr17/ScheduleManagement/model/entity"
)

type ShiftRepositoryMemory struct {
	ShiftData        []entity.ShiftSchedule
	ShiftRequestData []entity.ShiftRequest
}

func NewShiftRepositoryMemory() ShiftRepository {
	rp := new(ShiftRepositoryMemory)
	rp.ShiftData = make([]entity.ShiftSchedule, 0)
	return rp
}

func NewShiftRequestRepositoryMemory() ShiftRequestRepository {
	rp := new(ShiftRepositoryMemory)
	rp.ShiftRequestData = make([]entity.ShiftRequest, 0)
	return rp
}

/* Implements ShiftRepository */

func (r *ShiftRepositoryMemory) InsertShiftSchedule(request entity.ShiftSchedule) (entity.ShiftSchedule, error) {
	return entity.ShiftSchedule{}, nil
}

func (r *ShiftRepositoryMemory) UpdateShiftSchedule(id string, request entity.ShiftSchedule) (entity.ShiftSchedule, error) {
	return entity.ShiftSchedule{}, nil
}

func (r *ShiftRepositoryMemory) GetShiftSchedule(filter model.Filter) ([]entity.ShiftSchedule, error) {
	return []entity.ShiftSchedule{}, nil
}

/* Implements ShiftRequestRepository */

func (r *ShiftRepositoryMemory) InsertShiftRequest(request entity.ShiftRequest) (entity.ShiftRequest, error) {
	return entity.ShiftRequest{}, nil
}

func (r *ShiftRepositoryMemory) ApproveShiftRequest(id string) (entity.ShiftRequest, error) {
	return entity.ShiftRequest{}, nil
}

func (r *ShiftRepositoryMemory) RejectShiftRequest(id string) (entity.ShiftRequest, error) {
	return entity.ShiftRequest{}, nil
}

func (r *ShiftRepositoryMemory) GetShiftRequest(filter model.Filter) ([]entity.ShiftRequest, error) {
	return []entity.ShiftRequest{}, nil
}
