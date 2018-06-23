package main

import (
	"./versionapp"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	// Framework Echo
	app := echo.New()

	// Middleware App
	app.Use(middleware.Logger())
	app.Use(middleware.Recover())

	// Routes
	app.GET("/", versionapp.CurrentVersion)

	// Server
	app.Logger.Fatal(app.Start(":8080"))
}
