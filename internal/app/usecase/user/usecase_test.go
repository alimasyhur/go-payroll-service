package user_test

import (
	"testing"

	"github.com/alimasyhur/go-payroll-service/config"
	mockRepo "github.com/alimasyhur/go-payroll-service/internal/app/repository/mocks"
	"github.com/alimasyhur/go-payroll-service/internal/app/usecase/user"

	"github.com/stretchr/testify/assert"
)

func TestNewUsecase(t *testing.T) {

	t.Run("ShouldPanicWhenUserRepositoryIsNil", func(t *testing.T) {
		assert.Panics(t, func() {
			user.NewUsecase().Validate()
		})
	})

	t.Run("ShouldPanicWhenAppConfigIsNil", func(t *testing.T) {
		userRepository := new(mockRepo.User)

		assert.Panics(t, func() {
			user.NewUsecase().
				SetUserRepository(userRepository).
				Validate()
		})
	})

	t.Run("ShouldNotPanic", func(t *testing.T) {
		userRepository := new(mockRepo.User)
		cfg := config.Config{}

		assert.NotPanics(t, func() {
			user.NewUsecase().
				SetConfig(cfg.App).
				SetUserRepository(userRepository).
				Validate()
		})
	})
}
