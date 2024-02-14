package hello

import (
	oapicodegen "openapi/internal/infra/oapicodegen/hello"

	"github.com/labstack/echo/v4"
)

type Handler struct {
	oapicodegen.ServerInterface
}

func RegisterHandlers(e *echo.Echo) {
	oapicodegen.RegisterHandlers(e, &Handler{})
}
