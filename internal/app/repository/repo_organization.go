package repository

import (
	"context"

	"gorm.io/gorm"

	"github.com/weanan/weanan-service/internal/app/entity"
)

type Organization interface {
	GetById(ctx context.Context, id uint) (entity entity.Organization, err error)
	GetAll(ctx context.Context) (entities []entity.Organization, err error)
	Create(ctx context.Context, entity *entity.Organization) (err error)
	Update(ctx context.Context, entity *entity.Organization) (err error)
	Delete(ctx context.Context, entity *entity.Organization) (err error)
}

type organization struct {
	db *gorm.DB
}

func NewOrganizationRepository(db *gorm.DB) Organization {
	if db == nil {
		panic("database is nil")
	}

	return &organization{db}
}

func (r *organization) GetById(ctx context.Context, id uint) (entity entity.Organization, err error) {
	err = r.db.WithContext(ctx).Last(&entity, id).Error
	return
}

func (r *organization) GetAll(ctx context.Context) (entities []entity.Organization, err error) {
	err = r.db.WithContext(ctx).Find(&entities).Error
	return
}

func (r *organization) Create(ctx context.Context, entity *entity.Organization) (err error) {
	err = r.db.WithContext(ctx).Create(entity).Error
	return
}

func (r *organization) Update(ctx context.Context, entity *entity.Organization) (err error) {
	err = r.db.WithContext(ctx).Save(entity).Error
	return
}

func (r *organization) Delete(ctx context.Context, entity *entity.Organization) (err error) {
	err = r.db.WithContext(ctx).Delete(entity).Error
	return
}
