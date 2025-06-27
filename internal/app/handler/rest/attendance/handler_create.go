package attendance

import (
	"net/http"

	"github.com/alimasyhur/go-payroll-service/internal/app/usecase/attendance"
	"github.com/alimasyhur/go-payroll-service/internal/pkg/apperror"
	"github.com/alimasyhur/go-payroll-service/internal/pkg/response"
	"github.com/alimasyhur/go-payroll-service/internal/pkg/validator"
	"github.com/labstack/echo/v4"
)

func (h *handler) CreateAttendance(c echo.Context) (err error) {
	userUUID := c.Get("user_uuid").(string)
	ip := c.RealIP()

	var req attendance.AttendanceRequest
	req.UserUUID = userUUID
	req.IP = ip

	if reqErr := validator.Validate(c, &req); reqErr != nil {
		err = apperror.New(http.StatusBadRequest, reqErr)
		return
	}

	resp, err := h.attendanceUsecase.CreateAttendance(c.Request().Context(), req)
	if err != nil {
		return
	}

	return response.ResponseSuccess(c, resp)
}
