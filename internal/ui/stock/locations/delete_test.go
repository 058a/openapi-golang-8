package locations_test

import (
	"openapi/internal/infra/database"
	oapicodegen "openapi/internal/infra/oapicodegen/stock/location"
	"testing"

	"github.com/labstack/echo/v4"
	_ "github.com/lib/pq"

	"github.com/google/uuid"

	"net/http"
)

func TestDeleteOk(t *testing.T) {
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

	// Given
	postReqBody := &oapicodegen.PostStockLocationJSONRequestBody{
		Name: "test",
	}
	ctx := e.Post(postReqBody)
	if err := h.PostStockLocation(ctx); err != nil {
		t.Fatal(err)
	}
	defer e.rec.Result().Body.Close()

	if e.rec.Code != http.StatusCreated {
		t.Fatalf("%T %d want %d", e.rec.Code, e.rec.Code, http.StatusCreated)
	}

	postResBody, err := Response[oapicodegen.Created](e.rec.Result())
	if err != nil {
		t.Fatal(err)
	}

	// When
	ctx = e.Delete(postResBody.Id)
	if err := h.DeleteStockLocation(ctx, postResBody.Id); err != nil {
		t.Fatal(err)
	}
	defer e.rec.Result().Body.Close()

	// Then
	if e.rec.Code != http.StatusOK {
		t.Errorf("want %d, got %d", http.StatusOK, e.rec.Code)
	}
}

func TestDeleteNotFound(t *testing.T) {
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
	id := uuid.New()
	ctx := e.Delete(id)
	err = h.DeleteStockLocation(ctx, id)

	// Then
	if err == nil {
		t.Fatalf("expected not nil, actual nil")
	} else if err.(*echo.HTTPError).Code != http.StatusNotFound {
		t.Errorf("%T %d want %d", err.(*echo.HTTPError).Code, err.(*echo.HTTPError).Code, http.StatusNotFound)
	}
}
