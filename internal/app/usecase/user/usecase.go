package user

import (
	"context"

	"github.com/alimasyhur/go-payroll-service/config"
	"github.com/alimasyhur/go-payroll-service/internal/app/repository"
)

type UserUsecase interface {
	Login(ctx context.Context, req LoginRequest) (resp LoginResponse, err error)
	GenerateUserSeed(ctx context.Context) (err error)
}
type usecase struct {
	config         *config.AppConfig
	userRepository repository.User
}

func NewUsecase() *usecase {
	return &usecase{}
}

func (uc *usecase) SetConfig(c config.AppConfig) *usecase {
	uc.config = &c
	return uc
}

func (uc *usecase) SetUserRepository(r repository.User) *usecase {
	uc.userRepository = r
	return uc
}

func (uc *usecase) Validate() UserUsecase {
	if uc.config == nil {
		panic("app config is nil")
	}
	if uc.userRepository == nil {
		panic("user repository is nil")
	}

	return uc
}
