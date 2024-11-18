package mail

import (
	"qezde/notification/pkg/errors"
	"qezde/notification/pkg/utils/email"
	"regexp"
)

type WelcomeMailRequest struct {
	Email string `json:"email" binding:"required"`
	Code  string `json:"code"  binding:"required"`
}

func (r *WelcomeMailRequest) Validate() errors.Error {
	if r.Email == "" || !email.ValidateEmail(r.Email) {
		return errors.New("WELCOME_EMAIL_REQUEST_VALIDATION_ERROR", "invalid email format or invalid format", errors.TagBadRequest)
	}

	if r.Code == "" || len(r.Code) != 6 {
		return errors.New("WELCOME_EMAIL_REQUEST_VALIDATION_ERROR", "invalid code length", errors.TagBadRequest)
	}

	if matched, _ := regexp.MatchString("^[0-9]+$", r.Code); !matched {
		return errors.New("WELCOME_EMAIL_REQUEST_VALIDATION_ERROR", "code should only contain digits", errors.TagBadRequest)
	}

	return errors.Nil
}
