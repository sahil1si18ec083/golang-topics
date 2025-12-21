package utils

import (
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type UserClaims struct {
	Email string `json:"email"`
	ID    int64  `json:"id"`
	jwt.RegisteredClaims
}

func GenerateJWT(email string, userId int64) (string, error) {
	appEnv := os.Getenv("JWT_SECRET_KEY")
	fmt.Println(appEnv, "appenv")
	if appEnv == "" {
		fmt.Println("chal ")
	}

	claims := &UserClaims{
		Email: email,
		ID:    userId,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte("secretKey"))
	if err != nil {
		return "", err
	}
	return tokenString, nil

}
