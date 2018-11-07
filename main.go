package main

import (
	"flag"
	"fmt"
	"go-box/api"
	"go-box/common"
	"go-box/core"
	"go-box/db"
	"go-box/file"
	"go-box/micro"
	"os"

	"github.com/labstack/echo"
)

const (
	version = "0.1"
)

var (
	// Build information should only be set by -ldflags.
	BuildDate    string
	BuildGitHash string

	h       bool
	v       bool
	cfgpath string
	e       *echo.Echo
)

func cmd() {
	flag.BoolVar(&h, "h", false, "this help")
	flag.BoolVar(&v, "v", false, "show version")
	flag.StringVar(&cfgpath, "c", "config.json", "set cfg and start")
	flag.Parse()

	if v {
		fmt.Println("Version:", version)
		if len(BuildDate) > 0 {
			fmt.Println("BuildDate:", BuildDate)
			fmt.Println("BuildGitHash:", BuildGitHash)
		}
		os.Exit(0)
	}

	if h {
		flag.Usage()
		os.Exit(0)
	}
}

func main() {
	// cmd
	cmd()

	// config
	common.InitConfig(cfgpath)

	// db
	db.InitDB()

	// file
	file.InitFile()

	// core
	core.InitCore()

	// api
	api.InitApi()

	// ws
	common.InitWS()

	// grpc
	config := common.GlobalConfig
	if config.GrpcEnable {
		micro.RunGrpc(config.GrpcAddr)
	}

	// this will block forever
	api.StartAPP()
}
