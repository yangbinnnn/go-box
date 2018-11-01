package api

import (
	"net/http"

	"github.com/labstack/echo"
)

/**
 * @api {GET} /ping ping server
 * @apiGroup root
 * @apiName ping
 * @apiDescription return pong
 */
func ping(c echo.Context) error {
	return c.String(http.StatusOK, "pong")
}
