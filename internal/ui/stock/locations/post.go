package locations

import (
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"

	app "openapi/internal/app/stock/location"
	oapicodegen "openapi/internal/infra/oapicodegen/stock/location"
)

// PostStockLocation is a function that handles the HTTP POST request for creating a new stock item.
func (h *Handler) PostStockLocation(ctx echo.Context) error {
	// Binding
	req := &oapicodegen.PostStockLocationJSONRequestBody{}
	if err := ctx.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	// Validation
	if err := ctx.Validate(req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	reqDto, err := app.NewCreateRequest(req.Name)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	// Main Process
	newId := uuid.New()
	resDto, err := app.Create(reqDto, h.Repository, newId)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	res := &oapicodegen.Created{Id: resDto.Id}

	// Postcondition
	if err := ctx.Validate(res); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return ctx.JSON(http.StatusCreated, res)
}
