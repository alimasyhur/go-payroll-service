package reimbursement

import (
	"net/http"

	"github.com/alimasyhur/go-payroll-service/internal/app/usecase/reimbursement"
	"github.com/alimasyhur/go-payroll-service/internal/pkg/apperror"
	"github.com/alimasyhur/go-payroll-service/internal/pkg/response"
	"github.com/alimasyhur/go-payroll-service/internal/pkg/validator"
	"github.com/labstack/echo/v4"
)

func (h *handler) CreateReimbursement(c echo.Context) (err error) {
	userUUID := c.Get("user_uuid").(string)
	ip := c.RealIP()

	var req reimbursement.ReimbursementRequest
	req.UserUUID = userUUID
	req.IP = ip

	if reqErr := validator.Validate(c, &req); reqErr != nil {
		err = apperror.New(http.StatusBadRequest, reqErr)
		return
	}

	resp, err := h.reimbursementUsecase.CreateReimbursement(c.Request().Context(), req)
	if err != nil {
		return
	}

	return response.ResponseSuccess(c, resp)
}
