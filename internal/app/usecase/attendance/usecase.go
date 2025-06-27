package attendance

import (
	"context"

	"github.com/alimasyhur/go-payroll-service/internal/app/repository"
)

type AttendanceUsecase interface {
	CreateAttendancePeriod(ctx context.Context, req AttendancePeriodRequest) (resp AttendancePeriodResponse, err error)
}
type usecase struct {
	attendancePeriodRepository repository.AttendancePeriod
}

func NewUsecase() *usecase {
	return &usecase{}
}

func (uc *usecase) SetAttendancePeriodRepository(r repository.AttendancePeriod) *usecase {
	uc.attendancePeriodRepository = r
	return uc
}

func (uc *usecase) Validate() AttendanceUsecase {
	if uc.attendancePeriodRepository == nil {
		panic("attendance period repository is nil")
	}

	return uc
}
