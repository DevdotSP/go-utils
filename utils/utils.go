package utils

import (
	"crypto/rand"
	"encoding/hex"
	"log"
	"math/big"
	"os"
	"time"

	"github.com/gofiber/fiber/v3"
	"golang.org/x/crypto/bcrypt"
)

// HashPassword encrypts the password using bcrypt
func HashData(data string) (string, error) {
	hashedBytes, err := bcrypt.GenerateFromPassword([]byte(data), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedBytes), nil
}

// ComparePassword compares a hashed password with a plain text password
func CompareData(hashedData, plainData string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedData), []byte(plainData))
	return err == nil
}

// GeneratePasswordExpiry sets password expiry (default: 90 days)
func GeneratePasswordExpiry() time.Time {
	return time.Now().AddDate(0, 3, 0) // Expires in 3 months
}

// GetEnv returns the value of the environment variable or a default value.
func GetEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}

func GetResponseTime(c fiber.Ctx) string {
	connTime, ok := c.Locals("connTime").(time.Time)
	if !ok {
		connTime = time.Now()
	}
	return connTime.Format(time.DateTime)
}


// generateVerificationToken creates a random verification token
func GenerateVerificationToken() string {
	b := make([]byte, 16) // 16 bytes for a 32-character hex string
	rand.Read(b)
	return hex.EncodeToString(b)
}


// Helper function to extract the JWT from the Authorization header
func ExtractToken(headers map[string]string) (string, error) {
	token := headers["Authorization"]
	if token == "" || len(token) < 7 || token[:7] != "Bearer " {
		return "", fiber.NewError(fiber.StatusUnauthorized, "No token provided")
	}
	return token[7:], nil
}

const (
	lowerChars   = "abcdefghijklmnopqrstuvwxyz"
	upperChars   = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	digitChars   = "0123456789"
	specialChars = "!@#$%^&*"
	allChars     = lowerChars + upperChars + digitChars + specialChars
)

func GenerateRandomPassword() string {
	length := 6 + randInt(3) // Random length between 6 and 8
	password := make([]byte, length)

	// Ensure at least one character from each category
	password[0] = lowerChars[randInt(len(lowerChars))]
	password[1] = upperChars[randInt(len(upperChars))]
	password[2] = digitChars[randInt(len(digitChars))]
	password[3] = specialChars[randInt(len(specialChars))]

	// Fill the remaining with random characters
	for i := 4; i < length; i++ {
		password[i] = allChars[randInt(len(allChars))]
	}

	// Shuffle the password
	shuffle(password)

	return string(password)
}

func randInt(max int) int {
	n, err := rand.Int(rand.Reader, big.NewInt(int64(max)))
	if err != nil {
		log.Fatal(err)
	}
	return int(n.Int64())
}

func shuffle(data []byte) {
	for i := range data {
		j := randInt(len(data))
		data[i], data[j] = data[j], data[i]
	}
}