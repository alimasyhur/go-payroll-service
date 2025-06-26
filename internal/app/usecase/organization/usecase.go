package organization

import (
	"context"

	"github.com/weanan/weanan-service/internal/app/repository"
	"github.com/weanan/weanan-service/internal/app/wrapper/beeceptor"
)

type OrganizationUsecase interface {
	GetById(ctx context.Context, id uint) (resp OrganizationResponse, err error)
	Create(ctx context.Context, req CreateUpdateOrganizationRequest) (resp OrganizationResponse, err error)
}

type usecase struct {
	organizationRepository repository.Organization
	beeceptorWrapper       beeceptor.BeeceptorWrapper
}

func NewUsecase() *usecase {
	return &usecase{}
}

func (s *usecase) SetOrganizationRepository(repo repository.Organization) *usecase {
	s.organizationRepository = repo
	return s
}

func (s *usecase) SetBeeceptorWrapper(wrapper beeceptor.BeeceptorWrapper) *usecase {
	s.beeceptorWrapper = wrapper
	return s
}

func (s *usecase) Validate() OrganizationUsecase {
	if s.organizationRepository == nil {
		panic("organizationRepository is nil")
	}

	if s.beeceptorWrapper == nil {
		panic("beeceptorWrapper is nil")
	}

	return s
}
