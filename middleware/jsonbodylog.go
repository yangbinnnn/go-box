package middleware

import (
	"strings"

	"github.com/labstack/echo"
	"github.com/labstack/gommon/log"
)

func JsonBodyHandler(c echo.Context, req []byte, res []byte) {
	reqStr := "NotJsonContent"
	resStr := "NotJsonContent"
	if isJsonReq(c) {
		reqStr = string(req)
	}
	if isJsonRes(c) {
		resStr = string(res)
	}
	data := map[string]interface{}{
		"url": c.Request().URL,
		"req": reqStr,
		"res": resStr,
	}
	log.Infoj(data)
}

func JsonBodySkipper(c echo.Context) bool {
	if isJsonReq(c) || isJsonRes(c) {
		return false
	}
	return true
}

func isJsonReq(c echo.Context) bool {
	reqCt := strings.ToLower(c.Request().Header.Get("Content-Type"))
	return strings.HasPrefix(reqCt, "application/json")
}

func isJsonRes(c echo.Context) bool {
	resCt := strings.ToLower(c.Response().Header().Get("Content-Type"))
	return strings.HasPrefix(resCt, "application/json")
}
