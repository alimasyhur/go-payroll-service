package payroll

import (
	"github.com/alimasyhur/go-payroll-service/internal/app/usecase/payroll"
	"github.com/labstack/echo/v4"
)

type PayrollHandler interface {
	CreatePayroll(c echo.Context) error
}

type handler struct {
	payrollUsecase payroll.PayrollUsecase
}

func NewHandler() *handler {
	return &handler{}
}

func (h *handler) SetPayrollUsecase(uc payroll.PayrollUsecase) *handler {
	h.payrollUsecase = uc
	return h
}

func (h *handler) Validate() PayrollHandler {
	if h.payrollUsecase == nil {
		panic("payroll usecase is nil")
	}
	return h
}
