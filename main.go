package main

import (
	"github.com/labstack/echo"
)

func main() {
	e := echo.New()

	e.Static("/", "views")
	e.Logger.Fatal(e.Start(":1323"))
}