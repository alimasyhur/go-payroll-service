package attendance

import (
	"github.com/alimasyhur/go-payroll-service/internal/app/usecase/attendance"
	"github.com/labstack/echo/v4"
)

type AttendanceHandler interface {
	CreateAttendancePeriod(c echo.Context) error
}

type handler struct {
	attendanceUsecase attendance.AttendanceUsecase
}

func NewHandler() *handler {
	return &handler{}
}

func (h *handler) SetAttendanceUsecase(uc attendance.AttendanceUsecase) *handler {
	h.attendanceUsecase = uc
	return h
}

func (h *handler) Validate() AttendanceHandler {
	if h.attendanceUsecase == nil {
		panic("attendance usecase is nil")
	}
	return h
}
