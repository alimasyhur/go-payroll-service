package payroll

import (
	"context"
	"fmt"
	"time"

	"github.com/alimasyhur/go-payroll-service/internal/app/entity"
	"github.com/google/uuid"
)

func (uc *usecase) CreatePayroll(ctx context.Context, req CreatePayrollRequest) (resp CreatePayrollResponse, err error) {
	now := time.Now()

	period, err := uc.attendancePeriodRepository.GetOneByUUID(ctx, req.PeriodUUID)
	if err != nil {
		return resp, fmt.Errorf("unable to get Period. %s", err.Error())
	}

	if period.IsClosed {
		return resp, fmt.Errorf("attendance period is already closed")
	}

	payroll, _ := uc.payrollRepository.GetOneByPeriodUUID(ctx, period.UUID)
	if payroll.UUID != "" {
		return resp, fmt.Errorf("payroll is already exist")
	}

	newPayroll := entity.Payroll{
		UUID:                 uuid.New().String(),
		AttendancePeriodUUID: req.PeriodUUID,
		ProcessedAt:          now,
		IP:                   req.IP,
		CreatedBy:            req.UserUUID,
		CreatedAt:            now,
		UpdatedAt:            now,
	}

	if _, err := uc.payrollRepository.CreatePayroll(ctx, newPayroll); err != nil {
		return resp, fmt.Errorf("unable to create payroll. %s", err.Error())
	}

	employees, err := uc.userRepository.GetListByRole(ctx, "employee")
	if err != nil {
		return resp, fmt.Errorf("unable to get employee. %s", err.Error())
	}

	for _, employee := range employees {
		payslipPayload := GeneratePayslipRequest{
			UserUUID:    employee.UUID,
			PayrollUUID: newPayroll.UUID,
			PeriodUUID:  req.PeriodUUID,
			StartDate:   period.StartDate.Format(time.DateOnly),
			EndDate:     period.EndDate.Format(time.DateOnly),
		}

		if err := uc.generatePayslipForEmployee(ctx, payslipPayload); err != nil {
			return resp, fmt.Errorf("unable to generate payslip. %s", err.Error())
		}
	}

	period.IsClosed = true
	if err := uc.attendancePeriodRepository.UpdateAttendancePeriod(ctx, period); err != nil {
		return resp, fmt.Errorf("unable to update period. %s", err.Error())
	}

	resp.UUID = newPayroll.UUID

	return resp, nil
}

func (uc *usecase) generatePayslipForEmployee(ctx context.Context, req GeneratePayslipRequest) (err error) {
	salary, _ := uc.employeeSalaryRepository.GetOneByUserUUID(ctx, req.UserUUID)

	workDays, _ := uc.attendanceRepository.GetWorkdaysByUserDaterange(ctx, req.UserUUID, req.StartDate, req.EndDate)

	overtimes, _ := uc.overtimeRepository.GetListByUserDaterange(ctx, req.UserUUID, req.StartDate, req.EndDate)
	var totalOvertime float64
	dailySalary := salary.Amount / 22
	for _, overtime := range overtimes {
		totalOvertime += float64(overtime.Hours) * 2 * (dailySalary / 8)
	}

	reimbursements, _ := uc.reimbursementRepository.GetListByUserDaterange(ctx, req.UserUUID, req.StartDate, req.EndDate)

	var totalReimburse float64
	for _, r := range reimbursements {
		totalReimburse += r.Amount
	}

	fmt.Println(salary, workDays, overtimes)

	total := float64(workDays)*dailySalary + totalOvertime + totalReimburse

	payslip := entity.Payslip{
		UUID:        uuid.New().String(),
		PayrollUUID: req.PayrollUUID,
		UserUUID:    req.UserUUID,
		WorkDays:    workDays,
		BaseSalary:  salary.Amount,
		Overtime:    totalOvertime,
		Reimburse:   totalReimburse,
		Total:       total,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	_, err = uc.payslipRepository.CreatePayslip(ctx, payslip)

	return err
}
