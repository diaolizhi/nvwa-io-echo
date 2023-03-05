package dashboard

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"nvwa-io/base"
	"nvwa-io/database/models"
)

func GetAppDetail(c echo.Context) error {
	app := models.App{}
	err := base.DB.Where("id = ?", 1).First(&app).Error
	if err != nil {
		return c.String(http.StatusNotFound, "not fund")
	}

	return c.JSON(http.StatusOK, &app)
}
