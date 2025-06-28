package repository

import (
	"context"
	"fmt"

	"github.com/alimasyhur/go-payroll-service/internal/app/entity"
	"github.com/alimasyhur/go-payroll-service/internal/pkg/logger"
	"gorm.io/gorm"
)

type Reimbursement interface {
	CreateReimbursement(ctx context.Context, payload entity.Reimbursement) (result entity.Reimbursement, err error)
	GetListByUserDaterange(ctx context.Context, userUUID, startDate, endDate string) (results []entity.Reimbursement, err error)
}

type reimbursement struct {
	tableName string
	db        *gorm.DB
}

func NewReimbursementRepository(db *gorm.DB) Reimbursement {
	if db == nil {
		panic("db is nil")
	}

	return &reimbursement{
		tableName: "reimbursements",
		db:        db,
	}
}

func (r *reimbursement) CreateReimbursement(ctx context.Context, payload entity.Reimbursement) (entity.Reimbursement, error) {
	err := r.db.WithContext(ctx).Table(r.tableName).Create(&payload).Error
	if err != nil {
		logger.Log.Error(ctx, fmt.Sprintf("ReimbursementRepository.CreateReimbursement error: %s", err.Error()), payload)
		err = fmt.Errorf("ReimbursementRepository.CreateReimbursement error: %s", err.Error())
		return entity.Reimbursement{}, err
	}

	return payload, nil
}

func (r *reimbursement) GetListByUserDaterange(ctx context.Context, userUUID, startDate, endDate string) (results []entity.Reimbursement, err error) {
	err = r.db.WithContext(ctx).Table(r.tableName).
		Where("user_uuid = ? AND date BETWEEN ? AND ?", userUUID, startDate, endDate).
		Find(&results).Error

	if err != nil {
		logger.Log.Error(ctx, err.Error(), "OvertimeRepository.GetListByUserDaterange", map[string]interface{}{
			"user_uuid":  userUUID,
			"start_date": startDate,
			"end_date":   endDate,
		})
		return
	}

	return
}
