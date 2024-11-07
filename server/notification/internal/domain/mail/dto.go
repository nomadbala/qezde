package mail

import (
	"errors"
	"qezde/notification/pkg/utils"
)

type WelcomeMailRequest struct {
	Email string `json:"email" binding:"required"`
	Code  string `json:"code"  binding:"required"`
}

func (r *WelcomeMailRequest) Validate() error {
	if r.Email == "" || !utils.ValidateEmail(r.Email) {
		return errors.New("invalid email format or invalid email")
	}

	if r.Code == "" || len(r.Code) != 6 {
		return errors.New("invalid code length")
	}

	return nil
}
