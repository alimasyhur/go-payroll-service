package reimbursement

import (
	"context"
	"fmt"
	"time"

	"github.com/alimasyhur/go-payroll-service/internal/app/entity"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func (uc *usecase) CreateReimbursement(ctx context.Context, req ReimbursementRequest) (resp ReimbursementResponse, err error) {
	if req.Amount <= 0 {
		return resp, fmt.Errorf("invalid amount")
	}

	now := time.Now()
	today := now.Format(time.DateOnly)

	closedPeriod, err := uc.attendancePeriodRepository.GetOneClosedByDate(ctx, today)
	if err != nil && err != gorm.ErrRecordNotFound {
		return resp, fmt.Errorf("unable to get Period. %s", err.Error())
	}

	if closedPeriod.UUID != "" {
		return resp, fmt.Errorf("attendance period is closed")
	}

	reimbursement := entity.Reimbursement{
		UUID:        uuid.New().String(),
		UserUUID:    req.UserUUID,
		Amount:      req.Amount,
		Date:        now,
		Description: req.Description,
		IP:          req.IP,
		CreatedBy:   req.UserUUID,
		UpdatedBy:   req.UserUUID,
		CreatedAt:   now,
		UpdatedAt:   now,
	}

	newReimbursement, err := uc.reimbursementRepository.CreateReimbursement(ctx, reimbursement)
	if err != nil {
		return resp, fmt.Errorf("unable to create reimbursement. %s", err.Error())
	}

	auditLog := entity.AuditLog{
		UUID:       uuid.New().String(),
		UserUUID:   req.UserUUID,
		Action:     "create",
		Entity:     "reimbursements",
		EntityUUID: newReimbursement.UUID,
		IP:         req.IP,
		RequestID:  req.RequestID,
		CreatedAt:  now,
	}

	uc.auditLogRepository.CreateAuditLog(ctx, auditLog)

	resp = ReimbursementResponse{
		UUID:        newReimbursement.UUID,
		UserUUID:    newReimbursement.UserUUID,
		Amount:      newReimbursement.Amount,
		Date:        newReimbursement.Date,
		Description: newReimbursement.Description,
		IP:          newReimbursement.IP,
		CreatedBy:   newReimbursement.CreatedBy,
		UpdatedBy:   newReimbursement.UpdatedBy,
		CreatedAt:   newReimbursement.CreatedAt,
		UpdatedAt:   newReimbursement.UpdatedAt,
	}

	return resp, nil
}
