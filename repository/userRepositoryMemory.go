package repository

import "github.com/mahdifr17/ScheduleManagement/model/entity"

// implements user repository
type UserRepositoryMemory struct {
	UserData []entity.User
}

func (r *UserRepositoryMemory) GetUser(username, role string) (entity.User, error) {
	return entity.User{}, nil
}
