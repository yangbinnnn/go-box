package middleware

import (
	"fmt"
	"go-box/db"
	"net/http"

	"github.com/labstack/echo"
)

type tokenAuth struct {
	TokenName string
}

func (ta *tokenAuth) Process(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		cookie, err := c.Request().Cookie(ta.TokenName)
		if err != nil {
			return unAuthRequest(c, fmt.Sprintf("Auth cookie %s not found", ta.TokenName))
		}

		uid, _ := db.GetUserToken(cookie.Value)
		if uid == "" {
			return unAuthRequest(c, "Auth token expired")
		}

		// save valid uid for later use
		c.Set("uid", uid)

		return next(c)
	}
}

func unAuthRequest(c echo.Context, errMsg string) error {
	m := make(map[string]interface{})
	m["stat"] = "ERROR"
	m["message"] = errMsg
	return c.JSON(http.StatusUnauthorized, m)
}
