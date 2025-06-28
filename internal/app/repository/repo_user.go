package repository

import (
	"context"
	"fmt"

	"github.com/alimasyhur/go-payroll-service/internal/app/entity"
	"github.com/alimasyhur/go-payroll-service/internal/pkg/logger"
	"gorm.io/gorm"
)

type User interface {
	GetOneByUUID(ctx context.Context, uuid string) (result entity.User, err error)
	GetOneByUsername(ctx context.Context, username string) (result entity.User, err error)
	GetListByRole(ctx context.Context, role string) (results []entity.User, err error)
	CreateUser(ctx context.Context, payload entity.User) (result entity.User, err error)
	CreateEmployeeSalary(ctx context.Context, payload entity.EmployeeSalary) (result entity.EmployeeSalary, err error)
}

type user struct {
	tableName string
	db        *gorm.DB
}

func NewUserRepository(db *gorm.DB) User {
	if db == nil {
		panic("db is nil")
	}

	return &user{
		tableName: "users",
		db:        db,
	}
}

func (r *user) GetOneByUUID(ctx context.Context, uuid string) (result entity.User, err error) {
	err = r.db.WithContext(ctx).Table(r.tableName).Select(
		"uuid",
		"username",
		"password",
		"role",
	).
		Where("uuid = ?", uuid).
		Take(&result).Error

	if err != nil {
		logger.Log.Error(ctx, err.Error(), "UserRepository.GetOneByUUID", map[string]interface{}{
			"uuid": uuid,
		})
		return
	}

	return
}

func (r *user) GetListByRole(ctx context.Context, role string) (results []entity.User, err error) {
	err = r.db.WithContext(ctx).Table(r.tableName).Select(
		"uuid",
		"username",
		"role",
	).
		Where("role = ?", role).
		Find(&results).Error

	if err != nil {
		logger.Log.Error(ctx, err.Error(), "UserRepository.GetListByRole", map[string]interface{}{
			"role": role,
		})
		return
	}

	return
}

func (r *user) GetOneByUsername(ctx context.Context, username string) (result entity.User, err error) {
	err = r.db.WithContext(ctx).Table(r.tableName).Select(
		"uuid",
		"username",
		"password",
		"role",
	).
		Where("username = ?", username).
		Take(&result).Error

	if err != nil {
		logger.Log.Error(ctx, err.Error(), "UserRepository.GetOneByUsername", map[string]interface{}{
			"username": username,
		})
		return
	}

	return
}

func (r *user) CreateUser(ctx context.Context, payload entity.User) (entity.User, error) {
	err := r.db.WithContext(ctx).Table(r.tableName).Create(&payload).Error
	if err != nil {
		logger.Log.Error(ctx, fmt.Sprintf("UserRepository.CreateUser error: %s", err.Error()), payload)
		err = fmt.Errorf("UserRepository.CreateUser error: %s", err.Error())
		return entity.User{}, err
	}

	return payload, nil
}

func (r *user) CreateEmployeeSalary(ctx context.Context, payload entity.EmployeeSalary) (entity.EmployeeSalary, error) {
	err := r.db.WithContext(ctx).Table("employee_salaries").Create(&payload).Error
	if err != nil {
		logger.Log.Error(ctx, fmt.Sprintf("UserRepository.CreateEmployeeSalary error: %s", err.Error()), payload)
		err = fmt.Errorf("UserRepository.CreateEmployeeSalary error: %s", err.Error())
		return entity.EmployeeSalary{}, err
	}

	return payload, nil
}
