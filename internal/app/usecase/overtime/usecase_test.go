package overtime_test

import (
	"testing"

	mockRepo "github.com/alimasyhur/go-payroll-service/internal/app/repository/mocks"
	"github.com/alimasyhur/go-payroll-service/internal/app/usecase/overtime"

	"github.com/stretchr/testify/assert"
)

func TestNewUsecase(t *testing.T) {

	t.Run("ShouldPanicWhenAttendancePeriodRepositoryIsNil", func(t *testing.T) {
		assert.Panics(t, func() {
			overtime.NewUsecase().Validate()
		})
	})

	t.Run("ShouldPanicWhenOvertimeRepositoryIsNil", func(t *testing.T) {
		attendancePeriodRepository := new(mockRepo.AttendancePeriod)

		assert.Panics(t, func() {
			overtime.NewUsecase().
				SetAttendancePeriodRepository(attendancePeriodRepository).
				Validate()
		})
	})

	t.Run("ShouldNotPanic", func(t *testing.T) {
		attendancePeriodRepository := new(mockRepo.AttendancePeriod)
		overtimeRepository := new(mockRepo.Overtime)

		assert.NotPanics(t, func() {
			overtime.NewUsecase().
				SetAttendancePeriodRepository(attendancePeriodRepository).
				SetOvertimeRepository(overtimeRepository).
				Validate()
		})
	})
}
