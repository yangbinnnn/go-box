package api

import (
	"go-box/common"
	"go-box/core"
	"net/http"

	"github.com/labstack/echo"
)

func userRegister(c echo.Context) error {
	u := &common.User{}
	if err := c.Bind(u); err != nil {
		return err
	}

	if u.Email == "" || u.Name == "" {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"errmsg": "email or name missing",
		})
	}

	if err := core.UserRegister(u.Email, u.Name); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"errmsg": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"state": "OK",
	})
}

func userInfo(c echo.Context) error {
	email := c.QueryParam("email")
	if email == "" {
		return c.JSON(http.StatusBadRequest, nil)
	}

	u, err := core.UserInfo(email)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"errmsg": err.Error(),
		})
	}
	return c.JSON(http.StatusOK, u)
}
