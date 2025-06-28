package user

import (
	"context"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

func (uc *usecase) Login(ctx context.Context, req LoginRequest) (resp LoginResponse, err error) {
	user, err := uc.userRepository.GetOneByUsername(ctx, req.Username)
	if err != nil {
		return resp, fmt.Errorf("invalid username or password. %s", err.Error())
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		return resp, fmt.Errorf("invalid credentials. %s", err.Error())
	}

	claims := jwt.MapClaims{
		"user_uuid": user.UUID,
		"role":      user.Role,
		"exp":       time.Now().Add(24 * time.Hour).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	jwtSecret := uc.config.JwtSecret
	signedToken, err := token.SignedString([]byte(jwtSecret))
	if err != nil {
		return resp, fmt.Errorf("fail to generate token")
	}

	resp = LoginResponse{
		Token: signedToken,
	}

	return
}
