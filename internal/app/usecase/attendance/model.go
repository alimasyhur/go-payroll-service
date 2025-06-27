package attendance

import "time"

type AttendancePeriodRequest struct {
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
