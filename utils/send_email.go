package utils

import (
	"fmt"
	"os"
	"strconv"

	"gopkg.in/gomail.v2"
)

// EmailType constants
const (
	VerifyEmail     = "verify"
	ForgotPassword  = "forgot"
)

// GoogleSendEmail sends HTML email for verification or forgot password
func GoogleSendEmail(to, subject, link, emailType string) error {
	from := os.Getenv("MAIL_FROM_ADDRESS")
	password := os.Getenv("MAIL_PASSWORD")
	smtpHost := os.Getenv("MAIL_HOST")
	smtpPortStr := os.Getenv("MAIL_PORT")
	username := os.Getenv("MAIL_USERNAME")

	// Default port
	port := 587
	if smtpPortStr != "" {
		if parsedPort, err := strconv.Atoi(smtpPortStr); err == nil {
			port = parsedPort
		}
	}

	// Generate email body based on type
	htmlBody := generateEmailBody(emailType, link)

	m := gomail.NewMessage()
	m.SetHeader("From", from)
	m.SetHeader("To", to)
	m.SetHeader("Subject", subject)
	m.SetBody("text/html", htmlBody)

	dialer := gomail.NewDialer(smtpHost, port, username, password)

	if err := dialer.DialAndSend(m); err != nil {
		return fmt.Errorf("❌ failed to send email to %s: %w", to, err)
	}

	fmt.Println("✅ Email sent successfully to:", to)
	return nil
}

// generateEmailBody returns a formatted HTML string depending on the email type
func generateEmailBody(emailType, link string) string {
	var actionText, message string

	switch emailType {
	case ForgotPassword:
		actionText = "Reset Password"
		message = "You requested to reset your password. Click the button below to proceed:"
	case VerifyEmail:
		actionText = "Verify Email"
		message = "Please verify your email by clicking the button below:"
	default:
		actionText = "Open Link"
		message = "Click the button below:"
	}

	return fmt.Sprintf(`
		<html>
		<body>
			<p>%s</p>
			<a href="%s" style="display: inline-block; padding: 10px 20px; background-color: #4CAF50; color: white; text-align: center; text-decoration: none; border-radius: 5px;">
				%s
			</a>
			<br><br>
			<p>If you did not request this, you can safely ignore this email.</p>
			<p>Thank you,<br>Support Team</p>
			<img src="https://storage.googleapis.com/test_store_rapa/images/1739771740_163.jpg" alt="Footer Banner" style="width: 100%%; max-width: 600px;">
		</body>
		</html>
	`, message, link, actionText)
}
