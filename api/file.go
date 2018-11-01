package api

import (
	"go-box/file"
	"net/http"
	"os"

	"github.com/labstack/echo"
)

func upload(c echo.Context) error {
	name := c.Param("name")

	_, err := file.FileStore.Put(name, c.Request().Body)
	if err != nil {
		return err
	}

	return c.String(http.StatusOK, "")
}

func download(c echo.Context) error {
	name := c.Param("name")
	c.Response().Header().Set(echo.HeaderContentType, echo.MIMEOctetStream)
	c.Response().WriteHeader(http.StatusOK)
	_, err := file.FileStore.Get(name, c.Response().Writer)
	if err != nil {
		if os.IsNotExist(err) {
			return c.String(http.StatusNotFound, "")
		}
		return err
	}
	return nil
}
