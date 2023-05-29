package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"api/pkg/http/api"
)

func main() {
	r := echo.New()

	r.Use(middleware.Logger())
	r.Use(middleware.Recover())

	api.InjectApiRoutes(r)

	r.Logger.Fatal(r.Start(":8080"))
}
