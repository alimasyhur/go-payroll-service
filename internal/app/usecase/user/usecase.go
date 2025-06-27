package user

import (
	"context"

	"github.com/alimasyhur/go-payroll-service/internal/app/repository"
)

type UserUsecase interface {
	Login(ctx context.Context, req LoginRequest) (resp LoginResponse, err error)
	GenerateUserSeed(ctx context.Context) (err error)
}
type usecase struct {
	userRepository repository.User
}

func NewUsecase() *usecase {
	return &usecase{}
}

func (uc *usecase) SetUserRepository(r repository.User) *usecase {
	uc.userRepository = r
	return uc
}

func (uc *usecase) Validate() UserUsecase {
	if uc.userRepository == nil {
		panic("user repository is nil")
	}

	return uc
}
