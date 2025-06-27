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

func (uc *usecase) CreateAttendance(ctx context.Context, req AttendanceRequest) (resp AttendanceResponse, err error) {
	now := time.Now()
	today := now.Format(time.DateOnly)

	if now.Weekday() == time.Saturday || now.Weekday() == time.Sunday {
		errMessage := "cannot submit attendance on weekend"
		return resp, errors.New(errMessage)
	}

	closedPeriod, err := uc.attendancePeriodRepository.GetOneClosedByDate(ctx, today)
	if err != nil && err != gorm.ErrRecordNotFound {
		return resp, fmt.Errorf("unable to get Period. %s", err.Error())
	}

	if closedPeriod.UUID != "" {
		return resp, fmt.Errorf("attendance period is closed")
	}

	attendanceToday, err := uc.attendanceRepository.GetOneByUserDate(ctx, req.UserUUID, today)
	if err != nil && err != gorm.ErrRecordNotFound {
		return resp, fmt.Errorf("unable to get attendance. %s", err.Error())
	}

	if attendanceToday.UUID != "" {
		emptyTime := time.Time{}.Format(time.TimeOnly)
		if attendanceToday.ClockOut != emptyTime {
			return resp, fmt.Errorf("already checkout today")
		}

		clockin := attendanceToday.ClockIn
		attendanceToday.ClockIn = clockin
		attendanceToday.ClockOut = now.Format(time.TimeOnly)
		attendanceToday.UpdatedBy = req.UserUUID
		attendanceToday.IP = req.IP
		attendanceToday.UpdatedAt = now

		if err := uc.attendanceRepository.UpdateAttendance(ctx, attendanceToday); err != nil {
			return resp, fmt.Errorf("unable to update attendance. %s", err.Error())
		}

		resp = AttendanceResponse{
			UUID:      attendanceToday.UUID,
			UserUUID:  attendanceToday.UserUUID,
			Date:      attendanceToday.Date.String(),
			ClockIn:   attendanceToday.ClockIn,
			ClockOut:  attendanceToday.ClockOut,
			IP:        attendanceToday.IP,
			CreatedBy: attendanceToday.CreatedBy,
			UpdatedBy: attendanceToday.UpdatedBy,
			CreatedAt: attendanceToday.CreatedAt,
			UpdatedAt: attendanceToday.UpdatedAt,
		}
		return resp, nil
	}

	payloadAttendance := entity.Attendance{
		UUID:      uuid.New().String(),
		UserUUID:  req.UserUUID,
		Date:      now,
		ClockIn:   now.Format(time.TimeOnly),
		ClockOut:  time.Time{}.Format(time.TimeOnly),
		IP:        req.IP,
		CreatedBy: req.UserUUID,
		UpdatedBy: req.UserUUID,
		CreatedAt: now,
		UpdatedAt: now,
	}

	attendance, err := uc.attendanceRepository.CreateAttendance(ctx, payloadAttendance)
	if err != nil {
		return resp, fmt.Errorf("unable to create attendance. %s", err.Error())
	}

	resp = AttendanceResponse{
		UUID:      attendance.UUID,
		UserUUID:  attendance.UserUUID,
		Date:      attendance.Date.String(),
		ClockIn:   attendance.ClockIn,
		ClockOut:  attendance.ClockOut,
		IP:        attendance.IP,
		CreatedBy: attendance.CreatedBy,
		UpdatedBy: attendance.UpdatedBy,
		CreatedAt: attendance.CreatedAt,
		UpdatedAt: attendance.UpdatedAt,
	}

	return resp, nil
}
