package api

import (
	"io"
	"os"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/labstack/gommon/log"

	"go-box/common"
)

var (
	e *echo.Echo
)

func StartAPP() {
	config := common.GlobalConfig
	e.Logger.Fatal(e.Start(config.HTTPAddr))
}

func InitApi() {
	config := common.GlobalConfig
	logfile, _ := os.OpenFile(config.LogPath, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
	tee := io.MultiWriter(os.Stdout, logfile)
	logconfig := middleware.LoggerConfig{
		Output: tee,
	}

	e = echo.New()
	e.Logger.SetLevel(log.DEBUG)

	// middlewares
	e.Use(middleware.CORS())
	e.Use(middleware.LoggerWithConfig(logconfig))

	// web
	e.File("/", config.WebIndex)
	e.Static("/static", config.WebStatic)

	// api
	srv := e.Group(config.APIPrefix)
	srv.GET("/ping", ping)
	srv.GET("/add", add)
}
