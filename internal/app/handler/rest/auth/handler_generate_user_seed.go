package auth

import (
	"context"
)

func (h *handler) GenerateUserSeed(c context.Context) (err error) {
	err = h.userUsecase.GenerateUserSeed(c)

	return err
}
