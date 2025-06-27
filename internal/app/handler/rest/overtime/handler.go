package overtime

import (
	"github.com/alimasyhur/go-payroll-service/internal/app/usecase/overtime"
	"github.com/labstack/echo/v4"
)

type OvertimeHandler interface {
	CreateOvertime(c echo.Context) error
}

type handler struct {
	overtimeUsecase overtime.OvertimeUsecase
}

func NewHandler() *handler {
	return &handler{}
}

func (h *handler) SetOvertimeUsecase(uc overtime.OvertimeUsecase) *handler {
	h.overtimeUsecase = uc
	return h
}

func (h *handler) Validate() OvertimeHandler {
	if h.overtimeUsecase == nil {
		panic("overtime usecase is nil")
	}
	return h
}
