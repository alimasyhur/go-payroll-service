package auth

import (
	"context"

	"github.com/alimasyhur/go-payroll-service/internal/app/usecase/user"
	"github.com/labstack/echo/v4"
)

type AuthHandler interface {
	Login(c echo.Context) error
	GenerateUserSeed(c context.Context) error
}

type handler struct {
	userUsecase user.UserUsecase
}

func NewHandler() *handler {
	return &handler{}
}

func (h *handler) SetAuthUsecase(uc user.UserUsecase) *handler {
	h.userUsecase = uc
	return h
}

func (h *handler) Validate() AuthHandler {
	if h.userUsecase == nil {
		panic("user usecase is nil")
	}
	return h
}
