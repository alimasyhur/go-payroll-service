package repository

import (
	"context"
	"fmt"

	"github.com/alimasyhur/go-payroll-service/internal/app/entity"
	"github.com/alimasyhur/go-payroll-service/internal/pkg/logger"
	"gorm.io/gorm"
)

type Payslip interface {
	GetOneByUserPayrollUUID(ctx context.Context, userUUID, payrollUUID string) (result entity.Payslip, err error)
	GetListByPayrollUUID(ctx context.Context, payrollUUID string) (results []entity.Payslip, err error)
	GetOneDetailByPayrollUUID(ctx context.Context, payrollUUID string) (result entity.PayslipDetail, err error)
	CreatePayslip(ctx context.Context, payload entity.Payslip) (result entity.Payslip, err error)
}

type payslip struct {
	tableName string
	db        *gorm.DB
}

func NewPayslipRepository(db *gorm.DB) Payslip {
	if db == nil {
		panic("db is nil")
	}

	return &payslip{
		tableName: "payslips",
		db:        db,
	}
}

func (r *payslip) GetOneDetailByPayrollUUID(ctx context.Context, payrollUUID string) (result entity.PayslipDetail, err error) {
	err = r.db.WithContext(ctx).Table(r.tableName).Select(
		"payslips.uuid AS uuid",
		"payroll_uuid AS payroll_uuid",
		"payrolls.attendance_period_uuid AS period_uuid",
		"attendance_periods.start_date AS period_start_date",
		"attendance_periods.end_date AS period_end_date",
	).
		Joins(`JOIN payrolls ON payslips.payroll_uuid = payrolls.uuid`).
		Joins(`JOIN attendance_periods ON payrolls.attendance_period_uuid = attendance_periods.uuid`).
		Where("payslips.payroll_uuid = ?", payrollUUID).
		Take(&result).Error

	if err != nil {
		logger.Log.Error(ctx, err.Error(), "PayslipRepository.GetOneDetailByPayrollUUID", map[string]interface{}{
			"payroll_uuid": payrollUUID,
		})
		return
	}

	return
}

func (r *payslip) GetOneByUserPayrollUUID(ctx context.Context, userUUID, payrollUUID string) (result entity.Payslip, err error) {
	err = r.db.WithContext(ctx).Table(r.tableName).Select(
		"uuid",
		"payroll_uuid",
		"user_uuid",
		"work_days",
		"base_salary",
		"overtime",
		"reimburse",
		"total",
		"created_at",
		"updated_at",
	).
		Where("user_uuid = ? AND payroll_uuid = ?", userUUID, payrollUUID).
		Take(&result).Error

	if err != nil {
		logger.Log.Error(ctx, err.Error(), "PayslipRepository.GetOneByUserPayrollUUID", map[string]interface{}{
			"payroll_uuid": payrollUUID,
			"user_uuid":    userUUID,
		})
		return
	}

	return
}

func (r *payslip) GetListByPayrollUUID(ctx context.Context, payrollUUID string) (results []entity.Payslip, err error) {
	err = r.db.WithContext(ctx).Table(r.tableName).Select(
		"uuid",
		"payroll_uuid",
		"user_uuid",
		"work_days",
		"base_salary",
		"overtime",
		"reimburse",
		"total",
		"created_at",
		"updated_at",
	).
		Where("payroll_uuid = ?", payrollUUID).
		Find(&results).Error

	if err != nil {
		logger.Log.Error(ctx, err.Error(), "PayslipRepository.GetListByPayrollUUID", map[string]interface{}{
			"payroll_uuid": payrollUUID,
		})
		return
	}

	return
}

func (r *payslip) CreatePayslip(ctx context.Context, payload entity.Payslip) (entity.Payslip, error) {
	err := r.db.WithContext(ctx).Table(r.tableName).Create(&payload).Error
	if err != nil {
		logger.Log.Error(ctx, fmt.Sprintf("PayslipRepository.CreatePayslip error: %s", err.Error()), payload)
		err = fmt.Errorf("PayslipRepository.CreatePayslip error: %s", err.Error())
		return entity.Payslip{}, err
	}

	return payload, nil
}
