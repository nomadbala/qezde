package service

import (
	"bytes"
	"fmt"
	"github.com/resend/resend-go/v2"
	"html/template"
	"os"
	"path/filepath"
	"qezde/notification/pkg/errors"
)

const (
	BaseEmailTemplate    = "base_email.html"
	WelcomeEmailTemplate = "welcome_email.html"
)

var (
	ErrLoadingTemplates = errors.New("RESEND_EMAIL_SERVICE_ERROR", fmt.Sprintf("failed on loading templates from %s", "pkg/resend/templates"))
	ErrLoadingTemplate  = errors.New("RESEND_EMAIL_SERVICE_ERROR", fmt.Sprintf("failed on loading %s email template", WelcomeEmailTemplate))
	ErrParsingTemplate  = errors.New("RESEND_EMAIL_SERVICE_ERROR", fmt.Sprintf("failed on parsing %s email template", WelcomeEmailTemplate))
	ErrFillingTemplate  = errors.New("RESEND_EMAIL_SERVICE_ERROR", fmt.Sprintf("failed on filling %s email template", WelcomeEmailTemplate))
	ErrSendingEmail     = errors.New("RESEND_EMAIL_SERVICE_ERROR", "failed on sending email")
)

func (s *Service) LoadMailTemplate(filename string) (string, errors.Error) {
	path := filepath.Join("pkg/resend/templates", filename)
	data, err := os.ReadFile(path)

	if err != nil {
		return "", ErrLoadingTemplates
	}

	return string(data), errors.Nil
}

func (s *Service) SendWelcomeEmail(receiver, code string) errors.Error {
	templateData, err := s.LoadMailTemplate(WelcomeEmailTemplate)
	if err != errors.Nil {
		return ErrLoadingTemplate
	}

	parsedTemplate, parsErr := template.New("email").Parse(templateData)
	if parsErr != nil {
		return ErrParsingTemplate
	}

	var body bytes.Buffer
	data := struct {
		Code string
	}{Code: code}

	if err := parsedTemplate.Execute(&body, data); err != nil {
		return ErrFillingTemplate
	}

	//attachment := &resend.Attachment{
	//	Path:     "data:image/svg+xml;base64,PHN2ZyB3aWR0aD0iMTc5IiBoZWlnaHQ9IjE4NSIgdmlld0JveD0iMCAwIDE3OSAxODUiIGZpbGw9Im5vbmUiIHhtbG5zPSJodHRwOi8vd3d3LnczLm9yZy8yMDAwL3N2ZyI+CjxwYXRoIGZpbGwtcnVsZT0iZXZlbm9kZCIgY2xpcC1ydWxlPSJldmVub2RkIiBkPSJNMTM1LjQyNCAxODMuMDRMOTYuNTEyIDE0OS4yNDhDODkuMzQ0IDE1MC40NDMgODIuNjAyNyAxNTEuMDQgNzYuMjg4IDE1MS4wNEM1MC44NTg3IDE1MS4wNCAzMS43NDQgMTQ1LjA2NyAxOC45NDQgMTMzLjEyQzYuMzE0NjcgMTIxLjAwMyAwIDEwMS42MzIgMCA3NS4wMDhDMCA0OC44OTYgNi45OTczMyAyOS44NjY3IDIwLjk5MiAxNy45MkMzNC45ODY3IDUuOTczMzMgNTMuNTA0IDAgNzYuNTQ0IDBDMTI1LjAxMyAwIDE0OS4yNDggMjQuMTQ5MyAxNDkuMjQ4IDcyLjQ0OEMxNDkuMjQ4IDk4LjkwMTMgMTQyLjU5MiAxMTguODY5IDEyOS4yOCAxMzIuMzUyTDE1Ny42OTYgMTUyLjA2NEwxMzUuNDI0IDE4My4wNFpNNzUuMjY0IDEyMC44MzJDODAuMDQyNyAxMjAuODMyIDg0LjA1MzMgMTIwLjA2NCA4Ny4yOTYgMTE4LjUyOEM5MC41Mzg3IDExNi44MjEgOTMuMDk4NyAxMTMuOTIgOTQuOTc2IDEwOS44MjRDOTguNTYgMTAyLjgyNyAxMDAuMzUyIDkxLjM5MiAxMDAuMzUyIDc1LjUyQzEwMC4zNTIgNTkuOTg5MyA5OC41NiA0OC43MjUzIDk0Ljk3NiA0MS43MjhDOTEuMzkyIDM0LjU2IDg0LjgyMTMgMzAuODA1MyA3NS4yNjQgMzAuNDY0QzcxLjY4IDMwLjQ2NCA2Ny45MjUzIDMwLjg5MDcgNjQgMzEuNzQ0QzYwLjI0NTMgMzIuNDI2NyA1Ny44NTYgMzMuNDUwNyA1Ni44MzIgMzQuODE2QzU0Ljc4NCAzNy4yMDUzIDUzLjE2MjcgNDIuMzI1MyA1MS45NjggNTAuMTc2QzUwLjk0NCA1Ny44NTYgNTAuNDMyIDY2LjMwNCA1MC40MzIgNzUuNTJDNTAuNDMyIDg5Ljg1NiA1MS43OTczIDEwMC4zNTIgNTQuNTI4IDEwNy4wMDhDNTYuNDA1MyAxMTEuOTU3IDU5LjA1MDcgMTE1LjU0MSA2Mi40NjQgMTE3Ljc2QzY1Ljg3NzMgMTE5LjgwOCA3MC4xNDQgMTIwLjgzMiA3NS4yNjQgMTIwLjgzMlpNNzYuMzUyIDgzLjg0OEM3My4xNTIgODMuODA1MyA3MC4zMTQ3IDgzLjIwOCA2Ny44NCA4Mi4wNTZDNjcuODQgODAuMTM2IDY3LjczMzMgNzcuOTgxMyA2Ny41MiA3NS41OTJDNjcuMzQ5MyA3My4xNiA2Ny4wMjkzIDcwLjg3NzMgNjYuNTYgNjguNzQ0TDY3LjA3MiA2OC40ODhDNjkuNDE4NyA2OC45NTczIDcxLjYxNiA2OS4xOTIgNzMuNjY0IDY5LjE5MkM3NS4zMjggNjkuMTkyIDc2LjY1MDcgNjkuMDQyNyA3Ny42MzIgNjguNzQ0Qzc4LjYxMzMgNjguNDAyNyA3OS4xMDQgNjcuOTEyIDc5LjEwNCA2Ny4yNzJDNzkuMTA0IDY2LjYzMiA3OC41MjggNjYuMTYyNyA3Ny4zNzYgNjUuODY0Qzc2LjI2NjcgNjUuNTY1MyA3NC43OTQ3IDY1LjQxNiA3Mi45NiA2NS40MTZDNzEuODkzMyA2NS40MTYgNzAuMzM2IDY1LjUwMTMgNjguMjg4IDY1LjY3Mkw2Ni44OCA1Ny44NjRDNjcuODYxMyA1Ni45MjUzIDY4LjgyMTMgNTYuMzA2NyA2OS43NiA1Ni4wMDhDNzAuNzQxMyA1NS42NjY3IDcxLjc4NjcgNTUuNDk2IDcyLjg5NiA1NS40OTZDNzUuNDk4NyA1NS40OTYgNzcuNzE3MyA1Ni4wOTMzIDc5LjU1MiA1Ny4yODhDODEuNDI5MyA1OC40ODI3IDgyLjgzNzMgNjAuMDQgODMuNzc2IDYxLjk2Qzg0LjcxNDcgNjMuODggODUuMTg0IDY1Ljg4NTMgODUuMTg0IDY3Ljk3NkM4NS4xODQgNzEuMjE4NyA4NC40MzczIDczLjc3ODcgODIuOTQ0IDc1LjY1NkM4MS40OTMzIDc3LjQ5MDcgNzkuMzgxMyA3OC42IDc2LjYwOCA3OC45ODRMNzYuNjcyIDgzLjUyOEw3Ni4zNTIgODMuODQ4Wk03Mi40NDggOTcuMDk2QzcwLjk1NDcgOTcuMDk2IDY5Ljc2IDk2LjY0OCA2OC44NjQgOTUuNzUyQzY3Ljk2OCA5NC44MTMzIDY3LjUyIDkzLjU1NDcgNjcuNTIgOTEuOTc2QzY3LjUyIDkwLjYxMDcgNjcuOTY4IDg5LjQ1ODcgNjguODY0IDg4LjUyQzY5LjgwMjcgODcuNTgxMyA3MC45NTQ3IDg3LjExMiA3Mi4zMiA4Ny4xMTJDNzMuNzI4IDg3LjExMiA3NC45MDEzIDg3LjU4MTMgNzUuODQgODguNTJDNzYuODIxMyA4OS40NTg3IDc3LjMxMiA5MC42MTA3IDc3LjMxMiA5MS45NzZDNzcuMzEyIDkzLjU1NDcgNzYuODQyNyA5NC44MTMzIDc1LjkwNCA5NS43NTJDNzUuMDA4IDk2LjY0OCA3My44NTYgOTcuMDk2IDcyLjQ0OCA5Ny4wOTZaTTE2OC4zMiAxODQuNzEyQzE3My44NDMgMTg0LjcxMiAxNzguMzIgMTgwLjIzNSAxNzguMzIgMTc0LjcxMkMxNzguMzIgMTY5LjE4OSAxNzMuODQzIDE2NC43MTIgMTY4LjMyIDE2NC43MTJDMTYyLjc5NyAxNjQuNzEyIDE1OC4zMiAxNjkuMTg5IDE1OC4zMiAxNzQuNzEyQzE1OC4zMiAxODAuMjM1IDE2Mi43OTcgMTg0LjcxMiAxNjguMzIgMTg0LjcxMloiIGZpbGw9IndoaXRlIi8+Cjwvc3ZnPgo=",
	//	Filename: "logo.svg",
	//}

	params := &resend.SendEmailRequest{
		From:    s.Resend.Sender,
		To:      []string{receiver},
		Subject: s.Resend.Subject,
		Html:    body.String(),
		//Attachments: []*resend.Attachment{attachment},
	}

	if _, err := s.Resend.Client.Emails.Send(params); err != nil {
		return ErrSendingEmail
	}

	return errors.Nil
}

//func (s *Service) SendScheduledEmail(receiver, message, scheduledAt string) error {
//	templateData, err := s.LoadMailTemplate("base_email.html")
//	if err != errors.Nil {
//		return fmt.Errorf()
//	}
//
//	parsedTemplate, parseErr := template.New("email").Parse(templateData)
//	if parseErr != nil {
//		return parseErr
//	}
//
//	var body bytes.Buffer
//	data := struct {
//		Message string
//	}{Message: message}
//
//	if execErr := parsedTemplate.Execute(&body, data); execErr != nil {
//		return execErr
//	}
//
//	params := &resend.SendEmailRequest{
//		From:        s.Resend.Sender,
//		To:          []string{receiver},
//		Subject:     s.Resend.Subject,
//		Html:        body.String(),
//		ScheduledAt: scheduledAt,
//	}
//
//	_, sendErr := s.Resend.Client.Emails.Send(params)
//	return sendErr
//}
