package repository

import "github.com/mahdifr17/ScheduleManagement/model/entity"

// implements user repository
type UserRepositorySqlite struct {
	UserData []entity.User
}

func (r *UserRepositorySqlite) GetUser(username, role string) (entity.User, error) {
	return entity.User{}, nil
}
