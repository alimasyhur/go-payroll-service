package attendance

import (
	"context"

	"github.com/alimasyhur/go-payroll-service/internal/app/repository"
)

type AttendanceUsecase interface {
	CreateAttendancePeriod(ctx context.Context, req AttendancePeriodRequest) (resp AttendancePeriodResponse, err error)
	CreateAttendance(ctx context.Context, req AttendanceRequest) (resp AttendanceResponse, err error)
}
type usecase struct {
	attendancePeriodRepository repository.AttendancePeriod
	attendanceRepository       repository.Attendance
}

func NewUsecase() *usecase {
	return &usecase{}
}

func (uc *usecase) SetAttendancePeriodRepository(r repository.AttendancePeriod) *usecase {
	uc.attendancePeriodRepository = r
	return uc
}

func (uc *usecase) SetAttendanceRepository(r repository.Attendance) *usecase {
	uc.attendanceRepository = r
	return uc
}

func (uc *usecase) Validate() AttendanceUsecase {
	if uc.attendancePeriodRepository == nil {
		panic("attendance period repository is nil")
	}
	if uc.attendanceRepository == nil {
		panic("attendance repository is nil")
	}

	return uc
}
