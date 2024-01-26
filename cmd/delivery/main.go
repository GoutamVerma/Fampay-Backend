package main

import (
	handlers "github.com/GoutamVerma/FamPay-Backend/cmd/delivery/handlers"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	handlers.RegisterHandlers(e)

	e.Logger.Fatal(e.Start(":1323"))
}
