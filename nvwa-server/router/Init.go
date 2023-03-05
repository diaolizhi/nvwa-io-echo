package router

import (
	"github.com/labstack/echo/v4"
	"nvwa-io/controllers"
)

func Init(e *echo.Echo) {
	e.GET("/", index.Hello)
}
