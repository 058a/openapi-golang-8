package locations

import (
	"database/sql"
	"fmt"
	oapicodegen "openapi/internal/infra/oapicodegen/stock/location"
	infra "openapi/internal/infra/repository/sqlboiler/stock/location"

	domain "openapi/internal/domain/stock/location"

	"github.com/labstack/echo/v4"
)

type Handler struct {
	oapicodegen.ServerInterface
	Repository domain.IRepository
}

func RegisterHandlers(e *echo.Echo, db *sql.DB) error {
	if db == nil {
		return fmt.Errorf("db is nil")
	}
	repo, err := infra.NewRepository(db)
	if err != nil {
		return err
	}

	oapicodegen.RegisterHandlers(e, &Handler{Repository: repo})
	return nil
}
