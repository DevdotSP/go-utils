package utils

import (
	"fmt"
	"os"
	"strconv"

	"gopkg.in/gomail.v2"
)

func GoogleSendEmail(to, subject, link string) error {
	// Load SMTP configuration from environment variables
	from := os.Getenv("MAIL_FROM_ADDRESS")
	password := os.Getenv("MAIL_PASSWORD")
	smtpHost := os.Getenv("MAIL_HOST")
	smtpPortStr := os.Getenv("MAIL_PORT")
	username := os.Getenv("MAIL_USERNAME")

	// Parse SMTP port with fallback to 587
	port := 587
	if smtpPortStr != "" {
		if parsedPort, err := strconv.Atoi(smtpPortStr); err == nil {
			port = parsedPort
		}
	}

	// Compose HTML email
	htmlBody := fmt.Sprintf(`
		<html>
		<body>
			<p>Please verify your email by clicking the button below:</p>
			<a href="%s" style="display: inline-block; padding: 10px 20px; background-color: #4CAF50; color: white; text-align: center; text-decoration: none; border-radius: 5px;">
				Verify Email
			</a>
			<br><br>
			<p>Thank you,</p>
			<p>Support Team</p>
			<img src="https://storage.googleapis.com/test_store_rapa/images/1739771740_163.jpg" alt="Footer Banner" style="width: 100%%; max-width: 600px;">
		</body>
		</html>
	`, link)

	// Create new email
	m := gomail.NewMessage()
	m.SetHeader("From", from)
	m.SetHeader("To", to)
	m.SetHeader("Subject", subject)
	m.SetBody("text/html", htmlBody)

	// Send using SMTP dialer
	dialer := gomail.NewDialer(smtpHost, port, username, password)

	if err := dialer.DialAndSend(m); err != nil {
		return fmt.Errorf("❌ failed to send email to %s: %w", to, err)
	}

	fmt.Println("✅ Email sent successfully to:", to)
	return nil
}
