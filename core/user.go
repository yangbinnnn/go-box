package core

import (
	"go-box/common"
	"go-box/db"

	"github.com/segmentio/ksuid"
)

func UserRegister(email, name string) {
	u, err := db.GetUser(email)
	common.Verify(err != nil && err.Error() != "not found", common.ErrServerException)
	common.Verify(u != nil, common.ErrUserAlreadyExists)
	common.CheckError(db.AddUser(email, name))
}

func UserInfo(email string) *common.User {
	u, err := db.GetUser(email)
	common.Verify(err != nil && err.Error() == "not found", common.ErrUserNotFound)
	common.CheckError(err)
	return u
}

func UserLogin(email string) (token string) {
	token = ksuid.New().String()
	common.CheckError(db.AddUserToken(token, email, 60*60*24))
	return token
}
