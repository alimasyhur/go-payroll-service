package rest

import (
	"github.com/labstack/echo/v4"

	"github.com/alimasyhur/go-payroll-service/internal/app/container"
	"github.com/alimasyhur/go-payroll-service/internal/app/handler/rest/auth"
	"github.com/alimasyhur/go-payroll-service/internal/app/handler/rest/health_check"
)

func SetupRouter(server *echo.Echo, container *container.Container) {
	// inject handler with usecase via container
	healthCheckHandler := health_check.NewHandler().Validate()
	authHandler := auth.NewHandler().
		SetAuthUsecase(container.UserUsecase).
		Validate()

	server.GET("/health-check", healthCheckHandler.Check)
	server.POST("/login", authHandler.Login)
}
