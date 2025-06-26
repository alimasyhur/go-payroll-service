package organization

import (
	"github.com/labstack/echo/v4"

	"github.com/weanan/weanan-service/internal/app/usecase/organization"
	"github.com/weanan/weanan-service/internal/pkg/response"
	"github.com/weanan/weanan-service/internal/pkg/validator"
)

func (h *handler) Create(c echo.Context) (err error) {
	ctx := c.Request().Context()

	req := organization.CreateUpdateOrganizationRequest{}
	if err = validator.Validate(c, &req); err != nil {
		return
	}

	resp, err := h.organizationUsecase.Create(ctx, req)
	if err != nil {
		return
	}

	return response.ResponseSuccess(c, resp)
}
