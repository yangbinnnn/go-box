package core

import (
	"errors"
	"go-box/common"
	"go-box/db"
)

var (
	ErrUserExists   = errors.New("User already exsists")
	ErrUserNotFound = errors.New("User not found")
)

func UserRegister(email, name string) error {
	u, _ := UserInfo(email)
	if u != nil {
		return ErrUserExists
	}

	return db.AddUser(email, name)
}

func UserInfo(email string) (*common.User, error) {
	u, err := db.GetUser(email)
	if err != nil {
		return nil, ErrUserNotFound
	}

	return u, nil
}
