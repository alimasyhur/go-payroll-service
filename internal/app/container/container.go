package container

import (
	"github.com/alimasyhur/go-payroll-service/config"
	"github.com/alimasyhur/go-payroll-service/internal/app/driver"
	"github.com/alimasyhur/go-payroll-service/internal/app/repository"
	"github.com/alimasyhur/go-payroll-service/internal/app/usecase/attendance"
	"github.com/alimasyhur/go-payroll-service/internal/app/usecase/overtime"
	"github.com/alimasyhur/go-payroll-service/internal/app/usecase/user"
	"github.com/alimasyhur/go-payroll-service/internal/pkg/logger"
)

type Container struct {
	Config            config.Config
	UserUsecase       user.UserUsecase
	AttendanceUsecase attendance.AttendanceUsecase
	OvertimeUsecase   overtime.OvertimeUsecase
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
	attendanceRepository := repository.NewAttendanceRepository(db)
	overtimeRepository := repository.NewOvertimeRepository(db)

	// Setup Usecase
	userUsecase := user.NewUsecase().
		SetUserRepository(userRepository).
		Validate()

	attendancePeriodUsecase := attendance.NewUsecase().
		SetAttendancePeriodRepository(attendancePeriodRepository).
		SetAttendanceRepository(attendanceRepository).
		Validate()

	overtimeUsecase := overtime.NewUsecase().
		SetAttendancePeriodRepository(attendancePeriodRepository).
		SetOvertimeRepository(overtimeRepository).
		Validate()

	return &Container{
		Config:            cfg,
		UserUsecase:       userUsecase,
		AttendanceUsecase: attendancePeriodUsecase,
		OvertimeUsecase:   overtimeUsecase,
	}
}
