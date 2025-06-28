package entity

import (
	"time"
)

type AuditLog struct {
	UUID       string    `gorm:"column:uuid;type:uuid;default:uuid_generate_v4();primaryKey" json:"uuid"`
	UserUUID   string    `gorm:"column:user_uuid"`
	Action     string    `gorm:"column:action"`
	Entity     string    `gorm:"column:entity"`
	EntityUUID string    `gorm:"column:entity_uuid"`
	IP         string    `gorm:"column:ip"`
	RequestID  string    `gorm:"column:request_id"`
	CreatedAt  time.Time `gorm:"column:created_at" json:"created_at"`
}

func (e *AuditLog) TableName() string {
	return "audit_logs"
}
