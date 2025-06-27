package rest

import (
	"net/http"
	"runtime"

	"github.com/alimasyhur/go-payroll-service/internal/pkg/apperror"
	"github.com/alimasyhur/go-payroll-service/internal/pkg/response"
	"github.com/labstack/echo/v4"
)

func errorHandler(err error, c echo.Context) {
	// Need this, because somehow if both of error handler and body dump middleware used
	// It will be printed error twice

	if c.Get("error-handled") != nil {
		return
	}

	c.Set("error-handled", true)

	status := http.StatusBadRequest
	resp := response.DefaultResponse{
		Success: false,
		Message: err.Error(),
	}

	if ae, ok := err.(*apperror.ApplicationError); ok {
		status = ae.Status
		resp.Message = ae.Message
	} else if ae, ok := err.(*echo.HTTPError); ok {
		status = ae.Code
	} else if _, ok := err.(runtime.Error); ok {
		status = http.StatusInternalServerError
	}

	err = c.JSON(status, resp)
}
