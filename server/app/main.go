package main

import (
	"github.com/h-hiroki/surveillance_gogo/server/app/handlers"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	e := echo.New()

	// クロスドメイン対策。これ以外の方法で対応したい。。。
	e.Use(middleware.CORS())

	// テスト用にJSONを返却する　削除予定
	e.GET("/health_check", handlers.HealthCheck)
	e.GET("/test", handlers.GetTest)

	e.Logger.Fatal(e.Start(":1323"))
}