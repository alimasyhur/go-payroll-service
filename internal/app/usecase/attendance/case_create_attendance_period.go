package attendance

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/alimasyhur/go-payroll-service/internal/app/entity"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func (uc *usecase) CreateAttendancePeriod(ctx context.Context, req AttendancePeriodRequest) (resp AttendancePeriodResponse, err error) {
	startDateFormat, err := time.Parse(time.DateOnly, req.StartDate)
	if err != nil {
		errMessage := fmt.Sprintf("error parsing start date: %s", err.Error())
		return resp, errors.New(errMessage)
	}
	endDateFormat, err := time.Parse(time.DateOnly, req.EndDate)
	if err != nil {
		errMessage := fmt.Sprintf("error parsing end date: %s", err.Error())
		return resp, errors.New(errMessage)
	}

	existPeriod, err := uc.attendancePeriodRepository.GetOneByDaterange(ctx, req.StartDate, req.EndDate)
	if err != nil && err != gorm.ErrRecordNotFound {
		return resp, fmt.Errorf("invalid start_date and end_date. %s", err.Error())
	}

	if existPeriod.UUID != "" {
		return resp, fmt.Errorf("attendance period is exist")
	}

	period := entity.AttendancePeriod{
		UUID:      uuid.New().String(),
		StartDate: startDateFormat,
		EndDate:   endDateFormat,
		IsClosed:  false,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	if _, err = uc.attendancePeriodRepository.CreateAttendancePeriod(ctx, period); err != nil {
		return resp, err
	}

	resp = AttendancePeriodResponse{
		UUID:      uuid.New().String(),
		StartDate: startDateFormat,
		EndDate:   endDateFormat,
		IsClosed:  false,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	return resp, nil
}
