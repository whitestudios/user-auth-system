package auth

import (
	"fmt"
	"regexp"
)

type UserReqDto struct {
	Email    string
	Password string
}

func (u UserReqDto) EmailIsValid() bool {
	emailRegex := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	re := regexp.MustCompile(emailRegex)
	return re.MatchString(u.Email)
}

func (u UserReqDto) PasswordIsValid() bool {
	length := len(u.Password)

	switch {
	case length < 8:
		return false
	case length > 50:
		return false
	default:
		return true
	}
}

func (u UserReqDto) Validate() error {
	if !u.EmailIsValid() {
		return fmt.Errorf("Invalid email")
	}

	if len(u.Password) < 8 || len(u.Password) > 150 {
		return fmt.Errorf("Invalid password, the password field must be between 8 and 150 chars")
	}

	return nil
}
