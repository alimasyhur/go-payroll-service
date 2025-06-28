package entity

import (
	"time"
)

type Reimbursement struct {
	UUID        string    `gorm:"column:uuid;type:uuid;default:uuid_generate_v4();primaryKey" json:"uuid"`
	UserUUID    string    `gorm:"column:user_uuid" json:"user_uuid"`
	Amount      float64   `gorm:"column:amount" json:"amount"`
	Date        time.Time `gorm:"column:date" json:"date"`
	Description string    `gorm:"column:description" json:"description"`
	IP          string    `gorm:"column:ip" json:"ip"`
	CreatedBy   string    `gorm:"column:created_by" json:"created_by"`
	UpdatedBy   string    `gorm:"column:updated_by" json:"updated_by"`
	CreatedAt   time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt   time.Time `gorm:"column:updated_at" json:"updated_at"`
}

func (e *Reimbursement) TableName() string {
	return "reimbursements"
}
