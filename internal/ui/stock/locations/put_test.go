package locations_test

import (
	"openapi/internal/infra/database"
	oapicodegen "openapi/internal/infra/oapicodegen/stock/location"
	"strings"
	"testing"

	"github.com/labstack/echo/v4"

	"github.com/google/uuid"

	"net/http"
)

func TestPutOk(t *testing.T) {
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

	postResBody, err := Response[oapicodegen.Created](e.rec.Result())
	if err != nil {
		t.Fatal(err)
	}

	// When
	putReqBody := &oapicodegen.PutStockLocationJSONRequestBody{
		Name: "test",
	}
	ctx = e.Put(postResBody.Id, putReqBody)
	err = h.PutStockLocation(ctx, postResBody.Id)

	// Then
	if err != nil {
		t.Fatal(err)
	}

	if e.rec.Code != http.StatusOK {
		t.Errorf("%T %d want %d", e.rec.Code, e.rec.Code, http.StatusOK)
	}
}

func TestPutNotFound(t *testing.T) {
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
	putReqBody := &oapicodegen.PutStockLocationJSONRequestBody{
		Name: "test",
	}
	id := uuid.New()
	ctx := e.Put(id, putReqBody)
	err = h.PutStockLocation(ctx, id)

	// Then
	if err == nil {
		t.Fatalf("expected not nil, actual nil")
	} else if err.(*echo.HTTPError).Code != http.StatusNotFound {
		t.Errorf("%T %d want %d", err.(*echo.HTTPError).Code, err.(*echo.HTTPError).Code, http.StatusNotFound)
	}
}

func TestPutBadRequestNameEmpty(t *testing.T) {
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

	postResBody, err := Response[oapicodegen.Created](e.rec.Result())
	if err != nil {
		t.Fatal(err)
	}

	// When
	putReqBody := &oapicodegen.PutStockLocationJSONRequestBody{
		Name: "",
	}
	ctx = e.Put(postResBody.Id, putReqBody)
	err = h.PutStockLocation(ctx, postResBody.Id)

	// Then
	if err == nil {
		t.Fatalf("expected not nil, actual nil")
	} else if err.(*echo.HTTPError).Code != http.StatusBadRequest {
		t.Errorf("%T %d want %d", err.(*echo.HTTPError).Code, err.(*echo.HTTPError).Code, http.StatusBadRequest)
	}
}

func TestPutBadRequestNameMaxLengthOver(t *testing.T) {
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

	postResBody, err := Response[oapicodegen.Created](e.rec.Result())
	if err != nil {
		t.Fatal(err)
	}

	// When
	putReqBody := &oapicodegen.PutStockLocationJSONRequestBody{
		Name: strings.Repeat("a", 101),
	}
	ctx = e.Put(postResBody.Id, putReqBody)
	err = h.PutStockLocation(ctx, postResBody.Id)

	// Then
	if err == nil {
		t.Fatalf("expected not nil, actual nil")
	} else if err.(*echo.HTTPError).Code != http.StatusBadRequest {
		t.Errorf("%T %d want %d", err.(*echo.HTTPError).Code, err.(*echo.HTTPError).Code, http.StatusBadRequest)
	}
}
