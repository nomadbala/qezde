package service

import (
	"bytes"
	"fmt"
	"github.com/resend/resend-go/v2"
	"html/template"
	"os"
	"path/filepath"
	"qezde/notification/internal/domain/mail"
	"qezde/notification/pkg/errors"
)

func (s *Service) FetchEmailTemplate(filename string) (*template.Template, errors.Error) {
	s.templateCache.RLock()
	tmpl, exists := s.templateCache.templates[filename]
	s.templateCache.RUnlock()

	if exists {
		return tmpl, errors.Nil
	}

	templatePath := filepath.Join("pkg/resend/templates", filename)
	templateData, err := os.ReadFile(templatePath)
	if err != nil {
		return nil, errors.New("RESEND_EMAIL_SERVICE_ERROR", fmt.Sprintf("failed to load template %s from %s: %v", filename, "pkg/resend/templates", err), errors.TagInternalServerError)
	}

	parsedTemplate, err := template.New(filename).Parse(string(templateData))
	if err != nil {
		return nil, errors.New("RESEND_EMAIL_SERVICE_ERROR", fmt.Sprintf("failed to parse template %s: %v", filename, err), errors.TagInternalServerError)
	}

	s.templateCache.Lock()
	s.templateCache.templates[filename] = parsedTemplate
	s.templateCache.Unlock()

	return parsedTemplate, errors.Nil
}

func (s *Service) SendWelcomeEmail(request mail.WelcomeMailRequest) errors.Error {
	if err := request.Validate(); err != errors.Nil {
		return err
	}

	tmpl, err := s.FetchEmailTemplate("welcome_email.html")
	if err != errors.Nil {
		return err
	}

	var body bytes.Buffer
	if err := tmpl.Execute(&body, struct{ Code string }{Code: request.Code}); err != nil {
		return errors.New("RESEND_EMAIL_SERVICE_ERROR", fmt.Sprintf("failed to execute template %s: %v", "welcome_email.html", err), errors.TagInternalServerError)
	}

	params := &resend.SendEmailRequest{
		From:    s.Resend.Sender,
		To:      []string{request.Email},
		Subject: s.Resend.Subject,
		Html:    body.String(),
	}

	if _, err := s.Resend.Client.Emails.Send(params); err != nil {
		return errors.New("RESEND_EMAIL_SERVICE_ERROR", "failed on sending email", errors.TagInternalServerError)
	}

	return errors.Nil
}
