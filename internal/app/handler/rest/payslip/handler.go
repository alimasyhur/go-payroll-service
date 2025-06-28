package payslip

import (
	"github.com/alimasyhur/go-payroll-service/internal/app/usecase/payslip"
	"github.com/labstack/echo/v4"
)

type PayslipHandler interface {
	GetOnePayslip(c echo.Context) error
	GetSummary(c echo.Context) error
}

type handler struct {
	payslipUsecase payslip.PayslipUsecase
}

func NewHandler() *handler {
	return &handler{}
}

func (h *handler) SetPayslipUsecase(uc payslip.PayslipUsecase) *handler {
	h.payslipUsecase = uc
	return h
}

func (h *handler) Validate() PayslipHandler {
	if h.payslipUsecase == nil {
		panic("payslip usecase is nil")
	}
	return h
}
