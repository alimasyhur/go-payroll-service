package payslip

import (
	"context"
	"errors"
	"fmt"
)

func (uc *usecase) GetSummary(ctx context.Context, req GetSummaryRequest) (resp GetSummaryResponse, err error) {
	payslips, err := uc.payslipRepository.GetListByPayrollUUID(ctx, req.PayrollUUID)
	if err != nil {
		return resp, fmt.Errorf("unable to get list payslips. %s", err.Error())
	}

	if len(payslips) == 0 {
		return resp, errors.New("no payslips found for this payroll")
	}

	var summaries []EmployeePayslip
	var total float64

	for _, p := range payslips {
		user, err := uc.userRepository.GetOneByUUID(ctx, p.UserUUID)
		if err != nil {
			continue
		}

		summary := EmployeePayslip{
			UserID:     p.UserUUID,
			Username:   user.Username,
			BaseSalary: p.BaseSalary,
			WorkDays:   p.WorkDays,
			Total:      p.Total,
		}
		summaries = append(summaries, summary)
		total += p.Total
	}

	result := GetSummaryResponse{
		PayrollUUID:      req.PayrollUUID,
		TotalEmployees:   len(summaries),
		TotalTakeHomePay: total,
		Employees:        summaries,
	}

	return result, nil
}
