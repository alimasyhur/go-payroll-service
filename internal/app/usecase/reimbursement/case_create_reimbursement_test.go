package reimbursement_test

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/alimasyhur/go-payroll-service/internal/app/entity"
	mockRepo "github.com/alimasyhur/go-payroll-service/internal/app/repository/mocks"
	"gorm.io/gorm"

	"github.com/alimasyhur/go-payroll-service/internal/app/usecase/reimbursement"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCreateTemplate(t *testing.T) {

	cases := map[string]struct {
		ShouldError                 bool
		Request                     reimbursement.ReimbursementRequest
		AttendancePeriod            entity.AttendancePeriod
		Reimbursement               entity.Reimbursement
		GetOneClosedByDateResponse  error
		CreateReimbursementResponse error
		CreateAuditLogResponse      error
		Response                    error
	}{
		"ShouldErrorWhenAmountZero": {
			ShouldError: true,
			Request: reimbursement.ReimbursementRequest{
				RequestID:   "payroll-service",
				UserUUID:    "69ae7080-f05e-4ee9-98d3-2bf41aef9d2b",
				Date:        "2025-06-28",
				Amount:      0,
				IP:          "127.0.0.1",
				Description: "Beli Tiket Pesawat",
			},
			Response: fmt.Errorf("invalid amount"),
		},
		"ShouldErrorWhenInvalidDate": {
			ShouldError: true,
			Request: reimbursement.ReimbursementRequest{
				RequestID:   "payroll-service",
				UserUUID:    "69ae7080-f05e-4ee9-98d3-2bf41aef9d2b",
				Date:        "Tanggalan",
				Amount:      1000000,
				IP:          "127.0.0.1",
				Description: "Beli Tiket Pesawat",
			},
			Response: fmt.Errorf("unable to parse date"),
		},
		"ShouldErrorWhenAttendancePeriodIsClosed": {
			ShouldError: true,
			Request: reimbursement.ReimbursementRequest{
				RequestID:   "payroll-service",
				UserUUID:    "69ae7080-f05e-4ee9-98d3-2bf41aef9d2b",
				Date:        "2025-06-20",
				Amount:      1000000,
				IP:          "127.0.0.1",
				Description: "Beli Tiket Pesawat",
			},
			AttendancePeriod: entity.AttendancePeriod{
				UUID:      "69ae7080-f05e-4ee9-98d3-2bf41aef9d2b",
				StartDate: time.Now().AddDate(0, 0, -10),
				EndDate:   time.Now().AddDate(0, 0, 10),
				IsClosed:  true,
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
			},
			GetOneClosedByDateResponse: nil,
			Response:                   fmt.Errorf("attendance period is closed"),
		},
		"ShouldErrorWhenCreateReimbursement": {
			ShouldError: true,
			Request: reimbursement.ReimbursementRequest{
				RequestID:   "payroll-service",
				UserUUID:    "69ae7080-f05e-4ee9-98d3-2bf41aef9d2b",
				Date:        "2025-06-20",
				Amount:      1000000,
				IP:          "127.0.0.1",
				Description: "Beli Tiket Pesawat",
			},
			AttendancePeriod: entity.AttendancePeriod{
				UUID:      "69ae7080-f05e-4ee9-98d3-2bf41aef9d2b",
				StartDate: time.Now().AddDate(0, 0, -10),
				EndDate:   time.Now().AddDate(0, 0, 10),
				IsClosed:  false,
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
			},
			Reimbursement:               entity.Reimbursement{},
			GetOneClosedByDateResponse:  nil,
			CreateReimbursementResponse: gorm.ErrInvalidData,
			Response:                    nil,
		},
		"ShouldNotError": {
			ShouldError: false,
			Request: reimbursement.ReimbursementRequest{
				RequestID:   "payroll-service",
				UserUUID:    "69ae7080-f05e-4ee9-98d3-2bf41aef9d2b",
				Date:        "2025-06-20",
				Amount:      1000000,
				IP:          "127.0.0.1",
				Description: "Beli Tiket Pesawat",
			},
			AttendancePeriod: entity.AttendancePeriod{
				UUID:      "69ae7080-f05e-4ee9-98d3-2bf41aef9d2b",
				StartDate: time.Now().AddDate(0, 0, -10),
				EndDate:   time.Now().AddDate(0, 0, 10),
				IsClosed:  false,
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
			},
			Reimbursement:               entity.Reimbursement{UUID: "69ae7080-f05e-4ee9-98d3-2bf41aef9d2c"},
			GetOneClosedByDateResponse:  nil,
			CreateReimbursementResponse: nil,
			CreateAuditLogResponse:      nil,
			Response:                    nil,
		},
	}

	for v, test := range cases {
		t.Run(v, func(t *testing.T) {
			attendancePeriodRepository := new(mockRepo.AttendancePeriod)
			reimbursementRepository := new(mockRepo.Reimbursement)
			auditLogRepository := new(mockRepo.AuditLog)

			usecase := reimbursement.NewUsecase().
				SetAttendancePeriodRepository(attendancePeriodRepository).
				SetReimbursementRepository(reimbursementRepository).
				SetAuditLogRepository(auditLogRepository).
				Validate()

			attendancePeriodRepository.On("GetOneClosedByDate", mock.Anything, mock.Anything).Return(test.AttendancePeriod, test.GetOneClosedByDateResponse)
			reimbursementRepository.On("CreateReimbursement", mock.Anything, mock.Anything).Return(test.Reimbursement, test.CreateReimbursementResponse)
			auditLogRepository.On("CreateAuditLog", mock.Anything, mock.Anything).Return(test.CreateAuditLogResponse)

			_, err := usecase.CreateReimbursement(context.Background(), test.Request)

			if test.ShouldError {
				assert.NotNil(t, err)
			} else {
				assert.Nil(t, err)
			}
		})
	}
}
