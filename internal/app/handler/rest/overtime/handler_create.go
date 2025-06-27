package overtime

import (
	"net/http"

	"github.com/alimasyhur/go-payroll-service/internal/app/usecase/overtime"
	"github.com/alimasyhur/go-payroll-service/internal/pkg/apperror"
	"github.com/alimasyhur/go-payroll-service/internal/pkg/response"
	"github.com/alimasyhur/go-payroll-service/internal/pkg/validator"
	"github.com/labstack/echo/v4"
)

func (h *handler) CreateOvertime(c echo.Context) (err error) {
	userUUID := c.Get("user_uuid").(string)
	ip := c.RealIP()

	var req overtime.OvertimeRequest
	req.UserUUID = userUUID
	req.IP = ip

	if reqErr := validator.Validate(c, &req); reqErr != nil {
		err = apperror.New(http.StatusBadRequest, reqErr)
		return
	}

	resp, err := h.overtimeUsecase.CreateOvertime(c.Request().Context(), req)
	if err != nil {
		return
	}

	return response.ResponseSuccess(c, resp)
}
