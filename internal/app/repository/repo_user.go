package repository

import (
	"context"
	"fmt"

	"github.com/alimasyhur/go-payroll-service/internal/app/entity"
	"github.com/alimasyhur/go-payroll-service/internal/pkg/logger"
	"gorm.io/gorm"
)

type User interface {
	GetOneByUsername(ctx context.Context, username string) (result entity.User, err error)
	CreateUser(ctx context.Context, payload entity.User) (result entity.User, err error)
	CreateUserSalary(ctx context.Context, payload entity.UserSalary) (result entity.UserSalary, err error)
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

func (r *user) CreateUserSalary(ctx context.Context, payload entity.UserSalary) (entity.UserSalary, error) {
	err := r.db.WithContext(ctx).Table("user_salaries").Create(&payload).Error
	if err != nil {
		logger.Log.Error(ctx, fmt.Sprintf("UserRepository.CreateUserSalary error: %s", err.Error()), payload)
		err = fmt.Errorf("UserRepository.CreateUserSalary error: %s", err.Error())
		return entity.UserSalary{}, err
	}

	return payload, nil
}
