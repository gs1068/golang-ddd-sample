package infra

import (
	"github.com/gs1068/golang_ddd_sample/domain/model"
	"github.com/gs1068/golang_ddd_sample/domain/repository"
	"github.com/jinzhu/gorm"
)

type UserRepository struct {
	Conn *gorm.DB
}

func NewUserRepository(conn *gorm.DB) repository.UserRepository {
	return &UserRepository{Conn: conn}
}

func (ur *UserRepository) Create(user *model.User) (*model.User, error) {
	if err := ur.Conn.Create(&user).Error; err != nil {
		return nil, err
	}

	return user, nil
}

func (ur *UserRepository) FindByID(id int) (*model.User, error) {
	user := &model.User{ID: id}

	if err := ur.Conn.First(&user).Error; err != nil {
		return nil, err
	}

	return user, nil
}

func (ur *UserRepository) FindAll() (*[]model.User, error) {
	user := &[]model.User{}

	if err := ur.Conn.Find(&user).Error; err != nil {
		return nil, err
	}

	return user, nil
}

func (ur *UserRepository) Update(user *model.User) (*model.User, error) {
	if err := ur.Conn.Model(&user).Update(&user).Error; err != nil {
		return nil, err
	}

	return user, nil
}

func (tr *UserRepository) Delete(user *model.User) error {
	if err := tr.Conn.Delete(&user).Error; err != nil {
		return err
	}

	return nil
}
