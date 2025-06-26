package container

import (
	"github.com/weanan/weanan-service/config"
	"github.com/weanan/weanan-service/internal/app/driver"
	"github.com/weanan/weanan-service/internal/app/repository"
	"github.com/weanan/weanan-service/internal/app/usecase/organization"
	"github.com/weanan/weanan-service/internal/app/wrapper/beeceptor"
	"github.com/weanan/weanan-service/internal/pkg/logger"
)

type Container struct {
	Config              config.Config
	OrganizationUsecase organization.OrganizationUsecase
}

func Setup() *Container {
	// Load Config
	cfg := config.Load()

	logger.NewLogger(logger.Option{IsEnable: cfg.Logger.IsEnable})

	// Setup Driver
	db, _ := driver.NewMySQLDatabase(cfg.DB)

	// Setup Repository
	organizationRepository := repository.NewOrganizationRepository(db)

	// Setup Wrapper
	beeceptorWrapper := beeceptor.NewWrapper().SetConfig(cfg.Beeceptor).Setup().Validate()

	// Setup Usecase
	organizationUsecase := organization.NewUsecase().
		SetOrganizationRepository(organizationRepository).
		SetBeeceptorWrapper(beeceptorWrapper).
		Validate()

	return &Container{
		Config:              cfg,
		OrganizationUsecase: organizationUsecase,
	}
}
