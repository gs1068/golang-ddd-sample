package mock_repository

import (
	"fmt"
	reflect "reflect"
	"testing"

	gomock "github.com/golang/mock/gomock"
	model "github.com/gs1068/golang_ddd_sample/domain/model"
	"github.com/gs1068/golang_ddd_sample/usecase"
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
	errExpected := fmt.Errorf("ユーザーネームを入力してください")
	defer mockCtrl.Finish()
	var args = &model.User{
		UserName: "",
	}

	mockUser := NewMockUserRepository(mockCtrl)
	mockUser.EXPECT().Create(args).Return(nil, err)
	userUsecase := usecase.NewUserUsecase(mockUser)

	result, err := userUsecase.Create("")
	if err != errExpected {
		t.Error("Actual Create() is not same as expected")
	}

	// fmt.Print("エラー", err, "期待", errExpected)
	// fmt.Print("エラー", reflect.TypeOf(err) == reflect.TypeOf(errExpected))
	// fmt.Print(result)

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Actual Create() is not same as expected")
	}
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
