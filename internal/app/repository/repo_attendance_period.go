package repository

import (
	"context"
	"fmt"

	"github.com/alimasyhur/go-payroll-service/internal/app/entity"
	"github.com/alimasyhur/go-payroll-service/internal/pkg/logger"
	"gorm.io/gorm"
)

type AttendancePeriod interface {
	GetOneByUUID(ctx context.Context, uuid string) (result entity.AttendancePeriod, err error)
	GetOneByDaterange(ctx context.Context, startDate, endDate string) (result entity.AttendancePeriod, err error)
	GetOneClosedByDate(ctx context.Context, date string) (result entity.AttendancePeriod, err error)
	CreateAttendancePeriod(ctx context.Context, payload entity.AttendancePeriod) (result entity.AttendancePeriod, err error)
	UpdateAttendancePeriod(ctx context.Context, data entity.AttendancePeriod) error
}

type attendancePeriod struct {
	tableName string
	db        *gorm.DB
}

func NewAttendancePeriodRepository(db *gorm.DB) AttendancePeriod {
	if db == nil {
		panic("db is nil")
	}

	return &attendancePeriod{
		tableName: "attendance_periods",
		db:        db,
	}
}

func (r *attendancePeriod) GetOneByUUID(ctx context.Context, uuid string) (result entity.AttendancePeriod, err error) {
	err = r.db.WithContext(ctx).Table(r.tableName).Select(
		"uuid",
		"start_date",
		"end_date",
		"is_closed",
		"created_at",
		"updated_at",
	).
		Where("uuid = ?", uuid).
		Take(&result).Error

	if err != nil {
		logger.Log.Error(ctx, err.Error(), "AttendancePeriodRepository.GetOneByUUID", map[string]interface{}{
			"uuid": uuid,
		})
		return
	}

	return
}

func (r *attendancePeriod) GetOneByDaterange(ctx context.Context, startDate, endDate string) (result entity.AttendancePeriod, err error) {
	err = r.db.WithContext(ctx).Table(r.tableName).Select(
		"uuid",
		"start_date",
		"end_date",
		"is_closed",
		"created_at",
		"updated_at",
	).
		Where("start_date = ? AND end_date = ?", startDate, endDate).
		Take(&result).Error

	if err != nil {
		logger.Log.Error(ctx, err.Error(), "AttendancePeriodRepository.GetOneByDaterange", map[string]interface{}{
			"start_date": startDate,
			"end_date":   endDate,
		})
		return
	}

	return
}

func (r *attendancePeriod) GetOneClosedByDate(ctx context.Context, date string) (result entity.AttendancePeriod, err error) {
	err = r.db.WithContext(ctx).Table(r.tableName).Select(
		"uuid",
		"start_date",
		"end_date",
		"is_closed",
		"created_at",
		"updated_at",
	).
		Where("? BETWEEN start_date AND end_date AND is_closed = true", date).
		Take(&result).Error

	if err != nil {
		logger.Log.Error(ctx, err.Error(), "AttendancePeriodRepository.GetOneClosedByDaterange", map[string]interface{}{
			"date": date,
		})
		return
	}

	return
}

func (r *attendancePeriod) CreateAttendancePeriod(ctx context.Context, payload entity.AttendancePeriod) (entity.AttendancePeriod, error) {
	err := r.db.WithContext(ctx).Table(r.tableName).Create(&payload).Error
	if err != nil {
		logger.Log.Error(ctx, fmt.Sprintf("AttendancePeriodRepository.CreateAttendancePeriod error: %s", err.Error()), payload)
		err = fmt.Errorf("AttendancePeriodRepository.CreaCreateAttendancePeriodteUser error: %s", err.Error())
		return entity.AttendancePeriod{}, err
	}

	return payload, nil
}

func (r *attendancePeriod) UpdateAttendancePeriod(ctx context.Context, data entity.AttendancePeriod) error {
	err := r.db.WithContext(ctx).Table(r.tableName).Save(&data).Error

	if err != nil {
		logger.Log.Error(ctx, "AttendancePeriodRepository.UpdateAttendancePeriod", map[string]interface{}{
			"data": data,
		})
		return err
	}

	return nil
}
