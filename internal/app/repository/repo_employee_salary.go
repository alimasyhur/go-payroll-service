package repository

import (
	"context"

	"github.com/alimasyhur/go-payroll-service/internal/app/entity"
	"github.com/alimasyhur/go-payroll-service/internal/pkg/logger"
	"gorm.io/gorm"
)

type EmployeeSalary interface {
	GetOneByUserUUID(ctx context.Context, userUUID string) (result entity.EmployeeSalary, err error)
}

type employeeSalary struct {
	tableName string
	db        *gorm.DB
}

func NewEmployeeSalaryRepository(db *gorm.DB) EmployeeSalary {
	if db == nil {
		panic("db is nil")
	}

	return &employeeSalary{
		tableName: "employee_salaries",
		db:        db,
	}
}

func (r *employeeSalary) GetOneByUserUUID(ctx context.Context, userUUID string) (result entity.EmployeeSalary, err error) {
	err = r.db.WithContext(ctx).Table(r.tableName).Select(
		"uuid",
		"user_uuid",
		"amount",
		"active",
		"effective_date",
		"created_at",
		"updated_at",
	).
		Where("user_uuid = ? AND active = TRUE", userUUID).
		Take(&result).Error

	if err != nil {
		logger.Log.Error(ctx, err.Error(), "OvertimeRepository.GetOneByUserDate", map[string]interface{}{
			"user_uuid": userUUID,
		})
		return
	}

	return
}
