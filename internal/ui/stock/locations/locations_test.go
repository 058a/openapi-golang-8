package locations_test

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"openapi/internal/infra/database"
	"openapi/internal/infra/env"
	oapicodegen "openapi/internal/infra/oapicodegen/stock/location"
	infra "openapi/internal/infra/repository/sqlboiler/stock/location"
	"openapi/internal/infra/validator"
	"openapi/internal/ui/stock/locations"
	"testing"

	"github.com/google/uuid"
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

type RequestHelper struct {
	client *http.Client
}

type Request struct {
	context  echo.Context
	recorder *httptest.ResponseRecorder
}

func NewHandler(t *testing.T, db *sql.DB) (*locations.Handler, error) {
	t.Helper()

	repo, err := infra.NewRepository(db)
	if err != nil {
		return nil, err
	}

	return &locations.Handler{Repository: repo}, nil
}

type EchoHelper struct {
	e   *echo.Echo
	t   *testing.T
	rec *httptest.ResponseRecorder
}

func NewEchoHelper(t *testing.T) *EchoHelper {
	t.Helper()

	e := echo.New()

	e.Validator = validator.NewCustomValidator()

	return &EchoHelper{e: e, t: t}
}

func (h *EchoHelper) Post(reqBody *oapicodegen.PostStockLocationJSONRequestBody) echo.Context {
	h.t.Helper()

	reqBodyJson, _ := json.Marshal(reqBody)
	req := httptest.NewRequest(http.MethodPost, "/stock/locations", bytes.NewBuffer(reqBodyJson))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	h.rec = httptest.NewRecorder()

	return h.e.NewContext(req, h.rec)
}

func (h *EchoHelper) Put(stockLocationsId uuid.UUID, reqBody *oapicodegen.PutStockLocationJSONRequestBody) echo.Context {
	h.t.Helper()

	reqBodyJson, _ := json.Marshal(reqBody)
	req := httptest.NewRequest(http.MethodPut, fmt.Sprintf("/stock/locations/%s", stockLocationsId), bytes.NewBuffer(reqBodyJson))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	h.rec = httptest.NewRecorder()

	return h.e.NewContext(req, h.rec)
}

func (h *EchoHelper) Delete(stockLocationsId uuid.UUID) echo.Context {
	h.t.Helper()

	req := httptest.NewRequest(http.MethodDelete, fmt.Sprintf("/stock/locations/%s", stockLocationsId), nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	h.rec = httptest.NewRecorder()

	return h.e.NewContext(req, h.rec)
}

func NewRequest[I any](method string, path string, reqBody *I) *Request {
	e := echo.New()

	e.Validator = validator.NewCustomValidator()

	reqBodyJson, _ := json.Marshal(reqBody)
	req := httptest.NewRequest(method, path, bytes.NewBuffer(reqBodyJson))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	rec := httptest.NewRecorder()

	ctx := e.NewContext(req, rec)
	return &Request{
		context:  ctx,
		recorder: rec,
	}
}

func Response[T any](res *http.Response) (*T, error) {
	resBodyByte, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	resBody := new(T)
	json.Unmarshal(resBodyByte, resBody)
	return resBody, nil
}

func (h *RequestHelper) Post(reqBody *oapicodegen.PostStockLocationJSONRequestBody) (*http.Response, error) {
	reqBodyJson, _ := json.Marshal(reqBody)
	req, err := http.NewRequest(
		http.MethodPost,
		env.GetServiceUrl()+"/stock/locations",
		bytes.NewBuffer(reqBodyJson),
	)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	res, err := h.client.Do(req)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (h *RequestHelper) Put(stockLocationsId uuid.UUID, reqBody *oapicodegen.PutStockLocationJSONRequestBody) (*http.Response, error) {
	reqBodyJson, _ := json.Marshal(reqBody)
	req, err := http.NewRequest(
		http.MethodPut,
		env.GetServiceUrl()+"/stock/locations/"+stockLocationsId.String(),
		bytes.NewBuffer(reqBodyJson),
	)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	res, err := h.client.Do(req)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (h *RequestHelper) Delete(stockLocationsId uuid.UUID) (*http.Response, error) {
	req, err := http.NewRequest(
		http.MethodDelete,
		env.GetServiceUrl()+"/stock/locations/"+stockLocationsId.String(),
		nil,
	)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	res, err := h.client.Do(req)
	if err != nil {
		return nil, err
	}

	return res, nil
}

type ResponseConvertHelper struct{}

func (h *ResponseConvertHelper) AsCreated(res *http.Response) (*oapicodegen.Created, error) {
	resBodyByte, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	resBody := &oapicodegen.Created{}
	json.Unmarshal(resBodyByte, &resBody)
	if resBody.Id == uuid.Nil {
		return nil, fmt.Errorf("expected not empty, actual empty")
	}
	return resBody, nil
}

func (h *ResponseConvertHelper) AsBadRequest(res *http.Response) (*oapicodegen.BadRequestResponse, error) {
	resBodyByte, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	resBody := &oapicodegen.BadRequest{}
	json.Unmarshal(resBodyByte, &resBody)

	return resBody, nil
}
