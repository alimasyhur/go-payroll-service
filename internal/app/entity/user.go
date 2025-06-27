package entity

import (
	"time"
)

type User struct {
	UUID      string    `gorm:"column:uuid;type:uuid;default:uuid_generate_v4();primaryKey" json:"uuid"`
	Username  string    `gorm:"column:username;uniqueIndex" json:"username"`
	Password  string    `gorm:"column:password" json:"password"`
	Role      string    `gorm:"column:role;type:enum('admin', 'employee')" json:"role"`
	CreatedAt time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at" json:"updated_at"`
}

func (e *User) TableName() string {
	return "users"
}
