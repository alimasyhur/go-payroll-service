package payslip

import (
	"net/http"

	"github.com/alimasyhur/go-payroll-service/internal/app/usecase/payslip"
	"github.com/alimasyhur/go-payroll-service/internal/pkg/apperror"
	"github.com/alimasyhur/go-payroll-service/internal/pkg/response"
	"github.com/alimasyhur/go-payroll-service/internal/pkg/validator"
	"github.com/labstack/echo/v4"
)

func (h *handler) GetSummary(c echo.Context) (err error) {
	userUUID := c.Get("user_uuid").(string)

	var req payslip.GetSummaryRequest
	req.UserUUID = userUUID
	req.PayrollUUID = c.Param("payroll_uuid")

	if reqErr := validator.Validate(c, &req); reqErr != nil {
		err = apperror.New(http.StatusBadRequest, reqErr)
		return
	}

	resp, err := h.payslipUsecase.GetSummary(c.Request().Context(), req)
	if err != nil {
		return
	}

	return response.ResponseSuccess(c, resp)
}
