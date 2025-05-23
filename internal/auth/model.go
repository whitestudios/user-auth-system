package auth

import (
	"fmt"
	"regexp"
)

type UserCreateDto struct {
	Email    string
	Password string
}

func (u UserCreateDto) EmailIsValid() bool {
	emailRegex := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	re := regexp.MustCompile(emailRegex)
	return re.MatchString(u.Email)
}

func (u UserCreateDto) PasswordIsValid() error {
	length := len(u.Password)

	switch {
	case length < 8:
		return fmt.Errorf("password too short: need at least 8 characters, got %d", length)
	case length > 50:
		return fmt.Errorf("password too long: maximum 50 characters allowed, got %d", length)
	default:
		return nil
	}
}

func (u UserCreateDto) Validate() error {
	if !u.EmailIsValid() {
		return fmt.Errorf("Invalid email")
	}

	if len(u.Password) < 8 || len(u.Password) > 150 {
		return fmt.Errorf("Invalid password, the password field must be between 8 and 150 chars")
	}

	return nil
}
