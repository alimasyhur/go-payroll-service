package overtime

import (
	"context"
	"fmt"
	"time"

	"github.com/alimasyhur/go-payroll-service/internal/app/entity"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func (uc *usecase) CreateOvertime(ctx context.Context, req OvertimeRequest) (resp OvertimeResponse, err error) {
	if req.Hours <= 0 || req.Hours > 3 {
		return resp, fmt.Errorf("overtime exceededs limit")
	}

	now := time.Now()
	_, err = time.Parse(time.DateOnly, req.Date)
	if err != nil {
		return resp, fmt.Errorf("unable to parse date. %s", err.Error())
	}

	closedPeriod, err := uc.attendancePeriodRepository.GetOneClosedByDate(ctx, req.Date)
	if err != nil && err != gorm.ErrRecordNotFound {
		return resp, fmt.Errorf("unable to get Period. %s", err.Error())
	}

	if closedPeriod.UUID != "" {
		return resp, fmt.Errorf("attendance period is closed")
	}

	overtimeToday, err := uc.overtimeRepository.GetOneByUserDate(ctx, req.UserUUID, req.Date)
	if err != nil && err != gorm.ErrRecordNotFound {
		return resp, fmt.Errorf("unable to get overtime. %s", err.Error())
	}

	if overtimeToday.UUID != "" {
		return resp, fmt.Errorf("overtime was submitted this day")
	}

	overtime := entity.Overtime{
		UUID:      uuid.New().String(),
		UserUUID:  req.UserUUID,
		Date:      now,
		Hours:     req.Hours,
		CreatedBy: req.UserUUID,
		UpdatedBy: req.UserUUID,
		IP:        req.IP,
		CreatedAt: now,
		UpdatedAt: now,
	}

	newOvertime, err := uc.overtimeRepository.CreateOvertime(ctx, overtime)
	if err != nil {
		return resp, fmt.Errorf("unable to create overtime. %s", err.Error())
	}

	resp = OvertimeResponse{
		UUID:      newOvertime.UUID,
		UserUUID:  newOvertime.UserUUID,
		Date:      newOvertime.Date,
		Hours:     newOvertime.Hours,
		IP:        newOvertime.IP,
		CreatedBy: newOvertime.CreatedBy,
		UpdatedBy: newOvertime.UpdatedBy,
		CreatedAt: newOvertime.CreatedAt,
		UpdatedAt: newOvertime.UpdatedAt,
	}

	return resp, nil
}
