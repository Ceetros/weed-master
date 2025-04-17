package Repositories

import (
	"Api/Data/Models"
)

type IUserRepository interface {
	GetUserByEmail(email string) (Models.User, error)
	CreateUser(user Models.User) error
}
