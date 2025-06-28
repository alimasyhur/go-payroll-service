package reimbursement_test

import (
	"testing"

	mockRepo "github.com/alimasyhur/go-payroll-service/internal/app/repository/mocks"
	"github.com/alimasyhur/go-payroll-service/internal/app/usecase/reimbursement"

	"github.com/stretchr/testify/assert"
)

func TestNewUsecase(t *testing.T) {

	t.Run("ShouldPanicWhenAttendancePeriodRepositoryIsNil", func(t *testing.T) {
		assert.Panics(t, func() {
			reimbursement.NewUsecase().Validate()
		})
	})

	t.Run("ShouldPanicWhenReimbursementRepositoryIsNil", func(t *testing.T) {
		attendancePeriodRepository := new(mockRepo.AttendancePeriod)

		assert.Panics(t, func() {
			reimbursement.NewUsecase().
				SetAttendancePeriodRepository(attendancePeriodRepository).
				Validate()
		})
	})

	t.Run("ShouldPanicWhenAuditLogRepositoryIsNil", func(t *testing.T) {
		attendancePeriodRepository := new(mockRepo.AttendancePeriod)
		reimbursementRepository := new(mockRepo.Reimbursement)

		assert.Panics(t, func() {
			reimbursement.NewUsecase().
				SetAttendancePeriodRepository(attendancePeriodRepository).
				SetReimbursementRepository(reimbursementRepository).
				Validate()
		})
	})

	t.Run("ShouldNotPanic", func(t *testing.T) {
		attendancePeriodRepository := new(mockRepo.AttendancePeriod)
		reimbursementRepository := new(mockRepo.Reimbursement)
		auditLogRepository := new(mockRepo.AuditLog)

		assert.NotPanics(t, func() {
			reimbursement.NewUsecase().
				SetAttendancePeriodRepository(attendancePeriodRepository).
				SetReimbursementRepository(reimbursementRepository).
				SetAuditLogRepository(auditLogRepository).
				Validate()
		})
	})
}
