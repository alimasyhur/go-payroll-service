package organization_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	mockRepository "github.com/weanan/weanan-service/internal/app/repository/mocks"
	"github.com/weanan/weanan-service/internal/app/usecase/organization"
	mockBeeceptor "github.com/weanan/weanan-service/internal/app/wrapper/beeceptor/mocks"
)

func TestNewUsecase(t *testing.T) {
	beeceptorWrapper := new(mockBeeceptor.BeeceptorWrapper)
	organizationRepo := new(mockRepository.Organization)

	t.Run("ShouldPanicWhenOrganizationRepoIsNil", func(t *testing.T) {
		assert.Panics(t, func() {
			organization.NewUsecase().Validate()
		})
	})

	t.Run("ShouldPanicWhenBeeceptorWrapperIsNil", func(t *testing.T) {
		assert.Panics(t, func() {
			organization.NewUsecase().SetOrganizationRepository(organizationRepo).Validate()
		})
	})

	t.Run("ShouldNotPanic", func(t *testing.T) {
		assert.NotPanics(t, func() {
			organization.NewUsecase().SetOrganizationRepository(organizationRepo).SetBeeceptorWrapper(beeceptorWrapper).Validate()
		})
	})
}
