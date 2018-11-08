package api

import (
	"go-box/common"
	"go-box/core"
	"net/http"
	"time"

	"github.com/labstack/echo"
)

func userRegister(c echo.Context) error {
	u := &common.User{}
	if err := c.Bind(u); err != nil {
		return err
	}

	if u.Email == "" || u.Name == "" {
		return BadRequest(c, "email or name missing")
	}

	if err := core.UserRegister(u.Email, u.Name); err != nil {
		return BadRequest(c, err.Error())
	}

	return OKRequest(c)
}

func userInfo(c echo.Context) error {
	email := c.Get("uid").(string)
	if email == "" {
		return BadRequest(c, "user not login")
	}

	u, err := core.UserInfo(email)
	if err != nil {
		return BadRequest(c, err.Error())
	}
	return OKRequestWith(c, u)
}

func userLogin(c echo.Context) error {
	u := &common.User{}
	if err := c.Bind(u); err != nil {
		return err
	}

	if u.Email == "" || u.Name == "" {
		return BadRequest(c, "email or name missing")
	}

	token, err := core.UserLogin(u.Email)
	if err != nil {
		return err
	}
	cookie := &http.Cookie{
		Name:    "auth-token",
		Value:   token,
		Expires: time.Now().Add(time.Duration(60*60) * time.Second),
	}
	c.SetCookie(cookie)
	return OKRequest(c)
}
