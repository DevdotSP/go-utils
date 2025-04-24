package utils

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"log"
	"math/big"
	"os"
	"regexp"
	"strings"
	"time"

	"github.com/gofiber/fiber/v3"
	"github.com/joho/godotenv"
	"golang.org/x/crypto/bcrypt"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)


const (
	lowerChars   = "abcdefghijklmnopqrstuvwxyz"
	upperChars   = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	digitChars   = "0123456789"
	specialChars = "!@#$%^&*"
	allChars     = lowerChars + upperChars + digitChars + specialChars
)

var (
	slugRegex = regexp.MustCompile(`[^a-z0-9]+`)
	emailRegex = regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,}$`)
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

func GenerateTokenExpiry() time.Time {
	return time.Now().Add(5 * time.Minute) // Expires in 5 months
}

// LoadEnv loads environment variables from a .env file.
// It logs an error and stops execution if it can't load the .env file.
func LoadEnv() error {
	// Load the environment variables from the .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
		return fmt.Errorf("Error loading .env file")
	}
	log.Println("Environment variables loaded successfully")

	// Optionally, check if any required variables are missing
	return nil
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
// Helper function to extract the JWT from the Authorization header
func ExtractToken(ctx fiber.Ctx) (string, error) {
	token := ctx.Get("Authorization")
	if token == "" || len(token) < 7 || token[:7] != "Bearer " {
		return "", fiber.NewError(fiber.StatusUnauthorized, "No token provided")
	}
	return token[7:], nil
}




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


func GenerateUUID() string {
	b := make([]byte, 16)
	rand.Read(b)
	// Set version (4) and variant bits
	b[6] = (b[6] & 0x0f) | 0x40
	b[8] = (b[8] & 0x3f) | 0x80
	return fmt.Sprintf("%08x-%04x-%04x-%04x-%012x",
		b[0:4], b[4:6], b[6:8], b[8:10], b[10:])
}


func IsValidEmail(email string) bool {
	return emailRegex.MatchString(email)
}

// Invalid password, should have at least 8 characters long, a mix of uppercase and lowercase letters and at least one special character (@ or .)
// Validate password
func IsPasswordValid(password string) bool {
	hasEightLen := false
	hasUpperChar := false
	hasLowerChar := false
	hasSpecialChar := false
	if len(password) >= 8 {
		hasEightLen = true
	}

	upperString := regexp.MustCompile(`[A-Z]`)
	lowerString := regexp.MustCompile(`[a-z]`)
	specialString := regexp.MustCompile(`[!@#$%^&*(.)]`)

	hasUpperChar = upperString.MatchString(password)
	hasLowerChar = lowerString.MatchString(password)
	hasSpecialChar = specialString.MatchString(password)

	return hasEightLen && hasUpperChar && hasLowerChar && hasSpecialChar
}

func CurrentTimestamp() string {
	return time.Now().Format("2006-01-02 15:04:05")
}


func SlugifyString(input string) string {
	s := strings.ToLower(input)
	s = slugRegex.ReplaceAllString(s, "-")
	return strings.Trim(s, "-")
}

func IsUniqueConstraintError(err error) bool {
	// Check if error contains specific keywords indicating a unique constraint violation
	return err != nil && (strings.Contains(err.Error(), "unique constraint"))
}

func ToTitleCase(input string) string {
	caser := cases.Title(language.English)
	return caser.String(strings.ToLower(input))
}

func Print[T any](val T) {
	fmt.Println(val)
}
