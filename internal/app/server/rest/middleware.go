package rest

import (
	"bytes"
	"io"
	"net/http"
	"os"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"github.com/alimasyhur/go-payroll-service/internal/app/container"
	"github.com/alimasyhur/go-payroll-service/internal/pkg/logger"
	pkgValidator "github.com/alimasyhur/go-payroll-service/internal/pkg/validator"
)

func SetupMiddleware(server *echo.Echo, container *container.Container) {
	server.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{echo.GET, echo.POST, echo.PUT, echo.DELETE, echo.PATCH, echo.OPTIONS},
		AllowHeaders:     []string{"Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With, X-App-Token, X-Client-Id"},
		ExposeHeaders:    []string{"Content-Length", "Access-Control-Allow-Origin"},
		AllowCredentials: true,
	}))

	server.Use(middleware.Recover())
	server.Use(SetLoggerMiddleware(container))
	server.Use(LoggerMiddleware(container))

	server.HTTPErrorHandler = errorHandler
	server.Validator = &DataValidator{ValidatorData: pkgValidator.SetupValidator()}
}

func SetLoggerMiddleware(container *container.Container) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			if SkipLoggerMiddleware(c.Path()) {
				return next(c)
			}

			cfg := container.Config
			ctxLogger := logger.Context{
				ServiceName:    cfg.App.Name,
				ServiceVersion: cfg.App.Version,
				ServicePort:    cfg.App.HttpPort,
				ReqMethod:      c.Request().Method,
				ReqURI:         c.Request().URL.String(),
			}

			var bodyByte []byte
			if c.Request().Body != nil {
				bodyByte, _ = io.ReadAll(c.Request().Body)
				ctxLogger.ReqBody = string(bodyByte)

				c.Request().Body = io.NopCloser(bytes.NewBuffer(bodyByte))
			}

			request := c.Request()

			ctx := logger.InjectCtx(request.Context(), ctxLogger)
			c.SetRequest(request.WithContext(ctx))

			logger.Log.Info(ctx, "Request Header", c.Request().Header)

			if !logger.IsSkipLog(c.Request().Header.Get("Content-Type")) {
				logger.Log.Info(ctx, "Request Body", string(bodyByte))
			} else {
				logger.Log.Info(ctx, "Request Not Log Because Unsupported Content-Type")
			}

			return next(c)
		}
	}
}

func LoggerMiddleware(container *container.Container) echo.MiddlewareFunc {
	return middleware.BodyDump(func(c echo.Context, reqBody, resBody []byte) {
		if SkipLoggerMiddleware(c.Path()) {
			return
		}

		// log request header, body & response
		ctx := c.Request().Context()

		if !logger.IsSkipLog(c.Response().Header().Get("Content-Type")) {
			logger.Log.Info(ctx, "Response Body", string(resBody))
		} else {
			logger.Log.Info(ctx, "Response Not Log Because Unsupported Content-Type")
		}
	})
}

func AdminOnlyMiddleware() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			role := c.Get("role")
			if role == nil || role.(string) != "admin" {
				return c.JSON(http.StatusForbidden, echo.Map{
					"error": "admin access only",
				})
			}
			return next(c)
		}
	}
}

func JWTAuthMiddleware() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			authHeader := c.Request().Header.Get("Authorization")
			if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
				return echo.NewHTTPError(http.StatusUnauthorized, "missing or invalid token")
			}

			tokenStr := strings.TrimPrefix(authHeader, "Bearer ")
			token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
				return []byte(os.Getenv("JWT_SECRET")), nil
			})

			if err != nil || !token.Valid {
				return echo.NewHTTPError(http.StatusUnauthorized, "invalid token")
			}

			claims := token.Claims.(jwt.MapClaims)
			userUUID := claims["user_uuid"].(string)
			role := claims["role"].(string)

			c.Set("user_uuid", userUUID)
			c.Set("role", role)

			return next(c)
		}
	}
}

type DataValidator struct {
	ValidatorData *validator.Validate
}

func (cv *DataValidator) Validate(i interface{}) error {
	return cv.ValidatorData.Struct(i)
}

func SkipLoggerMiddleware(path string) bool {
	switch path {
	case "/", "/metrics", "/favicon.ico":
		return true
	}

	return false
}
