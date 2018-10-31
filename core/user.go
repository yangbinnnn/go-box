package core

import (
	"errors"
	"go-box/common"
	"time"
)

var (
	ErrUserExists   = errors.New("User already exsists")
	ErrUserNotFound = errors.New("User not found")
)

func UserRegister(email, name string) error {
	n, _ := UserTable.FindId(email).Count()
	if n != 0 {
		return ErrUserExists
	}

	u := &common.User{Email: email, Name: name, RegisterTime: time.Now()}
	return UserTable.Insert(u)
}

func UserInfo(email string) (*common.User, error) {
	u := new(common.User)

	err := UserTable.FindId(email).One(u)
	if err != nil {
		return nil, ErrUserNotFound
	}

	return u, nil
}
