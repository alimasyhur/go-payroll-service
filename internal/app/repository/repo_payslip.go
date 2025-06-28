package repository

import (
	"context"
	"fmt"

	"github.com/alimasyhur/go-payroll-service/internal/app/entity"
	"github.com/alimasyhur/go-payroll-service/internal/pkg/logger"
	"gorm.io/gorm"
)

type Payslip interface {
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

func (r *payslip) CreatePayslip(ctx context.Context, payload entity.Payslip) (entity.Payslip, error) {
	err := r.db.WithContext(ctx).Table(r.tableName).Create(&payload).Error
	if err != nil {
		logger.Log.Error(ctx, fmt.Sprintf("PayslipRepository.CreatePayslip error: %s", err.Error()), payload)
		err = fmt.Errorf("PayslipRepository.CreatePayslip error: %s", err.Error())
		return entity.Payslip{}, err
	}

	return payload, nil
}
