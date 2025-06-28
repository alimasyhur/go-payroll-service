package reimbursement

import (
	"context"

	"github.com/alimasyhur/go-payroll-service/internal/app/repository"
)

type ReimbursementUsecase interface {
	CreateReimbursement(ctx context.Context, req ReimbursementRequest) (resp ReimbursementResponse, err error)
}
type usecase struct {
	attendancePeriodRepository repository.AttendancePeriod
	reimbursementRepository    repository.Reimbursement
	auditLogRepository         repository.AuditLog
}

func NewUsecase() *usecase {
	return &usecase{}
}

func (uc *usecase) SetAttendancePeriodRepository(r repository.AttendancePeriod) *usecase {
	uc.attendancePeriodRepository = r
	return uc
}

func (uc *usecase) SetReimbursementRepository(r repository.Reimbursement) *usecase {
	uc.reimbursementRepository = r
	return uc
}

func (uc *usecase) SetAuditLogRepository(r repository.AuditLog) *usecase {
	uc.auditLogRepository = r
	return uc
}

func (uc *usecase) Validate() ReimbursementUsecase {
	if uc.attendancePeriodRepository == nil {
		panic("attendance period repository is nil")
	}
	if uc.reimbursementRepository == nil {
		panic("reimbursement repository is nil")
	}
	if uc.auditLogRepository == nil {
		panic("audit log repository is nil")
	}

	return uc
}
