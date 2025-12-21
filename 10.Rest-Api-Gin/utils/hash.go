package utils

import "golang.org/x/crypto/bcrypt"

func HashPssword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hash), nil

}

func ComparePassword(hashedPasswordFromDB string, suppliedPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPasswordFromDB), []byte(suppliedPassword))

}
