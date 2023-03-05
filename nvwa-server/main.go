package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"nvwa-io/router"
)

func main() {
	// 初始化配置
	initConfig()
	// 初始化数据库
	initDB()

	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	router.Init(e)

	// Start server
	e.Logger.Fatal(e.Start(":9000"))
}
