package router

import (
	"github.com/labstack/echo/v4"
	"nvwa-io/controllers"
	"nvwa-io/controllers/dashboard"
)

func Init(e *echo.Echo) {
	e.GET("/", index.Hello)

	// 应用详情
	e.GET("/api/dashboard/app/detail", dashboard.GetAppDetail)
}
