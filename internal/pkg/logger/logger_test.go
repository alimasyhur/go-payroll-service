package logger_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/weanan/weanan-service/internal/pkg/logger"
)

func TestInjectCtx(t *testing.T) {
	ctxVal := logger.Context{
		ServiceName: "logger",
		Tag:         "xxx",
		ReqMethod:   "POST",
		ReqURI:      "/",
	}

	t.Run("Success on nil context", func(t *testing.T) {
		ctx := logger.InjectCtx(nil, ctxVal)
		assert.NotEmpty(t, ctx)

		ctxValGet := logger.ExtractCtx(ctx)
		assert.EqualValues(t, ctxVal, ctxValGet)
	})

	t.Run("Success", func(t *testing.T) {
		ctx := logger.InjectCtx(context.Background(), ctxVal)
		assert.NotEmpty(t, ctx)

		ctxValGet := logger.ExtractCtx(ctx)
		assert.EqualValues(t, ctxVal, ctxValGet)
	})
}

// TestInjectCtx2 test if inject multiple times it still got the last one
func TestInjectCtx2(t *testing.T) {
	ctxVal1 := logger.Context{
		ServiceName: "logger",
		Tag:         "xxx",
		ReqMethod:   "POST",
		ReqURI:      "/",
	}

	ctxVal2 := logger.Context{
		ServiceName: "logger2",
		Tag:         "yyy",
		ReqMethod:   "GET",
		ReqURI:      "/",
	}

	t.Run("Last context win", func(t *testing.T) {
		ctx := logger.InjectCtx(context.Background(), ctxVal1)
		assert.NotEmpty(t, ctx)

		ctx = logger.InjectCtx(ctx, ctxVal2)
		assert.NotEmpty(t, ctx)

		ctxValGet := logger.ExtractCtx(ctx)
		assert.EqualValues(t, ctxVal2, ctxValGet)
	})

	t.Run("Redefine context with same struct name doesn't matter", func(t *testing.T) {
		// inject value 1 on same context
		ctx := logger.InjectCtx(context.Background(), ctxVal1)
		assert.NotEmpty(t, ctx)

		// redefine struct with same name and same variable name,
		type ctxKeyLogger struct{}

		var ctxKey = ctxKeyLogger{}

		// set value 2, should return value 2 if using ctx key
		ctx = context.WithValue(ctx, ctxKey, ctxVal2)
		assert.NotEmpty(t, ctx)

		// unless using same variable as above
		val, ok := ctx.Value(ctxKey).(logger.Context)
		assert.EqualValues(t, val, ctxVal2)
		assert.True(t, ok)

		// should still return context value 1 on extract
		ctxValGet := logger.ExtractCtx(ctx)
		assert.EqualValues(t, ctxVal1, ctxValGet)

	})
}

func TestExtractLogCtx(t *testing.T) {
	t.Run("Success on nil", func(t *testing.T) {
		ctxVal := logger.ExtractCtx(nil)
		assert.Empty(t, ctxVal)
	})

	t.Run("Success", func(t *testing.T) {
		ctxVal := logger.ExtractCtx(context.Background())
		assert.Empty(t, ctxVal)
	})
}
