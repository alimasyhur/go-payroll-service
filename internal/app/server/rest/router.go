package rest

import (
	"github.com/labstack/echo/v4"

	"github.com/alimasyhur/go-payroll-service/internal/app/container"
	"github.com/alimasyhur/go-payroll-service/internal/app/handler/rest/attendance"
	"github.com/alimasyhur/go-payroll-service/internal/app/handler/rest/auth"
	"github.com/alimasyhur/go-payroll-service/internal/app/handler/rest/health_check"
	"github.com/alimasyhur/go-payroll-service/internal/app/handler/rest/overtime"
	"github.com/alimasyhur/go-payroll-service/internal/app/handler/rest/payroll"
	"github.com/alimasyhur/go-payroll-service/internal/app/handler/rest/payslip"
	"github.com/alimasyhur/go-payroll-service/internal/app/handler/rest/reimbursement"
)

func SetupRouter(server *echo.Echo, container *container.Container) {
	// inject handler with usecase via container
	healthCheckHandler := health_check.NewHandler().Validate()
	authHandler := auth.NewHandler().
		SetAuthUsecase(container.UserUsecase).
		Validate()
	attendanceHandler := attendance.NewHandler().
		SetAttendanceUsecase(container.AttendanceUsecase).
		Validate()
	overtimeHandler := overtime.NewHandler().
		SetOvertimeUsecase(container.OvertimeUsecase).
		Validate()
	reimbursementHandler := reimbursement.NewHandler().
		SetReimbursementUsecase(container.ReimbursementUsecase).
		Validate()
	payrollHandler := payroll.NewHandler().
		SetPayrollUsecase(container.PayrollUsecase).
		Validate()
	payslipHandler := payslip.NewHandler().
		SetPayslipUsecase(container.PayslipUsecase).
		Validate()

	server.GET("/health-check", healthCheckHandler.Check)
	server.POST("/login", authHandler.Login)

	private := server.Group("")
	{
		private.Use(JWTAuthMiddleware())
		private.POST("/attendances", attendanceHandler.CreateAttendance)
		private.POST("/overtimes", overtimeHandler.CreateOvertime)
		private.POST("/reimbursements", reimbursementHandler.CreateReimbursement)
		private.GET("/payslips/:payroll_uuid", payslipHandler.GetOnePayslip)

		admin := private.Group("")
		{
			admin.Use(AdminOnlyMiddleware())
			admin.POST("/attendance-periods", attendanceHandler.CreateAttendancePeriod)
			admin.POST("/payrolls/run", payrollHandler.CreatePayroll)
		}
	}

}
