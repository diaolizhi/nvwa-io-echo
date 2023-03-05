package index

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func Hello(c echo.Context) error {
	return c.HTML(http.StatusOK, "<h1>Hello nvwa!</h1>")
}
