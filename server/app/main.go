package main

import (
	"github.com/h-hiroki/surveillance_gogo/server/app/handlers"
	"github.com/labstack/echo"
)

func main() {
	e := echo.New()

	// テスト用にJSONを返却する　削除予定
	e.File("/", "../../front/index.html")
	e.GET("/test", handlers.GetTest)


	e.Logger.Fatal(e.Start(":1323"))
}