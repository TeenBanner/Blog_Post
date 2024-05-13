package UserDomain

import (
	"github.com/TeenBanner/Inventory_system/User/Domain/model"
)

// userStorage use user methods on DB
type UserStorage interface {
	CreateUser(user model2.User) error
	GetUserByEmail(email string) (model2.User, error)
	UpdateUserName(email, name string) error
	UpdateUserEmail(ActualEmail, NewEmail string) error
	GetAllUsers() ([]model2.User, error)
}

// AdminStorage extends userStorge for Admin type functions
