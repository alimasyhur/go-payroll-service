package attendance

import "time"

type AttendancePeriodRequest struct {
	UserUUID  string `json:"user_uuid"`
	IP        string `json:"ip"`
	RequestID string `json:"request_id"`
	StartDate string `json:"start_date" validate:"required"`
	EndDate   string `json:"end_date" validate:"required"`
}

type AttendancePeriodResponse struct {
	UUID      string    `json:"uuid"`
	StartDate time.Time `json:"start_date"`
	EndDate   time.Time `json:"end_date"`
	IsClosed  bool      `json:"is_closed"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type AttendanceRequest struct {
	UserUUID string `json:"user_uuid" validate:"required"`
	IP       string `json:"ip" validate:"required"`
}

type AttendanceResponse struct {
	UUID      string    `json:"uuid"`
	UserUUID  string    `json:"user_uuid"`
	Date      string    `json:"date"`
	ClockIn   string    `json:"clockin"`
	ClockOut  string    `json:"clockout"`
	IP        string    `json:"ip"`
	CreatedBy string    `json:"created_by"`
	UpdatedBy string    `json:"updated_by"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
