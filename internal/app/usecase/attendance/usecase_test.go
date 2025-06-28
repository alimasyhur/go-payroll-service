package attendance_test

import (
	"testing"

	mockRepo "github.com/alimasyhur/go-payroll-service/internal/app/repository/mocks"
	"github.com/alimasyhur/go-payroll-service/internal/app/usecase/attendance"

	"github.com/stretchr/testify/assert"
)

func TestNewUsecase(t *testing.T) {

	t.Run("ShouldPanicWhenAttendancePeriodRepositoryIsNil", func(t *testing.T) {
		assert.Panics(t, func() {
			attendance.NewUsecase().Validate()
		})
	})

	t.Run("ShouldPanicWhenAttendanceRepositoryIsNil", func(t *testing.T) {
		attendancePeriodRepository := new(mockRepo.AttendancePeriod)

		assert.Panics(t, func() {
			attendance.NewUsecase().
				SetAttendancePeriodRepository(attendancePeriodRepository).
				Validate()
		})
	})

	t.Run("ShouldPanicWhenAuditLogRepositoryIsNil", func(t *testing.T) {
		attendancePeriodRepository := new(mockRepo.AttendancePeriod)
		attendanceRepository := new(mockRepo.Attendance)

		assert.Panics(t, func() {
			attendance.NewUsecase().
				SetAttendancePeriodRepository(attendancePeriodRepository).
				SetAttendanceRepository(attendanceRepository).
				Validate()
		})
	})

	t.Run("ShouldNotPanic", func(t *testing.T) {
		attendancePeriodRepository := new(mockRepo.AttendancePeriod)
		attendanceRepository := new(mockRepo.Attendance)

		auditLogRepository := new(mockRepo.AuditLog)

		assert.NotPanics(t, func() {
			attendance.NewUsecase().
				SetAttendancePeriodRepository(attendancePeriodRepository).
				SetAttendanceRepository(attendanceRepository).
				SetAuditLogRepository(auditLogRepository).
				Validate()
		})
	})
}
