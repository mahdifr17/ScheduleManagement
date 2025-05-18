package repository

import "github.com/mahdifr17/ScheduleManagement/model/entity"

type ShiftRepositoryMemory struct {
	ShiftData []entity.ShiftSchedule
}

func (r *ShiftRepositoryMemory) AddShiftSchedule(request entity.ShiftSchedule) (*entity.ShiftSchedule, error) {
	return new(entity.ShiftSchedule), nil
}
