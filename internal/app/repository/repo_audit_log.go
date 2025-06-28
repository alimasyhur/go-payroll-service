package repository

import (
	"context"
	"fmt"

	"github.com/alimasyhur/go-payroll-service/internal/app/entity"
	"github.com/alimasyhur/go-payroll-service/internal/pkg/logger"
	"gorm.io/gorm"
)

type AuditLog interface {
	CreateAuditLog(ctx context.Context, payload entity.AuditLog) (err error)
}

type auditLog struct {
	tableName string
	db        *gorm.DB
}

func NewAuditLogRepository(db *gorm.DB) AuditLog {
	if db == nil {
		panic("db is nil")
	}

	return &auditLog{
		tableName: "audit_logs",
		db:        db,
	}
}

func (r *auditLog) CreateAuditLog(ctx context.Context, payload entity.AuditLog) error {
	err := r.db.WithContext(ctx).Table(r.tableName).Create(&payload).Error
	if err != nil {
		logger.Log.Error(ctx, fmt.Sprintf("AuditLogRepository.CreateAuditLog error: %s", err.Error()), payload)
		err = fmt.Errorf("AuditLogRepository.CreateAuditLog error: %s", err.Error())
		return err
	}

	return nil
}
