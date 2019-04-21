package main

import (
	"github.com/h-hiroki/surveillance_gogo/server/app/handlers"
	"github.com/labstack/echo"
)

func main() {
	e := echo.New()

	// テスト用にJSONを返却する　削除予定
	//e.File("/", "../../front/build/index.html")
	e.GET("/health_check", handlers.HealthCheck)
	e.GET("/test", handlers.GetTest)

	e.Logger.Fatal(e.Start(":1323"))
}