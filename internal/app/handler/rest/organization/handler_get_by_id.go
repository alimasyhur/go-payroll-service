package organization

import (
	"github.com/labstack/echo/v4"
	"github.com/spf13/cast"

	"github.com/weanan/weanan-service/internal/pkg/response"
)

func (h *handler) GetByID(c echo.Context) (err error) {
	ctx := c.Request().Context()

	id := c.Param("id")

	resp, err := h.organizationUsecase.GetById(ctx, cast.ToUint(id))
	if err != nil {
		return
	}

	return response.ResponseSuccess(c, resp)
}
