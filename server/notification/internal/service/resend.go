package service

import (
	"bytes"
	"github.com/resend/resend-go/v2"
	"html/template"
	"os"
	"path/filepath"
)

func (s *Service) LoadMailTemplate(filename string) (string, error) {
	path := filepath.Join("pkg/resend/templates", filename)
	data, err := os.ReadFile(path)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

func (s *Service) SendWelcomeEmail(receiver, code string) error {
	templateData, err := s.LoadMailTemplate("welcome_email.html")
	if err != nil {
		return err
	}

	parsedTemplate, parseErr := template.New("email").Parse(templateData)
	if parseErr != nil {
		return parseErr
	}

	var body bytes.Buffer
	data := struct {
		Code string
	}{Code: code}

	if execErr := parsedTemplate.Execute(&body, data); execErr != nil {
		return execErr
	}

	params := &resend.SendEmailRequest{
		From:    s.Resend.Sender,
		To:      []string{receiver},
		Subject: s.Resend.Subject,
		Html:    body.String(),
	}

	_, sendErr := s.Resend.Client.Emails.Send(params)
	return sendErr
}
