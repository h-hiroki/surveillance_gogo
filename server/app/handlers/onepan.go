package handlers

import (
	"github.com/h-hiroki/surveillance_gogo/server/app/crawler"
	"github.com/labstack/echo"
	"net/http"
)

func OnePan(context echo.Context) error {
	crawler.CheckDiff()
	return context.JSON(http.StatusOK, "onepan ok")
}
