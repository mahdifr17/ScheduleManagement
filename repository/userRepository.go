package repository

import "github.com/mahdifr17/ScheduleManagement/model/entity"

type UserRepository interface {
	GetUser(username, role string) (entity.User, error)
}
