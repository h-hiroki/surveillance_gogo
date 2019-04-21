package handlers

import (
	"github.com/h-hiroki/surveillance_gogo/server/app/models"
	"github.com/labstack/echo"
	"net/http"
)

func GetTest(context echo.Context) error {
	return context.JSON(http.StatusOK, models.GetTest())
}
