package organization_test

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/weanan/weanan-service/internal/app/entity"
	mockRepository "github.com/weanan/weanan-service/internal/app/repository/mocks"
	"github.com/weanan/weanan-service/internal/app/usecase/organization"
	"github.com/weanan/weanan-service/internal/app/wrapper/beeceptor"
	mockBeeceptor "github.com/weanan/weanan-service/internal/app/wrapper/beeceptor/mocks"
	"github.com/weanan/weanan-service/internal/pkg/apperror"
)

func TestGetById(t *testing.T) {
	cases := map[string]struct {
		ShouldError                    bool
		IsAppError                     bool
		ExpectedAppErrorStatus         int
		GetOrganizationResp            entity.Organization
		GetOrganizationErr             error
		GetOrganizationSubdistrictResp beeceptor.GetOrganizationSubdistrictResponse
		GetOrganizationSubdistrictErr  error
	}{
		"ShouldErrorWhenGetOrganizationError": {
			ShouldError:            true,
			IsAppError:             true,
			ExpectedAppErrorStatus: 422,
			GetOrganizationErr:     errors.New("get organization error"),
		},
		"ShouldErrorWhenGetSubdistrictByOrganizationNameError": {
			ShouldError:                   true,
			IsAppError:                    true,
			ExpectedAppErrorStatus:        400,
			GetOrganizationSubdistrictErr: errors.New("get subdistrict organization error"),
		},
		"ShouldSuccess": {
			ShouldError: false,
		},
	}

	for v, test := range cases {
		t.Run(v, func(t *testing.T) {
			beeceptorWrapper := new(mockBeeceptor.BeeceptorWrapper)
			organizationRepo := new(mockRepository.Organization)

			organizationRepo.On("GetById", mock.Anything, mock.Anything).Return(test.GetOrganizationResp, test.GetOrganizationErr).Once()
			beeceptorWrapper.On("GetSubdistrictByOrganizationName", mock.Anything, mock.Anything).Return(test.GetOrganizationSubdistrictResp, test.GetOrganizationSubdistrictErr).Once()

			usecase := organization.NewUsecase().SetOrganizationRepository(organizationRepo).SetBeeceptorWrapper(beeceptorWrapper).Validate()

			resp, err := usecase.GetById(context.Background(), 100)

			if test.ShouldError {
				assert.NotNil(t, err)

				if test.IsAppError {
					var appErr *apperror.ApplicationError
					assert.ErrorAs(t, err, &appErr)

					appErr = err.(*apperror.ApplicationError)
					assert.Equal(t, test.ExpectedAppErrorStatus, appErr.Status)
				}
			} else {
				assert.Nil(t, err)
				assert.NotNil(t, resp)
			}
		})
	}
}
