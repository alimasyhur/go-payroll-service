package entity

import (
	"time"
)

type Payslip struct {
	UUID        string    `gorm:"column:uuid;type:uuid;default:uuid_generate_v4();primaryKey" json:"uuid"`
	PayrollUUID string    `gorm:"column:payroll_uuid"`
	UserUUID    string    `gorm:"column:user_uuid"`
	WorkDays    int64     `gorm:"column:work_days"`
	BaseSalary  float64   `gorm:"column:base_salary" json:"base_salary"`
	Overtime    float64   `gorm:"column:overtime" json:"overtime"`
	Reimburse   float64   `gorm:"column:reimburse" json:"reimburse"`
	Total       float64   `gorm:"column:total" json:"total"`
	CreatedAt   time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt   time.Time `gorm:"column:updated_at" json:"updated_at"`
}

func (e *Payslip) TableName() string {
	return "payslips"
}
