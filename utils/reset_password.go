package utils

import "fmt"

// ResetPassword handles password hashing and expiry update without checking old password
func ResetPassword(model PasswordUpdatable, newPwd string) error {
	if newPwd == "" {
		return fmt.Errorf("new password is required")
	}

	hashedPwd, err := HashData(newPwd)
	if err != nil {
		return fmt.Errorf("failed to hash password: %w", err)
	}

	model.SetPassword(hashedPwd)
	model.SetPwdExpiredDate(GeneratePasswordExpiry())

	return nil
}