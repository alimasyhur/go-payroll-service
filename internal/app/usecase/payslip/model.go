package payslip

type GetOnePayslipRequest struct {
	UserUUID    string `json:"user_uuid" validate:"required"`
	PayrollUUID string `json:"payroll_uuid" validate:"required"`
}

type GetOnePayslipResponse struct {
	UUID                string              `json:"uuid"`
	PayrollUUID         string              `json:"payroll_uuid"`
	UserUUID            string              `json:"user_uuid"`
	BaseSalary          float64             `json:"base_salary"`
	DailySalary         float64             `json:"daily_salary"`
	WorkDays            int64               `json:"work_days"`
	Attendance          float64             `json:"attendance"`
	Overtime            float64             `json:"overtime"`
	OvertimeDetail      []OvertimeItem      `json:"overtime_detail"`
	Reimbursement       float64             `json:"reimbursement"`
	ReimbursementDetail []ReimbursementItem `json:"reimbursement_detail"`
	Total               float64             `json:"total"`
}

type OvertimeItem struct {
	Date  string  `json:"date"`
	Hours float64 `json:"hours"`
	Value float64 `json:"value"`
}

type ReimbursementItem struct {
	Date        string  `json:"date"`
	Amount      float64 `json:"amount"`
	Description string  `json:"description"`
}

type GetSummaryRequest struct {
	UserUUID    string `json:"user_uuid" validate:"required"`
	PayrollUUID string `json:"payroll_uuid" validate:"required"`
}

type GetSummaryResponse struct {
	PayrollUUID      string            `json:"payroll_uuid"`
	TotalEmployees   int               `json:"total_employees"`
	TotalTakeHomePay float64           `json:"total_take_home_pay"`
	Employees        []EmployeePayslip `json:"employees"`
}

type EmployeePayslip struct {
	UserID     string  `json:"user_id"`
	Username   string  `json:"username"`
	BaseSalary float64 `json:"base_salary"`
	WorkDays   int64   `json:"work_days"`
	Total      float64 `json:"total"`
}
