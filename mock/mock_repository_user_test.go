package mock_repository

import (
	"errors"
	reflect "reflect"
	"testing"

	gomock "github.com/golang/mock/gomock"
	model "github.com/gs1068/golang-ddd-sample/domain/model"
	"github.com/gs1068/golang-ddd-sample/usecase"
	"github.com/stretchr/testify/assert"
)

var err error
var errNotFound = errors.New("record not found")
var errEmptyUserName = errors.New("ユーザーネームを入力してください")
var expected = &model.User{
	ID:       1,
	UserName: "test",
}

func TestUserSuccessCreate(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	var args = &model.User{
		UserName: "test",
	}

	mockUser := NewMockUserRepository(mockCtrl)
	mockUser.EXPECT().Create(args).Return(expected, err)
	userUsecase := usecase.NewUserUsecase(mockUser)

	result, returnErr := userUsecase.Create("test")
	if returnErr != nil {
		t.Error("Actual Create() is not same as expected")
	}
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Actual Create() is not same as expected")
	}
}

func TestUserFailureCreate(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	mockUser := NewMockUserRepository(mockCtrl)
	userUsecase := usecase.NewUserUsecase(mockUser)
	_, returnErr := userUsecase.Create("")
	if returnErr == nil {
		t.Error("Actual Create() is not same as expected")
	}
	assert.Equal(t, returnErr, errEmptyUserName)
}

func TestUserFindAllSuccess(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	var expected = &[]model.User{}
	mockUser := NewMockUserRepository(mockCtrl)
	mockUser.EXPECT().FindAll().Return(expected, err)
	userUsecase := usecase.NewUserUsecase(mockUser)
	result, returnErr := userUsecase.FindAll()
	if returnErr != nil {
		t.Error("Actual Create() is not same as expected")
	}

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Actual Create() is not same as expected")
	}
}

func TestUserDeleteSeccess(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	mockUser := NewMockUserRepository(mockCtrl)
	mockUser.EXPECT().FindByID(gomock.Any()).Return(nil, err)
	mockUser.EXPECT().Delete(gomock.Any()).Return(nil)
	userUsecase := usecase.NewUserUsecase(mockUser)
	returnErr := userUsecase.Delete(1)
	if returnErr != nil {
		t.Error("Actual Delete() is not same as expected")
	}
}

func TestUserDeleteFailure(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	userID := 44444444
	var expected = &model.User{}
	mockUser := NewMockUserRepository(mockCtrl)
	mockUser.EXPECT().FindByID(userID).Return(expected, errNotFound)
	userUsecase := usecase.NewUserUsecase(mockUser)
	returnErr := userUsecase.Delete(userID)
	if returnErr == nil {
		t.Error("Actual Delete() is not same as expected")
	}
}

func TestUserFIndByIDSuccess(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	var userID = 1
	mockUser := NewMockUserRepository(mockCtrl)
	mockUser.EXPECT().FindByID(userID).Return(expected, err)
	userUsecase := usecase.NewUserUsecase(mockUser)
	result, returnErr := userUsecase.FindByID(userID)
	if returnErr != nil {
		t.Error("Actual FindByID() is not same as expected")
	}

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Actual FindByID() is not same as expected")
	}
}

func TestUserFindByIDFailure(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	userID := 44444444
	mockUser := NewMockUserRepository(mockCtrl)
	mockUser.EXPECT().FindByID(userID).Return(nil, errNotFound)
	userUsecase := usecase.NewUserUsecase(mockUser)
	_, returnErr := userUsecase.FindByID(userID)
	if returnErr == nil {
		t.Error("Actual FindByID() is not same as expected")
	}
}

func TestUserUpdateSuccess(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	var userID = 1
	var args = &model.User{
		ID:       userID,
		UserName: "test",
	}
	mockUser := NewMockUserRepository(mockCtrl)
	mockUser.EXPECT().FindByID(userID).Return(expected, err)
	mockUser.EXPECT().Update(args).Return(expected, err)
	userUsecase := usecase.NewUserUsecase(mockUser)
	result, returnErr := userUsecase.Update(userID, "test")
	if returnErr != nil {
		t.Error("Actual Update() is not same as expected")
	}
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Actual Create() is not same as expected")
	}
}

func TestUserUpdateFailureNotFound(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	userID := 44444444
	mockUser := NewMockUserRepository(mockCtrl)
	mockUser.EXPECT().FindByID(userID).Return(nil, errNotFound)
	userUsecase := usecase.NewUserUsecase(mockUser)
	_, returnErr := userUsecase.Update(userID, "test")
	if returnErr == nil {
		t.Error("Actual Update() is not same as expected")
	}
	if !reflect.DeepEqual(errNotFound, returnErr) {
		t.Errorf("Actual Create() is not same as expected")
	}
}

func TestUserUpdateFailureEmptyUserName(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	var userID = 1
	mockUser := NewMockUserRepository(mockCtrl)
	mockUser.EXPECT().FindByID(userID).Return(expected, err)
	userUsecase := usecase.NewUserUsecase(mockUser)
	_, returnErr := userUsecase.Update(userID, "")
	if returnErr == nil {
		t.Error("Actual Update() is not same as expected")
	}
	if !reflect.DeepEqual(returnErr, errEmptyUserName) {
		t.Errorf("Actual Create() is not same as expected")
	}
}
