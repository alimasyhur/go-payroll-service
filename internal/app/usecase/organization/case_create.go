package organization

import (
	"context"
	"net/http"

	"go.elastic.co/apm/v2"

	"github.com/weanan/weanan-service/internal/app/entity"
	"github.com/weanan/weanan-service/internal/pkg/apperror"
)

func (u *usecase) Create(ctx context.Context, req CreateUpdateOrganizationRequest) (resp OrganizationResponse, err error) {
	span, ctx := apm.StartSpan(ctx, "usecase.Create", "custom")
	defer span.End()

	organizationData := entity.Organization{
		Name:     req.Name,
		Location: req.Location,
	}

	err = u.organizationRepository.Create(ctx, &organizationData)
	if err != nil {
		err = apperror.New(http.StatusUnprocessableEntity, err)
		return
	}

	resp = OrganizationResponse{
		ID:       organizationData.ID,
		Name:     organizationData.Name,
		Location: organizationData.Location,
	}

	return
}
