package entity

import (
	"time"
)

type UserSalary struct {
	UUID          string    `gorm:"column:uuid;type:uuid;default:uuid_generate_v4();primaryKey"`
	UserUUID      string    `gorm:"column:user_uuid;type:uuid;not null;index"`
	Amount        float64   `gorm:"column:amount;not null"`
	Active        bool      `gorm:"column:active;default:true"`
	EffectiveDate time.Time `gorm:"column:effective_date" json:"effective_date"`
	CreatedAt     time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt     time.Time `gorm:"column:updated_at" json:"updated_at"`
}
