package subscriber

import (
	"context"

	"github.com/weanan/weanan-service/internal/app/container"
	"github.com/weanan/weanan-service/internal/pkg/logger"
)

func LoggerMiddleware(ctx context.Context, container *container.Container, topic string, req interface{}) (newCtx context.Context) {
	cfg := container.Config

	ctxLogger := logger.Context{
		ServiceName:    cfg.App.Name,
		ServiceVersion: cfg.App.Version,
		ServicePort:    cfg.App.HttpPort,
		ReqMethod:      "Subscriber",
		ReqURI:         topic,
	}

	newCtx = logger.InjectCtx(ctx, ctxLogger)

	logger.Log.Info(newCtx, "Request Body", req)

	return
}
