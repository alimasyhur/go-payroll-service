package repository

import (
	"context"
	"fmt"

	"github.com/alimasyhur/go-payroll-service/internal/app/entity"
	"github.com/alimasyhur/go-payroll-service/internal/pkg/logger"
	"gorm.io/gorm"
)

type Attendance interface {
	CreateAttendance(ctx context.Context, payload entity.Attendance) (result entity.Attendance, err error)
	GetOneByUserDate(ctx context.Context, userUUID, date string) (result entity.Attendance, err error)
	UpdateAttendance(ctx context.Context, data entity.Attendance) error
	GetWorkdaysByUserDaterange(ctx context.Context, userUUID, startDate, endDate string) (total int64, err error)
}

type attendance struct {
	tableName string
	db        *gorm.DB
}

func NewAttendanceRepository(db *gorm.DB) Attendance {
	if db == nil {
		panic("db is nil")
	}

	return &attendance{
		tableName: "attendances",
		db:        db,
	}
}

func (r *attendance) GetOneByUserDate(ctx context.Context, userUUID, date string) (result entity.Attendance, err error) {
	err = r.db.WithContext(ctx).Table(r.tableName).Select(
		"uuid",
		"date",
		"clockin",
		"clockout",
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
		logger.Log.Error(ctx, err.Error(), "AttendanceRepository.GetOneByUserDate", map[string]interface{}{
			"date":      date,
			"user_uuid": userUUID,
		})
		return
	}

	return
}

func (r *attendance) CreateAttendance(ctx context.Context, payload entity.Attendance) (entity.Attendance, error) {
	err := r.db.WithContext(ctx).Table(r.tableName).Create(&payload).Error
	if err != nil {
		logger.Log.Error(ctx, fmt.Sprintf("AttendanceRepository.CreateAttendance error: %s", err.Error()), payload)
		err = fmt.Errorf("AttendanceRepository.CreaCreateAttendance error: %s", err.Error())
		return entity.Attendance{}, err
	}

	return payload, nil
}

func (r *attendance) UpdateAttendance(ctx context.Context, data entity.Attendance) error {
	err := r.db.WithContext(ctx).Table(r.tableName).Save(&data).Error

	if err != nil {
		logger.Log.Error(ctx, "AttendanceRepository.UpdateAttendance", map[string]interface{}{
			"data": data,
		})
		return err
	}

	return nil
}

func (r *attendance) GetWorkdaysByUserDaterange(ctx context.Context, userUUID, startDate, endDate string) (total int64, err error) {
	err = r.db.WithContext(ctx).Table(r.tableName).
		Where("user_uuid = ? AND date BETWEEN ? AND ?", userUUID, startDate, endDate).
		Count(&total).Error

	if err != nil {
		logger.Log.Error(ctx, err.Error(), "AttendanceRepository.GetCountByUserDaterange", map[string]interface{}{
			"user_uuid":  userUUID,
			"start_date": startDate,
			"end_date":   endDate,
		})
		return
	}

	return
}
