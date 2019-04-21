package main

import (
	"github.com/h-hiroki/surveillance_gogo/server/app/handlers"
	"github.com/labstack/echo"
)

func main() {
	e := echo.New()

	// テスト用にJSONを返却する　削除予定
	e.GET("/test", handlers.GetTest)
	e.File("/", "public/")

	e.Logger.Fatal(e.Start(":1323"))
}