package locations_test

import (
	"database/sql"
	"openapi/internal/infra/database"
	"openapi/internal/ui/stock/locations"
	"testing"

	"github.com/labstack/echo/v4"
)

func TestRegisterHandlers(t *testing.T) {
	e := echo.New()

	db, err := database.Open()
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()

	if err := locations.RegisterHandlers(e, db); err != nil {
		t.Fatal(err)
	}
}

func TestRegisterHandlersFail(t *testing.T) {
	e := echo.New()
	db := &sql.DB{}

	// When
	if err := locations.RegisterHandlers(e, nil); err == nil {
		t.Fatal(err)
	}

	if err := locations.RegisterHandlers(nil, db); err == nil {
		t.Fatal(err)
	}
}
