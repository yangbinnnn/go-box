package api

import (
	"go-box/core"
	"net/http"

	"github.com/labstack/echo"
)

func add(c echo.Context) error {
	a := int(QueryParamMustInt(c, "a"))
	b := int(QueryParamMustInt(c, "b"))
	s := core.Add(a, b)
	return c.JSON(http.StatusOK, map[string]int{
		"a":   a,
		"b":   b,
		"sum": s,
	})
}
