package main

import (
	"openapi/internal/infra/database"
	"openapi/internal/infra/validator"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	hello "openapi/internal/ui/hello"
	locations "openapi/internal/ui/stock/locations"
)

func main() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.Validator = validator.NewCustomValidator()

	hello.RegisterHandlers(e)

	db, err := database.Open()
	if err != nil {
		e.Logger.Fatal(err)
	}
	defer db.Close()

	if err := locations.RegisterHandlers(e, db); err != nil {
		e.Logger.Fatal(err)
	}

	e.Logger.Fatal(e.Start(":1323"))
}
