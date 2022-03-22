package mock_repository

import (
	reflect "reflect"
	"testing"

	gomock "github.com/golang/mock/gomock"
	model "github.com/gs1068/golang_ddd_sample/domain/model"
	"github.com/gs1068/golang_ddd_sample/usecase"
)

var args = &model.User{
	UserName: "test",
}

func TestCreate(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	var want = &model.User{}
	var err error

	mockUser := NewMockUserRepository(mockCtrl)
	mockUser.EXPECT().Create(args).Return(want, err)

	userUsecase := usecase.NewUserUsecase(mockUser)
	result, err := userUsecase.Create("test")

	if err != nil {
		t.Error("Actual FindAll() is not same as expected")
	}

	if !reflect.DeepEqual(result, want) {
		t.Errorf("Actual FindAll() is not same as expected")
	}
}
