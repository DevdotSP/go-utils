package utils

import (
	"fmt"
	"time"
)

// PasswordUpdatable defines required methods for a model to support password updates.
type PasswordUpdatable interface {
	GetPassword() string
	SetPassword(string)
	SetPwdExpiredDate(time.Time)
}

// TryUpdatePassword checks the old password, hashes the new one, and updates the model.
// It returns a bool indicating whether the password was changed.
func TryUpdatePassword(model PasswordUpdatable, oldPwd, newPwd string) (bool, error) {
	if oldPwd == "" || newPwd == "" {
		return false, nil // Nothing to do
	}

	if !CompareData(model.GetPassword(), oldPwd) {
		return false, fmt.Errorf("old password is incorrect")
	}

	hashedPwd, err := HashData(newPwd)
	if err != nil {
		return false, fmt.Errorf("password hashing failed: %w", err)
	}

	model.SetPassword(hashedPwd)
	model.SetPwdExpiredDate(GeneratePasswordExpiry())

	return true, nil
}
