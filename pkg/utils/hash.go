package utils

import "golang.org/x/crypto/bcrypt"

func HashPassword(password string) (string, error) {
	//the higher the cost, the heavier it gets
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

// compare a password with a hash
func CheckPasswordHash(password string, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
