package api

import (
	"go-box/file"
	"net/http"
	"os"

	"github.com/labstack/echo"
)

/**
 * @api {PUT} /api/file/:name upload file
 * @apiGroup file
 * @apiName upload
 * @apiParam {String} name file name
 * @apiExample {curl} Example usage:
 *     curl -i -X PUT -d 'hello world!' http://127.0.0.1:8000/api/file/hello.txt
 */
func upload(c echo.Context) error {
	name := c.Param("name")

	_, err := file.FileStore.Put(name, c.Request().Body)
	if err != nil {
		return err
	}

	return c.String(http.StatusOK, "")
}

/**
 * @api {GET} /api/file/:name download file
 * @apiGroup file
 * @apiName download
 * @apiParam {String} name file name
 * @apiExample {curl} Example usage:
 *     curl -i http://127.0.0.1:8000/api/file/hello.txt
 */
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
