package overtime

import (
	"context"

	"github.com/alimasyhur/go-payroll-service/internal/app/repository"
)

type OvertimeUsecase interface {
	CreateOvertime(ctx context.Context, req OvertimeRequest) (resp OvertimeResponse, err error)
}
type usecase struct {
	attendancePeriodRepository repository.AttendancePeriod
	overtimeRepository         repository.Overtime
}

func NewUsecase() *usecase {
	return &usecase{}
}

func (uc *usecase) SetAttendancePeriodRepository(r repository.AttendancePeriod) *usecase {
	uc.attendancePeriodRepository = r
	return uc
}

func (uc *usecase) SetOvertimeRepository(r repository.Overtime) *usecase {
	uc.overtimeRepository = r
	return uc
}

func (uc *usecase) Validate() OvertimeUsecase {
	if uc.attendancePeriodRepository == nil {
		panic("attendance period repository is nil")
	}
	if uc.overtimeRepository == nil {
		panic("overtime repository is nil")
	}

	return uc
}
