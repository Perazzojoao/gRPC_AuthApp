package util

import "golang.org/x/crypto/bcrypt"

func GenerateHash(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}

func CompareHash(password_1, password_2 string) error {
	return bcrypt.CompareHashAndPassword([]byte(password_1), []byte(password_2))
}
