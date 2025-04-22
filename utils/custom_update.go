package utils

import (
	"fmt"
	"gorm.io/gorm"
)

// UpdateRecord updates a record in the database using a transaction (tx).
func UpdateRecordTX[T any](tx *gorm.DB, model *T, column any, value any, updates map[string]interface{}) error {
	var valueStr string

	// Convert value to string or int for the query
	switch v := value.(type) {
	case string:
		valueStr = v
	case int:
		valueStr = fmt.Sprintf("%d", v)
	default:
		return fmt.Errorf("unsupported value type: %T", value)
	}

	// Convert column to string and ensure proper formatting for SQL query
	columnStr, ok := column.(string)
	if !ok {
		return fmt.Errorf("column must be a string, got %T", column)
	}

	// Attempt to find the record first
	result := tx.Model(model).Where(fmt.Sprintf("%s = ?", columnStr), valueStr).Updates(updates)
	if result.Error != nil {
		return fmt.Errorf("failed to update record: %w", result.Error)
	}

	// If no records were updated, insert a new one
	if result.RowsAffected == 0 {
		if err := tx.Create(model).Error; err != nil {
			return fmt.Errorf("failed to insert new record: %w", err)
		}
	}

	return nil
}




