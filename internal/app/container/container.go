package container

import (
	"github.com/alimasyhur/go-payroll-service/config"
	"github.com/alimasyhur/go-payroll-service/internal/app/driver"
	"github.com/alimasyhur/go-payroll-service/internal/app/repository"
	"github.com/alimasyhur/go-payroll-service/internal/app/usecase/attendance"
	"github.com/alimasyhur/go-payroll-service/internal/app/usecase/user"
	"github.com/alimasyhur/go-payroll-service/internal/pkg/logger"
)

type Container struct {
	Config            config.Config
	UserUsecase       user.UserUsecase
	AttendanceUsecase attendance.AttendanceUsecase
}

func Setup() *Container {
	// Load Config
	cfg := config.Load()

	logger.NewLogger(logger.Option{IsEnable: cfg.Logger.IsEnable})

	// Setup Driver
	db, _ := driver.NewPostgreSQLDatabase(cfg.DB)

	// Setup Repository
	userRepository := repository.NewUserRepository(db)
	attendancePeriodRepository := repository.NewAttendancePeriodRepository(db)

	// Setup Usecase
	userUsecase := user.NewUsecase().
		SetUserRepository(userRepository).
		Validate()

	attendancePeriodUsecase := attendance.NewUsecase().
		SetAttendancePeriodRepository(attendancePeriodRepository).
		Validate()

	return &Container{
		Config:            cfg,
		UserUsecase:       userUsecase,
		AttendanceUsecase: attendancePeriodUsecase,
	}
}
