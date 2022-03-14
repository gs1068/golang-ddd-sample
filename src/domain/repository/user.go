package repository

import (
	"github.com/gs1068/golang_ddd_sample/domain/model"
)

type UserRepository interface {
	Create(user *model.User) (*model.User, error)
	FindByID(id int) (*model.User, error)
	Update(user *model.User) (*model.User, error)
	Delete(user *model.User) error
}
