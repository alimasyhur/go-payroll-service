package repository

import (
	"context"
	"fmt"

	"github.com/alimasyhur/go-payroll-service/internal/app/entity"
	"github.com/alimasyhur/go-payroll-service/internal/pkg/logger"
	"gorm.io/gorm"
)

type Payroll interface {
	GetOneByPeriodUUID(ctx context.Context, periodUUID string) (result entity.Payroll, err error)
	CreatePayroll(ctx context.Context, payload entity.Payroll) (result entity.Payroll, err error)
}

type payroll struct {
	tableName string
	db        *gorm.DB
}

func NewPayrollRepository(db *gorm.DB) Payroll {
	if db == nil {
		panic("db is nil")
	}

	return &payroll{
		tableName: "payrolls",
		db:        db,
	}
}

func (r *payroll) GetOneByPeriodUUID(ctx context.Context, periodUUID string) (result entity.Payroll, err error) {
	err = r.db.WithContext(ctx).Table(r.tableName).Select(
		"uuid",
		"attendance_period_uuid",
		"processed_at",
		"ip",
		"created_by",
		"created_at",
		"updated_at",
	).
		Where("attendance_period_uuid = ?", periodUUID).
		Take(&result).Error

	if err != nil {
		logger.Log.Error(ctx, err.Error(), "PayrollRepository.GetOneByPeriodUUID", map[string]interface{}{
			"attendance_period_uuid": periodUUID,
		})
		return
	}

	return
}

func (r *payroll) CreatePayroll(ctx context.Context, payload entity.Payroll) (entity.Payroll, error) {
	err := r.db.WithContext(ctx).Table(r.tableName).Create(&payload).Error
	if err != nil {
		logger.Log.Error(ctx, fmt.Sprintf("PayrollRepository.CreatePayroll error: %s", err.Error()), payload)
		err = fmt.Errorf("PayrollRepository.CreatePayroll error: %s", err.Error())
		return entity.Payroll{}, err
	}

	return payload, nil
}
