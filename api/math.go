package api

import (
	"go-box/core"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

func add(c echo.Context) error {
	a, _ := strconv.Atoi(c.QueryParam("a"))
	b, _ := strconv.Atoi(c.QueryParam("b"))
	s := core.Add(a, b)
	return c.JSON(http.StatusOK, map[string]int{
		"a":   a,
		"b":   b,
		"sum": s,
	})
}
