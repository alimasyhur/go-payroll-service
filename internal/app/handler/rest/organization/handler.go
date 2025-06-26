package organization

import (
	"github.com/labstack/echo/v4"

	"github.com/weanan/weanan-service/internal/app/usecase/organization"
)

type OrganizationHandler interface {
	Create(c echo.Context) (err error)
	GetByID(c echo.Context) (err error)
}

type handler struct {
	organizationUsecase organization.OrganizationUsecase
}

func NewHandler() *handler {
	return &handler{}
}

func (h *handler) SetOrganizationUsecase(usecase organization.OrganizationUsecase) *handler {
	h.organizationUsecase = usecase
	return h
}

func (h *handler) Validate() OrganizationHandler {
	if h.organizationUsecase == nil {
		panic("organizationUsecase is nil")
	}

	return h
}
