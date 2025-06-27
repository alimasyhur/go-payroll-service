package user

import (
	"context"
	"log"
	"math/rand"
	"strconv"
	"time"

	"github.com/alimasyhur/go-payroll-service/internal/app/entity"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

func (uc *usecase) GenerateUserSeed(ctx context.Context) (err error) {
	password, _ := bcrypt.GenerateFromPassword([]byte("payroll123"), bcrypt.DefaultCost)

	admin := entity.User{
		UUID:     uuid.NewString(),
		Username: "admin",
		Password: string(password),
		Role:     "admin",
	}

	_, err = uc.userRepository.CreateUser(ctx, admin)

	for i := 1; i <= 100; i++ {
		user := entity.User{
			UUID:     uuid.New().String(),
			Username: "employee" + strconv.Itoa(i),
			Password: string(password),
			Role:     "employee",
		}
		if _, err = uc.userRepository.CreateUser(ctx, user); err != nil {
			log.Println("Failed to seed user:", err)
		}

		salary := entity.UserSalary{
			UUID:          uuid.New().String(),
			UserUUID:      user.UUID,
			Amount:        float64(rand.Intn(5_000_000) + 5_000_000),
			Active:        true,
			EffectiveDate: time.Now(),
			CreatedAt:     time.Now(),
		}
		if _, err = uc.userRepository.CreateUserSalary(ctx, salary); err != nil {
			log.Println("Failed to seed user salary:", err)
		}
	}

	return
}
