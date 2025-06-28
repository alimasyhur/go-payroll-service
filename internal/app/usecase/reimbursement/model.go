package reimbursement

import "time"

type ReimbursementRequest struct {
	UserUUID    string  `json:"user_uuid" validate:"required"`
	Date        string  `json:"date" validate:"required"`
	Amount      float64 `json:"amount" validate:"required"`
	IP          string  `json:"ip" validate:"required"`
	Description string  `json:"description"`
}

type ReimbursementResponse struct {
	UUID        string    `json:"uuid"`
	UserUUID    string    `json:"user_uuid"`
	Amount      float64   `json:"amount"`
	Date        time.Time `json:"date"`
	Description string    `json:"description"`
	IP          string    `json:"ip"`
	CreatedBy   string    `json:"created_by"`
	UpdatedBy   string    `json:"updated_by"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
