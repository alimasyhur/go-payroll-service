package jwt

import (
	"os"
	"time"

	"github.com/alimasyhur/go-payroll-service/internal/app/entity"
	"github.com/golang-jwt/jwt/v5"
)

var jwtKey = []byte(os.Getenv("JWT_SECRET")) // set di .env

func GenerateJWT(user entity.User) (string, error) {
	claims := jwt.MapClaims{
		"user_uuid": user.UUID,
		"role":      user.Role,
		"exp":       time.Now().Add(24 * time.Hour).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtKey)
}
