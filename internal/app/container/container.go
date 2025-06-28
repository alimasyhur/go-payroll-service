package container

import (
	"github.com/alimasyhur/go-payroll-service/config"
	"github.com/alimasyhur/go-payroll-service/internal/app/driver"
	"github.com/alimasyhur/go-payroll-service/internal/app/repository"
	"github.com/alimasyhur/go-payroll-service/internal/app/usecase/attendance"
	"github.com/alimasyhur/go-payroll-service/internal/app/usecase/overtime"
	"github.com/alimasyhur/go-payroll-service/internal/app/usecase/payroll"
	"github.com/alimasyhur/go-payroll-service/internal/app/usecase/reimbursement"
	"github.com/alimasyhur/go-payroll-service/internal/app/usecase/user"
	"github.com/alimasyhur/go-payroll-service/internal/pkg/logger"
)

type Container struct {
	Config               config.Config
	UserUsecase          user.UserUsecase
	AttendanceUsecase    attendance.AttendanceUsecase
	OvertimeUsecase      overtime.OvertimeUsecase
	ReimbursementUsecase reimbursement.ReimbursementUsecase
	PayrollUsecase       payroll.PayrollUsecase
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
	reimbursementRepository := repository.NewReimbursementRepository(db)
	employeeSalaryRepository := repository.NewEmployeeSalaryRepository(db)
	payrollRepository := repository.NewPayrollRepository(db)
	payslipRepository := repository.NewPayslipRepository(db)

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

	reimbursementUsecase := reimbursement.NewUsecase().
		SetAttendancePeriodRepository(attendancePeriodRepository).
		SetReimbursementRepository(reimbursementRepository).
		Validate()

	payrollUsecase := payroll.NewUsecase().
		SetAttendancePeriodRepository(attendancePeriodRepository).
		SetOvertimeRepository(overtimeRepository).
		SetAttendanceRepository(attendanceRepository).
		SetReimbursementRepository(reimbursementRepository).
		SetEmployeeSalaryRepository(employeeSalaryRepository).
		SetPayrollRepository(payrollRepository).
		SetPayslipRepository(payslipRepository).
		SetUserRepository(userRepository).
		Validate()

	return &Container{
		Config:               cfg,
		UserUsecase:          userUsecase,
		AttendanceUsecase:    attendancePeriodUsecase,
		OvertimeUsecase:      overtimeUsecase,
		ReimbursementUsecase: reimbursementUsecase,
		PayrollUsecase:       payrollUsecase,
	}
}
