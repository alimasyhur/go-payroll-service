package entity

import (
	"time"
)

type Overtime struct {
	UUID      string    `gorm:"column:uuid;type:uuid;default:uuid_generate_v4();primaryKey" json:"uuid"`
	UserUUID  string    `gorm:"column:user_uuid"`
	Date      time.Time `gorm:"column:date"`
	Hours     float64   `gorm:"column:hours" json:"hours"`
	IP        string    `gorm:"column:ip" json:"ip"`
	CreatedBy string    `gorm:"column:created_by" json:"created_by"`
	UpdatedBy string    `gorm:"column:updated_by" json:"updated_by"`
	CreatedAt time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at" json:"updated_at"`
}

func (e *Overtime) TableName() string {
	return "overtimes"
}
