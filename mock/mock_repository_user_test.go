package mock_repository

import (
	"errors"
	"fmt"
	reflect "reflect"
	"testing"

	gomock "github.com/golang/mock/gomock"
	model "github.com/gs1068/golang_ddd_sample/domain/model"
	"github.com/gs1068/golang_ddd_sample/usecase"
	"github.com/stretchr/testify/assert"
)

var err error
var expected = &model.User{}

func TestUserSuccessCreate(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
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
	fmt.Print("エラーbool,", reflect.TypeOf(err) == reflect.TypeOf(errExpected))
	if err == nil {
		t.Error("Actual Create() is not same as expected")
	}
	assert.NotNil(t, err)

	// if !reflect.DeepEqual(result, expected) {
	// 	t.Errorf("Actual Create() is not same as expected")
	// }
}

// func TestUserFailureCreate2(t *testing.T) {
// 	mockCtrl := gomock.NewController(t)
// 	errExpected := fmt.Errorf("ユーザーネームを入力してください")
// 	defer mockCtrl.Finish()
// 	var want = &model.User{
// 		UserName: "",
// 	}

// 	var args = &model.User{
// 		UserName: "",
// 	}

// 	mockUser := NewMockUserRepository(mockCtrl)
// 	mockUser.EXPECT().Create(want).Return(expected, err)
// 	userRepository := repository.UserRepository(mockUser)
// 	result, err := userRepository.Create(args)

// 	fmt.Println(err, errExpected)
// 	fmt.Print(result)

// 	if err == errExpected {
// 		t.Error("Actual Create() is not same as expected")
// 	}

// 	if !reflect.DeepEqual(result, expected) {
// 		t.Errorf("Actual Create() is not same as expected")
// 	}
// }
