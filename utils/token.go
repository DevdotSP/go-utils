package utils

import (
	"errors"
	"log"
	"os"
	"sync"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// Error messages
var (
	ErrInvalidSigningMethod = errors.New("invalid signing method")
	ErrInvalidToken         = errors.New("invalid token")
	ErrTokenNotFound        = errors.New("token not found")
)

// TokenInfo holds expiration details
type TokenInfo struct {
	Expiration time.Time
}

// ActiveTokens stores valid tokens
var activeTokens = sync.Map{}

// Load secret key from environment variable
var jwtSecret = []byte(os.Getenv("JWT_SECRET"))

// GenerateJWT creates a JWT token and removes the old one if provided
func GenerateJWT(userID int, currentToken string) (string, error) {
	// Remove old token if exists
	if currentToken != "" {
		activeTokens.Delete(currentToken)
	}

	// Define claims
	expirationTime := time.Now().Add(24 * time.Hour)
	claims := jwt.MapClaims{
		"user_id": userID,
		"exp":     expirationTime.Unix(),
		"iat":     time.Now().Unix(),
	}

	// Create token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Sign with secret key
	signedToken, err := token.SignedString(jwtSecret)
	if err != nil {
		return "", err
	}

	// Store the token in activeTokens
	activeTokens.Store(signedToken, TokenInfo{Expiration: expirationTime})

	return signedToken, nil
}

// ValidateToken checks if a token is valid
func ValidateToken(token string) (jwt.MapClaims, error) {
	// Check if the token exists and is not expired
	if info, ok := activeTokens.Load(token); !ok {
		return nil, ErrTokenNotFound
	} else if info.(TokenInfo).Expiration.Before(time.Now()) {
		activeTokens.Delete(token)
		return nil, ErrTokenNotFound
	}

	// Parse token
	parsedToken, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, ErrInvalidSigningMethod
		}
		return jwtSecret, nil
	})

	if err != nil {
		return nil, err
	}

	// Extract claims
	if claims, ok := parsedToken.Claims.(jwt.MapClaims); ok && parsedToken.Valid {
		return claims, nil
	}

	return nil, ErrInvalidToken
}

// ExtractCustomerIDFromToken extracts the user_id from JWT
func ExtractCustomerIDFromToken(tokenString string) (uint, error) {
	claims, err := ValidateToken(tokenString)
	if err != nil {
		return 0, err
	}

	customerID, ok := claims["user_id"].(float64)
	if !ok {
		return 0, errors.New("invalid or missing user_id in token")
	}

	return uint(customerID), nil
}



// DeleteToken removes a token
func DeleteToken(token string) error {
	_, loaded := activeTokens.LoadAndDelete(token)
	if !loaded {
		return ErrTokenNotFound
	}
	log.Printf("Token %s has been removed.", token)
	return nil
}

// CleanupExpiredTokens periodically removes expired tokens
func CleanupExpiredTokens() {
	for {
		time.Sleep(10 * time.Minute) // Adjust as needed
		now := time.Now()

		activeTokens.Range(func(key, value interface{}) bool {
			tokenInfo := value.(TokenInfo)
			if tokenInfo.Expiration.Before(now) {
				activeTokens.Delete(key)
				log.Printf("Expired token %s removed.", key)
			}
			return true
		})
	}
}

// StartCleanupRoutine runs CleanupExpiredTokens in the background
func StartCleanupRoutine() {
	go CleanupExpiredTokens()
}
