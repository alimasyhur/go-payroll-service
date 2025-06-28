package reimbursement

import (
	"github.com/alimasyhur/go-payroll-service/internal/app/usecase/reimbursement"
	"github.com/labstack/echo/v4"
)

type ReimbursementHandler interface {
	CreateReimbursement(c echo.Context) error
}

type handler struct {
	reimbursementUsecase reimbursement.ReimbursementUsecase
}

func NewHandler() *handler {
	return &handler{}
}

func (h *handler) SetReimbursementUsecase(uc reimbursement.ReimbursementUsecase) *handler {
	h.reimbursementUsecase = uc
	return h
}

func (h *handler) Validate() ReimbursementHandler {
	if h.reimbursementUsecase == nil {
		panic("reimbursement usecase is nil")
	}
	return h
}
