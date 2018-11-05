package db

import (
	"go-box/common"
	"time"

	"github.com/globalsign/mgo"
)

var (
	UserTable *mgo.Collection
)

func GetUser(email string) (*common.User, error) {
	u := new(common.User)

	err := UserTable.FindId(email).One(u)
	if err != nil {
		return nil, err
	}

	return u, nil
}

func AddUser(email, name string) error {
	u := &common.User{Email: email, Name: name, RegisterTime: time.Now()}
	return UserTable.Insert(u)
}
