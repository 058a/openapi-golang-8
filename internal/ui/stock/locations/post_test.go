package locations_test

import (
	"openapi/internal/infra/database"
	oapicodegen "openapi/internal/infra/oapicodegen/stock/location"
	"strings"
	"testing"

	"github.com/labstack/echo/v4"
	_ "github.com/lib/pq"

	"github.com/google/uuid"

	"net/http"
)

func TestPostCreated(t *testing.T) {
	t.Parallel()

	// Setup
	db, err := database.Open()
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()

	h, err := NewHandler(t, db)
	if err != nil {
		t.Fatal(err)
	}

	e := NewEchoHelper(t)

	// When
	postReqBody := &oapicodegen.PostStockLocationJSONRequestBody{
		Name: "test",
	}
	ctx := e.Post(postReqBody)
	err = h.PostStockLocation(ctx)

	// Then
	if err != nil {
		t.Fatal(err)
	}
	defer e.rec.Result().Body.Close()

	if e.rec.Code != http.StatusCreated {
		t.Errorf("%T %d want %d", e.rec.Code, e.rec.Code, http.StatusCreated)
	}

	postResBody, err := Response[oapicodegen.Created](e.rec.Result())
	if err != nil {
		t.Fatal(err)
	}

	if postResBody.Id == uuid.Nil {
		t.Errorf("expected not empty, actual empty")
	}
}

func TestPostBadRequestNameEmpty(t *testing.T) {
	t.Parallel()

	// Setup
	db, err := database.Open()
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()

	h, err := NewHandler(t, db)
	if err != nil {
		t.Fatal(err)
	}

	e := NewEchoHelper(t)

	// When
	postReqBody := &oapicodegen.PostStockLocationJSONRequestBody{
		Name: "",
	}
	ctx := e.Post(postReqBody)
	err = h.PostStockLocation(ctx)

	// Then
	if err == nil {
		t.Fatalf("expected not nil, actual nil")
	}

	if err.(*echo.HTTPError).Code != http.StatusBadRequest {
		t.Errorf("%T %d want %d", err.(*echo.HTTPError).Code, err.(*echo.HTTPError).Code, http.StatusBadRequest)
	}
}

func TestPostBadRequestNameMaxLengthOver(t *testing.T) {
	t.Parallel()

	// Setup
	db, err := database.Open()
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()

	h, err := NewHandler(t, db)
	if err != nil {
		t.Fatal(err)
	}

	e := NewEchoHelper(t)

	// When
	postReqBody := &oapicodegen.PostStockLocationJSONRequestBody{
		Name: strings.Repeat("a", 101),
	}
	ctx := e.Post(postReqBody)
	err = h.PostStockLocation(ctx)

	// Then
	if err == nil {
		t.Fatalf("expected not nil, actual nil")
	}

	if err.(*echo.HTTPError).Code != http.StatusBadRequest {
		t.Errorf("%T %d want %d", err.(*echo.HTTPError).Code, err.(*echo.HTTPError).Code, http.StatusBadRequest)
	}
}
