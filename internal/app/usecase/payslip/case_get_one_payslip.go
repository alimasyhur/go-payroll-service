package payslip

import (
	"context"
	"fmt"
	"time"
)

func (uc *usecase) GetOnePayslip(ctx context.Context, req GetOnePayslipRequest) (resp GetOnePayslipResponse, err error) {
	payslip, err := uc.payslipRepository.GetOneByUserPayrollUUID(ctx, req.UserUUID, req.PayrollUUID)
	if err != nil {
		return resp, fmt.Errorf("unable to get payslip. %s", err.Error())
	}

	daily := payslip.BaseSalary / 22
	attendance := payslip.WorkDays * int64(daily)

	payslipDetail, err := uc.payslipRepository.GetOneDetailByPayrollUUID(ctx, req.PayrollUUID)
	if err != nil {
		return resp, fmt.Errorf("unable to get payslip detail. %s", err.Error())
	}

	overtimes, err := uc.overtimeRepository.GetListByUserDaterange(ctx, req.UserUUID, payslipDetail.PeriodStartDate, payslipDetail.PeriodEndDate)
	if err != nil {
		return resp, fmt.Errorf("unable to get list overtimes. %s", err.Error())
	}

	var overtimeItems []OvertimeItem
	for _, ot := range overtimes {
		val := float64(ot.Hours) * 2 * (daily / 8)
		overtimeItems = append(overtimeItems, OvertimeItem{
			Date:  ot.Date.Format(time.DateOnly),
			Hours: ot.Hours,
			Value: val,
		})
	}

	reimbursements, err := uc.reimbursementRepository.GetListByUserDaterange(ctx, req.UserUUID, payslipDetail.PeriodStartDate, payslipDetail.PeriodEndDate)
	if err != nil {
		return resp, fmt.Errorf("unable to get list reimbursements. %s", err.Error())
	}

	var reimburseItems []ReimbursementItem
	var reimTotal float64
	for _, r := range reimbursements {
		reimTotal += r.Amount
		reimburseItems = append(reimburseItems, ReimbursementItem{
			Date:        r.Date.Format(time.DateOnly),
			Amount:      r.Amount,
			Description: r.Description,
		})
	}

	resp = GetOnePayslipResponse{
		UUID:                payslip.UUID,
		PayrollUUID:         payslip.PayrollUUID,
		UserUUID:            payslip.UserUUID,
		WorkDays:            payslip.WorkDays,
		BaseSalary:          payslip.BaseSalary,
		DailySalary:         daily,
		Attendance:          float64(attendance),
		Overtime:            payslip.Overtime,
		OvertimeDetail:      overtimeItems,
		Reimbursement:       payslip.Reimburse,
		ReimbursementDetail: reimburseItems,
		Total:               payslip.Total,
	}

	return resp, nil
}
