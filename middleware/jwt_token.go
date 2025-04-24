package middleware

import (
	"github.com/gofiber/fiber/v3"
	"github.com/DevdotSP/go-utils/utils" // Update with your actual repo path
)

// JWTAuthMiddleware checks the Authorization header for a valid JWT token
func JWTAuthMiddleware(c fiber.Ctx) error {
	// Extract the token using the helper function
	token, err := utils.ExtractToken(c)
	if err != nil {
		return err // Return the error to respond with an unauthorized status
	}

	// Validate the token using the ValidateToken function
	claims, err := utils.ValidateToken(token)
	if err != nil {
		return fiber.NewError(fiber.StatusUnauthorized, "Invalid or expired token")
	}

	// Optionally, store claims in locals for access in the next handlers
	c.Locals("claims", claims)

	// Proceed to the next handler
	return c.Next()
}
