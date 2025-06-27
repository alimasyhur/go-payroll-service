package mocks

import (
	"bytes"
	"net/http"
	"net/http/httptest"

	"github.com/labstack/echo/v4"

	"github.com/alimasyhur/go-payroll-service/internal/app/server/rest"
	"github.com/alimasyhur/go-payroll-service/internal/pkg/validator"
)

func MockEcho(method, path string, headers http.Header, body []byte) (c echo.Context, rec *httptest.ResponseRecorder) {
	e := echo.New()

	e.Validator = &rest.DataValidator{ValidatorData: validator.SetupValidator()}

	rec = httptest.NewRecorder()
	req := httptest.NewRequest(method, path, bytes.NewBuffer(body))
	req.Header = headers

	c = e.NewContext(req, rec)
	return
}
