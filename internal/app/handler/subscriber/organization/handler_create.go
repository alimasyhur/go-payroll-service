package organization

import (
	"context"
	"encoding/json"

	"cloud.google.com/go/pubsub"
	"github.com/weanan/weanan-service/internal/app/usecase/organization"
)

func (h *handler) Create(ctx context.Context, msg pubsub.Message) (err error) {
	req := organization.CreateUpdateOrganizationRequest{}

	body, err := json.Marshal(msg.Data)
	if err != nil {
		return
	}

	err = json.Unmarshal(body, &req)
	if err != nil {
		return
	}

	_, err = h.organizationUsecase.Create(ctx, req)
	return
}
