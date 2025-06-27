package overtime

import "time"

type OvertimeRequest struct {
	UserUUID string  `json:"user_uuid" validate:"required"`
	IP       string  `json:"ip" validate:"required"`
	Hours    float64 `json:"hours" validate:"required"`
	Date     string  `json:"date" validate:"required"`
}

type OvertimeResponse struct {
	UUID      string    `json:"uuid"`
	UserUUID  string    `json:"user_uuid"`
	Date      time.Time `json:"date"`
	Hours     float64   `json:"hours"`
	IP        string    `json:"ip"`
	CreatedBy string    `json:"created_by"`
	UpdatedBy string    `json:"updated_by"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
