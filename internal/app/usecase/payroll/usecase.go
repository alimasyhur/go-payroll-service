package payroll

import (
	"context"

	"github.com/alimasyhur/go-payroll-service/internal/app/repository"
)

type PayrollUsecase interface {
	CreatePayroll(ctx context.Context, req CreatePayrollRequest) (resp CreatePayrollResponse, err error)
}
type usecase struct {
	attendancePeriodRepository repository.AttendancePeriod
	overtimeRepository         repository.Overtime
	attendanceRepository       repository.Attendance
	reimbursementRepository    repository.Reimbursement
	employeeSalaryRepository   repository.EmployeeSalary
	payrollRepository          repository.Payroll
	payslipRepository          repository.Payslip
	userRepository             repository.User
}

func NewUsecase() *usecase {
	return &usecase{}
}

func (uc *usecase) SetAttendancePeriodRepository(r repository.AttendancePeriod) *usecase {
	uc.attendancePeriodRepository = r
	return uc
}

func (uc *usecase) SetAttendanceRepository(r repository.Attendance) *usecase {
	uc.attendanceRepository = r
	return uc
}

func (uc *usecase) SetOvertimeRepository(r repository.Overtime) *usecase {
	uc.overtimeRepository = r
	return uc
}

func (uc *usecase) SetReimbursementRepository(r repository.Reimbursement) *usecase {
	uc.reimbursementRepository = r
	return uc
}

func (uc *usecase) SetEmployeeSalaryRepository(r repository.EmployeeSalary) *usecase {
	uc.employeeSalaryRepository = r
	return uc
}

func (uc *usecase) SetPayrollRepository(r repository.Payroll) *usecase {
	uc.payrollRepository = r
	return uc
}

func (uc *usecase) SetPayslipRepository(r repository.Payslip) *usecase {
	uc.payslipRepository = r
	return uc
}

func (uc *usecase) SetUserRepository(r repository.User) *usecase {
	uc.userRepository = r
	return uc
}

func (uc *usecase) Validate() PayrollUsecase {
	if uc.attendancePeriodRepository == nil {
		panic("attendance period repository is nil")
	}
	if uc.overtimeRepository == nil {
		panic("overtime repository is nil")
	}
	if uc.attendanceRepository == nil {
		panic("attendance repository is nil")
	}
	if uc.reimbursementRepository == nil {
		panic("reimbursement repository is nil")
	}
	if uc.employeeSalaryRepository == nil {
		panic("employee salary repository is nil")
	}
	if uc.payrollRepository == nil {
		panic("payroll repository is nil")
	}
	if uc.payslipRepository == nil {
		panic("payslip repository is nil")
	}
	if uc.userRepository == nil {
		panic("user repository is nil")
	}

	return uc
}
