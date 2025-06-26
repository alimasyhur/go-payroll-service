package beeceptor

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/weanan/weanan-service/config"
)

func TestNewWrapper(t *testing.T) {
	cfg := config.BeeceptorConfig{}

	t.Run("ShouldPanicWhenLoggerIsNil", func(t *testing.T) {
		assert.Panics(t, func() {
			NewWrapper().SetConfig(cfg).Validate()
		})
	})

	t.Run("ShouldNotPanic", func(t *testing.T) {
		assert.NotPanics(t, func() {
			NewWrapper().SetConfig(cfg).Setup().Validate()
		})
	})
}
