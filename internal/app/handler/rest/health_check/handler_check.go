package health_check

import (
	"net/http"

	"github.com/alimasyhur/go-payroll-service/internal/pkg/response"
	"github.com/labstack/echo/v4"
)

func (h *handler) Check(c echo.Context) error {
	resp := response.DefaultResponse{
		Success: true,
		Message: "OK",
	}

	return c.JSON(http.StatusOK, resp)
}
