package main

import "github.com/labstack/echo/v4"

func main() {
	e := echo.New()
	e.File("/", "index.html")
	e.Static("/", "./")
	e.Logger.Fatal(e.Start(":8081"))
}
