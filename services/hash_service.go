package services

import "golang.org/x/crypto/bcrypt"

func Hash(pass string) (string, error) {
	encoded, err := bcrypt.GenerateFromPassword([]byte(pass), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(encoded), nil
}

func Compare(pass, hash string) error {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(pass))
}
