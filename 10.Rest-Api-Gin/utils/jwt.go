package utils

import (
	"errors"
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
	secretkey := os.Getenv("JWT_SECRET_KEY")
	if secretkey == "" {
		return "", errors.New("JWT_SECRET_KEY not set in environment")
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
	tokenString, err := token.SignedString([]byte(secretkey))
	if err != nil {
		return "", err
	}
	return tokenString, nil

}
func VerifyToken(tokenString string) (*UserClaims, error) {
	secretKey := os.Getenv("JWT_SECRET_KEY")
	if secretKey == "" {
		return nil, errors.New("JWT_SECRET_KEY not set")
	}

	token, err := jwt.ParseWithClaims(
		tokenString,
		&UserClaims{},
		func(token *jwt.Token) (interface{}, error) {

			// 🔐 Enforce HS256
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}

			return []byte(secretKey), nil
		},
	)

	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*UserClaims)

	if !ok || !token.Valid {
		return nil, errors.New("invalid token")
	}
	userId := claims.ID
	fmt.Println(userId, "hello")

	return claims, nil
}
