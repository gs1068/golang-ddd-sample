package model

import "errors"

type User struct {
	ID       int
	UserName string
}

func NewUser(username string) (*User, error) {
	if username == "" {
		return nil, errors.New("ユーザーネームを入力してください")
	}

	user := &User{
		UserName: username,
	}
	return user, nil
}

func (u *User) Set(username string) error {
	if username == "" {
		return errors.New("ユーザーネームを入力してください")
	}

	u.UserName = username
	return nil
}
