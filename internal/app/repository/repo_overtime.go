package repository

import (
	"context"
	"fmt"

	"github.com/alimasyhur/go-payroll-service/internal/app/entity"
	"github.com/alimasyhur/go-payroll-service/internal/pkg/logger"
	"gorm.io/gorm"
)

type Overtime interface {
	CreateOvertime(ctx context.Context, payload entity.Overtime) (result entity.Overtime, err error)
	GetOneByUserDate(ctx context.Context, userUUID, date string) (result entity.Overtime, err error)
	UpdateOvertime(ctx context.Context, data entity.Overtime) error
}

type overtime struct {
	tableName string
	db        *gorm.DB
}

func NewOvertimeRepository(db *gorm.DB) Overtime {
	if db == nil {
		panic("db is nil")
	}

	return &overtime{
		tableName: "overtimes",
		db:        db,
	}
}

func (r *overtime) GetOneByUserDate(ctx context.Context, userUUID, date string) (result entity.Overtime, err error) {
	err = r.db.WithContext(ctx).Table(r.tableName).Select(
		"uuid",
		"date",
		"hours",
		"user_uuid",
		"ip",
		"created_by",
		"updated_by",
		"created_at",
		"updated_at",
	).
		Where("date = ? AND user_uuid = ? ", date, userUUID).
		Take(&result).Error

	if err != nil {
		logger.Log.Error(ctx, err.Error(), "OvertimeRepository.GetOneByUserDate", map[string]interface{}{
			"date":      date,
			"user_uuid": userUUID,
		})
		return
	}

	return
}

func (r *overtime) CreateOvertime(ctx context.Context, payload entity.Overtime) (entity.Overtime, error) {
	err := r.db.WithContext(ctx).Table(r.tableName).Create(&payload).Error
	if err != nil {
		logger.Log.Error(ctx, fmt.Sprintf("OvertimeRepository.CreateOvertime error: %s", err.Error()), payload)
		err = fmt.Errorf("OvertimeRepository.CreaCreateAttCreateOvertimeendance error: %s", err.Error())
		return entity.Overtime{}, err
	}

	return payload, nil
}

func (r *overtime) UpdateOvertime(ctx context.Context, data entity.Overtime) error {
	err := r.db.WithContext(ctx).Table(r.tableName).Save(&data).Error

	if err != nil {
		logger.Log.Error(ctx, "OvertimeRepository.UpdateUpdateOvertimeAttendance", map[string]interface{}{
			"data": data,
		})
		return err
	}

	return nil
}
