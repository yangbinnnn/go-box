package api

import (
	"net/http"
	"reflect"
	"strings"

	"github.com/labstack/echo"
)

func OKRequest(c echo.Context) error {
	return c.JSON(http.StatusOK, nil)
}

func OKRequestWith(c echo.Context, o interface{}) error {
	m := map[string]interface{}{
		"stat": "OK",
	}
	on := reflect.Indirect(reflect.ValueOf(o)).Type().Name()
	if len(on) <= 1 {
		on = strings.ToLower(on)
	} else {
		on = strings.ToLower(string(on[0])) + on[1:]
	}
	m[on] = o
	return c.JSON(http.StatusOK, m)
}

func BadRequest(c echo.Context, message string) error {
	m := map[string]interface{}{
		"stat":    "ERROR",
		"message": message,
	}
	return c.JSON(http.StatusBadRequest, m)
}
