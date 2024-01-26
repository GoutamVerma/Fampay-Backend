package main

import (
	handlers "github.com/GoutamVerma/FamPay-Backend/cmd/delivery/handlers"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	// Creating a new instance of the Echo framework
	e := echo.New()

	// Adding a logger and revcover middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Registering the handlers for different routes
	handlers.RegisterHandlers(e)

	// Starting the server on port 1323 and logging any fatal errors
	e.Logger.Fatal(e.Start(":1323"))
}
