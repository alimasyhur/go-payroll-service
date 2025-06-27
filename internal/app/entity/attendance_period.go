package entity

import (
	"time"
)

type AttendancePeriod struct {
	UUID      string    `gorm:"type:uuid;default:uuid_generate_v4();primaryKey" json:"uuid"`
	StartDate time.Time `gorm:"column:start_date" json:"start_date"`
	EndDate   time.Time `gorm:"column:end_date" json:"end_date"`
	IsClosed  bool      `gorm:"column:is_closed" json:"is_closed"`
	CreatedAt time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at" json:"updated_at"`
}

func (e *AttendancePeriod) TableName() string {
	return "attendance_periods"
}
