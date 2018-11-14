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

	core.UserRegister(u.Email, u.Name)
	return OKRequest(c)
}

func userInfo(c echo.Context) error {
	email := c.Get("uid").(string)
	if email == "" {
		return BadRequest(c, "user not login")
	}

	u := core.UserInfo(email)
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

	token := core.UserLogin(u.Email)
	cookie := &http.Cookie{
		Name:    "auth-token",
		Value:   token,
		Expires: time.Now().Add(time.Duration(60*60) * time.Second),
	}
	c.SetCookie(cookie)
	return OKRequest(c)
}
