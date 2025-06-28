package payslip_test

import (
	"testing"

	mockRepo "github.com/alimasyhur/go-payroll-service/internal/app/repository/mocks"
	"github.com/alimasyhur/go-payroll-service/internal/app/usecase/payslip"

	"github.com/stretchr/testify/assert"
)

func TestNewUsecase(t *testing.T) {

	t.Run("ShouldPanicWhenAttendancePeriodRepositoryIsNil", func(t *testing.T) {
		assert.Panics(t, func() {
			payslip.NewUsecase().Validate()
		})
	})

	t.Run("ShouldPanicWhenOvertimeRepositoryIsNil", func(t *testing.T) {
		attendancePeriodRepository := new(mockRepo.AttendancePeriod)

		assert.Panics(t, func() {
			payslip.NewUsecase().
				SetAttendancePeriodRepository(attendancePeriodRepository).
				Validate()
		})
	})

	t.Run("ShouldPanicWhenAttendanceRepositoryIsNil", func(t *testing.T) {
		attendancePeriodRepository := new(mockRepo.AttendancePeriod)
		overtimeRepository := new(mockRepo.Overtime)

		assert.Panics(t, func() {
			payslip.NewUsecase().
				SetAttendancePeriodRepository(attendancePeriodRepository).
				SetOvertimeRepository(overtimeRepository).
				Validate()
		})
	})

	t.Run("ShouldPanicWhenReimbursementRepositoryIsNil", func(t *testing.T) {
		attendancePeriodRepository := new(mockRepo.AttendancePeriod)
		overtimeRepository := new(mockRepo.Overtime)
		attendanceRepository := new(mockRepo.Attendance)

		assert.Panics(t, func() {
			payslip.NewUsecase().
				SetAttendancePeriodRepository(attendancePeriodRepository).
				SetOvertimeRepository(overtimeRepository).
				SetAttendanceRepository(attendanceRepository).
				Validate()
		})
	})

	t.Run("ShouldPanicWhenEmployeeSalaryRepositoryIsNil", func(t *testing.T) {
		attendancePeriodRepository := new(mockRepo.AttendancePeriod)
		overtimeRepository := new(mockRepo.Overtime)
		attendanceRepository := new(mockRepo.Attendance)
		reimbursementRepository := new(mockRepo.Reimbursement)

		assert.Panics(t, func() {
			payslip.NewUsecase().
				SetAttendancePeriodRepository(attendancePeriodRepository).
				SetOvertimeRepository(overtimeRepository).
				SetAttendanceRepository(attendanceRepository).
				SetReimbursementRepository(reimbursementRepository).
				Validate()
		})
	})

	t.Run("ShouldPanicWhenPayrollRepositoryIsNil", func(t *testing.T) {
		attendancePeriodRepository := new(mockRepo.AttendancePeriod)
		overtimeRepository := new(mockRepo.Overtime)
		attendanceRepository := new(mockRepo.Attendance)
		reimbursementRepository := new(mockRepo.Reimbursement)
		employeeSalaryRepository := new(mockRepo.EmployeeSalary)

		assert.Panics(t, func() {
			payslip.NewUsecase().
				SetAttendancePeriodRepository(attendancePeriodRepository).
				SetOvertimeRepository(overtimeRepository).
				SetAttendanceRepository(attendanceRepository).
				SetReimbursementRepository(reimbursementRepository).
				SetEmployeeSalaryRepository(employeeSalaryRepository).
				Validate()
		})
	})

	t.Run("ShouldPanicWhenPayslipRepositoryIsNil", func(t *testing.T) {
		attendancePeriodRepository := new(mockRepo.AttendancePeriod)
		overtimeRepository := new(mockRepo.Overtime)
		attendanceRepository := new(mockRepo.Attendance)
		reimbursementRepository := new(mockRepo.Reimbursement)
		employeeSalaryRepository := new(mockRepo.EmployeeSalary)
		payrollRepository := new(mockRepo.Payroll)

		assert.Panics(t, func() {
			payslip.NewUsecase().
				SetAttendancePeriodRepository(attendancePeriodRepository).
				SetOvertimeRepository(overtimeRepository).
				SetAttendanceRepository(attendanceRepository).
				SetReimbursementRepository(reimbursementRepository).
				SetEmployeeSalaryRepository(employeeSalaryRepository).
				SetPayrollRepository(payrollRepository).
				Validate()
		})
	})

	t.Run("ShouldPanicWhenUserRepositoryIsNil", func(t *testing.T) {
		attendancePeriodRepository := new(mockRepo.AttendancePeriod)
		overtimeRepository := new(mockRepo.Overtime)
		attendanceRepository := new(mockRepo.Attendance)
		reimbursementRepository := new(mockRepo.Reimbursement)
		employeeSalaryRepository := new(mockRepo.EmployeeSalary)
		payrollRepository := new(mockRepo.Payroll)
		payslipRepository := new(mockRepo.Payslip)

		assert.Panics(t, func() {
			payslip.NewUsecase().
				SetAttendancePeriodRepository(attendancePeriodRepository).
				SetOvertimeRepository(overtimeRepository).
				SetAttendanceRepository(attendanceRepository).
				SetReimbursementRepository(reimbursementRepository).
				SetEmployeeSalaryRepository(employeeSalaryRepository).
				SetPayrollRepository(payrollRepository).
				SetPayslipRepository(payslipRepository).
				Validate()
		})
	})

	t.Run("ShouldNotPanic", func(t *testing.T) {
		attendancePeriodRepository := new(mockRepo.AttendancePeriod)
		overtimeRepository := new(mockRepo.Overtime)
		attendanceRepository := new(mockRepo.Attendance)
		reimbursementRepository := new(mockRepo.Reimbursement)
		employeeSalaryRepository := new(mockRepo.EmployeeSalary)
		payrollRepository := new(mockRepo.Payroll)
		payslipRepository := new(mockRepo.Payslip)
		userRepository := new(mockRepo.User)

		assert.NotPanics(t, func() {
			payslip.NewUsecase().
				SetAttendancePeriodRepository(attendancePeriodRepository).
				SetOvertimeRepository(overtimeRepository).
				SetAttendanceRepository(attendanceRepository).
				SetReimbursementRepository(reimbursementRepository).
				SetEmployeeSalaryRepository(employeeSalaryRepository).
				SetPayrollRepository(payrollRepository).
				SetPayslipRepository(payslipRepository).
				SetUserRepository(userRepository).
				Validate()
		})
	})
}
