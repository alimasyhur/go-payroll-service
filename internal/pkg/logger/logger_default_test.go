package logger_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/weanan/weanan-service/internal/pkg/logger"
)

func TestNewLogg(t *testing.T) {
	// with active config
	opt := logger.Option{
		IsEnable: true,
	}
	log := logger.NewLogger(opt)
	assert.NotNil(t, log)

	// with inactive config
	opt = logger.Option{
		IsEnable: false,
	}
	log = logger.NewLogger(opt)
	assert.NotNil(t, log)
}

func TestLoggerDebug(t *testing.T) {
	// its not test, because there's no assertation
	opt := logger.Option{
		IsEnable: true,
	}
	log := logger.NewLogger(opt)
	ctx := logger.InjectCtx(context.Background(), logger.Context{
		ServiceName: "logger",
		Tag:         "xxx",
		ReqMethod:   "POST",
		ReqURI:      "/",
	})

	assert.NotPanics(t, func() {
		log.Debug(ctx, "log Debug", "detail log Debug")
	})
}

func TestLoggerError(t *testing.T) {
	// its not test, because there's no assertation
	opt := logger.Option{
		IsEnable: true,
	}
	log := logger.NewLogger(opt)
	ctx := logger.InjectCtx(context.Background(), logger.Context{
		ServiceName: "logger",
		Tag:         "xxx",
		ReqMethod:   "POST",
		ReqURI:      "/",
	})

	assert.NotPanics(t, func() {
		log.Error(ctx, "log Error", "detail log Error")
	})
}

func TestLoggerInfo(t *testing.T) {
	// its not test, because there's no assertation
	opt := logger.Option{
		IsEnable: true,
	}
	log := logger.NewLogger(opt)
	ctx := logger.InjectCtx(context.Background(), logger.Context{
		ServiceName: "logger",
		Tag:         "xxx",
		ReqMethod:   "POST",
		ReqURI:      "/",
	})

	assert.NotPanics(t, func() {
		log.Info(ctx, "log info", "detail log info")
	})
}

func TestLoggerWarn(t *testing.T) {
	// its not test, because there's no assertation
	opt := logger.Option{
		IsEnable: true,
	}
	log := logger.NewLogger(opt)
	ctx := logger.InjectCtx(context.Background(), logger.Context{
		ServiceName: "logger",
		Tag:         "xxx",
		ReqMethod:   "POST",
		ReqURI:      "/",
	})

	assert.NotPanics(t, func() {
		log.Warn(ctx, "log Warn", "detail log Warn")
	})
}

func TestLoggerPanic(t *testing.T) {
	// its not test, because there's no assertation
	opt := logger.Option{
		IsEnable: true,
	}
	log := logger.NewLogger(opt)
	ctx := logger.InjectCtx(context.Background(), logger.Context{
		ServiceName: "logger",
		Tag:         "xxx",
		ReqMethod:   "POST",
		ReqURI:      "/",
	})

	assert.Panics(t, func() {
		log.Panic(ctx, "log Panic", "detail log Panic")
	})
}
