package attendance

import (
	"net/http"

	"github.com/alimasyhur/go-payroll-service/internal/app/usecase/attendance"
	"github.com/alimasyhur/go-payroll-service/internal/pkg/apperror"
	"github.com/alimasyhur/go-payroll-service/internal/pkg/response"
	"github.com/alimasyhur/go-payroll-service/internal/pkg/validator"
	"github.com/labstack/echo/v4"
)

func (h *handler) CreateAttendancePeriod(c echo.Context) (err error) {
	var req attendance.AttendancePeriodRequest
	if reqErr := validator.Validate(c, &req); reqErr != nil {
		err = apperror.New(http.StatusBadRequest, reqErr)
		return
	}

	resp, err := h.attendanceUsecase.CreateAttendancePeriod(c.Request().Context(), req)
	if err != nil {
		return
	}

	return response.ResponseSuccess(c, resp)
}
