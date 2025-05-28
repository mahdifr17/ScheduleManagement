package repository

import (
	"database/sql"

	"github.com/mahdifr17/ScheduleManagement/model"
	"github.com/mahdifr17/ScheduleManagement/model/entity"
)

type (
	ShiftRepositorySqlite struct {
		db *sql.DB
	}
)

func NewShiftRepositorySqlite(db *sql.DB) ShiftRepository {
	rp := new(ShiftRepositorySqlite)
	rp.db = db
	return rp
}

func NewShiftRequestRepositorySqlite(db *sql.DB) ShiftRequestRepository {
	rp := new(ShiftRepositorySqlite)
	rp.db = db
	return rp
}

/* Implements ShiftRepository */

func (r *ShiftRepositorySqlite) InsertShiftSchedule(request entity.ShiftSchedule) (entity.ShiftSchedule, error) {
	return entity.ShiftSchedule{}, nil
}

func (r *ShiftRepositorySqlite) UpdateShiftSchedule(id string, request entity.ShiftSchedule) (entity.ShiftSchedule, error) {
	return entity.ShiftSchedule{}, nil
}

func (r *ShiftRepositorySqlite) GetShiftSchedule(filter model.Filter) ([]entity.ShiftSchedule, error) {
	return []entity.ShiftSchedule{}, nil
}

/* Implements ShiftRequestRepository */

func (r *ShiftRepositorySqlite) InsertShiftRequest(request entity.ShiftRequest) (entity.ShiftRequest, error) {
	return entity.ShiftRequest{}, nil
}

func (r *ShiftRepositorySqlite) ApproveShiftRequest(id string) (entity.ShiftRequest, error) {
	return entity.ShiftRequest{}, nil
}

func (r *ShiftRepositorySqlite) RejectShiftRequest(id string) (entity.ShiftRequest, error) {
	return entity.ShiftRequest{}, nil
}

func (r *ShiftRepositorySqlite) GetShiftRequest(filter model.Filter) ([]entity.ShiftRequest, error) {
	return []entity.ShiftRequest{}, nil
}
