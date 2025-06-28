package payroll

type CreatePayrollRequest struct {
	PeriodUUID string `json:"period_uuid" validate:"required"`
	UserUUID   string `json:"user_uuid" validate:"required"`
	IP         string `json:"ip" validate:"required"`
	RequestID  string `json:"request_id" validate:"required"`
}

type CreatePayrollResponse struct {
	UUID string `json:"uuid"`
}

type GeneratePayslipRequest struct {
	UserUUID    string `json:"user_uuid" validate:"required"`
	PayrollUUID string `json:"payroll_uuid" validate:"required"`
	PeriodUUID  string `json:"period_uuid" validate:"required"`
	StartDate   string `json:"start_date" validate:"required"`
	EndDate     string `json:"end_date" validate:"required"`
}
