package auth

import (
	"net/http"

	"github.com/alimasyhur/go-payroll-service/internal/app/usecase/user"
	"github.com/alimasyhur/go-payroll-service/internal/pkg/apperror"
	"github.com/alimasyhur/go-payroll-service/internal/pkg/response"
	"github.com/alimasyhur/go-payroll-service/internal/pkg/validator"
	"github.com/labstack/echo/v4"
)

func (h *handler) Login(c echo.Context) (err error) {
	var req user.LoginRequest
	if reqErr := validator.Validate(c, &req); reqErr != nil {
		err = apperror.New(http.StatusBadRequest, reqErr)
		return
	}

	resp, err := h.userUsecase.Login(c.Request().Context(), req)
	if err != nil {
		return
	}

	return response.ResponseSuccess(c, resp.Token)
}
