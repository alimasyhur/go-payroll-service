package entity

import (
	"time"
)

type Payroll struct {
	UUID                 string    `gorm:"column:uuid;type:uuid;default:uuid_generate_v4();primaryKey" json:"uuid"`
	AttendancePeriodUUID string    `gorm:"column:attendance_period_uuid"`
	ProcessedAt          time.Time `gorm:"column:processed_at"`
	IP                   string    `gorm:"column:ip" json:"ip"`
	CreatedBy            string    `gorm:"column:created_by" json:"created_by"`
	CreatedAt            time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt            time.Time `gorm:"column:updated_at" json:"updated_at"`
}

func (e *Payroll) TableName() string {
	return "payrolls"
}
