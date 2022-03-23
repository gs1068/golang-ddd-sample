package mock_repository

import (
	"errors"
	"fmt"
	reflect "reflect"
	"testing"

	gomock "github.com/golang/mock/gomock"
	model "github.com/gs1068/golang-ddd-sample/domain/model"
	"github.com/gs1068/golang-ddd-sample/usecase"
	"github.com/stretchr/testify/assert"
)

var err error

func TestUserSuccessCreate(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	var expected = &model.User{}
	var args = &model.User{
		UserName: "test",
	}

	mockUser := NewMockUserRepository(mockCtrl)
	mockUser.EXPECT().Create(args).Return(expected, err)
	userUsecase := usecase.NewUserUsecase(mockUser)

	result, err := userUsecase.Create("test")
	if err != nil {
		t.Error("Actual Create() is not same as expected")
	}

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Actual Create() is not same as expected")
	}
}

func TestUserFailureCreate(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	errExpected := errors.New("ユーザーネームを入力してください")
	defer mockCtrl.Finish()

	mockUser := NewMockUserRepository(mockCtrl)
	userUsecase := usecase.NewUserUsecase(mockUser)
	_, err := userUsecase.Create("")
	if err == nil {
		t.Error("Actual Create() is not same as expected")
	}
	assert.Equal(t, err, errExpected)
}

func TestUserFindAllSuccess(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	var expected = &[]model.User{}
	mockUser := NewMockUserRepository(mockCtrl)
	mockUser.EXPECT().FindAll().Return(expected, err)
	userUsecase := usecase.NewUserUsecase(mockUser)
	result, err := userUsecase.FindAll()
	if err != nil {
		t.Error("Actual Create() is not same as expected")
	}

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Actual Create() is not same as expected")
	}
}

func TestUserDeleteSeccess(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	var expected = &model.User{}
	mockUser := NewMockUserRepository(mockCtrl)
	mockUser.EXPECT().FindByID(gomock.Any()).Return(expected, err)
	mockUser.EXPECT().Delete(gomock.Any()).Return(err)
	userUsecase := usecase.NewUserUsecase(mockUser)
	err := userUsecase.Delete(1)
	if err != nil {
		t.Error("Actual Delete() is not same as expected")
	}
}

func TestUserDeleteFailure(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	findID := 44444
	// var args = &model.User{
	// 	ID:       0,
	// 	UserName: "",
	// }
	var expected = &model.User{}
	mockUser := NewMockUserRepository(mockCtrl)
	// mockUser.EXPECT().FindByID(findID).Return(expected, err)
	// fmt.Print("エラー", err)
	// mockUser.EXPECT().Delete(expected).Return(err)
	userUsecase := usecase.NewUserUsecase(mockUser)

	err := userUsecase.Delete(findID)

	fmt.Print(expected)

	fmt.Print("エラー", err)
	if err != nil {
		t.Error("Actual Delete() is not same as expected")
	}
}
