package organization

import (
	"context"
	"net/http"

	"github.com/spf13/cast"
	"github.com/weanan/weanan-service/internal/pkg/apperror"
	"go.elastic.co/apm/v2"
)

func (u *usecase) GetById(ctx context.Context, id uint) (resp OrganizationResponse, err error) {
	span, ctx := apm.StartSpan(ctx, "usecase.GetById", "custom")
	defer span.End()

	organizationData, err := u.organizationRepository.GetById(ctx, id)
	if err != nil {
		err = apperror.New(http.StatusUnprocessableEntity, err)
		return
	}

	subdisctrict, err := u.beeceptorWrapper.GetSubdistrictByOrganizationName(ctx, organizationData.Name)
	if err != nil {
		err = apperror.New(http.StatusBadRequest, err)
		return
	}

	resp = OrganizationResponse{
		ID:       organizationData.ID,
		Name:     organizationData.Name,
		Location: organizationData.Location,
		Subdisctrict: &SubdisctrictDetail{
			ID:       cast.ToUint(subdisctrict.Data.ID),
			Name:     subdisctrict.Data.Location,
			Location: subdisctrict.Data.Name,
		},
	}

	return
}
