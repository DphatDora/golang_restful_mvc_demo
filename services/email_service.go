package services

import (
	"fmt"
	"os"
	"strconv"

	"gopkg.in/gomail.v2"
)

type EmailService interface {
	SendWelcomeEmail(to string, name string) error
}

type emailService struct {
	dialer *gomail.Dialer
}

func NewEmailService() (EmailService, error) {
	// Validate required environment variables
	requiredEnvVars := []string{"SMTP_HOST", "SMTP_USER", "SMTP_PASSWORD", "SMTP_FROM"}
	for _, envVar := range requiredEnvVars {
		if os.Getenv(envVar) == "" {
			return nil, fmt.Errorf("missing required environment variable: %s", envVar)
		}
	}

	// Get SMTP port from environment variable or use default
	port := 587
	if portStr := os.Getenv("SMTP_PORT"); portStr != "" {
		var err error
		port, err = strconv.Atoi(portStr)
		if err != nil {
			return nil, fmt.Errorf("invalid SMTP_PORT: %v", err)
		}
	}

	dialer := gomail.NewDialer(
		os.Getenv("SMTP_HOST"),
		port,
		os.Getenv("SMTP_USER"),
		os.Getenv("SMTP_PASSWORD"),
	)

	return &emailService{dialer: dialer}, nil
}

func (s *emailService) SendWelcomeEmail(to string, name string) error {
	if to == "" || name == "" {
		return fmt.Errorf("invalid email parameters: to=%s, name=%s", to, name)
	}

	m := gomail.NewMessage()
	m.SetHeader("From", os.Getenv("SMTP_FROM"))
	m.SetHeader("To", to)
	m.SetHeader("Subject", "Welcome to Our Service")
	m.SetBody("text/html", fmt.Sprintf(`
        <h1>Welcome %s!</h1>
        <p>Thank you for registering with our service.</p>
    `, name))

	return s.dialer.DialAndSend(m)
}
