package usecase

import (
	"github.com/gs1068/golang_ddd_sample/domain/model"
	"github.com/gs1068/golang_ddd_sample/domain/repository"
)

type UserUsecase interface {
	Create(usename string) (*model.User, error)
	FindByID(id int) (*model.User, error)
	Update(id int, username string) (*model.User, error)
	Delete(id int) error
}

type userUsecase struct {
	userRepo repository.UserRepository
}

func NewUserUsecase(userRepo repository.UserRepository) UserUsecase {
	return &userUsecase{userRepo: userRepo}
}

func (uu *userUsecase) Create(username string) (*model.User, error) {
	user, err := model.NewUser(username)
	if err != nil {
		return nil, err
	}

	createdUser, err := uu.userRepo.Create(user)
	if err != nil {
		return nil, err
	}

	return createdUser, nil
}

func (uu *userUsecase) FindByID(id int) (*model.User, error) {
	foundUser, err := uu.userRepo.FindByID(id)
	if err != nil {
		return nil, err
	}

	return foundUser, nil
}

func (uu *userUsecase) Update(id int, username string) (*model.User, error) {
	targetUser, err := uu.userRepo.FindByID(id)
	if err != nil {
		return nil, err
	}

	err = targetUser.Set(username)
	if err != nil {
		return nil, err
	}

	updatedUser, err := uu.userRepo.Update(targetUser)
	if err != nil {
		return nil, err
	}

	return updatedUser, nil
}

func (uu *userUsecase) Delete(id int) error {
	user, err := uu.userRepo.FindByID(id)
	if err != nil {
		return err
	}

	err = uu.userRepo.Delete(user)
	if err != nil {
		return err
	}

	return nil
}
