package repository

import (
	"github.com/mahdifr17/ScheduleManagement/model"
	"github.com/mahdifr17/ScheduleManagement/model/entity"
)

type (
	ShiftRepository interface {
		InsertShiftSchedule(request entity.ShiftSchedule) (entity.ShiftSchedule, error)
		UpdateShiftSchedule(id string, request entity.ShiftSchedule) (entity.ShiftSchedule, error)
		GetShiftSchedule(filter model.Filter) ([]entity.ShiftSchedule, error)

		// GetAllFinalizedShift() []entity.ShiftSchedule
		// GetAssignedShift() []entity.ShiftSchedule
		// GetAvailableShift() []entity.ShiftSchedule
	}

	ShiftRequestRepository interface {
		InsertShiftRequest(request entity.ShiftRequest) (entity.ShiftRequest, error)
		ApproveShiftRequest(id string) (entity.ShiftRequest, error)
		RejectShiftRequest(id string) (entity.ShiftRequest, error)
		GetShiftRequest(filter model.Filter) ([]entity.ShiftRequest, error)

		// GetAllShiftRequest() ([]entity.ShiftRequest, error)
		// GetEmployeeShiftRequest() ([]entity.ShiftRequest, error)
	}
)
